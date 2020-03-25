package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

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

//TestGravatarAvatar ...
func TestGravatarAvatar(t *testing.T) {
	var gravAvatar = new(GravatarAvatar)
	var c = new(client)
	c.userData = map[string]interface{}{"userid": "b1522df375addd5bbcadf1edac1a3671"}

	url, err := gravAvatar.GetAvatarURL(c)

	if err != nil {
		t.Error("gravAvatar.GetAvatarURL should not return an error", err)
	}

	if url != "//www.gravatar.com/avatar/b1522df375addd5bbcadf1edac1a3671" {
		t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {

	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer os.Remove(filename)
	//create instance of FileSysmteAvatar
	filesystemAvatar := new(FileSystemAvatar)
	var c = new(client)
	c.userData = map[string]interface{}{"userid": "b1522df375addd5bbcadf1edac1a3671"}

	url, err := filesystemAvatar.GetAvatarURL(c)
	if err != nil {
		t.Errorf("filesystemAvatar.GetAvatarURL should not return error")
	}

	if url != "/avatars/abc.jpg" {
		t.Errorf("filesystemAvatar.GetAvatarURL returned the wrong url")
	}

}
