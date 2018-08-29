# Qiitaのoauth認証

## About
### RedirectAuthrizeClient(c *gin.Context, clientID string, scope string state string)
### GetAccessTokenClient(c *gin.Context, clientID string, clientSecret string) (*CredentialInfo)
### CredentialInfo
##### `.Token string`
- 認証に必要なアクセストークン
##### `.Scopes string`
- APIアクセスのスコープ
##### `.ClientID`
- クライアントID