package concepts

const (
	//SearchDomainRange 查询空间范围
	SearchDomainRange = `SELECT ${column_name} AS space_col FROM ${ve_catalog_id}.${schema_name}.${name}`
	//SearchTimeRange 查询时间范围
	SearchTimeRange = `SELECT  MIN(to_unixtime(${column_name})) AS min_time, MAX(to_unixtime(${column_name})) AS max_time FROM ${ve_catalog_id}.${schema_name}.${name}`
)

//CatalogInfo  数据资源目录信息
const CatalogInfo = `lookup on entity_data_catalog where entity_data_catalog.id==${catalog_id} YIELD vertex as data_catalog`

//Catalog2DepartmentInfo 数据目录查询所属部门
const Catalog2DepartmentInfo = `lookup on entity_data_catalog where entity_data_catalog.id==${catalog_id}
    YIELD id(vertex) as id  | GO from $-.id
      OVER  entity_department_2_entity_data_catalog BIDIRECT
      where "entity_department" in tags($$)
    YIELD  $$ as  entity_department`

//Catalog2SourceBusinessScene 数据目录查询来源业务场景
const Catalog2SourceBusinessScene = `lookup on entity_data_catalog where entity_data_catalog.id==${catalog_id}
    YIELD id(vertex) as id  | GO from $-.id
    OVER  entity_catalog_2_entity_business_scene_source BIDIRECT
      where "entity_business_scene" in tags($$)
    YIELD  $$ as  entity_business_scene_source`

//Catalog2RelatedBusinessScene 数据目录查询关联业务场景
const Catalog2RelatedBusinessScene = `lookup on entity_data_catalog where entity_data_catalog.id==${catalog_id}
    YIELD id(vertex) as id  | GO from $-.id
    OVER  entity_catalog_2_entity_business_scene_related BIDIRECT
      where "entity_business_scene" in tags($$)
    YIELD  $$ as  entity_business_scene_related`

//Catalog2BusinessObjects 数据目录查询业务对象
const Catalog2BusinessObjects = `lookup on entity_data_catalog where entity_data_catalog.id==${catalog_id}
    YIELD id(vertex) as id  | GO from $-.id
    OVER  entity_business_object_2_entity_data_catalog BIDIRECT
      where "entity_business_object" in tags($$)
    YIELD  $$ as  entity_business_object`

//Catalog2StandardTable  数据资源目录查询标准表
const Catalog2StandardTable = `lookup on entity_data_catalog where entity_data_catalog.id==${catalog_id}
    YIELD id(vertex) as id  | GO from $-.id
    OVER  entity_standard_table_2_entity_data_catalog BIDIRECT
      where "entity_standard_table" in tags($$)
    YIELD  $$ as  entity_standard_table`

//CatalogFields   查询数据资源目录信息项
const CatalogFields = `lookup on entity_data_catalog where entity_data_catalog.id==${catalog_id}
    YIELD id(vertex) as id  | GO from $-.id
    OVER  entity_data_catalog_2_entity_data_catalog_column BIDIRECT
      where "entity_data_catalog_column" in tags($$)
    YIELD  $$ as  entity_data_catalog_column `

//StandardTable2Indicators 标准表名称和字段名查询关联的业务指标
const StandardTable2Indicators = `lookup on entity_standard_table where entity_standard_table.id==${id}
    YIELD id(vertex) as id  | GO 2 steps from $-.id
    OVER  entity_business_form_2_entity_standard_table,
          entity_business_form_standard_2_entity_field BIDIRECT
      where  ("entity_field" in tags($$) and properties($$).name_en==${column_name})
    YIELD id($$) as id | go from $-.id
    OVER  entity_field_2_entity_business_indicator
      where ( "entity_business_indicator" in tags($$))
    YIELD   $$ as  indicator `

//DataCatalog2RelatedBusinessForm  数据资源目录查询引用业务表
const DataCatalog2RelatedBusinessForm = `lookup on entity_data_catalog where entity_data_catalog.id==${catalog_id}
    YIELD id(vertex) as id  | GO 3 steps from $-.id
    OVER  entity_standard_table_2_entity_data_catalog,
          entity_business_form_2_entity_standard_table,
          entity_business_form_standard_2_self BIDIRECT
      where "entity_business_form_standard" in tags($$)
    YIELD  $$ as  entity_business_form_standard `

//BusinessForm2DataCatalog 引用业务表查询数据资源编目
const BusinessForm2DataCatalog = `lookup on entity_business_form_standard where entity_business_form_standard.id in ${form_ids}
    YIELD id(vertex) as id  | GO 2 steps from $-.id 
    OVER  entity_business_form_2_entity_standard_table,
          entity_standard_table_2_entity_data_catalog BIDIRECT
      where "entity_data_catalog" in tags($$)  
    YIELD  $$ as  entity_data_catalog`

//DataCatalog2Indicator 单个数资源目录查询业务指标
const DataCatalog2Indicator = `lookup on entity_data_catalog where entity_data_catalog.id == ${catalog_id}
    YIELD id(vertex) as id  | GO 4 steps from $-.id
    OVER  entity_standard_table_2_entity_data_catalog,
          entity_business_form_2_entity_standard_table,
          entity_business_form_standard_2_entity_field,
          entity_field_2_entity_business_indicator BIDIRECT
      where "entity_business_indicator" in tags($$)
    YIELD  $$ as  entity_business_indicator `

//DataCatalog2SomeIndicator 关联数据资源目录的指定指标
const DataCatalog2SomeIndicator = `lookup on entity_data_catalog where entity_data_catalog.id in ${catalog_ids}
    YIELD id(vertex) as id  | GO 4 steps from $-.id
    OVER  entity_standard_table_2_entity_data_catalog,
          entity_business_form_2_entity_standard_table,
          entity_business_form_standard_2_entity_field,
          entity_field_2_entity_business_indicator BIDIRECT
       where ("entity_business_indicator" in tags($$)  and properties($$).indicator_id in ${indicator_ids} )
    YIELD  $$ as  entity_business_indicator `

//CommonIndicator2DataCatalog 查询到共同的指标，反查数据编目
const CommonIndicator2DataCatalog = ` lookup on entity_business_indicator  where  entity_business_indicator.indicator_id==${indicator_id}
     YIELD id(vertex) as id  | GO 4 steps from $-.id
     OVER  entity_field_2_entity_business_indicator,
          entity_business_form_standard_2_entity_field,
          entity_business_form_2_entity_standard_table,
          entity_standard_table_2_entity_data_catalog BIDIRECT
      where ("entity_data_catalog" in tags($$) and properties($$).id != ${catalog_id} and properties($$).id  in ${catalog_ids} )
    YIELD  $$ as  entity_data_catalog`

//StandardFields2BusinessFormFields 多个标准表字段查询业务表字段
const StandardFields2BusinessFormFields = `lookup on entity_standard_table where entity_standard_table.id==${table_id}
  YIELD id(vertex) as id  | GO 2 steps from $-.id
    OVER  entity_business_form_2_entity_standard_table,
          entity_business_form_standard_2_entity_field BIDIRECT
    where  ("entity_field" in tags($$)  and properties($$).name_en in ${column_name_slice})
    YIELD  $$ as entity_field`

//StandardField2BusinessFormField 单个标准表字段查询业务表字段
const StandardField2BusinessFormField = `lookup on entity_standard_table where entity_standard_table.id==${table_id}
  YIELD id(vertex) as id  | GO 2 steps from $-.id
    OVER  entity_business_form_2_entity_standard_table,
          entity_business_form_standard_2_entity_field BIDIRECT
    where  ("entity_field" in tags($$)  and properties($$).name_en == ${column_name})
    YIELD  $$ as entity_field`

//BusinessFormField2BusinessFormWithinBusinessObject  业务表字段查询业务表_有业务对象关系
const BusinessFormField2BusinessFormWithinBusinessObject = `lookup on entity_field where  entity_field.field_id==${field_id} 
  			 YIELD id(vertex) as id  | GO 2 steps from $-.id 
         over entity_business_form_standard_2_entity_field, 
         entity_business_form_2_entity_business_object BIDIRECT
      where  ("entity_business_object" in tags($$)  and  properties($$).name in ${business_object_name_slice})
       YIELD $^ as business_form_standard`

//BusinessForm2BusinessModel 业务表查询主干业务
const BusinessForm2BusinessModel = `lookup on entity_business_form_standard where entity_business_form_standard.id==${id}
      YIELD id(vertex) as id  | GO from $-.id  over
       entity_business_form_2_entity_business_model BIDIRECT
        where "entity_business_model" in tags($$)
         YIELD $$ as business_model`

//BusinessForm2BusinessFlowchart 业务表查询业务流程
const BusinessForm2BusinessFlowchart = `lookup on entity_business_form_standard where entity_business_form_standard.id==${id}
      YIELD id(vertex) as id  | GO 2 steps from $-.id  over
        entity_business_form_2_entity_business_model,
        entity_business_model_2_entity_flowchart BIDIRECT
         where "entity_flowchart" in tags($$)
           YIELD $$ as business_model`

//BusinessFormField2BusinessFormWithoutBusinessObject  业务表字段查询业务表_无业务对象关系
const BusinessFormField2BusinessFormWithoutBusinessObject = `lookup on entity_field where  entity_field.ref_id==${ref_id} 
  			 YIELD id(vertex) as id  | GO 2 steps from $-.id 
         over entity_business_form_standard_2_entity_field, 
         entity_business_form_2_entity_business_object BIDIRECT
      where  ("entity_business_object" in tags($$) and  properties($$).name not in ${business_object_name_slice} )
       YIELD $^ as business_form_standard`

//BusinessRefId2Indicators 业务表引用父字段查询子字段的业务指标
const BusinessRefId2Indicators = `lookup  on  entity_field  where entity_field.ref_id==${field_id} and
		entity_field.business_form_id==${business_form_id}
       YIELD id(vertex) as id  |  GO  from $-.id
       OVER entity_field_2_entity_business_indicator BIDIRECT
       where  ("entity_business_indicator" in tags($$))    
       YIELD $$ as indicator`
