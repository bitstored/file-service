package repo

import (
	"context"
	cdb "github.com/bitstored/repository/couchbaserepo"
)

type Repository struct {
	next cdb.Repository
}
