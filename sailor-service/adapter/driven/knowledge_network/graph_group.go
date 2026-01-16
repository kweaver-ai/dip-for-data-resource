package knowledge_network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

type SubGraphBody struct {
	SubgraphId int            `json:"subgraph_id"`
	Name       string         `json:"name"`
	Entity     []*GraphEntity `json:"entity"`
	Edge       []*GraphEdge   `json:"edge"`
}

type SubGraphReq struct {
	GraphId    int            `json:"graph_id"`
	OntologyId int            `json:"ontology_id"`
	Name       string         `json:"name"`
	Entity     []*GraphEntity `json:"entity"`
	Edge       []*GraphEdge   `json:"edge"`
}

type SubGraphResp struct {
	Message    string `json:"message"`
	SubgraphId int    `json:"subgraph_id"`
}

type SubGraphInfo struct {
	Entity    []*GraphEntity `json:"entity"`
	Edge      []*GraphEdge   `json:"edge"`
	EntityNum int            `json:"entity_num"`
	EdgeNum   int            `json:"edge_num"`
	ID        int            `json:"id"`
	Name      string         `json:"name"`
}

type DeleteSubGraphReq struct {
	GraphId     int   `json:"graph_id"`
	SubGraphIDs []int `json:"subgraph_ids"`
}

func (a *ad) CreateSubGraph(ctx context.Context, subGraph *SubGraphReq) (id int, err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	rawURL := a.baseUrl + `/api/builder/v1/graph/subgraph`

	resp, err := httpPostDo[commonResp[SubGraphResp]](ctx, rawURL, subGraph, nil, a)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return resp.Res.SubgraphId, nil
}

func (a *ad) UpdateSubGraph(ctx context.Context, knid int, subGraphs []*SubGraphBody) (err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	rawURL := a.baseUrl + fmt.Sprintf(`/api/builder/v1/graph/subgraph/savenocheck/%v`, knid)

	_, err = httpPostDo[commonResp[any]](ctx, rawURL, subGraphs, nil, a)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (a *ad) GetSubGraph(ctx context.Context, knid int, subGraphName string) (subGraphSlice []*SubGraphInfo, err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	rawURL := a.baseUrl + fmt.Sprintf(`/api/builder/v1/graph/subgraph?graph_id=%v&subgraph_name=%v`, knid, url.QueryEscape(subGraphName))
	if subGraphName == "" {
		rawURL = a.baseUrl + fmt.Sprintf(`/api/builder/v1/graph/subgraph?graph_id=%v&subgraph_name=&return_all=True`, knid)
	}

	realURL, _ := url.Parse(rawURL)

	resp, err := httpGetDoV2[commonResp[[]*SubGraphInfo]](ctx, realURL, a)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return resp.Res, nil
}

func (a *ad) DeleteSubGraph(ctx context.Context, subGraph *DeleteSubGraphReq) (err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	rawURL := a.baseUrl + `/api/builder/v1/graph/subgraph/delete`

	_, err = httpPostDo[commonResp[any]](ctx, rawURL, subGraph, nil, a)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
