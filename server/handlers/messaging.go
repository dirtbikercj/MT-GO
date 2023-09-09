package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"fmt"
	"net/http"
)

func MessagingFriendList(w http.ResponseWriter, r *http.Request) {
	friends := database.GetAccountByUID(services.GetSessionID(r)).Friends
	body := services.ApplyResponseBody(friends)
	services.ZlibJSONReply(w, body)
}

func MessagingDialogList(w http.ResponseWriter, r *http.Request) {
	dialogues := database.GetDialogueByUID(services.GetSessionID(r))

	data := []interface{}{}
	for _, dialogue := range *dialogues {
		data = append(data, dialogue)
	}

	body := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, body)
}

type FriendRequestMailbox struct {
	Err  int           `json:"err"`
	Data []interface{} `json:"data"`
}

func MessagingFriendRequestInbox(w http.ResponseWriter, r *http.Request) {
	friends := database.GetAccountByUID(services.GetSessionID(r)).FriendRequestInbox
	body := &FriendRequestMailbox{
		Data: friends,
	}
	services.ZlibJSONReply(w, body)
}

func MessagingFriendRequestOutbox(w http.ResponseWriter, r *http.Request) {
	friends := database.GetAccountByUID(services.GetSessionID(r)).FriendRequestOutbox
	body := &FriendRequestMailbox{
		Data: friends,
	}
	services.ZlibJSONReply(w, body)
}

func MessagingMailDialogInfo(w http.ResponseWriter, r *http.Request) {
	parsedData := services.GetParsedBody(r)
	dialogId, _ := parsedData.(map[string]interface{})["dialogId"].(string)

	dialogues := *database.GetDialogueByUID(services.GetSessionID(r))
	dialog, ok := dialogues[dialogId]
	if !ok {
		fmt.Println("Dialogue does not exist")
	}

	dialogInfo := dialog.CreateQuestDialogueInfo()

	body := services.ApplyResponseBody(dialogInfo)
	services.ZlibJSONReply(w, body)
}
