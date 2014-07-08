package voice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type VoiceAccount struct {
	AccountSid string
	Token      string
}

// Represents the data used in creating an outbound voice message.
// "From" & "To" are required attributes.
// Either a ApplicationSid or a Url must also be provided.
// Visit https://www.twilio.com/docs/api/rest/making-calls#post-parameters for more details.
type Post struct {
	From           string
	To             string
	Body           string
	Url            string
	StatusCallback string
	ApplicationSid string
}

func validatePost(p Post) error {
	if p.From == "" || p.To == "" {
		return errors.New("Both \"From\" and \"To\" must be set in Post.")
	}
	if p.ApplicationSid == "" && p.Url == "" {
		return errors.New("Either \"ApplicationSid\" or \"Url\" must be set.")
	}
	return nil
}

// Represents the callback sent everytime the status of the call is updated.
type Callback struct {
	standardRequest
	RecordingUrl string
	CallDuration string
}

func (p Post) getReader() io.Reader {
	vals := url.Values{}
	if p.To != "" {
		vals.Set("To", p.To)
	}
	if p.From != "" {
		vals.Set("From", p.From)
	}
	if p.Url != "" {
		vals.Set("Url", p.Url)
	}
	if p.ApplicationSid != "" {
		vals.Set("ApplicationSid", p.ApplicationSid)
	}

	return strings.NewReader(vals.Encode())

}

// Internal function for sending the post request to twilio.
func (act VoiceAccount) makeCall(dest string, msg Post, resp *Response) error {
	// send post request to twilio
	c := http.Client{}
	req, err := http.NewRequest("POST", dest, msg.getReader())

	if err != nil {
		return errors.New(fmt.Sprintf("Error with req => %s", err.Error()))
	}
	req.SetBasicAuth(act.AccountSid, act.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	twilioResp, err := c.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Error with resp => %s", err.Error()))
	}
	if twilioResp.StatusCode != 201 {
		return errors.New(fmt.Sprintf("Error recieved from Twilio => %s", twilioResp.Status))
	}

	// parse twilio response
	bodyBytes, err := ioutil.ReadAll(twilioResp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while reading json from buffer => %s", err.Error()))
	}
	err = json.Unmarshal(bodyBytes, resp)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes)))
	}
	return nil
}

// Sends a post request to Twilio to send a voice request.
func (act VoiceAccount) Call(p Post) (Response, error) {
	err := validatePost(p)
	if err != nil {
		return Response{}, errors.New(fmt.Sprintf("Error validating voice post => %s.\n", err.Error()))
	}

	var r Response
	err = act.makeCall(fmt.Sprintf(postUrl, act.AccountSid), p, &r)

	return r, nil
}
