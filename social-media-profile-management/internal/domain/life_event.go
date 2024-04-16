package domain

type LifeEvent struct {
	LifeEventID int
	EventName   string
	EventType   EventType
	EventDate   string // Assuming the date as string, could be time.Time for more utility
}

func (le LifeEvent) ID() int {
	return le.LifeEventID
}

type EventType string
