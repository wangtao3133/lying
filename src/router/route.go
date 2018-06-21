package router

import (
	"controller"
	"github.com/labstack/echo"
)

func Route(c *echo.Echo) {
	s := new(controller.Login)
	c.GET("captcha", s.Captcha, CheckCaptchaRequestNum) // 验证码路由
	c.POST("login", s.SignIn, CheckLoginRequestErr)     // 登录

	r := new(controller.Register)
	c.POST("/register", r.Register)       // 注册
	c.GET("/register/account", r.Account) // 帐号唯一性校验

	e := c.Group("", CheckToken, PowerValid, AccessLog) // 开启中间件校验

	e.GET("logout", s.LogOut) // 退出登录

	// 角色
	role := new(controller.Role)
	e.GET("/role", role.List)     // 角色列表
	e.GET("/roleDrop", role.Drop) // 角色下拉框

	// 登录人操作
	e.PUT("password", s.ModifyPassword) // 修改登陆人密码
	e.GET("user", s.Info)               // 展示登陆人信息
	e.PUT("user", s.ModifyInfo)         // 修改登陆人信息

	// 账号管理
	a := new(controller.Account)
	e.GET("account/list", a.List)           // 账号列表
	e.DELETE("account", a.Delete)           // 删除账号
	e.POST("account", a.Add)                // 添加账号
	e.GET("account/info", a.Info)           // 账号详情
	e.PUT("account", a.Update)              // 修改账号
	e.GET("examine", a.ExamineList)         // 审批列表
	e.DELETE("examine", a.DelExamine)       // 删除审批账号
	e.PUT("examine", a.UpdateExamineStatus) // 修改审批状态

	l := new(controller.Ldap)
	e.GET("ldap", l.Info)      // 登录认证详情
	e.POST("ldap", l.Add)      // 添加登录认证
	e.GET("ldap/test", l.Auth) // 测试登录认证

	e.GET("/home", s.HomeView)            // 首页查看
	e.GET("/home/work", s.HomeWork)       // 首页工单
	e.GET("/total/sale", s.HomeTotalSale) // 首页销售额趋势
	// 下拉框
	drop := new(controller.Drop)
	e.GET("/systemVersion/drop", drop.SystemVersionDrop)  // 系统版本下拉框
	e.GET("/area/drop", drop.AreaDrop)                    // 区域下拉框
	e.GET("/room/drop", drop.RoomDrop)                    // 机房下拉框
	e.GET("/cabinet/drop", drop.CabinetDrop)              // 机柜下拉框
	e.GET("/leader/drop", drop.GroupLeader)               // 维护组组长下拉框
	e.GET("/member/drop", drop.GroupMember)               // 维护组组员下拉框
	e.GET("/duty/drop", drop.Duty)                        // 产品负责人下拉框
	e.GET("/server/drop", drop.Server)                    // 服务器配置下拉框
	e.GET("/order/drop", drop.Order)                      // 订单下拉框
	e.GET("/custom/drop", drop.Custom)                    // 客服下拉框
	e.GET("/line/drop", drop.Line)                        // 线路下拉框
	e.GET("/ips/drop", drop.Ips)                          // IP段下拉框
	e.GET("/ip/drop", drop.Ip)                            // ip列表下拉框
	e.GET("/user/drop", drop.User)                        // 客户下拉框
	e.GET("/template/drop", drop.TemplateDrop)            // 模版列表
	e.GET("/term/drop", drop.Term)                        // 维护组下拉框
	e.GET("/assign/drop", drop.Assign)                    // 指派人下拉框
	e.GET("/site/drop", drop.Site)                        // 推广人员站点下拉框
	e.GET("/supplierr/drop", drop.Supplierr)              // 供应商下拉框
	e.GET("/supplier/account/drop", drop.SupplierAccount) // 供应商账号下拉框

	// 角色权限
	rolePower := new(controller.RolePower)
	e.GET("/role/power", rolePower.List)          // 获取角色对应权限
	e.POST("/role/power", rolePower.AddRolePower) // 给角色分配权限

	// 角色菜单
	roleMenu := new(controller.RoleMenu)
	e.GET("/role/menu", roleMenu.List) // 获取角色对应菜单
	e.POST("/role/menu", roleMenu.Add) // 给角色分配菜单

	// 权限管理
	power := new(controller.Power)
	e.GET("/power/list", power.List)           // 权限列表
	e.GET("/power/info", power.Info)           // 权限详情
	e.PUT("/power", power.Update)              // 修改权限
	e.PUT("/power/status", power.UpdateStatus) // 修改权限状态
	e.POST("/power", power.Add)                // 增加权限
	e.DELETE("/power", power.Delete)           // 删除权限

	// 菜单管理
	menu := new(controller.Menu)
	e.GET("/menu/navigation", menu.MenuNavigation) // 菜单导航
	e.GET("/menu/list", menu.List)                 // 菜单列表
	e.GET("/menu/info", menu.Info)                 // 菜单详情
	e.PUT("/menu", menu.Update)                    // 修改菜单
	e.PUT("/menu/status", menu.UpdateStatus)       // 修改菜单状态
	e.POST("/menu", menu.Add)                      // 增加菜单
	e.DELETE("/menu", menu.Delete)                 // 删除菜单

	// 客户管理
	u := new(controller.User)
	e.GET("/user/list", u.List)       // 客户列表
	e.PUT("/user/update", u.Update)   // 修改客户
	e.POST("/user", u.Add)            // 增加客户
	e.DELETE("/user", u.Delete)       // 删除客户
	e.GET("/user/info", u.Info)       // 客户详情
	e.GET("/user/order", u.OrderInfo) // 客户订单详情

	// 产品管理
	ssl := new(controller.Ssl)
	e.GET("/ssl/list", ssl.List) // 证书列表
	e.POST("/ssl", ssl.Add)      // 添加证书
	e.PUT("/ssl", ssl.Change)    // 修改证书
	e.GET("/ssl", ssl.Info)      // 获取证书
	e.DELETE("/ssl", ssl.Delete) // 删除证书

	// 订单管理
	o := new(controller.Order)
	e.GET("/order/list", o.List)             // 订单列表
	e.POST("/order", o.Add)                  // 新增订单
	e.GET("/order/total", o.TotalMoney)      // 订单总销售额
	e.DELETE("/order", o.Delete)             // 删除订单
	e.GET("/order/info", o.Info)             // 订单详情
	e.PUT("/order", o.RenewRent)             // 订单续费
	e.PUT("/order/refund", o.Refund)         // 订单退款
	e.PUT("/order/feedback", o.Feedback)     // 编辑订单退款原因
	e.GET("/order/feedback", o.FeedbackInfo) // 获取订单退款原因
	e.PUT("/order/delay", o.Delay)           // 订单延期

	// 业务管理
	bss := new(controller.Business)
	e.GET("/business/host", bss.HostList)          // 主机列表
	e.GET("/business/server", bss.ServerList)      // 服务器业务列表
	e.GET("/business/cdn", bss.CdnList)            // cdn业务列表
	e.DELETE("/business", bss.Del)                 // 删除业务
	e.POST("/business/server", bss.ServerAdd)      // 添加服务器+ip业务
	e.POST("/business/cdn", bss.CdnAdd)            // 添加cdn业务
	e.GET("/business/server/info", bss.ServerInfo) // 服务器业务详情
	e.GET("/business/cdn/info", bss.CdnInfo)       // cdn业务详情
	e.PUT("/business/cdn", bss.CdnUpdate)          // 修改cdn业务
	e.GET("/business/info", bss.Info)              // 服务器业务修改前的详情
	e.PUT("/business/server", bss.ServerUpdate)    // 修改服务器业务

	// 新增服务器业务中的IP配置
	i := new(controller.Ip)
	e.POST("/ip/conf", i.Add)   // 新增ip配置
	e.DELETE("/ip/conf", i.Del) // 删除ip配置

	// 订单日志列表
	ol := new(controller.OrderLog)
	e.GET("/order/log", ol.List) // 订单日志

	// 区域
	area := new(controller.Area)
	e.GET("/area/list", area.List) // 区域列表
	e.POST("/area", area.Add)      // 区域添加
	e.PUT("/area", area.Update)    // 区域修改
	e.DELETE("/area", area.Delete) // 区域删除
	e.GET("/area/info", area.Info) // 区域详情

	// 机房
	room := new(controller.Room)
	e.GET("/room/list", room.List)                   // 机房列表
	e.POST("/room", room.Add)                        // 机房添加
	e.PUT("/room", room.Update)                      // 机房修改
	e.DELETE("/room", room.Delete)                   // 机房删除
	e.GET("/room/info", room.Info)                   // 机房详情
	e.POST("/room/salePerson", room.SalePersonAdd)   // 添加销售人员
	e.DELETE("/room/salePerson", room.SalePersonDel) // 删除销售人员

	// 机柜
	cabinet := new(controller.Cabinet)
	e.POST("/cabinet", cabinet.Add)                        // 添加机柜
	e.PUT("/cabinet", cabinet.Update)                      // 机柜修改
	e.DELETE("/cabinet", cabinet.Delete)                   // 机柜删除
	e.GET("/temperature", cabinet.GetTemperature)          // 获取预警温度
	e.PUT("/temperature", cabinet.UpdateTemperature)       // 设置预警温度
	e.GET("/cabinet/list", cabinet.List)                   // 机柜列表
	e.GET("/cabinet/info", cabinet.Info)                   // 机柜详情
	e.GET("/cabinet/visualization", cabinet.Visualization) // 可视化机柜
	e.PUT("/cabinet/visualization", cabinet.DeviceSort)    // 修改排序

	// 线路
	line := new(controller.Line)
	e.POST("/cabinet/line", line.Add)             // 添加主线路(机柜修改页面的添加)
	e.PUT("/cabinet/line", line.Update)           // 修改主线路(机柜修改页面的添加)
	e.DELETE("/cabinet/line", line.Delete)        // 删除线路
	e.POST("/cabinet/lineChild", line.LcAdd)      // 添加子线路
	e.PUT("/cabinet/lineChild", line.LcUpdate)    // 修改子线路
	e.DELETE("/cabinet/lineChild", line.LcDelete) // 删除子线路

	// 主机
	host := new(controller.Host)
	e.POST("/host", host.Add)                    // 添加主机
	e.GET("/host/list", host.List)               // 主机列表
	e.PUT("/host", host.Update)                  // 修改主机
	e.DELETE("/host", host.Delete)               // 删除主机
	e.GET("/host/info", host.Info)               // 主机详情
	e.PUT("/host/status", host.Status)           // 更改主机状态
	e.POST("/host/hardDisk", host.HdAdd)         // 添加硬盘
	e.PUT("/host/hardDisk", host.HdUpdate)       // 修改硬盘
	e.DELETE("/host/hardDisk", host.HdDelete)    // 删除硬盘
	e.POST("/host/networkCard", host.NcAdd)      // 添加网卡
	e.PUT("/host/networkCard", host.NcUpdate)    // 修改网卡
	e.DELETE("/host/networkCard", host.NcDelete) // 删除网卡
	e.GET("/ipmi/password", host.IpmiPassword)   // 查看ipmi密码
	e.PUT("/ipmi/password", host.PasswordUpdate) // 修改ipmi密码

	// 节点分组区域
	ng := new(controller.NodeGroup)
	e.POST("/node/group", ng.Add)      // 添加节点分组区域
	e.GET("/node/group", ng.List)      // 节点分组区域列表
	e.GET("/node/group/drop", ng.Drop) // 节点分组区域下拉框
	e.PUT("/node/group", ng.Update)    // 修改节点分组区域
	e.DELETE("/node/group", ng.Delete) // 删除节点分组区域

	// 节点分组解析
	ngp := new(controller.NodeGroupParse)
	e.POST("/node/group/parse", ngp.Add)      // 添加节点分组解析
	e.GET("/node/group/parse", ngp.List)      // 节点分组解析列表
	e.PUT("/node/group/parse", ngp.Update)    // 修改节点分组解析
	e.DELETE("/node/group/parse", ngp.Delete) // 删除节点分组解析
	e.GET("/node/group/info", ngp.NodeInfo)   // 节点详情
	e.GET("/ngp/drop", ngp.Drop)              // 解析分组下拉框

	// 节点
	n := new(controller.Node)
	e.GET("/node/drop", n.Drop)                       // 节点下拉框
	e.GET("/node/ip/drop", n.IpDrop)                  // 网卡ip下拉框
	e.GET("/node/system/drop", n.SystemDrop)          // 系统下拉框
	e.POST("/node", n.Add)                            // 添加节点
	e.GET("/node", n.List)                            // 节点列表
	e.GET("/node/password", n.SeePassword)            // 查看远程卡密码
	e.DELETE("/node", n.Del)                          // 删除节点
	e.PUT("/node", n.Update)                          // 修改节点
	e.PUT("/node/ipmipassword", n.UpdateIpmiPassword) // 修改节点ipmi密码

	// 路由器
	route := new(controller.Route)
	e.GET("/route", route.List)               // 路由器列表
	e.POST("/route", route.Add)               // 添加
	e.GET("/route/info", route.Info)          // 路由器详情
	e.GET("/route/see", route.SeePassword)    // 路由器查看密码
	e.PUT("/route/see", route.PasswordUpdate) // 路由器修改密码
	e.PUT("/route", route.Update)             // 修改路由器
	e.DELETE("/route", route.Delete)          // 删除路由器

	// 交换机
	sw := new(controller.Switchboard)
	e.GET("/switchboard", sw.List)               // 交换机列表
	e.POST("/switchboard", sw.Add)               // 添加交换机
	e.GET("/switchboard/info", sw.Info)          // 交换机详情
	e.GET("/switchboard/see", sw.SeePassword)    // 交换机查看密码
	e.PUT("/switchboard/see", sw.PasswordUpdate) // 交换机修改密码
	e.PUT("/switchboard", sw.Update)             // 修改交换机
	e.DELETE("/switchboard", sw.Delete)          // 删除交换机

	// 防火墙
	fw := new(controller.Firewall)
	e.GET("/firewall", fw.List)               // 防火墙列表
	e.POST("/firewall", fw.Add)               // 添加防火墙
	e.GET("/firewall/info", fw.Info)          // 防火墙详情
	e.GET("/firewall/see", fw.SeePassword)    // 防火墙查看密码
	e.PUT("/firewall/see", fw.PasswordUpdate) // 防火墙修改密码
	e.PUT("/firewall", fw.Update)             // 修改防火墙
	e.DELETE("/firewall", fw.Delete)          // 删除防火墙

	// 供应商
	su := new(controller.Supplier)
	e.GET("/supplier", su.List)                     // 供应商列表
	e.GET("/supplier/info", su.Info)                // 供应商详情
	e.POST("/supplier", su.Add)                     // 添加供应商
	e.PUT("/supplier", su.Update)                   // 修改供应商
	e.DELETE("/supplier", su.Delete)                // 删除供应商
	e.POST("/supplier/account", su.AddAccount)      // 添加账号
	e.PUT("/supplier/account", su.UpdateAccount)    // 修改账号
	e.DELETE("/supplier/account", su.DeleteAccount) // 删除账号
	e.GET("/supplier/password", su.SeePassword)     // 查看密码

	// 域名
	re := new(controller.Realm)
	e.GET("/realm", re.List)      // 域名列表
	e.POST("/realm", re.Add)      // 添加域名
	e.PUT("/realm", re.Update)    // 修改域名
	e.DELETE("/realm", re.Delete) // 删除域名

	// 备份
	b := new(controller.Backups)
	e.POST("/backups", b.Upload)          // 上传
	e.GET("/backups/list", b.BackupsList) // 备份列表
	e.DELETE("/backups", b.BackupsDelete) // 删除
	e.GET("/backups", b.BackupsDownLoad)  // 下载

	// 维护组
	t := new(controller.Term)
	e.GET("/term", t.List)                 // 维护组列表
	e.PUT("/term/top", t.IsTop)            // 置顶维护组
	e.POST("/term", t.Add)                 // 新增维护组
	e.GET("/term/info", t.Info)            // 维护组详情
	e.PUT("/term", t.Update)               // 修改维护组
	e.DELETE("/term", t.Disband)           // 解散维护组
	e.POST("/term/band", t.BandGroup)      // 绑定维护组
	e.GET("/term/list", t.DeviceList)      // 设备列表
	e.GET("/term/examine", t.ExamineList)  // 设备待审批列表
	e.PUT("/term/pass", t.ExaminePass)     // 审批通过
	e.PUT("/term/refuse", t.ExamineRefuse) // 审批拒绝
	e.PUT("/term/mark", t.MarkUpdate)      // 更新标签
	e.PUT("/term/transfer", t.Transfer)    // 设备移交

	// 维护组动态
	tt := new(controller.TermTrend)
	e.GET("/term/trend", tt.List)      // 组动态列表
	e.GET("/term/trend/info", tt.Info) // 组动态查看

	// 设备接口
	p := new(controller.Port)
	e.POST("/port", p.Band)     // 接口绑定
	e.GET("/port/list", p.List) // 接口列表
	e.GET("/port/info", p.Info) // 接口详情
	e.PUT("/port", p.Update)    // 接口修改
	e.DELETE("/port", p.Remove) // 接口删除

	// 高防套餐
	high := new(controller.High)
	e.POST("/high", high.Add)      // 添加高防套餐
	e.GET("/high", high.List)      // 高防套餐列表
	e.PUT("/high", high.Update)    // 高防套餐修改
	e.DELETE("/high", high.Delete) // 高防套餐删除

	// 无限域名套餐
	domain := new(controller.Domain)
	e.POST("/domain", domain.Add)      // 添加无限域名套餐
	e.GET("/domain", domain.List)      // 无限域名套餐列表
	e.PUT("/domain", domain.Update)    // 无限域名套餐修改
	e.DELETE("/domain", domain.Delete) // 无限域名套餐删除

	// 无限连接数套餐
	connection := new(controller.Connection)
	e.POST("/connection", connection.Add)      // 添加无限连接数套餐
	e.GET("/connection", connection.List)      // 无限连接数套餐列表
	e.PUT("/connection", connection.Update)    // 无限连接数套餐修改
	e.DELETE("/connection", connection.Delete) // 无限连接数套餐删除

	// 日志
	operation := new(controller.Operation)
	e.GET("/operation/log", operation.List)  // 操作日志列表
	e.GET("/entry/log", operation.LoginList) // 登陆日志列表

	// 消息中心
	nc := new(controller.Notice)
	e.GET("/notice/list", nc.List) // 消息列表
	e.PUT("/notice", nc.Update)    // 更改消息状态
	e.GET("/notice", nc.NoneRead)  // 顶部未读消息
	e.DELETE("/notice", nc.Delete) // 消息删除

	// 推广管理
	sc := new(controller.Spread)
	e.GET("/spread/list", sc.List)             // 推广明细列表
	e.GET("/spread/roleId", sc.GetRoleId)      // 获取登陆人的角色id
	e.GET("/spread/traffic", sc.SiteTraffic)   // 网站流量
	e.GET("/spread/trend", sc.SiteTrend)       // 网站流量趋势
	e.GET("/spread/make", sc.MakeSpreadUrl)    // 生成推广链接(推广人员操作)
	e.POST("/spread", sc.Add)                  // 添加推广人员信息(推广人员操作)
	e.GET("/spread/aInfo", sc.AInfo)           // 管理员查看推广人员详情
	e.GET("/spread/info", sc.Info)             // 获取推广人员名称和链接(推广人员操作)
	e.PUT("/spread", sc.Update)                // 修改站点和qq(推广人员操作)
	c.POST("/spread/count", sc.AddCount)       // 添加推广链接点击次数(推广人员操作)
	e.GET("/spread/lInfo", sc.LInfo)           // 获取推广链接点击列表详情(推广人员操作)
	e.GET("/spread/laInfo", sc.LaInfo)         // 获取推广链接点击列表详情
	e.GET("/spread/page", sc.PageView)         // 网站推广访问量(推广人员操作)
	e.GET("/spread/aPage", sc.PageViewByAdmin) // 网站推广访问量

	// 服务器配置
	server := new(controller.Server)
	e.POST("/server", server.Add)      // 添加服务器配置
	e.GET("/server", server.List)      // 服务器列表配置
	e.PUT("/server", server.Update)    // 服务器配置修改
	e.DELETE("/server", server.Delete) // 服务器配置删除

	// 工单
	work := new(controller.WorkList)
	e.POST("/work", work.Add)                  // 添加工单
	e.GET("/work/list", work.List)             // 工单列表
	e.GET("/work/info", work.Info)             // 工单详情
	e.PUT("/work/update", work.Update)         // 工单修改
	e.GET("/work/person", work.WorkNum)        // 个人工单数量
	e.PUT("/work/close", work.Close)           // 关闭工单
	e.POST("/work/reply", work.Reply)          // 回复工单
	e.GET("/order/exist", work.ExistByOrderId) // 查询订单号是否存在

	// 回收站
	recovery := new(controller.Recovery)
	e.GET("/recovery/user", recovery.UserList)            // 客户回收列表
	e.GET("/recovery/order", recovery.OrderList)          // 订单回收列表
	e.PUT("/recovery/user", recovery.UserRenewOrDelete)   // 客户恢复或删除
	e.PUT("/recovery/order", recovery.OrderRenewOrDelete) // 订单恢复或删除

	// 交接班
	shift := new(controller.Shift)
	e.GET("/shift", shift.List)           // 交接班列表
	e.POST("/shift", shift.Add)           // 添加交接班
	e.POST("/shift/extend", shift.Extend) // 追加交接班描述内容
	e.GET("/shift/info", shift.Info)      // 交接班详情
	e.DELETE("/shift", shift.Del)         // 删除交接班

	// 知识库
	kb := new(controller.KnowledgeBase)
	e.GET("/knowledge/base", kb.List)                // 知识库列表
	e.POST("/knowledge/base", kb.Add)                // 新增知识库
	e.DELETE("/knowledge/base", kb.Del)              // 删除知识库
	e.PUT("/knowledge/base/praise", kb.UpdatePraise) // 知识库点赞
	e.GET("/knowledge/base/info", kb.Info)           // 知识库详情
	e.PUT("/knowledge/base", kb.Update)              // 编辑知识库
	e.GET("/kb/history/info", kb.HistoryInfo)        // 知识库历史详情
	e.GET("/kb/history/version", kb.HistoryVer)      // 知识库历史详情上一版本或下一版本
}
