package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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

// Represents the callback sent everytime the status of the message is updated.
// Visit https://www.twilio.com/docs/api/rest/sending-messages#status-callback-parameter for more detaiils
type Callback struct {
	standardRequest
	MessageStatus string
	ErrorCode     string
}

// Internal function for sending the post request to twilio.
func sendSms(url string, msg io.Reader, resp *Response) error {
	// send post request to twilio
	c := &http.Client{}
	twilioResp, err := c.Post(url, "application/json", msg)
	if twilioResp.StatusCode != 200 {
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
func Send(p Post) (Response, error) {
	// marshal json string
	body, err := json.Marshal(p)
	if err != nil {
		return Response{}, errors.New(fmt.Sprintf("Error encoding json => %s", err.Error()))
	}

	var r Response
	jReader := bytes.NewBuffer(body)
	err = sendSms(postUrl, jReader, &r)

	return r, nil
}
