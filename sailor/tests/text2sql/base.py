from abc import ABC, abstractmethod
from typing import Any


class TestBase(ABC):

    @abstractmethod
    def get_metadata(
        self
    ) -> tuple[Any, Any]:
        """
        获取表基本信息和样例信息
        """

    @abstractmethod
    def build_prompt(
        self,
        query: str,
        ample: Any,
        ddl: Any,
        error_code: Any
    ) -> str:
        """
        构建基于text2sql的prompt
        """

    @abstractmethod
    def use_llm(
        self,
        prompt: str
    ) -> str:
        """
        调用大模型去获取prompt的执行结果
        """

    @abstractmethod
    def query_database(
        self
    ) -> str:
        """
        查询数据库
        """

    @abstractmethod
    def check_sql(
        self,
        sql: str,
        tables: list
    ) -> str:
        """
        校验SQL
        """

    @abstractmethod
    def accuracy_rate(
        self,
        predicted: Any,
        label: Any
    ) -> float:
        """
        计算准确率
        """
