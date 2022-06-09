package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/spf13/viper"
	"log"
	"todoAPI/entity"
)

type repo struct{}

var collectionName string = viper.GetString("collection_name")

func NewFirestoreRepo() TodoRepository {
	return &repo{}
}

func createFirestoreClient() (context.Context, *firestore.Client, error) {
	ctx := context.Background()
	projectId := viper.Get("project_id").(string)
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Unable to create friestore client. \n\t Project: %s \n\tError %v", projectId, err)
		return nil, nil, err
	}

	return ctx, client, nil
}

func (r repo) Create(todo *entity.Todo) (*entity.Todo, error) {
	ctx, client, err := createFirestoreClient()
	if err != nil {
		log.Fatalf("Unable to create friestore client. Error %v", err)
		return nil, err
	}
	defer client.Close()

	newTodo := map[string]interface{}{
		"Id":          todo.Id,
		"Title":       todo.Title,
		"Description": todo.Description,
		"Done":        todo.Done,
		"Deleted":     false,
	}
	collection := client.Collection(collectionName)
	doc := collection.Doc(todo.Id)
	_, err = doc.Set(ctx, newTodo)
	if err != nil {
		log.Printf("Unable to save todo to friestore. \n\t Todo: %v \n\tError %v", newTodo, err)
		return nil, err
	}
	return todo, nil
}

func (r repo) Get(id string) (*entity.Todo, error) {
	ctx, client, err := createFirestoreClient()
	if err != nil {
		log.Fatalf("Unable to create friestore client. Error %v", err)
		return nil, err
	}
	defer client.Close()
	collection := client.Collection(collectionName)

	doc, err := collection.Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Unable to get todo from friestore. \n\t Id: %v \n\tError %v", id, err)
		return nil, err
	}

	data := doc.Data()

	if data["Deleted"].(bool) {
		err = errors.New("Item not found")
		return nil, err
	}
	todo := entity.Todo{
		Id:          data["Id"].(string),
		Title:       data["Title"].(string),
		Description: data["Description"].(string),
		Done:        data["Done"].(bool),
	}

	return &todo, nil
}

func (r repo) GetAll() ([]entity.Todo, error) {
	ctx, client, err := createFirestoreClient()
	if err != nil {
		log.Printf("Unable to create friestore client. Error %v", err)
		return nil, err
	}
	defer client.Close()

	var todos []entity.Todo

	collection := client.Collection(collectionName)

	documents, err := collection.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Failed to get Documents from firestore")
		return nil, err
	}

	for _, doc := range documents {
		data := doc.Data()
		todo := entity.Todo{
			Id:          data["Id"].(string),
			Title:       data["Title"].(string),
			Description: data["Description"].(string),
			Done:        data["Done"].(bool),
		}

		if !data["Deleted"].(bool) {
			todos = append(todos, todo)
		}
	}

	return todos, nil
}

func (r repo) Update(todo *entity.Todo) (*entity.Todo, error) {
	ctx, client, err := createFirestoreClient()
	if err != nil {
		log.Fatalf("Unable to create friestore client. Error %v", err)
		return nil, err
	}
	defer client.Close()

	todoUpdate := map[string]interface{}{
		"Title":       todo.Title,
		"Description": todo.Description,
		"Done":        todo.Done,
	}
	collection := client.Collection(collectionName)
	doc := collection.Doc(todo.Id)
	updates := []firestore.Update{}

	for k, value := range todoUpdate {
		update := firestore.Update{
			Path:  k,
			Value: value,
		}

		updates = append(updates, update)
	}

	_, err = doc.Update(ctx, updates)
	if err != nil {
		log.Printf("Unable to save todo to friestore. \n\t Todo: %v \n\tError %v", todoUpdate, err)
		return nil, err
	}
	return todo, nil
}

func (r repo) Delete(id string) error {
	ctx, client, err := createFirestoreClient()
	if err != nil {
		log.Fatalf("Unable to create friestore client. Error %v", err)
		return err
	}
	defer client.Close()
	collection := client.Collection(collectionName)

	docRef := collection.Doc(id)
	_, err = docRef.Get(ctx)
	if err != nil {
		log.Printf("Unable to get todo from friestore. \n\t Id: %v \n\tError %v", id, err)
		return err
	}

	updates := []firestore.Update{
		firestore.Update{
			Path:  "Deleted",
			Value: true,
		},
	}

	_, err = docRef.Update(ctx, updates)
	if err != nil {
		log.Printf("Unable to delete todo on friestore. \n\t Id: %v \n\tError %v", id, err)
		return err
	}

	return nil
}
