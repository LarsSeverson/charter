package user

import (
	"github.com/LarsSeverson/charter/services/user/internal/app/user/lookup"
	"github.com/LarsSeverson/charter/services/user/internal/app/user/port"
)

// UseCases embeds each capability, so its use cases are promoted:
// useCases.Get comes from lookup.Lookup.
type UseCases struct {
	lookup.Lookup
}

func New(deps port.Deps) UseCases {
	useCases := UseCases{
		Lookup: lookup.New(deps),
	}
	return useCases
}
