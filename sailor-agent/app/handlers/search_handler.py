# -*- coding:utf-8 -*-
import sys

from fastapi import APIRouter, Request, Body, Path
from fastapi.responses import JSONResponse

from app.cores.prompt.manage.api_base import API
from app.routers.agent_temp_router import *

from config import settings

# sys.path.append("/mnt/pan/zkn/code_agent/feature_633396")

search_api = APIRouter()


@search_api.get(SearchInfoRouter)
async def search_info(request: Request):
    if not verify_token(request):  # 验证token
        return JSONResponse(status_code=401, content={"message": "Unauthorized"})


    return JSONResponse({"res": {
        "adp_agent_key": settings.ADP_AGENT_KEY,
        "adp_business_domain_id": settings.ADP_BUSINESS_DOMAIN_ID
    }}, status_code=200)

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


