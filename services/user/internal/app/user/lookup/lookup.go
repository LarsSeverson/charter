package lookup

import "github.com/LarsSeverson/charter/services/user/internal/app/user/port"

type Lookup struct {
	Get Get
}

func New(deps port.Deps) Lookup {
	get := Get{
		users: deps.Users,
	}

	lookup := Lookup{
		Get: get,
	}
	return lookup
}
