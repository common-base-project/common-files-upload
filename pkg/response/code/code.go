package code

/*
  @Author : zggong
*/

var (
	// 通用
	InternalServerError = &Errno{Errno: 60001, Errmsg: "内部服务器错误"}
	ExistError          = &Errno{Errno: 160003, Errmsg: "数据已存在"}
	NotExistError       = &Errno{Errno: 10002, Errmsg: "数据不存在"}
	ParamError          = &Errno{Errno: 10001, Errmsg: "参数不正确"}
	BindError           = &Errno{Errno: 160006, Errmsg: "绑定失败"}

	// 成功
	Success = &Errno{Errno: 0, Errmsg: "请求成功"}

	// 未知失败
	UnknownError = &Errno{Errno: 199999, Errmsg: "未知错误"}

	// namespace
	CreateNameSpaceError = &Errno{Errno: 160001, Errmsg: "创建空间失败"}
	UpdateNameSpaceError = &Errno{Errno: 160002, Errmsg: "更新空间失败"}
	DeleteNameSpaceError = &Errno{Errno: 160003, Errmsg: "删除空间失败"}
	SelectNameSpaceError = &Errno{Errno: 160005, Errmsg: "查询空间失败"}

	// project
	CreateProjectError = &Errno{Errno: 160011, Errmsg: "创建项目失败"}
	UpdateProjectError = &Errno{Errno: 160012, Errmsg: "更新项目失败"}
	DeleteProjectError = &Errno{Errno: 160013, Errmsg: "删除项目失败"}
	SelectProjectError = &Errno{Errno: 160015, Errmsg: "查询项目失败"}

	// template
	CreateTemplateError = &Errno{Errno: 160021, Errmsg: "创建模版失败"}
	UpdateTemplateError = &Errno{Errno: 160022, Errmsg: "更新模版失败"}
	DeleteTemplateError = &Errno{Errno: 160023, Errmsg: "删除模版失败"}
	SelectTemplateError = &Errno{Errno: 160025, Errmsg: "查询模版失败"}

	// field
	CreateFieldError = &Errno{Errno: 160031, Errmsg: "添加模版字段失败"}
	UpdateFieldError = &Errno{Errno: 160032, Errmsg: "更新模版字段失败"}
	DeleteFieldError = &Errno{Errno: 160033, Errmsg: "删除模版字段失败"}
	SelectFieldError = &Errno{Errno: 160035, Errmsg: "查询模版字段失败"}

	// workflow
	CreateWorkflowError = &Errno{Errno: 160051, Errmsg: "创建工作流失败"}
	UpdateWorkflowError = &Errno{Errno: 160052, Errmsg: "更新工作流失败"}
	DeleteWorkflowError = &Errno{Errno: 160053, Errmsg: "删除工作流失败"}
	SelectWorkflowError = &Errno{Errno: 160055, Errmsg: "查询工作流失败"}

	// status
	CreateStatusError = &Errno{Errno: 160061, Errmsg: "创建工作流状态失败"}
	UpdateStatusError = &Errno{Errno: 160062, Errmsg: "更新工作流状态失败"}
	DeleteStatusError = &Errno{Errno: 160063, Errmsg: "删除工作流状态失败"}
	SelectStatusError = &Errno{Errno: 160065, Errmsg: "查询工作流状态失败"}

	// transform
	CreateTransformError = &Errno{Errno: 160071, Errmsg: "创建工作流流转失败"}
	UpdateTransformError = &Errno{Errno: 160072, Errmsg: "更新工作流流转失败"}
	DeleteTransformError = &Errno{Errno: 160073, Errmsg: "删除工作流流转失败"}
	SelectTransformError = &Errno{Errno: 160075, Errmsg: "查询工作流流转失败"}

	// ticket
	SelectTicketError     = &Errno{Errno: 160081, Errmsg: "查询工单数据失败"}
	ProcessingTicketError = &Errno{Errno: 160082, Errmsg: "处理工单失败"}
	TransferTicketError   = &Errno{Errno: 160083, Errmsg: "转交工单失败"}

	// ticket transform history
	SelectHistoryError = &Errno{Errno: 160091, Errmsg: "获取工单流转历史错误"}
	CreateHistoryError = &Errno{Errno: 160092, Errmsg: "创建工单流转历史错误"}

	// ticket feedback
	AddFeedbackError = &Errno{Errno: 160101, Errmsg: "添加工单注释或者留言失败"}
	GetFeedbackError = &Errno{Errno: 160102, Errmsg: "查询工单注释或者留言失败"}

	// group 1606xx
	CreateGroupError = &Errno{Errno: 160601, Errmsg: "创建用户组失败"}
	UpdateGroupError = &Errno{Errno: 160602, Errmsg: "更新用户组失败"}
	DeleteGroupError = &Errno{Errno: 160603, Errmsg: "删除用户组失败"}
	SelectGroupError = &Errno{Errno: 160604, Errmsg: "查询用户组失败"}

	// user 1607xx
	CreateUserError = &Errno{Errno: 170701, Errmsg: "创建用户失败"}
	UpdateUserError = &Errno{Errno: 170702, Errmsg: "更新用户失败"}
	DeleteUserError = &Errno{Errno: 170703, Errmsg: "删除用户失败"}
	SelectUserError = &Errno{Errno: 170704, Errmsg: "查询用户失败"}

	// depart 1608xx
	CreateDepartError = &Errno{Errno: 160801, Errmsg: "创建部门失败"}
	UpdateDepartError = &Errno{Errno: 160802, Errmsg: "更新部门失败"}
	DeleteDepartError = &Errno{Errno: 160803, Errmsg: "删除部门失败"}
	SelectDepartError = &Errno{Errno: 160804, Errmsg: "查询部门失败"}

	// usergroup 1709xx
	CreateUserGropuError = &Errno{Errno: 160901, Errmsg: "创建用户和组关联失败"}
	UpdateUserGroupError = &Errno{Errno: 160902, Errmsg: "更新用户和组关联失败"}
	DeleteUserGroupError = &Errno{Errno: 160903, Errmsg: "删除用户和组关联失败"}
	SelectUserGroupError = &Errno{Errno: 160904, Errmsg: "查询用户和组关联失败"}

	EndTicketError = &Errno{Errno: 160084, Errmsg: "结束工单失败"}

	// help
	SelectHelpContentError = &Errno{Errno: 161001, Errmsg: "获取帮助文档失败"}
	UpdateHelpContentError = &Errno{Errno: 161002, Errmsg: "更新帮助文档失败"}

	// upload
	UploadError = &Errno{Errno: 161002, Errmsg: "文档上传失败"}

	// ticket data
	ProcessingTicketDataError       = &Errno{Errno: 300001, Errmsg: "处理数据 失败"}
	ProcessingTicketCreateDictError = &Errno{Errno: 300002, Errmsg: "创建目录 失败"}
	ProcessingTicketDateError       = &Errno{Errno: 300003, Errmsg: "日期输入错误 失败"}
)
