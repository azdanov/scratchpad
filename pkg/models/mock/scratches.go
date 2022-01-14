package mock

import (
	"time"

	"github.com/azdanov/scratchpad/pkg/models"
)

var mockScratch = &models.Scratch{
	ID:      1,
	Title:   "Lorem",
	Content: "Ipsum...",
	Created: time.Now(),
	Expires: time.Now(),
}

type ScratchModel struct{}

func (m *ScratchModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

func (m *ScratchModel) Get(id int) (*models.Scratch, error) {
	switch id {
	case 1:
		return mockScratch, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *ScratchModel) Latest() ([]*models.Scratch, error) {
	return []*models.Scratch{mockScratch}, nil
}
