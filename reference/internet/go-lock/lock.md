# Go 語言互斥鎖設計實現

## Mutex 介紹

- Lock(): 多次上鎖會 `panic()`
- Unlock(): 多次解鎖會 `panic()`
- tryLock(): 1.18新增, 非租塞式的取鎖操作

Mutex的結構

```go
type Mutex struct {
    state int32 // 表示互斥鎖的狀態, 複合型字段
    sema uint32 // 信號量變量, 用來控制等待 goroutine 的阻塞休眠和喚醒
}

const (
    mutexLocked =1 << iota // 表示互斥鎖的鎖定狀態
    mutexWoken // 表示從正常模式被喚醒
    mutexStraving // 當前的互斥鎖進入飢餓狀態
    mutexWaiterShift = iota // 當前互斥鎖上等待者的數量
)
```

## Lock 加鎖

```go
func (m *Mutex) Lock() {
    // 判斷當前鎖的狀態, 如果鎖是完全空閒的, 即 m.state 為 0, 將 m.state 的值賦為 1
    if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
        if race.Enabled() {
            race.Acquire(unsafe.Pointer(m))
        }
        return
    }
    m.lockSlow()
}
```

上面的代碼主要兩部分邏輯:

- 通過`CAS`判斷當前鎖的狀態, 也就是`state`字段的低 1 位, 如果鎖是完全空閒的, 即 m.state 為 0, 則對其加鎖, 將 m.state 的值賦為 1.
- 若當前鎖已經被其他`goroutine`加鎖, 則進行`lockSlow`方法嘗試通過自旋或飢餓狀態下飢餓`goroutine`競爭方式等待鎖的釋放, 我們在下面介紹 lockSlow 方法:

lockSlow 代碼段有點長, 主體是一個`for`循環, 其主要邏輯可以分為以下三部分:

- 狀態初始化
- 判斷是否符合自旋條件, 符合條件進行自旋操作
- 搶鎖準備期望狀態
- 通過`CAS`操作更新期望狀態

### 初始化狀態

```go
func (m *Mutex) lockSlow() {
    var waitStartTime int64 // 用來計算 waiter 的等待時間
    starving := false // 飢餓模式標誌, 如果等待市場超過 1ms, starving 設置為 true, 後續操作會把 Mutex 也標記為飢餓狀態
    awoke := false // awoke 表示攜程是否喚醒, 當 goroutine 在自旋時, 相當於 CPU 上已經有在等鎖的攜程. 為避免 Mutex 解鎖時要在喚醒狀態, Mutex 處于喚醒狀態後, 要把本攜程的 awoke 也設置為 TRUE.
    iter := 0 // 用於記錄攜程的自旋次數
    old := m.state // 記錄當前鎖的狀態
    // ...
}
```

### 自旋

```go
for {
    // 判斷是否允許進入自旋 兩個條件
    // 條件 1: 當前鎖不能處於解狀態
    // 條件 2: 在 runtime_canSpin 內實現, 邏輯是在多核 CPU 運行, 自旋的次數小於 4
    if old & (mutexLocked | mutexStarving) == mutexLocked && runtime_canSpin(iter) {
        // !awoke 判斷當前 goroutine 不是在喚醒狀態
        // old & mutexWoken == 0: 表示沒有其他正在喚醒的 goroutine
        // old >> mutexWaiterShift != 0: 表示等待隊列中有正在等待的 goroutine
        // atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken): 嘗試將當前鎖的低 2 位的 Woken 狀態位設置為 1, 表示已經被喚醒, 這是為了通知在解鎖 Unlock()中不要在喚醒其他的 waiter 了.
        if !awoke && old & mutexWoken == 0 && old >> mutexWaiterShift != 0 && atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken)  {
            awoke = true // 設置當前 goroutine 喚醒成功
        }
    }
    // 進行自旋
    runtime_doSpin()
    // 自選次數
    iter++
    // 記錄當前鎖的狀態
    old = m.state
    continue
}
```

我們想讓當前`goroutine`進入自旋的原因是我們樂觀的認為`當前正在持有鎖的 goroutine 能在較短的時間內歸還鎖`, 所以我們需要一些條件來判斷, mutex 的判斷條件我們在文字描述下:

- `old & (mutexLocked|mutexStarving) == mutexLocked`用來判斷鎖是否處於正常模式且加鎖, 為什麼要這麼判斷呢?
  - `mutexLocked` 二進制表示為 `0001`
  - `mutexStarving` 二進制表示為 `0100`
  - `mutexLocked|mutexStarving` 二進制則為 `0101`
  - 如果當前處於飢餓模式, 低三位一定會是 1, 如果當前處於加鎖模式, 低一位一定會是 1
  - runtime_canSpin() 方法用來判斷是否符合自旋條件

`runtime_canSpin`:

```go
const active_spin = 4
func sync_runtime_canSpin(i int) bool {
    // i >= active_spin: 自旋次數 超過 4 , 退出
    // ncpu <= 1: 單核 CPU, 退出
    // gmaxprocs <= 1+xx: GOMAXPROC <=1 , 退出
    if i >= active_spin || ncpu <= 1 || gomaxprocs <= int32(sched.npidle+sched.nmspining) +1 {
        return false
    }
    // 沒有處理器 p, 或者處理器 p 的運行隊列不為空, 退出
    if p := getg().m.p.ptr(); !runqempty(p){
        return false
    }
    return true
}
```

`runtime_doSpin`

```go
const active_spin_cnt = 30 // 循環次數 30 次
func sync_runtime_doSpin() {
    procyield(active_spin_cnt)
}
```

```asm
TEXT runtime.procyield(SB),NOSPLIT, $0-0
    MOVL cycles+0(FP), AX
again:
    PAUSE
    SUBL $1, AX
    JNZ again
    RET
```

循環次數被設置為`30`次, 自旋操作就是執行 30 次`PAUSE`指令, 通過該指令佔用 CPU 並消費 CPU 時間, 進行忙等待.

`自旋`就是為了優化`等待組賽->喚醒->參與搶佔鎖`這個過程的不高效, 期望程序在`自旋`過程中鎖能被釋放掉

### 搶鎖準備期望

自旋邏輯處理好後開始根據上下文計算當前互斥鎖最新的狀態, 根據不同條件來計算 mutexLocked, mutexStarving, mutexWoken, mutexWaiterShift.

計算 mutexLocked 的值:

```go
{
    // ...

    // 基於 old 狀態聲明到一個新狀態
    new := old
    // 新狀態處於非飢餓的條件下才可加鎖
    if old & mutexStarving == 0 {
        new |= mutexLocked
    }
    // ...
}
```

計算 mutexWaiterShift 的值:

```go
// 如果 old 已經處於加鎖或者飢餓狀態, 則等待者按照 FIFO 的順序排隊
if old & (mutexLocked | mutexStarving) != 0 {
    new += 1 << mutexWaiterShift
}
```

計算 mutexStarving 的值:

```go
if starving && old & mutexLocked != 0 {
    new != mutexStarving
}
```

計算 mutexWoken 的值:

```go
// 當前 goroutine 的 waither 被喚醒, 則重置 flag
if awoke {
    // 喚醒狀態不一致, 直接拋出異常 
    if new & mutexWoken == 0 {
        panic("sync: inconsisten mutex state")
    }
    // 新狀態清除喚醒標記, 因為後面的 goroutine 只會阻塞或者搶鎖成功
    // 如果是掛起狀態, 那就需要等待其他釋放鎖的 goroutine 來喚醒.
    // 假如其他 goroutine 在 unlock 的時候發現 Woken 的位置不是 0, 則就不回去喚醒, 那該 goroutine 就無法在被喚醒後加鎖 
    new &^= mutexWoken
}
```

### 通過 CAS 操作更新期望狀態

上面我們已經得到了鎖的期望狀態, 接下來通過 CAS 將鎖的狀態進行更新

```go
// 嘗試將鎖的狀態更新到最新狀態
if atomic.CompareAndSwapInt32(&m.state, old, new) {
    // 如果原來鎖的狀態是沒有加鎖的並且不處於飢餓狀態, 則表示當前 goroutine 已經獲取到鎖了, 直接退出即可
    if old & (mutexLocked | mutexStarving) == 0 {
        break // locked and mutex with CAS
    }
    // 到這裡就表示 goroutine 還沒有獲取到鎖, waitStartTime 是 goroutine 開始等待的時間, waitStartTime != 0 就表示當前 goroutine 已經等待過了, 則需要將其放置在等待隊列的隊頭, 否則就排到隊列隊尾
    queueLifo := waitStartTime != 0 
    if waitStartTime == 0 {
        waitStartTime = runtime_nanotime()
    }
    // 阻塞等待
    runtime_SemacquireMutex(&m.sema, queueLifo, 1)
    // 被信號量喚醒後檢查當前 goroutine 是否應該表示為飢餓
    // 1. 當前 goroutine 已經飢餓
    // 2. goroutine 已經等待了 1ms 以上
    starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
    // 再次獲取當前鎖的狀態
    old = m.state
    if old & mutexStarving != 0 {
        // 如果當前鎖即不是被獲取也不是被喚醒狀態, 或者等待隊列為空 這代表鎖狀態產生了不一致的問題
        if old & (mutexLocked | mutexWoken) != 0 || old >> mutexWaiterShift == 0 {
            throw("sync: inconsistent mutex state")
        }

        // 當前 goroutine 已經獲取了鎖, 等待隊列-1
        delta := int32(mutexLocked -1 << mutexWaiterShift)
        // 當前 goroutine 非飢餓狀態 或者 等待隊列只剩下一個 waiter, 則退出飢餓模式(清除飢餓標識位)
        if !starving || old >> mutexWaiterShift == 1 {
            delta -= mutexStarving

            // 更新狀態值並終止 for 循環, 拿到鎖退出
            atomic.AddInt32(&m.state, delta)
            break
        }
        // 設置當前 goroutine 為喚醒狀態, 且重置自旋次數
        awoke = true
        iter = 0
    } else {
        // 鎖被其他 goroutine 佔用了, 還原狀態繼續 for 循環
        old = m.state
    }
}
```

## 解鎖

相對於加鎖操作, 解鎖的邏輯就沒有這麼複雜了, 接下來我們來看一看 Unlock 的邏輯:

```go
func (m *Mutex) Unlock() {
    // Fast path: drop lock bit.
    // 如果 m.state == 0 , 則表示當前鎖已經完全空閒, 結束解鎖
    // m.state != 0 則說明當前鎖沒有被佔用, 但還有等待中且未被喚醒的 goroutine, 需要進行一系列喚醒操作, 這部分邏輯就在 unlockSlow 方法內:
    new := atomic.AddInt32(&m.state, -mutexLocked)
    if new != 0 {
        // Outlined slow path to allow inlining the fast path
        // To hide unlockSlow during tracing we skip one extra frame when tracing GoUnblock.
        m.unlockSlow(new)
    }
}
```

unlockSlow():

```go
func (m *Mutex) unlockSlow(new int32) {
    // 這裡表示解鎖了一個沒被上鎖的鎖, 則直接 Panic
    if (new + mutexLocked) & mutexLocked == 0 {
        panic("sync: unlock of unlocked mutex")
    }

    // 正常模式的釋放鎖邏輯
    if new & mutexStarving == 0 {
        old := new
        for {
            // 如果沒有等待者則直接返回即可
            if old >> mutexWaiterShift == 0 {
                return
            }
            // 如果鎖處於加鎖狀態, 則表示已經有 goroutine 獲取到了鎖, 可以返回 
            if  old & (mutexLocked) != 0 {
                return
            }
            // 如果鎖處於喚醒狀態, 這就表明 有等待的 goroutine 被喚醒了, 不用嘗試獲取其他 goroutine 了
            if old & (mutexWoken) != 0 {
                return
            }
            // 如果鎖處於飢餓模式, 鎖之後會直接給等待對頭 goroutine
            if old & mutexStarving != 0 {
                return 
            }
            // 搶佔喚醒標誌位, 這裡是想要把加鎖的狀態設置為喚醒, 然後 waiter 隊列-1 
            new = (old-1 << mutexWaiterShift) | mutexWoken
            if atomic.CompareAndSwapInt32(&m.state, old, new) {
                // 搶佔成功喚醒一個 goroutine
                runtime_Semrelease(&m.sema, false, 1)
                return
            }
            // 執行搶佔不成功時重新更新一下狀態信息, 下次 for 循環繼續處理
            old = m.state
        }
    } else {
        runtime_Semrelease(&m.sema, true, 1)
    }
}
```

## 非阻塞加鎖

非常簡潔的 TryLock():

```go
func (m *Mutex) TryLock() bool {
    // 記錄當前狀態
     old := m.state
    // 處於加鎖狀態 / 飢餓狀態直接失敗
    if old & (mutexLocked| mutexStarving) != 0 {
        return false
    }
    // 嘗試獲取鎖, 獲取失敗直接失敗
    if !atomic.CompareAndSwapInt32(&m.state, old,old|mutexLocked) {
        return false
    }

    return true
}
```
