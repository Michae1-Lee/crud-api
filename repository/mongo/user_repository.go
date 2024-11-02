package mongo

import (
	"context"
	"crud-api/domain"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	client *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{client: client.Database("user").Collection("user")}
}

func (repo *UserRepository) CreateUser(user domain.User) error {
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": user}

	_, err := repo.client.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("can't put user: %w", err)
	}
	return err
}

func (repo *UserRepository) GetUser(id int) (domain.User, error) {
	var user domain.User
	filter := bson.M{"_id": id}
	err := repo.client.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repo *UserRepository) DeleteUser(id int) error {
	filter := bson.M{"_id": id}
	_, err := repo.client.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) FindByLogin(login string) (domain.User, error) {
	var user domain.User
	filter := bson.M{"login": login}
	err := repo.client.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}

func (repo *UserRepository) Find(login string, password string) (domain.User, error) {
	var user domain.User
	filter := bson.M{"login": login, "password": password}
	err := repo.client.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}
