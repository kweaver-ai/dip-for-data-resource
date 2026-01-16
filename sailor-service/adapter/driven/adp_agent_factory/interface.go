package adp_agent_factory

import "context"

type Repo interface {
	AgentList(ctx context.Context, req AgentListReq) (resp *AgentListResp, err error)

	AgentListV2(ctx context.Context, req AgentListReq) (resp *AgentListResp, err error)
}

type AgentListReq struct {
	Name                string   `json:"name"`
	Ids                 []string `json:"ids"`
	AgentKeys           []string `json:"agent_keys"`
	ExcludeAgentKeys    []string `json:"exclude_agent_keys"`
	CategoryId          string   `json:"category_id"`
	PublishToBe         string   `json:"publish_to_be"`
	CustomSpaceId       string   `json:"custom_space_id"`
	Size                int      `json:"size"`
	PaginationMarkerStr string   `json:"pagination_marker_str"`
	IsToCustomSpace     int      `json:"is_to_custom_space"`
	IsToSquare          int      `json:"is_to_square"`
	BusinessDomainIds   []string `json:"business_domain_ids"`
}

//type AgentListReq struct {
//	Name             string   `form:"name" json:"name"`                       // 根据名称模糊查询
//	IDs              []string `json:"ids"`                                    // 根据ID查询
//	AgentKeys        []string `json:"agent_keys"`                             // 根据智能体标识查询
//	ExcludeAgentKeys []string `json:"exclude_agent_keys"`                     // 排除智能体标识
//	CategoryID       string   `form:"category_id" json:"category_id"`         // 分类ID
//	PublishToBe      string   `form:"publish_to_be" json:"publish_to_be"`     // 发布为标识("api_agent", "web_sdk_agent", "skill_agent")
//	CustomSpaceID    string   `form:"custom_space_id" json:"custom_space_id"` // 自定义空间ID
//
//	IsToCustomSpace int `form:"is_to_custom_space" json:"is_to_custom_space"` // 获取发布到自定义空间的智能体
//
//	IsToSquare int `form:"is_to_square" json:"is_to_square"` // 获取发布到广场的智能体
//
//	BusinessDomainIDs []string `json:"business_domain_ids"` // 业务域ID数组，如果不传，会使用header中的"x-business-domain"。如果该header也没有传，会默认使用"公共业务域"进行过滤
//
//	Size int `form:"size,default=10" json:"size" binding:"numeric,max=1000"` // 每页显示数量
//
//	PaginationMarkerStr string `form:"pagination_marker_str" json:"pagination_marker_str"` // 上一次查询的最后一条记录对应的pagination_marker_str
//}

type AgentItem struct {
	Id              string `json:"id"`
	Key             string `json:"key"`
	IsBuiltIn       int    `json:"is_built_in"`
	IsSystemAgent   int    `json:"is_system_agent"`
	Name            string `json:"name"`
	Profile         string `json:"profile"`
	Version         string `json:"version"`
	AvatarType      int    `json:"avatar_type"`
	Avatar          string `json:"avatar"`
	PublishedAt     int64  `json:"published_at"`
	PublishedBy     string `json:"published_by"`
	PublishedByName string `json:"published_by_name"`
	PublishInfo     struct {
		IsApiAgent      int `json:"is_api_agent"`
		IsSdkAgent      int `json:"is_sdk_agent"`
		IsSkillAgent    int `json:"is_skill_agent"`
		IsDataFlowAgent int `json:"is_data_flow_agent"`
	} `json:"publish_info"`
	BusinessDomainId string `json:"business_domain_id"`
	ListStatus       string `json:"list_status"`
	AFAgentID        string `json:"af_agent_id"`
}

type AgentListResp struct {
	Entries             []AgentItem `json:"entries"`
	PaginationMarkerStr string      `json:"pagination_marker_str"`
	IsLastPage          bool        `json:"is_last_page"`
}
