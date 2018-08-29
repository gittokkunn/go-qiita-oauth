package qiita_oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


type CredentialInfo struct {
	Token string `json:"token"`
	ClientID string `json:"client_id"`
	Scopes []string `json:"scopes"`
}

type AuthInfo struct {
	Code string `json:"code"`
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
}

var(
	ClientID string
	ClientSecret string
	AccessToken string
)

// Qiitaで認証済みか判定
func LoginHome(c *gin.Context) {
	if AccessToken != "" {
		c.HTML(http.StatusOK, "index.html", nil)
	}else{
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

// Qiita認証画面にリダイレクト
func RedirectAuthrize(c *gin.Context) {
	EnvLoad()
	scope := ""
	state := ""
	RedirectAuthrizeClient(c, ClientID, scope, state)
}

// クライアントIDを指定してリダイレクト
func RedirectAuthrizeClient(c *gin.Context, clientID string, scope string, state string) {
	authURL := "https://qiita.com/api/v2/oauth/authorize?client_id=" + clientID + "&scope=" + scope + "&state=" + state
	c.Redirect(http.StatusMovedPermanently, authURL)
}

// アクセストークンを取得
func GetAccessToken(c *gin.Context) {
	EnvLoad()
	cre := GetAccessTokenClient(c, ClientID, ClientSecret)
	fmt.Println(cre.Token)
	c.Redirect(http.StatusMovedPermanently, "/")
}

// クライアントID, クライアントパスをしていしてアクセストークンを取得
func GetAccessTokenClient(c *gin.Context, clientID string, clientSecret string) (*CredentialInfo){
	code := c.Request.URL.Query().Get("code")
	state := c.Request.URL.Query().Get("state")
	if state == "" {
		fmt.Println("state is empty")
	}
	authtoken, err := json.Marshal(AuthInfo{Code: code, ClientID: clientID, ClientSecret: clientSecret})
	resp, err := http.Post("https://qiita.com/api/v2/access_tokens", "application/json", bytes.NewBuffer(authtoken))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	var cre *CredentialInfo
	json.Unmarshal(byteArray, &cre)
	err = setAccessToken(cre)
	if err != nil {
		panic(err)
	}
	fmt.Println(AccessToken)
	return cre
}

func setAccessToken(cre *CredentialInfo) error {
	AccessToken = cre.Token
	if AccessToken == "" {
		err := errors.New("accessToken is empty")
		return err
	}
	return nil
}
