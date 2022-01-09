package mysql

import (
	"database/sql"
	"github.com/azdanov/scratchpad/pkg/models"
)

type ScratchModel struct {
	DB *sql.DB
}

func (m *ScratchModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *ScratchModel) Get(id int) (*models.Scratch, error) {
	return nil, nil
}

func (m *ScratchModel) Latest() ([]*models.Scratch, error) {
	return nil, nil
}
