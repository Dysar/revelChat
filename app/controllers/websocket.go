package controllers

import (
	"fmt"
	"time"

	"github.com/revel/revel"

	"github.com/revel/examples/chat/app/chatroom"

	"revelChat/app/twilio"
)

type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) Room(user string) revel.Result {
	return c.Render(user)
}

func (c WebSocket) RoomSocket(user string, ws revel.ServerWebSocket) revel.Result {

	//add token to message to detect if the admin is injecting a message (it's MD5 of Rick)
	authToken := "17ad55a9b8384777496330d23e59d520"

	// Make sure the websocket is valid.
	if ws == nil {
		return nil
	}

	// Join the room.
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	chatroom.Join(user)
	//chatroom.Say(user, fmt.Sprintf("Hey, %s, I accept commands including Estonian cell numbers. Once received I will call the number.", user))
	defer chatroom.Leave(user)

	// Send down the archive.
	for _, event := range subscription.Archive {
		if ws.MessageSendJSON(&event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := ws.MessageReceiveJSON(&msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket or the chatroom.
	for {
		select {
		case event := <-subscription.New:
			if ws.MessageSendJSON(&event) != nil {
				// They disconnected.
				return nil
			}

			//runs all the time
			// go func() {
			// 	time.Sleep(2 * time.Second)
			// 	chatroom.Say(user, authToken+"Welcome! You can type in an Estonian cell phone number and I will call it."+
			// 		"Also you can chat with other fellows in the chat. Enjoy :)")
			// }()
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, say something.
			c.Log.Info(fmt.Sprintf("chatroom is saying %s %s", user, msg))
			chatroom.Say(user, msg)

			//go routine is needed for sleeping (simulated thinking of the chatbot)
			go func() {
				// If the message contains a phone number approach

				// If the message contains phone number execute Twilio logic
				number, err := twilio.ExtractNumber(msg)
				time.Sleep(2 * time.Second)
				if err == nil {
					c.Log.Info("the message contains a phone number!")
					msg = twilio.CallNumber(number)
					msg = authToken + msg
					chatroom.Say(user, msg)
				} else if twilio.IsBadNumberInMessage(msg) {
					chatroom.Say(user, authToken+"The number is not valid.. Try again?")
				}
			}()
		}
	}
}
