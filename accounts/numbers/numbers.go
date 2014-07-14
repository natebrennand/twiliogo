package numbers

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

const (
	localURL    = "https://api.twilio.com/2010-04-01/Accounts/%s/AvailablePhoneNumbers/%s/local.json"    // takes an AccountSid & IsoCountryCode
	tollfreeURL = "https://api.twilio.com/2010-04-01/Accounts/%s/AvailablePhoneNumbers/%s/tollfree.json" // takes an AccountSid & IsoCountryCode
	mobileURL   = "https://api.twilio.com/2010-04-01/Accounts/%s/AvailablePhoneNumbers/%s/mobile.json"   // takes an AccountSid & IsoCountryCode
	allURL      = "https://api.twilio.com/2010-04-01/Accounts/%s/AvailablePhoneNumbers.json"             // takes an AccountSid
)

type Account struct {
	AccountSid string
	Token      string
	Client     http.Client
}

func (act Account) GetSid() string {
	return act.AccountSid
}
func (act Account) GetToken() string {
	return act.Token
}
func (act Account) GetClient() http.Client {
	return act.Client
}

var validateIsoCountry = regexp.MustCompile(`^(AU|AT|BH|BE|BR|BG|CA|CY|CZ|DK|DO|SV|EE|FI|FR|GR|HK|IE|IL|IT|JP|LV|LT|LU|MT|MX|NL|NO|NZ|PE|PL|PT|PR|RO|SK|ZA|ES|SE|CH|GB|US)$`).MatchString

// Can filter Local
// https://www.twilio.com/docs/api/rest/available-phone-numbers#local-instance-properties
type Number struct {
	FriendlyName string            `json:"friendly_name"`
	PhoneNumber  string            `json:"phone_number"`
	Lata         string            `json:"lata"`
	RateCenter   string            `json:"rate_center"`
	Latitude     *common.JSONFloat `json:"latitude"`
	Longitude    *common.JSONFloat `json:"longitude"`
	Region       string            `json:"region"`
	PostalCode   string            `json:"postalCode"`
	IsoCountry   string            `json:"iso_country"`
	Capabilities struct {
		Voice bool `json:"voice,string"`
		SMS   bool `json:",string"`
		MMS   bool `json:",string"`
	} `json:"capabilities"`
}

type NumberList struct {
	URI                   string    `json:"uri"`
	AvailablePhoneNumbers *[]Number `json:"available_phone_numbers"`
}

type localSearchFilter struct {
	AreaCode     string
	Contains     string
	NearNumber   string
	InPostalCode string
	InRegion     string
	InRateCenter string
	SmsEnabled   *bool
	MmsEnabled   *bool
	VoiceEnabled *bool
	Distance     *int64
	InLata       *int64
	NearLat      *float64
	NearLong     *float64
}

func (l localSearchFilter) getQueryString() string {
	v := url.Values{}
	if l.AreaCode != "" {
		v.Add("AreaCode", l.AreaCode)
	}
	if l.Contains != "" {
		v.Add("Contains", l.Contains)
	}
	if l.NearNumber != "" {
		v.Add("NearNumber", l.NearNumber)
	}
	if l.InPostalCode != "" {
		v.Add("InPostalCode", l.InPostalCode)
	}
	if l.InRegion != "" {
		v.Add("InRegion", l.InRegion)
	}
	if l.InRateCenter != "" {
		v.Add("InRateCenter", l.InRateCenter)
	}
	if l.SmsEnabled != nil {
		v.Add("SmsEnabled", strconv.FormatBool(*l.SmsEnabled))
	}
	if l.MmsEnabled != nil {
		v.Add("MmsEnabled", strconv.FormatBool(*l.MmsEnabled))
	}
	if l.VoiceEnabled != nil {
		v.Add("VoiceEnabled", strconv.FormatBool(*l.VoiceEnabled))
	}
	if l.Distance != nil {
		v.Add("Distance", strconv.FormatInt(*l.Distance, 10))
	}
	if l.InLata != nil {
		v.Add("InLata", strconv.FormatInt(*l.InLata, 10))
	}
	if l.NearLat != nil && l.NearLong != nil {
		v.Add("NearLatLong", strconv.FormatFloat(*l.NearLat, 'f', 6, 64)+","+strconv.FormatFloat(*l.NearLong, 'f', 6, 64))
	}
	return v.Encode()
}

func (act Account) GetLocalNumbers(l localSearchFilter, iso string) (NumberList, error) {
	var nl NumberList
	if !validateIsoCountry(iso) {
		return nl, errors.New("Invalid country ISO")
	}
	err := common.SendGetRequest(fmt.Sprintf(localURL, act.AccountSid, iso)+l.getQueryString(), act, &nl)
	return nl, err
}

type tollfreeSearchFilter struct {
	AreaCode string
	Contains string
}

func (l tollfreeSearchFilter) getQueryString() string {
	v := url.Values{}
	if l.AreaCode != "" {
		v.Add("AreaCode", l.AreaCode)
	}
	if l.Contains != "" {
		v.Add("Contains", l.Contains)
	}
	return v.Encode()
}

func (act Account) GetTollfreeNumbers(l tollfreeSearchFilter, iso string) (NumberList, error) {
	var nl NumberList
	if !validateIsoCountry(iso) {
		return nl, errors.New("Invalid country ISO")
	}
	err := common.SendGetRequest(fmt.Sprintf(tollfreeURL, act.AccountSid, iso)+l.getQueryString(), act, &nl)
	return nl, err
}

type mobileSearchFilter struct {
	Contains     string
	SmsEnabled   *bool
	MmsEnabled   *bool
	VoiceEnabled *bool
}

func (l mobileSearchFilter) getQueryString() string {
	v := url.Values{}
	if l.Contains != "" {
		v.Add("Contains", l.Contains)
	}
	if l.SmsEnabled != nil {
		v.Add("SmsEnabled", strconv.FormatBool(*l.SmsEnabled))
	}
	if l.MmsEnabled != nil {
		v.Add("MmsEnabled", strconv.FormatBool(*l.MmsEnabled))
	}
	if l.VoiceEnabled != nil {
		v.Add("VoiceEnabled", strconv.FormatBool(*l.VoiceEnabled))
	}
	return v.Encode()
}

func (act Account) GetMobileNumbers(l mobileSearchFilter, iso string) (NumberList, error) {
	var nl NumberList
	if !validateIsoCountry(iso) {
		return nl, errors.New("Invalid country ISO")
	}
	err := common.SendGetRequest(fmt.Sprintf(mobileURL, act.AccountSid, iso)+l.getQueryString(), act, &nl)
	return nl, err
}

type CountryList struct {
	common.ListResponseCore
	Countries struct {
		CountryCode    string `json:"country_code"`
		Country        string `json:"country"`
		URI            string `json:"uri"`
		SubresourceURL struct {
			Local    string `json:"local"`
			Tollfree string `json:"toll_free"`
			Mobile   string `json:"mobile"`
		} `json:"subresource_uris"`
	} `json:"countries"`
}

func (act Account) GetAllNumbers() (CountryList, error) {
	var cl CountryList
	err := common.SendGetRequest(fmt.Sprintf(allURL, act.AccountSid), act, &cl)
	return cl, err
}
