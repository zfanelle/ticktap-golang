package model

type EventDTO struct {
	Host                  int    `json:"host"`
	EventName             string `json:"event_name"`
	MaximumTicketCapacity int    `json:"maximum_ticket_capacity"`
}

func (eventDTO EventDTO) DTOtoevent() Event {
	return Event{Host: eventDTO.Host, EventName: eventDTO.EventName, MaximumTicketCapacity: eventDTO.MaximumTicketCapacity}
}

type Event struct {
	Id                    int    `json:"id"`
	Host                  int    `json:"host"`
	EventName             string `json:"event_name"`
	MaximumTicketCapacity int    `json:"maximum_ticket_capacity"`
}

type EventDB struct {
	Id                    int    `db:"id"`
	Host                  int    `db:"host"`
	EventName             string `db:"event_name"`
	MaximumTicketCapacity int    `db:"maximum_ticket_capacity"`
}

func (eventDB EventDB) DBtoevent() Event {
	return Event{Id: eventDB.Id, Host: eventDB.Host, EventName: eventDB.EventName, MaximumTicketCapacity: eventDB.MaximumTicketCapacity}
}
