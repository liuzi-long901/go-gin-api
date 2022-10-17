package receiver

import (
	"encoding/json"
	"jassue-gin/bootstrap"
	"jassue-gin/global"
	"jassue-gin/sender"
	"time"

	"github.com/x-funs/go-fun"
)

func RabbitSimple() {
	// goroutines
	goroutines := 1
	//if config.Cfg.RabbitQueue.Simple.Goroutines > 0 {
	//	goroutines = config.Cfg.RabbitQueue.Simple.Goroutines
	//}

	queueName := global.App.Config.RabbitmqQueue.Simple.Name
	simple, _ := bootstrap.NewSimple(queueName)

	for i := 0; i < goroutines; i++ {
		no := i
		go func() {
			if msgs, err := simple.Receive(); err == nil {
				for msg := range msgs {
					var demo sender.Demo
					if e := json.Unmarshal(msg.Body, &demo); e == nil {
						bootstrap.Info("Received simple message:", bootstrap.Int("no", no), bootstrap.Any("demo", demo))

						// 模拟耗时
						s := fun.RandomInt(300, 800)
						time.Sleep(time.Millisecond * time.Duration(s))

						// 消息确认
						_ = msg.Ack(false)
					}
				}
			} else {
				bootstrap.Error("Received topic error")
			}
		}()
	}
}

func RabbitTopic() {
	// goroutines
	goroutines := 1
	//if config.Cfg.RabbitQueue.Topic.Goroutines > 0 {
	//	goroutines = config.Cfg.RabbitQueue.Topic.Goroutines
	//}

	exchangeName := global.App.Config.RabbitmqQueue.Topic.Exchange
	queueName := global.App.Config.RabbitmqQueue.Simple.Name
	routingKeys := fun.SliceTrim(global.App.Config.RabbitmqQueue.Topic.RoutingKeys)

	topic, _ := bootstrap.NewTopic(exchangeName)

	for i := 0; i < goroutines; i++ {
		no := i
		go func() {
			if msgs, err := topic.ReceiveWithRoutingKeys(queueName, routingKeys); err == nil {
				for msg := range msgs {
					var demo sender.Demo
					if e := json.Unmarshal(msg.Body, &demo); e == nil {
						bootstrap.Info("Received topic message:", bootstrap.Int("no", no), bootstrap.Any("demo", demo))

						// 模拟耗时
						s := fun.RandomInt(300, 800)
						time.Sleep(time.Millisecond * time.Duration(s))

						// 消息确认
						_ = msg.Ack(false)
					}
				}
			} else {
				bootstrap.Error("Received topic error")
			}
		}()
	}
}
