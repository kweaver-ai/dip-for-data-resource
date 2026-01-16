import json

import requests
import urllib3

from app.logs.logger import logger

urllib3.disable_warnings()


def retrieve_data_properties(base_url: str, authorization: str, top: int = 1, **kwargs):
    # 输入一个字符串，召回所有相关数据资产的名称与描述,top控制每一个数据资产类型的最大数量，默认1，即每种资产最多召回1个
    logger.debug("采用认知搜索获取结果......")
    body = {"size": top, "keyword": kwargs["search_str"]}
    headers = {"Authorization": authorization}
    try:
        af_service_info = requests.request(
            method="POST", headers=headers, json=body, verify=False,
            url=f"{base_url}/api/data-catalog/frontend/v1/data-catalog/search/cog"
        )
        database = []
        services = []
        if af_service_info.json()["entries"] is None:
            logger.debug("认知搜索获取结果为空，采用普通检索......")
            af_database_info = requests.request(
                method="POST", headers=headers, json=body, verify=False,
                url=f"{base_url}/api/data-catalog/frontend/v1/data-catalog/search"
            )
            af_service_info = requests.request(
                method="POST", headers=headers, json=body, verify=False,
                url=f"{base_url}/api/data-application-service/frontend/v1/services/search"
            )
            try:
                for entries in af_database_info.json()["entries"]:
                    database.append(
                        {
                            "数据资产名": entries["raw_title"],
                            "描述": entries["raw_description"],
                            "输入参数描述": get_column_info(entries['id'], headers, base_url)
                        }
                    )
            except TypeError:
                pass
            try:
                for entries in af_service_info.json()["entries"]:
                    services.append(
                        {
                            "中文名": entries["raw_name"],
                            "描述": entries["raw_description"],
                            "输入参数描述": get_interface_parameter_info(entries['id'], headers, base_url),
                            "是否具有接口的调用权限": "具有" if entries["audit_status"] == "pass" else "不具有"
                        }
                    )
            except TypeError:
                pass
        else:
            for entries in af_service_info.json()["entries"]:
                if entries["type"] == "interface_svc":
                    audit_status = requests.get(
                        url=f"{base_url}/api/data-application-service/frontend/v1/services/{entries['code']}",
                        headers=headers, verify=False,
                    ).json()["service_apply"]["audit_status"]
                    services.append(
                        {
                            "中文名": entries["raw_title"],
                            "描述": entries["raw_description"],
                            "输入参数描述": get_interface_parameter_info(entries['code'], headers, base_url),
                            "是否具有接口的调用权限": "具有" if audit_status == "pass" else "不具有"
                        }
                    )
                elif entries["type"] == "data_catalog":
                    database.append(
                        {
                            "中文名": entries["raw_title"],
                            "描述": entries["raw_description"],
                            "输入参数描述": get_column_info(entries['id'], headers, base_url)
                        }
                    )
        logger.debug("召回的结果是：{}".format({"数据目录": database, "接口服务": services}))
        return {"数据目录": database, "接口服务": services}
    except KeyError:
        logger.debug("召回的结果异常......")
        return "用户登录已过期, 请重新登录或者跟换 Authorization。"


def get_column_info(idx: str, headers: dict, base_url: str) -> dict:
    column_info = requests.request(
        method="GET", headers=headers, verify=False,
        url=f"{base_url}/api/data-catalog/frontend/v1/data-catalog/{idx}/column"
    )
    en2cn = {
        entries["name_en"]: entries["name_cn"] + entries["ai_description"]
        for entries in column_info.json()["entries"]
    }
    return en2cn


def get_interface_parameter_info(idx: str, headers: dict, base_url: str) -> dict:
    service_info = requests.request(
        method="GET", headers=headers, verify=False,
        url=f"{base_url}/api/data-application-service/frontend/v1/services/{idx}"
    )
    en2cn = {
        entries["en_name"]: entries["cn_name"] + entries["description"]
        for entries in service_info.json()["service_param"]["data_table_request_params"]
    }
    return en2cn

