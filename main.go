package govkm

import (
	"context"

	"github.com/oklookat/govkm/schema"
	"github.com/oklookat/vantuz"
)

const (
	_errPrefix = "govkm"
)

// Получить Client для запросов к API, и ID текущего пользователя.
//
// Если токен неверный,
// в error будет завернут schema.ResponseError
// (см. schema.ResponseError.IsUnauthorized).
func New(accessToken string) (*Client, error) {
	httpCl := vantuz.C().SetAuthorization("Bearer " + accessToken)
	cl := &Client{
		Http: httpCl,
	}
	cl.SetUserAgent("okhttp/5.0.0-alpha.2")
	cl.Http.SetGlobalHeader("X-App-Id", "android")
	cl.Http.SetGlobalHeader("X-Client-Version", "10477")

	usr, err := cl.UserInfo(context.Background())
	if err != nil {
		return nil, err
	}
	cl.CurrentUserId = usr.Data.User.APIID

	return cl, err
}

// Клиент для запросов к API.
type Client struct {
	// ID текущего пользователя.
	CurrentUserId schema.ID

	// Отправляет запросы.
	Http *vantuz.Client
}

// Установить user agent для запросов.
func (c *Client) SetUserAgent(val string) {
	c.Http.SetUserAgent(val)
}
