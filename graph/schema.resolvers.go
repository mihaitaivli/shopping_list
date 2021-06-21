package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/mihaitaivli/shopping_list/data_utils"
	"github.com/mihaitaivli/shopping_list/graph/generated"
	"github.com/mihaitaivli/shopping_list/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) AddUser(ctx context.Context, input model.NewUser) (*string, error) {
	// defer client.Disconnect(ctx)
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
	collection := client.Database("localDb").Collection("Users")

	objId, convertErr := primitive.ObjectIDFromHex(input.ID)
	if convertErr != nil {
		panic("Error converting string to ObjectId")
	}

	filter := bson.M{"_id": bson.M{"$eq": objId}}
	update := bson.M{"$set": input}

	res, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Printf("Error while updating userId: %s - %v\n", input.ID, err)
		return nil, err
	}

	count := res.ModifiedCount
	if count == 1 {
		return &input.ID, nil
	}
	return nil, errors.New("unsuccessful insert")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*string, error) {
	collection := client.Database("localDb").Collection("Users")
	objId, convertErr := primitive.ObjectIDFromHex(userID)
	if convertErr != nil {
		panic("Error converting string to ObjectId")
	}
	res, err := collection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		fmt.Printf("Error deleting user with id: #{userID}\n")
		return nil, errors.New("unsuccessful delete")
	}

	if res.DeletedCount != 1 {
		fmt.Printf("No user with id: #{userID}\n")
		return nil, errors.New("unsuccessful delete")
	}

	return &userID, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var client = data_utils.CreateClient()
