package sms

import (
	"encoding/json"
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

func TestBuildMessageSuccess(t *testing.T) {
	var msg Message
	err := json.Unmarshal([]byte(testSmsResponseFixtureString), &msg)
	if err != nil {
		t.Errorf("Building Message from json string failed with error => %s", err.Error())
	}
	if msg.AccountSid != testSmsResponseFixtureAccountSid {
		t.Error("Building Message from json string failed, AccountSid not properly set.")
	}
}

func TestBuildMessageFailure(t *testing.T) {
	var msg Message
	err := json.Unmarshal([]byte(testSmsResponseFixtureString[0:35]), &msg)
	if err == nil {
		t.Error("Building Message from json string should've failed")
	}
}

func TestBuildMessageListSuccess(t *testing.T) {
	var ml MessageList
	err := json.Unmarshal([]byte(testSmsListFixtureString), &ml)
	if err != nil {
		t.Errorf("Building Message List from json string failed with error => %s", err.Error())
	}
	if ml.Total != testSmsListFixture.Total {
		t.Error("Building Message from json string failed, total not properly set.")
	}
}

func TestBuildMessageListFailure(t *testing.T) {
	var ml MessageList
	err := json.Unmarshal([]byte(testSmsListFixtureString[0:35]), &ml)
	if err == nil {
		t.Error("Building Message from json string should've failed")
	}
}
