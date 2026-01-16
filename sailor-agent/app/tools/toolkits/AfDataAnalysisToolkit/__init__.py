from app.callbacks import AfAgentToolCallbackHandler
from app.tools.agent_tools import AFText2SQLTool, AFInterfaceTool, AFKnowledgeSearchTool, AFKnowledgeEnhancementTool
from app.tools.toolkits import InstructionBookInsideToolkit
from app.tools.agent_tools import AfHumanInputTool


def get_toolkit(url, token):
    tool1 = AFInterfaceTool(url=url, token=token)
    tool2 = AFText2SQLTool(token=token)
    tool3 = AFKnowledgeSearchTool(top=3, url=url, token=token)
    tool4 = AFKnowledgeEnhancementTool()

    af_data_analysis_toolkit = InstructionBookInsideToolkit()
    af_data_analysis_toolkit.set_toolkit_instruction('''
    当你判断用户最新的对话正在询问你一个新的数据分析领域的问题之后，首先基于背景知识将相对指代转化为绝对指代，然后你必须要尝试按以下流程步骤仔细思考并安排确定你的下一步action：

    步骤1:在判断用户询问你一个新的数据分析领域的问题之后的第一步action使用"通过文本召回相关接口服务及表格工具"工具，召回与这个问题最相关的数据表以及数据接口,然后在下一步action执行步骤2；
    步骤2:根据召回的接口信息的输入参数与返回参数信息以及问句信息，判断是否需要知识增强工具获得更多背景知识，如果是，在下一个action中对问题使用"背景知识增强工具"并更新问题后再在之后的下一步action中执行步骤3，如果不是，则直接执行步骤3；
    步骤3:判断召回的相关接口服务及表格能否回答当前问题，如果能，那么挑选你觉得有用的一组接口服务和数据表,接着在下一步action中从你挑出来的一组接口服务中挑一个接口服务，尝试生成对应的入参json格式代码片段，如果必要的入参可以填写完整，那么在下一步action中使用"通过JSON调用接口工具"获得信息，如果不能，那么在下一步action中使用"询问人类"工具尝试获得你觉得缺失了的输入参数，然后再再之后一步action中使用"通过JSON调用接口工具"获得信息；
    步骤4：判断Observation的内容是否已经有助于回答问题的最终答案，如果是，那么下一个action执行步骤7，如果不是，那么下一个action回到步骤3执行重复动作但是不要与之前已经调用过的你挑出来的接口重复，如果步骤3中挑出来的接口服务已经尝试完，那么下一步action执行步骤5；
    步骤5：判断哪个数据表可能有助于回答问题，然后在下一步action通过传入问题与表名，使用"数据库检索工具"尝试获得回答问题的依据；
    步骤6：判断上一步action执行工具后获得的Observation是否已经有助于回答问题的最终答案，如果是，那么下一个action执行步骤7，如果不是，那么下一个action回到步骤5执行重复动作但是不要与之前已经调用过的你挑出来的数据表重复，如果步骤5中挑出来的接口服务已经尝试完，那么下一步action执行步骤7；
    步骤7：回答最终答案，时刻谨记在最后的答案前回复`Final Answer`这个短语！在最后的答案中一定要首先说出你调用了哪些接口和使用了哪些数据表来回答问题。即使你无法回答或者回答失败了，那么也回答你查询到了的步骤1中获得的相关数据表和数据接口可能有助于回答问题。
    注意！永远不要假设或者编造任何参数，字段，接口，数据目录名称！！！！！永远不要在Action中使用不存在的工具或者编造工具。
    ''')

    a_tool_human = AfHumanInputTool(callbacks=[AfAgentToolCallbackHandler()])
    # a_tool_human.description = (
    #     "当通过观察observation后，你觉得难以难以得出下一步action，或者难以回答问题，或者需要补充信息，你可以询问用户来协助指引你思考。该工具的输入参数是一个你希望向人类提的用来协助引导你思考下一步Action的问题。每当使用此工具后，请一定在下一步Action中将上一步Action的Observation作为一个回复人类的最终回答原原本本地输出出来并结束对话。")

    af_data_analysis_toolkit.set_tools([a_tool_human,tool1, tool2, tool3, tool4])
    # af_data_analysis_toolkit.set_tools([a_tool_human, tool1, tool2, tool3, tool4])
    # af_data_analysis_toolkit.set_tools([a_tool_human])


    return af_data_analysis_toolkit


