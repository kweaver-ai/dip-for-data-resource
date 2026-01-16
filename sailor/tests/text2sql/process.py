from typing import Any, Optional
from langchain_openai import ChatOpenAI
from app.cores.prompt.text2sql import GENERATE_SQL_TEMPLATE
from app.cores.text2sql.t2s_api import Services
from app.cores.text2sql.t2s_reshape import Reshape
from tests.text2sql.dataset import read_json, read_csv, write_jsonl, read_jsonl
from tests.text2sql.base import TestBase
import re
import ast
import json
import pymysql
from decimal import Decimal
from datetime import datetime


class TestText2SQL(TestBase):
    template: str = GENERATE_SQL_TEMPLATE
    services: Services = Services()
    parse: Reshape = Reshape()

    def __init__(
            self,
            metadata_path: str,
            sample_path: str
    ):
        super().__init__()
        self.metadata_path = metadata_path
        self.sample_path = sample_path

    def get_metadata(
            self
    ) -> tuple[Any, Any]:
        sample = read_json(self.sample_path)
        metadata = read_csv(self.metadata_path)
        return sample, metadata

    def get_metadata_code_table(
            self
    ) -> Any:
        metadata = read_jsonl(self.metadata_path)
        return metadata

    def build_prompt(
            self,
            query: str,
            ddl_and_sample_prompt: Optional[str] = "",
            error_code: Optional[str] = ""
    ) -> str:
        ddl_and_sample = f"""
           {ddl_and_sample_prompt}
            """
        prompt = self.template.replace(
            "{{ddl_and_sample}}", ddl_and_sample
        ).replace(
            "{{query}}", query
        ).replace(
            "{{error_code}}", error_code
        )

        return prompt

    def use_llm(
            self,
            prompt: str
    ) -> str:
        model = ChatOpenAI(
            model_name="Qwen2-72B-Chat",
            openai_api_key="EMPTY",
            openai_api_base="http://192.168.152.11:8303/v1",
            max_tokens=2000,
            temperature=0.01,
        )
        res = model.invoke(prompt).content
        return res

    def _connection_sakila(self):
        self.config = {
            'host': '10.4.109.228',
            'user': 'root',
            'password': '',
            'database': 'sakila',
            'charset': 'utf8mb4'
        }
        return pymysql.connect(**self.config)

    def _connection_china_biz_registry(self):
        self.config = {
            'host': '10.4.109.226',
            'user': 'root',
            'password': '',
            'database': 'china_biz_registry',
            'charset': 'utf8mb4'
        }
        return pymysql.connect(**self.config)

    def check_sql_resutl(self, results):
        converted_result = []
        for item in results:
            if isinstance(item, Decimal):
                converted_result.append(str(item))
            elif isinstance(item, datetime):
                converted_result.append(item.strftime('%Y-%m-%d %H:%M:%S'))
            else:
                converted_result.append(item)
        return converted_result

    def query_database(
            self,
            sql: str
    ) -> list:

        # connection = self._get_connection()
        connection = self._connection_china_biz_registry()
        try:
            with connection.cursor() as cursor:
                cursor.execute(sql)
                results = cursor.fetchall()
                if len(results) == 1:
                    results = list(results[0])
                elif len(results) == 0:
                    results = []
                else:
                    results = [list(row) for row in results]
                return results

        finally:
            connection.close()

        return sql_results

    @staticmethod
    def build_final_result(query, qa_generated_sql, qa_generated_sql_exec_result, af_generated_sql):
        result = {
            f"{query}": {
                "x": {
                    "sql": qa_generated_sql,
                    "exec_result": ast.literal_eval(qa_generated_sql_exec_result)
                },
                "k": {
                    "sql": af_generated_sql,
                    "exec_result": ""
                }
            }
        }
        return result

    def accuracy_rate(
            self,
            data
    ) -> float:
        match_count = 0
        total_count = len(data.items())
        for question, content in data.items():
            x_exec_result = content['x']['exec_result']
            k_exec_result = content['k']['exec_result']

            # 对比exec_result是否一致
            if x_exec_result == k_exec_result:
                match_count += 1
        return "{:.2f}".format(match_count / total_count)

    @staticmethod
    def extract_table_info(ddl):
        table_name_end = ddl.find('(')
        table_name = ddl[:table_name_end].strip()
        fields = ddl[table_name_end + 1:-1].split(',')
        return table_name, fields

    @staticmethod
    def format_ddl(table_name, fields):
        # 格式化字段
        formatted_fields = []
        for field in fields:
            field = field.strip()
            if 'comment' in field:
                field_parts = field.split('comment')
                field_def = field_parts[0].strip()
                field_comment = field_parts[1].strip()
                formatted_field = f"{field_def} comment {field_comment}\n"
                formatted_fields.append(formatted_field)
            else:
                formatted_fields.append(f"{field}\n")

        # 组合表名和格式化后的字段
        formatted_sql = f"{table_name}\n(\n" + "".join(formatted_fields)[:-1] + "\n);"
        return formatted_sql

    @staticmethod
    def build_ddl_sample_prompt(tables_schema, code_table):
        #  flag 1 表示数据库样例数据  2 码表等探查结果信息
        prompt = """已知的表信息的DDL语句，其comment中包含了表的中文业务定义和表字段的中文业务定义：
                """
        index = 0
        for table_name, table_ddl in tables_schema.items():
            ddl_table_name, fields = TestText2SQL.extract_table_info(table_ddl)
            formatted_ddl = TestText2SQL.format_ddl(ddl_table_name, fields)
            # sample_list = [sample[table_name]]
            # sample_str = json.dumps(sample_list, indent=2)
            # prompt += f"""这是第{index + 1}张表：表名：{table_name}，其DDL语句如下：
            #             {formatted_ddl}
            #
            #             这是第{index + 1}张：表名：{table_name}，其样例数据如下：
            #             {code_table}
            #             """
            index += 1
            prompt += f"""这是第{index}张表：表名：{table_name}，其DDL语句如下：
                           {formatted_ddl}

                          这是第{index}张：表名：{table_name}，其码表如下：
                          {code_table[table_name]}
                          """

        prompt += f"""
        
        另外需要注意：
        1. 你只能从我提供给你的表名列表 {list(tables_schema.keys())} 里去生成合适的SQL，不要使用未提供的表名
        2. 码表内的 comment 是码表值，请利用码表值辅助生成合适的SQL
        """
        return prompt

    @staticmethod
    def default(o):
        import decimal
        if isinstance(o, decimal.Decimal):
            return float(o)

    @staticmethod
    def read_write_json_file(filename, flag=0, content=None):
        if flag == 0:
            with open(filename, 'r', encoding='utf-8') as f:
                data = json.load(f)
                return data
        elif flag == 1:
            with open(filename, 'w', encoding='utf-8') as f:
                json.dump(content, f, ensure_ascii=False, default=TestText2SQL.default, indent=4)
                print('打印完成')

    @staticmethod
    def extract_table_name(
            ddl: Any
    ) -> list:
        pattern = r"CREATE TABLE\s+(\w+)\s+\("
        table_names = re.findall(pattern, ddl, re.DOTALL)
        return table_names

    async def check_sql(
            self,
            sql: str,
            tables: list
    ) -> str:
        sql = await self.parse.check_sql_v3(
            sql=sql,
            detail=("", "", tables),
            config={"sql_limit": 3}
        )
        return sql

    async def invoke(
            self,
    ):
        case_infos = self.get_metadata_code_table()  # 返回数据探查结果样例数json文件、数据pandas csv文件对象
        for index, case in enumerate(case_infos):
            question = case['question']
            tables_schema = case['tables_schema']
            golden_sql = case['golden_sql']
            golden_sql_result = case['golden_sql_result']
            code_table = case['code_table']
            tables_ddl = list(tables_schema.values())

            ddl_sample_prompt = self.build_ddl_sample_prompt(tables_schema, code_table)

            prompt = self.build_prompt(
                query=question,
                ddl_and_sample_prompt=ddl_sample_prompt
            )
            # print(f'prompt : {prompt}')
            sql = self.use_llm(prompt)

            case['pred_sql'] = await self.check_sql(sql, tables_ddl)
            case['pred_sql_result'] =  self.query_database(case['pred_sql'])


        return case_infos


if __name__ == '__main__':
    async def main():
        # sample_path = "./dataset/database_sample.json"
        # sample_path = "./dataset/data_explore.json"
        metadata_path = "dataset/code_table_benchmark.jsonl"
        demo = TestText2SQL(metadata_path, '')
        results = await demo.invoke()

        write_jsonl('dataset/code_table_benchmark.imp.jsonl', results)
        # print(f"+++++++++++++++++\n{results}")
        # 将初步执行结果写入JSON文件
        # filename = "./dataset/final_result_sample_20240813.json"
        # filename = "dataset/final_result_codetable_20240813.json"
        # demo.read_write_json_file(filename, flag=1, content=results)
        # print("+++++++++++++++++")
        # 读取json文件，查询数据库执行sql并将结果写入json新文件里
        # data = demo.read_write_json_file(filename)
        # exec_llm_gen_sql_results = demo.query_database(data)
        # filename = "./dataset/final_result_sample_20240813.json"
        # filename = "./dataset/final_result_now_20240812.json"
        # demo.read_write_json_file(filename, flag=1, content=exec_llm_gen_sql_results)
        # print("+++++++++++++++++")
        # 计算SQL执行结果一致率
        # filename = "./dataset/final_result_ago_audit.json"
        # filename = "./dataset/final_result_now_audit.json"
        # data = demo.read_write_json_file(filename)
        # print(f' {demo.accuracy_rate(data)}')


    import asyncio

    asyncio.run(main())
