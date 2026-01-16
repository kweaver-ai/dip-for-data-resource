package intelligence

var Example1 = `["姓名:str","年龄:int","城市:str","主键1:str","主键2:int"]`
var Example2 = `{"count": 1,"sample_data":[{"姓名":"张三","年龄":"25","城市":"西安","主键1":"38619cb5-81c3-406d-adb6-57f369938034", "主键2":"1"}]}`

var SampleDataPromptNoExample = `样例数据生成任务描述：根据表头生成样例数据：比如表头为%s,结果为：%s ` +
	`那么现在执行样例数据生成任务,用报文回复结果，表头为%s,请生成5条样例数据("count":5),` +
	`并保证str类型及%s每条数据没有任何重复，结果为:`

var SampleDataPromptWithExample = `样例数据生成任务描述：根据表头生成样例数据：比如表头为%s,结果为：%s ` +
	`那么现在执行样例数据生成任务,用报文回复结果，表头为%s,表中已有样例为:%s,请再生成5条样例数据("count":5),` +
	`并保证str类型及%s每条数据没有任何重复，结果为:`
