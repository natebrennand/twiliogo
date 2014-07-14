package twiliogo

import (
	"fmt"
	"github.com/natebrennand/twiliogo/applications"
	"github.com/natebrennand/twiliogo/conference"
	"github.com/natebrennand/twiliogo/notifications"
	"github.com/natebrennand/twiliogo/recording"
	"github.com/natebrennand/twiliogo/sms"
	"github.com/natebrennand/twiliogo/sms/shortcodes"
	"github.com/natebrennand/twiliogo/transcription"
	"github.com/natebrennand/twiliogo/voice"

	"net/http"
	"os"
	"regexp"
)

const (
	twilioAccount = "TWILIO_ACCOUNT"
	twilioToken   = "TWILIO_TOKEN"
)

var validateAccountSid = regexp.MustCompile("^AC[a-z0-9]{32}$").MatchString

type Account struct {
	AccountSid string
	Token      string
	// redundancy for usability
	Sms            sms.Account
	ShortCode      shortcodes.Account
	Voice          voice.Account
	Recordings     recording.Account
	Transcriptions transcription.Account
	Conferences    conference.Account
	Applications   applications.Account
	Notifications  notifications.Account
}

func NewAccount(sid, token string) Account {
	if !validateAccountSid(sid) {
		panic("Invalid Account Sid")
	}
	return Account{
		AccountSid: sid,
		Token:      token,
		Sms: sms.Account{
			AccountSid: sid,
			Token:      token,
			Client:     http.Client{},
		},
		ShortCode: shortcodes.Account{
			AccountSid: sid,
			Token:      token,
			Client:     http.Client{},
		},
		Voice: voice.Account{
			AccountSid: sid,
			Token:      token,
			Client:     http.Client{},
		},
		Recordings: recording.Account{
			AccountSid: sid,
			Token:      token,
			Client:     http.Client{},
		},
		Transcriptions: transcription.Account{
			AccountSid: sid,
			Token:      token,
			Client:     http.Client{},
		},
		Applications: applications.Account{
			AccountSid: sid,
			Token:      token,
			Client:     http.Client{},
		},
		Conferences: conference.Account{
			AccountSid: sid,
			Token:      token,
			Client:     http.Client{},
		},
		Notifications: notifications.Account{
			AccountSid: sid,
			Token:      token,
			Client:     http.Client{},
		},
	}
}

func NewAccountFromEnv() Account {
	sid := os.Getenv(twilioAccount)
	token := os.Getenv(twilioToken)
	if sid == "" {
		panic(fmt.Sprintf("You must set the environment variable %s.", twilioAccount))
	}
	if token == "" {
		panic(fmt.Sprintf("You must set the environment variable %s.", twilioToken))
	}

	return NewAccount(sid, token)
}
