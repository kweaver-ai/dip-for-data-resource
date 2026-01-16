package data_change_mq

import (
	"encoding/json"
	"fmt"
	"time"
)

type TDataCatalogBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				AuditAdvice      string         `json:"audit_advice"`
				AuditApplySn     int64          `json:"audit_apply_sn"`
				AuditState       int            `json:"audit_state"`
				Code             string         `json:"code"`
				CreatedAt        interface{}    `json:"created_at"`
				CreatorName      string         `json:"creator_name"`
				CreatorUid       string         `json:"creator_uid"`
				CurrentVersion   int            `json:"current_version"`
				DataKind         int            `json:"data_kind"`
				DataKindFlag     int            `json:"data_kind_flag"`
				DataRange        int            `json:"data_range"`
				DeleteName       string         `json:"delete_name"`
				DeleteUid        string         `json:"delete_uid"`
				DeletedAt        interface{}    `json:"deleted_at"`
				Description      string         `json:"description"`
				FileCount        int            `json:"file_count"`
				FlowApplyId      string         `json:"flow_apply_id"`
				FlowId           string         `json:"flow_id"`
				FlowName         string         `json:"flow_name"`
				FlowNodeId       string         `json:"flow_node_id"`
				FlowNodeName     string         `json:"flow_node_name"`
				FlowType         int            `json:"flow_type"`
				FlowVersion      string         `json:"flow_version"`
				ForwardVersionId int            `json:"forward_version_id"`
				GroupId          int            `json:"group_id"`
				GroupName        string         `json:"group_name"`
				Id               NumberOrString `json:"id"`
				IsCanceled       interface{}    `json:"is_canceled"`
				IsIndexed        int            `json:"is_indexed"`
				LabelFlag        int            `json:"label_flag"`
				OpenCondition    string         `json:"open_condition"`
				OpenType         int            `json:"open_type"`
				Orgcode          string         `json:"orgcode"`
				Orgname          string         `json:"orgname"`
				OwnerId          string         `json:"owner_id"`
				OwnerName        string         `json:"owner_name"`
				PhysicalDeletion interface{}    `json:"physical_deletion"`
				ProcDefKey       string         `json:"proc_def_key"`
				PublishFlag      int            `json:"publish_flag"`
				PublishedAt      interface{}    `json:"published_at"`
				RelCatalogFlag   int            `json:"rel_catalog_flag"`
				RelEventFlag     int            `json:"rel_event_flag"`
				SharedCondition  string         `json:"shared_condition"`
				SharedMode       int            `json:"shared_mode"`
				SharedType       int            `json:"shared_type"`
				Source           int            `json:"source"`
				SrcEventFlag     int            `json:"src_event_flag"`
				State            int            `json:"state"`
				SyncFrequency    string         `json:"sync_frequency"`
				SyncMechanism    int            `json:"sync_mechanism"`
				SystemFlag       int            `json:"system_flag"`
				TableCount       int            `json:"table_count"`
				TableType        int            `json:"table_type"`
				ThemeId          int            `json:"theme_id"`
				ThemeName        string         `json:"theme_name"`
				Title            string         `json:"title"`
				UpdateCycle      int            `json:"update_cycle"`
				UpdatedAt        interface{}    `json:"updated_at"`
				UpdaterName      string         `json:"updater_name"`
				UpdaterUid       string         `json:"updater_uid"`
				Version          string         `json:"version"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type NumberOrString struct {
	Value interface{} // 使用 interface{} 存储原始值
}

func (ns *NumberOrString) UnmarshalJSON(data []byte) error {
	// 尝试将JSON数据解析为JSON Number
	var num json.Number
	err := json.Unmarshal(data, &num)
	if err == nil {
		// 成功解析为数字，转换为 int64
		value, err := num.Int64()
		if err != nil {
			// 如果转换为 int64 失败，则保留为原始JSON Number（这通常不会发生，除非数字过大）
			ns.Value = num
		} else {
			ns.Value = value
		}
		return nil
	}

	// 如果不是JSON Number，则尝试解析为字符串
	var str string
	err = json.Unmarshal(data, &str)
	if err == nil {
		ns.Value = str
		return nil
	}

	// 如果两种解析都失败，则返回错误
	return fmt.Errorf("cannot unmarshal NumberOrString: %w", err)
}

type TDataCatalogBody2 struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				AuditAdvice      string      `json:"audit_advice"`
				AuditApplySn     int64       `json:"audit_apply_sn"`
				AuditState       int         `json:"audit_state"`
				Code             string      `json:"code"`
				CreatedAt        interface{} `json:"created_at"`
				CreatorName      string      `json:"creator_name"`
				CreatorUid       string      `json:"creator_uid"`
				CurrentVersion   int         `json:"current_version"`
				DataKind         int         `json:"data_kind"`
				DataKindFlag     int         `json:"data_kind_flag"`
				DataRange        int         `json:"data_range"`
				DeleteName       string      `json:"delete_name"`
				DeleteUid        string      `json:"delete_uid"`
				DeletedAt        interface{} `json:"deleted_at"`
				Description      string      `json:"description"`
				FileCount        int         `json:"file_count"`
				FlowApplyId      string      `json:"flow_apply_id"`
				FlowId           string      `json:"flow_id"`
				FlowName         string      `json:"flow_name"`
				FlowNodeId       string      `json:"flow_node_id"`
				FlowNodeName     string      `json:"flow_node_name"`
				FlowType         int         `json:"flow_type"`
				FlowVersion      string      `json:"flow_version"`
				ForwardVersionId int         `json:"forward_version_id"`
				GroupId          int         `json:"group_id"`
				GroupName        string      `json:"group_name"`
				Id               string      `json:"id"`
				IsCanceled       interface{} `json:"is_canceled"`
				IsIndexed        int         `json:"is_indexed"`
				LabelFlag        int         `json:"label_flag"`
				OpenCondition    string      `json:"open_condition"`
				OpenType         int         `json:"open_type"`
				Orgcode          string      `json:"orgcode"`
				Orgname          string      `json:"orgname"`
				OwnerId          string      `json:"owner_id"`
				OwnerName        string      `json:"owner_name"`
				PhysicalDeletion interface{} `json:"physical_deletion"`
				ProcDefKey       string      `json:"proc_def_key"`
				PublishFlag      int         `json:"publish_flag"`
				PublishedAt      interface{} `json:"published_at"`
				RelCatalogFlag   int         `json:"rel_catalog_flag"`
				RelEventFlag     int         `json:"rel_event_flag"`
				SharedCondition  string      `json:"shared_condition"`
				SharedMode       int         `json:"shared_mode"`
				SharedType       int         `json:"shared_type"`
				Source           int         `json:"source"`
				SrcEventFlag     int         `json:"src_event_flag"`
				State            int         `json:"state"`
				SyncFrequency    string      `json:"sync_frequency"`
				SyncMechanism    int         `json:"sync_mechanism"`
				SystemFlag       int         `json:"system_flag"`
				TableCount       int         `json:"table_count"`
				TableType        int         `json:"table_type"`
				ThemeId          int         `json:"theme_id"`
				ThemeName        string      `json:"theme_name"`
				Title            string      `json:"title"`
				UpdateCycle      int         `json:"update_cycle"`
				UpdatedAt        interface{} `json:"updated_at"`
				UpdaterName      string      `json:"updater_name"`
				UpdaterUid       string      `json:"updater_uid"`
				Version          string      `json:"version"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type ServiceBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				Id                 int         `json:"id"`
				ServiceName        string      `json:"service_name"`
				ServiceId          string      `json:"service_id"`
				ServiceCode        string      `json:"service_code"`
				ServicePath        string      `json:"service_path"`
				Status             string      `json:"status"`
				AuditType          string      `json:"audit_type"`
				AuditStatus        string      `json:"audit_status"`
				ApplyId            string      `json:"apply_id"`
				ProcDefKey         string      `json:"proc_def_key"`
				AuditAdvice        string      `json:"audit_advice"`
				BackendServiceHost string      `json:"backend_service_host"`
				BackendServicePath string      `json:"backend_service_path"`
				DepartmentId       string      `json:"department_id"`
				DepartmentName     string      `json:"department_name"`
				OwnerId            string      `json:"owner_id"`
				OwnerName          string      `json:"owner_name"`
				SubjectDomainId    string      `json:"subject_domain_id"`
				SubjectDomainName  string      `json:"subject_domain_name"`
				CreateModel        string      `json:"create_model"`
				HttpMethod         string      `json:"http_method"`
				ReturnType         string      `json:"return_type"`
				Protocol           string      `json:"protocol"`
				FileId             string      `json:"file_id"`
				Description        interface{} `json:"description"`
				DeveloperId        string      `json:"developer_id"`
				DeveloperName      string      `json:"developer_name"`
				RateLimiting       int         `json:"rate_limiting"`
				Timeout            int         `json:"timeout"`
				ServiceType        string      `json:"service_type"`
				FlowId             string      `json:"flow_id"`
				FlowName           string      `json:"flow_name"`
				FlowNodeId         string      `json:"flow_node_id"`
				FlowNodeName       string      `json:"flow_node_name"`
				OnlineTime         time.Time   `json:"online_time"`
				CreateTime         time.Time   `json:"create_time"`
				UpdateTime         time.Time   `json:"update_time"`
				DeleteTime         int         `json:"delete_time"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type ServiceParamBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				AliasName    string    `json:"alias_name"`
				CnName       string    `json:"cn_name"`
				CreateTime   time.Time `json:"create_time"`
				DataType     string    `json:"data_type"`
				DefaultValue string    `json:"default_value"`
				DeleteTime   int       `json:"delete_time"`
				Description  string    `json:"description"`
				EnName       string    `json:"en_name"`
				Id           int64     `json:"id"`
				Masking      string    `json:"masking"`
				Operator     string    `json:"operator"`
				ParamType    string    `json:"param_type"`
				Required     string    `json:"required"`
				Sequence     int       `json:"sequence"`
				ServiceId    string    `json:"service_id"`
				Sort         string    `json:"sort"`
				UpdateTime   time.Time `json:"update_time"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type IndicatorBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			ClassName string `json:"class_name"`
			TableName string `json:"table_name"`
			Entities  []struct {
				AnalysisDimension        string      `json:"analysis_dimension"`
				AtomicIndicatorId        interface{} `json:"atomic_indicator_id"`
				BusinessIndicatorId      string      `json:"business_indicator_id"`
				BusinessIndicatorName    string      `json:"business_indicator_name"`
				Code                     string      `json:"code"`
				CreatedAt                interface{} `json:"created_at"`
				CreatorName              string      `json:"creator_name"`
				CreatorUid               string      `json:"creator_uid"`
				DeletedAt                interface{} `json:"deleted_at"`
				DeleterName              string      `json:"deleter_name"`
				DeleterUid               string      `json:"deleter_uid"`
				Description              string      `json:"description"`
				DimensionModelId         int64       `json:"dimension_model_id"`
				Expression               string      `json:"expression"`
				Id                       int64       `json:"id"`
				IndicatorType            string      `json:"indicator_type"`
				IndicatorUnit            string      `json:"indicator_unit"`
				MgntDepId                string      `json:"mgnt_dep_id"`
				MgntDepName              string      `json:"mgnt_dep_name"`
				ModifierRestrict         interface{} `json:"modifier_restrict"`
				ModifierRestrictRelation string      `json:"modifier_restrict_relation"`
				Name                     string      `json:"name"`
				OwnerName                string      `json:"owner_name"`
				OwnerUid                 string      `json:"owner_uid"`
				ReferCount               int         `json:"refer_count"`
				SubjectDomainId          string      `json:"subject_domain_id"`
				TimeRestrict             interface{} `json:"time_restrict"`
				UpdateCycle              string      `json:"update_cycle"`
				UpdatedAt                interface{} `json:"updated_at"`
				UpdaterName              string      `json:"updater_name"`
				UpdaterUid               string      `json:"updater_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type DimensionModelBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			ClassName string `json:"class_name"`
			TableName string `json:"table_name"`
			Entities  []struct {
				Canvas         string      `json:"canvas"`
				CreatedAt      interface{} `json:"created_at"`
				CreatorName    string      `json:"creator_name"`
				CreatorUid     string      `json:"creator_uid"`
				DeletedAt      interface{} `json:"deleted_at"`
				DeleterName    string      `json:"deleter_name"`
				DeleterUid     string      `json:"deleter_uid"`
				Description    string      `json:"description"`
				DimModelConfig string      `json:"dim_model_config"`
				Id             int64       `json:"id"`
				Name           string      `json:"name"`
				ReferCount     int         `json:"refer_count"`
				UpdatedAt      interface{} `json:"updated_at"`
				UpdaterName    string      `json:"updater_name"`
				UpdaterUid     string      `json:"updater_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type SubjectDomainBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				CreatedAt    time.Time   `json:"created_at"`
				CreatedByUid string      `json:"created_by_uid"`
				DeletedAt    int         `json:"deleted_at"`
				Description  string      `json:"description"`
				DomainId     int64       `json:"domain_id"`
				Id           string      `json:"id"`
				LabelId      int         `json:"label_id"`
				Name         string      `json:"name"`
				Owners       interface{} `json:"owners"`
				Path         string      `json:"path"`
				PathId       string      `json:"path_id"`
				RefId        string      `json:"ref_id"`
				StandardId   int         `json:"standard_id"`
				Type         int         `json:"type"`
				Unique       int         `json:"unique"`
				UpdatedAt    time.Time   `json:"updated_at"`
				UpdatedByUid string      `json:"updated_by_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type DomainBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessSystem interface{} `json:"business_system"`
				CreatedAt      time.Time   `json:"created_at"`
				CreatedByUid   int         `json:"created_by_uid"`
				DeletedAt      int         `json:"deleted_at"`
				DepartmentId   interface{} `json:"department_id"`
				Description    interface{} `json:"description"`
				Id             string      `json:"id"`
				Level          int         `json:"level"`
				ModelId        string      `json:"model_id"`
				Name           string      `json:"name"`
				ParentType     int         `json:"parent_type"`
				Path           string      `json:"path"`
				PathId         string      `json:"path_id"`
				Pid            int64       `json:"pid"`
				Type           int         `json:"type"`
				UpdatedAt      time.Time   `json:"updated_at"`
				UpdatedByUid   int         `json:"updated_by_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type InfoSystemBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				InfoStstemId int64     `json:"info_ststem_id"`
				Name         string    `json:"name"`
				DeletedAt    int       `json:"deleted_at"`
				UpdatedAt    time.Time `json:"updated_at"`
				UpdatedByUid string    `json:"updated_by_uid"`
				Id           string    `json:"id"`
				Description  struct {
					String string      `json:"String"`
					Valid  interface{} `json:"Valid"`
				} `json:"description"`
				CreatedAt    time.Time `json:"created_at"`
				CreatedByUid string    `json:"created_by_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type BusinessModelBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessDomainId interface{} `json:"business_domain_id"`
				BusinessModelId  string      `json:"business_model_id"`
				CreatedAt        time.Time   `json:"created_at"`
				CreatedByUid     string      `json:"created_by_uid"`
				DeletedAt        int         `json:"deleted_at"`
				DepartmentId     interface{} `json:"department_id"`
				Description      string      `json:"description"`
				Id               int64       `json:"id"`
				MainBusinessId   string      `json:"main_business_id"`
				Name             string      `json:"name"`
				ProjectId        string      `json:"project_id"`
				Status           int         `json:"status"`
				TaskId           string      `json:"task_id"`
				UpdatedAt        time.Time   `json:"updated_at"`
				UpdatedByUid     string      `json:"updated_by_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type BusinessFlowchartBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessModelId string    `json:"business_model_id"`
				CreatedAt       time.Time `json:"created_at"`
				CreatedByUid    int       `json:"created_by_uid"`
				DeletedAt       int       `json:"deleted_at"`
				Description     string    `json:"description"`
				FilePathId      int64     `json:"file_path_id"`
				FlowchartId     string    `json:"flowchart_id"`
				FlowchartLevel  int       `json:"flowchart_level"`
				Id              int64     `json:"id"`
				MainBusinessId  string    `json:"main_business_id"`
				Name            string    `json:"name"`
				Path            string    `json:"path"`
				PathId          string    `json:"path_id"`
				Status          int       `json:"status"`
				UpdatedAt       time.Time `json:"updated_at"`
				UpdatedByUid    int       `json:"updated_by_uid"`
				Version         int       `json:"version"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type BusinessFlowchartComponentBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				ComponentId   string    `json:"component_id"`
				CreatedAt     time.Time `json:"created_at"`
				Description   string    `json:"description"`
				DiagramId     string    `json:"diagram_id"`
				DiagramName   string    `json:"diagram_name"`
				FlowchartId   string    `json:"flowchart_id"`
				Function      string    `json:"function"`
				Id            int       `json:"id"`
				MxcellId      string    `json:"mxcell_id"`
				Name          string    `json:"name"`
				Parent        string    `json:"parent"`
				Source        string    `json:"source"`
				Target        string    `json:"target"`
				Type          int       `json:"type"`
				Value         string    `json:"value"`
				ValueRelation string    `json:"value_relation"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type FormViewBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			ClassName string `json:"class_name"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessName            string      `json:"business_name"`
				CreatedAt               time.Time   `json:"created_at"`
				CreatedByUid            string      `json:"created_by_uid"`
				DatasourceId            string      `json:"datasource_id"`
				DeletedAt               int         `json:"deleted_at"`
				DepartmentId            interface{} `json:"department_id"`
				Description             interface{} `json:"description"`
				EditStatus              int         `json:"edit_status"`
				ExploreJobId            interface{} `json:"explore_job_id"`
				ExploreJobVersion       interface{} `json:"explore_job_version"`
				ExploreTimestampId      interface{} `json:"explore_timestamp_id"`
				ExploreTimestampVersion interface{} `json:"explore_timestamp_version"`
				FormViewId              int64       `json:"form_view_id"`
				Id                      string      `json:"id"`
				MetadataFormId          string      `json:"metadata_form_id"`
				OwnerId                 interface{} `json:"owner_id"`
				PublishAt               interface{} `json:"publish_at"`
				SceneAnalysisId         string      `json:"scene_analysis_id"`
				Status                  int         `json:"status"`
				SubjectId               interface{} `json:"subject_id"`
				TechnicalName           string      `json:"technical_name"`
				Type                    int         `json:"type"`
				UniformCatalogCode      string      `json:"uniform_catalog_code"`
				UpdatedAt               time.Time   `json:"updated_at"`
				UpdatedByUid            string      `json:"updated_by_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type BusinessFormStandardBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessFormId       string      `json:"business_form_id"`
				BusinessModelId      string      `json:"business_model_id"`
				CreatedAt            time.Time   `json:"created_at"`
				CreatedByUid         int         `json:"created_by_uid"`
				DataKind             int         `json:"data_kind"`
				DeletedAt            int         `json:"deleted_at"`
				Description          string      `json:"description"`
				FieldCount           int         `json:"field_count"`
				FieldFiller          string      `json:"field_filler"`
				FieldFillingDate     interface{} `json:"field_filling_date"`
				Filler               string      `json:"filler"`
				FillingDate          interface{} `json:"filling_date"`
				FormType             int         `json:"form_type"`
				FromTableId          string      `json:"from_table_id"`
				Guideline            string      `json:"guideline"`
				Id                   int64       `json:"id"`
				MapTableId           string      `json:"map_table_id"`
				Name                 string      `json:"name"`
				OpenAttribute        int         `json:"open_attribute"`
				OpenCondition        string      `json:"open_condition"`
				Published            interface{} `json:"published"`
				RefBusinessObject    string      `json:"ref_business_object"`
				RelatedBusinessScene interface{} `json:"related_business_scene"`
				ResourceTag          interface{} `json:"resource_tag"`
				SharedAttribute      int         `json:"shared_attribute"`
				SharedCondition      string      `json:"shared_condition"`
				SharedMode           int         `json:"shared_mode"`
				Source               int         `json:"source"`
				SourceBusinessScene  interface{} `json:"source_business_scene"`
				SourceSystem         interface{} `json:"source_system"`
				UpdateCycle          interface{} `json:"update_cycle"`
				UpdatedAt            time.Time   `json:"updated_at"`
				UpdatedByUid         int         `json:"updated_by_uid"`
				Warn                 interface{} `json:"warn"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type BusinessFormFieldStandardBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessFormId              string      `json:"business_form_id"`
				BusinessFormName            string      `json:"business_form_name"`
				CodeTable                   string      `json:"code_table"`
				ConfidentialAttribute       int         `json:"confidential_attribute"`
				CreatedAt                   time.Time   `json:"created_at"`
				CreatedByUid                int         `json:"created_by_uid"`
				DataAccuracy                interface{} `json:"data_accuracy"`
				DataLength                  interface{} `json:"data_length"`
				DataType                    string      `json:"data_type"`
				DeletedAt                   int         `json:"deleted_at"`
				Description                 string      `json:"description"`
				EncodingRule                string      `json:"encoding_rule"`
				FieldId                     string      `json:"field_id"`
				FieldRelationship           string      `json:"field_relationship"`
				FormulateBasis              int         `json:"formulate_basis"`
				Id                          int64       `json:"id"`
				IsCurrentBusinessGeneration interface{} `json:"is_current_business_generation"`
				IsIncrementalField          interface{} `json:"is_incremental_field"`
				IsPrimaryKey                interface{} `json:"is_primary_key"`
				IsRequired                  interface{} `json:"is_required"`
				IsStandardizationRequired   interface{} `json:"is_standardization_required"`
				LabelId                     string      `json:"label_id"`
				Name                        string      `json:"name"`
				NameEn                      string      `json:"name_en"`
				OpenAttribute               int         `json:"open_attribute"`
				OpenCondition               string      `json:"open_condition"`
				RefId                       string      `json:"ref_id"`
				SensitiveAttribute          int         `json:"sensitive_attribute"`
				SharedAttribute             int         `json:"shared_attribute"`
				SharedCondition             string      `json:"shared_condition"`
				SourceId                    string      `json:"source_id"`
				SourceIdPublished           string      `json:"source_id_published"`
				StandardId                  string      `json:"standard_id"`
				StandardStatus              int         `json:"standard_status"`
				Unit                        string      `json:"unit"`
				UpdatedAt                   time.Time   `json:"updated_at"`
				UpdatedByUid                int         `json:"updated_by_uid"`
				ValueRange                  string      `json:"value_range"`
				ValueRangeType              int         `json:"value_range_type"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type StandardInfoBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type  string `json:"type"`
		Graph string `json:"graph"`
		Name  string `json:"name"`
		Model struct {
			Id         int64  `json:"id"`
			Name       string `json:"name"`
			NameEn     string `json:"name_en"`
			DataType   string `json:"data_type"`
			DataLength struct {
				Int32 int         `json:"Int32"`
				Valid interface{} `json:"Valid"`
			} `json:"data_length"`
			DataAccuracy struct {
				Int32 int         `json:"Int32"`
				Valid interface{} `json:"Valid"`
			} `json:"data_accuracy"`
			ValueRange     string `json:"value_range"`
			FormulateBasis int    `json:"formulate_basis"`
			CodeTable      string `json:"code_table"`
		} `json:"model"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"payload"`
}

type FormViewFieldBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			ClassName string `json:"class_name"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessName      string      `json:"business_name"`
				BusinessTimestamp interface{} `json:"business_timestamp"`
				CodeTableId       interface{} `json:"code_table_id"`
				DataAccuracy      int         `json:"data_accuracy"`
				DataLength        int         `json:"data_length"`
				DataType          string      `json:"data_type"`
				DeletedAt         int         `json:"deleted_at"`
				FormViewFieldId   int64       `json:"form_view_field_id"`
				FormViewId        string      `json:"form_view_id"`
				Id                string      `json:"id"`
				IsNullable        string      `json:"is_nullable"`
				OriginalDataType  string      `json:"original_data_type"`
				PrimaryKey        interface{} `json:"primary_key"`
				Standard          interface{} `json:"standard"`
				StandardCode      interface{} `json:"standard_code"`
				Status            int         `json:"status"`
				TechnicalName     string      `json:"technical_name"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type TReportBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				FCode           string      `json:"f_code"`
				FCreatedAt      time.Time   `json:"f_created_at"`
				FCreatedByUid   interface{} `json:"f_created_by_uid"`
				FCreatedByUname interface{} `json:"f_created_by_uname"`
				FDvTaskId       string      `json:"f_dv_task_id"`
				FExploreType    int         `json:"f_explore_type"`
				FFinishedAt     time.Time   `json:"f_finished_at"`
				FId             int64       `json:"f_id"`
				FLatest         int         `json:"f_latest"`
				FQueryParams    string      `json:"f_query_params"`
				FReason         string      `json:"f_reason"`
				FResult         interface{} `json:"f_result"`
				FSchema         string      `json:"f_schema"`
				FStatus         int         `json:"f_status"`
				FTable          string      `json:"f_table"`
				FTableId        string      `json:"f_table_id"`
				FTaskId         int64       `json:"f_task_id"`
				FTaskVersion    int         `json:"f_task_version"`
				FTotalNum       interface{} `json:"f_total_num"`
				FTotalSample    int         `json:"f_total_sample"`
				FTotalScore     interface{} `json:"f_total_score"`
				FVeCatalog      string      `json:"f_ve_catalog"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type TDataCatalogInfo struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				CatalogId int64  `json:"catalog_id"`
				Id        int64  `json:"id"`
				InfoKey   string `json:"info_key"`
				InfoType  int    `json:"info_type"`
				InfoValue string `json:"info_value"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type BusinessIndicatorBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessModelId    string    `json:"business_model_id"`
				CalculationFormula string    `json:"calculation_formula"`
				Code               string    `json:"code"`
				CreatedAt          time.Time `json:"created_at"`
				CreatorName        string    `json:"creator_name"`
				CreatorUid         string    `json:"creator_uid"`
				DeletedAt          int       `json:"deleted_at"`
				DeleterName        string    `json:"deleter_name"`
				DeleterUid         string    `json:"deleter_uid"`
				Description        string    `json:"description"`
				Id                 int64     `json:"id"`
				IndicatorId        string    `json:"indicator_id"`
				Name               string    `json:"name"`
				StatisticalCaliber string    `json:"statistical_caliber"`
				StatisticsCycle    string    `json:"statistics_cycle"`
				Unit               string    `json:"unit"`
				UpdatedAt          time.Time `json:"updated_at"`
				UpdaterName        string    `json:"updater_name"`
				UpdaterUid         string    `json:"updater_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type DataSourceBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				Password     string    `json:"password"`
				CreatedByUid string    `json:"created_by_uid"`
				CreatedAt    time.Time `json:"created_at"`
				UpdatedAt    time.Time `json:"updated_at"`
				Id           string    `json:"id"`
				Port         int       `json:"port"`
				Schema       string    `json:"schema"`
				UpdatedByUid string    `json:"updated_by_uid"`
				Host         string    `json:"host"`
				CatalogName  string    `json:"catalog_name"`
				Type         int       `json:"type"`
				DatabaseName string    `json:"database_name"`
				DataSourceId int       `json:"data_source_id"`
				Name         string    `json:"name"`
				Username     string    `json:"username"`
				SourceType   int       `json:"source_type"`
				TypeName     string    `json:"type_name"`
				InfoSystemId struct {
					String string      `json:"String"`
					Valid  interface{} `json:"Valid"`
				} `json:"info_system_id"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type TDataElementInfo struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				State         interface{} `json:"state"`
				DisableReason interface{} `json:"disableReason"`
				Version       int         `json:"version"`
				Deleted       interface{} `json:"deleted"`
				AuthorityId   string      `json:"authorityId"`
				CreateTime    string      `json:"createTime"`
				CreateUser    string      `json:"createUser"`
				UpdateTime    string      `json:"updateTime"`
				UpdateUser    string      `json:"updateUser"`
				Id            string      `json:"id"`
				Code          string      `json:"code"`
				NameEn        string      `json:"nameEn"`
				NameCn        string      `json:"nameCn"`
				Synonym       string      `json:"synonym"`
				StdType       int         `json:"stdType"`
				DataType      int         `json:"dataType"`
				DataLength    int         `json:"dataLength"`
				DataPrecision int         `json:"dataPrecision"`
				DictCode      string      `json:"dictCode"`
				Description   string      `json:"description"`
				CatalogId     string      `json:"catalogId"`
				RuleId        interface{} `json:"ruleId"`
				LabelId       interface{} `json:"labelId"`
			} `json:"entities"`
			UpdatedAt interface{} `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type TBusinessIndicator struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessModelId    string    `json:"business_model_id"`
				CalculationFormula string    `json:"calculation_formula"`
				Code               string    `json:"code"`
				CreatedAt          time.Time `json:"created_at"`
				CreatorName        string    `json:"creator_name"`
				CreatorUid         string    `json:"creator_uid"`
				DeletedAt          int       `json:"deleted_at"`
				DeleterName        string    `json:"deleter_name"`
				DeleterUid         string    `json:"deleter_uid"`
				Description        string    `json:"description"`
				Id                 int64     `json:"id"`
				IndicatorId        string    `json:"indicator_id"`
				Name               string    `json:"name"`
				NameEn             string    `json:"name_en"`
				StatisticalCaliber string    `json:"statistical_caliber"`
				StatisticsCycle    string    `json:"statistics_cycle"`
				Unit               string    `json:"unit"`
				UpdatedAt          time.Time `json:"updated_at"`
				UpdaterName        string    `json:"updater_name"`
				UpdaterUid         string    `json:"updater_uid"`
				VersionId          string    `json:"version_id"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type TRule struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				State         *string     `json:"state,omitempty"`
				DisableReason interface{} `json:"disableReason"`
				Version       interface{} `json:"version"`
				Deleted       interface{} `json:"deleted"`
				CreateTime    string      `json:"createTime"`
				CreateUser    string      `json:"createUser"`
				UpdateTime    string      `json:"updateTime"`
				UpdateUser    string      `json:"updateUser"`
				Id            int64       `json:"id"`
				Name          string      `json:"name"`
				CatalogId     int         `json:"catalogId"`
				OrgType       int         `json:"orgType"`
				Description   interface{} `json:"description"`
				RuleType      string      `json:"ruleType"`
				Expression    string      `json:"expression"`
				DepartmentIds string      `json:"departmentIds"`
				ThirdDeptId   string      `json:"thirdDeptId"`
			} `json:"entities"`
			UpdatedAt interface{} `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type TLabelCategory struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				CreatedAt     time.Time `json:"created_at"`
				CreatorName   string    `json:"creator_name"`
				CreatorUid    string    `json:"creator_uid"`
				DeletedAt     int       `json:"deleted_at"`
				FAuditStatus  string    `json:"f_audit_status"`
				FDescription  string    `json:"f_description"`
				FRangeTypeKey string    `json:"f_range_type_key"`
				FSort         int       `json:"f_sort"`
				FState        *int      `json:"f_state"`
				Id            int64     `json:"id"`
				Name          string    `json:"name"`
				UpdatedAt     time.Time `json:"updated_at"`
				UpdaterName   string    `json:"updater_name"`
				UpdaterUid    string    `json:"updater_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type TLabel struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				CategoryId  int64     `json:"category_id"`
				CreatedAt   time.Time `json:"created_at"`
				CreatorName string    `json:"creator_name"`
				CreatorUid  string    `json:"creator_uid"`
				DeletedAt   int       `json:"deleted_at"`
				FSort       int       `json:"f_sort"`
				Id          int64     `json:"id"`
				Name        string    `json:"name"`
				Path        string    `json:"path"`
				Pid         int64     `json:"pid"`
				UpdatedAt   time.Time `json:"updated_at"`
				UpdaterName string    `json:"updater_name"`
				UpdaterUid  string    `json:"updater_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type UserBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				Name   string `json:"name"`
				Status int    `json:"status"`
				Id     string `json:"id"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type ObjectBody struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				Id        string    `json:"id"`
				PathId    string    `json:"path_id"`
				Path      string    `json:"path"`
				CreatedAt time.Time `json:"created_at"`
				UpdatedAt time.Time `json:"updated_at"`
				Name      string    `json:"name"`
				Type      int       `json:"type"`
				Attribute string    `json:"attribute"`
				DeletedAt int       `json:"deleted_at"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type DAnalysisDims struct {
	AnalysisDimList []*DAnalysisDim
}

type DAnalysisDim struct {
	TableId       string `json:"table_id"`
	FieldId       string `json:"field_id"`
	BusinessName  string `json:"business_name"`
	TechnicalName string `json:"technical_name"`
	DataType      string `json:"data_type"`
}

type TModelLabelRecRel struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				CreatedAt       time.Time `json:"created_at"`
				CreatedName     string    `json:"created_name"`
				CreatorUid      string    `json:"creator_uid"`
				DeletedAt       int       `json:"deleted_at"`
				Description     string    `json:"description"`
				Id              int64     `json:"id"`
				Name            string    `json:"name"`
				RelatedModelIds string    `json:"related_model_ids"`
				UpdatedAt       time.Time `json:"updated_at"`
				UpdaterName     string    `json:"updater_name"`
				UpdaterUid      string    `json:"updater_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}

type TGraphModel struct {
	Header struct {
	} `json:"header"`
	Payload struct {
		Type    string `json:"type"`
		Content struct {
			Type      string `json:"type"`
			TableName string `json:"table_name"`
			Entities  []struct {
				BusinessName  string    `json:"business_name"`
				CatalogId     int       `json:"catalog_id"`
				CreatedAt     time.Time `json:"created_at"`
				CreatorUid    string    `json:"creator_uid"`
				DataViewId    string    `json:"data_view_id"`
				DeletedAt     int       `json:"deleted_at"`
				Description   string    `json:"description"`
				GradeLabelId  string    `json:"grade_label_id"`
				GraphId       int       `json:"graph_id"`
				Id            string    `json:"id"`
				ModelId       int       `json:"model_id"`
				ModelType     int       `json:"model_type"`
				SubjectId     string    `json:"subject_id"`
				TechnicalName string    `json:"technical_name"`
				UpdatedAt     time.Time `json:"updated_at"`
				UpdaterName   string    `json:"updater_name"`
				UpdaterUid    string    `json:"updater_uid"`
			} `json:"entities"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"content"`
	} `json:"payload"`
}
