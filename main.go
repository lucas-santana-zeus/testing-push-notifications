package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"log"
	"os"
)

var token = os.Getenv("SMARTPHONE_TOKEN")

func main() {
	fmt.Println(token)

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalln("new app firebase: ", err)
	}
	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalln("app messagin: ", err)
	}

	response, err := fcmClient.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Vai chover meu parceiro!",
			Body:  "Abre o guarda chuvas aí meu patrão",
		},
		Token: token,
	})
	if err != nil {
		log.Fatalln("cloud message send: ", err)
	}

	fmt.Println(response)
}
