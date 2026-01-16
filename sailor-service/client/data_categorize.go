package client

type LogicalViewDatacategorizeReq struct {
	ViewId            string `json:"view_id"`
	ViewTechnicalName string `json:"view_technical_name"`
	ViewBusinessName  string `json:"view_business_name"`
	ViewDesc          string `json:"view_desc"`
	SubjectId         string `json:"subject_id"`
	ViewFields        []struct {
		ViewFieldId            string `json:"view_field_id"`
		ViewFieldTechnicalName string `json:"view_field_technical_name"`
		ViewFieldBusinessName  string `json:"view_field_business_name"`
		StandardCode           string `json:"standard_code"`
	} `json:"view_fields"`
	ExploreSubjectIds     []string `json:"explore_subject_ids"`
	ViewSourceCatalogName string   `json:"view_source_catalog_name"`
}

type LogicalViewDataCategorizeResp struct {
	Res struct {
		Answers struct {
			ViewId     string `json:"view_id"`
			ViewFields []struct {
				ViewFieldId string `json:"view_field_id"`
				RelSubjects []struct {
					SubjectId string `json:"subject_id"`
					Score     string `json:"score"`
				} `json:"rel_subjects"`
			} `json:"view_fields"`
		} `json:"answers"`
	} `json:"res"`
}
