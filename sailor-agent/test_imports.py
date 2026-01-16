try:
    from app.cores.af_agents.agent_service import AFAgentService
    from app.models.agent_models import AFAgentListReqBody, PutOnAFAgentReqBody, PullOffAFAgentReqBody
    from app.utils.get_token import get_token
    print("All imports successful!")
except Exception as e:
    print(f"Import error: {e}")
    import traceback
    traceback.print_exc()
