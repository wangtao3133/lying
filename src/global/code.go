package global

var ErrorMsg = map[string]map[string]string{
	"app": {
		"file_sha256_not_found": "程序sha256文件不存在",
		"file_sha256_error":     "程序sha256文件内容错误",
		"file_md5_not_found":    "程序md5文件不存在",
		"file_md5_error":        "程序md5文件内容错误",
	},
	"db": {
		"error": "数据库错误",
	},
	"redis": {
		"error":     "redis错误",
		"empty_key": "redis key不存在",
	},
	"system": {
		"error":               "系统错误",
		"get_exec_path_error": "获取程序运行路径错误",
		"get_fonts_error":     "获取字体文件错误",
		"range":               "是否安装系统值不合法",
		"not_found":           "操作系统不存在",
	},
	"permission": {"denied": "你没有权限"},
	"input":      {"error_type": "传参类型错误"},
	"add":        {"error": "添加失败"},
	"update": {
		"error": "修改失败",
		"null":  "没有数据被修改",
	},
	"delete": {
		"error": "删除失败",
	},
	"base64": {
		"decode_error": "base64解密失败",
	},
	"uuid": {
		"init_error": "生成uuid字符串错误",
	},
	"json": {
		"unmarshal_error": "json解析失败",
	},
	"ldap": {
		"server_connect_error":   "ldap服务器连接失败",
		"server_connect_success": "ldap服务器连接成功",
		"non_config":             "ldap未配置",
		"create_error":           "创建ldap账号失败",
		"config_error":           "ldap配置失败",
	},
	"role": {
		"not_found":  "角色不存在",
		"id_invalid": "角色id不合法",
	},
	"login": {
		"failure_or_timeout":            "登录失败或超时",
		"incorrect_account_or_password": "账号或密码不正确",
		"log_error":                     "添加登录日志失败",
		"error_too_more":                "账号密码错误次数超过限制,请稍候再试",
	},
	"downline": {
		"error": "踢线失败",
	},
	"valid": {
		"error":   "valid验证方法错误",
		"success": "校验成功",
	},
	"captcha": {
		"init_error":               "初始化验证码对象错误",
		"non_existence_or_expired": "验证码不存在或已过期",
		"error":                    "验证码错误",
		"param_error":              "验证码参数错误",
	},
	"code": {
		"length": "code码长度错误",
	},
	"key": {
		"length": "key长度错误",
	},
	"account": {
		"match":                    "账号不合法",
		"not_found":                "账号不存在",
		"is_exist":                 "账号已存在",
		"is_refuse":                "账号被拒绝",
		"is_approve":               "账号待审核",
		"is_disable":               "账号被禁用",
		"cannot_delete_on_enable":  "账号非禁用状态,不能删除",
		"cannot_delete_on_examine": "账号已审批通过,不能删除",
		"examine_error":            "账号审批失败",
		"nullable":                 "账号长度不合法",
		"rangesize":                "账号长度不合法",
	},
	"username": {
		"rangesize": "用户名不合法",
		"maxsize":   "用户名不合法",
		"is_exist":  "客户名已存在",
		"nullable":  "客户名不合法",
	},
	"roleid": {
		"required":  "角色id必选",
		"range":     "角色id不合法",
		"not_found": "角色id不存在",
	},
	"email": {
		"email":    "无效的email地址",
		"nullable": "无效的email地址",
	},
	"wechat": {
		"error":    "无效的微信号",
		"maxsize":  "微信号长度超过限制",
		"nullable": "微信不合法",
	},
	"qq": {
		"error":    "无效的qq号",
		"nullable": "qq不合法",
		"maxsize":  "qq长度超过限制",
	},
	"mobile": {
		"mobile": "无效的手机号",
		"phone":  "无效的手机号",
	},
	"job": {
		"maxsize": "job长度超过限制",
	},
	"password": {
		"rangesize":       "密码长度不合法",
		"invalid":         "密码格式不合法",
		"nullable":        "密码长度不合法",
		"cannot_the_same": "新旧密码不能相同",
		"inequality":      "两次密码不一致",
		"error":           "密码错误",
	},
	"oldpassword": {
		"rangesize": "旧密码长度不合法",
		"error":     "原密码错误",
		"newpassword_cannot_the_same_old": "新密码不能与旧密码不同",
	},
	"newpassword": {
		"rangesize": "新密码长度不合法",
	},
	"replypassword": {
		"rangesize": "重复密码长度不合法",
		"nullable":  "重复密码长度不合法",
	},
	"repassword": {
		"rangesize": "重复密码长度不合法",
	},
	"avatar": {
		"invalid":       "头像格式不正确",
		"data_too_long": "头像大小不能超过1M",
		"type_error":    "头像类型不正确",
	},
	"style": {
		"rangesize": "style长度不合法",
		"is_exist":  "style已存在",
	},
	"price": {
		"minfloat64": "price不合法",
		"maxfloat64": "price不合法",
	},
	"remark": {
		"rangesize": "备注长度不合法",
		"maxsize":   "备注长度超过限制",
		"nullable":  "备注长度不合法",
	},
	"authtype": {
		"min": "认证方式不合法",
	},
	"id": {
		"min":                     "id不合法",
		"maxsize":                 "id不合法",
		"minsize":                 "id不合法",
		"invalid":                 "id不合法",
		"rangesize":               "id不合法",
		"not_found":               "id不存在",
		"required":                "id必填",
		"cannot_eq_zero":          "id不能为0",
		"contain_not_exist_by_id": "有id不存在",
	},
	"status": {
		"range": "状态值不合法",
	},
	"usestatus": {
		"range":                     "使用状态值不合法",
		"cannot_change_by_server":   "主机已被租用,服务器不可更改",
		"cannot_change_to_hostrent": "不能更改主机使用状态",
	},
	"examinestatus": {
		"range": "审批状态值不合法",
	},
	"examine": {
		"error": "审批失败",
	},
	"temperature": {
		"range": "温度不合法",
	},
	"title": {
		"is_exist":              "名称已存在",
		"match":                 "%s不合法",
		"chinese":               "%s必须为中文",
		"rangesize":             "%s长度不合法",
		"cannot_change_by_host": "机柜下存在主机,不能修改名称",
		"maxsize":               "%s长度超过限制",
	},
	"omit": {
		"match": "%s不合法",
	},
	"title_omit": {
		"is_exist":                 "名称和缩写已存在",
		"cannot_change_by_room":    "该区域下存在机房,不允许修改名称和缩写",
		"cannot_change_by_cabinet": "该机房下存在机柜,不允许修改名称和缩写",
	},
	"address": {
		"rangesize": "地址长度超过限制",
	},
	"area": {
		"cannot_delete_by_room":  "该区域下存在机房,不允许删除",
		"cannot_delete_by_wrong": "区域信息有误，不能删除",
		"cannot_update_by_wrong": "区域信息有误，不能更改区域",
		"old_area_not_found":     "信息有误，原区域不存在",
	},
	"area_room": {
		"cannot_change_by_device": "该机柜下存在设备,不允许修改区域或机房",
		"id_error":                "区域或机房id不合法",
	},
	"room": {
		"cannot_up_area_by_cabinet": "该机房下存在机柜,不允许修改区域",
		"cannot_delete_by_cabinet":  "该机房下存在机柜,不允许删除",
		"cannot_delete_by_wrong":    "机房信息有误，不能删除",
		"cannot_update_by_wrong":    "机房信息有误，不能更改机房",
		"old_area_not_found":        "信息有误，原机房不存在",
	},
	"cabinet": {
		"cannot_delete_by_device": "该机柜下存在设备,不允许删除",
		"off_capacity":            "机柜容量不足",
		"location_is_wrong":       "请求位置信息有误",
		"cannot_update":           "不允许更改机柜",
		"cannot_delete_by_wrong":  "机柜信息有误，不能删除",
		"cannot_update_by_wrong":  "机柜信息有误，不能更改机柜",
		"old_area_not_found":      "信息有误，原机柜不存在",
	},
	"upload": {
		"not_file":      "上传文件不存在",
		"is_exist":      "文件已存在",
		"type_error":    "文件类型不合法",
		"data_too_long": "文件大小超过限制",
		"error":         "上传失败",
	},
	"type": {
		"range":   "类型不合法",
		"invalid": "类型不合法",
	},
	"deviceid": {
		"min":       "设备id不合法",
		"invalid":   "设备id不合法",
		"not_found": "设备id不存在",
		"maxsize":   "设备id不合法",
	},
	"areaid": {
		"required":  "区域id必填",
		"min":       "区域id不合法",
		"invalid":   "区域id不合法",
		"not_found": "区域id不存在",
	},
	"roomid": {
		"required":  "机房id必填",
		"min":       "机房id不合法",
		"not_found": "机房id不存在",
	},
	"cabinetid": {
		"required":  "机柜id必填",
		"min":       "机柜id不合法",
		"not_found": "机柜id不存在",
	},
	"linechildid": {
		"required": "子线路id必填",
	},
	"earlywarning": {
		"range": "预警温度不合法",
	},
	"sizelong": {
		"range": "长度超过限制",
	},
	"sizewith": {
		"range": "宽度超过限制",
	},
	"sizehigh": {
		"range": "高度超过限制",
	},
	"electric": {
		"range": "电量超过限制",
	},
	"socket": {
		"range": "插排超过限制",
	},
	"flowsupport": {
		"range": "流量报告支持不合法",
	},
	"clapboard": {
		"range": "隔板不合法",
	},
	"devicesort": {
		"status_type_deviceid_invalid": "状态、类型、设备id不能为0",
		"data_error":                   "请求数据与机柜容量不相符",
		"device_error":                 "机柜设备信息有误",
	},
	"device": {
		"not_found":             "设备不存在",
		"cannot_delete_by_port": "该设备已有端口被连接,不能删除",
		"cannot_move_by_order":  "该设备有业务,不能移交",
		"move_error":            "设备移交失败",
	},
	"brand": {
		"rangesize": "品牌长度超过限制",
	},
	"cpu": {
		"rangesize": "cpu不合法",
	},
	"ip": {
		"range":                "ip个数不合法",
		"is_exist":             "ip已存在",
		"invalid":              "ip不合法",
		"ip":                   "ip不合法",
		"not_found":            "ip不存在",
		"host_line_no_cabinet": "主机和线路不在同一机柜下",
		"line_ip_not_found":    "线路下的ip不存在",
	},
	"line": {
		"not_found": "线路不存在",
	},
	"mac": {
		"mac":      "mac地址不合法",
		"is_exist": "mac地址已存在",
	},
	"memory": {
		"range": "内存你不合法",
	},
	"harddisk": {
		"rangesize": "硬盘不合法",
		"not_found": "硬盘不存在",
	},
	"bandwidth": {
		"range": "带宽不合法",
	},
	"problem": {
		"range": "问题不合法",
	},
	"orderid": {
		"min":     "订单id不合法",
		"invalid": "订单id不合法",
	},
	"describe": {
		"maxsize": "描述不合法",
	},
	"creatorid": {
		"min": "创建者id不合法",
	},
	"level": {
		"min":     "级别不合法",
		"range":   "等级超过限制",
		"invalid": "等级不合法",
	},
	"complete": {
		"max": "完成度不合法",
	},
	"time": {
		"start_time_cannot_gt_end_time":         "开始时间不能大于结束时间",
		"enter_right_time":                      "请输入正确的时间",
		"get_time_zone_fail":                    "获取时区失败",
		"unlawful_time_conversion":              "时间转换不合法",
		"year_month_cannot_get_nowtime":         "年月不能大于当前时间",
		"year_month_cannot_than_oneyear":        "年月不能提前一年以上时间",
		"unlawful_time":                         "时间不合法",
		"time_must_be_selected_or_not_selected": "开始时间和结束时间必须都选择或者都不选择",
		"time_must_after_today":                 "到期时间必须在今天之后",
	},
	"starttime": {
		"minsize":           "开始时间不合法",
		"cannot_lg_endtime": "开始时间不能大于结束时间",
		"cannot_eq_endtime": "开始时间不能等于结束时间",
	},
	"endtime": {
		"minsize": "结束时间不合法",
	},
	"name": {
		"rangesize": "名称不合法",
		"nullable":  "名称不合法",
		"maxsize":   "名称长度超过限制",
	},
	"skype": {
		"nullable": "skype不合法",
		"maxsize":  "skype长度超过限制",
	},
	"salesid": {
		"not_found": "销售人员id不存在",
	},
	"image": {
		"error":            "图片格式不正确",
		"limit":            "头像大小不能超过1M",
		"type":             "头像类型不正确",
		"exceed_the_limit": "最多只能添加五张图片",
	},
	"assign": {
		"range":     "指派类型不合法",
		"required":  "指派者信息必填",
		"not_exist": "指派者们id有误",
	},
	"workid": {
		"min": "工单id不合法",
	},
	"content": {
		"rangesize": "回复长度不合法",
		"maxsize":   "搜索内容长度超过限制",
		"invalid":   "搜索内容不合法",
	},
	"work": {
		"not_allow_close":      "不允许关闭",
		"not_allow_update":     "不允许修改",
		"cannot_close_by_role": "非超管、管理员、创建者不能关闭工单",
	},
	"site": {
		"rangesize":  "站点不合法",
		"site_exist": "站点已存在",
	},
	"spread": {
		"quantity_ceiling":        "数量达到上限，不能再添加",
		"not_found":               "推广不存在",
		"not_found_by_spreadcode": "推广码不存在",
	},
	"rent": {
		"minfloat64": "月租不合法",
		"maxfloat64": "月租不合法",
	},
	"ddos": {
		"range": "DDOS流量防御不合法",
	},
	"cc": {
		"range": "CC防御不合法",
	},
	"sourceipnum": {
		"range": "源ip个数不合法",
	},
	"domainnum": {
		"range": "域名限制个数不合法",
	},
	"ddoscc": {
		"range": "DDOS+CC防御不合法",
	},
	"userconnect": {
		"range": "同户同时连接数不合法",
	},
	"sourceconnect": {
		"rangesize": "源站连接方式不合法",
	},
	"downtime": {
		"range": "宕机自动切换不合法",
	},
	"order": {
		"range":                 "是否有订单值不合法",
		"cost":                  "原价不合法",
		"discount":              "折扣不合法",
		"discount_price":        "优惠金额不合法",
		"server_conf_not_exist": "服务器配置不存在",
		"cs_not_exist":          "客服不存在",
		"ts_not_exist":          "技术支持不存在",
		"duty_not_exist":        "产品负责人不存在",
		"user_not_exist":        "客户不存在",
		"site":                  "站点不合法",
		"cdn_type":              "套餐不合法",
		"cdn_not_exist":         "cdn套餐不存在",
		"stack_deficiency":      "服务器库存不足",
		"not_found":             "订单不存在",
		"unable_to_delete":      "订单不存在或未过期,无法删除",
		"renew_error":           "续费失败",
		"refund_error":          "退款失败",
		"feedback_error":        "修改订单过期原因失败",
		"delay_error":           "修改订单延期失败",
		"update_normal_error":   "修改状态为进行中失败",
		"update_timeout_error":  "修改状态为已过期失败",
		"update_ready_error":    "修改状态为待续费失败",
		"renew_endtime_error":   "续费订单的结束时间不能小于进行中订单的结束时间",
		"refund_money_error":    "退款金额不能大于实付金额",
		"price_money_error":     "优惠金额不能大于原价",
		"cannot_refund":         "有续费订单,无法退款",
	},
	"feedback": {
		"rangesize": "原因内容不合法",
	},
	"refundmoney": {
		"minfloat64": "退款金额不合法",
	},
	"refundremark": {
		"rangesize": "退款原因不合法",
	},
	"servertype": {
		"range": "服务类型不合法",
	},
	"ifuse": {
		"range": "是否试用不合法",
	},
	"csid": {
		"min": "客服不合法",
	},
	"tsid": {
		"min": "技术支持不合法",
	},
	"userid": {
		"min": "客户不合法",
	},
	"dutyid": {
		"min":      "产品负责人不合法",
		"required": "客服id不合法",
	},
	"cdntype": {
		"range": "cdn套餐类型不合法",
	},
	"ordertype": {
		"range": "订单状态不合法",
	},
	"year": {
		"min": "年份不合法",
	},
	"month": {
		"min":   "月份不合法",
		"range": "月份不合法",
	},
	"day": {
		"min": "天数不合法",
	},
	"templateid": {
		"not_found": "模板id不存在",
	},
	"sn": {
		"maxsize": "sn长度超过限制",
	},
	"authorize": {
		"maxsize": "授权码长度超过限制",
	},
	"capacity": {
		"range": "容量超过限制",
	},
	"manageip": {
		"ip": "管理id不合法",
	},
	"enable": {
		"maxsize": "enable长度超过限制",
	},
	"seepassword": {
		"rangesize": "查看密码长度超过限制",
		"error":     "查看密码错误",
		"nullable":  "查看密码不合法",
	},
	"model": {
		"rangesize": "型号长度超过限制",
	},
	"systemversion": {
		"max": "系统版本不合法",
	},
	"kernelversion": {
		"maxsize": "系统内核版本长度超过限制",
	},
	"cpumodel": {
		"rangesize": "cpu型号长度超过限制",
		"maxsize":   "cpu型号长度超过限制",
	},
	"cpunum": {
		"range": "cpu个数超过限制",
	},
	"cpucore": {
		"range": "cpu核心数超过限制",
	},
	"cpuline": {
		"range": "cpu线程数超过限制",
	},
	"cpufreq": {
		"minfloat64": "cpu主频不合法",
		"maxfloat64": "cpu主频不合法",
	},
	"cpugl": {
		"range": "cpu取值方式不合法",
	},
	"memorygl": {
		"range": "内存取值方式不合法",
	},
	"hdgl": {
		"range": "硬盘取值方式不合法",
	},
	"memorytype": {
		"range": "内存类型不合法",
	},
	"memorysize": {
		"max": "内存大小超过限制",
	},
	"raidcardmodel": {
		"rangesize": "raid卡型号超过限制",
	},
	"raidcardlevel": {
		"rangesize": "raid卡等级不合法",
		"invalid":   "raid卡等级不合法",
	},
	"ipmiip": {
		"ip":       "ipmi ip不合法",
		"is_exist": "ipmi ip已存在",
	},
	"ipmimac": {
		"mac":      "ipmi mac不合法",
		"is_exist": "ipmi mac已存在",
	},
	"ipmiaccount": {
		"rangesize": "ipmi账号长度超过限制",
	},
	"ipmiuser": {
		"nullable": "ipmi账号长度超过限制",
	},
	"ipmipassword": {
		"rangesize": "ipmi密码长度超过限制",
		"nullable":  "ipmi密码长度超过限制",
	},
	"server": {
		"cannot_delete_by_stack": "服务器还有库存,不能删除",
		"cannot_delete_by_host":  "服务器下有主机绑定，不能删除",
		"cannot_delete_by_order": "服务器下有订单绑定，不能删除",
		"host_maybe_use":         "主机可能已经被使用，不能修改其服务器配置或使用状态",
	},
	"serverid": {
		"min":           "服务器配置id不合法",
		"not_found":     "服务器配置不存在",
		"old_not_found": "主机原本的服务器配置不存在",
	},
	"hostid": {
		"invalid":   "主机id不合法",
		"min":       "主机id不合法",
		"not_found": "主机id不存在",
	},
	"host": {
		"not_found":                "主机不存在",
		"cannot_delete_by_order":   "主机存在订单,不能删除",
		"cannot_delete_by_network": "主机网卡已被绑定,不能删除",
		"cannot_delete_by_ip":      "主机存在ip业务,不能删除",
		"cannot_delete_by_ipmiip":  "主机ipmi远程卡端口已被绑定,不能删除",
		"status_error":             "主机状态不正确",
		"close_error":              "主机关闭失败",
		"open_error":               "主机开启失败",
		"rangesize":                "主机不合法",
		"no_cabinet":               "主机没有查到机柜",
		"partial_exist":            "主机不全存在",
	},
	"issystem": {
		"range": "是否为详细搜索值不合法",
	},
	"startip": {
		"ip":                 "开始ip不合法",
		"maxsize":            "开始ip不合法",
		"not_empty_by_endip": "结束ip不为空,开始ip必填",
		"is_exist":           "开始ip已存在",
	},
	"endip": {
		"ip":                   "结束ip不合法",
		"maxsize":              "结束ip不合法",
		"not_empty_by_startip": "开始ip不为空,结束ip必填",
		"is_exist":             "结束ip已存在",
	},
	"startip_endip": {
		"not_the_same":    "开始ip和结束ip不能相同",
		"not_the_segment": "开始ip和结束ip不在同一ip段",
	},
	"hdtype": {
		"range": "硬盘类型不合法",
	},
	"hdcapacity": {
		"min": "硬盘容量不合法",
	},
	"rate": {
		"range": "接口速率超过限制",
	},
	"ismonitor": {
		"range": "监控开关不合法",
	},
	"port": {
		"range":            "端口值不合法",
		"invalid":          "端口不合法",
		"is_bind":          "端口已被绑定",
		"cannot_bind_host": "该端口不能绑定主机",
		"cannot_bind_self": "该端口不能绑定自身",
		"min":              "端口不合法",
		"cannot_delete_by_wrong": "使用端口信息有误，不能删除",
		"old_area_not_found":     "信息有误，原连接端口不存在",
	},
	"selfid": {
		"min": "设备自身id不合法",
	},
	"selftype": {
		"range": "设备自身类型不合法",
	},
	"portid": {
		"min":       "端口id不合法",
		"not_found": "端口不存在",
	},
	"baseda": {
		"rangesize": "base da长度超过限制",
	},
	"interbandwidth": {
		"range": "国际带宽超过限制",
	},
	"chinabandwidth": {
		"range": "国内带宽超过限制",
	},
	"localbandwidth": {
		"range": "本地带宽超过限制",
	},
	"networkaddress": {
		"ip":       "网络地址不合法",
		"is_exist": "网络地址已存在",
	},
	"subnetmask": {
		"ip": "子网掩码不合法",
	},
	"gateway": {
		"ip":       "网关不合法",
		"is_exist": "网关已存在",
	},
	"lineid": {
		"min":       "线路id不合法",
		"not_found": "线路id不存在",
	},
	"route": {
		"rangesize": "路由长度超过限制",
	},
	"routetype": {
		"range": "路由方式不合法",
	},
	"icon": {
		"rangesize": "icon长度超过限制",
	},
	"parentid": {
		"min":       "所属父级id不合法",
		"not_found": "所属父级id不存在",
	},
	"menu": {
		"cannot_delete_by_child": "该菜单下有子菜单,禁止删除",
	},
	"isself": {
		"range": "是否自有值不合法",
	},
	"defensevalue": {
		"range": "节点防御值超过限制",
	},
	"defensetype": {
		"range": "防御类型超过限制",
	},
	"groupid": {
		"min": "分组区域id不合法",
	},
	"ipmi": {
		"cannot_empty_by_noself": "非自用节点必做填写完整ipmi数据",
		"is_bind":                "ipmi已被绑定",
	},
	"nodegroup": {
		"not_found":             "节点分组区域不存在",
		"cannot_delete_by_node": "该分组区域下有节点,不能删除",
	},
	"operationaccount": {
		"nullable": "操作人员账号不合法",
	},
	"export": {
		"max": "是否导出不合法",
	},
	"pid": {
		"min":       "产品id不合法",
		"not_found": "产品id不存在",
	},
	"pport": {
		"min":     "自身端口不合法",
		"invalid": "自身端口不合法",
	},
	"cport": {
		"invalid": "端口不合法",
	},
	"cid": {
		"rangesize": "端口连接id不合法",
		"invalid":   "端口连接id不合法",
	},
	"ptype": {
		"range":   "产品类型不合法",
		"invalid": "产品类型不合法",
	},
	"ctype": {
		"range": "链接的产品类型不合法",
	},
	"ntype": {
		"range": "接口类型不合法",
	},
	"switchboard": {
		"only_link_host_or_switchboard": "交换机普通端口只能连接主机或者交换机",
		"only_link_route_or_firewall":   "交换机sfp端口只能连接路由器或者防火墙",
	},
	"network": {
		"not_found": "网卡不存在",
		"is_bind":   "网卡已被绑定",
	},
	"routeid": {
		"not_found":                "路由器不存在",
		"not_allow_update_cabinet": "不允许修改所属机柜",
		"not_allow_delete":         "不允许删除",
	},
	"switchboardid": {
		"not_found":                "交换机不存在",
		"not_allow_update_cabinet": "不允许修改所属机柜",
		"not_allow_delete":         "不允许删除",
	},
	"modelname": {
		"rangesize": "模块名长度超过限制",
	},
	"modelname_title": {
		"is_exist": "模块名和权限名已存在",
	},
	"istop": {
		"range": "是否置顶不合法",
		"error": "置顶失败",
	},
	"leaderid": {
		"not_found": "组长id不存在",
	},
	"termid": {
		"not_found": "组员不存在",
		"min":       "维护组id不合法",
	},
	"term": {
		"not_found":                          "维护组不存在",
		"cannot_delete_by_order":             "该维护组存在业务,不能解散",
		"bind_error":                         "维护组绑定失败",
		"device_cannot_delete_by_term_wrong": "设备维护组信息有误不能删除",
		"not_found_by_device":                "设备绑定的维护组信息不存在",
	},
	"devicetype": {
		"range":   "设备类型不合法",
		"invalid": "设备类型不合法",
	},
	"addnotice": {
		"range": "新增设备是否通知不合法",
	},
	"changenotice": {
		"range": "设备变更是否通知不合法",
	},
	"leadercheck": {
		"range": "设备迁入是否需要组长审核不合法",
	},
	"tag": {
		"too_more": "标签个数超过限制",
		"invalid":  "标签不合法",
	},
	"mold": {
		"range": "设备类型不合法",
	},
	"groupparse": {
		"error":        "分组解析不合法",
		"id_not_found": "分组解析id不存在",
	},
	"contact": {
		"error": "客户联系方式不可全部为空",
	},
	"duty": {
		"error":        "客服不存在",
		"obtain_error": "获取客服名失败",
	},
	"phone": {
		"maxsize":  "联系电话不合法",
		"is_exist": "联系电话已存在",
	},
	"user": {
		"error":             "新增客户失败",
		"get_order_failure": "获取客户订单详情失败",
		"get_failure":       "获取客户详情失败",
		"delete_failure":    "客户有订单,无法删除",
	},
	"qqname": {
		"nullable": "qq昵称不合法",
	},
	"nodegroupparse": {
		"cannot_delete_by_high":       "该分组区域下有高防套餐,不能删除",
		"cannot_delete_by_domain":     "该分组区域下有无限域名套餐,不能删除",
		"cannot_delete_by_connection": "该分组区域下有无限域名套餐,不能删除",
	},
	"business": {
		"not_found":                  "业务不存在",
		"ip":                         "源ip填写不全",
		"domain":                     "域名填写不全",
		"cdn_order_not_found":        "cdn订单不存在",
		"server_order_not_found":     "服务器订单不存在",
		"ip_order_not_found":         "ip订单不全存在",
		"ip_cannot_gt_rent_num":      "ip数据个数不能大于ip订单租用个数",
		"host_cannot_gt_rent_num":    "主机个数不能大于服务器订单租用个数",
		"ip_partial_exist":           "ip不全存在",
		"ip_conf_host_partial_exist": "ip配置中主机不全存在",
		"exist_to_server_order":      "服务器订单已存在",
		"exist_to_cdn_order":         "cdn订单已存在",
		"exist_to_ip_order":          "ip订单已存在",
		"old_order_cannot_update":    "原ip订单不能修改",
		"ip_conf_cannot_repeat":      "ip配置不能有重复",
	},
	"serverorderid": {
		"min": "服务器订单id不合法",
	},
	"cdnOrderId": {
		"min": "cdn订单id不合法",
	},
	"domainParse": {
		"rangesize": "域名解析不合法",
	},
	"cdn": {
		"has_order": "该套餐有订单绑定，不能删除",
	},
	"spreadcode": {
		"length": "推广码长度不合法",
	},
	"adminid": {
		"min": "推广人id不合法",
	},
	"notice": {
		"has_not_belong_notice": "有不属于登陆人的消息",
	},
	"field": {
		"rangesize": "所属领域不合法",
	},
	"article": {
		"required": "知识库文章不合法",
	},
	"publishtime": {
		"min": "发布时间不合法",
	},
	"knowledgebase": {
		"not_found":           "知识库不存在",
		"update_praise_error": "点赞失败",
		"already_praise":      "你已经点赞,不可重复操作",
		"admin_id_error":      "owner不合法",
	},
	"supplierid": {
		"min": "供应商id不合法",
		"cannot_delete_on_realm": "有域名选择,不能删除",
		"account_is_supplier":    "选择账号不属于所选供应商",
	},
	"shift": {
		"range":                   "%s不合法",
		"id_invalid_or_not_owner": "交接班id不合法或者不属于自己添加的交接班",
	},
	"shiftid": {
		"min": "%s不合法",
	},
	"service": {
		"min":              "客服id不合法",
		"numbers_too_more": "提醒客服个数超过限制",
		"invalid":          "所选客服不合法",
	},
	"attachment": {
		"numbers_too_more": "图片附件个数超过限制",
	},
	"accountid": {
		"min": "账号id不合法",
	},
	"becomeday": {
		"min": "到期提醒天数不合法",
	},
}

// 返回码总体说明:
// ret = 0: 正确返回
// ret > 0: 请求错误
// ret < 0: 系统内部错误
var ReturnCode = map[int]string{
	0:   "操作成功",
	-1:  "数据库错误",
	-2:  "redis错误",
	-3:  "系统错误",
	-4:  "ldap服务器连接失败",
	-5:  "角色不存在",
	-7:  "登录失败或超时",
	-8:  "没有数据被修改",
	-9:  "你没有权限",
	-10: "code不存在",
	1:   "传参类型错误",
	10:  "添加失败",
	11:  "修改失败",
	12:  "删除失败",
	13:  "绑定失败",
	14:  "审批失败",
	15:  "更新失败",
	16:  "移交失败",
	// 账号
	100: "账号格式不正确",
	101: "账号被禁用",
	102: "账号未通过审核",
	103: "账号已存在",
	104: "ldap账号同步失败",
	105: "账号添加失败",
	106: "账号修改失败",
	107: "账号删除失败",
	108: "账号非禁用状态,不能删除",
	109: "账号已通过审批,不能删除",
	110: "账号不存在",
	111: "账号审批失败",
	112: "审批的账号id不正确",
	113: "账号或密码不正确",
	114: "账号注册已被拒绝",
	115: "账号注册失败",
	116: "账号待审核",
	117: "账号密码错误次数超过限制,请稍候再试",
	118: "账号已存在",
	// 密码
	120: "密码格式不正确",
	121: "重复密码格式不正确",
	122: "重复密码不一致",
	123: "原密码不正确",
	124: "新密码和原密码不能相同",
	125: "原密码格式不正确",
	126: "新密码格式不正确",

	// 验证码
	130: "验证码获取次数超过限制,请稍候再试",
	131: "验证码参数不正确",
	132: "验证码key不正确",
	133: "验证码不存在或已过期",
	134: "验证码不正确",
	// 头像
	140: "头像格式不正确",
	141: "头像大小不能超过1M",
	142: "获取头像内容失败",
	143: "头像类型不正确",
	// 真实姓名
	150: "姓名格式不正确",
	// 手机号
	151: "手机号格式不正确",
	// 邮箱
	152: "邮箱格式不正确",
	// 职位
	153: "职位格式不正确",
	// 备注
	154: "备注格式不正确",
	// 状态格式不正确
	155: "状态格式不正确",
	156: "认证方式不正确",
	157: "审核状态格式不正确",
	// ldap
	160: "ldap主机格式不正确",
	161: "ldap端口格式不正确",
	162: "ldap名称格式不正确",
	163: "ldap da格式不正确",
	164: "ldap未配置",
	165: "ldap配置失败",
	166: "ldap服务器连接成功",
	// id
	169: "id不能为0",
	170: "id不合法",
	171: "id不存在",
	172: "区域id不合法",
	173: "区域不存在",
	174: "销售人员id有误",
	175: "机房id不合法",
	176: "机房不存在",
	177: "位置信息有误",
	178: "机柜id不合法",
	179: "机柜不存在",

	// title
	180: "名称不合法",
	181: "名称已存在",
	182: "名称和缩写已存在",

	// omit
	190: "缩写不合法",
	191: "不允许修改缩写",
	192: "该区域下存在机房,不允许修改名称和缩写",
	193: "该区域下存在机房,不允许删除",
	194: "该机柜下存在设备,不允许修改区域或机房",

	// address
	200: "详细地址不合法",

	// skype
	210: "skype不合法",

	// wechat
	211: "微信不合法",

	// qq
	212: "QQ不合法",

	// 不允许
	220: "不允许更改区域",
	221: "不允许删除",
	222: "不允许修改名称",
	223: "不允许更改机柜",

	// ip
	230: "ip不合法",

	// 预警温度
	231: "预警温度不合法",

	// 机柜尺寸
	232: "长度不合法",
	233: "宽度不合法",
	234: "高度不合法",

	// electric
	235: "电量不合法",

	// socket
	236: "排插不合法",

	// flowSupport
	237: "流量报告不合法",

	// clapBoard
	238: "隔板不合法",

	// 带宽  inter_bandwidth ,china_bandwidth,local_bandwidth
	239: "国际带宽不合法",
	240: "国内带宽不合法",
	241: "本地带宽不合法",

	// ip
	242: "网络地址不合法",
	243: "子网掩码不合法",
	244: "网关不合法",
	245: "开始ip不合法",
	246: "结束ip不合法",
	247: "网络地址已存在",
	248: "网关已存在",
	249: "开始ip已存在",
	250: "结束ip已存在",

	// type
	251: "类型不合法",

	// status
	252: "状态不合法",

	// use_status
	253: "使用状态不合法",

	// temperature
	254: "温度不合法",
	255: "状态、类型、设备id不能为0",
	256: "请求数据与机柜容量不相符",

	// enable密码
	257: "enable密码不合法",

	// template_id
	258: "模板id不合法",
	259: "模板id不存在",

	// 查询类型
	260: "查询类型不合法",
	261: "ip已存在",

	// 备份上传
	262: "文件类型不合法",
	263: "文件大小不能超过3m",
	264: "产品类型不合法",
	265: "主机没有上传备份功能",
	266: "备份id不存在",
	267: "设备id不合法",

	// 端口绑定
	268: "接口不合法",
	269: "接口已被绑定",
	270: "接口连接id不合法",
	271: "不能绑定主机接口",
	272: "产品端口不能跟自身接口绑定",
	273: "网卡已被绑定",
	274: "主机ipmi已被绑定",
	275: "连接设备不存在",
	276: "连接接口不合法",
	277: "连接设备类型不合法",
	278: "接口已绑定",
	279: "接口类型不合法",
	280: "接口不存在",

	// 备份
	281: "备份不存在",
	282: "有设备不存在",
	283: "交换机SFP接口只能与路由器和防火墙接口连接",
	284: "交换机普通端口只能连接主机或者交换机",
	285: "机柜设备信息有误",

	// 主机管理
	300: "ipmi_ip已存在",
	301: "ipmi_mac已存在",
	302: "硬盘添加失败",
	303: "主机不存在",
	304: "消息添加失败",
	305: "硬盘id不合法",
	306: "品牌不合法",
	307: "型号不合法",
	308: "授权码不合法",
	309: "sn不合法",
	310: "主机备注不合法",
	311: "系统不合法",
	312: "版本号不合法",
	313: "内核版本不合法",
	314: "cpu型号不合法",
	315: "cpu个数不合法",
	316: "cpu核心不合法",
	317: "cpu线程不合法",
	318: "cpu主频不合法",
	319: "内存类型不合法",
	320: "内存大小不合法",
	321: "raid卡型号不合法",
	322: "raid卡等级不合法",
	323: "ipmi ip不合法",
	324: "ipmi物理地址不合法",
	325: "ipmi账号不合法",
	326: "ipmi密码不合法",
	327: "主机用户不合法",
	328: "主机用户密码不合法",
	329: "租金不合法",
	330: "主机id不合法",
	331: "是否自用不合法",
	332: "容量不合法",
	333: "硬盘类型不合法",
	334: "硬盘容量不合法",
	335: "主机添加失败",
	336: "添加日志失败",
	337: "订单id不合法",
	338: "开始ip不为空，结束ip必填",
	339: "结束ip不为空，开始ip必填",
	340: "开始ip不能和结束ip相同",
	341: "IP段不一致",
	342: "主机存在订单，不能删除",
	343: "网卡被绑定，不能删除",
	344: "主机存在绑定ip业务，不能删除",
	345: "主机ipmi被绑定，不能删除",
	346: "主机删除失败",
	347: "关闭主机失败",
	348: "开启主机失败",
	349: "硬盘不存在",
	350: "硬盘删除失败",
	351: "网卡ip已存在",
	352: "网卡mac已存在",
	353: "网卡添加失败",
	354: "修改网卡失败",
	355: "查看密码不合法",
	356: "监控开关不合法",
	357: "查看密码错误",
	358: "机柜容量不足",
	359: "主机修改失败",
	360: "主机状态不正确",
	361: "网卡不存在",
	362: "网卡删除失败",
	363: "系统是否安装参数不合法",
	364: "是否为详细搜索不合法",
	365: "是否用业务参数不合法",
	366: "搜索类型不合法",
	367: "cpu核心取值范围方式不合法",
	368: "内存取值范围方式不合法",
	369: "硬盘取值范围方式不合法",
	370: "开始IP不合法",
	371: "结束IP不合法",
	372: "主机状态不合法",
	373: "网卡名称不合法",
	374: "接口速率不合法",
	375: "物理地址不合法",
	376: "搜索内容不合法",
	377: "修改硬盘失败",
	378: "租用状态不允许修改",

	// 节点管理
	400: "节点防御值不合法",
	401: "防御类型不合法",
	402: "分组区域id不合法",
	403: "分组区域名称不合法",
	404: "分组区域名称已存在",
	405: "添加分组区域失败",
	406: "节点分组区域不存在",
	407: "存在节点的分组区域无法删除",
	408: "删除分组区域失败",
	409: "解析分组名称不合法",
	410: "解析分组不存在",
	411: "删除分组解析失败",
	412: "添加分组解析失败",
	413: "分组解析名称已存在",
	414: "修改分组解析失败",
	415: "是否自有不合法",
	416: "远程卡数据不完整",
	417: "节点名称已存在",
	418: "网卡ip不存在",
	419: "节点id不合法",
	420: "节点不存在",
	421: "删除节点失败",
	422: "添加节点失败",
	423: "节点ip已存在",
	424: "操作系统不存在",

	// 维护组管理
	440: "维护组名称不合法",
	441: "维护组id不合法",
	442: "维护组是否置顶不合法",
	443: "置顶维护组失败",
	444: "组名不合法",
	445: "组长不合法",
	446: "组名已存在",
	447: "新增维护组失败",
	448: "组长不存在",
	449: "组员不全存在",
	450: "维护组不存在",
	451: "备注不合法",
	452: "新增设备通知不合法",
	453: "设备变更通知不合法",
	454: "设备迁入需要组长审核不合法",
	455: "是否解散不合法",
	456: "设备类型不合法",
	457: "添加组动态失败",
	458: "标签字数不合法",
	459: "标签最多添加三条",
	460: "维护组有业务,解散失败",
	461: "解散失败",
	462: "设备有绑定业务，不能移交",

	// memory
	470: "内存不合法",

	// Bandwidth
	471: "带宽不合法",

	// price
	472: "价格不合法",

	// severId
	473: "服务器配置不合法",
	474: "服务器配置不存在",
	475: "不允许更改服务器配置",

	// CDN套餐
	480: "套餐名不合法",
	481: "月租不合法",
	482: "DDOS流量防御不合法",
	483: "CC防御不合法",
	484: "源IP个数不合法",
	485: "域名限制个数不合法",
	486: "DDOS+CC防御不合法",
	487: "同户同时连接数不合法",
	488: "源站连接方式奴合法",
	499: "套餐名已存在",
}
