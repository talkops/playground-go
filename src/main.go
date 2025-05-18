package main

import (
    "fmt"
    "github.com/talkops/sdk-go"
    "math/rand"
    "time"
)

var extension *talkops.Extension

var color *talkops.Parameter
var email *talkops.Parameter

func onBoot(args map[string]interface{}) {
	fmt.Println("onBoot")
	fmt.Println(args["language"])
	fmt.Println(color.GetValue())
	fmt.Println(email.GetValue())
}

func onSession(args map[string]interface{}) {
	fmt.Println("onSession")
	fmt.Println(args["language"])
	fmt.Println(color.GetValue())
	fmt.Println(email.GetValue())
}

func onEnable(args map[string]interface{}) {
	fmt.Println("onEnable")
	fmt.Println(args["language"])
	fmt.Println(color.GetValue())
	fmt.Println(email.GetValue())
}

func onDisable(args map[string]interface{}) {
	fmt.Println("onDisable")
	fmt.Println(args["language"])
	fmt.Println(color.GetValue())
	fmt.Println(email.GetValue())
}

func enableAlarm(args map[string]interface{}) string {
	extension.EnableAlarm()
	return "Done."
}

func receiveRandomDice(args map[string]interface{}) string {
	return fmt.Sprintf("%d", rand.Intn(6)+1)
}

func receiveRandomDiceAsynchronously(args map[string]interface{}) string {
	time.Sleep(10 * time.Second)
	return fmt.Sprintf("%d", rand.Intn(6)+1)
}

func receiveRandomDiceMessage(args map[string]interface{}) string {
	msg := fmt.Sprintf("%d", rand.Intn(6)+1)
	extension.SendMessage(msg)
	return "Done."
}

func receiveRandomDiceNotification(args map[string]interface{}) string {
	notif := fmt.Sprintf("%d", rand.Intn(6)+1)
	extension.SendNotification(notif)
	return "Done."
}

func receiveRandomImage(args map[string]interface{}) string {
	url := fmt.Sprintf("https://picsum.photos/seed/%d/640/480", rand.Intn(100)+1)
	extension.SendMedias([]*talkops.Media{
		talkops.NewAttachment(url, "test.jpg"),
		talkops.NewImage(url),
	})
	return "Done."
}

func receiveRandomVideo(args map[string]interface{}) string {
	videos := []string{
		"https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
		"https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4",
		"https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/Sintel.mp4",
		"https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/TearsOfSteel.mp4",
	}
	extension.SendMedias([]*talkops.Media{
		talkops.NewVideo(videos[rand.Intn(len(videos))]),
	})
	return "Done."
}

func main() {
	color = talkops.NewParameter("COLOR").
		SetDescription("The color used for test.").
		SetType("select").
		SetDefaultValue("blue").
		SetAvailableValues([]string{"red","orange","yellow","green","blue","indigo","violet"})

	email = talkops.NewParameter("EMAIL").
		SetDescription("The email used for test.").
		SetType("email").
		SetDefaultValue("john.doe@example.com")

	extension = talkops.NewExtension()
	extension.
		SetName("Go Playground").
		SetIcon("https://talkops.app/images/extensions/playground-go.png").
		SetCategory("utility").
		SetDemo(true).
		SetFeatures([]string{
			"Enable the alarm",
			"Receive a random dice",
			"Receive a random dice asynchronously",
			"Receive a random dice as message",
			"Receive a random dice as notification",
			"Receive a random image",
			"Receive a random video",
		}).
		SetParameters([]*talkops.Parameter{
			color,
			email,
		}).
		SetFunctionSchemas([]map[string]interface{}{
			// {
			// 	"name": "get_weather",
			// 	"description": "Get weather at the given location",
			// 	"parameters": map[string]interface{}{
			// 		"type": "object",
			// 		"properties": map[string]interface{}{
			// 			"location": map[string]string{
			// 				"type": "string",
			// 				"description": "The location",
			// 			},
			// 		},
			// 		"required": []string{"location"},
			// 	},
			// },
			{
				"name": "enable_alarm",
				"description": "Enable the alarm.",
			},
			{
				"name": "receive_random_dice",
				"description": "Receive a random dice.",
			},
			{
				"name": "receive_random_dice_asynchronously",
				"description": "Receive a random dice asynchronously.",
			},
			{
				"name": "receive_random_dice_message",
				"description": "Receive a random dice as message.",
			},
			{
				"name": "receive_random_dice_notification",
				"description": "Receive a random dice as notification.",
			},
			{
				"name": "receive_random_image",
				"description": "Receive a random image.",
			},
			{
				"name": "receive_random_video",
				"description": "Receive a random video.",
			},
		}).
		SetFunctions(map[string]func(args map[string]interface{}) string{
			"enable_alarm": enableAlarm,
			"receive_random_dice": receiveRandomDice,
			"receive_random_dice_asynchronously": receiveRandomDiceAsynchronously,
			"receive_random_dice_message": receiveRandomDiceMessage,
			"receive_random_dice_notification": receiveRandomDiceNotification,
			"receive_random_image": receiveRandomImage,
			"receive_random_video": receiveRandomVideo,
		}).
		On("boot", onBoot).
		On("session", onSession).
		On("enable", onEnable).
		On("disable", onDisable).
		Start()
}
