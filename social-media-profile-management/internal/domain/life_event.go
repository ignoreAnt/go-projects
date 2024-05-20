package domain

type LifeEvent struct {
	ID        int
	EventName string
	EventType EventType
	EventDate string // Assuming the date as string, could be time.Time for more utility
}

type EventType string
