package data_change_mq

import (
	"context"
	"strconv"
	"strings"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

func (c *consumer) deleteEntity(ctx context.Context, graphConfigName string, entityId string, entityIdKey string, entityName string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	graphData := []map[string]string{
		{entityIdKey: entityId},
	}

	_, err = c.adProxy.DeleteEntity(ctx, graphData, entityName, graphIdInt)
	if err != nil {
		return err
	}

	return nil
}

func compareVersions(v1, v2 string) int {
	versions1 := strings.Split(v1, ".")
	versions2 := strings.Split(v2, ".")

	len1, len2 := len(versions1), len(versions2)

	for i := 0; i < len1 || i < len2; i++ {
		var num1, num2 int
		var err1, err2 error

		if i < len1 {
			num1, err1 = strconv.Atoi(versions1[i])
			if err1 != nil {
				// 如果转换失败，可以认为是较小的版本号
				return -1
			}
		}

		if i < len2 {
			num2, err2 = strconv.Atoi(versions2[i])
			if err2 != nil {
				// 如果转换失败，可以认为是较大的版本号
				return 1
			}
		}

		// 如果一个版本有更多的部分，那么未定义的部分默认为 0
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}

	// 如果所有部分都相等，则版本号相等
	return 0
}
