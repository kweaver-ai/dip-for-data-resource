package client

type RecViewReq struct {
	RecommendViewTypes []int `json:"recommend_view_types"`
	Table              struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"table"`
	Fields []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"fields"`
}

type RecViewResp struct {
	Res struct {
		Answers struct {
			Views []struct {
				Id       string  `json:"id"`
				HitScore float64 `json:"hit_score"`
				Reason   string  `json:"reason"`
			} `json:"views"`
		} `json:"answers"`
	} `json:"res"`
}
