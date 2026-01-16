# -*- coding: utf-8 -*-
"""
@Time ： 2024/7/24 10:17
@Auth ： Xia.wang
@File ：build_test_report.py
@IDE ：PyCharm
"""
import json
from sql_similarity_checker import cosine_similarity_text
import re
import pymysql
from mo_sql_parsing import parse
import unicodedata
from tests.text2sql.dataset import save_to_excel, read_jsonl


class ResultDataBuilder:
    def __init__(self, result_set_filepath):
        self.data = []
        # self.review_dataset_filepath = review_dataset_filepath
        self.result_set_filepath = result_set_filepath
        self.dataframe = None

    def add_sql_data(self, tables_schema, question, complexity, sql_difficulty_level, sql_type,
                     golden_sql, golden_sql_result, pred_sql, pred_sql_result, sql_similarity,
                     is_match_sql, is_match_sql_result, sql_difference, is_code_table=0, code_table=None):
        entry = {
            "complexity": complexity,
            "tables_schema": tables_schema,
            "question": question,
            "sql_difficulty_level": sql_difficulty_level,
            "sql_type": sql_type,
            "golden_sql": golden_sql,
            "golden_sql_result": golden_sql_result,
            "pred_sql": pred_sql,
            "pred_sql_result": pred_sql_result,
            "sql_similarity": sql_similarity,
            "sql_is_match": is_match_sql,
            "sql_result_is_match": is_match_sql_result,
            "sql_difference": sql_difference
        }
        if is_code_table == 1:
            entry.update({"code_table": code_table})
        self.data.append(entry)

    def write_to_json(self, filename):
        print(f'filename : {filename}')
        with open(filename, 'w', encoding='utf-8') as f:
            json.dump(self.data, f, ensure_ascii=False, indent=4)

    def metadata(self):
        # review_dataset = read_jsonl(self.review_dataset_filepath)
        result_set = read_jsonl(self.result_set_filepath)
        return result_set

    def is_all_match(self, golden, pred):
        if golden == pred:
            return 1
        else:
            return 0

    def _get_connection(self):
        self.config = {
            'host': '10.4.109.228',
            'user': 'root',
            'password': '',
            'database': 'sakila',
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

    def sql_difficulty_level(self, sql_keyword_count):
        level = ["简单", "基础", "中等", "难", "非常难"]
        if not sql_keyword_count:
            return level[0]
        elif sql_keyword_count.get('JOIN', 0) > 3:
            return level[4]
        elif sql_keyword_count.get('JOIN', 0) == 2:
            return level[3]
        elif sql_keyword_count.get('JOIN', 0) < 2:
            return level[1]
        elif 'WHERE' in sql_keyword_count:
            return level[1]
        else:
            return "人工审核"

    def sql_type(self, golden_sql):
        SQL_KEYWORD = ['JOIN', 'GROUP BY', 'ORDER BY', 'LIMIT', 'COUNT', 'LIKE', 'WHERE', 'LEFT JOIN', 'RIGHT JOIN',
                       'IN',
                       'NOT IN', 'BETWEEN', 'AND', 'OR', 'NOT', 'SUM', 'AVG', 'MAX', 'MIN', 'HAVING', 'INNER JOIN',
                       'FULL JOIN']
        sql_keyword_count = dict.fromkeys(SQL_KEYWORD, 0)
        for keyword in sql_keyword_count:
            sql_keyword_count[keyword] = golden_sql.count(keyword)
        sql_keyword_count = {keyword: count for keyword, count in sql_keyword_count.items() if count > 0}
        return sql_keyword_count

    def query_database(
            self, sql):
        connection = self._get_connection()
        with connection.cursor() as cursor:
            try:
                cursor.execute(sql)
                result = cursor.fetchall()
                if len(result) == 1:
                    result = list(result[0])
                elif len(result) == 0:
                    result = []
                else:
                    result = [list(row) for row in result]
                converted_result = self.check_sql_resutl(result)
            except Exception as e:
                result = str(e)

        connection.close()
        return converted_result

    def sql_parsing(self, golden_sql, pred_sql):
        try:
            golden_sql_parse = parse(golden_sql)
            pred_sql_parse = parse(pred_sql)
            return self.dict_difference(golden_sql_parse, pred_sql_parse)
        except Exception:
            print(f"对比差异报错: {golden_sql, pred_sql}")
            pass

    def dict_difference(self, dict1, dict2):
        diff = {}
        for key in set(dict1.keys()).union(dict2.keys()):
            if key in dict1 and key in dict2:
                if isinstance(dict1[key], dict) and isinstance(dict2[key], dict):
                    nested_diff = self.dict_difference(dict1[key], dict2[key])
                    if nested_diff:
                        diff[key] = nested_diff
                elif isinstance(dict1[key], list) and isinstance(dict2[key], list):
                    if dict1[key] != dict2[key]:
                        diff[key] = (dict1[key], dict2[key])
                elif dict1[key] != dict2[key]:
                    diff[key] = (dict1[key], dict2[key])
            elif key in dict1:
                diff[key] = (dict1[key], None)
            else:
                diff[key] = (None, dict2[key])
        return diff

    def invoke(self):
        result_set = self.metadata()
        for case in result_set:
            complexity = case['complexity']
            tables_schema = case['tables_schema']
            question = case['question']
            golden_sql = case['golden_sql']
            golden_sql_result = case['golden_sql_result']
            pred_sql = case['pred_sql']
            pred_sql_result = case['pred_sql_result']
            sql_type = self.sql_type(golden_sql)
            sql_difficulty_level = self.sql_difficulty_level(sql_type)
            sql_similarity = cosine_similarity_text(golden_sql, pred_sql)
            is_match_sql = self.is_all_match(golden_sql, pred_sql)
            is_match_sql_result = self.is_all_match(golden_sql_result, pred_sql_result)
            sql_difference = self.sql_parsing(golden_sql, pred_sql)
            self.add_sql_data(complexity=complexity, tables_schema=tables_schema, question=question,
                              golden_sql=golden_sql, golden_sql_result=golden_sql_result, pred_sql=pred_sql,
                              pred_sql_result=pred_sql_result, sql_type=sql_type,
                              sql_difficulty_level=sql_difficulty_level, sql_similarity=sql_similarity,
                              is_match_sql=is_match_sql,
                              is_match_sql_result=is_match_sql_result, sql_difference=sql_difference, is_code_table=1,
                              code_table=case['code_table'])


if __name__ == '__main__':
    # database_sample_result_path = "./dataset/final_result_sample_20240813.json"
    # data_explore_path = "./dataset/final_result_codetable_20240813.json"
    # review_dataset_path = "dataset/baseline20240712.jsonl"
    result_set_filepath = "./dataset/code_table_benchmark.imp.jsonl"
    builder = ResultDataBuilder(result_set_filepath)

    builder.invoke()

    filename = './dataset/code_table_benchmark.xlsx'
    save_to_excel(filename, builder.data)
