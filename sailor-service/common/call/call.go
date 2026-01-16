package call

import (
	"bytes"
	"context"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func CallInternal(ctx context.Context, client *http.Client, errorMsg string, method string, urlStr string, postReq any, res any) (err error) {
	return call(ctx, client, errorMsg, method, urlStr, postReq, res, "")
}
func call(ctx context.Context, client *http.Client, errorMsg string, method string, urlStr string, postReq any, res any, auth string) (err error) {
	head := http.Header{}
	if auth != "" {
		head.Set("Authorization", auth)
	}
	var reqBody io.Reader
	switch method {
	case http.MethodGet:
		reqBody = nil
	case http.MethodDelete:
		reqBody = nil
	case http.MethodPost:
		jsonReq, err := jsoniter.Marshal(postReq)
		if err != nil {
			log.WithContext(ctx).Error(errorMsg+" json.Marshal error", zap.Error(err))
			return err
		}
		reqBody = bytes.NewReader(jsonReq)
		head.Set("Content-Type", "application/json")
	}

	request, _ := http.NewRequest(method, urlStr, reqBody)
	request.Header = head

	resp, err := client.Do(request.WithContext(ctx))
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"client.Do error", zap.Error(err))
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"io.ReadAll", zap.Error(err))
		return
	}
	log.Infof(errorMsg+" body:%s \n ", body)
	if resp.StatusCode != http.StatusOK {
		return StatusCodeNotOK(errorMsg, resp.StatusCode, body)
	}

	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.WithContext(ctx).Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return
	}
	return nil
}
