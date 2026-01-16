//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/kweaver-ai/dip-for-data-resource/sailor-service/adapter/driven"
	"github.com/kweaver-ai/dip-for-data-resource/sailor-service/adapter/driver"
	"github.com/kweaver-ai/dip-for-data-resource/sailor-service/common/settings"
	"github.com/kweaver-ai/dip-for-data-resource/sailor-service/domain"
	"github.com/kweaver-ai/dip-for-data-resource/sailor-service/infrastructure"
	af_go_frame "github.com/kweaver-ai/idrm-go-frame"
)

func InitApp(httpCfg settings.HttpConf) (*af_go_frame.App, func(), error) {
	panic(
		wire.Build(
			driver.RouterSet,
			driver.Set,
			driven.Set,
			domain.Set,
			infrastructure.Set,
			newApp,
			toTransportServer,
		),
	)
}
