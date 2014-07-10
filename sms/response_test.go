package sms

import (
	"testing"
)

func TestValidateMessageSid(t *testing.T) {
	if MsgSid("SMa2ff4e37c7cb43b49a820f2d7e3ee135").Validate() != true {
		t.Error("Validation failed on valid SMS sid")
	}

	if MsgSid("MMa2ff4e37c7cb43b49a820f2d7e3ee135").Validate() != true {
		t.Error("Validation failed on valid SMS sid")
	}
}
