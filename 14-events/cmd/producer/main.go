package main

import "github.com/felipecardosodeoliveira/Golang/14-events/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "O mundo é dos devs", "amq.direct")
}
