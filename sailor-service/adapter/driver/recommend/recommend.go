package recommend

import domain "github.com/kweaver-ai/dip-for-data-resource/sailor-service/domain/recommend"

type Service struct {
	uc domain.UseCase
}

func NewService(uc domain.UseCase) *Service {
	return &Service{uc: uc}
}
