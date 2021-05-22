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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetScenarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var scenarios []models.Scenario

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created_at", -1}})

	var collection = helper.ConnectDB().Database("bulman").Collection("scenarios")
	cur, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var scenario models.Scenario
		err := cur.Decode(&scenario)
		if err != nil {
			log.Fatal(err)
		}
		scenarios = append(scenarios, scenario)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(scenarios)
}

func GetScenariosByProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var scenarios []models.Scenario
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{}
	filter["project.$id"] = id

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created_at", -1}})

	jsonString, _ := json.Marshal(filter)
	log.Printf("mgo query: %s\n", jsonString)

	var collection = helper.ConnectDB().Database("bulman").Collection("scenarios")
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var scenario models.Scenario
		err := cur.Decode(&scenario)
		if err != nil {
			log.Fatal(err)
		}
		scenarios = append(scenarios, scenario)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(scenarios)
}

func GetScenario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var scenario models.Scenario
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}

	var collection = helper.ConnectDB().Database("bulman").Collection("scenarios")
	err := collection.FindOne(context.TODO(), filter).Decode(&scenario)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(scenario)
}

func CreateScenario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var scenario models.Scenario

	_ = json.NewDecoder(r.Body).Decode(&scenario)

	if scenario.CreatedAt.IsZero() {
		scenario.CreatedAt = time.Now()
	}

	scenario.Project.Ref = "project"

	var collection = helper.ConnectDB().Database("bulman").Collection("scenarios")
	result, err := collection.InsertOne(context.TODO(), scenario)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateScenario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var scenario models.Scenario
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&scenario)

	update := bson.D{
		{"$set", bson.D{
			{"name", scenario.Name},
			{"state", scenario.State},
			{"project", scenario.Project},
			{"url", scenario.URL},
			{"method", scenario.Method},
			{"payload", scenario.Payload},
			{"headers", scenario.Headers},
			{"parameters", scenario.Parameters},
		}},
	}

	var collection = helper.ConnectDB().Database("bulman").Collection("scenarios")
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&scenario)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	scenario.ID = id

	json.NewEncoder(w).Encode(scenario)
}

func DeleteScenario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	var collection = helper.ConnectDB().Database("bulman").Collection("scenarios")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}
