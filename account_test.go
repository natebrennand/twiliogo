package twiliogo

import (
	"os"
	"testing"
)

const (
	testAccountSidString    = "AC5ef8732a3c49700934481addd5ce1659"
	testBadAccountSidString = "invalid token"
	testTokenString         = "token"
)

func TestNewAccount(t *testing.T) {
	act, tkn := testAccountSidString, testTokenString
	a := NewAccount(act, tkn)

	if a.AccountSid != act {
		t.Error("AccountSID not set")
	} else if a.Token != tkn {
		t.Error("Account token not set")
	}
}

func TestInvalidNewAccount(t *testing.T) {
	handledPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				handledPanic = true
			}
		}()
		act, tkn := testBadAccountSidString, testTokenString
		NewAccount(act, tkn)
	}()
	if handledPanic != true {
		t.Error("NewAccount did not panic when passed an invalid AccountSid")
	}
}

// Helper method for setting environment variables on tests.
func setEnvVariables(sid, token string, t *testing.T) {
	err := os.Setenv(twilioAccount, sid)
	if err != nil {
		t.Errorf("Test failed to set environment variable %s to %s\n", twilioAccount, sid)
	}
	err = os.Setenv(twilioToken, token)
	if err != nil {
		t.Errorf("Test failed to set environment variable %s to %s\n", twilioToken, token)
	}
}

func TestNewAccountFromEnv(t *testing.T) {
	setEnvVariables(testAccountSidString, testTokenString, t)
	a := NewAccountFromEnv()
	if a.AccountSid != testAccountSidString {
		t.Error("NewAccountFromEnv failed to properly set the AccountSid")
	}
	if a.Token != testTokenString {
		t.Error("NewAccountFromEnv failed to properly set the Token")
	}
}

func TestInvalidNewAccountFromEnv(t *testing.T) {
	handledPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				handledPanic = true
			}
		}()
		setEnvVariables(testBadAccountSidString, testTokenString, t)
		NewAccountFromEnv()
	}()
	if handledPanic != true {
		t.Error("NewAccount did not panic when passed an invalid AccountSid")
	}
}
