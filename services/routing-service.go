package services

import "log"

var RouteServiceInstance RouteService = routeService{}

type RouteService interface {
	HandlePing() ([]byte, error)
}
type routeService struct{}

func (rs routeService) HandlePing() ([]byte, error) {
	log.Println("FROM SERVICE")
	return []byte("PONG"), nil
}
