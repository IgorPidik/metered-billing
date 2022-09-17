package graph

import "billing_service/app/handlers"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	HitsHandler     *handlers.HitsHandler
	InvoicesHandler *handlers.InvoicesHandler
}
