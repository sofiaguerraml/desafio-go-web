package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	list   []domain.Ticket
}

func NewRouter(router *gin.Engine, list []domain.Ticket) *Router {
	return &Router{router: router, list: list}
}

func (r *Router) MapRoutes() {
	repo := tickets.NewRepository(r.list)
	service := tickets.NewService(repo)
	ticketHandler := handler.NewService(service)
	tickets := r.router.Group("/tickets")

	{
		tickets.GET("/getByCountry/:dest", ticketHandler.GetTicketsByCountry())
		tickets.GET("/getAverage/:dest", ticketHandler.AverageDestination())
	}
}
