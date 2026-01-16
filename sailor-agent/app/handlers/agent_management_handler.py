from fastapi import APIRouter, Request, Body, Depends

from app.cores.af_agents.agent_service import AFAgentService
from app.models.agent_models import AFAgentListReqBody, PutOnAFAgentReqBody, PullOffAFAgentReqBody
from app.utils.get_token import get_token

AgentManagementRouter = APIRouter()
agent_service = AFAgentService()


@AgentManagementRouter.post("/agent/list")
async def get_agent_list(request: Request, req_body: AFAgentListReqBody = Body(...), token: str = Depends(get_token)):
    """获取智能体列表"""
    try:
        resp = agent_service.get_agent_list(req_body, token)
        return resp
    except Exception as e:
        return {"entries": [], "pagination_marker_str": "", "is_last_page": True}


@AgentManagementRouter.put("/agent/put-on")
async def put_on_agent(request: Request, req_body: PutOnAFAgentReqBody = Body(...), token: str = Depends(get_token)):
    """上架智能体"""
    try:
        resp = agent_service.put_on_agent(req_body)
        return resp
    except Exception as e:
        return {"res": {"status": "error", "message": str(e)}}


@AgentManagementRouter.put("/agent/pull-off")
async def pull_off_agent(request: Request, req_body: PullOffAFAgentReqBody = Body(...), token: str = Depends(get_token)):
    """下架智能体"""
    try:
        resp = agent_service.pull_off_agent(req_body)
        return resp
    except Exception as e:
        return {"res": {"status": "error", "message": str(e)}}
