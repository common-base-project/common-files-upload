package code

/*
  @Author : zggong
*/

var (
	// 通用
	InternalServerError = &Errno{ErrNo: 60001, ErrMsg: "内部服务器错误"}
	ExistError          = &Errno{ErrNo: 160003, ErrMsg: "数据已存在"}
	NotExistError       = &Errno{ErrNo: 10002, ErrMsg: "数据不存在"}
	ParamError          = &Errno{ErrNo: 10001, ErrMsg: "参数不正确"}
	BindError           = &Errno{ErrNo: 160006, ErrMsg: "绑定失败"}

	// 成功
	Success = &Errno{ErrNo: 0, ErrMsg: "请求成功"}

	// 未知失败
	UnknownError = &Errno{ErrNo: 199999, ErrMsg: "未知错误"}

	// namespace
	CreateNameSpaceError = &Errno{ErrNo: 160001, ErrMsg: "创建空间失败"}
	UpdateNameSpaceError = &Errno{ErrNo: 160002, ErrMsg: "更新空间失败"}
	DeleteNameSpaceError = &Errno{ErrNo: 160003, ErrMsg: "删除空间失败"}
	SelectNameSpaceError = &Errno{ErrNo: 160005, ErrMsg: "查询空间失败"}

	// project
	CreateProjectError = &Errno{ErrNo: 160011, ErrMsg: "创建项目失败"}
	UpdateProjectError = &Errno{ErrNo: 160012, ErrMsg: "更新项目失败"}
	DeleteProjectError = &Errno{ErrNo: 160013, ErrMsg: "删除项目失败"}
	SelectProjectError = &Errno{ErrNo: 160015, ErrMsg: "查询项目失败"}

	// template
	CreateTemplateError = &Errno{ErrNo: 160021, ErrMsg: "创建模版失败"}
	UpdateTemplateError = &Errno{ErrNo: 160022, ErrMsg: "更新模版失败"}
	DeleteTemplateError = &Errno{ErrNo: 160023, ErrMsg: "删除模版失败"}
	SelectTemplateError = &Errno{ErrNo: 160025, ErrMsg: "查询模版失败"}

	// field
	CreateFieldError = &Errno{ErrNo: 160031, ErrMsg: "添加模版字段失败"}
	UpdateFieldError = &Errno{ErrNo: 160032, ErrMsg: "更新模版字段失败"}
	DeleteFieldError = &Errno{ErrNo: 160033, ErrMsg: "删除模版字段失败"}
	SelectFieldError = &Errno{ErrNo: 160035, ErrMsg: "查询模版字段失败"}

	// workflow
	CreateWorkflowError = &Errno{ErrNo: 160051, ErrMsg: "创建工作流失败"}
	UpdateWorkflowError = &Errno{ErrNo: 160052, ErrMsg: "更新工作流失败"}
	DeleteWorkflowError = &Errno{ErrNo: 160053, ErrMsg: "删除工作流失败"}
	SelectWorkflowError = &Errno{ErrNo: 160055, ErrMsg: "查询工作流失败"}

	// status
	CreateStatusError = &Errno{ErrNo: 160061, ErrMsg: "创建工作流状态失败"}
	UpdateStatusError = &Errno{ErrNo: 160062, ErrMsg: "更新工作流状态失败"}
	DeleteStatusError = &Errno{ErrNo: 160063, ErrMsg: "删除工作流状态失败"}
	SelectStatusError = &Errno{ErrNo: 160065, ErrMsg: "查询工作流状态失败"}

	// transform
	CreateTransformError = &Errno{ErrNo: 160071, ErrMsg: "创建工作流流转失败"}
	UpdateTransformError = &Errno{ErrNo: 160072, ErrMsg: "更新工作流流转失败"}
	DeleteTransformError = &Errno{ErrNo: 160073, ErrMsg: "删除工作流流转失败"}
	SelectTransformError = &Errno{ErrNo: 160075, ErrMsg: "查询工作流流转失败"}

	// ticket
	SelectTicketError     = &Errno{ErrNo: 160081, ErrMsg: "查询工单数据失败"}
	ProcessingTicketError = &Errno{ErrNo: 160082, ErrMsg: "处理工单失败"}
	TransferTicketError   = &Errno{ErrNo: 160083, ErrMsg: "转交工单失败"}

	// ticket transform history
	SelectHistoryError = &Errno{ErrNo: 160091, ErrMsg: "获取工单流转历史错误"}
	CreateHistoryError = &Errno{ErrNo: 160092, ErrMsg: "创建工单流转历史错误"}

	// ticket feedback
	AddFeedbackError = &Errno{ErrNo: 160101, ErrMsg: "添加工单注释或者留言失败"}
	GetFeedbackError = &Errno{ErrNo: 160102, ErrMsg: "查询工单注释或者留言失败"}

	// group 1606xx
	CreateGroupError = &Errno{ErrNo: 160601, ErrMsg: "创建用户组失败"}
	UpdateGroupError = &Errno{ErrNo: 160602, ErrMsg: "更新用户组失败"}
	DeleteGroupError = &Errno{ErrNo: 160603, ErrMsg: "删除用户组失败"}
	SelectGroupError = &Errno{ErrNo: 160604, ErrMsg: "查询用户组失败"}

	// user 1607xx
	CreateUserError = &Errno{ErrNo: 170701, ErrMsg: "创建用户失败"}
	UpdateUserError = &Errno{ErrNo: 170702, ErrMsg: "更新用户失败"}
	DeleteUserError = &Errno{ErrNo: 170703, ErrMsg: "删除用户失败"}
	SelectUserError = &Errno{ErrNo: 170704, ErrMsg: "查询用户失败"}

	// depart 1608xx
	CreateDepartError = &Errno{ErrNo: 160801, ErrMsg: "创建部门失败"}
	UpdateDepartError = &Errno{ErrNo: 160802, ErrMsg: "更新部门失败"}
	DeleteDepartError = &Errno{ErrNo: 160803, ErrMsg: "删除部门失败"}
	SelectDepartError = &Errno{ErrNo: 160804, ErrMsg: "查询部门失败"}

	// usergroup 1709xx
	CreateUserGropuError = &Errno{ErrNo: 160901, ErrMsg: "创建用户和组关联失败"}
	UpdateUserGroupError = &Errno{ErrNo: 160902, ErrMsg: "更新用户和组关联失败"}
	DeleteUserGroupError = &Errno{ErrNo: 160903, ErrMsg: "删除用户和组关联失败"}
	SelectUserGroupError = &Errno{ErrNo: 160904, ErrMsg: "查询用户和组关联失败"}

	EndTicketError = &Errno{ErrNo: 160084, ErrMsg: "结束工单失败"}

	// help
	SelectHelpContentError = &Errno{ErrNo: 161001, ErrMsg: "获取帮助文档失败"}
	UpdateHelpContentError = &Errno{ErrNo: 161002, ErrMsg: "更新帮助文档失败"}

	// upload
	UploadError = &Errno{ErrNo: 161002, ErrData: nil, ErrMsg: "文档上传失败"}

	// ticket data
	ProcessingTicketDataError       = &Errno{ErrNo: 300001, ErrMsg: "处理数据 失败"}
	ProcessingTicketCreateDictError = &Errno{ErrNo: 300002, ErrMsg: "创建目录 失败"}
	ProcessingTicketDateError       = &Errno{ErrNo: 300003, ErrMsg: "日期输入错误 失败"}
)
