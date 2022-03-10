package tweet

import (
	"errors"
	"fmt"
)

type tweetValidator func(*Tweet) error

// TweetMaxCharacters ...
const TweetMaxCharacters = 140

// ErrCharactersExceeded ...
var ErrCharactersExceeded = errors.New("max characters exceeded")

// TweetLengthValidator ...
var TweetLengthValidator = func(t *Tweet) error {
	tweetCharacters := len(t.Content)

	if tweetCharacters > TweetMaxCharacters {
		return fmt.Errorf("%w: got: %d", ErrCharactersExceeded, tweetCharacters)
	}

	return nil
}
