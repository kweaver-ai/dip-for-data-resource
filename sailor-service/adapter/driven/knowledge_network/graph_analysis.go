package knowledge_network

import "context"

//GraphAnalysis   自定义认知服务，表单推荐等接口
func (a *ad) GraphAnalysis(ctx context.Context, serviceId string, content any) (*GraphAnalysisResp, error) {
	uri := a.baseUrl + "/api/engine/v1/open/custom-search/services/" + serviceId
	headers := make(map[string][]string)
	headers["source-type"] = []string{"2"}
	return httpPostDo[GraphAnalysisResp](ctx, uri, content, headers, a)
}
