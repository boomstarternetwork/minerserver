package dbstore

import (
	"fmt"
	"os"
	"testing"

	"github.com/boomstarternetwork/bestore"

	_ "github.com/lib/pq"
)

const postgresConnStr = "postgres://testing:password@localhost:5432/testing" +
	"?sslmode=disable"

var s *DBStore

func initTestingStore() {
	var err error
	s, err = New(postgresConnStr, "testing")
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"Failed to create store: %s. Postgres connection string: %s\n",
			err.Error(), postgresConnStr)
		os.Exit(1)
	}
}

func createTestingTables() {
	s.gdb.AutoMigrate(&bestore.Project{})
}

func dropTestingTables() {
	s.gdb.DropTableIfExists(&bestore.Project{})
}

func TestMain(m *testing.M) {
	initTestingStore()
	os.Exit(m.Run())
}
