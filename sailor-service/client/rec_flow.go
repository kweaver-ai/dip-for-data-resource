package client

type RecFlowReq struct {
	BusinessModelId string `json:"business_model_id"`
	Node            struct {
		Id          string `json:"id"`
		MxcellId    string `json:"mxcell_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"node"`
	ParentNode struct {
		Id          string   `json:"id"`
		MxcellId    string   `json:"mxcell_id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Tables      []string `json:"tables"`
	} `json:"parent_node"`
	Flowchart struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Nodes       []struct {
			Id          string `json:"id"`
			MxcellId    string `json:"mxcell_id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"nodes"`
	} `json:"flowchart"`
}

type RecFlowResp struct {
	Flowcharts []struct {
		Id       string  `json:"id"`
		HitScore float64 `json:"hit_score"`
		Reason   string  `json:"reason"`
	} `json:"flowcharts"`
}
