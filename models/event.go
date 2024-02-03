package models

import (
	"time"

	"github.com/personal-financial-app/internal/repository"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}
 
func (e Event) Save() error{
	// later on we will add to the database	
	query := `INSERT INTO events(name , description , location , dateTime , user_id)
	VALUES(?,?,?,?,?)`
	stmt , err := repository.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	result , err := stmt.Exec(e.Name , e.Description , e.Location , e.DateTime , e.UserID)
	if err != nil{
		return err
	}
	id , err := result.LastInsertId() // get the last inserted id to fetch the last inserted data
	e.ID = id
	return err
}

// what is the difference between the Query and the Exec

func GetAllEvents() ([]Event , error){
	query := "SELECT * FROM events"
	rows , err := repository.DB.Query(query)
	if err != nil{
		return nil , err
	}

	defer rows.Close()

	var events []Event

	for rows.Next(){
		var event Event
		err := rows.Scan(&event.ID , &event.Name , &event.Description , &event.Location , &event.DateTime , &event.UserID)
		if err != nil{
			return nil , err
		}
		
		events = append(events, event)
	}
	return events , nil
}