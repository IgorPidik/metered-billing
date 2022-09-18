package graph_utils

import (
	"billing_service/app/graph/model"
	db_models "billing_service/app/models"
	"time"
)

func MapInvoiceToResponse(invoice *db_models.Invoice) *model.Invoice {
	return &model.Invoice{
		UUID:       invoice.UUID.String(),
		CustomerID: int(invoice.CustomerID),
		Hits:       MapHitsToResponse(invoice.Hits),
	}
}

func MapInvoicesToResponse(invoices []*db_models.Invoice) []*model.Invoice {
	var responseInvoices []*model.Invoice
	for _, invoice := range invoices {
		responseInvoices = append(responseInvoices, MapInvoiceToResponse(invoice))
	}
	return responseInvoices
}

func MapHitToResponse(hit *db_models.APIHit) *model.APIHit {
	return &model.APIHit{
		UUID:       hit.UUID.String(),
		CustomerID: int(hit.CustomerID),
		ServiceID:  int(hit.ServiceID),
		Timestamp:  hit.Timestamp.Format(time.RFC3339),
	}
}
func MapHitsToResponse(hits []*db_models.APIHit) []*model.APIHit {
	var responseHits []*model.APIHit
	for _, hit := range hits {
		responseHits = append(responseHits, MapHitToResponse(hit))
	}
	return responseHits
}
