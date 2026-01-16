# # -*- coding: utf-8 -*-
# # @Time    : 2024/1/22 16:03
# # @Author  : Glen.lv
# # @File    : test_asset_search
# # @Project : copilot
#
# import asyncio
# import json
# import os
# import time
# import unittest
# from unittest import mock
#
# import ast
# import math
# import re
# import jieba
# import ahocorasick
# import asyncio
# import time
# import numpy
# import json
# # import Levenshtein
# from itertools import groupby
# from typing import List, Dict
# from sklearn.feature_extraction.text import CountVectorizer
# from sklearn.metrics.pairwise import cosine_similarity
# from starlette import status
#
# # from anydata import (Builder,CogEngine)
# # from anydata.infra.opensearch import Opensearch
# # from anydata import Opensearch
# # from anydata import *
# from app.logs.logger import logger
# from app.cores.cognitive_search.sdk_utils import (ad_builder_download_lexicon, ad_builder_get_kg_info,ad_opensearch_connector)
# from app.cores.cognitive_search.sdk_utils import (ad_cog_engine_cognitive_service_config_call,ad_cog_engine_custom_search_service_call)
# # from app.cores.cognitive_search.sdk_utils import ad_cog_engine_connector
# from app.cores.cognitive_search import asset_search
#
#
# class TestAssetSearch(unittest.TestCase):
#
#     def test_run(self):
#
#         from app.handlers.cognitive_search_handler import AssetSearchParams
#
#         inputs = AssetSearchParams(**inputs)
#         outputs = asyncio.run(asset_search.run_func(inputs))
#         # print(type(outputs))
#         # print(outputs)
#         assert 'entities' in outputs
#         assert 'count' in outputs
#         assert 'answer' in outputs
#         assert 'subgraphs' in outputs
#         assert 'query_cuts' in outputs
#
#         for x in outputs['entities']:
#             print(x)
#         print(len(outputs['entities']))
#
#
#
# if __name__ == '__main__':
#     unittest.main()
