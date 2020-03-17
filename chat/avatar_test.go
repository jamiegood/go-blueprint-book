package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar = new(AuthAvatar)
	var c = new(client)

	url, err := authAvatar.GetAvatarURL(c)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
	}

	if url != "" {
		t.Error("url should be empty but it's not")
	}

	//set the url
	testURL := "http://testing"
	c.userData = map[string]interface{}{"avatar_url": testURL}
	url, err = authAvatar.GetAvatarURL(c)

	if url != testURL {
		t.Error("testurl is not correct")

	}

	//return url, err

}
