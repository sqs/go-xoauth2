package xoauth2

import (
	"testing"
)

func Test_OAuth2String(t *testing.T) {
	expected := "user=alice\001auth=Bearer some-access-token\001\001"
	actual := OAuth2String("alice", "some-access-token")
	if expected != actual {
		t.Error("OAuth2String failed: ", actual, " != ", expected)
	}
}

func Test_XOAuth2String(t *testing.T) {
	expected := "dXNlcj1hbGljZQFhdXRoPUJlYXJlciBzb21lLWFjY2Vzcy10b2tlbgEB"
	actual := XOAuth2String("alice", "some-access-token")
	if expected != actual {
		t.Error("XOAuth2String failed: ", actual, " != ", expected)
	}
}
