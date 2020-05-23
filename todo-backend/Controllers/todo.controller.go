package Controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"maera/todos/Helpers"
	"maera/todos/Models"
	"net/http"
	"time"
)

func GetTodos(w http.ResponseWriter, r *http.Request){

	w.Header().Add("content-type","application/json")
	w.Header().Add("Access-Control-Allow-Origin","*")
	w.Header().Add("Access-Control-Allow-Methods","GET,POST,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers","*")

	collection := Helpers.DB.Collection("ToDoItems")
	cursor,_ := collection.Find(context.TODO(),bson.M{})

	var results []Models.ToDoItem

	for cursor.Next(context.TODO()){
		var item Models.ToDoItem
		err := cursor.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results,item)

	}
	json.NewEncoder(w).Encode(results)
}

func CreateTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application/json")
	w.Header().Add("Access-Control-Allow-Origin","*")
	w.Header().Add("Access-Control-Allow-Methods","GET,POST,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers","*")
	var item Models.ToDoItem
	json.NewDecoder(r.Body).Decode(&item)
	log.Println(item)
	collection := Helpers.DB.Collection("ToDoItems")
	ctx, _ := context.WithTimeout(context.Background(),10*time.Second)
	result, err :=collection.InsertOne(ctx,item)
	log.Println(err)
	json.NewEncoder(w).Encode(result)
}

func DeleteTodo(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	response.Header().Add("Access-Control-Allow-Origin","*")
	response.Header().Add("Access-Control-Allow-Methods","*")
	response.Header().Add("Access-Control-Allow-Headers","*")
	response.Header().Add("content-type","application/json")
	collection := Helpers.DB.Collection("ToDoItems")
	ctx, _ := context.WithTimeout(context.Background(),10*time.Second)
	vars := mux.Vars(request)
	id,_ := primitive.ObjectIDFromHex(vars["id"])
	collection.DeleteOne(ctx, bson.M{"_id":id})
}
