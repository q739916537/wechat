package main

import (
	"fmt"
	"net/http"

	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

func serveWechat(rw http.ResponseWriter, req *http.Request) {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:          "wxb951ff3ba8358a14",
		AppSecret:      "5e60d1f791765b4da1f6ba8487a01121",
		Token:          "qaz739916537",
		EncodingAESKey: "cAsnU6tOZLsHCxnX8erFHXF8ioZZSf1LvGxMBOng1r4",
		Cache:          memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	user := officialAccount.GetUser()
	officialAccount.GetMenu()
	fmt.Println(user)
	// 传入request和responseWriter
	server := officialAccount.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func main() {
	http.HandleFunc("/", serveWechat)
	fmt.Println("wechat server listener at", ":8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}
}
