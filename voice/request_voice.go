package voice

type standardRequest struct {
	CallSid       string
	AccountSid    string
	From          string
	To            string
	CallStatus    string
	APIVersion    string
	Direction     string
	ForwardedFrom string
	CallerName    string

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
