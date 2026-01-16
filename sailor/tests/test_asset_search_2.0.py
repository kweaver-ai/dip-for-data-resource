# -*- coding: utf-8 -*-
# @Time    : 2024/1/22 16:03
# @Author  : Glen.lv
# @File    : test_asset_search
# @Project : copilot

import asyncio
import json
import os
import time
import unittest
from unittest import mock

import ast
import math
import re
import jieba
import ahocorasick
import asyncio
import time
import numpy
import json
# import Levenshtein
from itertools import groupby
from typing import List, Dict
from sklearn.feature_extraction.text import CountVectorizer
from sklearn.metrics.pairwise import cosine_similarity
from starlette import status

# from anydata import (Builder,CogEngine)
# from anydata.infra.opensearch import Opensearch
# from anydata import Opensearch
# from anydata import *
from app.logs.logger import logger
from app.cores.cognitive_search.sdk_utils import (ad_builder_download_lexicon, ad_builder_get_kg_info,ad_opensearch_connector)
from app.cores.cognitive_search.sdk_utils import (ad_cog_engine_cognitive_service_config_call,ad_cog_engine_custom_search_service_call)
# from app.cores.cognitive_search.sdk_utils import ad_cog_engine_connector
from app.cores.cognitive_search import re_asset_search


class TestAssetSearch(unittest.TestCase):

    def test_run(self):
        inputs = {
            "query": "你是谁",
            "limit": 4,
            "stopwords": [],
            "stop_entities": [],
            "filter": {
                "data_kind": "0",
                # "update_cycle": "[0,1,2,5]",
                "update_cycle": "[-1]",
                "shared_type": "[-1]",
                "start_time": "1600122122",
                "end_time": "1800122122",
                "asset_type": "[-1]",
                "stop_entity_infos": [
                    {
                        "class_name": "",
                        "names": [
                            ""
                        ]
                    }
                ]
            },
            "ad_appid": "",
            "kg_id": 19,
            "entity2service": {
                "businessobject": {
                    "relation": "关联",
                    "service": "0b61a9252ed64439a9a06a94a1f6cd1d",
                    "weight": 1.6
                },
                "catalogtag": {
                    "relation": "打标",
                    "service": "0b61a9252ed64439a9a06a94a1f6cd1d",
                    "weight": 1.5
                },
                "data_explore_report": {
                    "relation": "包含",
                    "service": "f70bd69da0d04c4cbc8888563406b470",
                    "weight": 1.7
                },
                "datacatalog": {
                    "relation": "",
                    "service": "c1b8cae604dd4672b0a6ddb5f7ea2586",
                    "weight": 4
                },
                "dataowner": {
                    "relation": "管理",
                    "service": "0b61a9252ed64439a9a06a94a1f6cd1d",
                    "weight": 1.4
                },
                "datasource": {
                    "relation": "包含",
                    "service": "f70bd69da0d04c4cbc8888563406b470",
                    "weight": 1
                },
                "department": {
                    "relation": "管理",
                    "service": "0b61a9252ed64439a9a06a94a1f6cd1d",
                    "weight": 1.3
                },
                "domain": {
                    "relation": "包含",
                    "service": "f70bd69da0d04c4cbc8888563406b470",
                    "weight": 1
                },
                "info_system": {
                    "relation": "关联",
                    "service": "0b61a9252ed64439a9a06a94a1f6cd1d",
                    "weight": 1.2
                },
                "metadata_table": {
                    "relation": "编目",
                    "service": "0b61a9252ed64439a9a06a94a1f6cd1d",
                    "weight": 1.9
                },
                "metadata_table_field": {
                    "relation": "包含",
                    "service": "5b2e39e26f054aaaa49f6b6db1e8bd1f",
                    "weight": 1.8
                },
                "metadataschema": {
                    "relation": "包含",
                    "service": "5b2e39e26f054aaaa49f6b6db1e8bd1f",
                    "weight": 1
                },
                "subdomain": {
                    "relation": "包含",
                    "service": "5b2e39e26f054aaaa49f6b6db1e8bd1f",
                    "weight": 1
                }
            },
            "required_resource": {
                "lexicon_actrie": {
                    "lexicon_id": "3"
                },
                "stopwords": {
                    "lexicon_id": "4"
                }
            }
        }

        from app.handlers.cognitive_search_handler import AssetSearchParams

        inputs = AssetSearchParams(**inputs)
        outputs = asyncio.run(re_asset_search.run_func(inputs))
        # print(type(outputs))
        # print(outputs)
        assert 'entities' in outputs
        assert 'count' in outputs
        assert 'answer' in outputs
        assert 'subgraphs' in outputs
        assert 'query_cuts' in outputs

        for x in outputs['entities']:
            print(x)
        print(len(outputs['entities']))



if __name__ == '__main__':
    unittest.main()
