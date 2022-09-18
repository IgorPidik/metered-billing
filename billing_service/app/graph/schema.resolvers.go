package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"billing_service/app/graph/generated"
	"billing_service/app/graph/model"
	utils "billing_service/app/graph/utils"
	"context"

	"github.com/google/uuid"
)

// Invoices is the resolver for the invoices field.
func (r *queryResolver) Invoices(ctx context.Context) ([]*model.Invoice, error) {
	invoices, err := r.InvoicesHandler.GetInvoices()
	if err != nil {
		return nil, err
	}

	var responseInvoices []*model.Invoice
	for _, invoice := range invoices {
		responseInvoices = append(responseInvoices, utils.MapInvoiceToResponse(invoice))
	}
	return responseInvoices, nil
}

// Invoice is the resolver for the invoice field.
func (r *queryResolver) Invoice(ctx context.Context, invoiceUUID string) (*model.Invoice, error) {
	convertedUUID, err := uuid.Parse(invoiceUUID)
	if err != nil {
		return nil, err
	}

	invoice, dbErr := r.InvoicesHandler.GetInvoice(convertedUUID)
	if dbErr != nil {
		return nil, dbErr
	}

	if invoice != nil {
		return utils.MapInvoiceToResponse(invoice), nil
	}

	return nil, nil
}

// InvoicesForCustomer is the resolver for the invoicesForCustomer field.
func (r *queryResolver) InvoicesForCustomer(ctx context.Context, customerID int) ([]*model.Invoice, error) {
	invoices, err := r.InvoicesHandler.GetInvoicesForCustomerId(uint(customerID))
	if err != nil {
		return nil, err
	}

	return utils.MapInvoicesToResponse(invoices), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
