package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/juanfgs/canvas/graph/generated"
	"github.com/juanfgs/canvas/graph/model"
	"github.com/juanfgs/canvas/internal/canvas"
	"github.com/juanfgs/canvas/internal/canvas/shapes"
)

// CreateCanvas is the resolver for the createCanvas field.
func (r *mutationResolver) addShape(ctx context.Context, input model.NewRectangle) (*model.Canvas, error) {
	var rectangle shapes.Rectangle
	return &model.Canvas{ID: canvasID, Name: canvas.Name}, nil
}


// CreateCanvas is the resolver for the createCanvas field.
func (r *mutationResolver) CreateCanvas(ctx context.Context, input model.NewCanvas) (*model.Canvas, error) {
	var canvas canvas.Canvas
	canvas.Name = input.Name
	canvasID := canvas.Save()
	return &model.Canvas{ID: canvasID, Name: canvas.Name, Contents: canvas.Contents}, nil
}

// Canvases is the resolver for the canvases field.
func (r *queryResolver) Canvases(ctx context.Context) ([]*model.Canvas, error) {
	var canvases []*model.Canvas
	placeholderCanvas := model.Canvas{
		Name: "Canvas ID",
	}
	canvases = append(canvases, &placeholderCanvas)
	return canvases, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
