package govkm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/oklookat/govkm/schema"
	"github.com/oklookat/vantuz"
)

// genApiPath создает URL-адрес для запроса к API, используя заданный путь.
//
// Пример использования: genApiPath([]string{"users", iToString(1234), "playlists", "list"})
//
// Результат: https://api.music.yandex.net/users/1234/playlists/list
func genApiPath(paths ...string) string {
	if len(paths) == 0 {
		return schema.ApiUrl
	}

	base := schema.ApiUrl + "/" + paths[0]
	for i := 1; i < len(paths); i++ {
		if len(paths[i]) == 0 {
			continue
		}
		base += "/" + paths[i]
	}

	return base
}

var (
	// Странная ошибка.
	ErrNilResponse = errors.New(_errPrefix + ": nil http response")
)

func checkResponse[T any](resp *vantuz.Response, respError *schema.ResponseError, data *schema.Response[T]) error {
	if resp == nil {
		return ErrNilResponse
	}
	if resp.IsSuccess() {
		return nil
	}
	if respError != nil && len(respError.Error()) > 0 {
		return fmt.Errorf("%s: %w", _errPrefix, respError)
	}
	return fmt.Errorf("%s: %w", _errPrefix, schema.NewErrorWithStatusCode(resp.StatusCode))
}

func getLiked[T any](ctx context.Context, client *Client, entityName string, limit, offset int) (*schema.Response[T], error) {
	// "liked/" - со слешем в конце. Так и задумано.
	endpoint := genApiPath("user", entityName, "liked/")
	params := getLimitOffset(limit, offset)
	data := &schema.Response[T]{}
	respErr := &schema.ResponseError{}
	resp, err := client.Http.R().
		SetQueryParams(params).
		SetError(respErr).
		SetResult(&data).
		Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, respErr, data)
	}
	return data, err
}

func getEntityById[T any](ctx context.Context, client *Client, entityName string, id schema.ID) (*schema.Response[T], error) {
	endpoint := genApiPath(entityName, id.String())
	data := &schema.Response[T]{}
	respErr := &schema.ResponseError{}
	resp, err := client.Http.R().
		SetError(respErr).
		SetResult(&data).
		Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, respErr, data)
	}
	return data, err
}

func getAny[T any](ctx context.Context, client *Client, queryParams url.Values, paths ...string) (*schema.Response[T], error) {
	endpoint := genApiPath(paths...)
	data := &schema.Response[T]{}
	respErr := &schema.ResponseError{}
	req := client.Http.R().
		SetError(respErr).
		SetResult(&data)
	if queryParams != nil {
		req.SetQueryParams(queryParams)
	}
	resp, err := req.Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, respErr, data)
	}
	return data, err
}

func writeableAny[T any](ctx context.Context, client *Client, method string, body url.Values, paths ...string) (*schema.Response[T], error) {
	endpoint := genApiPath(paths...)
	data := &schema.Response[T]{}
	respErr := &schema.ResponseError{}

	req := client.Http.R().
		SetError(respErr).
		SetResult(&data)

	if body != nil {
		req.SetFormUrlValues(body)
	}

	var (
		resp *vantuz.Response
		err  error
	)
	switch method {
	case http.MethodPost:
		resp, err = req.Post(ctx, endpoint)
	case http.MethodPut:
		resp, err = req.Put(ctx, endpoint)
	case http.MethodDelete:
		resp, err = req.Delete(ctx, endpoint)
	}

	if err == nil {
		err = checkResponse(resp, respErr, data)
	}
	return data, err
}

func getLimitOffset(limit, offset int) url.Values {
	params := url.Values{}
	if offset > 0 {
		params.Set("offset", strconv.Itoa(offset))
	}
	if limit == 0 {
		limit = 1
	}
	params.Set("limit", strconv.Itoa(limit))
	return params
}

func searchEntity[T any](ctx context.Context, client *Client,
	query, entityType string, limit, offset int) (*schema.Response[T], error) {

	params := getLimitOffset(limit, offset)
	params.Set("q", query)

	return getAny[T](ctx, client, params, "search", entityType, "/")
}

func likeUnlikeEntity(
	ctx context.Context,
	client *Client,
	entityType string,
	entityId schema.ID,
	like bool,
) (*schema.Response[any], error) {

	endpoint := genApiPath(entityType, entityId.String(), "like")
	data := &schema.Response[any]{}
	respErr := &schema.ResponseError{}
	req := client.Http.R().SetError(respErr).SetResult(&data)
	var resp *vantuz.Response
	var err error
	if like {
		resp, err = req.Put(ctx, endpoint)
	} else {
		resp, err = req.Delete(ctx, endpoint)
	}
	if err == nil {
		err = checkResponse(resp, respErr, data)
	}
	return data, err
}

func valuesFileId(id []schema.ID) url.Values {
	params := url.Values{}
	for _, id := range id {
		params.Add("file_id", id.String())
	}
	return params
}
