package knowledge_network

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type ListKnowledgeNetworkReq struct {
	Order string `json:"order"`       // 默认按照开始时间从新至旧排序，接受参数为：'desc'（从新到旧），'asc'（从旧到新）
	Page  int    `json:"page,string"` // 页码
	Size  int    `json:"size,string"` // 每页数量
	Rule  string `json:"rule"`        // 排序字段，可选start_time, end_time
}

type ListKnowledgeNetworkResp struct {
	Res struct {
		Count int `json:"count"`
		Df    []struct {
			Color             string `json:"color"`
			CreationTime      string `json:"creation_time"`
			CreatorId         string `json:"creator_id"`
			CreatorName       string `json:"creator_name"`
			FinalOperator     string `json:"final_operator"`
			GroupColumn       int    `json:"group_column"`
			Id                int    `json:"id"`
			IdentifyId        string `json:"identify_id"`
			IntelligenceScore string `json:"intelligence_score"`
			KnwDescription    string `json:"knw_description"`
			KnwName           string `json:"knw_name"`
			OperatorName      string `json:"operator_name"`
			ToBeUploaded      int    `json:"to_be_uploaded"`
			UpdateTime        string `json:"update_time"`
		} `json:"df"`
	} `json:"res"`
}

func (a *ad) ListKnowledgeNetwork(ctx context.Context, req *ListKnowledgeNetworkReq) (*ListKnowledgeNetworkResp, error) {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	rawURL := a.baseUrl + `/api/builder/v1/knw/get_all`
	u, err := url.Parse(rawURL)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to parse url: %s, err: %v", rawURL, err)
		return nil, errors.Wrap(err, "parse url failed")
	}

	var m map[string]string
	if err = json.Unmarshal(lo.T2(json.Marshal(req)).A, &m); err != nil {
		log.WithContext(ctx).Errorf("json.Unmarshal(lo.T2(json.Marshal(req)).A, &m) failed, err: %v", err)
		return nil, errors.Wrap(err, `req param marshal json failed`)
	}

	values := url.Values{}
	for k, v := range m {
		values.Add(k, v)
	}

	u.RawQuery = values.Encode()

	return httpGetDoV2[ListKnowledgeNetworkResp](ctx, u, a)
}

type ListKnowledgeGraphReq struct {
	Filter string `json:"filter"`        // all
	KnwId  int    `json:"knw_id,string"` // 网络id
	Order  string `json:"order"`         // 默认按照开始时间从新至旧排序，接受参数为：'desc'（从新到旧），'asc'（从旧到新）
	Page   int    `json:"page,string"`   // 页码
	Size   int    `json:"size,string"`   // 每页数量
	Rule   string `json:"rule"`          // 排序字段，可选start_time, end_time
	Name   string `json:"name"`          // 筛选名字
}

type ListKnowledgeGraphResp struct {
	Res struct {
		Count int `json:"count"`
		Df    []struct {
			//CreateUser    string `json:"createUser"`
			//CreateTime    string `json:"create_time"`
			//Export        bool   `json:"export"`
			//GraphDbName   string `json:"graph_db_name"`
			Id int `json:"id"`
			//IsImport      bool   `json:"is_import"`
			//KgDesc        string `json:"kgDesc"`
			//KnowledgeType string `json:"knowledge_type"`
			KnwId int    `json:"knw_id"`
			Name  string `json:"name"`
			//OtlId         string `json:"otl_id"`
			//RabbitmqDs    int    `json:"rabbitmqDs"`
			//Status        string `json:"status"`
			//StepNum       int    `json:"step_num"`
			//Taskstatus    string `json:"taskstatus"`
			//UpdateTime    string `json:"updateTime"`
			//UpdateUser    string `json:"updateUser"`
		} `json:"df"`
	} `json:"res"`
}

func (a *ad) ListKnowledgeGraph(ctx context.Context, req *ListKnowledgeGraphReq) (*ListKnowledgeGraphResp, error) {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	rawURL := a.baseUrl + `/api/builder/v1/knw/get_graph_by_knw`
	u, err := url.Parse(rawURL)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to parse url: %s, err: %v", rawURL, err)
		return nil, errors.Wrap(err, "parse url failed")
	}

	var m map[string]string
	if err = json.Unmarshal(lo.T2(json.Marshal(req)).A, &m); err != nil {
		log.WithContext(ctx).Errorf("json.Unmarshal(lo.T2(json.Marshal(req)).A, &m) failed, err: %v", err)
		return nil, errors.Wrap(err, `req param marshal json failed`)
	}

	values := url.Values{}
	for k, v := range m {
		values.Add(k, v)
	}

	u.RawQuery = values.Encode()

	return httpGetDoV2[ListKnowledgeGraphResp](ctx, u, a)
}

type ListKnowledgeLexiconReq struct {
	KnowledgeId int    `json:"knowledge_id,string"` // 网络id
	Order       string `json:"order"`               // 默认按照开始时间从新至旧排序，接受参数为：'desc'（从新到旧），'asc'（从旧到新）
	Page        int    `json:"page,string"`         // 页码
	Size        int    `json:"size,string"`         // 每页数量
	Rule        string `json:"rule"`                // 排序字段，可选start_time, end_time
	Word        string `json:"word"`                // 筛选名字
}

type ListKnowledgeLexiconResp struct {
	Res struct {
		Count int `json:"count"`
		Df    []struct {
			Columns   []string `json:"columns"`
			ErrorInfo string   `json:"error_info"`
			Id        int      `json:"id"`
			Name      string   `json:"name"`
			Status    string   `json:"status"`
		} `json:"df"`
	} `json:"res"`
}

func (a *ad) ListKnowledgeLexicon(ctx context.Context, req *ListKnowledgeLexiconReq) (*ListKnowledgeLexiconResp, error) {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	rawURL := a.baseUrl + `/api/builder/v1/lexicon/getall`
	u, err := url.Parse(rawURL)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to parse url: %s, err: %v", rawURL, err)
		return nil, errors.Wrap(err, "parse url failed")
	}

	var m map[string]string
	if err = json.Unmarshal(lo.T2(json.Marshal(req)).A, &m); err != nil {
		log.WithContext(ctx).Errorf("json.Unmarshal(lo.T2(json.Marshal(req)).A, &m) failed, err: %v", err)
		return nil, errors.Wrap(err, `req param marshal json failed`)
	}

	values := url.Values{}
	for k, v := range m {
		values.Add(k, v)
	}

	u.RawQuery = values.Encode()

	return httpGetDoV2[ListKnowledgeLexiconResp](ctx, u, a)
}
