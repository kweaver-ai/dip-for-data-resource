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
from app.handlers.categorize_handler import DataCategorizeParams
from app.cores.categorize.categorize import dataCategorize


class TestStringMethods(unittest.TestCase):

    def test_categorize(self):
        data = {
            "query": {
                "view_id": "逻辑视图ID",
                "view_technical_name": "逻辑视图技术名称",
                "view_business_name": "逻辑视图业务名称",
                "view_desc": "逻辑视图描述",
                "subject_id": "585b15d7-99a8-4ab0-8684-a006ee4791fd",
                "view_fields": [
                    {
                    "view_field_id": "字段ID-01",
                    "view_field_technical_name": "字段技术名称",
                    "view_field_business_name": "性别",
                    "standard_code": "字段关联标准"
                    },
                    {
                        "view_field_id": "字段ID-02",
                        "view_field_technical_name": "字段技术名称",
                        "view_field_business_name": "姓名",
                        "standard_code": "1781188721228054530"
                    }
                ]
            },
            "graph_id": "40",
            "appid": ""
        }
        params = DataCategorizeParams(**data)
        result = asyncio.run(dataCategorize(params.query, params.graph_id, params.appid))
        logger.info(result)
        assert 'answers' in result



if __name__ == '__main__':
    unittest.main()