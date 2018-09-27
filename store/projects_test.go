package store

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_dbProjectsStore_List_success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub"+
			" database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("id-1", "name-1").
		AddRow("id-2", "name-2")

	wantProjects := []Project{
		{ID: "id-1", Name: "name-1"},
		{ID: "id-2", Name: "name-2"},
	}

	mock.ExpectQuery(`^SELECT (.+) FROM projects`).
		WillReturnRows(rows)

	ps := NewDBProjectsStore(db)

	projects, err := ps.List()

	if assert.NoError(t, err) {
		assert.Equal(t, wantProjects, projects)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_dbProjectsStore_List_error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub"+
			" database connection", err)
	}
	defer db.Close()

	wantError := errors.New("some error")

	mock.ExpectQuery(`^SELECT (.+) FROM projects`).
		WillReturnError(wantError)

	ps := NewDBProjectsStore(db)

	_, err = ps.List()

	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
