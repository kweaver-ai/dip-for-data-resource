import json
from typing import List, Type, Any
from urllib.parse import urlsplit

import requests
import urllib3
from langchain_core.tools import ToolException
from pydantic import BaseModel, Field, create_model

from app.logs.logger import logger
from app.models.tool_retrieval_models import (
    AFDataServiceInfoModel,
    AFDataServiceParameterModel,
)
from app.utils.get_token import init_token

urllib3.disable_warnings()

# zkn 这是一个类型映射，AF的数据类型不被完全兼容
map_type = {
    "int": "int", "long": "int", "string": "str",
    "float": "float", "double": "float", "boolean": "bool"
}

map_required = {
    "yes": "true",
    "no": "false"
}


def get_af_services(base_url: str, authorization: str, query: str, top: int = 100) -> List[AFDataServiceInfoModel]:
    """
    获取与query最相似的AF数据服务接口，返回接口信息列表
    调用AF字符串召回数据服务接口，实现输入一个query，召回一组最相似数据服务的功能，返回接口信息列表，默认召回top1
    """
    # 获取与query最相似的服务接口，但是没有 url
    flag = 1  # 认知搜索
    headers = {"Authorization": authorization}
    search_url = f"{base_url}/api/data-catalog/frontend/v1/data-catalog/search/cog"
    body = {"size": top, "keyword": query, "asset_type": "interface_svc", }
    af_service_info = requests.post(url=search_url,
                                    headers=headers,
                                    json=body,
                                    verify=False)
    # logger.debug("采用认知搜索获取结果......")
    if af_service_info.json()["entries"] is None:
        flag = 2  # 备选路线
        logger.debug("认知搜索获取结果为空，采用普通检索......")
        search_url = f"{base_url}/api/data-application-service/frontend/v1/services/search"
        af_service_info = requests.post(url=search_url,
                                        headers=headers,
                                        json=body,
                                        verify=False)

    data_services = []
    if af_service_info.json()["entries"]:
        for entries in af_service_info.json()["entries"]:
            idx = entries['code'] if flag == 1 else entries['id']
            # # 通过接口获取 AF数据服务接口的参数, 并拼接 url  # if: interface
            get_if_url = f"{base_url}/api/data-application-service/frontend/v1/services/{idx}"
            if_info = requests.get(url=get_if_url,
                                   headers=headers,
                                   verify=False).json()
            if_path = f"{base_url}/data-application-gateway{if_info['service_info']['service_path']}"
            if_params = [AFDataServiceParameterModel(
                name=param["en_name"],
                type=param["data_type"],
                required=param["required"],
                default=param["default_value"],
                operator=param["operator"],
                description=param["cn_name"]
            ) for param in if_info["service_param"]["data_table_request_params"]]
            data_services.append(AFDataServiceInfoModel(
                name=entries["raw_title"],
                # code=entries["id"],
                code=idx,
                path=if_path,
                type=if_info["service_info"]["http_method"],
                description=entries["raw_description"],
                parameters=if_params))
    return data_services


class AFServiceFunction:
    def __init__(self, app_id, app_secret, url_path, url_type):
        self.app_id = app_id
        self.app_secret = app_secret
        self.url_path = url_path
        self.url_type = url_type

    def __call__(self, **kwargs):
        if self.url_type == "post":
            data = {key: value for key, value in kwargs.items()}
            url_token = init_token(
                app_id=self.app_id,
                app_secret=self.app_secret,
                method=self.url_type,
                path=urlsplit(self.url_path)[2],
                body=data)
            response = requests.post(
                url=self.url_path,
                json=data,
                headers={"Authorization": url_token},
                verify=False).json()
            # logger.debug("###########{}############".format(response.json()))
        else:
            data = {key: value for key, value in kwargs.items()}
            url_token = init_token(
                app_id=self.app_id,
                app_secret=self.app_secret,
                method=self.url_type,
                path=urlsplit(self.url_path)[2],
                qs=data)
            url_param = "&".join([f"{key}={value}" for key, value in kwargs.items()])
            self.url_path += f"?{url_param}"
            response = requests.get(url=self.url_path, headers={"Authorization": url_token}, verify=False).json()
        return response


def get_func_from_af_service(af_service_info: AFDataServiceInfoModel, url: str, token: str) -> AFServiceFunction:
    """
    将一个AF数据服务接口转化成远程调用的函数
    """
    auth_url = f"{url}/api/data-application-service/frontend/v1/services/{af_service_info.code}/auth_info"
    auth_info = requests.request(method="GET",
                                 url=auth_url,
                                 headers={"Authorization": token},
                                 verify=False).json()
    app_id = auth_info["app_id"]  # app_id
    app_secret = auth_info["app_secret"]  # app_secret
    url_path = af_service_info.path  # 路径
    url_type = af_service_info.type  # 参数
    af_service_function = AFServiceFunction(app_id, app_secret, url_path, url_type)

    return af_service_function


def create_args_schema(json_dict: dict) -> Type[BaseModel]:
    field_dict = {
        prop: (
            json_dict[prop]["type"],
            Field(
                default=json_dict[prop].get("default", None),
                description=json_dict[prop]["description"]
            )
        ) for prop in json_dict
    }

    dynamic_model = create_model(
        'AFInterfaceServiceParameterModel',
        **field_dict,
    )
    return dynamic_model


def get_args_schema_from_af_service(af_service_info: AFDataServiceInfoModel) -> Type[BaseModel]:
    """
    解析一个AF数据服务接口信息，转化并动态生成对应的pydantic model用于规范该工具的输入参数
    """
    properties = "".join([
        """"%s":{"type": "%s", "required": %s,"description": "%s", "default": "%s"},"""
        % (param.name, map_type[param.type], map_required[param.required],
           param.description, param.default)
        for param in af_service_info.parameters
    ])

    json_dict = """{"type": "object", "properties": {%s}}""" % properties[:-1]
    json_dict = json.loads(json_dict)
    args_schema = create_args_schema(json_dict["properties"])

    return args_schema


def handle_tool_error(error: ToolException) -> str:
    return ("下面这个工具未能执行成功:"
            + error.args[0]
            + "请重试其他的工具。")


async def text2sql(token, **kwargs) -> str:
    from app.cores.text2sql.t2s import Text2SQL
    demo = Text2SQL(
        user="foo",
        appid="Nr8KsyyoK8x8B1Nk-vO",
        query=kwargs["question"],
        headers={"Authorization": token}
    )
    resp = await demo.call()
    try:
        # result = resp["table"]
        # return result
        return resp
    except Exception as e:
        return str(e)


def exec_func_of_service(url: str, token: str, **kwargs) -> str:
    af_service = get_af_services(url, token, kwargs["interface_name"], top=3)
    func_dict = {}
    for info in af_service:
        try:
            func = get_func_from_af_service(info, url, token)
            func_dict[info.name] = func
        except KeyError:
            continue
    exec_interface = func_dict.get(kwargs["interface_name"])
    if not exec_interface:
        return "没有找到接口"
    else:
        interface_res = exec_interface(**kwargs["params"])
        return interface_res


def knowledge_enhancement(**kwargs) -> str:
    return "未找到匹配的背景知识。"


async def exec_copilot_by_query(headers, **kwargs) -> Exception | Any:
    from app.utils.copilot_inputs import inputs
    from app.cores.text2sql.t2s_base import API
    import json
    inputs["query"] = kwargs["query"]
    url = "http://10.4.119.109:5005/api/copilot/v1/assistant/qa"
    api = API(
        url=url,
        headers=headers,
        method="POST",
        payload=inputs
    )
    try:
        resp = await api.call_async()
        resp = json.loads(resp)
        resp = resp["result"]["res"]
        return resp["text"] + resp["table"]
    except Exception as e:
        return e
