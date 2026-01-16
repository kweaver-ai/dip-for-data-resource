import time
from enum import Enum
from typing import Callable, Optional

import pandas as pd
from langchain.tools import BaseTool
from langchain_community.tools import HumanInputRun
from langchain_core.callbacks import CallbackManagerForToolRun, AsyncCallbackManagerForToolRun

from app.llm.chat_qa_model import get_llm
from app.models.agent_models import LogsPrefix
from app.models.tool_retrieval_models import AFInterfaceToolModel, AFText2SQLToolModel, AFKnowledgeSearchToolModel, \
    AFKnowledgeEnhancementToolModel, AfCopilotToolModel, AfSailorToolModel, JsonToPlotToolModel
from app.retriever.data_property_retriever import retrieve_data_properties
from app.session.redis_session import RedisHistorySession
from app.utils.chat2plot import chat2plot
from app.utils.chat2plot.g2_schema import G2PlotConfig
from app.utils.parse_tools_utils import *


def _print_func(text: str) -> None:
    print("\n")
    print(text)


def web_input_func(query):
    # 从聊天窗口表单获取一个用户输入，并且加载服务端对应的对话状态管理内的对话记录到memory
    # TODO
    # return "《" + query + "》\n上面《》内的内容是一个回复给用户的问句，请将《》内的内容在其前面加上Final Answer: 回复给用户，同时结束此次问答。"
    res = res_frame("用户马上就会看到你的提问，请结束此次问答，等待用户下次提问")
    return res


class ToolName(Enum):
    AfSailorToolName = "根据问题获取相关数据信息的工具"
    JsonToPlotToolName = "json2plot"
    AfHumanInputToolName = "询问用户"


class AfHumanInputTool(HumanInputRun):
    name: str = ToolName.AfHumanInputToolName.value
    description = '''当通过观察observation后，你觉得难以难以得出下一步action，或者难以回答问题，或者需要补充信息，你可以询问用户来协助指引
    你思考。该工具的输入参数是一个你希望向人类提的用来协助引导你思考下一步Action的问题。每当使用此工具后，结束对话。'''
    prompt_func: Callable = _print_func
    input_func: Callable = web_input_func

    def _run(
        self,
        query: str,
        run_manager: Optional[CallbackManagerForToolRun] = None,
    ) -> str:
        """Use the Human input tool."""
        self.prompt_func(query)
        return self.input_func(query)


class InterfaceTool(BaseTool):
    name: str = Field(..., description="工具名字")
    func: AFServiceFunction = Field(..., description="工具的执行函数")
    description: str = Field(..., description="工具描述")
    args_schema: Type[BaseModel] = Field(..., description="一个pydantic BaseModel, 用于约束工具的输入格式")

    def _run(self, **kwargs):  # 工具的执行方法
        return self.func(**kwargs)

    async def _arun(self, **kwargs):  # 工具的异步执行方法
        return self.func(**kwargs)


def get_tools_from_af_service(url: str, token: str, query: str = ""):
    """
    转化并封装所有AF数据服务接口成为Agent可执行调用的工具
    """
    tools = []
    try:
        af_service = get_af_services(url, token, query)
    except TypeError:
        return "抱歉，AI提问在您有权限的数据中并没有找到此问题的答案，您可以用其他问题再次发起提问。"
    except KeyError:
        return "用户登录已过期, 请重新登录或者跟换 Authorization。"
    for info in af_service:
        try:
            func = get_func_from_af_service(info, url, token)
            args_schema = get_args_schema_from_af_service(info)
            tool = InterfaceTool(name=info.name,
                                 func=func,
                                 description=info.description,
                                 args_schema=args_schema,
                                 handle_tool_error=handle_tool_error)
            tools.append(tool)
        except KeyError:
            continue
    return tools


class AFInterfaceTool(BaseTool):
    name: str = "通过JSON调用接口工具"
    description: str = "当用户的问题是一个具体的可被查询的问题时，可尝试调用接口来获取答案"
    args_schema: Type[BaseModel] = AFInterfaceToolModel
    url: str = ""
    token: str = ""

    def __init__(self, url: str, token: str, **kwargs: Any):
        super().__init__(**kwargs)
        self.url = url
        self.token = token

    def _run(self, **kwargs):  # 工具的执行方法
        return exec_func_of_service(url=self.url, token=self.token, **kwargs)

    async def _arun(self, **kwargs):  # 工具的异步执行方法
        return exec_func_of_service(url=self.url, token=self.token, **kwargs)


class AFText2SQLTool(BaseTool):
    name: str = "通过问题进行数据库检索工具"
    description: str = "工具的输入时一个用户的自然语言问题，不需要额外的将问题转化为SQL语句"
    args_schema: Type[BaseModel] = AFText2SQLToolModel
    token: str = ""

    def __init__(self, token: str, **kwargs: Any):
        super().__init__(**kwargs)
        self.token = token

    def _run(self, **kwargs):  # 工具的执行方法
        return text2sql(token=self.token, **kwargs)

    async def _arun(self, **kwargs):  # 工具的异步执行方法
        return await text2sql(token=self.token, **kwargs)


class AFKnowledgeSearchTool(BaseTool):
    top: int = Field(..., description="召回结果的top数")
    url: str = Field(..., description="环境的ip")
    token: str = Field(..., description="鉴权字符串")
    name: str = "通过文本召回相关接口服务及表格工具"
    description: str = "这是一个知识检索工具，通过输入一个query，可以获取到与query相关的数据目录和接口目录，其中数据目录可以通过SQL进行查询，接口目录可以通过调用进行查询"
    args_schema: Type[BaseModel] = AFKnowledgeSearchToolModel

    def __init__(self, top: int, url: str, token: str):
        super().__init__()
        self.top = top
        self.url = url
        self.token = token

    def _run(self, **kwargs):  # 工具的执行方法
        return retrieve_data_properties(base_url=self.url, authorization=self.token, top=1, **kwargs)

    async def _arun(self, **kwargs):  # 工具的异步执行方法
        return retrieve_data_properties(base_url=self.url, authorization=self.token, top=1, **kwargs)


class AFKnowledgeEnhancementTool(BaseTool):
    name: str = "背景知识增强工具"
    description: str = "获取与问题相关的知识背景，以此加强对问题的理解"
    args_schema: Type[BaseModel] = AFKnowledgeEnhancementToolModel

    def __init__(self):
        super().__init__()

    def _run(self, **kwargs):  # 工具的执行方法
        return knowledge_enhancement(**kwargs)

    async def _arun(self, **kwargs):  # 工具的异步执行方法
        return knowledge_enhancement(**kwargs)


class AfCopilotTool(BaseTool):
    name: str = "通过问题进行知识检索工具"
    description: str = "入参是一个问题，问题中没有任何指代性词，可以获得问题的相关信息"
    args_schema: Type[BaseModel] = AfCopilotToolModel
    headers: dict = None

    def __init__(self, headers: dict, **kwargs: Any):
        super().__init__(**kwargs)
        self.headers = headers

    def _run(self, **kwargs):  # 工具的执行方法
        return exec_copilot_by_query(self.headers, **kwargs)

    async def _arun(self, **kwargs):  # 工具的异步执行方法
        return await exec_copilot_by_query(self.headers, **kwargs)


def res_frame(text: str):
    return {
        "result": {
            "status": "answer",
            "res": {"cites": [], "table": [], "df2json": [""], "explain": [], "chart": [], "text": [text]}
        }
    }


class AfSailorTool(BaseTool):
    name: str = ToolName.AfSailorToolName.value
    description: str = """工具返回包含以下信息：
    1）如果执行正常，会将结果保存在 redis，可供其它工具直接读取，也可供用户直接读取，
    2）如果执行异常，会返回异常的原因。
    """
    args_schema: Type[BaseModel] = AfSailorToolModel
    session: Any = ""

    def __init__(
        self,
        metadata,
        **kwargs: Any
    ):
        super().__init__(**kwargs)
        self.metadata = metadata
        self.session = RedisHistorySession()

    def _service(
        self,
        **kwargs: Any
    ):
        from app.cores.prompt.manage.api_base import API, HTTPMethod
        self.metadata["query"] = kwargs.get("question")
        if kwargs.get("extraneous_information", None) is not None:
            self.metadata["query"] = kwargs.get("question") + kwargs["extraneous_information"]

        api = API(
            url="http://af-sailor:9797/api/af-sailor/v1/assistant/qa",
            # url="http://10.4.119.109:5005/api/af-sailor/v1/assistant/qa",
            method=HTTPMethod.POST,
            headers=self.metadata["Authorization"],
            payload=self.metadata
        )
        return api

    def _run(
        self,
        run_manager: Optional[CallbackManagerForToolRun] = None,
        *args,
        **kwargs: Any
    ):
        try:
            api = self._service(**kwargs)
            res = json.loads(api.call())
            if res["result"]["res"]["table"] != [""]:
                self.session.add_agent_logs(
                    session_id=LogsPrefix.AfSailor.value + self.metadata["session_id"],
                    logs=res
                )
                res = res_frame(
                    "已经成功获取了到了相关信息，并经过校验，成功保存到 redis，可供作图工具直接使用，也可以结束此次问答，并且不反问用户"
                )
            else:
                if not self.metadata["direct_qa"]:
                    self.session.add_agent_logs(
                        session_id=LogsPrefix.AfSailor.value + self.metadata["session_id"],
                        logs=res
                    )
                    res = res_frame("请结束此次问答，并且不反问用户")
                else:
                    for cite in res["result"]["res"]["cites"]:
                        if cite.get("type") == "indicator":
                            self.session.add_agent_logs(
                                session_id=LogsPrefix.AfSailor.value + self.metadata["session_id"],
                                logs=res
                            )
                            res = res_frame("请结束此次问答，并且不反问用户")
                            return res
                    # text2sql 执行正常，但是没有出表格结果，保存一下，防止后面报错
                    self.session.add_agent_logs(
                        session_id=LogsPrefix.AfSailor.value + self.metadata["session_id"],
                        logs=res
                    )

        except Exception as e:
            res = res_frame(
                "抱歉，可能由于网络延迟或当前服务器繁忙，当前回答尚未完成。若您还需要继续获取更多信息，可重新发起搜索")
        return res

    async def _arun(
        self,
        run_manager: Optional[AsyncCallbackManagerForToolRun] = None,
        *args,
        **kwargs: Any
    ):
        api = self._service(**kwargs)
        res = json.loads(await api.call_async())
        res["result"]["res"]["table"] = ""
        return res


class JsonToPlotTool(BaseTool):
    name: str = ToolName.JsonToPlotToolName.value
    description: str = """输入是 query 和 chart_type，然后工具根据 query 要求 和 chart_type 数据进行绘图，，目前 chart_type 仅支持 柱状图， 折线图，饼图。
    工具返回：1）如果执行正常，会将结果保存在 redis，用户会直接读取，
            2）如果执行异常，会返回异常的原因。
    """
    args_schema: Type[BaseModel] = JsonToPlotToolModel
    session: Any = ""

    def __init__(
        self,
        metadata,
        **kwargs: Any
    ):
        super().__init__(**kwargs)
        self.metadata = metadata
        self.session = RedisHistorySession()

    def _parse(
        self,
        **kwargs: Any
    ):
        qwen_llm = get_llm()
        af_sailor_res = self.session.get_agent_logs(
            session_id=LogsPrefix.AfSailor.value + self.metadata["session_id"]
        )
        df2json = af_sailor_res["result"]["res"]["df2json"][0]
        df = pd.read_json(df2json)
        print("===============")
        print(df)
        print(df.columns.values)
        copy_df = df.iloc[:5][:]
        print(copy_df)
        start = time.time()
        c2p = chat2plot(copy_df, language="中文", schema_definition=G2PlotConfig, chat=qwen_llm)
        question = f'用{kwargs["chart_type"]} 来表示 {kwargs["question"]}'
        if kwargs.get("extraneous_information", None) is not None:
            question = f'用{kwargs["chart_type"]} 来表示 {kwargs["question"]}，{kwargs["extraneous_information"]}'
        print(f"============================ {question}")
        config = c2p(question).config.dict()
        datas = df.to_json(orient="records", force_ascii=False)
        print("+++++++++++", time.time() - start, "++++++++++++++++")
        print(config)
        print(list(config.values()))
        for column in list(config.values())[1:]:
            if column not in df.columns.values and column is not None:
                return res_frame("作图工具作图失败，原因是：使用了多维度进行作图，请你再次调用作图工具，并在问题中提醒工具，只能基于问题中的一个维度生成配置")


        af_sailor_res["result"]["res"]["chart"] = [{"data": datas, "config": config}]
        self.session.add_agent_logs(
            session_id=LogsPrefix.Chat2Plot.value + self.metadata["session_id"],
            logs=af_sailor_res
        )
        result = res_frame("已经成功获取了到了完整的图表，并经过校验，成功保存到 redis，可以被用户直接读取")

        return result

    def _run(
        self,
        run_manager: Optional[CallbackManagerForToolRun] = None,
        *args,
        **kwargs: Any
    ):
        try:
            return self._parse(**kwargs)
        except Exception as e:
            print(e)
            res = res_frame("工具使用失败，你必须先通过工具获取相关数据，然后才能使用作图工具")
            return res

    def _arun(
        self,
        run_manager: Optional[AsyncCallbackManagerForToolRun] = None,
        *args,
        **kwargs: Any
    ):
        return self._parse(**kwargs)


if __name__ == '__main__':
    p = {
        "df2json": "{\"status\\/状态\":{\"0\":\"null\",\"1\":\"-\",\"2\":\"冻结\",\"3\":\"失效\",\"4\":\"股权变更\",\"5\":\"解除冻结\"},\"状态数量\":{\"0\":2,\"1\":56,\"2\":604,\"3\":32,\"4\":2,\"5\":212}}",
        "question": "每一种股权冻结的状态有多少个",
        "chart_type": "柱状图"
    }
    p = {
        "question": "根据学期和专业统计参加考试的人数",
        "chart_type": "柱状图"

    }

    s = JsonToPlotTool({"session_id": "77777777777777777"}).invoke(p)
    print(s)
