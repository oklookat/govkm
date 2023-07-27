package schema

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	ApiUrl = "https://api.moosic.io"
)

type ID string

func (e ID) String() string {
	return string(e)
}

type UnixTime int64

func (e UnixTime) Time() time.Time {
	return time.Unix(int64(e), 0)
}

type (
	// Ответ.
	Response[T any] struct {
		Extra *struct {
			// Может быть доступно при поиске конкретной сущности.
			// Например трека.
			//
			// Допустим вы вызвали SearchTrack с limit 5 и offset 0.
			// Тогда здесь offset будет равен 5.
			Offset int `json:"offset"`
		} `json:"extra"`

		Data T `json:"data"`
	}
)

// Ответ-ошибка.
type ResponseError struct {
	// Название ошибки.
	HError string `json:"error"`

	// Описание ошибки.
	ErrorDescription string `json:"error_description"`
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("%s (%s)", e.HError, e.ErrorDescription)
}

func (e ResponseError) IsUnauthorized() bool {
	return strings.EqualFold(e.HError, "unauthorized")
}

func (e ResponseError) IsNotFound() bool {
	return strings.Contains(e.HError, "invalid_param") || strings.Contains(e.HError, "not_found")
}

func NewErrorWithStatusCode(code int) ErrorWithStatusCode {
	return ErrorWithStatusCode{
		StatusCode: code,
	}
}

type ErrorWithStatusCode struct {
	StatusCode int
}

func (e ErrorWithStatusCode) Error() string {
	return strconv.Itoa(e.StatusCode)
}
