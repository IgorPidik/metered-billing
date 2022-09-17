package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"billing_service/app/graph/generated"
	"billing_service/app/graph/model"
	db_models "billing_service/app/models"
	"context"
)

func mapHitToResponse(hit *db_models.APIHit) *model.APIHit {
	return &model.APIHit{
		UUID:       hit.UUID.String(),
		CustomerID: int(hit.CustomerID),
		ServiceID:  int(hit.ServiceID),
	}
}

func mapHitsToResponse(hits []*db_models.APIHit) []*model.APIHit {
	var responseHits []*model.APIHit
	for _, hit := range hits {
		responseHits = append(responseHits, mapHitToResponse(hit))
	}
	return responseHits
}

// Hits is the resolver for the hits field.
func (r *queryResolver) Hits(ctx context.Context) ([]*model.APIHit, error) {
	hits, err := r.HitsHandler.GetHits()
	if err != nil {
		return nil, err
	}

	return mapHitsToResponse(hits), nil
}

// Invoices is the resolver for the invoices field.
func (r *queryResolver) Invoices(ctx context.Context) ([]*model.Invoice, error) {
	invoices, err := r.InvoicesHandler.GetInvoices()
	if err != nil {
		return nil, err
	}

	var responseInvoices []*model.Invoice
	for _, invoice := range invoices {
		responseInvoices = append(responseInvoices, &model.Invoice{
			UUID:       invoice.UUID.String(),
			CustomerID: int(invoice.CustomerID),
			Hits:       mapHitsToResponse(invoice.Hits),
		})
	}
	return responseInvoices, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
