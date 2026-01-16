package client

type TableCompletionTableInfoReqBody struct {
	Id            string `json:"id"`
	TechnicalName string `json:"technical_name"`
	BusinessName  string `json:"business_name"`
	Desc          string `json:"desc"`
	Database      string `json:"database"`
	Subject       string `json:"subject"`
	Columns       []struct {
		Id            string `json:"id"`
		TechnicalName string `json:"technical_name"`
		BusinessName  string `json:"business_name"`
		DataType      string `json:"data_type"`
		Comment       string `json:"comment"`
	} `json:"columns"`
	RequestType           int    `json:"request_type"`
	ViewSourceCatalogName string `json:"view_source_catalog_name"`
}

type TableCompletionTableInfoResp struct {
	Res struct {
		TaskId string `json:"task_id"`
	} `json:"res"`
}

type TableCompletionReqBody struct {
	Id            string `json:"id"`
	TechnicalName string `json:"technical_name"`
	BusinessName  string `json:"business_name"`
	Desc          string `json:"desc"`
	Database      string `json:"database"`
	Subject       string `json:"subject"`
	Columns       []struct {
		Id            string `json:"id"`
		TechnicalName string `json:"technical_name"`
		BusinessName  string `json:"business_name"`
		DataType      string `json:"data_type"`
		Comment       string `json:"comment"`
	} `json:"columns"`
	RequestType           int      `json:"request_type"`
	GenFieldIds           []string `json:"gen_field_ids"`
	ViewSourceCatalogName string   `json:"view_source_catalog_name"`
}

type TableCompletionResp struct {
	Res struct {
		TaskId string `json:"task_id"`
	} `json:"res"`
}
