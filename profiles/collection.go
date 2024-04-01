package profiles

const (
	settingsFile        = "settings.yaml"
	settingsDefaultFile = "settings.default.yaml"
)

type profileCollection[T any] struct {
	Current  string       `yaml:"current"`
	Profiles map[string]T `yaml:"profiles"`
}

func (s profileCollection[T]) getProfile(profileName string) (T, bool) {
	var empty T
	profile, ok := s.Profiles[profileName]
	if !ok {
		return empty, false
	}
	//profile.Profile = profileName
	return profile, ok
}
