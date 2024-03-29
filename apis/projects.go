package apis

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"bulman-api/helper"
	"bulman-api/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var projects []models.Project

	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "scenarios",
				"localField":   "_id",
				"foreignField": "project.$id",
				"as":           "scenariostotal",
			},
		},
		{
			"$addFields": bson.M{
				"scenariostotal": bson.M{"$size": "$scenariostotal"},
			},
		}, {
			"$sort": bson.D{{"created_at", -1}},
		},
	}

	var collection = helper.ConnectDB().Database("bulman").Collection("projects")
	cur, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var project models.Project
		err := cur.Decode(&project)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, project)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(projects)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var project models.Project
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}

	var collection = helper.ConnectDB().Database("bulman").Collection("projects")
	err := collection.FindOne(context.TODO(), filter).Decode(&project)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var project models.Project

	if project.CreatedAt.IsZero() {
		project.CreatedAt = time.Now()
	}

	_ = json.NewDecoder(r.Body).Decode(&project)

	var collection = helper.ConnectDB().Database("bulman").Collection("projects")
	result, err := collection.InsertOne(context.TODO(), project)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var project models.Project
	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&project)

	update := bson.D{
		{"$set", bson.D{
			{"name", project.Name},
			{"description", project.Description},
		}},
	}

	var collection = helper.ConnectDB().Database("bulman").Collection("projects")
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&project)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	project.ID = id

	json.NewEncoder(w).Encode(project)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	var collection = helper.ConnectDB().Database("bulman").Collection("projects")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}
