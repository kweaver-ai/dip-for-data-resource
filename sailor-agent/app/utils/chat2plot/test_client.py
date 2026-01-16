# -*- coding: utf-8 -*-
import pandas as pd
from app.utils.chat2plot import chat2plot
from langchain_community.chat_models import ChatOpenAI
import json
from app.utils.chat2plot.g2_schema import G2PlotConfig

csv_path = "/mnt/pan/douziyu/测试.csv"
llm = ChatOpenAI(
        model_name="Qwen2-72B-Chat",
        openai_api_key="EMPTY",
        openai_api_base="http://192.168.152.11:8303/v1",
        max_tokens=20000,
        temperature=0
    )
df = pd.read_csv(csv_path,encoding="utf-8")
c2p = chat2plot(df,language="中文",schema_definition=G2PlotConfig, chat=llm)
res = c2p("展示一下大家各自的分数")
config = res.config.dict()
explanation = res.explanation
print(config)
print(explanation)
config = json.dumps(config,ensure_ascii=False)
datas = df.to_json(orient="records",force_ascii=False)
final_result = json.dumps({"data":datas,"config":config},ensure_ascii=False)
print(final_result)
