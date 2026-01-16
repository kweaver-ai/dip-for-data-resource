import uuid
from datetime import datetime

from app.cores.service.adp_service import ADPService
from app.db.tables import TAgent, get_engine
from sqlalchemy.orm import sessionmaker
from app.logs.logger import logger


class AFAgentService(object):
    LISTSTATUS = 1     # 列表模式

    def __init__(self):
        self.adp_service = ADPService()
        self.engine = get_engine()
        self.Session = sessionmaker(bind=self.engine)

    def get_agent_list(self, req, token):
        """获取智能体列表"""
        try:
            # 构建请求参数
            adp_req = {
                "name": req.name,
                "size": req.size,
                "pagination_marker_str": req.pagination_marker_str
            }

            # 获取已配置的智能体列表
            agent_list = self.get_agent_list_from_db()
            agent_keys = []
            agent_key_2_af_agent = {}
            for item in agent_list:
                agent_keys.append(item.adp_agent_key)
                agent_key_2_af_agent[item.adp_agent_key] = item.id

            # 如果是列表模式，添加智能体keys
            if req.list_flag == self.LISTSTATUS:
                adp_req["agent_keys"] = agent_keys
                adp_req["ids"] = []
                adp_req["exclude_agent_keys"] = []
                adp_req["business_domain_ids"] = []

            # 调用ADP服务获取智能体列表
            adp_resp = self.adp_service.agent_list(adp_req, token)

            # 处理返回值
            entries = adp_resp.get("entries", [])
            for item in entries:
                if item.get("key") in agent_key_2_af_agent:
                    item["af_agent_id"] = agent_key_2_af_agent[item.get("key")]
                    item["list_status"] = "put-on"
                else:
                    item["list_status"] = "pull-off"

            return {
                "entries": entries,
                "pagination_marker_str": adp_resp.get("pagination_marker_str", ""),
                "is_last_page": adp_resp.get("is_last_page", True)
            }
        except Exception as e:
            logger.error(f"Get agent list failed: {str(e)}")
            return {
                "entries": [],
                "pagination_marker_str": "",
                "is_last_page": True
            }

    def put_on_agent(self, req):
        """上架智能体"""
        try:
            session = self.Session()
            try:
                for item in req.agent_list:
                    # 检查是否已存在
                    existing_agent = session.query(TAgent).filter(
                        TAgent.adp_agent_key == item.agent_key,
                        TAgent.deleted_at == 0
                    ).first()

                    if not existing_agent:
                        # 创建新智能体
                        new_agent = TAgent(
                            agent_id=self._generate_agent_id(),
                            id=str(uuid.uuid4()),
                            adp_agent_key=item.agent_key,
                            deleted_at=0,
                            created_at=datetime.now(),
                            updated_at=datetime.now()
                        )
                        session.add(new_agent)

                session.commit()
                return {"res": {"status": "success"}}
            finally:
                session.close()
        except Exception as e:
            logger.error(f"Put on agent failed: {str(e)}")
            return {"res": {"status": "error", "message": str(e)}}

    def pull_off_agent(self, req):
        """下架智能体"""
        try:
            session = self.Session()
            try:
                # 软删除智能体
                agent = session.query(TAgent).filter(
                    TAgent.id == req.af_agent_id,
                    TAgent.deleted_at == 0
                ).first()

                if agent:
                    agent.deleted_at = int(datetime.now().timestamp())
                    agent.updated_at = datetime.now()
                    session.commit()

                return {"res": {"status": "success"}}
            finally:
                session.close()
        except Exception as e:
            logger.error(f"Pull off agent failed: {str(e)}")
            return {"res": {"status": "error", "message": str(e)}}

    def get_agent_list_from_db(self):
        """从数据库获取智能体列表"""
        try:
            session = self.Session()
            try:
                agents = session.query(TAgent).filter(TAgent.deleted_at == 0).all()
                return agents
            finally:
                session.close()
        except Exception as e:
            logger.error(f"Get agent list from db failed: {str(e)}")
            return []

    def _generate_agent_id(self):
        """生成智能体ID"""
        # 简单实现，实际项目中可能需要使用雪花算法
        return int(datetime.now().timestamp() * 1000)
