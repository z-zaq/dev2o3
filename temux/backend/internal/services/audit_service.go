package services

import (
	"temux/internal/models"
	"temux/internal/repository"
)

type AuditService struct {
	AuditRepo *repository.AuditRepository
}

func (s *AuditService) Log(
	userID int,
	action string,
	entity string,
	entityID int,
	ip string,
	userAgent string,
) error {

	return s.AuditRepo.Create(
		&models.AuditLog{
			UserID:    userID,
			Action:    action,
			Entity:    entity,
			EntityID:  entityID,
			IPAddress: ip,
			UserAgent: userAgent,
		},
	)
}
