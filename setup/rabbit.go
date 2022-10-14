package setup

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"jassue-gin/bootstrap"
	"jassue-gin/di"
	"jassue-gin/global"
	"strconv"
)

func Rabbit() {
	if connection, err := amqp.Dial("amqp://" + global.App.Config.Rabbitmq.UserName +
		":" + global.App.Config.Rabbitmq.Password +
		"@" + global.App.Config.Rabbitmq.Host +
		":" + strconv.Itoa(global.App.Config.Rabbitmq.Port)); err == nil {
		if channel, err := connection.Channel(); err == nil {

			//// Qos
			//if config.Cfg.Rabbit.Prefetch > 0 {
			//	_ = channel.Qos(config.Cfg.Rabbit.Prefetch, 0, false)
			//}

			di.SetRabbit(channel)

			bootstrap.Info("Setup rabbit success")
		} else {
			bootstrap.Fatal("Setup rabbit channel failed", bootstrap.String("err", err.Error()))
		}
	} else {
		bootstrap.Fatal("Setup rabbit connection failed", bootstrap.String("err", err.Error()))
	}
}
