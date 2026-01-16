package settings

import (
	"sync"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry"

	"os"
	"strings"

	"github.com/kweaver-ai/idrm-go-frame/core/options"
)

var (
	lock   = new(sync.RWMutex)
	once   = new(sync.Once)
	config *Config
)

func Init() {
	//初始化配置文件
	if config.ServerConf.SwagConf.Host == "" {
		config.ServerConf.SwagConf.Host = "0.0.0.0:8133"
	}
}

func initAI() {
	if config.AfSailorConf.URL == "" {
		config.AfSailorConf.URL = "http://af-sailor:9797"
	}

	if config.AfSailorAgentConf.URL == "" {
		config.AfSailorAgentConf.URL = "http://af-sailor-agent:9595"
	}

	if config.KafkaConf.Version == "" {
		config.KafkaConf.Version = "2.3.1"
	}
	if config.KafkaConf.URI == "" {
		config.KafkaConf.URI = "kafka-headless.resource:9097"
	}
	if config.KafkaConf.ClientId == "" {
		config.KafkaConf.ClientId = "af.sailor-service"
	}
	if config.KafkaConf.GroupId == "" {
		config.KafkaConf.GroupId = "af.sailor-service"
	}

	if config.DepServicesConf.UserMgmPrivateHost == "" {
		config.DepServicesConf.UserMgmPrivateHost = "http://user-management-private:30980"
	}
	if config.DepServicesConf.MetaDataMgmHost == "" {
		config.DepServicesConf.MetaDataMgmHost = "http://metadata-manage:80"
	}
	if config.DepServicesConf.AuthServiceHost == "" {
		config.DepServicesConf.AuthServiceHost = "http://auth-service:8155"
	}
	if config.DepServicesConf.DataCatalogHost == "" {
		config.DepServicesConf.DataCatalogHost = "http://data-catalog:8153"
	}
	if config.DepServicesConf.DataSubjectHost == "" {
		config.DepServicesConf.DataSubjectHost = "http://data-subject:8123"
	}
	if config.DepServicesConf.BasicSearchHost == "" {
		config.DepServicesConf.BasicSearchHost = "http://basic-search:8163"
	}

	if config.OpenSearchConf.ReadUri == "" {
		config.OpenSearchConf.ReadUri = "http://opensearch-read.resource:9200"
	}

	if config.OpenSearchConf.WriteUri == "" {
		config.OpenSearchConf.WriteUri = "http://opensearch-write.resource:9200"
	}

	if config.DepServicesConf.ADPAgentFactoryHost == "" {
		config.DepServicesConf.ADPAgentFactoryHost = "http://agent-factory:13020"
	}
}
func init() {
	once.Do(func() {
		config = new(Config)
	})
	initAI()
}

type Config struct {
	Telemetry                   telemetry.Config            `json:"telemetry"`
	DBOptions                   options.DBOptions           `json:"database"`
	ServerConf                  ServerConf                  `json:"server"`
	SysConf                     SysConf                     `json:"sys"`
	DepServicesConf             DepServicesConf             `json:"depServices"`
	OpenAIConf                  OpenAIConf                  `json:"openAI"`
	AnyDataConf                 AnyDataConf                 `json:"anyDataConf"`
	OAuth                       OAuth                       `json:"oauth"`
	KnowledgeNetworkBuild       KnowledgeNetworkBuild       `json:"knowledgeNetworkBuild"`
	KnowledgeNetworkResourceMap KnowledgeNetworkResourceMap `json:"knowledgeNetworkResourceMap"`
	Redis                       RedisConf                   `json:"redis"`
	KafkaConf                   KafkaConf                   `json:"kafkaConf"`
	AfSailorConf                AfSailorConf                `json:"afSailorConf"`
	AfSailorAgentConf           AfSailorAgentConf           `json:"afSailorAgentConf"`
	ProtonConf                  ProtonConf                  `json:"protonConf"`
	GraphModelConfig            GraphModel                  `json:"graphModelConfig"`
	OpenSearchConf              OpenSearchConf              `json:"opensearch"`
}

type OpenSearchConf struct {
	ReadUri     string `json:"readUri,omitempty"`
	WriteUri    string `json:"writeUri,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Sniff       bool   `json:"sniff"`
	Healthcheck bool   `json:"healthcheck"`
	Debug       bool   `json:"debug"`
	Highlight   struct {
		PreTag  string `json:"preTag"`
		PostTag string `json:"postTag"`
	} `json:"highlight"`
}

type GraphModel struct {
	AfDatasourceID int `json:"af_datasource_id"`
	KnwID          int `json:"knw_id"`
}

type SysConfMode string

const (
	SysConfModeDebug   SysConfMode = "debug"
	SysConfModeRelease SysConfMode = "release"
)

type SysConf struct {
	Mode   SysConfMode `json:"mode"`   // debug、release
	SelfIP string      `json:"selfIP"` // 自身ip，读取POD_IP
}

type HttpConf struct {
	Addr string `json:"addr"`
}

type ServerConf struct {
	HttpConf `json:"http"`
	SwagConf `json:"doc"`
}

type SwagConf struct {
	Host    string `json:"host"`
	Version string `json:"version"`
}

type DepServicesConf struct {
	VirtualizationEngineUrl string `json:"virtualizationEngineUrl"`
	AnyRobotTraceUrl        string `json:"anyRobotTraceUrl"`
	ConfigCenterHost        string `json:"configCenterHost"`
	UserMgmPrivateHost      string `json:"userMgmPrivateHost"`
	AnyDataRecUrl           string `json:"anyDataRecUrl"`
	MetaDataMgmHost         string `json:"metaDataMgmHost"`
	AuthServiceHost         string `json:"authServiceHost"`
	DataCatalogHost         string `json:"dataCatalogHost"`
	DataSubjectHost         string `json:"dataSubjectHost"`
	BasicSearchHost         string `json:"basicSearchHost"`
	ADPAgentFactoryHost     string `json:"adpAgentFactoryHost"`
}

type OpenAIConf struct {
	APIKey     string `json:"apiKey"`
	URL        string `json:"url"`
	APIVersion string `json:"apiVersion"`
	APIType    string `json:"apiType"`
}

type AnyDataConf struct {
	URL         string `json:"url"`
	AccountType string `json:"accountType"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Version     string `json:"version"`
}

type AfSailorConf struct {
	URL string `json:"url"`
}

type AfSailorAgentConf struct {
	URL string `json:"url"`
}

type ProtonConf struct {
	ApplicationUrl string `json:"applicationUrl"`
}

func GetConfig() *Config {
	initAI()
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ResetConfig(conf *Config) {
	lock.Lock()
	defer lock.Unlock()
	config = conf
}

type OAuth struct {
	HydraAdmin   string `json:"hydraAdmin"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type GraphAnalysis struct {
	ID           string            `json:"id,omitempty"`
	Name         string            `json:"name,omitempty"`
	FilePath     string            `json:"filePath,omitempty"`
	GraphID      string            `json:"graphId,omitempty"`
	FileBaseName string            `json:"-"`
	EntityMap    map[string]string `json:"EntityMap,omitempty"`
	Version      int32             `json:"version,omitempty"`
}

type GraphBuildInfo struct {
	ID                 string `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	FilePath           string `json:"filePath,omitempty"`
	KnowledgeNetworkID string `json:"knowledgeNetworkId,omitempty"`
	OldDatasourceID    string `json:"oldDatasourceId,omitempty"`
	NewDatasourceID    string `json:"newDatasourceId,omitempty"`
	BuildTaskType      string `json:"buildTaskType"` // 构建任务类型，定时构建中使用全量构建还是增量构建，full、increment
	Version            int32  `json:"version,omitempty"`
}

type WeightItem struct {
	Service  string  `json:"service"`
	Relation string  `json:"relation"`
	Weight   float64 `json:"weight"`
}

type KnowledgeNetworkBuild struct {
	GraphScheduledCron string `json:"graphScheduledCron"` // 图谱更新定时策略，cron表达式
	FileBaseDir        string `json:"fileBaseDir"`
	KnowledgeNetwork   []struct {
		ID      string `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Version int32  `json:"version,omitempty"`
	} `json:"knowledgeNetwork,omitempty"`
	Datasource []struct {
		ID                 string `json:"id,omitempty"`
		OldID              string `json:"oldId,omitempty"` // 不为空时，表示图谱文件中的数据源id，即需要替换掉的数据源
		Name               string `json:"name,omitempty"`
		Source             string `json:"source,omitempty"`
		DataType           string `json:"dataType,omitempty"`
		Address            string `json:"address,omitempty"`
		Port               int    `json:"port,omitempty"`
		User               string `json:"user,omitempty"`
		Password           string `json:"password,omitempty"`
		Path               string `json:"path,omitempty"`
		ExtractType        string `json:"extractType,omitempty"`
		ConnectType        string `json:"connectType,omitempty"`
		KnowledgeNetworkID string `json:"knowledgeNetworkId,omitempty"`
		Version            int32  `json:"version,omitempty"`
	} `json:"datasource,omitempty"`
	Graph            []GraphBuildInfo `json:"graph,omitempty"`
	GraphAnalysis    []GraphAnalysis  `json:"graphAnalysis,omitempty"`
	SynonymsLexicons []struct {
		ID                 string `json:"id,omitempty"`
		Name               string `json:"name,omitempty"`
		FilePath           string `json:"filePath,omitempty"`
		KnowledgeNetworkID string `json:"knowledgeNetworkId,omitempty"`
		Version            int32  `json:"version,omitempty"`
	} `json:"synonymsLexicons,omitempty"`
	CognitiveService []struct {
		ID                    string `json:"id,omitempty"`
		Name                  string `json:"name,omitempty"`
		FilePath              string `json:"filePath,omitempty"`
		OldGraphID            string `json:"oldGraphId,omitempty"`
		OldLexiconID          string `json:"oldLexiconId,omitempty"`
		NewLexiconID          string `json:"newLexiconId,omitempty"`
		OldStopwordsLexiconID string `json:"oldStopwordsLexiconId,omitempty"`
		NewStopwordsLexiconID string `json:"newStopwordsLexiconId,omitempty"`
		GraphID               string `json:"graphId,omitempty"`
		Version               int32  `json:"version,omitempty"`
	} `json:"cognitiveService,omitempty"`
}

func (k KnowledgeNetworkBuild) FilePath(path string) string {
	return strings.TrimSuffix(k.FileBaseDir, string(os.PathSeparator)) + string(os.PathSeparator) + path
}

type KnowledgeNetworkResourceMap struct {
	AFBusinessRelationsGraphConfigId                        string `json:"afBusinessRelationsGraphConfigId"`
	LineageGraphConfigId                                    string `json:"lineageGraphConfigId"`
	DataAssetsGraphConfigId                                 string `json:"dataAssetsGraphConfigId"`
	SmartRecommendationGraphConfigId                        string `json:"smartRecommendationGraphConfigId"`
	LogicalViewGraphConfigId                                string `json:"logicalViewGraphConfigId"`
	CognitiveSearchDataCatalogGraphConfigId                 string `json:"cognitiveSearchDataCatalogGraphConfigId"`
	CognitiveSearchDataResourceGraphConfigId                string `json:"cognitiveSearchDataResourceGraphConfigId"`
	DataComprehensionGraphAnalysisServiceConfigId           string `json:"dataComprehensionGraphAnalysisServiceConfigId"`
	AssetGraphAnalysisServiceConfigId                       string `json:"assetGraphAnalysisServiceConfigId"`
	DataAssetCognitiveSearchGraphAnalysisConfigId           string `json:"dataAssetCognitiveSearchGraphAnalysisConfigId"`
	CognitiveSearchDataCatalogGraphAnalysisConfigId         string `json:"cognitiveSearchDataCatalogGraphAnalysisConfigId"`
	CognitiveSearchResourceGraphAnalysisConfigId            string `json:"cognitiveSearchResourceGraphAnalysisConfigId"`
	FieldStandardRecommendConfigId                          string `json:"fieldStandardRecommendConfigId"`
	FlowRecommendConfigId                                   string `json:"flowchartRecommendConfigId"`
	FormRecommendConfigId                                   string `json:"formRecommendConfigId"`
	StandardCheckConfigId                                   string `json:"standardCheckConfigId"`
	AssetSubgraphSearchConfigId                             string `json:"assetSubgraphSearchConfigId"`
	AssetSubgraphEntitySearchConfigId                       string `json:"assetSubgraphEntitySearchConfigId"`
	CognitiveSearchDataCatalogSubgraphSearchConfigId        string `json:"cognitiveSearchDataCatalogSubgraphSearchConfigId"`
	CognitiveSearchDataCatalogSubgraphEntitySearchConfigId  string `json:"cognitiveSearchDataCatalogSubgraphEntitySearchConfigId"`
	CognitiveSearchDataResourceSubgraphSearchConfigId       string `json:"cognitiveSearchDataResourceSubgraphSearchConfigId"`
	CognitiveSearchDataResourceSubgraphEntitySearchConfigId string `json:"cognitiveSearchDataResourceSubgraphEntitySearchConfigId"`
}

type RedisConf struct {
	Addr             string `json:"addr,omitempty"`
	ClientName       string `json:"clientName,omitempty"`
	DB               int    `json:"DB,omitempty"`
	Password         string `json:"password,omitempty"`
	SentinelPassword string `json:"sentinelPassword,omitempty"`
	MaxRetries       int    `json:"maxRetries,omitempty"`
	MinRetryBackoff  int    `json:"minRetryBackoff,omitempty"` // 单位毫秒
	MaxRetryBackoff  int    `json:"maxRetryBackoff,omitempty"` // 单位毫秒
	DialTimeout      int    `json:"dialTimeout,omitempty"`     // 单位毫秒
	ReadTimeout      int    `json:"readTimeout,omitempty"`     // 单位毫秒
	WriteTimeout     int    `json:"writeTimeout,omitempty"`    // 单位毫秒
	PoolSize         int    `json:"poolSize,omitempty"`
	PoolTimeout      int    `json:"poolTimeout,omitempty"` // 单位毫秒
	MinIdleConns     int    `json:"minIdleConns,omitempty"`
	MaxIdleConns     int    `json:"maxIdleConns,omitempty"`
	ConnMaxIdleTime  int    `json:"connMaxIdleTime,omitempty"` // 单位分钟
	ConnMaxLifetime  int    `json:"connMaxLifetime,omitempty"` // 单位分钟
	MasterName       string `json:"masterName,omitempty"`
}

type KafkaConf struct {
	Version  string `json:"version"`
	URI      string `json:"uri,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	GroupId  string `json:"groupId,omitempty"`
}

// 认知搜索权重配置

type CognitiveSearchItem struct {
	SearchType  string
	Name        string
	RequestName string
	Relation    string
	Weight      float64
}

func GetCognitiveSearchResourceConfig() []CognitiveSearchItem {
	cognitiveSearchConfig := []CognitiveSearchItem{
		{"data-resource", "resource", "resource", "", 4},
		{"data-resource", "response_field", "response_field", "包含", 1.9},
		{"data-resource", "field", "field", "包含", 1.8},
		{"data-resource", "data_explore_report", "data_explore_report", "包含", 1.8},
		{"data-resource", "data_owner", "dataowner", "管理", 1.3},
		{"data-resource", "department", "department", "管理", 1.3},
		{"data-resource", "domain", "domain", "包含", 1.3},
		{"data-resource", "subdomain", "subdomain", "包含", 1},
		{"data-resource", "datasource", "datasource", "包含", 1},
		{"data-resource", "metadataschema", "metadataschema", "包含", 1},
	}
	return cognitiveSearchConfig
}

func GetCognitiveSearchCatalogConfig() []CognitiveSearchItem {
	cognitiveSearchConfig := []CognitiveSearchItem{
		{"data-catalog", "catalogtag", "catalogtag", "打标", 1.5},
		{"data-catalog", "data_explore_report", "data_explore_report", "包含", 1.8},
		{"data-catalog", "form_view_field", "form_view_field", "包含", 1.8},
		{"data-catalog", "datacatalog", "datacatalog", "包含", 4},
		{"data-catalog", "data_owner", "dataowner", "管理", 1.4},
		{"data-catalog", "datasource", "datasource", "管理", 1},
		{"data-catalog", "department", "department", "包含", 1.3},
		{"data-catalog", "info_system", "info_system", "包含", 1.2},
		{"data-catalog", "form_view", "form_view", "包含", 1.9},
		{"data-catalog", "metadataschema", "metadataschema", "包含", 1},
	}
	return cognitiveSearchConfig
}
