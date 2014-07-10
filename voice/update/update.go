package update

import (
	"net/http"
)

const (
	postUrl = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls/%s.json"
)

type Update struct {
	Url    string `json:"url"`
	Method string `json:"method"`
	Status string `json:"status"`
	// The below are optional
	FallbackUrl          string `json:"fallback_url"`
	FallbackMethod       string `json:"fallback_method"`
	StatusCallback       string `json:"status_callback"`
	StatusCallbackMethod string `json:"status_callback_method"`
}

func (u *Update) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while reading json from buffer => %s", err.Error()))
	}
	err = json.Unmarshal(bodyBytes, u)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes)))
	}
	return nil
}

// Internal function for sending the post request to twilio.
func (act VoiceAccount) postUpdate(dest string, msg Post, resp *Update) error {
	// send post request to twilio
	return common.SendPostRequest(dest, msg, act, resp, 201)
}

// Sends a post request to Twilio to modify a call
func (act VoiceAccount) Update(p Post, sid string) (Update, error) {
	var r Update
	err := act.postUpdate(fmt.Sprintf(postUrl, act.AccountSid, string(sid)), p, &r)
	return r, err
}
