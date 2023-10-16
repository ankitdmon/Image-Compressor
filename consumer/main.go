package main

import (
	"fmt"

	"github.com/ankitdmon/consumer/messaging"
)

func main(){
	fmt.Println("Hello Consumer");
	messaging.ConsumeFromRabbitMQ()
}