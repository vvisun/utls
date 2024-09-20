package errutil

var err_table = map[int32]string{
	Succ:           "成功",
	Sys_MidFeature: "功能尚未实现",
	Sys_UnKnown:    "未知错误",
	Sys_InvalidArg: "无效参数",
}
