# -*- coding: utf-8 -*-

"""
@Time ：2024/1/29 18:49
@Auth ：Danny.gao
@File ：test_recommend.py
@Desc ：
@Motto：ABC(Always Be Coding)
"""

import json
import unittest
import asyncio

from fastapi import BackgroundTasks

from app.cores.understand.understand import tableCompletion


class TestStringMethods(unittest.TestCase):

    def test_table_completion(self):
        inputs = {
            "id": "Table-ID-01",
            "technical_name": "T_ActiceCodeApplyImplement",
            "business_name": "",
            "desc": "",
            "database": "",
            "subject": "",
            "columns": [
                {"id": "Field-Id-01", "technical_name": "id", "business_name": "id", "data_type": "int", "comment": ""},
                {"id": "Field-Id-02", "technical_name": "activeCodeApplyId", "business_name": "id", "data_type": "int", "comment": ""},
            ],
            "view_source_catalog_name": "",
            "request_type": 1,
            "demo_data": {}
        }

        af_auth = ""
        appid = ''

        background_tasks = BackgroundTasks()

        res = asyncio.run(tableCompletion(background_tasks=background_tasks,
                                          query=inputs,
                                          appid='',
                                          af_auth="", only_for_table=False))
        assert 'task_id' in res


if __name__ == '__main__':
    unittest.main()