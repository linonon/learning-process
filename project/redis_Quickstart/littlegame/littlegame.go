package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type User struct {
	Id      string `json:"Id"`
	Balance int    `json:"balance"`
}
type UserBet struct {
	Id     string `json:"Id"`
	Round  int    `json:"round"`
	Amount int    `json:"amount"`
}

var Round = 0
var startTimeThisRound time.Time
var RC *redis.Client

const (
	RoundSecond    = 30
	DefaultBalance = 10000  // 初始化金額
	UserMember     = "game" // 存儲使用者的balance
	BetThisRound   = "bet_this_round"
)

func init() {
	RC = newClient()

	RC.Del(UserMember)
	RC.Del(BetThisRound)

	go GameServer()
}

func main() {
	router := gin.Default()

	router.RedirectFixedPath = true
	router.GET("/bet/:user", GetUserBalance)
	router.GET("/bet/:user/:amount", Bet)
	router.GET("/prize", GetCurrentPrize)
	router.GET("/bets", GetUserBets)
	router.Run(":80")
}
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       1,
	})
	pong, err := client.Ping().Result()
	log.Println(pong)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GameServer() {
	rand.Seed(time.Now().UTC().Unix())
	ticker := time.NewTicker(RoundSecond * time.Second)
	go func() {
		for {
			Round++
			startTimeThisRound = time.Now()
			log.Println(startTimeThisRound.Format("2006-01-02 15:04:05"), "\t round", Round, "start")
			_ = <-ticker.C

			var prizePool = getCurrentPrize()
			var userBets = getUserBets()
			if len(userBets) == 0 {
				log.Println("Round", Round, "沒有任何玩家下注")
				continue
			}

			winNum := rand.Intn(prizePool + 1)
			var winner string
			for _, userBet := range userBets {
				winNum -= userBet.Amount
				if winNum <= 0 {
					winner = userBet.Id
					break
				}
			}
			log.Println("獎金池:", prizePool, "\t 得主:", winner)

			RC.ZIncrBy(UserMember, float64(prizePool), winner)
			RC.Del(BetThisRound)
		}
	}()
}

func Bet(c *gin.Context) {
	var user User
	user.Id = c.Param("user")
	amountStr := c.Param("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		wrapResponse(c, nil, errors.New("下注金額有誤"))
		return
	}

	balance, err := RC.ZScore(UserMember, user.Id).Result()
	if err == redis.Nil {
		wrapResponse(c, nil, errors.New("查無此用戶，請先註冊"))
		return
	}
	user.Balance = int(balance)
	if amount <= 0 {
		wrapResponse(c, nil, errors.New("金額需要為正整數"))
		return
	}
	if amount > user.Balance {
		wrapResponse(c, nil, errors.New("餘額不足"))
		return
	}

	user.Balance -= amount
	RC.ZIncrBy(UserMember, float64(-amount), user.Id)
	RC.ZIncrBy(BetThisRound, float64(amount), user.Id)

	wrapResponse(c, user, nil)
}

func GetUserBalance(c *gin.Context) {
	var user User
	user.Id = c.Param("user")
	balance, err := RC.ZScore(UserMember, user.Id).Result()
	if err == redis.Nil { // 查無此人，新建用戶
		balance = DefaultBalance
		RC.ZAdd(UserMember, redis.Z{
			Score:  balance,
			Member: user.Id,
		})
	}
	user.Balance = int(balance)
	wrapResponse(c, user, nil)

}

func GetCurrentPrize(c *gin.Context) {
	wrapResponse(c, getCurrentPrize(), nil)
}

func getCurrentPrize() (prizePool int) {
	bets, _ := RC.ZRangeWithScores(BetThisRound, 0, -1).Result()
	for _, bet := range bets {
		var userBet UserBet
		userBet.Amount = int(bet.Score)
		prizePool += userBet.Amount
	}
	return
}

func GetUserBets(c *gin.Context) {
	UserBets := getUserBets()
	if len(UserBets) == 0 {
		wrapResponse(c, nil, errors.New("目前咩有紀錄"))
		return
	}
	wrapResponse(c, UserBets, nil)
}

func getUserBets() (userBets []UserBet) {
	bets, _ := RC.ZRangeWithScores(BetThisRound, 0, -1).Result()
	for _, bet := range bets {
		var userBet UserBet
		userBet.Id = fmt.Sprintf("%s", bet.Member)
		userBet.Amount = int(bet.Score)
		userBet.Round = Round
		userBets = append(userBets, userBet)
	}
	return
}

func wrapResponse(c *gin.Context, data interface{}, err error) {
	type ret struct {
		Status string      `json:"status"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
	}

	d := ret{
		Status: "ok",
		Msg:    "",
		Data:   []struct{}{},
	}

	if data != nil {
		d.Data = data
	}

	if err != nil {
		d.Status = "failed"
		d.Msg = err.Error()
	}

	c.JSON(http.StatusOK, d)
}
