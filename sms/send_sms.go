package sms

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type SmsAccount struct {
	AccountSid string
	Token      string
}

// Represents the data used in creating an outbound sms message.
// "From" & "To" are required attributes.
// Either a Body or a MediaUrl must also be provided.
// "StatusCallback" and "ApplicationSid" are both optional.
// Visit https://www.twilio.com/docs/api/rest/sending-messages#post for more details.
type Post struct {
	From           string
	To             string
	Body           string
	MediaUrl       string
	StatusCallback string
	ApplicationSid string
}

func validateSmsPost(p Post) error {
	if p.From == "" || p.To == "" {
		return errors.New("Both \"From\" and \"To\" must be set in Post.")
	}
	if p.Body == "" && p.MediaUrl == "" {
		return errors.New("Either \"Body\" or \"MediaUrl\" must be set.")
	}
	return nil
}

// Represents the callback sent everytime the status of the message is updated.
// Visit https://www.twilio.com/docs/api/rest/sending-messages#status-callback-parameter for more detaiils
type Callback struct {
	standardRequest
	MessageStatus string
	ErrorCode     string
}

// Internal function for sending the post request to twilio.
func (act SmsAccount) sendSms(destUrl string, msg Post, resp *Response) error {
	// send post request to twilio
	c := http.Client{}
	v := url.Values{}
	v.Set("To", msg.To)
	v.Set("From", msg.From)
	v.Set("Body", msg.Body)
	v.Set("MediaUrl", msg.MediaUrl)
	v.Set("StatusCallback", msg.StatusCallback)
	v.Set("ApplicationSid", msg.ApplicationSid)
	req, err := http.NewRequest("POST", destUrl, strings.NewReader(v.Encode()))
	req.SetBasicAuth(act.AccountSid, act.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	twilioResp, err := c.Do(req)

	if twilioResp.StatusCode != 201 {
		return errors.New(fmt.Sprintf("Error recieved from Twilio => %s", twilioResp.Status))
	}

	// parse twilio response
	bodyBytes, err := ioutil.ReadAll(twilioResp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while reading json from buffer => %s", err.Error()))
	}
	err = Unmarshal(bodyBytes, resp)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes)))
	}
	return nil
}

// Sends a post request to Twilio to send a sms request.
func (act SmsAccount) Send(p Post) (Response, error) {
	err := validateSmsPost(p)
	if err != nil {
		return Response{}, errors.New(fmt.Sprintf("Error validating sms post => %s.\n", err.Error()))
	}

	// marshal json string
	if err != nil {
		return Response{}, errors.New(fmt.Sprintf("Error encoding json => %s", err.Error()))
	}

	var r Response
	err = act.sendSms(fmt.Sprintf(postUrl, act.AccountSid), p, &r)

	return r, err
}
