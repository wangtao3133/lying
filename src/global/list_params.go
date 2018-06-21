package global

import (
	"github.com/go-xorm/xorm"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
)

// 列表参数
type ListParams struct {
	Page     int    `query:"page"`     // 页码数
	PageSize int    `query:"pageSize"` // 页面数量
	Limit    []int  // 分页
	OrderBy  string `query:"orderBy"` // 排序字段
	Desc     bool   `query:"desc"`    // 排序顺序 true 正 false 反
}

func (lp *ListParams) Make(model *xorm.Session) {
	if lp.Limit != nil {
		model.Limit(lp.Limit[0], lp.Limit[1])
	}
	if lp.OrderBy != "" {
		if lp.Desc == true {
			model.Desc(lp.OrderBy)
		} else {
			model.Asc(lp.OrderBy)
		}
	}
}

type Times struct {
	StartTime int64 `query:"startTime"` // 开始时间
	EndTime   int64 `query:"endTime"`   // 结束时间
}

// 根据时间查询
func (this *Times) Make(timeParam string, model *xorm.Session) {
	if this.StartTime > 0 && this.EndTime > 0 {
		model.Where(timeParam+">=?", this.StartTime).And(timeParam+"<=?", this.EndTime)
	} else if this.StartTime > 0 && this.EndTime == 0 {
		model.Where(timeParam+">=?", this.StartTime)
	} else if this.StartTime == 0 && this.EndTime > 0 {
		model.Where(timeParam+"<=?", this.EndTime)
	}
}

// 文件二进制流读取
func ReadByte(src *multipart.FileHeader) ([]byte, error) {
	fi, err := src.Open()
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)

	}
	return chunks, err
}

// int64数组去重
func RemoveRepeatNum(list []int64) []int64 {
	var x []int64
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}

// 字符串数组去重,去空
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	lenA := len(a)
	for i := 0; i < lenA; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

// id补零
func GetDeviceId(id int, ao, ro, an string) string {
	idStr := strconv.Itoa(id)
	var a string
	if len(idStr) < 6 {
		a = ao + "-" + ro + "-" + an + "-" + strings.Repeat("0", 6-len(idStr)) + idStr
	} else {
		a = ao + "-" + ro + "-" + an + "-" + idStr
	}
	return a
}

// id转int
func IdToInt(id string) (string, int) {
	// 分割id
	ar := strings.Split(id, "-")
	// 如果分割得到的数组不等于4
	if len(ar) != 4 {
		return "id.invalid", 0
	}
	// 将id转为int
	rid, err := strconv.Atoi(ar[3])
	if err != nil {
		Glogger.Error(err.Error())
		return "system.error", 0
	}
	return "valid.success", rid
}

// 消息提示 -- 主机温度预警通知
func Prompt(id, temperature string) string {
	x := "主机【" + id + "】温度超过【" + temperature + "】，请尽快处理。"
	return x
}

// 消息提示 -- 维护组绑定 -- 审核或者通知添加
func MaintainGroupBand(num, name string, ty, check int) string {
	var s, x string
	switch ty {
	case ProductHost:
		s = "主机"
	case ProductRouter:
		s = "路由器"
	case ProductSwitchboard:
		s = "交换机"
	case ProductFirewall:
		s = "防火墙"
	case ProductNode:
		s = "节点"
	default:
		s = "设备"
	}
	if check == IsCheck {
		x = "有 " + num + " 台" + s + "被分配到【" + name + "】，请及时审核。"
	} else {
		x = "有 " + num + " 台" + s + "被分配到【" + name + "】，请关注。"
	}
	return x
}

// 消息提示 -- 维护组 -- 设备移交
func MaintainGroupRemove(num string) string {
	x := "有 " + num + " 台设备被移交，请关注。"
	return x
}

// 消息提示 -- 维护组修改
func MaintainGroupUpdate(termName string, add, leader int) string {
	var x string
	if add == IsAdd {
		if leader == IsLeader {
			x = "你被设置为维护组【" + termName + "】组长。"
		} else {
			x = "你被添加到维护组【" + termName + "】。"
		}
	} else {
		x = "你被移出维护组【" + termName + "】。"
	}
	return x
}

// 消息提示 -- 维护组解散
func MaintainGroupDisband(name string) string {
	x := "你所在的维护组【" + name + " 】已解散。"
	return x
}

// 消息提示 -- 维护组设备修改
func MaintainGroupDeviceUpdate(name, id string, ty int) string {
	var s string
	switch ty {
	case ProductHost:
		s = "主机"
	case ProductRouter:
		s = "路由器"
	case ProductSwitchboard:
		s = "交换机"
	case ProductFirewall:
		s = "防火墙"
	case ProductNode:
		s = "节点"
	default:
		s = "设备"
	}
	x := "【" + name + "】的【" + s + "】【" + id + "】被修改了配置，请关注。"
	return x
}

// 消息提示 -- 维护组设备修改
func MaintainGroupDeviceDelete(name, id string, ty int) string {
	var s string
	switch ty {
	case ProductHost:
		s = "主机"
	case ProductRouter:
		s = "路由器"
	case ProductSwitchboard:
		s = "交换机"
	case ProductFirewall:
		s = "防火墙"
	case ProductNode:
		s = "节点"
	default:
		s = "设备"
	}
	x := "【" + name + "】的【" + s + "】【" + id + "】被删除。"
	return x
}

// 消息提示 -- 工单添加
func WorkAdd(name string) string {
	x := "你被指派了一个新工单【" + name + "】，请关注。"
	return x
}

// 消息提示 -- 工单修改
func WorkListUpdate(name string) string {
	x := "工单【" + name + "】被修改，请关注。"
	return x
}

// 消息提示 -- 工单关闭
func WorkDelete(name string) string {
	x := "工单【" + name + "】已关闭"
	return x
}

// 消息提示 -- 订单待续费提示
func OrderRenewalFee(oid, day string) string {
	x := "订单【" + oid + "】还剩" + day + "天到期，请及时通知客户续费。"
	return x
}

// 消息提示 -- 订单到期提示
func OrderExpire(oid string) string {
	x := "订单【" + oid + "】已到期，请关注。"
	return x
}

// 消息提示 -- 订单续费提示
func OrderRenewSuccess(oid string) string {
	x := "订单【" + oid + "】续费成功，请关注。"
	return x
}

// 消息提示 -- 订单退款提示
func OrderRefundSuccess(oid string) string {
	x := "订单【" + oid + "】已退款，请关注。"
	return x
}

// 消息提示 -- 订单延期提示
func OrderDelaySuccess(oid string) string {
	x := "订单【" + oid + "】延期成功，请关注。"
	return x
}

// 消息提示 -- 订单删除提示
func OrderDeleteSuccess(oid string) string {
	x := "订单【" + oid + "】已被删除，请关注。"
	return x
}

// 消息提示 -- 订单新增提示
func OrderAddSuccess(oid string) string {
	x := "你被指派了一个新订单【" + oid + "】，请关注。"
	return x
}

// 消息提示 -- 交接班提醒
func ShiftSuccess(name, t, title string) string {
	x := "【" + name + "】【" + t + "】的交班记录【" + title + "】提醒你查看。"
	return x
}

// 消息提示 -- 域名到期
func PromptDomain(name, t string) string {
	x := "域名【" + name + "】还剩【" + t + "】天到期，请及时续费。"
	return x
}

// 主机raid卡等级校验
func CheckRaidLevel(raidLevel string) bool {
	if raidLevel == "未使用" || raidLevel == "0" || raidLevel == "1" || raidLevel == "10" || raidLevel == "01" ||
		raidLevel == "3" || raidLevel == "4" || raidLevel == "5" || raidLevel == "6" {
		return true
	}
	return false
}
