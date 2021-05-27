package service

import (
	config "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/config"
	ticket "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/model"
	repository "github.com/zfanelle/ticktap-golang/internal/app/ticketingservice/repository"
)

func CreateTicket(appConfig config.AppConfig, newTicket ticket.Ticket) {

	repository.CreateTicket(appConfig, newTicket)

}

func GetTicket(appConfig config.AppConfig, ticketID int) ticket.Ticket {

	ticket := repository.GetTicket(appConfig, ticketID)

	return ticket

}

func GetAllTickets(appConfig config.AppConfig) []ticket.Ticket {

	tickets := repository.GetAllTickets(appConfig)

	return tickets
}
