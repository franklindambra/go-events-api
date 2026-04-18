package models

import "time"

//binding struct tag works with the ShouldBindJSON method in main, err genreated if data missing required field
type Event struct {
	ID int
	Name string `binding:"required"`
	Descrtiption string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID  int
}

var events []Event = []Event{}


func (event Event)Save(){
		events = append(events, event)
}


func GetAllEvents()[]Event{
		return events
}
