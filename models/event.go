package models

import (
	"fmt"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    string `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	//later store in database
	query := `Insert into events(name, description, location, dateTime, user_id)
	values(?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return err

}

func (e Event) Update() error {
	query := `Update events set name=?, description=?, location=?, dateTime=?, user_id=? where id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return err
	}
	return err
}

func (e Event) Delete() error {
	query := `Delete from events where id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "select * from events"
	rows, err := db.DB.Query(query)
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
		fmt.Println(event)
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (Event, error) {
	query := "select * from events where id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (e Event) Register(userId int64) error {
	query := `Insert into registrations(event_id, user_id) values(?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := `Delete from registrations where event_id=? and user_id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	return err
}
