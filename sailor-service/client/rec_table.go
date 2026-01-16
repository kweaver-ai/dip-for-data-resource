package client

type RecTableReq struct {
	BusinessModelId   string `json:"business_model_id"`
	BusinessModelName string `json:"business_model_name"`

	Domain struct {
		DomainId     string `json:"domain_id"`
		DomainName   string `json:"domain_name"`
		DomainPath   string `json:"domain_path"`
		DomainPathId string `json:"domain_path_id"`
	} `json:"domain"`

	Dept struct {
		DeptId     string `json:"dept_id"`
		DeptName   string `json:"dept_name"`
		DeptPath   string `json:"dept_path"`
		DeptPathId string `json:"dept_path_id"`
	} `json:"dept"`

	InfoSystem []InfoSystemItem `json:"info_system"`

	Table struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"table"`

	Fields []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"fields"`
	Key string `json:"key"`
}

type InfoSystemItem struct {
	InfoSystemId   string `json:"info_system_id"`
	InfoSystemName string `json:"info_system_name"`
	InfoSystemDesc string `json:"info_system_desc"`
}

type RecTableResp struct {
	Tables []struct {
		ID       string  `json:"id"`
		HitScore float64 `json:"hit_score"`
		Reason   string  `json:"reason"`
	} `json:"tables"`
}
