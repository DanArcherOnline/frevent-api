package models

import (
	"time"

	"github.com/frevent/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	saveEventQuery := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	preparedSaveEventQuery, err := db.DB.Prepare(saveEventQuery)
	if err != nil {
		return err
	}
	defer preparedSaveEventQuery.Close()
	result, err := preparedSaveEventQuery.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	getEventsQuery := "SELECT * FROM events"
	rows, err := db.DB.Query(getEventsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	getEventQuery := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(getEventQuery, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Update() error {
	updateEventQuery := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	preparedUpdateEventQuery, err := db.DB.Prepare(updateEventQuery)
	if err != nil {
		return err
	}
	defer preparedUpdateEventQuery.Close()

	_, err = preparedUpdateEventQuery.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err

}

func (e *Event) Delete() error {
	deleteEventQuery := `DELETE FROM events WHERE id = ?`
	preparedDeleteEventQuery, err := db.DB.Prepare(deleteEventQuery)
	if err != nil {
		return err
	}
	defer preparedDeleteEventQuery.Close()
	_, err = preparedDeleteEventQuery.Exec(e.ID)
	return err
}
