package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
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

func (p Post) ToFormEncoded() io.Reader {
	v := url.Values{}
	v.Set("To", p.To)
	v.Set("From", p.From)
	if p.Body != "" {
		v.Set("Body", p.Body)
	}
	if p.MediaUrl != "" {
		v.Set("MediaUrl", p.MediaUrl)
	}
	if p.StatusCallback != "" {
		v.Set("StatusCallback", p.StatusCallback)
	}
	if p.ApplicationSid != "" {
		v.Set("ApplicationSid", p.ApplicationSid)
	}
	return strings.NewReader(v.Encode())
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

// Internal function for sending the post request to twilio.
func (act SmsAccount) sendSms(destUrl string, msg Post, resp *Response) error {
	// send post request to twilio
	c := http.Client{}
	req, err := http.NewRequest("POST", destUrl, msg.ToFormEncoded())

	req.SetBasicAuth(act.AccountSid, act.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	twilioResp, err := c.Do(req)

	if twilioResp.StatusCode != 201 {
		var (
			twilioErr common.Error
			buf       bytes.Buffer
		)
		_, err := buf.ReadFrom(req.Body)
		if err != nil {
			return errors.New(fmt.Sprintf("Twilio error encountered, failure while reading body => %s", err.Error()))
		}

		err = json.Unmarshal(buf.Bytes(), &twilioErr)
		if err != nil {
			return errors.New(fmt.Sprintf("Twilio error encountered, failure while parsing => %s", err.Error()))
		}
		return twilioErr
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
