# -*- coding: utf-8 -*-

import os

from langchain_community.chat_models import ChatOpenAI

from config import settings


def create_local_agent_openai_style_chat_model(
    model: str = "",
    openai_api_base: str = ""
):
    openai_api_base, model_name = os.path.split(settings.AISHU_READER_LLM.rstrip("/"))

    llm = ChatOpenAI(
        model_name=model_name,
        openai_api_key="EMPTY",
        openai_api_base=openai_api_base,
        max_tokens=2000,
        temperature=0.01,
    )

    return llm


agent_openai_style_chat_model = create_local_agent_openai_style_chat_model()
small_model = create_local_agent_openai_style_chat_model()

