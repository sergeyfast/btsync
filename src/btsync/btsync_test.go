package btsync

import (
	"testing"
)

const (
	testHost     = "10.10.10.2"
	testPort     = "8888"
	testUser     = "admin"
	testPassword = "admin"
)

var c = newClient()

func newClient() *Client {
	Debug = true
	return NewClient(testHost, testPort, testUser, testPassword)
}

// First Test
func TestRequestToken(t *testing.T) {
	if !c.RequestToken() {
		t.Error("Can't get token")
	}
}


func TestFolders(t *testing.T) {
	fi, err := c.Folders()
	if err != nil {
		t.Error(err)
	}

	if len(fi.Folders) == 0 {
		t.Error("No sync folders")
	}
}

func TestGenerateSecret(t *testing.T) {
	s, err := c.GenerateSecret()
	if err != nil {
		t.Error(err)
	}

	if s.Secret == "" || s.ROSecret == "" {
		t.Fail()
	}
}
