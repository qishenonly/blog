package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendWith(t *testing.T) {
	tokendata := VerifyTokenData{
		Token: "SAasfFGERWRxqFWEASc",
	}
	err := SendWithTemplate("confirm_email", tokendata, "1050026498@qq.com", "激活帐号！")
	assert.Nil(t, err)

}
