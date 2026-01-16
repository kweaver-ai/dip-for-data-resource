package intelligence

import "context"

type UseCase interface {
	TableSampleData(ctx context.Context, req *SampleDataReq) (*SampleDataResp, error)
}

type SampleItem map[string]any

type SampleDataReq struct {
	SampleDataReqBody `param_type:"body"`
}

type SampleDataReqBody struct {
	Titles  []string     `json:"titles" form:"titles" binding:"required"`
	Example []SampleItem `json:"example" form:"example"`
	Differs []string     `json:"differs" form:"differs"`
}

type SampleDataResp struct {
	Count      int          `json:"count"`
	SampleData []SampleItem `json:"sample_data"`
}
