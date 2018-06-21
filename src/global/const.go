package global

var (
	VERSION   string // 版本号,通过编译参数改变,见build.go
	BUILDTIME string // 编译时间,通过编译参数获取
)

const (
	TablePrefix = "cmdb_v2_"                                     // 数据表前缀
	EncryptSalt = "ocpB8nZG5yBWrMfJDsM2fRB5L5LERMF47A6PWAC4wpM=" // 加密盐,不能小于32位
)

// 错误码
const (
	ErrNotNull     = "error key不能为空"
	ErrNotFound    = "error key不存在"
	ErrFormatError = "error key格式不正确"
)

// 帐号状态
const (
	StatusEnable  = 1 // 启用
	StatusDisable = 2 // 禁用
	StatusTrue    = 1 // 是否必须改密码(是)
	StatusFalse   = 2 // 是否必须改密码(否)
	StatusNone    = 1 // 未读
	StatusRead    = 2 // 已读
	StatusUse     = 1 // 使用
	StatusNoUse   = 2 // 未使用
	Disband       = 1 // 解散
)

// 注册方式
const (
	RegNormal  = 1 // 正常注册
	RegByAdmin = 2 // 管理员添加
	RegByLdap  = 3 // ldap创建
)

// 帐号审批状态
const (
	ExamineApprove = 1 // 待审批
	ExaminePass    = 2 // 已同意
	ExamineRefuse  = 3 // 已拒绝
)

// 角色
const (
	RoleSA  = 1  // 超级管理员
	RoleA   = 2  // 管理员
	RoleTL  = 3  // 技术主管
	RoleASL = 4  // 售后主管
	RolePSL = 5  // 售前主管
	RoleSD  = 6  // 销售主管
	RoleTS  = 7  // 技术支持
	RoleASS = 8  // 售后客服
	RolePSS = 9  // 售前客服
	RoleSP  = 10 // 销售推广
	RolePM  = 11 // 产品经理
	RoleOU  = 12 // 普通用户
)

// 操作日志设备类型
const (
	OPERATION_HOST        = 1 // 主机
	OPERATION_ROUTER      = 2 // 路由器
	OPERATION_SWITCHBOARD = 3 // 交换机
	OPERATION_FIREWALL    = 4 // 防火墙
)

// 订单服务类型
const (
	OrderServer = 1 // 服务器租用
	OrderIp     = 2 // ip租用
	OrderCdn    = 3 // cdn租用
)

// 是否试用
const (
	IfUse   = 1 // 是
	IfNoUse = 2 // 否
)

// cdn套餐
const (
	CdnHigh       = 1 // 高防
	CdnConnection = 2 // 无限制连接
	CdnDomain     = 3 // 无限制域名
)

// 订单状态
const (
	OrderNormal        = 1 // 进行中
	OrderReady         = 2 // 待续费
	OrderTimeout       = 3 // 已过期
	OrderAddAbnormal   = 4 // 新增未进行
	OrderRenewAbnormal = 5 // 续费未进行
)

// 订单日志类型
const (
	OrderAdd      = 1 // 创建
	OrderRefund   = 2 // 退款
	OrderDelay    = 3 // 延期
	OrderRenew    = 4 // 续费
	OrderDel      = 5 // 删除
	Orderrecovery = 6 // 恢复
)

// 业务类型
const (
	BusinessServerIp = 1 // 服务器+ip业务
	BusinessCdn      = 2 // cdn业务
)

// 产品类型
const (
	ProductHost        = 1 // 主机
	ProductRouter      = 2 // 路由器
	ProductSwitchboard = 3 // 交换机
	ProductFirewall    = 4 // 防火墙
	ProductNode        = 5 // 节点
)

// 路由请求方式
const (
	ROUTE_GET    = 1
	ROUTE_POST   = 2
	ROUTE_PUT    = 3
	ROUTE_DELETE = 4
)

// 冗余字段修改
const (
	IsNeed   = 1 // 需要修改
	IsNoNeed = 2 // 不需要修改
)

// 硬盘类型
const (
	HD_HHD = 1 // 硬盘类型 hhd
	HD_SSD = 2 // 硬盘类型 ssd
)

// 是否安装系统
const (
	IS_SYSTEM    = 1 // 是
	IS_NO_SYSTEM = 2 // 否
)

// 是否有业务
const (
	IS_BUSS = 1 // 有
	NO_BUSS = 2 // 没有
)

// 主机列表搜索框选项
const (
	TYPE_ID    = 1 // ID
	TYPE_ORDER = 2 // 订单
	TYPE_IP    = 3 // IP
)

// 列表搜索框选项
const (
	TypeId    = 1 // ID
	TypeOrder = 2 // 订单
	TypeIp    = 3 // IP
)

// 端口类型
const (
	PortNozzle  = 1 // 接口
	PortIpmi    = 2 // IPMI
	PortNetwork = 3 // 网卡
)

// 登录状态
const (
	LOGIN_SUCCESS = 1 // 成功
	LOGIN_FAIL    = 2 // 失败
)

// 接口类型
const (
	NozzleRj45 = 1 // RJ45类型
	NozzleSfp  = 2 // sfp类型
)

// 详细搜索
const (
	SEARCH_GT = 1 // 大于
	SEARCH_LT = 2 // 小于
)

// 消息变更类型
const (
	MessageResource = 1 // 资源池消息
	MessageTerm     = 2 // 维护组
	MessageOrder    = 3 // 订单
	MessageWork     = 4 // 工单
	MessageShift    = 5 // 交接班
)

// 时间戳常量
const (
	ThirtyDays   = 30 * 24 * 60 * 60 // 30天
	FIFTEEN_DAYS = 15 * 24 * 60 * 60 // 15天
	SEVEN_DAYS   = 7 * 24 * 60 * 60  // 7天
	HALF_HOUR    = 30 * 60           // 三十分钟
	OneDay       = 24 * 60 * 60      // 1天
	FiveDay      = 5 * 24 * 60 * 60  // 5天
)

// 负责人变更消息类型
const (
	BAND_DUTY    = 1 // 绑定主负责人
	REMOVE_DUTY  = 2 // 解除主要负责人
	BAND_SPARE   = 3 // 绑定副负责人
	REMOVE_SPARE = 4 // 解除副负责人
)

// 操作类型
const (
	TYPE_ADD    = 1 // 添加硬盘
	TYPE_DELETE = 2 // 删除硬盘
	TYPE_UPDATE = 3 // 修改
)

// 内存大小
const (
	THREE_MB = 3 * 1024 * 1024 // 3mb
)

// 是否为主机详细搜索
const (
	IS_DETAIL   = 1 // 是详细搜索
	IS_NODELETE = 2 // 不是详细搜索
)

// 推广类型
const (
	SPREAD_PC  = 1 // pc
	SPREAD_WAP = 2 // wap
)

// 查询类型
const (
	TYPE_DAY   = 1 // 今日
	TYPE_MONTH = 2 // 月份
)

// ip类型
const (
	IpNetwork = 1 // 网络地址
	IpGateway = 2 // 网关
)

// 主机租用状态
const (
	HostUse    = 1 // 自用
	HostNoRent = 2 // 未租用
	HostRent   = 3 // 已租用
)

// 是否高温
const (
	HighTem   = 1 // 高温
	NormalTem = 2 // 正常
)

// 监控
const (
	MonitorOn  = 1 // 开启
	MonitorOff = 2 // 关闭
)

// 是否自有
const (
	IsSelf   = 1 // 是
	IsNoSelf = 2 // 否
)

// 是否置顶
const (
	IsTop   = 1 // 是
	IsNoTop = 2 // 否
)

// 是否需要审核
const (
	IsCheck   = 1 // 是
	IsNoCheck = 2 // 否
)

// 主机状态
const (
	HostOn  = 1 // 开启
	HostOff = 2 // 关闭
)

// 防御类型
const (
	StandAlone = 1 // 单机防御
	HardGuard  = 2 // 硬防
	Colony     = 3 // 集群
)

// 维护组动态操作类型
const (
	Modify   = 1 // 更新
	Delete   = 2 // 删除
	Transfer = 3 // 移交
)

const TwoThousand = 2000 // 2000常量

// 是否置灰
const (
	IsNoAsh = 1 // 不置灰
	IsAsh   = 2 // 置灰
)

// 工单状态
const (
	ToBeTreated = 1 // 待处理
	HaveInHand  = 2 // 进行中
	Completed   = 3 // 已完成
	WorkClosed  = 4 // 已关闭
)

// 工单时间状态
const (
	DateFutureExpires = 1 // 未过期
	DateExpired       = 2 // 已过期
)

// 指派类型
const (
	Designate  = 1 // 指派
	BeAssigned = 2 // 被指派
)

// 是否能修改
const (
	CanChange    = 1 // 能修改
	CanNotChange = 2 // 不能修改
)

// 工单操作
const (
	WorkCreate = 1 // 创建
	WorkUpdate = 2 // 修改
	WorkClose  = 3 // 关闭
)

// 工单操作名称
const (
	WorkCreateMsg = "创建工单"
	WorkCloseMsg  = "关闭工单"
)

// 工单优先级
const (
	WorkLow         = 1 // 低
	WorkOrdinary    = 2 // 普通
	WorkHigh        = 3 // 高
	WorkUrgent      = 4 // 紧急
	WorkImmediately = 5 // 立刻
)

// 业务是否过期
const (
	IsTimeout   = 1 // 是
	IsNoTimeout = 0 // 否
)

// 是否为组长
const (
	IsLeader  = 1 // 是组长
	NotLeader = 2 // 不是组长
)

// 是否被添加
const (
	IsAdd  = 1 // 是
	NotAdd = 2 // 否
)

// 是否属于自己
const (
	IsOwner    = 1 // 是
	IsNotOwner = 2 // 否
)

// 首页销售额趋势筛选
const (
	ThisWeek  = 1 // 本周
	ThisMonth = 2 // 本月
	ThisYear  = 3 // 今年
)

const (
	IsPraise      = 1 // 已点赞
	IsDel         = 1 // 可以删除
	IsNoDel       = 2 // 不可以删除
	IsBeforeVer   = 1 // 有上一版本
	IsNoBeforeVer = 2 // 没有上一版本
	IsNextVer     = 1 // 有下一版本
	IsNoNextVer   = 2 // 没有下一版本
)

// 交接班图片附件个数
const (
	ShiftAttachmentNum = 10 // 交接班图片附件个数不能超过10个
	ShiftServicesNum   = 3  // 交接班提醒客服个数不能超过3个
)

// 历史记录版本
const (
	BeforeVersion = 1 // 上一版本
	NextVersion   = 2 // 下一版本
)

// 主机预警温度
const (
	HostEarlyWarningTemperature = 60
)
