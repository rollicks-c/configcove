package profiles

import (
	"fmt"
	"github.com/rollicks-c/secretblend"
)

func (m Manager[T]) inject(profileData T) T {
	enrichedData, err := secretblend.Inject(profileData)
	if err != nil {
		panic(fmt.Sprintf("error injecting secrets: %s", err))
	}
	return enrichedData
}
