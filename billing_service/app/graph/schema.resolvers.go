package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"billing_service/app/graph/generated"
	"billing_service/app/graph/model"
	"context"
)

// Hits is the resolver for the hits field.
func (r *queryResolver) Hits(ctx context.Context) ([]*model.APIHit, error) {
	hits, err := r.HitsHandler.GetHits()
	if err != nil {
		return nil, err
	}

	var responseHits []*model.APIHit
	for _, hit := range hits {
		responseHits = append(responseHits, &model.APIHit{
			UUID:       hit.UUID.String(),
			CustomerID: int(hit.CustomerID),
			ServiceID:  int(hit.ServiceID),
		})
	}
	return responseHits, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
