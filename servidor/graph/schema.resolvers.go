package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kissmarkrivas/go_sum/graph/generated"
	"github.com/kissmarkrivas/go_sum/graph/model"
)

// CreateSuma is the resolver for the createSuma field.
func (r *mutationResolver) CreateSuma(ctx context.Context, input model.SumInput) (*model.Suma, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateSuma is the resolver for the updateSuma field.
func (r *mutationResolver) UpdateSuma(ctx context.Context, sumaID int, input model.SumInput) (*model.Suma, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteSuma is the resolver for the deleteSuma field.
func (r *mutationResolver) DeleteSuma(ctx context.Context, sumaID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Usuarios is the resolver for the usuarios field.
func (r *queryResolver) Usuarios(ctx context.Context) ([]*model.Suma, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
