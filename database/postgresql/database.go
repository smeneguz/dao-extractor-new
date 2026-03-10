package postgresql

import (
	"context"

	"github.com/dao-portal/flux/database"
	"github.com/dao-portal/flux/database/postgresql"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"

	dbtypes "github.com/dao-portal/extractor/database/postgresql/types"
	"github.com/dao-portal/extractor/types"
)

// DB represents a PostgreSQL database that can be used to store information related
// to the DAOs.
type DB struct {
	*postgresql.Database
	defaultPagination *types.Pagination
}

func Builder(
	ctx context.Context,
	id string,
	rawConfig []byte,
) (database.Database, error) {
	db, err := postgresql.DatabaseBuilder(ctx, id, rawConfig)
	if err != nil {
		return nil, err
	}
	postgresqlDatabase := db.(*postgresql.Database)

	var config dbtypes.DBConfig
	err = yaml.Unmarshal(rawConfig, &config)
	if err != nil {
		return nil, err
	}

	var pagination *types.Pagination
	if config.Pagination != nil {
		pagination = types.NewPagination(config.Pagination.Limit, config.Pagination.Offset)
	} else {
		pagination = types.NewPagination(100, 0)
	}

	myDB := &DB{
		Database:          postgresqlDatabase,
		defaultPagination: pagination,
	}
	return myDB, nil
}

// NewDatabaseFromURL creates a new PostgreSQL database from a URL.
func NewDatabaseFromURL(
	logger zerolog.Logger,
	url string,
) (*DB, error) {
	cfg := postgresql.DefaultConfig().WithURL(url)
	db, err := postgresql.NewDatabase(logger, &cfg)
	if err != nil {
		return nil, err
	}
	return &DB{
		Database:          db,
		defaultPagination: types.NewPagination(100, 0),
	}, nil
}
