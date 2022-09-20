package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juanfgs/canvas/graph/generated"
	"github.com/juanfgs/canvas/graph/model"
)

// CreateCanvas is the resolver for the createCanvas field.
func (r *mutationResolver) CreateCanvas(ctx context.Context, input model.NewCanvas) (*model.Canvas, error) {
	panic(fmt.Errorf("not implemented: CreateCanvas - createCanvas"))
}

// Canvases is the resolver for the canvases field.
func (r *queryResolver) Canvases(ctx context.Context) ([]*model.Canvas, error) {
	var links []*model.Canvas
	placeholderCanvas := model.Canvas{
		Name: "Canvas ID",
	}
	links = append(links, &placeholderCanvas)
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
