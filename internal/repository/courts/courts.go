package courts

import (
	"bomond-tenis/internal/models"
	"github.com/jmoiron/sqlx"
)

type CourtRepository struct {
	db *sqlx.DB
}

func NewCourtsRepository(db *sqlx.DB) *CourtRepository {
	return &CourtRepository{db: db}
}

func (c CourtRepository) GetCourts() ([]models.Court, error) {
	panic("implement me")
}

func (c CourtRepository) BookCourt(id int) error {
	panic("implement me")
}

func (c CourtRepository) GetCourtById(id int) (models.Court, error) {
	panic("implement me")
}

func (c CourtRepository) DeleteCourt(id int) error {
	panic("implement me")
}
