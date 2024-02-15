現在我使用 go 語言在設計一個行為控制器, 但是現在遇到一些問題.

主要關注點是 actionHandleInstance, 它是一個 interface

```go
type ActionController struct {
    // ...

	// actionHandle 實例
	actionHandleInstance ActionHandleI
}

// ActionHandleI 玩家動作處理器, 用來處理玩家的動作, 每個遊戲需要實現自己的 ActionHandleI
type ActionHandleI interface {
	// 重置玩家动作
	Reset()

	// 處理玩家請求
	HandlePlayerReq(playerID uint64, req any) error

	// 處理玩家超時
	HandlePlayerTimeout(playerID uint64) error
}
```

然後使用 HandlePlayerReq 的地方只有一個:

```go
// 模拟接收玩家动作的函数
func (ac *ActionController) receiveAction(id uint64) {
	defer func() {
		ac.reciveActionWg.Done()
	}()

	uc, ok := ac.userInfoMap.Load(id)
	if !ok {
		log.Error("error on receiveAction(): user(%v) info not exist!\n", id)
		return
	}

	for {
		select {
		case <-uc.forceCloseChan:
			log.Info("force close receiveAction goroutine")
			return
		case req := <-uc.reqChan:
			log.Debug("receiveAction: receive player action: %v\n", req)

			// 先重新計算剩餘緩衝時間
			uc.countdown.RecalculateRemainingBuffer()

			// 再處理玩家邏輯
			if err := ac.actionHandleInstance.HandlePlayerReq(id, req); err != nil {
				log.Error("error on receiveAction(): handle player(%v) req(%v) failed: %v\n", id, req, err)
			}

			return
		case <-time.After(uc.countdown.TotalCountdown()):
			log.Info("receiveAction: player(%v) action timeout, do timeout handler\n", id)
			ac.actionHandleInstance.HandlePlayerTimeout(id)
			return
		}
	}
}
```

然後傳入消息的時候使用這個 function:
```go
// SendPlayerActionReq 將玩家動作請求發送到對應的 reqChan
func (ac *ActionController) SendPlayerActionReq(id uint64, req any) {
	uc, ok := ac.userInfoMap.Load(id)
	if !ok {
		log.Warn("SendPlayerActionReq: user(%v) info not exist!\n", id)
		return
	}

	if !uc.isUserTurn {
		log.Warn("SendPlayerActionReq: user(%v) not in turn, ignore\n", id)
		return
	}

		uc.reqChan <- req
}
```

現在有個問題是: 我的請求會通過 channel 傳入, 然後再被處理(uc.reqChan <- req)

但是我現在希望的是, 如果 actionHandler 有工作在處理的時候, 就不能傳進去, 請問這個功能要怎麼做?

