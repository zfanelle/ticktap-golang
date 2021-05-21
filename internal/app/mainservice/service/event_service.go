package service

import (
	config "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config"
	event "github.com/zfanelle/ticktap-golang/internal/app/mainservice/model"
	repository "github.com/zfanelle/ticktap-golang/internal/app/mainservice/repository"
)

func Createevent(appConfig config.AppConfig, newevent event.Event) {

	repository.CreateEvent(appConfig, newevent)

}

func Getevent(appConfig config.AppConfig, eventID int) event.Event {

	event := repository.GetEvent(appConfig, eventID)

	return event

}

func GetAllEvents(appConfig config.AppConfig) []event.Event {

	events := repository.GetAllEvents(appConfig)

	return events
}
