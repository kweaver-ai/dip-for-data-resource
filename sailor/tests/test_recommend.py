# -*- coding: utf-8 -*-

"""
@Time ：2024/1/29 18:49
@Auth ：Danny.gao
@File ：test_recommend.py
@Desc ：
@Motto：ABC(Always Be Coding)
"""
from app.logs.logger import logger
import unittest
import asyncio
from app.handlers.recommend_handler import RecommendTableParams, RecommendFlowParams, RecommendCodeParams, CheckCodeParams
from app.cores.recommend.recommend import recommendCode, recommendTable, recommendFlow, checkCode


class TestStringMethods(unittest.TestCase):

    def test_recommend_table(self):
        data = {
            "af_query": {
                "business_model_id": "业务模型id",
                "business_model_name": "业务模型名称",
                "domain": {
                    "domain_id": "业务域id",
                    "domain_name": "业务域名称",
                    "domain_path": "业务域层级",
                    "domain_path_id": "业务域层级id"
                },
                "dept": {
                    "dept_id": "组织部门id",
                    "dept_name": "组织部门名称",
                    "dept_path": "组织部门层级",
                    "dept_path_id": "组织部门层级id"
                },
                "info_system": [{
                    "info_system_id": "信息系统id",
                    "info_system_name": "信息系统名称",
                    "info_system_desc": "信息系统描述"
                }],
                "table": {
                    "name": "表单名称",
                    "description": "表单描述"
                 },
                "key": "不动产抵押表"
            },
            "graph_id": "601",
            "appid": ""
        }
        params = RecommendTableParams(**data)
        result = asyncio.run(recommendTable(params.af_query, params.graph_id, params.appid))
        logger.info(result)
        assert 'answers' in result
        assert 'tables' in result['answers']

    def test_recommend_flow(self):
        data = {
            "af_query": {
                "business_model_id": "业务模型id,当前理解为主干业务id,后期做调整",
                "node": {
                        "id": "数据库中流程节点ID,即uuid",
                        "mxcell_id": "前端节点ID",
                        "name": "节点名称 ",
                        "description": "节点描述"
                },
                "parent_node": {
                        "id": "数据库中流程节点ID,即uuid",
                        "mxcell_id": "前端节点ID ",
                        "name": "节点名称 ",
                        "description": "节点描述 ",
                        "flowchart_id": "所属流程图的ID ",
                        "tables": ["表单ID1", "表单ID2"]
                },
                "flowchart": {
                        "id": "当前需要推荐流程的节点所在流程图ID",
                        "name": "流程图名称",
                        "description": "流程图描述 ",
                        "business_model_id": "流程图所在业务模型,当前理解为主干业务id,后期做调整",
                        "nodes": [
                                {
                                        "id": "数据库中流程节点ID,即uuid",
                                        "mxcell_id": "前端节点ID",
                                        "name": "节点名称",
                                        "description": "节点描述 "
                                }, {
                                        "id": "数据库中流程节点ID, 即uuid",
                                        "mxcell_id" :"前端节点ID ",
                                        "name": "节点名称",
                                        "description": "节点描述"
                                }
                        ]
                }
            },
            "graph_id": "601",
            "appid": ""
        }
        params = RecommendFlowParams(**data)
        result = asyncio.run(recommendFlow(params.af_query, params.graph_id, params.appid))
        logger.info(result)
        assert 'answers' in result
        assert 'flowcharts' in result['answers']

    def test_recommend_code(self):
        data = {
            "query": {
                "table_id": "table-id-01",
                "table_name": "全员核酸检测数据",
                "table_desc": "",
                "table_fields": [
                    {
                        "table_field_id": "field-id-01",
                        "table_field_name": "人员所属区域标识"
                    },
                    {
                        "table_field_id": "field-id-02",
                        "table_field_name": "教育"
                    }
                ]
            },
            "graph_id": "601",
            "appid": ""
        }
        params = RecommendCodeParams(**data)
        result = asyncio.run(recommendCode(params.query, params.graph_id, params.appid))
        logger.info(result)
        assert 'answers' in result

    def test_check_code(self):
        data = {
            "check_af_query": [
                {
                    "business_model_id": "业务模型id",
                    "business_model_name": "业务模型名称",
                    "domain": {
                        "domain_id": "业务域id",
                        "domain_name": "业务域名称",
                        "domain_path": "业务域层级",
                        "domain_path_id": "业务域层级id"
                    },
                    "dept": {
                        "dept_id": "组织部门id",
                        "dept_name": "组织部门名称",
                        "dept_path": "组织部门层级",
                        "dept_path_id": "组织部门层级id"
                    },
                    "info_system": [{
                        "info_system_id": "信息系统id",
                        "info_system_name": "信息系统名称",
                        "info_system_desc": "信息系统描述"
                    }],
                    "tables":[
                        {
                            "table_id": "table-01",
                        "table_name": "全员核酸检测数据",
                        "table_desc": "",
                        "table_fields": [
                            {
                                "field_id": "field-01",
                                "table_field_name": "教育",
                                "table_field_desc": "",
                                "standard_id": "4"
                            }
                        ]
                        }
                    ]

                }
            ],
            "graph_id": "601",
            "appid": ""
        }
        params = CheckCodeParams(**data)
        result = asyncio.run(checkCode(params.check_af_query, params.graph_id, params.appid))
        logger.info(result)
        assert 'answers' in result


if __name__ == '__main__':
    unittest.main()