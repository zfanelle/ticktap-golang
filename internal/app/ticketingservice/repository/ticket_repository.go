package repsitory

import (
	"fmt"

	config "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/config"
	ticket "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/model"
)

func CreateTicket(appConfig config.AppConfig, newTicket ticket.Ticket) {

	db := appConfig.DB

	sql := fmt.Sprintf("INSERT INTO ticktap.ticket (event,account) VALUES (%d, %d)", newTicket.Event, newTicket.Account)

	db.MustExec(sql)

}

func GetTicket(appConfig config.AppConfig, TicketID int) ticket.Ticket {
	db := appConfig.DB

	tickets := []ticket.TicketDB{}

	sql := fmt.Sprintf("SELECT * FROM ticktap.ticket WHERE ID = %d", TicketID)

	db.Select(&tickets, sql)

	// Add in error checking

	ticket := tickets[0].DBtoTicket()

	return ticket
}

func GetAllTickets(appConfig config.AppConfig) []ticket.Ticket {
	db := appConfig.DB

	dbTickets := []ticket.TicketDB{}

	sql := "SELECT * FROM ticktap.ticket"

	db.Select(&dbTickets, sql)

	tickets := []ticket.Ticket{}

	// Add in error checking

	for _, ticket := range dbTickets {
		tickets = append(tickets, ticket.DBtoTicket())
	}

	return tickets
}
