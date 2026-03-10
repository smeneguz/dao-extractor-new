package postgresql_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	fluxsuite "github.com/dao-portal/flux/database/suite"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"

	"github.com/dao-portal/extractor/database/postgresql"
)

type DbTestSuite struct {
	fluxsuite.Suite
	database *postgresql.DB
}

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

func (suite *DbTestSuite) SetupSuite() {
	parserDb, err := postgresql.NewDatabaseFromURL(
		zerolog.New(os.Stdout),
		"postgres://dao-portal:password@localhost:6432/dao-portal?sslmode=disable&search_path=public",
	)
	suite.Require().NoError(err)
	dbCfg := parserDb.Cfg

	// Delete the public schema
	_, err = parserDb.SQL.Exec(fmt.Sprintf(`DROP SCHEMA %s CASCADE;`, dbCfg.GetSchema()))
	suite.Require().NoError(err)

	// Create the schema
	_, err = parserDb.SQL.Exec(fmt.Sprintf(`CREATE SCHEMA %s;`, dbCfg.GetSchema()))
	suite.Require().NoError(err)

	dirPath := "schema"
	dir, err := os.ReadDir(dirPath)
	suite.Require().NoError(err)

	for _, fileInfo := range dir {
		if !strings.HasSuffix(fileInfo.Name(), ".sql") {
			continue
		}

		file, err := os.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
		suite.Require().NoError(err)

		_, err = parserDb.SQL.Exec(string(file))
		if err != nil {
			err = fmt.Errorf("failed to execute %s: %w", fileInfo.Name(), err)
		}
		suite.Require().NoError(err)
	}

	// Create the truncate function
	stmt := fmt.Sprintf(`
CREATE OR REPLACE FUNCTION truncate_tables(username IN VARCHAR) RETURNS void AS $$
DECLARE
    table_statements CURSOR FOR
        SELECT tablename FROM pg_tables
        WHERE tableowner = username AND schemaname = '%[1]s';
    sequence_statements CURSOR FOR
        SELECT sequence_name FROM information_schema.sequences
        WHERE sequence_schema = '%[1]s';
BEGIN
    FOR stmt IN table_statements LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
    END LOOP;
    
    FOR seq IN sequence_statements LOOP
        EXECUTE 'ALTER SEQUENCE ' || quote_ident(seq.sequence_name) || ' RESTART WITH 1;';
    END LOOP;
END;
$$ LANGUAGE plpgsql;`, dbCfg.GetSchema())
	_, err = parserDb.SQL.Exec(stmt)
	suite.Require().NoError(err)

	suite.database = parserDb
	suite.InitDB(parserDb)

	suite.WithBeforeTestHook(func() {
		suite.T().Log("Truncating tables")
		_, err := suite.database.SQL.Exec(`SELECT truncate_tables('dao-portal')`)
		suite.Require().NoError(err)
	})
}
