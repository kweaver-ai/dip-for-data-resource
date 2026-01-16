package data_change_mq

import (
	"context"
	"strconv"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"github.com/pkg/errors"
)

func (c *consumer) createDataCatalog(ctx context.Context, graphConfigName string, entityId string, flowId string) (int, error) {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()
	state := 0

	//fmt.Println(settings.GetConfig().KnowledgeNetworkResourceMap.DataAssetsGraphConfigId)
	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return state, err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return state, err
	}

	dataCatalogInfo, err := c.dbRepo.GetEmptyDataCatalog()
	if entityId != "0" {
		dataCatalogInfo, err = c.dbRepo.GetDataCatalog(ctx, entityId)
	} else if flowId != "" {
		dataCatalogInfo, err = c.dbRepo.GetDataCatalogByFlowId(ctx, flowId)
	} else {
		return state, errors.Wrap(err, "can not find entity info "+entityId)
	}

	if err != nil {
		return state, err
	}

	if dataCatalogInfo == nil || dataCatalogInfo.Sid == "" {
		errors.Wrap(err, "can not find entity info "+entityId)
		state = 2
		return state, err
	}

	mainSearchCfg, err := c.adProxy.GetSearchConfig("datacatalog", "datacatalogid", entityId)
	mainEntityInfo, err := c.adProxy.FulltextSearchV2(ctx, graphId, "", mainSearchCfg)

	oInfoSystemId := ""
	oDepartmentId := ""
	oOwnerId := ""
	if len(mainEntityInfo.Res.Nodes) > 0 {
		node := mainEntityInfo.Res.Nodes[0]
		if len(node.Properties) > 0 {
			for _, property := range node.Properties[0].Props {
				if property.Name == "owner_id" {
					oOwnerId = property.Value
				}
				if property.Name == "info_system_id" {
					oInfoSystemId = property.Value
				}
				if property.Name == "department_id" {
					oDepartmentId = property.Value
				}
			}
		}

	}
	log.Infof("entity datacatlog graph property InfoSystemId:%s, DepartmentId:%s, OwnerId:%s", oInfoSystemId, oDepartmentId, oOwnerId)

	graphData := []map[string]any{
		{
			"ves_catalog_name":   dataCatalogInfo.VesCatalogName,
			"department_path":    dataCatalogInfo.DepartmentPath,
			"department_path_id": dataCatalogInfo.DepartmentPathId,
			"info_system_id":     dataCatalogInfo.InfoSystemId,
			"owner_id":           dataCatalogInfo.OwnerId,
			"info_system_name":   dataCatalogInfo.InfoSystemName,
			"department_id":      dataCatalogInfo.DepartmentId,
			"color":              dataCatalogInfo.Color,
			"department":         dataCatalogInfo.Department,
			"data_owner":         dataCatalogInfo.OwnerName,
			"metadata_schema":    dataCatalogInfo.MetadataSchema,
			"datasource":         dataCatalogInfo.Datasource,
			"update_cycle":       dataCatalogInfo.UpdateCycle,
			"published_at":       dataCatalogInfo.PublishedAt,
			"shared_type":        dataCatalogInfo.SharedType,
			"data_kind":          dataCatalogInfo.DataKind,
			"code":               dataCatalogInfo.Code,
			"asset_type":         dataCatalogInfo.AssetType,
			"description_name":   dataCatalogInfo.Description,
			"datacatalogname":    dataCatalogInfo.Name,
			"datacatalogid":      dataCatalogInfo.Sid,
		},
	}

	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "datacatalog")
	if err != nil {
		return state, err
	}

	entityId = dataCatalogInfo.Sid

	// 信息系统
	if dataCatalogInfo.InfoSystemId != "" {
		//searchCfg, err := c.adProxy.GetSearchConfig("info_system", "infosystemid", dataCatalogInfo.InfoSystemId)
		//entityInfo, err := c.adProxy.FulltextSearchV2(ctx, graphId, "", searchCfg)
		//if len(entityInfo.Res.Nodes) == 0 {
		//	infoSystem, err := c.dbRepo.GetInfoSystem(ctx, entityId)
		//	if err != nil {
		//		return err
		//	}
		//	if infoSystem == nil {
		//		errors.Wrap(err, "can not find data_view "+entityId)
		//		return err
		//	}
		//	infoSystemData := []map[string]any{
		//		{
		//			"infosystemid":     infoSystem.SysSid,
		//			"infosystemna":     infoSystem.InfoSystemName,
		//			"info_system_uuid": infoSystem.InfoSystemUuid,
		//		},
		//	}
		//	_, err = c.adProxy.InsertEntity(ctx, "entity", infoSystemData, graphIdInt, "info_system")
		//	if err != nil {
		//		return err
		//	}
		//
		//}

		if oInfoSystemId != "" && oInfoSystemId != dataCatalogInfo.InfoSystemId {
			deleteSideInfo1 := []map[string]map[string]string{
				{
					"start": {"info_system_uuid": oInfoSystemId, "_start_entity": "info_system"},
					"end":   {"datacatalogid": entityId, "_end_entity": "datacatalog"},
				},
			}
			_, err = c.adProxy.DeleteEdge(ctx, deleteSideInfo1, "info_system_2_datacatalog", graphIdInt)
		}

		sideInfo1 := []map[string]map[string]string{
			{
				"start":     {"info_system_uuid": dataCatalogInfo.InfoSystemId, "_start_entity": "info_system"},
				"end":       {"datacatalogid": entityId, "_end_entity": "datacatalog"},
				"edge_pros": {},
			},
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo1, graphIdInt, "info_system_2_datacatalog")
		if err != nil {
			return state, err
		}
	}

	// 部门
	if dataCatalogInfo.DepartmentId != "" {
		//searchCfg, err := c.adProxy.GetSearchConfig("department", "departmentid", dataCatalogInfo.DepartmentId)
		//entityInfo, err := c.adProxy.FulltextSearchV2(ctx, graphId, "", searchCfg)
		//if len(entityInfo.Res.Nodes) == 0 {
		//	department, err := c.dbRepo.GetDepartment(ctx, entityId)
		//	if err != nil {
		//		return err
		//	}
		//	if department == nil {
		//		errors.Wrap(err, "can not find data_view "+entityId)
		//		return err
		//	}
		//	departmentData := []map[string]any{
		//		{
		//			"departmentid":   department.Id,
		//			"departmentname": department.Name,
		//		},
		//	}
		//	_, err = c.adProxy.InsertEntity(ctx, "entity", departmentData, graphIdInt, "department")
		//	if err != nil {
		//		return err
		//	}
		//
		//}

		if oDepartmentId != "" && oDepartmentId != dataCatalogInfo.DepartmentId {
			deleteSideInfo2 := []map[string]map[string]string{
				{
					"start": {"departmentid": oDepartmentId, "_start_entity": "department"},
					"end":   {"datacatalogid": entityId, "_end_entity": "datacatalog"},
				},
			}
			_, err = c.adProxy.DeleteEdge(ctx, deleteSideInfo2, "department_2_datacatalog", graphIdInt)
		}
		sideInfo2 := []map[string]map[string]string{
			{
				"start":     {"departmentid": dataCatalogInfo.DepartmentId, "_start_entity": "department"},
				"end":       {"datacatalogid": entityId, "_end_entity": "datacatalog"},
				"edge_pros": {},
			},
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo2, graphIdInt, "department_2_datacatalog")
		if err != nil {
			return state, err
		}
	}

	// data owner
	if dataCatalogInfo.OwnerId != "" {
		//searchCfg, err := c.adProxy.GetSearchConfig("dataowner", "dataownerid", dataCatalogInfo.OwnerId)
		//entityInfo, err := c.adProxy.FulltextSearchV2(ctx, graphId, "", searchCfg)
		//if len(entityInfo.Res.Nodes) == 0 {
		//	dataOwner, err := c.dbRepo.GetDataOwner(ctx, entityId)
		//	if err != nil {
		//		return err
		//	}
		//	if dataOwner == nil {
		//		errors.Wrap(err, "can not find data_view "+entityId)
		//		return err
		//	}
		//	dataOwnerData := []map[string]any{
		//		{
		//			"dataownerid":   dataOwner.OwnerId,
		//			"dataownername": dataOwner.OwnerName,
		//		},
		//	}
		//	_, err = c.adProxy.InsertEntity(ctx, "entity", dataOwnerData, graphIdInt, "dataowner")
		//	if err != nil {
		//		return err
		//	}
		//
		//}
		if oOwnerId != "" && oOwnerId != dataCatalogInfo.OwnerId {
			deleteSideInfo3 := []map[string]map[string]string{
				{
					"start": {"dataownerid": oOwnerId, "_start_entity": "dataowner"},
					"end":   {"datacatalogid": entityId, "_end_entity": "datacatalog"},
				},
			}
			_, err = c.adProxy.DeleteEdge(ctx, deleteSideInfo3, "dataowner_2_datacatalog", graphIdInt)
		}
		sideInfo3 := []map[string]map[string]string{
			{
				"start":     {"dataownerid": dataCatalogInfo.OwnerId, "_start_entity": "dataowner"},
				"end":       {"datacatalogid": entityId, "_end_entity": "datacatalog"},
				"edge_pros": {},
			},
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo3, graphIdInt, "dataowner_2_datacatalog")
		if err != nil {
			return state, err
		}
	}

	// form_view
	formView2DataCatalog, err := c.dbRepo.GetFormView2DataCatalogByDataCatalog(ctx, dataCatalogInfo.Code)
	if err != nil {
		return state, err
	}
	if len(formView2DataCatalog) > 0 {
		sideInfo4 := []map[string]map[string]string{}
		for _, item := range formView2DataCatalog {
			addItem := map[string]map[string]string{
				"start":     {"formview_uuid": item.FormViewUuid, "_start_entity": "form_view"},
				"end":       {"datacatalogid": entityId, "_end_entity": "datacatalog"},
				"edge_pros": {},
			}
			sideInfo4 = append(sideInfo4, addItem)

			//searchCfg, err := c.adProxy.GetSearchConfig("form_view", "formview_uuid", dataCatalogInfo.OwnerId)
			//if err != nil {
			//	continue
			//}
			//entityInfo, err := c.adProxy.FulltextSearchV2(ctx, graphId, "", searchCfg)
			//if len(entityInfo.Res.Nodes) == 0 {
			//	formViewV2, err := c.dbRepo.GetFormViewV2(ctx, entityId)
			//	if err != nil {
			//		return err
			//	}
			//	if formViewV2 == nil {
			//		errors.Wrap(err, "can not find data_view "+entityId)
			//		return err
			//	}
			//	formViewV2Data := []map[string]any{
			//		{
			//			"description":    formViewV2.Description,
			//			"department_id":  formViewV2.DepartmentId,
			//			"subject_id":     formViewV2.SubjectId,
			//			"owner_id":       formViewV2.OwnerId,
			//			"publish_at":     formViewV2.PublishAt,
			//			"datasource_id":  formViewV2.DatasourceId,
			//			"type":           formViewV2.Type,
			//			"business_name":  formViewV2.BusinessName,
			//			"technical_name": formViewV2.TechnicalName,
			//			"formview_code":  formViewV2.FormviewCode,
			//			"formview_uuid":  formViewV2.FormviewUuid,
			//		},
			//	}
			//	_, err = c.adProxy.InsertEntity(ctx, "entity", formViewV2Data, graphIdInt, "form_view")
			//	if err != nil {
			//		return err
			//	}

			//}
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "formview_2_datacatalog")
		if err != nil {
			return state, err
		}
	}

	// catalogtag
	catalogTag2DataCatalog, err := c.dbRepo.GetCatalogTag2DataCatalogByDataCatalog(ctx, entityId)
	if err != nil {
		return state, err
	}
	if len(catalogTag2DataCatalog) > 0 {

		sideInfo5 := []map[string]map[string]string{}
		for _, item := range catalogTag2DataCatalog {
			addItem := map[string]map[string]string{
				"start":     {"catalogtagid": item.TagSid, "_start_entity": "catalogtag"},
				"end":       {"datacatalogid": entityId, "_end_entity": "datacatalog"},
				"edge_pros": {},
			}
			sideInfo5 = append(sideInfo5, addItem)

			searchCfg, err := c.adProxy.GetSearchConfig("catalogtag", "catalogtagid", item.TagSid)
			if err != nil {
				continue
			}
			entityInfo, err := c.adProxy.FulltextSearchV2(ctx, graphId, "", searchCfg)
			if len(entityInfo.Res.Nodes) == 0 {
				cataLogTag, err := c.dbRepo.GetCatalogTag(ctx, entityId)
				if err != nil {
					return state, err
				}
				if cataLogTag == nil {
					errors.Wrap(err, "can not find data_view "+entityId)
					return state, err
				}
				cataLogTagData := []map[string]any{
					{
						"catalogtagid":   cataLogTag.TagSid,
						"catalogtagname": cataLogTag.TagName,
						"tag_code":       cataLogTag.TagCode,
					},
				}
				_, err = c.adProxy.InsertEntity(ctx, "entity", cataLogTagData, graphIdInt, "catalogtag")
				if err != nil {
					return state, err
				}

			}
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo5, graphIdInt, "catalogtag_2_datacatalog")
		if err != nil {
			return state, err
		}
	}
	state = 1
	return state, nil
}

func (c *consumer) createCatalogTag(ctx context.Context, graphConfigName string, entityId string) error {
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

	entityInfo, err := c.dbRepo.GetCatalogTag(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"catalogtagid":   entityInfo.TagSid,
			"catalogtagname": entityInfo.TagName,
			"tag_code":       entityInfo.TagCode,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "catalogtag")
	if err != nil {
		return err
	}

	// 数据资源目录
	//catalogTag2DataCatalog, err := c.dbRepo.GetCatalogTag2DataCatalogByDataCatalog(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(catalogTag2DataCatalog) > 0 {
	//	sideInfo4 := []map[string]map[string]string{}
	//	for _, item := range catalogTag2DataCatalog {
	//		addItem := map[string]map[string]string{
	//			"start":     {"field_id": entityId, "_start_entity": "catalogtag"},
	//			"end":       {"datacatalogid": item.DataCatalogSid, "_end_entity": "datacatalog"},
	//			"edge_pros": {},
	//		}
	//		sideInfo4 = append(sideInfo4, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "catalogtag_2_datacatalog")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createInfoSystem(ctx context.Context, graphConfigName string, entityId string) error {
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

	entityInfo, err := c.dbRepo.GetInfoSystem(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"infosystemid":            entityInfo.SysSid,
			"infosystemname":          entityInfo.InfoSystemName,
			"info_system_uuid":        entityInfo.InfoSystemUuid,
			"info_system_description": entityInfo.InfoSystemDescription,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "info_system")
	if err != nil {
		return err
	}

	// 数据资源目录
	//infoSystem2DataCatalog, err := c.dbRepo.GetInfoSystem2DataCatalogByInfoSystem(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(infoSystem2DataCatalog) > 0 {
	//	sideInfo4 := []map[string]map[string]string{}
	//	for _, item := range infoSystem2DataCatalog {
	//		addItem := map[string]map[string]string{
	//			"start":     {"infosystemid": entityId, "_start_entity": "info_system"},
	//			"end":       {"datacatalogid": item.DataCatalogSid, "_end_entity": "datacatalog"},
	//			"edge_pros": {},
	//		}
	//		sideInfo4 = append(sideInfo4, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "info_system_2_datacatalog")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createDepartmentV2(ctx context.Context, graphConfigName string, entityId string) error {
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

	entityInfo, err := c.dbRepo.GetDepartmentV2(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"departmentid":   entityInfo.Id,
			"departmentname": entityInfo.Name,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "department")
	if err != nil {
		return err
	}

	// 数据资源目录
	//department2DataCatalog, err := c.dbRepo.GetDepartment2DataCatalogByDepartment(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(department2DataCatalog) > 0 {
	//	sideInfo4 := []map[string]map[string]string{}
	//	for _, item := range department2DataCatalog {
	//		addItem := map[string]map[string]string{
	//			"start":     {"departmentid": entityId, "_start_entity": "department"},
	//			"end":       {"datacatalogid": item.DataCatalogSid, "_end_entity": "datacatalog"},
	//			"edge_pros": {},
	//		}
	//		sideInfo4 = append(sideInfo4, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "department_2_datacatalog")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createDataOwnerV2(ctx context.Context, graphConfigName string, entityId string) error {
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

	entityInfo, err := c.dbRepo.GetDataOwnerV2(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"dataownerid":   entityInfo.OwnerId,
			"dataownername": entityInfo.OwnerName,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "dataowner")
	if err != nil {
		return err
	}

	// 数据资源目录
	//dataOwner2DataCatalog, err := c.dbRepo.GetDataOwner2DataCatalogByDataOwner(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(dataOwner2DataCatalog) > 0 {
	//	sideInfo4 := []map[string]map[string]string{}
	//	for _, item := range dataOwner2DataCatalog {
	//		addItem := map[string]map[string]string{
	//			"start":     {"dataownerid": entityId, "_start_entity": "dataowner"},
	//			"end":       {"datacatalogid": item.DataCatalogSid, "_end_entity": "datacatalog"},
	//			"edge_pros": {},
	//		}
	//		sideInfo4 = append(sideInfo4, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "dataowner")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) createFormViewV2(ctx context.Context, graphConfigName string, entityId string) error {
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

	entityInfo, err := c.dbRepo.GetFormViewV2(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"subject_path":       entityInfo.SubjectPath,
			"subject_path_id":    entityInfo.SubjectPathId,
			"department_path":    entityInfo.DepartmentPath,
			"department_path_id": entityInfo.DepartmentPathId,
			"subject_name":       entityInfo.SubjectName,
			"department":         entityInfo.Department,
			"owner_name":         entityInfo.OwnerName,
			"description":        entityInfo.Description,
			"department_id":      entityInfo.DepartmentId,
			"subject_id":         entityInfo.SubjectId,
			"owner_id":           entityInfo.OwnerId,
			"publish_at":         entityInfo.PublishAt,
			"datasource_id":      entityInfo.DatasourceId,
			"type":               entityInfo.Type,
			"business_name":      entityInfo.BusinessName,
			"technical_name":     entityInfo.TechnicalName,
			"formview_code":      entityInfo.FormviewCode,
			"formview_uuid":      entityInfo.FormviewUuid,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "form_view")
	if err != nil {
		return err
	}

	// 数据资源目录
	//formView2DataCatalog, err := c.dbRepo.GetFormView2DataCatalogByFormView(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(formView2DataCatalog) > 0 {
	//	sideInfo4 := []map[string]map[string]string{}
	//	for _, item := range formView2DataCatalog {
	//		addItem := map[string]map[string]string{
	//			"start":     {"formview_uuid": entityId, "_start_entity": "form_view"},
	//			"end":       {"datacatalogid": item.DataCatalogCode, "_end_entity": "datacatalog"},
	//			"edge_pros": {},
	//		}
	//		sideInfo4 = append(sideInfo4, addItem)
	//	}
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "formview_2_datacatalog")
	//	if err != nil {
	//		return err
	//	}
	//}

	// 库名称
	metadataSchema2MetadataTable, err := c.dbRepo.GetMetadataSchema2MetadataTableByFormView(ctx, entityId)
	if err != nil {
		return err
	}
	if len(metadataSchema2MetadataTable) > 0 {
		sideInfo4 := []map[string]map[string]string{}
		for _, item := range metadataSchema2MetadataTable {
			addItem := map[string]map[string]string{
				"start":     {"metadataschemaid": item.SchemaSid, "_start_entity": "metadataschema"},
				"end":       {"formview_uuid": entityId, "_end_entity": "form_view"},
				"edge_pros": {},
			}
			sideInfo4 = append(sideInfo4, addItem)

			searchCfg, err := c.adProxy.GetSearchConfig("metadataschema", "metadataschemaid", item.SchemaSid)
			if err != nil {
				continue
			}
			entityInfo, err := c.adProxy.FulltextSearchV2(ctx, graphId, "", searchCfg)
			if len(entityInfo.Res.Nodes) == 0 {
				DataSourceAndMetadataSchema, err := c.dbRepo.GetDataSourceAndMetadataSchemaByMetadataSchema(ctx, item.SchemaSid)
				if err != nil {
					return err
				}
				if DataSourceAndMetadataSchema == nil {
					errors.Wrap(err, "can not find metadata "+item.SchemaSid)
					return err
				}
				metadataSchemaData := []map[string]any{
					{
						"metadataschemaid":   DataSourceAndMetadataSchema.SchemaSid,
						"metadataschemaname": DataSourceAndMetadataSchema.SchemaName,
						"prefixname":         "库",
					},
				}
				_, err = c.adProxy.InsertEntity(ctx, "entity", metadataSchemaData, graphIdInt, "metadataschema")
				if err != nil {
					return err
				}

				//// 数据源，顺便创建
				//dataSourceData := []map[string]any{
				//	{
				//		"source_type_name":      DataSourceAndMetadataSchema.SourceTypeName,
				//		"source_type_code":      DataSourceAndMetadataSchema.SourceTypeCode,
				//		"data_source_type_name": DataSourceAndMetadataSchema.DataSourceTypeName,
				//		"prefixname":            DataSourceAndMetadataSchema.PrefixName,
				//		"datasourcename":        DataSourceAndMetadataSchema.DataSourceName,
				//		"datasourceid":          DataSourceAndMetadataSchema.DataSourceUuid,
				//	},
				//}
				//_, err = c.adProxy.InsertEntity(ctx, "entity", dataSourceData, graphIdInt, "datasource")
				//if err != nil {
				//	return err
				//}

				// 边数据
				dataSource2MetadataSchema := []map[string]map[string]string{
					{
						"start":     {"datasourceid": item.SchemaSid, "_start_entity": "datasource"},
						"end":       {"metadataschemaid": entityId, "_end_entity": "metadataschema"},
						"edge_pros": {},
					},
				}

				_, err = c.adProxy.InsertSide(ctx, "edge", dataSource2MetadataSchema, graphIdInt, "datasource_2_metadataschema")
				if err != nil {
					return err
				}

			}

		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "metadataschema_2_metadata_table")
		if err != nil {
			return err
		}
	}

	// 数据探查结果
	dataExploreReport2MetadataTable, err := c.dbRepo.GetDataExploreReport2MetadataTableByFormView(ctx, entityId)
	if err != nil {
		return err
	}
	if len(dataExploreReport2MetadataTable) > 0 {
		sideInfo4 := []map[string]map[string]string{}
		for _, item := range dataExploreReport2MetadataTable {
			addItem := map[string]map[string]string{
				"start": {
					"column_id":      item.ColumnId,
					"explore_item":   item.ExploreItem,
					"explore_result": item.ExploreResult,
					"_start_entity":  "data_explore_report"},
				"end":       {"formview_uuid": entityId, "_end_entity": "form_view"},
				"edge_pros": {},
			}
			sideInfo4 = append(sideInfo4, addItem)
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "data_explore_report_2_metadata_table")
		if err != nil {
			return err
		}
	}

	// 逻辑视图字段
	metadataTableField2MetadataTable, err := c.dbRepo.GetMetadataTableField2MetadataTableByFormView(ctx, entityId)
	if err != nil {
		return err
	}
	if len(metadataTableField2MetadataTable) > 0 {
		sideInfo4 := []map[string]map[string]string{}
		for _, item := range metadataTableField2MetadataTable {
			addItem := map[string]map[string]string{
				"start":     {"column_id": item.ColumnId, "_start_entity": "form_view_field"},
				"end":       {"formview_uuid": entityId, "_end_entity": "form_view"},
				"edge_pros": {},
			}
			sideInfo4 = append(sideInfo4, addItem)

			searchCfg, err := c.adProxy.GetSearchConfig("form_view_field", "column_id", item.ColumnId)
			if err != nil {
				continue
			}
			entityInfos, err := c.adProxy.FulltextSearchV2(ctx, graphId, "", searchCfg)
			if len(entityInfos.Res.Nodes) == 0 {
				dataViewFields, err := c.dbRepo.GetFormViewFieldV2(ctx, item.ColumnId)
				if err != nil {
					return err
				}
				if dataViewFields == nil {
					errors.Wrap(err, "can not find data_view filed "+item.ColumnId)
					return err
				}
				dataViewFieldsData := []map[string]any{
					{
						"formview_uuid":  dataViewFields.FormViewUuid,
						"technical_name": dataViewFields.TechnicalName,
						"business_name":  dataViewFields.BusinessName,
						"data_type":      dataViewFields.DataType,
						"column_id":      dataViewFields.ColumnId,
					},
				}
				_, err = c.adProxy.InsertEntity(ctx, "entity", dataViewFieldsData, graphIdInt, "form_view_field")
				if err != nil {
					return err
				}
			}
		}

		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "metadata_table_field_2_metadata_table")
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *consumer) deleteFormViewV2(ctx context.Context, graphConfigName string, entityId string, entityIdKey string, entityName string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	graphData := []map[string]string{
		{"formview_uuid": entityId},
	}

	_, err = c.adProxy.DeleteEntity(ctx, graphData, "form_view", graphIdInt)
	if err != nil {
		return err
	}

	dataView2Fileds, err := c.dbRepo.GetMetadataTableField2MetadataTableByFormView(ctx, entityId)
	if err != nil {
		return err
	}
	itemGraphData := []map[string]string{}
	for _, item := range dataView2Fileds {
		addItem := map[string]string{
			"column_id": item.ColumnId,
		}
		itemGraphData = append(itemGraphData, addItem)

	}

	if len(itemGraphData) > 0 {
		_, err = c.adProxy.DeleteEntity(ctx, itemGraphData, "form_view_field", graphIdInt)
	}

	return nil
}

func (c *consumer) createDataExploreReportV2(ctx context.Context, graphConfigName string, entityId string) error {
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

	entityInfo, err := c.dbRepo.GetDataExploreReportV2(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"column_id":      entityInfo.ColumnId,
			"explore_item":   entityInfo.ExploreItem,
			"column_name":    entityInfo.ColumnName,
			"explore_result": entityInfo.ExploreResult,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "data_explore_report")
	if err != nil {
		return err
	}

	// 数据资源目录
	dataOwner2DataCatalog, err := c.dbRepo.GetDataOwner2DataCatalogByDataOwner(ctx, entityId)
	if err != nil {
		return err
	}
	if len(dataOwner2DataCatalog) > 0 {
		sideInfo4 := []map[string]map[string]string{}
		for _, item := range dataOwner2DataCatalog {
			addItem := map[string]map[string]string{
				"start":     {"column_id": entityId, "_start_entity": "data_explore_report"},
				"end":       {"formview_uuid": item.DataCatalogSid, "_end_entity": "form_view"},
				"edge_pros": {},
			}
			sideInfo4 = append(sideInfo4, addItem)
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "data_explore_report_2_metadata_table")
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *consumer) deleteDataExploreReportV2(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	dataExploreReport, err := c.dbRepo.GetDataExploreReport(ctx, entityId)
	if err != nil {
		return err
	}
	if dataExploreReport == nil {
		errors.Wrap(err, "can not find data_view "+entityId)
		return err
	}
	graphData := []map[string]string{
		{
			"explore_result": dataExploreReport.ExploreResult,
			"explore_item":   dataExploreReport.ExploreItem,
			"column_id":      dataExploreReport.ColumnId,
		},
	}

	_, err = c.adProxy.DeleteEntity(ctx, graphData, "data_explore_report", graphIdInt)
	if err != nil {
		return err
	}

	return nil
}

func (c *consumer) createFormViewFieldV2(ctx context.Context, graphConfigName string, entityId string) error {
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

	entityInfo, err := c.dbRepo.GetFormViewFieldV2(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"formview_uuid":  entityInfo.FormViewUuid,
			"technical_name": entityInfo.TechnicalName,
			"business_name":  entityInfo.BusinessName,
			"data_type":      entityInfo.DataType,
			"column_id":      entityInfo.ColumnId,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "form_view_field")
	if err != nil {
		return err
	}

	// 逻辑视图
	if entityInfo.FormViewUuid != "" {
		sideInfo := []map[string]map[string]string{
			{
				"start":     {"column_id": entityId, "_start_entity": "form_view_field"},
				"end":       {"formview_uuid": entityInfo.FormViewUuid, "_end_entity": "form_view"},
				"edge_pros": {},
			},
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo, graphIdInt, "metadata_table_field_2_metadata_table")
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *consumer) createMetaDataSchemaV2(ctx context.Context, graphConfigName string, entityId string) error {
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

	entityInfo, err := c.dbRepo.GetMetadataSchemaV2(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"metadataschemaid":   entityInfo.SchemaSid,
			"metadataschemaname": entityInfo.SchemaName,
			"prefixname":         entityInfo.PrefixName,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "metadataschema")
	if err != nil {
		return err
	}

	// 逻辑视图
	metadataSchema2MetadataTable, err := c.dbRepo.GetMetadataSchema2MetadataTableByMetadataSchemaV2(ctx, entityId)
	if err != nil {
		return err
	}
	if len(metadataSchema2MetadataTable) > 0 {
		sideInfo4 := []map[string]map[string]string{}
		for _, item := range metadataSchema2MetadataTable {
			addItem := map[string]map[string]string{
				"start":     {"metadataschemaid": entityId, "_start_entity": "metadataschema"},
				"end":       {"formview_uuid": item.FormViewUuid, "_end_entity": "form_view"},
				"edge_pros": {},
			}
			sideInfo4 = append(sideInfo4, addItem)
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "metadataschema_2_metadata_table")
		if err != nil {
			return err
		}
	}

	// 数据源
	datasource2MetaDataSchema, err := c.dbRepo.GetDatasource2MetaDataSchemaByMetadataSchemaV2(ctx, entityId)
	if err != nil {
		return err
	}
	if len(datasource2MetaDataSchema) > 0 {
		sideInfo4 := []map[string]map[string]string{}
		for _, item := range datasource2MetaDataSchema {
			addItem := map[string]map[string]string{
				"start":     {"datasourceid": item.DataSourceUuid, "_start_entity": "datasource"},
				"end":       {"metadataschemaid": entityId, "_end_entity": "metadataschema"},
				"edge_pros": {},
			}
			sideInfo4 = append(sideInfo4, addItem)
		}
		_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "datasource_2_metadataschema")
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *consumer) createDataSourceV2(ctx context.Context, graphConfigName string, entityId string) error {
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

	entityInfo, err := c.dbRepo.GetDataSourceV2(ctx, entityId)
	if err != nil {
		return err
	}

	if entityInfo == nil {
		errors.Wrap(err, "can not find entity info "+entityId)
		return err
	}
	graphData := []map[string]any{
		{
			"source_type_name":      entityInfo.SourceTypeName,
			"source_type_code":      entityInfo.SourceTypeCode,
			"data_source_type_name": entityInfo.DataSourceTypeName,
			"prefixname":            entityInfo.PrefixName,
			"datasourcename":        entityInfo.DataSourceName,
			"datasourceid":          entityInfo.DataSourceUuid,
		},
	}

	//fmt.Println(graphIdInt, graphData)
	_, err = c.adProxy.InsertEntity(ctx, "entity", graphData, graphIdInt, "datasource")
	if err != nil {
		return err
	}

	// 库名称
	//datasource2MetaDataSchema, err := c.dbRepo.GetDatasource2MetaDataSchemaByDataSourceV2(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(datasource2MetaDataSchema) > 0 {
	//	sideInfo4 := []map[string]map[string]string{}
	//	for _, item := range datasource2MetaDataSchema {
	//		addItem := map[string]map[string]string{
	//			"start":     {"datasourceid": entityId, "_start_entity": "datasource"},
	//			"end":       {"metadataschemaid": item.SchemaSid, "_end_entity": "metadataschema"},
	//			"edge_pros": {},
	//		}
	//		sideInfo4 = append(sideInfo4, addItem)
	//
	//		dataSchema, err := c.dbRepo.GetMetadataSchema(ctx, entityId)
	//		if err != nil {
	//			continue
	//		}
	//		addInfo := []map[string]any{
	//			{
	//				"metadataschemaid":   dataSchema.SchemaSid,
	//				"metadataschemaname": dataSchema.SchemaName,
	//				"prefixname":         "库",
	//			},
	//		}
	//		_, err = c.adProxy.InsertEntity(ctx, "entity", addInfo, graphIdInt, "metadataschema")
	//		if err != nil {
	//			continue
	//		}
	//	}
	//
	//	_, err = c.adProxy.InsertSide(ctx, "edge", sideInfo4, graphIdInt, "datasource_2_metadataschema")
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (c *consumer) deleteDataSourceV2(ctx context.Context, graphConfigName string, entityId string) error {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	graphId, err := c.adCfgHelper.GetGraphId(ctx, graphConfigName)
	if err != nil {
		return err
	}

	graphIdInt, err := strconv.Atoi(graphId)
	if err != nil {
		return err
	}

	mainSearchCfg, err := c.adProxy.GetSearchConfig("datasource", "datasourceid", entityId)
	mainEntityInfo, err := c.adProxy.FulltextSearchV2(ctx, graphId, "", mainSearchCfg)

	if len(mainEntityInfo.Res.Nodes) == 0 {
		return nil
	}

	node := mainEntityInfo.Res.Nodes[0]

	edges, err := c.adProxy.NeighborSearchByEngineV2(ctx, node.Id, 1, graphId)
	if len(edges.Res.Nodes) > 0 {
		nNode := edges.Res.Nodes[0]
		metadataschemaid := ""
		if len(nNode.Properties) != 0 {
			for _, p := range nNode.Properties[0].Props {
				if p.Name == "metadataschemaid" {
					metadataschemaid = p.Value
				}
			}
			if metadataschemaid != "" {
				dItem := []map[string]string{
					{"metadataschemaid": metadataschemaid},
				}

				_, err = c.adProxy.DeleteEntity(ctx, dItem, "metadataschema", graphIdInt)
				if err != nil {
					return err
				}

			}
		}

	}

	graphData := []map[string]string{
		{"datasourceid": entityId},
	}
	entityName := "datasource"

	_, err = c.adProxy.DeleteEntity(ctx, graphData, entityName, graphIdInt)
	if err != nil {
		return err
	}
	// 删除库

	//dataSource2MetadataSchema, err := c.dbRepo.GetDataSource2MetadataSchema(ctx, entityId)
	//if err != nil {
	//	return err
	//}
	//if len(dataSource2MetadataSchema) > 0 {
	//	for _, item := range dataSource2MetadataSchema {
	//		dItem := []map[string]string{
	//			{"metadataschemaid": item.SchemaSid},
	//		}
	//
	//		_, err = c.adProxy.DeleteEntity(ctx, dItem, "metadataschema", graphIdInt)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//}

	return nil
}
