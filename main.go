package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gittokkunn/go-qiita-oauth/qiita_oauth"

	//"log"
)

func main() {
	r := gin.Default()
	//r.Static("/img", "./client/img")
	//r.Static("/javascript", "./client/javascript")
	r.LoadHTMLGlob("./index.html")
	r.GET("/", qiita_oauth.LoginHome)
	r.GET("/login", qiita_oauth.RedirectAuthrize)
	r.GET("/callback", qiita_oauth.GetAccessToken)
	r.Run(":3001")

}
