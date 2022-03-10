package domain

import "ddd-with-factory/internal/domain/tweet"

// TweetFactory ...
type TweetFactory interface {
	NewTweet(string) tweet.Tweet
}

// User is an Entity identified by the Username
type User struct {
	ID       int64
	Username string

	tweetFactory TweetFactory
}

// NewTweet ...
func (u *User) NewTweet(content string) (tweet.Tweet, error) {
	tweet := u.tweetFactory.NewTweet(content)

	if err := tweet.Validate(); err != nil {
		return tweet, err
	}

	if err := tweet.Save(); err != nil {
		return tweet, err
	}

	return tweet, nil
}
