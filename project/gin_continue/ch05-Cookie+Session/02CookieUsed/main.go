package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "Not Set"

			// name: cookie's name
			// value: cookie's value
			// maxAge: second
			// path: Only $PATH and $PATH/../../.. can use cookie. $PATH2 can't use $PATH's cookie.
			// domain: cross main domain to use cookie.
			// 			.google.com 			: Lv1 domain
			//			.www.google.com			: Lv2 domain
			//			.data.google.com		: Lv3 domain
			//			.any.data.google.com	: Lv4 domain
			// 			Lv(x) cookie can use in Lv(x+1) - (+Inf)
			// secure: Auto use https:// or not
			// httpOnly: For defending XSS attack, prevent Client's JS access cookis.
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie's value is: %s\n", cookie)
	})
	r.Run(":9900")
}
