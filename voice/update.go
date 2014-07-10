package voice

import (
	"errors"
	"io"
	"net/url"
	"strings"
)

type Update struct {
	Url                  string `json:"url"`
	Method               string `json:"method"`
	Status               string `json:"status"`
	FallbackUrl          string `json:"fallback_url"`
	FallbackMethod       string `json:"fallback_method"`
	StatusCallback       string `json:"status_callback"`
	StatusCallbackMethod string `json:"status_callback_method"`
}

func (p Update) GetReader() io.Reader {
	vals := url.Values{}
	if p.Url != "" {
		vals.Set("Url", p.Url)
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
	if p.FallbackUrl != "" {
		vals.Set("FallbackUrl", p.FallbackUrl)
	}
	if p.FallbackMethod != "" {
		vals.Set("FallbackMethod", p.FallbackMethod)
	}

	return strings.NewReader(vals.Encode())
}

func (p Update) Validate() error {
	if p.Url == "" && p.Method == "" && p.Status == "" {
		return errors.New("\"Url\" or \"Status\" or \"Method\" must all be set.")
	}
	return nil
}
