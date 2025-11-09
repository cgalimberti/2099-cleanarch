package web

import (
	"encoding/json"
	"net/http"

	"github.com/cgalimberti/2099-CleanArch/internal/usecase"
)

type OrderHandler struct {
	ListOrdersUseCase usecase.ListOrdersUseCase
}

func NewOrderHandler(listOrdersUseCase usecase.ListOrdersUseCase) *OrderHandler {
	return &OrderHandler{
		ListOrdersUseCase: listOrdersUseCase,
	}
}

func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	orders, err := h.ListOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
