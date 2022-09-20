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
func (r *mutationResolver) CreateCanvas(ctx context.Context, input model.NewCanvas) (*model.Canvas, error) {
	var canvas canvas.Canvas
	canvas.Name = input.Name
	canvasID := canvas.Create()

	return &model.Canvas{ID: canvasID, Name: canvas.Name}, nil
}

// AddShape is the resolver for the addShape field.
func (r *mutationResolver) AddShape(ctx context.Context, input model.NewRectangle) (*model.Canvas, error) {
	var canvas canvas.Canvas
	var responseShapes []*model.Rectangle
	canvas.Get(input.CanvasID)

	canvas.Contents = append(canvas.Contents, &shapes.Rectangle{
		X:       input.X,
		Y:       input.Y,
		Width:   input.Width,
		Height:  input.Height,
		Fill:    input.Fill,
		Outline: input.Outline,
	})

	canvas.Save()

	for _, rectangle := range canvas.Contents {
		responseShapes = append(responseShapes, &model.Rectangle{
			X:       rectangle.X,
			Y:       rectangle.Y,
			Width:   rectangle.Width,
			Height:  rectangle.Height,
			Fill:    rectangle.Fill,
			Outline: rectangle.Outline,
		})
	}

	return &model.Canvas{ID: canvas.ID, Name: canvas.Name, Contents: responseShapes}, nil
}

// Canvases is the resolver for the canvases field.
func (r *queryResolver) Canvases(ctx context.Context) ([]*model.Canvas, error) {
	var canvases []*model.Canvas
	var canvasRectangles []*model.Rectangle
	var canvasesData = canvas.GetAll()
	for _, canvas := range canvasesData {
		for _, rectangle := range canvas.Contents {
			canvasRectangles = append(canvasRectangles, &model.Rectangle{
				X:       rectangle.X,
				Y:       rectangle.Y,
				Width:   rectangle.Width,
				Height:  rectangle.Height,
				Fill:    rectangle.Fill,
				Outline: rectangle.Outline,
			})
		}
		canvases = append(canvases, &model.Canvas{
			ID:       canvas.ID,
			Name:     canvas.Name,
			Contents: canvasRectangles,
		})
	}
	return canvases, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
