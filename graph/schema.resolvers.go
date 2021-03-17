package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mihaitaivli/shopping_list/data_utils"
	"github.com/mihaitaivli/shopping_list/graph/generated"
	"github.com/mihaitaivli/shopping_list/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = data_utils.CreateClient()

func (r *mutationResolver) AddUser(ctx context.Context, input model.NewUser) (*string, error) {
	collection := client.Database("localDb").Collection("Users")

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		fmt.Printf("Error while inserting user: %v", err)
		return nil, err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return &id, nil
}

func (r *mutationResolver) EditUser(ctx context.Context, input model.EditUser) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LinkUser(ctx context.Context, input model.LinkUserInput) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UnlinkUser(ctx context.Context, linkedUserID string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddList(ctx context.Context, name string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditList(ctx context.Context, input model.EditList) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteList(ctx context.Context, input string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditSharedList(ctx context.Context, input model.EditSharedList) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddItem(ctx context.Context, input model.NewItem) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditItem(ctx context.Context, input model.EditItem) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteItem(ctx context.Context, itemID string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) List(ctx context.Context, listID string) (*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Lists(ctx context.Context, listIds []string) ([]*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, userIds []string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Item(ctx context.Context, itemID string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Items(ctx context.Context, itemIds []string) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
