package Nsq

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"wPan/v1/Config"
	"wPan/v1/Utils"
)

type NsqHandler struct{}

var P *nsq.Producer

func (s *NsqHandler) HandleMessage(message *nsq.Message) error {
	//解析传递的json数据
	var mapData map[string]interface{}
	_ = json.Unmarshal(message.Body, &mapData)
	var to []string
	for _, v := range mapData["to"].([]interface{}) {
		to = append(to, v.(string))
	}
	err := Utils.SendMail(to, mapData["subject"].(string), mapData["body"].(string))
	if err != nil {
		fmt.Println("发送失败")
	}
	return nil
}

func InitNSQ() {
	config := nsq.NewConfig()
	P, _ = nsq.NewProducer(Config.NSQSetting.Host, config)
	com, err := nsq.NewConsumer(Config.ServerSetting.Name, "email", config)
	if err != nil {
		fmt.Println(err)
	}
	com.AddHandler(&NsqHandler{})
	err = com.ConnectToNSQD(Config.NSQSetting.Host)

	if err != nil {
		fmt.Println(err)
	}
}

func SendNSQ(to []string, subject, body string) error {
	Command := make(map[string]interface{})
	Command["to"] = to
	Command["subject"] = subject
	Command["body"] = body
	data, err := json.Marshal(Command)
	err = P.Publish(Config.ServerSetting.Name, []byte(data))
	return err
}
