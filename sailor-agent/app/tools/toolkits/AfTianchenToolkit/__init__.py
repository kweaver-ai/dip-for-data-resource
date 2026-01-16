from app.callbacks import AfAgentToolCallbackHandler
from app.tools.agent_tools import AFText2SQLTool, AFInterfaceTool, AFKnowledgeSearchTool, AFKnowledgeEnhancementTool
from app.tools.toolkits import InstructionBookInsideToolkit
from app.tools.agent_tools import AfHumanInputTool


def get_toolkit(url, token):
    # tool1 = AFInterfaceTool(url=url, token=token)
    tool2 = AFText2SQLTool(token=token)
    # tool3 = AFKnowledgeSearchTool(top=3, url=url, token=token)
    # tool4 = AFKnowledgeEnhancementTool()
    tool2.name = "自然语言查询数据库获取必要信息的工具"
    tool2.description = "通过输入自然语言从数据库中获取更多可以用于回答问题的信息。"
    af_data_analysis_toolkit = InstructionBookInsideToolkit()
    af_data_analysis_toolkit.set_toolkit_instruction('''
    无论问什么问题，首先使用{0}获取更多信息。然后必须将使用{0}获取的信息以markdown表格形式组织展示。
    '''.format(tool2.name))
    af_data_analysis_toolkit.set_tools([ tool2])
    # af_data_analysis_toolkit.set_tools([a_tool_human])
    return af_data_analysis_toolkit
