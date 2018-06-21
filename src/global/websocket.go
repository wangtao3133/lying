package global

import (
	"encoding/json"
	"github.com/golang/net/websocket"
	"sync"
)

type Message struct {
	Id         int64  `json:"id"`          // 通知id
	Title      int8   `json:"title"`       // 通知标题(类型)
	Content    string `json:"description"` // 通知内容
	CreateTime int64  `json:"datetime"`    // 通知时间
	UserID     int64  `json:"user_id"`     // 通知对象用户ID
	Count      int64  `json:"count"`       // 未读消息总数
}

var MapToConn sync.Map // 全局同步map结构,存储token => id,和 id => *websocket.Conn

// 供服务端发送消息
// m 消息体
// 使用方法 m := Message{...}
// go SendNotice(m)
func SendNotice(m Message) {
	data, err := json.Marshal(m)
	if err != nil {
		Glogger.Error(err.Error())
		return
	}

	if c, ok := MapToConn.Load(m.UserID); ok {
		conn := c.(*websocket.Conn)
		_, err := conn.Write(data)
		if err != nil {
			conn.Close()
			MapToConn.Delete(m.UserID)
		}
	}
}
