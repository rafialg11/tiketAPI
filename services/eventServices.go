package services

import (
	"tiketAPI/configs"
	"tiketAPI/models"
)

// GetEvents is a function to get all events
func GetEvents() (interface{}, error) {
	var events []models.Event
	if err := configs.DB.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

// GetEventByID is a function to get event by ID
func GetEventByID(id string) (interface{}, error) {
	var event models.Event
	if err := configs.DB.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

// CreateEvent is a function to create event
func CreateEvent(event models.Event) (interface{}, error) {
	if err := configs.DB.Save(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

// UpdateEvent is a function to update event
func UpdateEvent(id string, event models.Event) (interface{}, error) {
	var updatedEvent models.Event
	if err := configs.DB.Where("id = ?", id).First(&updatedEvent).Error; err != nil {
		return nil, err
	}
	if err := configs.DB.Model(&updatedEvent).Updates(event).Error; err != nil {
		return nil, err
	}
	return updatedEvent, nil
}

// DeleteEvent is a function to delete event
func DeleteEvent(id string) (interface{}, error) {
	var event models.Event
	if err := configs.DB.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	if err := configs.DB.Delete(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}
