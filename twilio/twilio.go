package twilio

import(
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SendMessage(apiKey string, apiSecret string, message string) {
	post(apiKey, apiSecret, "Messages", url.Values{"From": {"+17205482998"}, "To": {"+13039104832"}, "Body": {message}})
}

func post(apiKey string, apiSecret string, endpoint string, urlValues url.Values) {
	endpointUrl := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/%s.json", apiKey, endpoint)

	req, err := http.NewRequest("POST", endpointUrl, strings.NewReader(urlValues.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(apiKey, apiSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, _ := client.Do(req)
	fmt.Println(resp)

	if err != nil {
		fmt.Println(err)
	}
}
