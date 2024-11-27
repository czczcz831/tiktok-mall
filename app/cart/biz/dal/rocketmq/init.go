package rocketmq

import (
	"os"
)

func Init() {
	os.Setenv("mq.consoleAppender.enabled", "true")

	err := clearCartConsumerInit()

	if err != nil {
		panic(err)
	}

}
