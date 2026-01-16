# -*- coding: utf-8 -*-
"""
Search Tools 模块

本模块包含数据搜索相关的工具，包括：
- AfSailorTool: 数据搜索工具
- MultiQuerySearchTool: 多查询搜索工具
- DataSourceFilterTool: 数据源过滤工具
- DataScopeCheckerTool: 数据范围检查工具
- DataSeekerIntentionRecognizerTool: 数据搜索意图识别工具
- DataSeekerReportWriterTool: 数据搜索报告撰写工具
"""

from .af_sailor import AfSailorTool
from .datasource_filter import DataSourceFilterTool
from .datasource_filter_v2 import DataSourceFilterToolV2
from .data_seeker_intention_recognizer import DataSeekerIntentionRecognizerTool
from .data_seeker_report_writer import DataSeekerReportWriterTool
from .base import QueryIntentionName
from data_retrieval.tools.base import (
    ToolName,
    # QueryIntentionName,
    ToolMultipleResult,
    ToolResult,
    LogResult,
    AFTool,
    LLMTool,
    construct_final_answer,
    async_construct_final_answer,
    api_tool_decorator,
)

__all__ = [
    # Tools
    "AfSailorTool",
    "DataSourceFilterTool",
    # "DataSourceFilterToolV2",
    # "DataScopeCheckerTool",
    # "DataSeekerIntentionRecognizerTool",
    # "DataSeekerReportWriterTool",
    # Base classes and utilities
    "ToolName",
    "QueryIntentionName",
    "ToolMultipleResult",
    "ToolResult",
    "LogResult",
    "AFTool",
    "LLMTool",
    "construct_final_answer",
    "async_construct_final_answer",
    "api_tool_decorator",
]

_TOOLS_MAPPING = {
    "af_sailor": AfSailorTool,
    # "multi_query_search": MultiQuerySearchTool,
    "datasource_filter": DataSourceFilterTool,
    # "datasource_filter_v2": DataSourceFilterToolV2,
    # "data_scope_checker": DataScopeCheckerTool,
    # "data_seeker_intention_recognizer": DataSeekerIntentionRecognizerTool,
    # "data_seeker_report_writer": DataSeekerReportWriterTool,
}
