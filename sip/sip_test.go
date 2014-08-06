package sip

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	validDomainSid     = "SD098e7b11c00d0ba152b1d3f084c4b776"
	validCredentialSid = "SC098e7b11c00d0ba152b1d3f084c4b776"
	validIPSid         = "IP098e7b11c00d0ba152b1d3f084c4b776"
	validALSid         = "AL95a47094615fe05b7c17e62a7877836c"
	validCLSid         = "CL32a3c49700934481addd5ce1659f04d2"
)

func TestValidateControlListUpdateSuccess(t *testing.T) {
	u := ControlListUpdate{IPAccessControlListSid: validALSid}
	if nil != u.Validate() {
		t.Error("Validation of valid control list update failed.")
	}
}

func TestValidateControlListUpdateFailure(t *testing.T) {
	u := ControlListUpdate{}
	if nil == u.Validate() {
		t.Error("Validation of control list update missing sid was supposed to fail.")
	}
}

func TestListDomains(t *testing.T) {
	// TODO: test HTTP request
}

func TestCreateDomain(t *testing.T) {
	// TODO: test HTTP request
}

func TestGetDomain(t *testing.T) {
	act := Account{}
	_, err := act.GetDomain("heyeyeyye")
	assert.Error(t, err)

	// TODO: test HTTP request
}

func TestUpdateDomain(t *testing.T) {
	act := Account{}
	n, err := ioutil.ReadAll(testValidNewDomainFixture.GetReader())
	assert.Nil(t, err)
	form := string(n)

	assert.Contains(t, form, "DomainName=dunder-mifflin-scranton.sip.twilio.com")
	assert.Contains(t, form, "FriendlyName=Scranton+Office")
	assert.Contains(t, form, "VoiceUrl=https%3A%2F%2Fdundermifflin.example.com%2Ftwilio%2Fapp.php")
	assert.Contains(t, form, "VoiceMethod=POST")
	assert.Contains(t, form, "VoiceFallbackMethod=POST")
	assert.Contains(t, form, "VoiceStatusCallbackMethod=POST")

	_, err = act.UpdateDomain(testValidNewDomainFixture, "heyeyeyye")
	assert.Error(t, err)

	// TODO: test HTTP request
}

func TestDeleteDomain(t *testing.T) {
	act := Account{}
	err := act.DeleteDomain("heyeyeyye")
	assert.Error(t, err)

	// TODO: test HTTP request

}

func TestGetMapping(t *testing.T) {
	act := Account{}
	_, err := act.GetMapping("gopher", "topher")
	assert.Error(t, err)

	// TODO: test HTTP request
}

func TestAddMapping(t *testing.T) {
	act := Account{}
	_, err := act.AddMapping(ControlListUpdate{IPAccessControlListSid: validALSid}, "heyeyeyye")
	assert.Error(t, err)

	// TODO: test HTTP request
}

func TestDeleteMapping(t *testing.T) {
	act := Account{}
	err := act.DeleteMapping("gopher", "topher")
	assert.Error(t, err)

	// TODO: test HTTP request
}

func TestListCredentialMappings(t *testing.T) {
	act := Account{}
	_, err := act.ListCredentialMappings("heyeyeyye")
	assert.Error(t, err)

	// TODO: test HTTP request
}

func TestValidateCredentialListUpdateSuccess(t *testing.T) {
	u := CredentialListUpdate{CredentialListSid: validCLSid}
	if nil != u.Validate() {
		t.Error("Validation of valid credential list update failed.")
	}
}

func TestValidateCredentialListUpdateFailure(t *testing.T) {
	u := CredentialListUpdate{}
	if nil == u.Validate() {
		t.Error("Validation of credential list update missing sid was supposed to fail.")
	}
}

func TestAddCredentialMapping(t *testing.T) {
	act := Account{}
	_, err := act.AddCredentialMapping(CredentialListUpdate{CredentialListSid: validCLSid}, "heyeyeyye")
	assert.Error(t, err)

	// TODO: test HTTP request
}

func TestDeleteCredentialMapping(t *testing.T) {
	act := Account{}
	err := act.DeleteCredentialMapping("gopher", "topher")
	assert.Error(t, err)

	// TODO: test HTTP request
}
