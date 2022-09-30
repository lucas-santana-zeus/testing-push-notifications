package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

var token = "cEUibWSJrU6AmXy-E1RQ3n:APA91bHO-Ywl9QlASxecaVorsiDFSbrHZRArXrRtJ9NLQtptLJbYwLGlakbIo_Asj1dPobnlTBLd_wCu19MBeuiMfD0FNj3G5Q1lbztIbxJbs4crBL3AXyojtF-obtxHfE3e7NyYhVw5"

func main() {
	sendToOneToken()

}

func sendToOneToken() {
	//Token ratinho
	// token := "c8bfpDJ2Crx3RAqflcKpXS:APA91bHRTFJV1vOxA-JKE9whYKvG1wqot7V7Sb9BK-MRaSER2JyybaDHbImogPUZ2HIfVv6l-jia8ODHCB4UfeToEyN-YCotU_hIktTyNVP_Jaoxb65HoWPcQ7pc3KkzSMBu5Q6kJYJ9"
	// token := "cEUibWSJrU6AmXy-E1RQ3n:APA91bHO-Ywl9QlASxecaVorsiDFSbrHZRArXrRtJ9NLQtptLJbYwLGlakbIo_Asj1dPobnlTBLd_wCu19MBeuiMfD0FNj3G5Q1lbztIbxJbs4crBL3AXyojtF-obtxHfE3e7NyYhVw5"
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

func sendToMoreThanOne() {
	registrationTokens := []string{
		"cEUibWSJrU6AmXy-E1RQ3n:APA91bHO-Ywl9QlASxecaVorsiDFSbrHZRArXrRtJ9NLQtptLJbYwLGlakbIo_Asj1dPobnlTBLd_wCu19MBeuiMfD0FNj3G5Q1lbztIbxJbs4crBL3AXyojtF-obtxHfE3e7NyYhVw5",
		"c8bfpDJ2Crx3RAqflcKpXS:APA91bHRTFJV1vOxA-JKE9whYKvG1wqot7V7Sb9BK-MRaSER2JyybaDHbImogPUZ2HIfVv6l-jia8ODHCB4UfeToEyN-YCotU_hIktTyNVP_Jaoxb65HoWPcQ7pc3KkzSMBu5Q6kJYJ9",
		"eqi7m_qeQgu4WQmv6MD3PL:APA91bHk6OGP3sRoNzx5hD4rIW72R0zyuJO6lybs2_UBBeV4rG40Pzkai6UMwlw9anI_5ftFZW95yk1nG59jdgQ0BDH2Lw5mQs59m2xj4If80BKhJ3xXokJnGUdmZ_cWHKnKrZsH_qxl",
	}
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalln("new app firebase: ", err)
	}
	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalln("app messagin: ", err)
	}
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "Vai chover meu parceiro!",
			Body:  "Abre o guarda chuvas aí meu patrão",
			// ImageURL: "https://cdn.brasildefato.com.br/media/1d3118aff09d4986f692c4ed897f6e12.jpg",
		},

		Tokens: registrationTokens,
	}

	br, err := fcmClient.SendMulticast(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}

	// See the BatchResponse reference documentation
	// for the contents of response.
	fmt.Printf("%d messages were sent successfully\n", br.SuccessCount)
	fmt.Println("Texto da resposta:", br.Responses)
	if br.FailureCount > 0 {
		var failedTokens []string
		for idx, resp := range br.Responses {
			if !resp.Success {
				// The order of responses corresponds to the order of the registration tokens.
				failedTokens = append(failedTokens, registrationTokens[idx])
			}
		}

		fmt.Printf("List of tokens that caused failures: %v\n", failedTokens)
	}

}

// func sendByTopic() {
// 	app, err := firebase.NewApp(context.Background(), nil)
// 	if err != nil {
// 		log.Fatalln("new app firebase: ", err)
// 	}
// 	fcmClient, err := app.Messaging(context.Background())
// 	if err != nil {
// 		log.Fatalln("app messagin: ", err)
// 	}
// 	// The topic name can be optionally prefixed with "/topics/".
// 	topic := "highScores"

// 	// See documentation on defining a message payload.
// 	message := &messaging.Message{
// 		Data: map[string]string{
// 			"score": "850",
// 			"time":  "2:45",
// 		},
// 		Topic: topic,
// 	}

// 	// Send a message to the devices subscribed to the provided topic.
// 	response, err := fcmClient.Send(context.Background(), message)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// Response is a message ID string.
// 	fmt.Println("Successfully sent message:", response)
// }
