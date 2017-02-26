package main

import (
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"github.com/shawncatz/go-destiny/destiny"
	"os"
)

var Applications = map[string]interface{}{
	"/echo/helloworld": alexa.EchoApplication{ // Route
		AppID:    "xxxxxxxx", // Echo App ID from Amazon Dashboard
		OnIntent: EchoIntentHandler,
		OnLaunch: EchoIntentHandler,
	},
	"/echo/destiny": alexa.EchoApplication{ // Route
		AppID:    "xxxxxxxx", // Echo App ID from Amazon Dashboard
		OnIntent: DestinyIntentHandler,
		OnLaunch: DestinyIntentHandler,
	},
}

func main() {
	alexa.Run(Applications, "3000")
}

func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")
}

func DestinyIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	client = destiny.New(os.Getenv("DESTINY_API"), os.Getenv("DESTINY_USER"), os.Getenv("DESTINY_PASS"), os.Getenv("DESTINY_TYPE"))
	echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")
}
