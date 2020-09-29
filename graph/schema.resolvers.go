package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/Pranjalya/vue-go-graphql/graph/generated"
	"github.com/Pranjalya/vue-go-graphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	jsonE, _ := json.Marshal(todo)
	log.Println(string(jsonE))
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input string) (*model.Todo, error) {
	var todo *model.Todo
	log.Println("INPUT", input)
	for i, t := range r.todos {
		if t.ID == input {
			r.todos[i].Done = !r.todos[i].Done
			todo = r.todos[i]
			break
		}
	}
	return todo, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
