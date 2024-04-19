package domain

type LifeEvent struct {
	LifeEventID int       `json:"life_event_id"`
	EventName   string    `json:"event_name"`
	EventType   EventType `json:"event_type"`
	EventDate   string    `json:"event_date"`
}

func (le LifeEvent) ID() int {
	return le.LifeEventID
}

type EventType string
