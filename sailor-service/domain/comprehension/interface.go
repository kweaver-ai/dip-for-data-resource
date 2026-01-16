package comprehension

import (
	"context"
)

type ReqArgs struct {
	CatalogID string `json:"catalog_id" form:"catalog_id" binding:"required"`
	Dimension string `json:"dimension" form:"dimension" binding:"required"`
}

type Q struct {
	Query string `json:"question" form:"question" binding:"required"`
}

type Domain interface {
	AIComprehension(ctx context.Context, catalogId string, dimension string) (any, error) //根据AI返回不同的理解内容
	AIComprehensionConfig() any                                                           //根据AI返回不同的理解内容
	SetAIComprehensionConfig(id string) string
}
