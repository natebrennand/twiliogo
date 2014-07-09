package sms

type standardRequest struct {
	MessageSid string
	SmsSid     string
	AccountSid string
	From       string
	To         string
	Body       string
	NumMedia   string

	// Only sent when media is associated wit the message.
	MediaContentType1 string `json=",omitempty"`
	MediaContentType2 string `json=",omitempty"`
	MediaContentType3 string `json=",omitempty"`
	MediaContentType4 string `json=",omitempty"`
	MediaUrl1         string `json=",omitempty"`
	MediaUrl2         string `json=",omitempty"`
	MediaUrl3         string `json=",omitempty"`
	MediaUrl4         string `json=",omitempty"`

	// Only sent when Twilio can look up the geographic data.
	FromCity   string `json=",omitempty"`
	FromState  string `json=",omitempty"`
	FromZip    string `json=",omitempty"`
	FromCountr string `json=",omitempty"`
	ToCity     string `json=",omitempty"`
	ToState    string `json=",omitempty"`
	ToZip      string `json=",omitempty"`
	ToCountry  string `json=",omitempty"`
}
