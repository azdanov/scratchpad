package mysql

import (
	"database/sql"

	"github.com/azdanov/scratchpad/pkg/models"
)

type ScratchModel struct {
	DB *sql.DB
}

func (m *ScratchModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO scratches (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ScratchModel) Get(id int) (*models.Scratch, error) {
	return nil, nil
}

func (m *ScratchModel) Latest() ([]*models.Scratch, error) {
	return nil, nil
}
