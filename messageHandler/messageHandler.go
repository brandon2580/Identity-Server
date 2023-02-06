package messageHandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Failed to read request body: %v", err)
		return
	}
	fmt.Fprintf(w, "Webhook received")
	fmt.Println("New message:", string(body))
}

func SendSMS(to, from, body string) error {
	accountSID, _ := os.LookupEnv("ACCOUNTSID")
	authToken, _ := os.LookupEnv("AUTHTOKEN")
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", from)
	msgData.Set("Body", body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/"+accountSID+"/Messages.json", &msgDataReader)
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println("Text message sent!")
		return nil
	} else {
		return fmt.Errorf("Text message failed with status code: %d", resp.StatusCode)
	}
}
