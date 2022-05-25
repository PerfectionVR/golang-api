package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/PerfectionVR/golang-api/graph/generated"
	"github.com/PerfectionVR/golang-api/graph/model"
)

func (r *queryResolver) Applications(ctx context.Context) ([]*model.Application, error) {
	if len(r.applications) == 0 {
		applications := getApplications()
		r.applications = applications
	}
	return r.applications, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
