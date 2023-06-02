package settings

import (
	portainer "github.com/portainer/portainer/api"
)

const (
	// BucketName represents the name of the bucket where this service stores data.
	BucketName  = "settings"
	settingsKey = "SETTINGS"
)

// Service represents a service for managing environment(endpoint) data.
type Service struct {
	connection portainer.Connection
}

// NewService creates a new instance of a service.
func NewService(connection portainer.Connection) (*Service, error) {
	return &Service{
		connection: connection,
	}, nil
}

func (service *Service) Tx(tx portainer.Transaction) ServiceTx {
	return ServiceTx{
		service: service,
		tx:      tx,
	}
}

// Settings retrieve the settings object.
func (service *Service) Settings() (*portainer.Settings, error) {
	settings := portainer.Settings{ID: 1}

	err := service.connection.GetByID(1, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}

// UpdateSettings persists a Settings object.
func (service *Service) UpdateSettings(settings *portainer.Settings) error {
	db := service.connection.GetDB()
	tx := db.Model(&portainer.Settings{}).Where(portainer.Settings{ID: 1}).Save(settings)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
