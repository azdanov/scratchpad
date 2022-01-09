package mysql

import (
	"database/sql"
	"errors"

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
	stmt := `SELECT id, title, content, created, expires FROM scratches
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	s := &models.Scratch{}

	row := m.DB.QueryRow(stmt, id)
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return s, nil
}

func (m *ScratchModel) Latest() ([]*models.Scratch, error) {
	return nil, nil
}
