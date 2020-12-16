package main

import (
	"fmt"
	"rabbitmq-demo/rabbitmq"
)

type TestPro struct {
	msgContent   string
}

// 实现发送者
func (t *TestPro) MsgContent() string {
	return t.msgContent
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
}

func main() {
	msg := fmt.Sprintf("这是测试任务")
	t1 := &TestPro{
		msg,
	}
	t2 := &TestPro{
		msg,
	}
	queueExchange := &rabbitmq.QueueExchange{
		"x",
		"",  //info 存在走route模式，为空走publish/sub
		"ex",
		"direct",
	}
	mq := rabbitmq.New(queueExchange)
	mq.RegisterProducer(t1)
	mq2 := rabbitmq.New(queueExchange)
	mq2.RegisterReceiver(t2)
	//重启消息不丢失
	mq2.StartConsumer()
	mq.StartProducer()
	//mq2.StartConsumer()

	select {

	}
	fmt.Println("done")
}