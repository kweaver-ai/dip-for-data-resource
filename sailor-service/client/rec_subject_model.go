package client

type RecSubjectModelReq struct {
	Query string `json:"query"`
}

type RecSubjectModelResp struct {
	Data []struct {
		ID             string `json:"id"`
		BusinessName   string `json:"business_name"`
		DataViewID     string `json:"data_view_id"`
		DisplayFieldID string `json:"display_field_id"`
		TechnicalName  string `json:"technical_name"`
	} `json:"data"`
}
