package sms

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	validMmsSid   = "MM800f449d0399ed014aae2bcc0cc2f2ec"
	validMediaSid = "ME85ebf7e12cb821f84b319340424dcb02"
	act           = Account{}
)

func TestValidateMediaSid(t *testing.T) {
	if !validateMediaSid(validMediaSid) {
		t.Error("validateMediaSid failed on a valid sid")
	}
	if validateMediaSid("ME85ebf7e12cb8") {
		t.Error("validateMediaSid validated an invalid sid")
	}
}

func TestValidateMmsSid(t *testing.T) {
	if !validateMmsSid(validMmsSid) {
		t.Error("validateMsSid failed on a valid sid")
	}
	if validateMmsSid("MM800f4") {
		t.Error("validateMmsSid validated an invalid sid")
	}
}

func TestMediaBuild(t *testing.T) {
	var m Media
	err := json.Unmarshal([]byte(testMediaFixtureString), &m)
	if err != nil {
		t.Errorf("Building Media from json string failed with error => %s", err.Error())
	}
	if m.AccountSid != testMediaFixture.AccountSid {
		t.Errorf(
			"Building Media from json string failed to properly set AccountSid, expected %s, found %s",
			testMediaFixture.AccountSid,
			m.AccountSid,
		)
	}
}

func TestMediaListBuild(t *testing.T) {
	var m MediaList
	err := json.Unmarshal([]byte(testMediaListFixtureString), &m)
	if err != nil {
		t.Errorf("Building Media from json string failed with error => %s", err.Error())
	}
	if len(*m.MediaList) != len(*testMediaListFixture.MediaList) {
		t.Errorf(
			"Building Media from json string failed to properly allocate the list of media, expected: %d, found %d",
			len(*testMediaListFixture.MediaList),
			len(*m.MediaList),
		)
	}
}

func TestGetMedia(t *testing.T) {
	_, err := act.GetMedia("sldkfj", "sldkfjls")
	assert.Error(t, err)

	_, err = act.GetMedia(validMmsSid, "sldkfjls")
	assert.Error(t, err)

	// TODO: test HTTP call
}

func TestGetMediaList(t *testing.T) {
	_, err := act.GetMediaList("sldkfj")
	assert.Error(t, err)

	// TODO: test HTTP call
}
