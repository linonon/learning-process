# Redis Database
安裝完默認有16個數據庫

## Redis 基本使用
- `set [key] [value]`  set
- `get [key]` get
- `selece 1` 切換到 1號 數據庫
- `dbsize` 查數據，默認當前
- `flushdb` 清空
- `del [key]` 刪除
- `setex [key] [sec] [value]` 一定時間刪除
- `mset`    一次性set多個
- `mget`    一次性get多個

## Redis 數據類型 和 CRUD

- 五大數據類型
    1.  String
    2.  Hash
    3.  List
    4.  Set
    5.  Zset (sorted set)

### String 
- key : value
- 二進制安全的，也可以存圖片(byte切片)
- value一個String最大512MB

### Hash
- `hset user1 [k1] [v1] [k2] [v2] ... ...`
- `hgetall` 獲得全部
- `hmset` & `hmget` 一次性設置
- `hexists key field` 查看是否存在
- key 不可重複

### List
- Head(left) end(right)
- 元素有序
- 元素可重複
- `lpush`：左邊push & `rpush`：右邊push
- `lpop` & `rpop`
- `lset` 固定位置放
- `lrange` 查看一定範圍的數據

### Set
- 無序集合
- 不可重複（如通過email註冊）
- `sadd`  & `smembers`/`sismember` 增改 & 查
- `srem` 刪除

## Connect Go to Redis

