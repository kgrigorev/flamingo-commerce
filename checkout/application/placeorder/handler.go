package placeorder

import (
	"context"

	"flamingo.me/flamingo-commerce/v3/checkout/domain/placeorder/process"
)

// Handler for handling PlaceOrder related commands
type Handler struct {
	coordinator *Coordinator
}

// Inject dependencies
func (h *Handler) Inject(
	c *Coordinator,
) *Handler {
	h.coordinator = c

	return h
}

// StartPlaceOrder handles start place order command
func (h *Handler) StartPlaceOrder(ctx context.Context, command StartPlaceOrderCommand) (*process.Context, error) {
	return h.coordinator.New(ctx, command.Cart, command.ReturnURL)
}

// RefreshPlaceOrder handles RefreshPlaceOrder command
func (h *Handler) RefreshPlaceOrder(ctx context.Context, command RefreshPlaceOrderCommand) (*process.Context, error) {
	p, err := h.coordinator.LastProcess(ctx)
	if err != nil {
		return nil, err
	}
	lastPlaceOrderCtx := p.Context()

	// proceed in state
	h.coordinator.Run(ctx)

	return &lastPlaceOrderCtx, nil
}

// RefreshPlaceOrderBlocking handles RefreshPlaceOrder blocking
func (h *Handler) RefreshPlaceOrderBlocking(ctx context.Context, command RefreshPlaceOrderCommand) (*process.Context, error) {
	return h.coordinator.RunBlocking(ctx)
}

// HasUnfinishedProcess checks for processes not in final state
func (h *Handler) HasUnfinishedProcess(ctx context.Context) (bool, error) {
	return h.coordinator.HasUnfinishedProcess(ctx)
}