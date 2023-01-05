package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var botUrl string

func main() {
	botToken := "5857442294:AAGoadu1j7o1bVYDzu0QmVVXviFYOl2UIfA"
	botApi := "https://api.telegram.org/bot"
	botUrl = botApi + botToken
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Printf("Error - %s", err.Error())
		}
		for _, update := range updates {
			err = respond(botUrl, update)
			offset = update.UpdateId + 1
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println(updates)
	}
}

func getUpdates(bitUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
