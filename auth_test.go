package hutplate

import (
	"testing"
)

func TestAuth_LoginWithValidCredsAndLogout(t *testing.T) {
	hut := SetupLoggingInAndGetHut(t)
	hut.Auth.Login("registered_email@example.com", "1234567890")
	got := hut.Auth.UserId()
	if got != 2 {
		t.Errorf("invalid user id expected 2 got %v", got)
	}

	if ! hut.Auth.Check() {
		t.Error("user logging in failed")
	}


	hut.Auth.Logout()

	if hut.Auth.Check() {
		t.Error("user logout failed")
	}

	if hut.Auth.UserId() != nil {
		t.Error("invalid user id after logout")
	}
}

func TestAuth_LoginWithInvalidCreds(t *testing.T) {
	hut := SetupLoggingInAndGetHut(t)
	hut.Auth.Login("registered_email@example.com", "123456789F")
	if hut.Auth.Check() || hut.Auth.UserId() != nil {
		t.Error("user logged in even after invalid password")
	}

	hut.Auth.Login("not_registered_email@example.com", "1234567890")
	if hut.Auth.Check() || hut.Auth.UserId() != nil {
		t.Error("user logged in even after invalid email")
	}


}

func SetupLoggingInAndGetHut(t *testing.T) Http {
	Config.GetUserWithCred = func(credential interface{}) (interface{}, string) {
		if credential == "registered_email@example.com" {
			// Hash of 1234567890
			hashedPassword := "$2a$14$KYv4f668vfwQy/0/OjRWD.cQtXiK1XPF/XMTZXZxyXmHq5ULjqleu"
			return 2, hashedPassword
		}
		return nil, ""
	}
	return GetARequest(t)
}

