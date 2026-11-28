package user

import (
	"github.com/LarsSeverson/charter/services/user/internal/app/user/lookup"
	"github.com/LarsSeverson/charter/services/user/internal/app/user/port"
)

type UseCases struct {
	lookup.Lookup
}

func New(deps port.Deps) UseCases {
	useCases := UseCases{
		Lookup: lookup.New(deps),
	}
	return useCases
}
