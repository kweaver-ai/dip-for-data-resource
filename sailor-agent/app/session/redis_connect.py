import redis
from redis.sentinel import Sentinel

from config import settings


class RedisConnect:
    def __init__(self):
        self.redis_cluster_mode = settings.REDIS_CONNECT_TYPE
        self.db = settings.REDIS_DB
        self.master_name = settings.REDIS_MASTER_NAME
        self.sentinel_user_name = settings.REDIS_SENTINEL_USER_NAME

        self.host = settings.REDIS_HOST
        self.sentinel_host = settings.REDIS_SENTINEL_HOST

        self.port = settings.REDIS_PORT
        self.sentinel_port = settings.REDIS_SENTINEL_PORT

        self.password = settings.REDIS_PASSWORD
        self.sentinel_password = settings.REDIS_SENTINEL_PASSWORD

    def connect(self):
        if self.redis_cluster_mode == "master-slave":
            pool = redis.ConnectionPool(
                host=self.host,
                port=self.port,
                password=self.password,
                db=self.db,
            )
            client = redis.StrictRedis(connection_pool=pool)
            return client
        if self.redis_cluster_mode == "sentinel":
            sentinel = Sentinel(
                [(self.sentinel_host, self.sentinel_port)],
                password=self.sentinel_password,
                sentinel_kwargs={
                    "password": self.sentinel_password,
                    "username": self.sentinel_user_name
                }
            )
            client = sentinel.master_for(
                self.master_name,
                password=self.sentinel_password,
                username=self.sentinel_user_name,
                db=self.db
            )
            return client


redis_client = RedisConnect()

