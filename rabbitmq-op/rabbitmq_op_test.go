package rabbitmq_op

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestRabbit(t *testing.T) {

	////Simple模式 发送者
	//rabbitmq := NewRabbitMQSimple("imoocSimple")
	//rabbitmq.PublishSimple("hello imooc!")
	////接收者
	//rabbitmq := NewRabbitMQSimple("imoocSimple")
	//rabbitmq.ConsumeSimple()


	//重启消息不丢
	//订阅模式发送者
	go func() {
		rabbitmq := NewRabbitMQPubSub("" + "extest3")
		for i := 10; i <= 100; i++ {
			rabbitmq.PublishPub2("订阅模式生产第" + strconv.Itoa(i) + "条数据")
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}()

	//接收者
	go func() {

		rabbitmq := NewRabbitMQPubSub("" + "extest3")
		rabbitmq.ReceiverSub2()
	}()



	//重启消息丢失
	//订阅模式发送者
	//go func() {
	//	rabbitmq := NewRabbitMQPubSub("" + "newProduct2")
	//	for i := 0; i <= 100; i++ {
	//		rabbitmq.PublishPub("订阅模式生产第" + strconv.Itoa(i) + "条数据")
	//		fmt.Println(i)
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//
	////接收者
	//go func() {
	//
	//	rabbitmq := NewRabbitMQPubSub("" + "newProduct2")
	//	rabbitmq.RecieveSub()
	//}()
	//
	//go func() {
	//	rabbitmq := NewRabbitMQPubSub("" + "newProduct2")
	//	rabbitmq.RecieveSub()
	//}()
	time.Sleep(time.Second*100)

	////路由模式发送者
	//imoocOne := NewRabbitMQRouting("exImooc", "imooc_one")
	//imoocTwo := NewRabbitMQRouting("exImooc", "imooc_two")
	//
	//for i := 0; i <= 10; i++ {
	//	imoocOne.PublishRouting("hello imooc one!" + strconv.Itoa(i))
	//	imoocTwo.PublishRouting("hello imooc two!" + strconv.Itoa(i))
	//	time.Sleep(1 * time.Second)
	//	fmt.Println(i)
	//}
	////接收者
	//rabbitmq := NewRabbitMQRouting("exImooc", "imooc_one")
	//rabbitmq.RecieveRouting()
	//
	////Topic模式发送者
	//imoocOne := NewRabbitMQTopic("exImoocTopic", "imooc.topic88.three")
	//imoocTwo := NewRabbitMQTopic("exImoocTopic", "imooc.topic88.four")
	//
	//for i := 0; i <= 10; i++ {
	//	imoocOne.PublishTopic("hello imooc topic three!" + strconv.Itoa(i))
	//	imoocTwo.PublishTopic("hello imooc topic four!" + strconv.Itoa(i))
	//	time.Sleep(1 * time.Second)
	//	fmt.Println(i)
	//}
	////Topic接收者
	//rabbitmq := NewRabbitMQTopic("exImoocTopic", "#")
	//rabbitmq.RecieveTopic()
}