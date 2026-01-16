package data_change_mq

import (
	"context"
	"strconv"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"github.com/pkg/errors"
)

func (c *consumer) createBRGEntityBusinessDomain(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityBusinessDomain(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":          entityInfo.Id,
			"name":        entityInfo.Name,
			"description": entityInfo.Description,
			"owners":      entityInfo.Owners,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_business_domain")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityThemeDomain(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityThemeDomain(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":          entityInfo.Id,
			"name":        entityInfo.Name,
			"description": entityInfo.Description,
			"owners":      entityInfo.Owners,
		},
	}

	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_theme_domain")
	if err != nil {
		return err
	}

	// 主题域分组
	entityThemeDomain2EntityBusinessDomain, err := c.dbRepo.GetBRGEdgeEntityThemeDomain2EntityBusinessDomain(ctx, entityId)
	if err != nil {
		return err
	}
	if len(entityThemeDomain2EntityBusinessDomain) > 0 {
		sideInfo := []map[string]map[string]string{}
		for _, item := range entityThemeDomain2EntityBusinessDomain {
			addItem := map[string]map[string]string{
				"start":     {"id": entityId, "_start_entity": "entity_theme_domain"},
				"end":       {"id": item.ThemeId, "_end_entity": "entity_business_domain"},
				"edge_pros": {},
			}
			sideInfo = append(sideInfo, addItem)
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entityThemeDomain2EntityBusinessDomain")
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *consumer) createBRGEntityBusinessObject(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityBusinessObject(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":          entityInfo.Id,
			"name":        entityInfo.Name,
			"description": entityInfo.Description,
			"owners":      entityInfo.Owners,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_business_object")
	if err != nil {
		return err
	}

	// 主题域
	entityBusinessObject2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessObject2EntityThemeDomain(ctx, entityId)
	if err != nil {
		return err
	}
	if len(entityBusinessObject2EntityThemeDomain) > 0 {
		sideInfo := []map[string]map[string]string{}
		for _, item := range entityBusinessObject2EntityThemeDomain {
			addItem := map[string]map[string]string{
				"start":     {"id": entityId, "_start_entity": "entity_business_object"},
				"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
				"edge_pros": {},
			}
			sideInfo = append(sideInfo, addItem)
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_object_2_entity_theme_domain")
		if err != nil {
			return err
		}
	}

	// 业务表
	entityBusinessObject2EntityBusinessForm, err := c.dbRepo.GetBRGEdgeEntityBusinessObject2EntityBusinessForm(ctx, entityId)
	if err != nil {
		return err
	}
	if len(entityBusinessObject2EntityBusinessForm) > 0 {
		sideInfo := []map[string]map[string]string{}
		for _, item := range entityBusinessObject2EntityBusinessForm {
			addItem := map[string]map[string]string{
				"start":     {"id": entityId, "_start_entity": "entity_business_object"},
				"end":       {"id": item.BusinessFormId, "_end_entity": "entity_business_form_standard"},
				"edge_pros": {},
			}
			sideInfo = append(sideInfo, addItem)
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_object_2_entity_business_form")
		if err != nil {
			return err
		}
	}

	// 数据资源目录
	entityBusinessObject2EntityDataCatalog, err := c.dbRepo.GetBRGEdgeEntityBusinessObject2EntityDataCatalog(ctx, entityId)
	if err != nil {
		return err
	}
	if len(entityBusinessObject2EntityDataCatalog) > 0 {
		sideInfo := []map[string]map[string]string{}
		for _, item := range entityBusinessObject2EntityDataCatalog {
			addItem := map[string]map[string]string{
				"start":     {"id": entityId, "_start_entity": "entity_business_object"},
				"end":       {"id": item.CatalogId, "_end_entity": "entity_data_catalog"},
				"edge_pros": {},
			}
			sideInfo = append(sideInfo, addItem)
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_object_2_entity_data_catalog")
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *consumer) createBRGEntityDataCatalog(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityDataCatalog(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":           entityInfo.Id,
			"code":         entityInfo.Code,
			"title":        entityInfo.Title,
			"group_id":     entityInfo.GroupId,
			"group_name":   entityInfo.GroupName,
			"theme_id":     entityInfo.ThemeId,
			"theme_name":   entityInfo.ThemeName,
			"description":  entityInfo.Description,
			"data_range":   entityInfo.DataRange,
			"update_cycle": entityInfo.UpdateCycle,
			"data_kind":    entityInfo.DataKind,
			"orgcode":      entityInfo.OrgCode,
			"orgname":      entityInfo.OrgName,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_data_catalog")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityInfoSystem(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityInfoSystem(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":        entityInfo.Id,
			"name":      entityInfo.Name,
			"path":      entityInfo.Path,
			"attribute": entityInfo.Attribute,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_info_system")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityBusinessScene(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityBusinessScene(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":        entityInfo.Id,
			"name":      entityInfo.Name,
			"path":      entityInfo.Path,
			"attribute": entityInfo.Attribute,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_business_scene")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityDataCatalogColumn(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityDataCatalogColumn(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":             entityInfo.Id,
			"catalog_id":     entityInfo.CatalogId,
			"column_name":    entityInfo.ColumnName,
			"name_cn":        entityInfo.NameCn,
			"description":    entityInfo.Description,
			"data_format":    entityInfo.DataFormat,
			"data_length":    entityInfo.DataLength,
			"datameta_id":    entityInfo.DatametaId,
			"datameta_name":  entityInfo.DatametaName,
			"ranges":         entityInfo.Ranges,
			"codeset_id":     entityInfo.CodesetId,
			"codeset_name":   entityInfo.CodesetName,
			"timestamp_flag": entityInfo.TimestampFlag,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_data_catalog_column")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntitySourceTable(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntitySourceTable(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":                entityInfo.Id,
			"name":              entityInfo.Name,
			"description":       entityInfo.Description,
			"schema_name":       entityInfo.SchemaName,
			"ve_catalog_id":     entityInfo.VeCatalogId,
			"metadata_table_id": entityInfo.MetadataTableId,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_source_table")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityDepartment(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityDepartment(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":   entityInfo.Id,
			"name": entityInfo.Name,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_department")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityStandardTable(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityStandardTable(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":                entityInfo.Id,
			"name":              entityInfo.Name,
			"description":       entityInfo.Description,
			"metadata_table_id": entityInfo.MetadataTableId,
			"schema_name":       entityInfo.SchemaName,
			"ve_catalog_id":     entityInfo.VeCatalogId,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_standard_table")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityBusinessFormStandard(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityBusinessFormStandard(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"id":                entityInfo.Id,
			"name":              entityInfo.Name,
			"description":       entityInfo.Description,
			"business_model_id": entityInfo.BusinessModelId,
			"guideline":         entityInfo.Guideline,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_business_form_standard")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityField(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityField(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"field_id":           entityInfo.FieldId,
			"business_form_id":   entityInfo.BusinessFormId,
			"business_form_name": entityInfo.BusinessFormName,
			"name":               entityInfo.Name,
			"name_en":            entityInfo.NameEn,
			"data_type":          entityInfo.DataType,
			"data_length":        entityInfo.DataLength,
			"value_range":        entityInfo.ValueRange,
			"field_relationship": entityInfo.FieldRelationship,
			"ref_id":             entityInfo.RefId,
			"standard_id":        entityInfo.StandardId,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_field")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityBusinessModel(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityBusinessModel(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"business_model_id": entityInfo.BusinessModelId,
			"main_business_id":  entityInfo.MainBusinessId,
			"name":              entityInfo.Name,
			"description":       entityInfo.Description,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_business_model")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityBusinessIndicator(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityBusinessIndicator(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"business_indicator_id": entityInfo.BusinessIndicatorId,
			"indicator_id":          entityInfo.IndicatorId,
			"business_model_id":     entityInfo.BusinessModelId,
			"name":                  entityInfo.Name,
			"desc":                  entityInfo.Desc,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_business_indicator")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityFlowchart(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityFlowchart(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"flowchart_id":      entityInfo.FlowchartId,
			"name":              entityInfo.Name,
			"description":       entityInfo.Description,
			"business_model_id": entityInfo.BusinessModelId,
			"main_business_id":  entityInfo.MainBusinessId,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_flowchart")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createBRGEntityFlowchartNode(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	entityInfo, err := c.dbRepo.GetBRGEntityFlowchartNode(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"node_id":      entityInfo.NodeId,
			"flowchart_id": entityInfo.FlowchartId,
			"diagram_id":   entityInfo.DiagramId,
			"diagram_name": entityInfo.DiagramName,
			"name":         entityInfo.Name,
			"description":  entityInfo.Description,
			"target":       entityInfo.Target,
			"source":       entityInfo.Source,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "entity_flowchart_node")
	if err != nil {
		return err
	}

	// 主题域
	//entityBusinessDomain2EntityThemeDomain, err := c.dbRepo.GetBRGEdgeEntityBusinessDomain2EntityThemeDomain(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(entityBusinessDomain2EntityThemeDomain) > 0 {
	//	sideInfo := []map[string]map[string]string{}
	//	for _, item := range entityBusinessDomain2EntityThemeDomain {
	//		addItem := map[string]map[string]string{
	//			"start":     {"id": entityId, "_start_entity": "entity_business_domain"},
	//			"end":       {"id": item.ThemeId, "_end_entity": "entity_theme_domain"},
	//			"edge_pros": {},
	//		}
	//		sideInfo = append(sideInfo, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "entity_business_domain_2_entity_theme_domain")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}
