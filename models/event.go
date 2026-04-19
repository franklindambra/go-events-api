package models

import (
	"time"

	"example.com/rest-api/db"
)

//binding struct tag works with the ShouldBindJSON method in main, err genreated if data missing required field
type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time
	UserID  int
}



func (event Event)Save()error{

	query := `
	INSERT INTO event(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	event.ID = id
	return err

}


func GetAllEvents()([]Event, error){
		query := "SELECT * FROM event"
		rows, err := db.DB.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var events []Event
		//.Next returns true as long as rows are left and then false thereafter
		for rows.Next() {
			var event Event
			err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID )
			if err != nil {
				return nil, err
			}
			events = append(events, event )
		}

		return events, nil

}


func GetEventByID(eventId int64)(*Event, error){

	query := `SELECT * FROM event WHERE id = ?`

	row := db.DB.QueryRow(query, eventId)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID )

	if err != nil {
	 return nil, err
	}


	return &event, nil

}


func (event Event) Update() error{
	query := `
	UPDATE event
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.
		Location, event.DateTime, event.ID)

	return err

}


func (event Event) Delete() error{
	query := `
	DELETE FROM event
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	return err

}
