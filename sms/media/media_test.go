package media

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestValidateMediaSid(t *testing.T) {
	if !validateMediaSid("ME85ebf7e12cb821f84b319340424dcb02") {
		t.Error("validateMediaSid failed on a valid sid")
	}
	if validateMediaSid("ME85ebf7e12cb8") {
		t.Error("validateMediaSid validated an invalid sid")
	}
}

func TestValidateMmsSid(t *testing.T) {
	if !validateMmsSid("MM800f449d0399ed014aae2bcc0cc2f2ec") {
		t.Error("validateMsSid failed on a valid sid")
	}
	if validateMmsSid("MM800f4") {
		t.Error("validateMmsSid validated an invalid sid")
	}
}

func TestMediaBuild(t *testing.T) {
	var m Media
	resp := http.Response{Body: ioutil.NopCloser(strings.NewReader(testMediaFixtureString))}
	err := m.Build(&resp)
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
	resp := http.Response{Body: ioutil.NopCloser(strings.NewReader(testMediaListFixtureString))}
	err := m.Build(&resp)
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
