# -*- coding:utf-8 -*-
import sys

from fastapi import APIRouter, Request, Body, Path
from fastapi.responses import JSONResponse
from starlette.responses import StreamingResponse
from app.cores.af_agents.models import AgentQAModel
from app.cores.af_agents.service import AgentService

from app.cores.prompt.manage.api_base import API
from app.routers.agent_temp_router import *

from config import settings

dpqa_v2 = APIRouter()


"""
转发adp 问答接口
"""

@dpqa_v2.post(AgentV2Router)
async def dpqa_chat_v2(request: Request, adp_agent_key: str = Path(...,  description="adp agent id"), params=Body(...)):
    if not verify_token(request):  # 验证token
        return JSONResponse(status_code=401, content={"message": "Unauthorized"})

    authorization = request.headers.get('Authorization')
    agent_server = AgentService()

    params["agent_key"] = adp_agent_key

    new_params = AgentQAModel.model_validate(params)


    return StreamingResponse(agent_server.stream(new_params, authorization), media_type="text/event-stream")


def verify_token(request: Request = None):
    token: str = request.headers.get("Authorization").split(" ")[1]
    api = API(
        data={"token": token},
        method="POST",
        url="{}/admin/oauth2/introspect".format(settings.HYDRA_URL),
        headers={
            "Content-Type": "application/x-www-form-urlencoded"
        }
    )
    res = api.call()
    return res.get("active", False)


