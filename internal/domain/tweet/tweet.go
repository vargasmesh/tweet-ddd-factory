package tweet

// Tweet ...
type Tweet struct {
	Content string

	repository Repository
	validators []tweetValidator
}

// Repository ...
type Repository interface {
	Save(Tweet) error
}

// Validate ...
func (t *Tweet) Validate() error {
	for _, validator := range t.validators {
		if err := validator(t); err != nil {
			return err
		}
	}

	return nil
}

// Save ...
func (t *Tweet) Save() error {
	return t.repository.Save(*t)
}
