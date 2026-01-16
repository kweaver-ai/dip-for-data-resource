import unittest
import asyncio
from app.utils.password import get_authorization
from app.cores.cognitive_assistant.qa import QA
from app.handlers.cognitive_assistant_handler import QAParamsModel
from app.utils.tp201 import inputs

class NestedDict(dict):
    def __getattr__(self, item):
        return self.get(item)


req = NestedDict({"headers": NestedDict(
    {
        "Authorization": get_authorization('https://10.4.109.201', "bank")
    }
)})


class TestQA(unittest.TestCase):
    @staticmethod
    async def sql():
        inputs["query"] = "学生违纪情况"
        chunk = ""
        qa = QA()
        resp = qa.stream(req, QAParamsModel(**inputs))
        async for chunk in resp:
            continue
        assert 'data: {"result": {"status": "ending"}}' in chunk

    @staticmethod
    async def svc():
        inputs["query"] = "根据学生维表id和课程维表id查询成绩,fk_course=6232e801, fk_student=5634abb6"
        chunk = ""
        qa = QA()
        resp = qa.stream(req, QAParamsModel(**inputs))
        async for chunk in resp:
            continue
        assert 'data: {"result": {"status": "ending"}}' in chunk

    def test_sql(self):
        loop = asyncio.get_event_loop()
        loop.run_until_complete(self.sql())

    def test_svc(self):
        loop = asyncio.get_event_loop()
        loop.run_until_complete(self.svc())


if __name__ == '__main__':
    unittest.main()
