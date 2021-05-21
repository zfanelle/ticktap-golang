package repsitory

import (
	"fmt"

	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	event "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
)

func CreateEvent(appConfig config.AppConfig, newEvent event.Event) {

	db := appConfig.DB

	sql := fmt.Sprintf("INSERT INTO ticktap.event (host,event_name,maximum_ticket_capacity) VALUES (\"%d\", \"%s\", \"%d\")", newEvent.Host, newEvent.EventName, newEvent.MaximumTicketCapacity)

	db.MustExec(sql)

}

func GetEvent(appConfig config.AppConfig, eventID int) event.Event {
	db := appConfig.DB

	events := []event.EventDB{}

	sql := fmt.Sprintf("SELECT * FROM ticktap.event WHERE ID = %d", eventID)

	db.Select(&events, sql)

	// Add in error checking

	event := events[0].DBtoevent()

	return event
}

func GetAllEvents(appConfig config.AppConfig) []event.Event {
	db := appConfig.DB

	dbevents := []event.EventDB{}

	sql := fmt.Sprintf("SELECT * FROM ticktap.event")

	db.Select(&dbevents, sql)

	events := []event.Event{}

	// Add in error checking

	for _, event := range dbevents {
		events = append(events, event.DBtoevent())
	}

	return events
}
