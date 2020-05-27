package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	model "webhook/module"
	"webhook/transformer"


)

func Send(notification model.Notification, dingtalkRobot string) (err error) {

	markdown, err := transformer.TransformToMarkdown(notification)

	if err != nil {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		dingtalkRobot,
		bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return err
}