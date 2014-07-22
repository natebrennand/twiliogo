package numbers

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/url"
	"regexp"
	"strconv"
)

var numbers = struct {
	local, tollfree, mobile, all string
}{
	local:    "/2010-04-01/Accounts/%s/AvailablePhoneNumbers/%s/local.json",    // takes an AccountSid & IsoCountryCode
	tollfree: "/2010-04-01/Accounts/%s/AvailablePhoneNumbers/%s/tollfree.json", // takes an AccountSid & IsoCountryCode
	mobile:   "/2010-04-01/Accounts/%s/AvailablePhoneNumbers/%s/mobile.json",   // takes an AccountSid & IsoCountryCode
	all:      "/2010-04-01/Accounts/%s/AvailablePhoneNumbers.json",             // takes an AccountSid
}

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

var validateIsoCountry = regexp.MustCompile(`^(AU|AT|BH|BE|BR|BG|CA|CY|CZ|DK|DO|SV|EE|FI|FR|GR|HK|IE|IL|IT|JP|LV|LT|LU|MT|MX|NL|NO|NZ|PE|PL|PT|PR|RO|SK|ZA|ES|SE|CH|GB|US)$`).MatchString

// Number represents the data associated with an available number resource.
//
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

// NumberList represents a list of numbers that matched a query and are avaiable.
type NumberList struct {
	URI                   string    `json:"uri"`
	AvailablePhoneNumbers *[]Number `json:"available_phone_numbers"`
	act                   *Account
}

// LocalSearchFilter allows detailed filtering while searching for a number to purchase.
//
// https://www.twilio.com/docs/api/rest/available-phone-numbers#local-get-advanced-filters
type LocalSearchFilter struct {
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

func (l LocalSearchFilter) getQueryString() string {
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

// GetLocalNumbers searches for local numbers that match the given filter.
//
// https://www.twilio.com/docs/api/rest/available-phone-numbers#local-instance
func (act Account) GetLocalNumbers(l LocalSearchFilter, iso string) (NumberList, error) {
	var nl NumberList
	if !validateIsoCountry(iso) {
		return nl, errors.New("Invalid country ISO")
	}
	err := common.SendGetRequest(fmt.Sprintf(numbers.local, act.AccountSid, iso)+l.getQueryString(), act, &nl)
	nl.act = &act
	return nl, err
}

// TollfreeSearchFilter allows filtereing for tollfree number searches.
type TollfreeSearchFilter struct {
	AreaCode string
	Contains string
}

func (l TollfreeSearchFilter) getQueryString() string {
	v := url.Values{}
	if l.AreaCode != "" {
		v.Add("AreaCode", l.AreaCode)
	}
	if l.Contains != "" {
		v.Add("Contains", l.Contains)
	}
	return v.Encode()
}

// GetTollfreeNumbers searches for avaiable tollfree numbers that match a given filter.
//
// https://www.twilio.com/docs/api/rest/available-phone-numbers#toll-free
func (act Account) GetTollfreeNumbers(l TollfreeSearchFilter, iso string) (NumberList, error) {
	var nl NumberList
	if !validateIsoCountry(iso) {
		return nl, errors.New("Invalid country ISO")
	}
	err := common.SendGetRequest(fmt.Sprintf(numbers.tollfree, act.AccountSid, iso)+l.getQueryString(), act, &nl)
	nl.act = &act
	return nl, err
}

// MobileSearchFilter allow filtering while searching for a mobile number.
type MobileSearchFilter struct {
	Contains     string
	SmsEnabled   *bool
	MmsEnabled   *bool
	VoiceEnabled *bool
}

func (l MobileSearchFilter) getQueryString() string {
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

// GetMobileNumbers returns all available mobile numbers matching the given filter.
//
// https://www.twilio.com/docs/api/rest/available-phone-numbers#mobile
func (act Account) GetMobileNumbers(l MobileSearchFilter, iso string) (NumberList, error) {
	var nl NumberList
	if !validateIsoCountry(iso) {
		return nl, errors.New("Invalid country ISO")
	}
	err := common.SendGetRequest(fmt.Sprintf(numbers.mobile, act.AccountSid, iso)+l.getQueryString(), act, &nl)
	nl.act = &act
	return nl, err
}

// CountryList lists all countries that currently have numbers available for purchase.
//
// https://www.twilio.com/docs/api/rest/available-phone-numbers#countries
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

// GetGlobalNumbers returns a list of all resources for sarching for numbers in other
// countries.
func (act Account) GetGlobalNumbers() (CountryList, error) {
	var cl CountryList
	err := common.SendGetRequest(fmt.Sprintf(numbers.all, act.AccountSid), act, &cl)
	return cl, err
}
