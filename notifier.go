package main

import (
	"bytes"
	"fmt"
	"github.com/kensodev/sns-parser"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	m := os.Args[1]
	parser := snsparser.NewSNSParser([]byte(m))

	failed, message := parser.IncludesMessage("Failed to deploy application")

	if failed {
		sendMessage(message)

	} else {
		fmt.Printf("Everything is OK, nothing to report in this message")
	}
}

func sendMessage(message snsparser.SNS) {
	data := getData(message)
	req, _ := http.NewRequest("POST", "SLACK_HOOK", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Message is 'Failed to deploy application', send to slack: ", string(body))
}

func getData(message snsparser.SNS) url.Values {
	data := url.Values{}
	jsonPayload := `
			{
				"channel": "#devs",
				"username": "webhookbot",
				"text": "ALERT: <!here> ElasticBeanstalk failed to deploy application %v",
				"icon_emoji": ":red_circle:"
			}
		`

	jsonMessage := fmt.Sprintf(jsonPayload, message.TopicArn)
	data.Set("payload", jsonMessage)
	return data
}