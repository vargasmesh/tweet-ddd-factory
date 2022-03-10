package service

import (
	"ddd-with-factory/internal/domain/tweet"
	user "ddd-with-factory/internal/domain/user"
)

// UserService is a Service Application
type UserService struct {
	userDataMapper UserDataMapper
}

// UserTweetMapper ...
type UserTweetMapper interface {
	AssociateUserTweet(user.User, tweet.Tweet) error
}

// UserDataMapper ...
type UserDataMapper interface {
	UserTweetMapper
	GetUserFromUsername(username string) (user.User, error)
}

// NewTweet ...
func (u *UserService) NewTweet(username string, content string) error {
	userModel, err := u.userDataMapper.GetUserFromUsername(username)
	if err != nil {
		return err
	}

	tweet, err := userModel.NewTweet(content)
	if err != nil {
		return err
	}

	return u.userDataMapper.AssociateUserTweet(userModel, tweet)
}
