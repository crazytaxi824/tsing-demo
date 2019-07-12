package action

import (
	"time"

	"github.com/dxvgef/filter"
	"github.com/dxvgef/tsing"
	"github.com/gbrlsnchs/jwt/v3"
)

type AccessToken struct {
	Data struct {
		ID       int64  `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
	} `json:"data,omitempty"`
	jwt.Payload
}

type Example struct{}

func (Example) SignJWT(ctx tsing.Context) error {
	var respData RespData
	var accessToken AccessToken
	accessToken.Data.ID = 123
	accessToken.Data.Username = "dxvgef"
	accessToken.ExpirationTime = jwt.NumericDate(time.Now().Add(time.Minute))

	alg := jwt.NewHS256([]byte("secret"))
	token, err := jwt.Sign(accessToken, alg)
	if err != nil {
		return ctx.Event(err)
	}
	respData.Data = string(token)

	return JSON(ctx, 200, respData)
}

func (Example) Admin(ctx tsing.Context) error {
	var reqData struct {
		username string
		password string
	}
	// 过滤多个元素
	err := filter.MSet(
		// 要过滤的元素
		filter.El(&reqData.username, filter.FromString(ctx.RawQueryValue("username"), "账号").
			RemoveSpace().MinLength(3).MaxLength(16)), // 要求最大长度
		filter.El(&reqData.password, filter.FromString(ctx.RawQueryValue("password"), "密码").
			MinLength(6).MaxLength(32).HasDigit().HasUpper().HasLower().HasSymbol(),
		),
	)
	if err != nil {
		return ctx.Event(err)
	}
	return String(ctx, 200, "Hello Tsing")
}
