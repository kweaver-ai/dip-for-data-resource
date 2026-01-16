package model

type DataViewFields struct {
	FormViewId    string `json:"form_view_id"`
	TechnicalName string `json:"technical_name"`
	BusinessName  string `json:"business_name"`
}

type SVCFields struct {
	ServiceId string `json:"service_id"`
	CnName    string `json:"cn_name"`
	EnName    string `json:"en_name"`
}

type IndicatorFields struct {
	Id                string `json:"id"`
	AnalysisDimension string `json:"analysis_dimension"`
}

type CatalogFormView struct {
	Code   string `json:"code"`
	SResId string `json:"s_res_id"`
}
