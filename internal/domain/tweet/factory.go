package tweet

// Factory ...
type Factory struct {
	lazyRepository func() Repository
}

// NewFactory ...
func NewFactory(
	lazyRepository func() Repository,
) *Factory {
	return &Factory{lazyRepository: lazyRepository}
}

// NewTweet ...
func (f *Factory) NewTweet(content string) Tweet {
	return Tweet{
		Content:    content,
		repository: f.lazyRepository(),
		validators: []tweetValidator{
			TweetLengthValidator,
		},
	}
}
