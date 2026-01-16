import os
from langchain_community.chat_models import ChatOpenAI
from data_retrieval.utils.llm import CustomChatOpenAI
from app.logs.logger import logger
from config import settings

# ad = PromptServices()
openai_api_base, model_name = os.path.split(
    settings.AISHU_READER_LLM.rstrip("/"))


def get_llm_openai():
    llm = CustomChatOpenAI(
        model_name=model_name,
        openai_api_key=ad.get_appid(),
        openai_api_base=openai_api_base,
        max_tokens=10000,        # 减小 token 限制以提高响应速度
        temperature=0.2,        # 降低温度使输出更确定性
        request_timeout=60,     # 添加请求超时设置
        verify_ssl=False,
        # top_p=0.95,             # 添加 top_p 采样
        # presence_penalty=0.1,   # 添加存在惩罚以减少重复
        # frequency_penalty=0.1,  # 添加频率惩罚以增加多样性
    )
    return llm


def from_openai_style(max_tokens=8000, input_temperature=0.2):
    from app.cores.prompt.manage.ad_service import PromptServices
    ad = PromptServices()
    # api_key = ad.get_appid()
    llm = CustomChatOpenAI(
        model_name=model_name,
        openai_api_key=ad.get_appid(),
        openai_api_base=openai_api_base,
        max_tokens=max_tokens,  # 减小 token 限制以提高响应速度
        temperature=input_temperature,  # 降低温度使输出更确定性
        request_timeout=60,  # 添加请求超时设置
        verify_ssl=False,
        # top_p=0.95,             # 添加 top_p 采样
        # presence_penalty=0.1,   # 添加存在惩罚以减少重复
        # frequency_penalty=0.1,  # 添加频率惩罚以增加多样性
    )
    return llm

def from_final_openai_style(max_tokens=8000, input_temperature=0.2):
    from app.cores.prompt.manage.ad_service import PromptServices
    f_openai_api_base, f_model_name = os.path.split(
        settings.FINAL_READER_LLM.rstrip("/"))
    ad = PromptServices()
    # api_key = ad.get_appid()
    llm = CustomChatOpenAI(
        model_name=f_model_name,
        openai_api_key=ad.get_appid(),
        openai_api_base=f_openai_api_base,
        max_tokens=max_tokens,  # 减小 token 限制以提高响应速度
        temperature=input_temperature,  # 降低温度使输出更确定性
        request_timeout=60,  # 添加请求超时设置
        verify_ssl=False,
        # top_p=0.95,             # 添加 top_p 采样
        # presence_penalty=0.1,   # 添加存在惩罚以减少重复
        # frequency_penalty=0.1,  # 添加频率惩罚以增加多样性
    )
    return llm

def from_aishu(max_tokens=2000, input_temperature=0.2):
    llm = ChatOpenAI(
        model_name=model_name,
        openai_api_key="EMPTY",
        openai_api_base=openai_api_base,
        max_tokens=max_tokens,
        temperature=input_temperature,
    )
    return llm
#
def get_llm(max_input_tokens=8000, input_temperature=0.2):
    llm_style = settings.LLM_STYLE
    ad_version = settings.AD_VERSION
    logger.info("llm max token is {} temperature is {}".format(max_input_tokens, input_temperature))

    if ad_version == "3.0.1.0":
        llm_tools, llm_final = from_openai_style(max_input_tokens, input_temperature), from_final_openai_style(max_input_tokens, input_temperature)
    else:
        if llm_style == "AISHU":
            llm_tools, llm_final = from_aishu(max_input_tokens, input_temperature), from_aishu(max_input_tokens, input_temperature)
        else:
            llm_tools, llm_final = from_openai_style(max_input_tokens, input_temperature), from_final_openai_style(max_input_tokens, input_temperature)

    return llm_tools, llm_final


if __name__ == '__main__':
    m1, m2 = get_llm()
    res = m2.invoke("你好,你是谁")
    print(res)
