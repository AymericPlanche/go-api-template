package integration_tests

import (
	"database/sql"
	"os/exec"

	"github.com/doug-martin/goqu/v9"
)

// Init a connection to the integration test database
func InitDatabase() *goqu.Database {
	cnn, err := sql.Open("pgx", "postgres://postgres:postgres@db:5432/integration")
	if err != nil {
		panic(err)
	}

	return goqu.Dialect("postgres").DB(cnn)
}

// Reset the integration database
func ResetDatabase() {
	cmd := exec.Command("/home/go-user/myapp/build/database/reset.sh", "integration")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
