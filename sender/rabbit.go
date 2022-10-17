package sender

import (
	"jassue-gin/bootstrap"
	"jassue-gin/global"
	"time"

	"github.com/x-funs/go-fun"
)

func RabbitSimple() {
	queueName := global.App.Config.RabbitmqQueue.Topic.Name
	if simple, err := bootstrap.NewSimple(queueName); err == nil {
		var id int

		for {
			id++
			msg := Demo{
				Id:   id,
				Name: fun.RandomLetter(4),
				Time: fun.Date(fun.DatetimeMilliPattern),
			}

			msgJson := fun.ToJson(msg)

			if err := simple.Send(fun.Bytes(msgJson)); err == nil {
				bootstrap.Info("Send simple success", bootstrap.String("msg", msgJson))
			} else {
				bootstrap.Error("Send simple error")
			}

			// sleep
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func RabbitTopic() {
	exchangeName := global.App.Config.RabbitmqQueue.Topic.Exchange
	if topic, err := bootstrap.NewTopic(exchangeName); err == nil {

		var id int
		var routingKey string

		for {
			id++

			// routingKey
			if id%2 == 0 {
				routingKey = "log.info"
			} else {
				routingKey = "log.warn"
			}

			msg := Demo{
				Id:   id,
				Name: fun.RandomLetter(4),
				Time: fun.Date(fun.DatetimeMilliPattern),
				Key:  routingKey,
			}

			msgJson := fun.ToJson(msg)

			if err := topic.Send(routingKey, fun.Bytes(msgJson)); err == nil {
				bootstrap.Info("Send topic success", bootstrap.String("key", routingKey), bootstrap.String("msg", msgJson))
			} else {
				bootstrap.Error("Send topic error")
			}

			// sleep
			time.Sleep(time.Millisecond * 10)
		}
	}
}

type Demo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
	Time string `json:"time"`
}
