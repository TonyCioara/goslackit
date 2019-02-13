package slack

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/nlopes/slack"
)

/*
   TODO: Change @BOT_NAME to the same thing you entered when creating your Slack application.
   NOTE: command_arg_1 and command_arg_2 represent optional parameteras that you define
   in the Slack API UI
*/
const helpMessage = "type in '@shuffle-bot shuffle or just @shuffle-bot'"

/*
   CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
   initiating the socket connection and returning the client.
   DO NOT EDIT THIS FUNCTION. This is a fully complete implementation.
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey) // this is some cool stuff
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

/*
   RespondToEvents waits for messages on the Slack client's incomingEvents channel,
   and sends a response when it detects the bot has been tagged in a message with @<botTag>.

   EDIT THIS FUNCTION IN THE SPACE INDICATED ONLY!
*/
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			message := strings.Replace(ev.Msg.Text, botTagString, "", -1)

			// TODO: Make your bot do more than respond to a help command. See notes below.
			// Make changes below this line and add additional funcs to support your bot's functionality.
			// sendHelp is provided as a simple example. Your team may want to call a free external API
			// in a function called sendResponse that you'd create below the definition of sendHelp,
			// and call in this context to ensure execution when the bot receives an event.

			// START SLACKBOT CUSTOM CODE
			// ===============================================================
			sendResponse(slackClient, message, ev.Channel)
			sendHelp(slackClient, message, ev.Channel)
			annoy(slackClient, message, ev.Channel)
			// ===============================================================
			// END SLACKBOT CUSTOM CODE
		default:

		}
	}
}

// sendHelp is a working help message, for reference.
func sendHelp(slackClient *slack.RTM, message, slackChannel string) {
	if strings.ToLower(message) != "help" {
		return
	}
	slackClient.SendMessage(slackClient.NewOutgoingMessage(helpMessage, slackChannel))
}

// sendResponse is NOT unimplemented --- write code in the function body to complete!

func sendResponse(slackClient *slack.RTM, message, slackChannel string) {

	if !strings.Contains(strings.ToLower(message), "scramble") {
		return
	}

	command := strings.ToLower(message)
	println("[RECEIVED] sendResponse:", command)

	runes := []rune(message)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := len(runes); i > 0; i-- {
		randIndex := r.Intn(i)
		runes[i-1], runes[randIndex] = runes[randIndex], runes[i-1]
	}

	newMessage := ""
	for _, rune := range runes {
		newMessage += string(rune)
	}

	slackClient.SendMessage(slackClient.NewOutgoingMessage(newMessage, slackChannel))

	// START SLACKBOT CUSTOM CODE
	// ===============================================================
	// TODO:
	//      1. Implement sendResponse for one or more of your custom Slackbot commands.
	//         You could call an external API here, or create your own string response. Anything goes!
	//      2. STRETCH: Write a goroutine that calls an external API based on the data received in this function.
	// ===============================================================
	// END SLACKBOT CUSTOM CODE
}

func annoy(slackClient *slack.RTM, message, slackChannel string) {

	newMessage := ""
	for _, char := range message {
		if unicode.IsLetter(char) {
			r := rand.Intn(2)
			if r == 1 {
				newMessage += string(unicode.ToLower(char))
			} else {
				newMessage += string(unicode.ToUpper(char))
			}
		} else {
			newMessage += string(char)
		}
	}

	slackClient.SendMessage(slackClient.NewOutgoingMessage(newMessage, slackChannel))
}
