package copilot

import domain "github.com/kweaver-ai/dip-for-data-resource/sailor-service/domain/copilot"

type Service struct {
	uc domain.UseCase
}

func NewService(uc domain.UseCase) *Service {
	return &Service{uc: uc}
}
