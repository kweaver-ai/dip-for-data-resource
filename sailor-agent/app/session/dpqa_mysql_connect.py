# -*- coding:utf-8 -*-
import dmPython
import pymysql
from config import settings
from app.logs.logger import logger
from app.cores.data_product_qa.dpqa_error import dpqa_error_code

class MariaDBHandler:
    def __init__(self):
        self.host = settings.DPQA_MYSQL_HOST.split(":")[0]
        self.port = settings.DPQA_MYSQL_HOST.split(":")[1]
        self.user = settings.DPQA_MYSQL_USER
        self.password = settings.DPQA_MYSQL_PASSWORD
        # self.database = settings.DPQA_MYSQL_DATABASE
        self.dbType = settings.DB_TYPE
        self.database = "af_main"
        self.connection = None
        self.cursor = None

    async def connect(self):
        try:
            if self.dbType.upper()=="DM8":
                self.connection = dmPython.connect(
                    user=self.user,
                    password=self.password,
                    server=self.host,
                    port=int(self.port),  # 达梦默认端口
                    schema=self.database,
                    autoCommit=True
                )
                self.cursor = self.connection.cursor()
                print("DM8 Successfully connected to the database")
            else:
                # 建立数据库连接
                self.connection = pymysql.connect(
                    host=self.host,
                    user=self.user,
                    password=self.password,
                    database=self.database,
                    port=int(self.port)
                )
                print("Successfully connected to the database")
                self.cursor = self.connection.cursor(pymysql.cursors.DictCursor)
        except pymysql.MySQLError as e:
            logger.info(f"Error while connecting to MariaDB: {e}")
            return dpqa_error_code.dbcon_error_json
        except (dmPython.Error, Exception) as err:
            logger.info(f"Error while connecting to DM8: {err}")
            return dpqa_error_code.dbcon_error_json

    async def disconnect(self):
        if self.cursor:
            self.cursor.close()
        if self.connection:
            self.connection.close()
        print("MariaDB connection is closed")

    async def query_data_by_id(self, id_value):
        try:
            sql_query = "SELECT * FROM agent WHERE agent_id = %s"
            self.cursor.execute(sql_query, (id_value,))
            result = self.cursor.fetchone()  # 获取单条记录
            if result:
                # print(f"Data for ID {id_value}: {result}")
                return result
            else:
                # print(f"No data found for ID {id_value}")
                return dpqa_error_code.nodata_error_json
        except pymysql.MySQLError as e:
            logger.info(f"Error while querying data: {e}")
            return dpqa_error_code.query_error_json

    async def query_special_field_info(self):
        try:
            sql_query = "SELECT * FROM agent_fields"
            self.cursor.execute(sql_query)
            result = self.cursor.fetchone()  # 获取单条记录
            if result:
                # print(f"Data for ID {id_value}: {result}")
                return json.loads(result["fields_info"])
            else:
                # print(f"No data found for ID {id_value}")
                return dict()
        except pymysql.MySQLError as e:
            logger.info(f"Error while querying data: {e}")
            return dict()


if __name__ == "__main__":
    # 创建MariaDBHandler实例
    db_handler = MariaDBHandler()
    # 连接数据库
    qs = db_handler.connect()
    if qs:
        print(qs)
    else:
        id_to_query = "d285a611-adcf-4cb5-8997-5775168d5449"
        result = db_handler.query_data_by_id(id_to_query)
        print(result)
        print(type(result))
        print("="*30)
        import json
        result = json.loads(result["config"])
        tools = result["tools"]
        personality = result["configs"]["task_desc"]
        background = result["configs"]["knowledge"]
        data_views_list = result["data_views"]
        print(tools)
        print("+"*30)
        print(personality)
        print("+" * 30)
        print(background)
        print("+" * 30)
        print(data_views_list)
        print("=" * 30)
        db_handler.disconnect()
