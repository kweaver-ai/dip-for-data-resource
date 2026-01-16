package constant

import (
	"database/sql/driver"
	"strconv"
	"time"

	"github.com/kweaver-ai/idrm-go-frame/core/errorx/agcodes"
	"github.com/kweaver-ai/idrm-go-frame/core/errorx/agerrors"
)

const (
	ServiceName = "AfSailorService"

	DefaultHttpRequestTimeout = 60 * time.Second

	CommonTimeFormat = "2006-01-02 15:04:05"
)

const (
	SortByCreatedAt = "created_at"
	SortByUpdatedAt = "updated_at"
)

const (
	TraceIdContextKey  = "traceID"
	UserInfoContextKey = "info"
	UserTokenKey       = "token"
	UserId             = "userId"
)

const (
	DataCatalogVersion  = "data-catalog"
	DataResourceVersion = "data-resource"
)

const (
	ChatReady      = "ready"
	ChatDataMarket = "data_market_chat"
	ChatGoing      = "chat"
	ChatDelete     = "delete"
)

type ModelID string

func NewModelID(id uint64) ModelID {
	return ModelID(strconv.FormatUint(id, 10))
}

func (m ModelID) Uint64() uint64 {
	if len(m) == 0 {
		return 0
	}

	uintId, err := strconv.ParseUint(string(m), 10, 64)
	if err != nil {
		coder := agcodes.New(ServiceName+".Public.InvalidParameter", "参数值异常", "", "ID需要修改为可解析为数字的字符串", err, "")
		panic(agerrors.NewCode(coder))
	}

	return uintId
}

// Value 实现数据库驱动所支持的值
// 没有该方法会将ModelID在驱动层转换后string，导致与数据库定义类型不匹配
func (m ModelID) Value() (driver.Value, error) {
	return m.Uint64(), nil
}
