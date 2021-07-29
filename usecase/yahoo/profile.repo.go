package yahoo

import (
	"context"

	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/entities"
)

///////////////////////////////////////////////////////////
// Repository Interface
///////////////////////////////////////////////////////////

// Reader interface
type Reader interface {
	FindYahooProfiles(ctx context.Context) ([]*entities.YahooProfile, error)
}

// Writer interface
type Writer interface{}

// Repo interface
type Repo interface {
	Reader
	Writer
}
