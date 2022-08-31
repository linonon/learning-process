package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

type LoginForm struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"passowrd" xml:"password" binding:"required"`
}
type SignUpForm struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎屬性，做到訂製
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New()
		// 第一個參數是備用的語言環境，後面的參數是應該支持的語言環境
		uni := ut.New(zhT, enT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			fmt.Println("1")
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			fmt.Println("2")
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			zh_translations.RegisterDefaultTranslations(v, trans)
		}
		return

	}
	return
}

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻譯器錯誤")
		return
	}
	router := gin.Default()
	router.POST("/loginJSON", func(c *gin.Context) {
		var loginForm LoginForm
		if err := c.ShouldBind(&loginForm); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}

			// fmt.Println(err.Error())

			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs.Translate(trans),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "login successd",
		})
	})

	router.POST("/signup", func(c *gin.Context) {
		var signFrom SignUpForm
		if err := c.ShouldBind(&signFrom); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "signup successd",
		})
	})
	router.Run(":8080")
}
