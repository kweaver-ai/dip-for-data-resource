from app.callbacks import AfAgentToolCallbackHandler
from app.tools.agent_tools import AfHumanInputTool
from app.tools.agent_tools import AfSailorTool, JsonToPlotTool
from app.tools.toolkits import InstructionBookInsideToolkit
from app.tools.agent_tools import ToolName

def get_toolkit(
    metadata,
    **kwargs
):
    sailor_tool = AfSailorTool(
        metadata=metadata,
        **kwargs
    )
    # human_tool = AfHumanInputTool(
    #     callbacks=[AfAgentToolCallbackHandler()]
    # )
    json_to_plot_tool = JsonToPlotTool(metadata=metadata)
    toolkit = InstructionBookInsideToolkit()
    toolkit.set_toolkit_instruction(
        f'''
        1. 开发者对工具做出了特别设定，当工具返回提示你结束此次问答时，请结束此次问答，并且不反问用户；
        2. 每种工具调用最多不超过两次；
        3. 如果用户提示直接作图，你可以首先调用 {ToolName.JsonToPlotToolName.value}，如果工具调用失败，你再根据反馈继续思考；
        4. 你需要仔细查看 先前的对话历史，并明确用户的意图，然后调用合适的工具；
        5. 当用户在重复一些内容时，你需要将这个内容传给 {ToolName.AfSailorToolName.value}，以获取更加符合用户意图的数据；
        6. 使用中文；
        7. 你看不到工具执行成功的结果，只能看到工具执行失败的结果，所以你要尽可能的让工具完成所有任务。
        8. 当用户的问题不明确的时候，你需要执行 {ToolName.AfSailorToolName.value}， 工具会返回搜索的结果，
        '''
    )

    # toolkit.set_tools([human_tool, sailor_tool])
    # toolkit.set_tools([sailor_tool, json_to_plot_tool, human_tool])
    toolkit.set_tools([sailor_tool, json_to_plot_tool])

    return toolkit
    # 当你拿到 {ToolName.JsonToPlotToolName.value} 的结果以后，其数据已经可以作为结果进行返回，不需要生成一个网址进行展示，
    # 如果用户要求用图进行表示，最终的 Final Answer 返回图数据即可。
#         3. 在使用 {ToolName.AfHumanInputToolName.value} 工具以后，不能再使用其它工具，请直接结束本次问答；