package sms

import (
	"testing"
)

func TestValidateMessageSid(t *testing.T) {
	if validateSmsSid("SMa2ff4e37c7cb43b49a820f2d7e3ee135") != true {
		t.Error("Validation failed on valid SMS sid")
	}

	if validateSmsSid("MMa2ff4e37c7cb43b49a820f2d7e3ee135") != true {
		t.Error("Validation failed on valid SMS sid")
	}
}
