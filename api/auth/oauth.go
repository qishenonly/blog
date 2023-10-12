package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"io/ioutil"
	"net/http"
)

var oauthStateString = "random"
var oauthConf = &oauth2.Config{
	ClientID:     "60be73a6e9c78546495d",
	ClientSecret: "4ea51a2af22ac7e4f0b2cc8ae7359eb84d7716a4",
	Scopes:       []string{"read:user", "user:email"},
	Endpoint:     github.Endpoint,
}

// AuthGithub godoc
// @Summary Github授权
// @Description Github授权
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Success 302 {string} string "{"success":true,"data":{},"msg":"授权成功"}"
// @Router /auth/github [get]
func (aa *AuthApi) AuthGithub(c *gin.Context) {
	url := oauthConf.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// AuthGithubCallback godoc
// @Summary Github授权回调
// @Description Github授权回调
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Param state query string true "state"
// @Param code query string true "code"
// @Success 302 {string} string "{"success":true,"data":{},"msg":"授权成功"}"
func (aa *AuthApi) AuthGithubCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		return
	}

	code := c.Query("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		return
	}

	if !token.Valid() {
		fmt.Println("Token not valid")
		return
	}

	client := oauthConf.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		fmt.Printf("Request failed with '%s'\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Failed to close response body with '%s'\n", err)
			return
		}
	}(resp.Body)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body with '%s'\n", err)
		return
	}

	c.String(http.StatusOK, "Github User Info: %s\n", content)
}
