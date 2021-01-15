package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Fatalf("%s: %s", msg, err)
//	}
//}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
const MQURL ="amqp://admin:admin@172.16.0.91:5672/"
func main() {
	conn, err := amqp.Dial(MQURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//body := bodyFrom(os.Args)
	for i:=1;i<10;i++ {
		body:=fmt.Sprintf("m-%d",i)
		// 将消息发送到延时队列上
		err = ch.Publish(
			"",           // exchange 这里为空则不选择 exchange
			"test_delay", // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
				Expiration:  "5000", // 设置五秒的过期时间
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s", body)
		time.Sleep(time.Second*5)
	}


}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
