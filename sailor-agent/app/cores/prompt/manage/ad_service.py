import base64
import json
import os
import traceback
from typing import Any, Dict, List
from urllib.parse import urljoin

from Crypto.Cipher import PKCS1_v1_5
from Crypto.PublicKey import RSA

from app.cores.prompt.manage.api_base import API
from app.cores.prompt.manage.prompt_to_anydata import *
from app.logs.logger import logger
from config import settings
from app.cores.data_mart_qa.models import DataMartQAModel

class PromptServices(object):
    ad_gateway_url: str = settings.AD_GATEWAY_URL
    data_catalog_url: str = "http://data-catalog:8153/"

    def __init__(self):
        self._gen_api_url()
        self._file_name()
        self.payload = payload
        self.path = os.path.abspath(
            os.path.join(os.path.abspath(__file__),
                         os.pardir, "prompt2id.json")
        )

    def _gen_api_url(self):
        self.get_appid_url = "/api/rbac/v1/user/login/appId"
        self.get_prompt_url = "/api/model-factory/v1/prompt/{prompt_id}"
        self.save_prompt_url = "/api/model-factory/v1/prompt/batch_add"
        self.ad_graph_search = "/api/engine/v1/custom-search/kgs/{kg_id}"
        self.get_prompt_item_url = "/api/model-factory/v1/prompt-item-source"
        self.get_prompt_id_url = "/api/model-factory/v1/prompt-source"
        self.search_id_url1 = "/api/data-catalog/frontend/v1/data-resources/search"
        self.search_id_url2 = "/api/data-catalog/frontend/v1/data-resources/searchForOper"
        self.search_id_url3 = "/api/data-catalog/frontend/v1/data-catalog/search"
        self.search_id_url4 = "/api/data-catalog/frontend/v1/data-catalog/operation/search"

    def _file_name(self):
        af_ip = settings.AF_IP
        self.file_name = "af_sailor_agent"
        if af_ip != "":
            self.file_name += "_" + af_ip

    @staticmethod
    def encrypt():
        password = f"{settings.AD_GATEWAY_PASSWORD}"
        pub_key = RSA.importKey(base64.b64decode(pub_key_ad))
        rsa = PKCS1_v1_5.new(pub_key)
        password = rsa.encrypt(password.encode("utf-8"))
        password_base64 = base64.b64encode(password).decode()
        return password_base64

    def get_appid(self):
        try:
            url = urljoin(self.ad_gateway_url, self.get_appid_url)
            api = API(
                url=url,
                payload={
                    "username": settings.AD_GATEWAY_USER,
                    "password": self.encrypt(),
                    "isRefresh": 0,
                },
                method="POST"
            )
            res = api.call()
            return res["res"]
        except Exception as e:
            print(e)
            print(
                f"获取appid错误, 账号：{settings.AD_GATEWAY_USER}， 密码：{settings.AD_GATEWAY_PASSWORD}")

    async def get_prompt(self, appid: str, prompt_id: str):
        try:
            url = urljoin(self.ad_gateway_url,
                          self.get_prompt_url.format(prompt_id=prompt_id))
            api = API(
                url=url,
                headers={"appid": appid},
            )
            res = await api.call_async()
            prompt = res["res"]["messages"]
            return prompt
        except Exception as e:
            print(e)
            return None

    def save_prompt_to_anydata(self):
        url = urljoin(self.ad_gateway_url, self.save_prompt_url)
        for items in self.payload:
            items["prompt_item_type_name"] = self.file_name

        api = API(
            url=url,
            payload=self.payload,
            headers={"appid": self.get_appid()},
            method="POST"
        )
        res = api.call()
        map_prompt = {
            key: value
            for key, value in res["res"][0]["prompt_list"].items()
        }
        map_prompt["file_name"] = self.file_name
        with open(self.path, 'w') as json_file:
            json.dump(map_prompt, json_file, indent=4, ensure_ascii=False)
        logger.info(
            msg=f"prompt成功写入：{self.ad_gateway_url}: AnyFabric, {self.file_name}")

        return map_prompt, self.ad_gateway_url

    @staticmethod
    def from_local(name):
        """get the prompt from local file

        :param name: prompt name
        :return: the prompt content
        """

        logger.info("get anydata prompt error: used local prompt")
        return prompt_map[name]

    async def from_anydata(self, appid: str, name: str) -> tuple:
        """get the prompt that has already been saved on anydata

        :param appid: anydata appid
        :param name: prompt name
        :return: the prompt content
        """
        try:
            logger.info(f"获取prompt的路径：{self.path}")
            with open(self.path, 'r') as file:
                data = json.load(file)
            prompt_id = data.get(name, None)
            print("当前prompt 分组名字：{}".format(self.file_name))
            assert prompt_id is not None, f"get {name} prompt error: {os.path.abspath(__file__)}"
            prompt = await self.get_prompt(prompt_id=prompt_id, appid=appid)
            if prompt is None:
                prompt = self.from_local(name)
            print(prompt)
            return prompt, prompt_id
        except Exception as e:
            logger.info(e)

    async def get_prompt_id(self, name: str) -> str:
        with open(self.path, 'r') as file:
            data = json.load(file)
        prompt_id = data.get(name, None)
        assert prompt_id is not None, f"get {name} prompt error: {os.path.abspath(__file__)}"

        return prompt_id

    def update_prompt_id(self) -> None:
        try:
            url = urljoin(self.ad_gateway_url, self.get_prompt_item_url)
            api = API(
                url=url,
                headers={"appid": self.get_appid()},
                params={
                    "is_management": "true",
                    "size": 10000,
                    "prompt_item_name": "AnyFabric",
                }
            )

            res = api.call()["res"]["data"]
            for item in res:
                if item["prompt_item_name"] == "AnyFabric":
                    res = item
                    break
            prompt_item_id = res["prompt_item_id"]
            for items in res["prompt_item_types"]:
                if items["name"] == self.file_name:
                    prompt_item_type_id = items["id"]
                    break
            url = urljoin(self.ad_gateway_url, self.get_prompt_id_url)
            api = API(
                url=url,
                headers={"appid": self.get_appid()},
                params={
                    "prompt_item_id": prompt_item_id,
                    "prompt_item_type_id": prompt_item_type_id,
                    "is_management": "true",
                    "rule": "create_time",
                    "order": "desc",
                    "deploy": "all",
                    "prompt_type": "all",
                    "size": 10000,
                    "page": 1,
                }
            )
            map_prompt = {}
            res = api.call()["res"]["data"]
            for msg in res:
                map_prompt[msg["prompt_name"]] = msg["prompt_id"]
            map_prompt["file_name"] = self.file_name
            with open(self.path, 'w') as json_file:
                json.dump(map_prompt, json_file, indent=4, ensure_ascii=False)
            logger.info("prompt id has been updated")
            logger.info(json.dumps(map_prompt, indent=4, ensure_ascii=False))
        except Exception as e:
            logger.info("提示词 id 更新错误")
            print(e)

    def get_all_prompt_item(self) -> tuple[list[Any], bool]:
        url = urljoin(self.ad_gateway_url, self.get_prompt_item_url)
        api = API(
            url=url,
            headers={"appid": self.get_appid()},
            params={
                "is_management": "true",
                "size": 10000,
            }
        )
        res = api.call()
        msg = [name["prompt_item_types"] for name in res["res"]
               ["data"] if name.get("prompt_item_name") == "AnyFabric"]
        name = []
        for x in msg:
            for y in x:
                name.append(y["name"])

        if self.file_name in name:
            return name, False

        return name, True

    async def fun_verify_online(self, headers: dict, ids: list, search_id_url: str) -> dict[Any, list[Any]]:
        # 查验资源是否可用是否为线上
        # 1是资源版非数据运营开发，2是资源版数据运营开发  ，3是目录版非数据运营开发，  4是目录版数据运行开发
        service_id = {}
        params = {"filter": {"ids": ids}}
        url = urljoin(
            self.data_catalog_url,
            search_id_url)
        api1 = API(
            url=url,
            headers=headers,
            method="POST",
            params=params
        )
        try:
            res = await api1.call_async()
            if res["total_count"] > 0:
                for i in res["entries"]:
                    if i["type"] in service_id.keys():
                        service_id[i["type"]].append(i['id'])
                    else:
                        service_id[i["type"]] = []
        except Exception as e:
            print('查询资源上下线内部接口过程报错，报错如下', e)
        return service_id

    async def check_resource_status(self, params: DataMartQAModel) -> dict[Any, list[Any]]:
        try:
            # 查验资源是否可用是否为线上
            # 1是资源版非数据运营开发，2是资源版数据运营开发  ，3是目录版非数据运营开发，  4是目录版数据运行开发
            headers = {"Authorization": params.token}
            ids = [source.id for source in params.resources]
            roles = params.roles
            if "data-operation-engineer" in roles or "data-development-engineer" in roles:
                if params.af_editions == "catalog":
                    res = await self.fun_verify_online(headers, ids, self.search_id_url4)
                else: 
                    res = await self.fun_verify_online(headers, ids, self.search_id_url2)
            else:
                if params.af_editions == "catalog":
                    res = await self.fun_verify_online(headers, ids, self.search_id_url3)
                else: 
                    res = await self.fun_verify_online(headers, ids, self.search_id_url1)
                
            
        except Exception as e:
            logger.info(traceback.format_exc())
            res = {}
            print('查询资源上下线外部接口过程报错，报错如下', e)
        return res


prompt_service = PromptServices()


