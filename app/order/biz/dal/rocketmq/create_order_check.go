package rocketmq

import (
	"log"

	"github.com/apache/rocketmq-clients/golang"
)

func CreateOrderCheck(msg *golang.MessageView) golang.TransactionResolution {

	log.Println("create order check")
	return golang.COMMIT
}
