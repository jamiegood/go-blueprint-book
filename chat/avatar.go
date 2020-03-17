package main

import (
	"crypto/md5"
	"errors"
)

// ErrNoAvatarURL is the error that is returned when the
// Avatar instance is unable to provide an avatar URL
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar url")

// Avatar represents types capable of representing user profile images
type Avatar interface {
	// takes in a client as we want to know the user to get, i.e. the gravatar url
	// returns a url
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar ...
type AuthAvatar struct{}

// UseAuthAvatar used in main
var UseAuthAvatar AuthAvatar

// GetAvatarURL returns an avatar url
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {

	if url, ok := c.userData["avatar_url"]; ok {
		return url.(string), nil

	}
	//return url.(string), ErrNoAvatarURL
	//return c.userData["avatar_url"], ErrNoAvatarURL
	return "", ErrNoAvatarURL
}

// GravatarAvatar ...
type GravatarAvatar struct{}

// GetAvatarURL ...
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {

	if email, ok := c.userData["Email"]; ok {
		m := md5.New()
	}
	return "test", nil
}
