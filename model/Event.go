package model

import (
	errors "github.com/lin-sel/event-manager/error"
	"github.com/lin-sel/event-manager/util"
)

// Event used to define event data
type Event struct {
	Base
	EventName        *string
	EventContactNo   *string
	EventDescription *string
	EventStartDate   *string
	EventEndDate     *string
}

// CreateEvent used to create new Event Object
func CreateEvent() *Event {
	return &Event{}
}

// ValidateEvent used to check Event valid or not
func (event *Event) ValidateEvent() error {
	if util.IsStringNil(event.EventName) {
		return errors.NewValidationError("event name must be specified")
	}
	if util.IsStringNil(event.EventContactNo) {
		return errors.NewValidationError("event Contact must be specified")
	}
	if util.IsStringNil(event.EventStartDate) {
		return errors.NewValidationError("event Start date must be specified")
	}
	return nil
}
