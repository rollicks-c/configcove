package profiles

import (
	"fmt"
	"github.com/rollicks-c/secretblend"
)

func (m Manager[T]) inject(profile T) T {
	//secretblend.AddProvider(vt.AsProvider(), "vault://")
	enrichedProfile, err := secretblend.Inject(profile)
	if err != nil {
		panic(fmt.Sprintf("error injecting secrets: %s", err))
	}
	return enrichedProfile
}
