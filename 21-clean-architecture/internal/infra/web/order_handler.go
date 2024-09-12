package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/felipecardosodeoliveira/Golang/21-clean-architecture/internal/entity"
	"github.com/felipecardosodeoliveira/Golang/21-clean-architecture/internal/usecase"
	"github.com/felipecardosodeoliveira/Golang/21-clean-architecture/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	if h.OrderCreatedEvent == nil {
		http.Error(w, "OrderCreatedEvent is not initialized", http.StatusInternalServerError)
		return
	}
	if h.EventDispatcher == nil {
		http.Error(w, "EventDispatcher is not initialized", http.StatusInternalServerError)
		return
	}
	if h.OrderRepository == nil {
		http.Error(w, "OrderRepository is not initialized", http.StatusInternalServerError)
		return
	}
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)

	fmt.Println(*createOrder)

	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
