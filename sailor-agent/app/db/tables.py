import os
import yaml

from sqlalchemy.pool import NullPool
from sqlalchemy import Column, String, create_engine, Integer, Text, DateTime, FLOAT, DECIMAL
from sqlalchemy.orm import declarative_base
from config import get_settings
from sqlalchemy.dialects.mysql import LONGTEXT, TEXT, DOUBLE

"""
    綠氨 数据库 表
"""

Base = declarative_base()


class AgentChatHistory(Base):
    __tablename__ = 'agent_chat_history'
    # __table_args__ = {
    #     'mysql_charset': 'utf8'
    # }
    chat_history_id = Column(Integer, primary_key=True, nullable=False, comment="历史id,雪花")
    id = Column(String, nullable=False, comment="uuid")
    title = Column(String, nullable=False, comment="问答标题")
    session_id = Column(String, nullable=False, comment="session_id")
    product_id = Column(String, nullable=False, comment="product_id")
    created_by_uid = Column(String, nullable=False, comment="创建者id")
    created_at = Column(DateTime, nullable=False, comment="创建时间")
    updated_at = Column(DateTime, nullable=False, comment="更新时间")
    deleted_at = Column(Integer, nullable=False, comment="删除时间")


class AgentChatHistoryDetails(Base):
    __tablename__ = 'agent_chat_history_details'
    # __table_args__ = {
    #     'mysql_charset': 'utf8'
    # }
    chat_history_details_id = Column(Integer, primary_key=True, nullable=False, comment="历史id,雪花")
    id = Column(String, nullable=False, comment="uuid")
    session_id = Column(String, nullable=False, comment="session_id")
    query = Column(String, nullable=False, comment="用户问题")
    answer = Column(TEXT, nullable=False, comment="机器人答案")
    like_status = Column(String, nullable=False, comment="点赞状态")
    created_at = Column(DateTime, nullable=False, comment="创建时间")
    updated_at = Column(DateTime, nullable=False, comment="更新时间")
    deleted_at = Column(Integer, nullable=False, comment="删除时间")


class TAgent(Base):
    __tablename__ = 't_agent'
    agent_id = Column(Integer, nullable=False, comment="af 智能体id")
    id = Column(String, primary_key=True, nullable=False, comment="智能体id")
    adp_agent_key = Column(String, nullable=False, comment="adp 智能体key")
    deleted_at = Column(Integer, nullable=False, comment="删除时间")
    created_at = Column(DateTime, nullable=False)
    updated_at = Column(DateTime, nullable=False)

def get_engine():
    from urllib import parse
    settings = get_settings()
    user = settings.DPQA_MYSQL_USER
    passwd = settings.DPQA_MYSQL_PASSWORD
    host = settings.DPQA_MYSQL_HOST
    # port = config["mysql"]["port"]
    database = settings.DPQA_MYSQL_DATABASE
    dbtype = settings.DB_TYPE
    if dbtype.upper() == "DM8":
        sqlalchemy_database_uri = 'dm+dmPython://{user}:{passwd}@{host}?schema={dbschema}'.format(
            user=user,
            passwd=passwd,  # 特殊字符@处理
            host=host,
            dbschema=database
        )
    else:

        sqlalchemy_database_uri = 'mysql+pymysql://{user}:{passwd}@{host}/{database}?charset=utf8'.format(
            user=user,
            passwd=passwd,  # 特殊字符@处理
            host=host,
            database=database
        )
    _engine = create_engine(sqlalchemy_database_uri,
                            # connect_args={'connect_timeout': 100},
                            poolclass=NullPool,
                            echo=True)
    # _engine = create_engine(url, **kwargs)

    return _engine