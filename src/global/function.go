package global

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// 获取运行程序路径
func GetExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		Glogger.Error(err.Error())
		return "", err
	}
	p, err := filepath.Abs(file)
	if err != nil {
		Glogger.Error(err.Error())
		return "", err
	}
	p = filepath.Dir(p)
	return p, nil
}

// 字符串ip转换为int
func StringIpToInt(ip string) int {
	var ips []string
	if ip == "::1" {
		ips = []string{"127", "0", "0", "1"}
	} else {
		ips = strings.Split(ip, ".")
	}
	ipInt := 0
	var pos uint = 24
	for _, ipSeg := range ips {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

// ip由数字转为字符串
func IpIntToString(i int) string {
	if i == 0 {
		return ""
	}
	i4 := i & 255
	i3 := i >> 8 & 255
	i2 := i >> 16 & 255
	i1 := i >> 24 & 255
	if i1 > 255 || i2 > 255 || i3 > 255 || i4 > 255 {
		return ""
	}
	ipstring := fmt.Sprintf("%d.%d.%d.%d", i1, i2, i3, i4)
	return ipstring
}

// md5加密密码
func Md5Password(p string) string {
	h := md5.New()
	io.WriteString(h, EncryptSalt)
	io.WriteString(h, p)
	io.WriteString(h, EncryptSalt)
	return hex.EncodeToString(h.Sum(nil))
}

// 密码校验
func CheckPassword(password string) bool {
	reg := regexp.MustCompile(`[a-z]`) // 查看是否有小写字母
	ss := reg.FindAllString(password, -1)
	if len(ss) > 0 {
		reg := regexp.MustCompile(`[A-Z]`) // 查看是否有大写字母
		ss := reg.FindAllString(password, -1)
		if len(ss) > 0 {
			reg := regexp.MustCompile(`[0-9]`) // 查看是否有数字
			ss := reg.FindAllString(password, -1)
			if len(ss) > 0 {
				return true
			}
		}
	}
	return false
}

// 区域缩写校验
func CheckAreaOmit(title string) (bool, error) {
	titleJs := "^[A-Z]{1,10}$"
	flag, err := regexp.MatchString(titleJs, title)
	return flag, err
}

// 机房缩写校验
func CheckRoomOmit(title string) (bool, error) {
	titleJs := "^[A-Z]{1}[A-Z0-9]{0,9}$"
	flag, err := regexp.MatchString(titleJs, title)
	return flag, err
}

// 机柜名称
func CheckCabinetTitle(title string) (bool, error) {
	titleJs := "^[A-Z]{1}[A-Z0-9]{0,19}$"
	flag, err := regexp.MatchString(titleJs, title)
	return flag, err
}

// 邮箱校验
func CheckEmail(email string) string {
	if email == "" {
		return "valid.success"
	}
	emailJS := "[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"
	flag, err := regexp.MatchString(emailJS, email)
	if err != nil {
		return "system.error"
	}
	if !flag {
		return "email.email"
	}
	return "valid.success"
}

// 微信校验
func CheckWechat(wechat string) string {
	if wechat == "" {
		return "valid.success"
	}
	wechatJS := "^[a-zA-Z]{1}[-_a-zA-Z0-9]{5,19}$"
	flag, err := regexp.MatchString(wechatJS, wechat)
	if err != nil {
		return "system.error"
	}
	if !flag {
		return "wechat.error"
	}
	return "valid.success"
}

// qq校验
func CheckQQ(qq string) string {
	if qq == "" {
		return "valid.success"
	}
	qqJS := "^[1-9]{1}[0-9]{4,9}$"
	flag, err := regexp.MatchString(qqJS, qq)
	if err != nil {
		return "db.error"
	}
	if !flag {
		return "qq.error"
	}
	return "valid.success"
}

// 联系电话校验
func CheckPhone(phone string) string {
	if phone == "" {
		return "valid.success"
	}
	phoneJS := "^^1([0-9][0-9]|14[57]|5[^4])\\d{8}$"
	flag, err := regexp.MatchString(phoneJS, phone)
	if err != nil {
		return "db.error"
	}
	if !flag {
		return "mobile.mobile"
	}
	return "valid.success"
}

// ip校验
func CheckIp(ip string) string {
	if ip == "" {
		return "valid.success"
	}
	IpJS := "^(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|[1-9])\\." +
		"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)\\." +
		"(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)\\." + "(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)$"
	flag, err := regexp.MatchString(IpJS, ip)
	if err != nil {
		return "system.error"
	}
	if !flag {
		return "ip.invalid"
	}
	return "valid.success"
}

// mac校验
// func CheckMac(mac string) (bool, error) {
// 	ok, err := regexp.MatchString("^([A-Fa-f0-9]{2}-){5}[A-Fa-f0-9]{2}$|^([A-Fa-f0-9]{2}:){5}[A-Fa-f0-9]{2}$",
// 		mac)
// 	return ok, err
// }

func AesEncrypt(str string) []byte {
	k := getKey()
	iv := []byte(k)[:aes.BlockSize]
	enc := make([]byte, len(str))
	nc, _ := aes.NewCipher(k)
	ns := cipher.NewCFBEncrypter(nc, iv)
	ns.XORKeyStream(enc, []byte(str))
	return enc
}

func AesDecrypt(src []byte) string {
	k := getKey()
	iv := []byte(k)[:aes.BlockSize]
	dec := make([]byte, len(src))
	abd, _ := aes.NewCipher([]byte(k))
	ns := cipher.NewCFBDecrypter(abd, iv)
	ns.XORKeyStream(dec, src)
	return string(dec)
}

func getKey() []byte {
	return []byte(EncryptSalt)[:32]
}

func ParseResponseMsg(key, alisa string) string {
	if key == "" {
		Glogger.Debug("error key 不能为空")
		return ErrNotNull
	}
	s := strings.Split(strings.ToLower(key), ".")
	if len(s) != 2 || s[0] == "" || s[1] == "" {
		Glogger.Debug("response key 格式不正确")
		return ErrFormatError
	}
	val, ok := ErrorMsg[s[0]]
	if !ok {
		Glogger.Debug("response key 不存在")
		return ErrNotFound
	}
	v, ok := val[s[1]]
	if !ok {
		Glogger.Debug("response key 不存在")
		return ErrNotFound
	}
	return strings.Replace(v, "%s", alisa, -1)
}

// 站点校验
func CheckSite(site string) string {
	if site == "" {
		return "valid.success"
	}
	siteJS := "^[a-zA-Z\\d]{1,20}$"
	flag, err := regexp.MatchString(siteJS, site)
	if err != nil {
		return "system.error"
	}
	if !flag {
		return "order.site"
	}
	return "valid.success"
}
