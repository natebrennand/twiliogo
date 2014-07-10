package voice

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"strings"
)

type Update struct {
	URL                  string `json:"url"`
	Method               string `json:"method"`
	Status               string `json:"status"`
	FallbackURL          string `json:"fallback_url"`
	FallbackMethod       string `json:"fallback_method"`
	StatusCallback       string `json:"status_callback"`
	StatusCallbackMethod string `json:"status_callback_method"`
}

func (p Update) GetReader() io.Reader {
	vals := url.Values{}
	if p.URL != "" {
		vals.Set("URL", p.URL)
	}
	if p.Status != "" {
		vals.Set("Status", p.Status)
	}
	if p.StatusCallback != "" {
		vals.Set("StatusCallback", p.StatusCallback)
	}
	if p.StatusCallbackMethod != "" {
		vals.Set("StatusCallbackMethod", p.StatusCallbackMethod)
	}
	if p.Method != "" {
		vals.Set("Method", p.Method)
	}
	if p.FallbackURL != "" {
		vals.Set("FallbackURL", p.FallbackURL)
	}
	if p.FallbackMethod != "" {
		vals.Set("FallbackMethod", p.FallbackMethod)
	}

	return strings.NewReader(vals.Encode())
}

func (p Update) Validate() error {
	if p.URL == "" && p.Method == "" && p.Status == "" {
		return errors.New("URL or Status or Method must all be set.")
	}
	return nil
}

// Internal function for sending the post request to twilio.
func (act VoiceAccount) postUpdate(dest string, msg Update, resp *Call) error {
	// send post request to twilio
	return common.SendPostRequest(dest, msg, act, resp, 200)
}

// Sends a post request to Twilio to modify a call
func (act VoiceAccount) Update(p Update, sid string) (Call, error) {
	var r Call
	err := act.postUpdate(fmt.Sprintf(updateURL, act.AccountSid, string(sid)), p, &r)
	return r, err
}
