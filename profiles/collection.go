package profiles

const (
	settingsFile        = "settings.yaml"
	settingsDefaultFile = "settings.default.yaml"
)

type profileCollection[T any] struct {
	Current  string                `yaml:"current"`
	Profiles map[string]Profile[T] `yaml:"profiles"`
}

func (s profileCollection[T]) getProfile(profileName string) (Profile[T], bool) {
	profile, ok := s.Profiles[profileName]
	if !ok {
		return *new(Profile[T]), false
	}
	return profile, ok
}
