package courts

import (
	"bomond-tenis/internal/models"
	"bomond-tenis/internal/repository"
)

type CourtService struct {
	repo repository.Courts
}

func NewCourtService(repo repository.Courts) *CourtService {
	return &CourtService{repo: repo}
}

func (c CourtService) GetCourts() ([]models.Court, error) {
	panic("implement me")
}

func (c CourtService) BookCourt(id int) error {
	panic("implement me")
}

func (c CourtService) GetCourtById(id int) (models.Court, error) {
	panic("implement me")
}

func (c CourtService) DeleteCourt(id int) error {
	panic("implement me")
}
