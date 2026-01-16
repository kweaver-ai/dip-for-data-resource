# -*- coding: utf-8 -*-
# @Time    : 2026/1/4 11:49
# @Author  : Glen.lv
# @File    : datasource_filter_v2
# @Project : af-agent

import json
import traceback
from io import StringIO
from textwrap import dedent
from typing import Optional, Type, Any, List, Dict
from collections import OrderedDict
from enum import Enum
import re
import asyncio

import pandas as pd
from data_retrieval.api.af_api import Services
from app.tools.search_tools.datasource_filter import DataSourceFilterTool
from langchain.tools import BaseTool
from langchain_core.callbacks import CallbackManagerForToolRun, AsyncCallbackManagerForToolRun
from langchain_core.pydantic_v1 import BaseModel, Field, PrivateAttr
from langchain_core.prompts import (
    ChatPromptTemplate,
    HumanMessagePromptTemplate
)
from langchain_core.messages import HumanMessage, SystemMessage

from pandas import Timestamp
from data_retrieval.logs.logger import logger
from data_retrieval.sessions import BaseChatHistorySession, CreateSession
from data_retrieval.tools.base import ToolName
from data_retrieval.tools.base import ToolResult, ToolMultipleResult, LLMTool, _TOOL_MESSAGE_KEY
from data_retrieval.tools.base import construct_final_answer, async_construct_final_answer
from data_retrieval.errors import Json2PlotError, ToolFatalError
from data_retrieval.tools.base import api_tool_decorator

# from data_retrieval.datasource.data_view import AFDataSource
from data_retrieval.datasource.af_indicator import AFIndicator

from data_retrieval.prompts.tools_prompts.datasource_filter_prompt import DataSourceFilterPrompt
from data_retrieval.utils.model_types import ModelType4Prompt
from data_retrieval.parsers.base import BaseJsonParser

from app.utils.llm_utils import estimate_tokens_safe

from fastapi import FastAPI, HTTPException


class DataSourceDescSchema(BaseModel):
    id: str = Field(description="数据资源的 id, 为一个字符串")
    title: str = Field(description="数据资源的名称")
    type: str = Field(description="数据资源的类型")
    description: str = Field(description="数据源的描述")
    columns: Any = Field(default=None, description="数据源的字段信息")


class ArgsModel(BaseModel):
    query: str = Field(default="", description="用户的完整查询需求，如果是追问，则需要根据上下文总结")
    search_tool_cache_key: str = Field(default="", description=f"""是前几轮问答 {ToolName.from_sailor.value} 工具结果的缓存 key，
    对应`search`工具结果的'result_cache_key','result_cache_key'形如'68a8a4f4b83c32adc3146acdb7b0ef40_CswwizwiRsmRsJ9BV69Gmg',
    不能编造该信息; 注意不是'数据资源的 ID',形如'7ce014e4-6f7e-4e4e-bcf7-6c7a09e339a7',不要混淆!""")
    # data_resource_list: Optional[List[str]] = Field(default=[], description=f"数据源的列表, 每个列表都是一个字典, 格式为: {DataSourceDescSchema.schema_json(ensure_ascii=False)}")


class DataSourceFilterToolV2(DataSourceFilterTool):
    name: str = ToolName.from_datasource_filter.value
    description: str = dedent(
        f"""数据资源过滤工具，如果用户针对上一论问答的结果做进一步追问的时候，可以使用该工具。一定要注意如果本轮问答使用了 {ToolName.from_sailor.value} 工具, 就不能再使用该工具!

参数:
- query: 查询语句
- search_tool_cache_key: 是前几轮问答 {ToolName.from_sailor.value} 工具结果的缓存 key，对应`search`工具结果的'result_cache_key'
,'result_cache_key'形如'68a8a4f4b83c32adc3146acdb7b0ef40_CswwizwiRsmRsJ9BV69Gmg',不能编造该信息; 注意不是'数据资源的 ID',
形如'7ce014e4-6f7e-4e4e-bcf7-6c7a09e339a7',不要混淆。

如果没有 search_tool_cache_key 信息(形如'68a8a4f4b83c32adc3146acdb7b0ef40_CswwizwiRsmRsJ9BV69Gmg'),你需要仔细甄别，千万不要将数据
资源的 ID(形如'7ce014e4-6f7e-4e4e-bcf7-6c7a09e339a7') 作为 search_tool_cache_key , 否则会出现严重错误!

注意: 在同一次问答中, 该工具不能与 {ToolName.from_sailor.value} 工具同时使用,只能使用其中的一个。
"""
    )
    args_schema: Type[BaseModel] = ArgsModel
    # with_sample: bool = True
    with_sample: bool = False
    data_resource_num_limit: int = -1  # 数据资源数量上限，-1代表不限制
    dimension_num_limit: int = -1  # 字段（维度）数量上限，-1代表不限制
    session_type: str = "redis"
    session: Optional[BaseChatHistorySession] = None
    batch_size: int = 10  # map-reduce 批次大小，每批处理的数据源数量（回退方案）
    max_tokens_per_chunk: Optional[int] = None  # 每个批次的最大 token 数，如果设置则按 token 数分块
    search_configs: Optional[Any] = None
    service: Any = None
    base_url: str = ""
    headers: Dict[str, str] = Field(default_factory=dict)  # HTTP 请求头

    token: str = ""
    user_id: str = ""
    background: str = ""


    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        logger.info(f'*args={args}, \n**kwargs={kwargs}')
        self.service = Services(base_url=self.base_url)
        if kwargs.get("session") is None:
            self.session = CreateSession(self.session_type)
        if kwargs.get("batch_size") is not None:
            self.batch_size = kwargs.get("batch_size")
        if kwargs.get("max_tokens_per_chunk") is not None:
            self.max_tokens_per_chunk = kwargs.get("max_tokens_per_chunk")
        if kwargs.get("search_configs") is not None:
            logger.info(f'search_configs={kwargs.get("search_configs")}')
            self.search_configs = kwargs.get("search_configs")
        else:
            raise ToolFatalError("search_configs is required")
        
        # 初始化 headers，使用 token 字段（可能来自 kwargs 或类默认值）
        token_value = kwargs.get("token") or self.token
        if token_value:
            self.headers = {"Authorization": token_value}
        else:
            self.headers = {}

    def _split_into_batches(
            self, 
            data_list: List[Any], 
            query: str,
            data_resource_list_description: str = ""
    ) -> List[List[Any]]:
        """
        将列表分成多个批次，优先按 token 数分块，失败则回退到按数量分块
        
        Args:
            data_list: 数据源列表
            query: 用户查询
            data_resource_list_description: 数据源列表描述
            
        Returns:
            分块后的列表
        """
        chunks = []
        
        if self.max_tokens_per_chunk is not None:
            try:
                # 估算固定 token 数：query + prompt模板 + 输出预留
                # query约100 tokens, data_resource_list_description约500 tokens, prompt约500 tokens, 输出预留1000 tokens
                estimated_fixed_tokens = 2100
                
                # 获取模型的最大 token 长度（如果有的话）
                model_max_tokens = None
                if self.llm and hasattr(self.llm, 'max_tokens'):
                    model_max_tokens = self.llm.max_tokens
                elif self.llm and hasattr(self.llm, 'model_kwargs') and 'max_tokens' in self.llm.model_kwargs:
                    model_max_tokens = self.llm.model_kwargs.get('max_tokens')
                
                # 计算可用 token 数
                if model_max_tokens:
                    available_tokens = min(
                        self.max_tokens_per_chunk, 
                        int(model_max_tokens) - estimated_fixed_tokens
                    )
                else:
                    available_tokens = self.max_tokens_per_chunk - estimated_fixed_tokens
                
                # 确保 available_tokens 至少为 100，避免过小的值
                if available_tokens < 100:
                    logger.warning(f'计算出的 available_tokens ({available_tokens}) 过小，使用默认值 1000')
                    available_tokens = 1000
                
                logger.info(f'max_tokens_per_chunk = {self.max_tokens_per_chunk}')
                logger.info(f'model_max_tokens = {model_max_tokens}')
                logger.info(f'available_tokens = {available_tokens}')
                
                # 按 token 数分块
                current_chunk = []
                current_chunk_tokens = 0

                for item in data_list:
                    # 估算当前项的 token 数
                    item_str = json.dumps(item, ensure_ascii=False, separators=(',', ':'))
                    estimated_tokens = estimate_tokens_safe(item_str)

                    # 如果单个数据项的 token 数就超过了可用 token，仍然添加到当前块（避免无限循环）
                    if estimated_tokens > available_tokens:
                        logger.warning(f'单个数据项的 token 数 ({estimated_tokens}) 超过可用 token ({available_tokens})，仍将添加到当前块')

                    if current_chunk_tokens + estimated_tokens > available_tokens and current_chunk:
                        # 当前块已满，开始新块
                        chunks.append(current_chunk)
                        current_chunk = [item]
                        current_chunk_tokens = estimated_tokens
                    else:
                        current_chunk.append(item)
                        current_chunk_tokens += estimated_tokens

                if current_chunk:
                    chunks.append(current_chunk)

                logger.info(f'按token数分块，共 {len(chunks)} 块，每块约 {available_tokens:.0f} tokens')
                
            except Exception as e:
                logger.warning(f'按token数分块失败: {e}，回退到按数量分块')
                chunks = [data_list[i:i + self.batch_size] for i in range(0, len(data_list), self.batch_size)]
        else:
            # 按数量分块
            chunks = [data_list[i:i + self.batch_size] for i in range(0, len(data_list), self.batch_size)]
        
        logger.info(f'数据分块完成，共 {len(chunks)} 块，每块约 {len(chunks[0]) if chunks else 0} 个数据项')
        return chunks

    def _config_chain(
            self,
            data_resource_list: List[dict] = [],
            data_resource_list_description: str = "",
            skip_refresh_cache_key: bool = False
    ):
        # 刷新结果缓存key（除非明确要求跳过）
        if not skip_refresh_cache_key:
            self.refresh_result_cache_key()

        system_prompt = DataSourceFilterPrompt(
            data_source_list=data_resource_list,
            prompt_manager=self.prompt_manager,
            language=self.language,
            data_source_list_description=data_resource_list_description,
            background=self.background
        )

        logger.debug(f"{ToolName.from_datasource_filter.value} -> model_type: {self.model_type}")

        if self.model_type == ModelType4Prompt.DEEPSEEK_R1.value:
            prompt = ChatPromptTemplate.from_messages(
                [
                    HumanMessage(
                        content="下面是你的任务，请务必牢记" + system_prompt.render(),
                        additional_kwargs={_TOOL_MESSAGE_KEY: ToolName.from_datasource_filter.value}
                    ),
                    HumanMessagePromptTemplate.from_template("{input}")
                ]
            )
        else:
            prompt = ChatPromptTemplate.from_messages(
                [
                    SystemMessage(
                        content=system_prompt.render(),
                        additional_kwargs={_TOOL_MESSAGE_KEY: ToolName.from_datasource_filter.value}
                    ),
                    HumanMessagePromptTemplate.from_template("{input}")
                ]
            )

        chain = (
                prompt
                | self.llm
                | BaseJsonParser()
        )
        return chain

    @construct_final_answer
    def _run(
            self,
            input: str,
            search_tool_cache_key: Optional[str] = "",
            # data_resource_list: Optional[List[str]] = [],
            run_manager: Optional[CallbackManagerForToolRun] = None,
    ):
        return asyncio.run(self._arun(
            query=input,
            search_tool_cache_key=search_tool_cache_key,
            #  data_resource_list=data_resource_list,
            run_manager=run_manager)
        )

    @async_construct_final_answer
    async def _arun(
            self,
            query: str,
            search_tool_cache_key: Optional[str] = "",
            # data_resource_list: Optional[List[str]] = [],
            run_manager: Optional[AsyncCallbackManagerForToolRun] = None,
    ):
        data_view_list, metric_list = OrderedDict(), OrderedDict()
        data_view_metadata, metric_metadata = {}, {}
        data_resource_list = []
        # data_source_list 改名为 data_resource_list
        if search_tool_cache_key:
            tool_res = self.session.get_agent_logs(
                search_tool_cache_key
            )
            if tool_res:
                data_resource_list = tool_res.get("cites", [])
                data_resource_list_description = tool_res.get("description", "")
            else:
                return {
                    "result": f"搜索工具的缓存 key 不存在: {search_tool_cache_key}"
                }

        # cites和最终处理完渲染提示词的变量 data_source_list 结构几乎完全相同，把 cites 中的 "fields" 字段改名为 "columns"

        # 转换 cites 格式为 data_source_list 格式：将 fields 改名为 columns
        def convert_cites_to_data_source_list(cites: List[dict]) -> List[dict]:
            """
            将 cites 格式转换为 data_source_list 格式
            主要转换：将 'fields' 字段改名为 'columns'
            
            Args:
                cites: cites 格式的数据源列表，包含 'fields' 字段
                
            Returns:
                data_source_list: 转换后的数据源列表，包含 'columns' 字段
            """
            data_source_list = []
            for cite in cites:
                # 创建新的字典，复制所有字段
                data_source = cite.copy()
                
                # 如果存在 'fields' 字段，改名为 'columns'
                if 'fields' in data_source:
                    data_source['columns'] = data_source.pop('fields')
                # 如果已经存在 'columns' 字段，保持不变
                # 如果两者都不存在，保持原样（后续可能会通过 API 获取）
                
                data_source_list.append(data_source)
            
            return data_source_list
        
        # 将 cites 格式转换为 data_source_list 格式
        data_resource_list = convert_cites_to_data_source_list(data_resource_list)

        # data_view_list 是cites中逻辑视图部分
        # metric_list 是cites中指标部分
        for data_resource in data_resource_list:
            if data_resource["type"] == "data_view":
                data_view_list[data_resource["id"]] = data_resource
            elif data_resource["type"] == "indicator":
                metric_list[data_resource["id"]] = data_resource
            else:
                return {
                    "result": f"数据资源类型错误: {data_resource['type']}"
                }

        logger.info(f'data_view_list={data_view_list}')

        if len(data_view_list) > 0:
            # 需要在 description 中加入样例数据和探查结果枚举值
            for hit in data_view_list.values():

                view_column_info_for_prompt, view_source_catalog_name = await self.service.get_view_column_info_for_prompt(
                    idx=hit['id'],
                    headers=self.headers
                )

                view_details = self.service.get_view_details_by_id(
                    view_id=hit['id'],
                    headers=self.headers
                )

                source, schema = view_source_catalog_name.split('.')
                source_dict = {
                    "source": source,  # 数据源在虚拟化引擎中的 catalog, 数据源配置的时候已经固定数据库了
                    "schema": schema,  # 固定都是default
                    "title": view_details['technical_name']  # 表名
                }

                sample = await self.service.get_view_sample_by_source(
                    source=source_dict,
                    headers=self.headers
                )

                try:
                    data_explore_rst = await self.service.get_data_explore(
                        entity_id=hit['id'],
                        headers=self.headers
                    )
                except Exception as e:
                    logger.warning(f'get_data_explore warning: 探查报告不存在')
                    data_explore_rst = []
                if data_explore_rst:
                    logger.info(f'data_explore_rst = {data_explore_rst}')

                # 构建字段ID到枚举值的映射
                field_enum_map = {}
                for explore_item in data_explore_rst:
                    field_id = explore_item['field_id']
                    if explore_item.get('details') and len(explore_item['details']) > 0:
                        result_str = explore_item['details'][0].get('result', '[]')
                        try:
                            enum_data = json.loads(result_str)
                            # 提取枚举值列表（只取key，过滤掉null值）
                            enum_values = [item['key'] for item in enum_data if item.get('key') is not None]
                            if enum_values:  # 只有当枚举值不为空时才添加
                                field_enum_map[field_id] = enum_values
                        except Exception as e:
                            pass

                # 构建技术名称到样例数据的映射
                sample_data = sample['data'][0] if sample.get('data') and len(sample['data']) > 0 else []
                sample_columns = sample.get('columns', [])
                field_sample_map = {}
                for idx, col in enumerate(sample_columns):
                    if idx < len(sample_data):
                        field_sample_map[col['name']] = sample_data[idx]

                # 整合字段信息、样例数据和枚举值
                def build_field_info():
                    """整合字段信息、样例数据和枚举值"""
                    fields_info = []
                    for field_info in view_column_info_for_prompt:
                        field_id = field_info['id']
                        technical_name = field_info['technical_name']

                        # 获取样例数据
                        sample_value = field_sample_map.get(technical_name, None)

                        # 获取枚举值
                        enum_values = field_enum_map.get(field_id, None)

                        field_data = {
                            'business_name': field_info['business_name'],
                            'technical_name': technical_name,
                            'comment': field_info['comment'],
                            'data_type': field_info['data_type'],
                            'sample_value': sample_value,
                            'enum_values': enum_values
                        }
                        fields_info.append(field_data)

                    return fields_info

                fields_info = build_field_info()
                
                # 格式化字段信息为长字符串
                def format_to_long_string(fields_info: List[dict], table_description: str = "") -> str:
                    """生成长字符串格式的提示词（包含表描述、字段信息、样例数据、枚举值）"""
                    # 开始构建字段信息部分
                    fields_text = ""

                    for idx, field in enumerate(fields_info, 1):
                        # 字段基本信息
                        field_info = f"{field['business_name']}（技术名称：{field['technical_name']}，数据类型：{field['data_type']}"

                        # 添加样例值
                        if field['sample_value'] is not None:
                            sample_str = str(field['sample_value'])
                            # 如果样例值太长，截断
                            if len(sample_str) > 80:
                                sample_str = sample_str[:80] + "..."
                            field_info += f"，样例值：{sample_str}"
                        else:
                            field_info += "，样例值：无"

                        # 添加枚举值（如果有）
                        if field['enum_values']:
                            max_enum_display = 6  # 枚举值最多显示6个，避免过长
                            if len(field['enum_values']) <= max_enum_display:
                                enum_str = '、'.join(map(str, field['enum_values']))
                                field_info += f"，枚举值：{enum_str}（共{len(field['enum_values'])}个）"
                            else:
                                enum_str = '、'.join(map(str, field['enum_values'][:max_enum_display]))
                                field_info += f"，枚举值：{enum_str}等（共{len(field['enum_values'])}个）"

                        field_info += "）"

                        if idx < len(fields_info):
                            fields_text += field_info + "；"
                        else:
                            fields_text += field_info

                    # 组合成完整的长字符串
                    full_prompt = f"{table_description} 字段信息、样例数据、部分字段的枚举值如下：{fields_text}"

                    return full_prompt
                
                # 获取原始描述
                original_description = hit.get('description', '')
                description_append_fields_info = format_to_long_string(
                    fields_info,
                    table_description=original_description
                )
                hit['description'] = description_append_fields_info



        if len(metric_list) > 0:
            metric_source = AFIndicator(
                indicator_list=list(metric_list.keys()),
                token=self.token,
                user_id=self.user_id
            )
            try:
                metric_metadata = metric_source.get_details(
                    input_query=query,
                    indicator_num_limit=self.data_resource_num_limit,
                    input_dimension_num_limit=self.dimension_num_limit
                )

                for k, v in metric_list.items():
                    for detail in metric_metadata["details"]:
                        if detail["id"] == k:
                            v["columns"] = {
                                dimension["technical_name"]: dimension["business_name"]
                                for dimension in detail.get("dimensions", [])
                            }
                            break

            except Exception as e:
                logger.error(f"获取指标元数据失败: {str(e)}")
                # raise ToolFatalError(f"获取指标元数据失败: {e}")

        if not data_view_list and not metric_list:
            return {
                "result": f"没有找到符合要求的数据源"
            }
            # raise ToolFatalError(f"没有找到符合要求的数据源")

        # 合并所有数据源
        all_data_resources = list(data_view_list.values()) + list(metric_list.values())
        view_ids = [data_view["id"] for data_view in data_view_list.values()]
        metric_ids = [metric["id"] for metric in metric_list.values()]

        # 在开始处理前刷新一次 cache key，确保所有批次使用同一个 key
        self.refresh_result_cache_key()

        estimated_tokens = estimate_tokens_safe(str(all_data_resources))
        logger.info(f'estimated_tokens = {estimated_tokens}')
        


        try:
            # Map-Reduce 模式：将数据源列表分成多个批次处理
            # 如果设置了 max_tokens_per_chunk 或数据源数量超过 batch_size，则使用 map-reduce
            # use_map_reduce = (
            #     self.max_tokens_per_chunk is not None or
            #     (len(all_data_resources) > self.batch_size and self.batch_size > 0)
            # )
            # 如果 search_configs 存在，使用它来计算 max_tokens_per_chunk
            if self.search_configs and hasattr(self.search_configs, 'sailor_search_qa_llm_input_len'):
                logger.info(
                    f'search_configs.sailor_search_qa_llm_input_len = {self.search_configs.sailor_search_qa_llm_input_len}')
                calculated_max_tokens = int(self.search_configs.sailor_search_qa_llm_input_len) * 0.8
                # 如果用户没有手动设置 max_tokens_per_chunk，则使用计算出的值
                if self.max_tokens_per_chunk is None:
                    self.max_tokens_per_chunk = int(calculated_max_tokens)
                    logger.info(f'根据 search_configs 自动设置 max_tokens_per_chunk = {self.max_tokens_per_chunk}')
                use_mapreduce = estimated_tokens > calculated_max_tokens
            else:
                use_mapreduce = False

            # 判断是否使用 map-reduce 模式
            use_map_reduce = (
                use_mapreduce or
                self.max_tokens_per_chunk is not None or 
                (len(all_data_resources) > self.batch_size and self.batch_size > 0)
            )
            
            if use_map_reduce:
                if self.max_tokens_per_chunk is not None:
                    logger.info(f"使用 map-reduce 模式处理，按 token 数分块 (max_tokens_per_chunk={self.max_tokens_per_chunk})")
                else:
                    logger.info(f"数据源数量 ({len(all_data_resources)}) 超过批次大小 ({self.batch_size})，使用 map-reduce 模式处理")
                
                # 将数据源列表分成多个批次（按 token 数或数量）
                batches = self._split_into_batches(
                    all_data_resources,
                    query=query,
                    data_resource_list_description=data_resource_list_description
                )
                logger.info(f"共分成 {len(batches)} 个批次进行处理")
                
                # Map 阶段：对每个批次并行处理
                async def process_batch(batch: List[dict], batch_index: int, total_batches: int) -> dict:
                    """处理单个批次的异步函数"""
                    logger.info(f"处理第 {batch_index+1}/{total_batches} 批次，包含 {len(batch)} 个数据源")
                    # 注意：这里不调用 refresh_result_cache_key，使用之前统一的 cache key
                    chain = self._config_chain(
                        data_resource_list=batch,
                        data_resource_list_description=data_resource_list_description,
                        skip_refresh_cache_key=True  # 跳过刷新 cache key
                    )
                    return await chain.ainvoke({"input": query})
                
                # 并行处理所有批次
                batch_tasks = [
                    process_batch(batch, i, len(batches))
                    for i, batch in enumerate(batches)
                ]
                batch_results = await asyncio.gather(*batch_tasks)
                
                # Reduce 阶段：合并所有批次的结果
                result_datasource_list = []
                for batch_result in batch_results:
                    if batch_result and "result" in batch_result:
                        for res in batch_result["result"]:
                            if res["id"] in view_ids:
                                # 结果中补充 title
                                res["title"] = data_view_list[res["id"]].get("title", "")
                                result_datasource_list.append(res)
                            elif res["id"] in metric_ids:
                                res["title"] = metric_list[res["id"]].get("title", "")
                                result_datasource_list.append(res)
                
                logger.info(f"map-reduce 处理完成，共筛选出 {len(result_datasource_list)} 个数据源")
            else:
                # 如果数据源数量不超过批次大小，使用原来的单次处理方式
                logger.info(f"数据源数量 ({len(all_data_resources)}) 未超过批次大小，使用单次处理")
                chain = self._config_chain(
                    data_resource_list=all_data_resources,
                    data_resource_list_description=data_resource_list_description,
                    skip_refresh_cache_key=True  # 使用之前统一的 cache key
                )
                result = await chain.ainvoke({"input": query})

                result_datasource_list = []
                for res in result["result"]:
                    if res["id"] in view_ids:
                        # 结果中补充 title
                        res["title"] = data_view_list[res["id"]].get("title", "")
                        result_datasource_list.append(res)
                    elif res["id"] in metric_ids:
                        res["title"] = metric_list[res["id"]].get("title", "")
                        result_datasource_list.append(res)

            logger.info(f"result_datasource_list: {result_datasource_list}")

            self.session.add_agent_logs(
                self._result_cache_key,
                logs={
                    "result": result_datasource_list,
                    "cites": [
                        {
                            "id": data_resource["id"],
                            "type": data_resource["type"],
                            "title": data_resource["title"],
                        } for data_resource in result_datasource_list
                    ]
                }
            )
        except Exception as e:
            logger.error(f"获取数据源失败: {str(e)}")
            raise ToolFatalError(f"获取数据源失败: {str(e)}")

        # 给大模型的数据
        return {
            "result": result_datasource_list,
            "result_cache_key": self._result_cache_key
        }

    def handle_result(
            self,
            result_cache_key: str,
            log: Dict[str, Any],
            ans_multiple: ToolMultipleResult
    ) -> None:
        tool_res = self.session.get_agent_logs(
            result_cache_key
        )
        if tool_res:
            log["result"] = tool_res

            # 替换 cites
            if tool_res.get("cites"):
                ans_multiple.cites = tool_res.get("cites", [])


if __name__ == "__main__":
    import asyncio
    from langchain_openai import ChatOpenAI
    from af_agent.prompts.manager.base import BasePromptManager
    from af_agent.sessions.in_memory_session import InMemoryChatSession
    
    # 创建模拟的数据源列表（基于用户提供的示例）
    test_data_resource_list = [

    ]
    
    # 创建模拟的 session，用于存储测试数据
    test_session = InMemoryChatSession()
    test_cache_key = "test_search_cache_key_12345"
    
    # 将测试数据存储到 session 中
    test_session.add_agent_logs(
        test_cache_key,
        logs={
            "cites": test_data_resource_list,
            "description": "测试数据源列表描述：包含部门、位置、等相关数据源"
        }
    )
    
    print("=" * 80)
    print("测试 1: 按 token 数分块")
    print("=" * 80)
    search_configs={'sailor_search_if_history_qa_enhance': '0', 'sailor_search_if_kecc': '1',
     'sailor_search_if_auth_in_find_data_qa': '0', 'direct_qa': 'false', 'sailor_vec_min_score_analysis_search': '0.5',
     'sailor_vec_knn_k_analysis_search': '100', 'sailor_vec_size_analysis_search': '100',
     'sailor_vec_min_score_kecc': '0.5', 'sailor_vec_knn_k_kecc': '20', 'sailor_vec_size_kecc': '20',
     'kg_id_kecc': '19475', 'sailor_vec_min_score_history_qa': '0.7', 'sailor_vec_knn_k_history_qa': '10',
     'sailor_vec_size_history_qa': '10', 'kg_id_history_qa': '19467', 'sailor_token_tactics_history_qa': '1',
     'sailor_search_qa_llm_temperature': '0.0000001', 'sailor_search_qa_llm_top_p': '1',
     'sailor_search_qa_llm_presence_penalty': '0', 'sailor_search_qa_llm_frequency_penalty': '0',
     'sailor_search_qa_llm_max_tokens': '16000', 'sailor_search_qa_llm_input_len': '8000',
     'sailor_search_qa_llm_output_len': '8000', 'sailor_search_qa_cites_num_limit': '100'}
    
    # 测试1: 按 token 数分块
    tool1 = DataSourceFilterToolV2(
        session=test_session,
        max_tokens_per_chunk=2000,  # 设置较小的 token 限制，确保会分块
        batch_size=10,
        search_configs=search_configs
    )
    
    # 设置 LLM 和 prompt_manager（可选，仅用于测试分块逻辑）
    tool1.llm = ChatOpenAI(
        model_name="Qwen-72B-Chat",
        openai_api_key="EMPTY",
        openai_api_base="http://192.168.173.19:8304/v1",
        max_tokens=8000,
        temperature=0
    )
    tool1.prompt_manager = BasePromptManager()
    
    # 测试分块功能
    query = "查找某部门相关信息"
    chunks = tool1._split_into_batches(
        test_data_resource_list,
        query=query,
        data_resource_list_description="测试数据源列表描述"
    )
    
    print(f"\n数据源总数: {len(test_data_resource_list)}")
    print(f"分块数量: {len(chunks)}")
    for i, chunk in enumerate(chunks):
        chunk_str = json.dumps(chunk, ensure_ascii=False, separators=(',', ':'))
        estimated_tokens = estimate_tokens_safe(chunk_str)
        print(f"  批次 {i+1}: {len(chunk)} 个数据源, 约 {estimated_tokens} tokens")
    
    print("\n" + "=" * 80)
    print("测试 2: 按数量分块（回退方案）")
    print("=" * 80)
    
    # 测试2: 按数量分块
    tool2 = DataSourceFilterToolV2(
        session=test_session,
        max_tokens_per_chunk=None,  # 不设置 token 限制，使用数量分块
        batch_size=2,        # 每批2个
        search_configs=search_configs
    )
    
    chunks2 = tool2._split_into_batches(
        test_data_resource_list,
        query=query,
        data_resource_list_description="测试数据源列表描述"
    )
    
    print(f"\n数据源总数: {len(test_data_resource_list)}")
    print(f"分块数量: {len(chunks2)}")
    for i, chunk in enumerate(chunks2):
        print(f"  批次 {i+1}: {len(chunk)} 个数据源")
    
    print("\n" + "=" * 80)
    print("测试 3: 测试单个数据源的 token 估算")
    print("=" * 80)
    
    # 测试单个数据源的 token 估算
    for i, data_resource in enumerate(test_data_resource_list[:2]):  # 只测试前2个
        item_str = json.dumps(data_resource, ensure_ascii=False, separators=(',', ':'))
        estimated_tokens = estimate_tokens_safe(item_str)
        print(f"\n数据源 {i+1} ({data_resource['id'][:8]}...):")
        print(f"  Title: {data_resource['title'][:50]}...")
        print(f"  估算 token 数: {estimated_tokens}")
        print(f"  JSON 长度: {len(item_str)} 字符")
    
    print("\n" + "=" * 80)
    print("测试 4: cites 格式转换为 data_source_list 格式")
    print("=" * 80)
    
    # 测试 cites 格式转换为 data_source_list 格式
    test_cites = [

    ]
    
    # 使用工具类中的转换函数（需要创建一个临时实例来访问方法）
    # 或者直接定义转换函数
    def convert_cites_to_data_source_list(cites: List[dict]) -> List[dict]:
        """将 cites 格式转换为 data_source_list 格式"""
        data_source_list = []
        for cite in cites:
            data_source = cite.copy()
            if 'fields' in data_source:
                data_source['columns'] = data_source.pop('fields')
            data_source_list.append(data_source)
        return data_source_list
    
    converted_list = convert_cites_to_data_source_list(test_cites)
    
    print(f"\ncites 格式数据源数量: {len(test_cites)}")
    print(f"转换后 data_source_list 数量: {len(converted_list)}")
    print("\n转换示例:")
    for i, (cite, converted) in enumerate(zip(test_cites, converted_list)):
        print(f"\n  数据源 {i+1}:")
        print(f"    cites 格式 - 字段名: {'fields' if 'fields' in cite else '无'}")
        print(f"    转换后格式 - 字段名: {'columns' if 'columns' in converted else '无'}")
        if 'fields' in cite:
            print(f"    cites.fields: {cite['fields']}")
        if 'columns' in converted:
            print(f"    data_source_list.columns: {converted['columns']}")
    
    print("\n" + "=" * 80)
    print("测试完成！")
    print("=" * 80)
    

