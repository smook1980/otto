package directory

import (
	"github.com/hashicorp/otto/appfile"
)

// InfraState is used to track the state of an infrastructure.
//
// This is required because the state of an infrastructure isn't binary.
// It can be not created at all, partially created, or fully created. We
// need to represent this in the directory.
type InfraState byte

const (
	InfraStateInvalid InfraState = 0
	InfraStatePartial InfraState = iota
	InfraStateReady
)

// Infra represents the data stored in the directory service about
// Infrastructures.
type Infra struct {
	// State is the state of this infrastructure. This is important since
	// it is possible for there to be a partial state. If we're in a
	// partial state then deploys and such can't go through yet.
	State InfraState

	// Outputs are the output data from the infrastructure step.
	// This is an opaque blob that is dependent on each infrastructure
	// type. Please refer to docs of a specific infra to learn more about
	// what values are here.
	Outputs map[string]string `json:"outputs"`
}

// InfraId generates the ID to use to reference an infrastructure in
// the directory. This is guaranteed to be a unique and consistent ID
// for the given inputs and should be used by multiple plugins (infra,
// apps, etc.) to generate a consistent ID for querying data.
func InfraId(infra *appfile.Infrastructure) string {
	// TODO: Make a more unique ID.
	return infra.Name
}
