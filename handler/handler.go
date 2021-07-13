package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"GolangEcho/models"

	"github.com/labstack/echo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
// for localhost mongoDB
// const connectionString = "mongodb://localhost:27017"
const connectionString = "mongodb+srv://<username>:<password>@cluster0.fkwok.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

// Database Name
const dbName = "test"

// Collection name
const collName = "todolist"

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

// GetAllTask get all the task route
func GetAllTask(c echo.Context) error {
	c.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	return c.JSON(http.StatusOK, payload)
}

func CreateTask(c echo.Context) error {
	c.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "POST")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDoList
	_ = json.NewDecoder(c.Request().Body).Decode(&task)
	// fmt.Println(task, r.Body)
	insertOneTask(task)

	return c.JSON(http.StatusOK, task)
}

// TaskComplete update task route
func TaskComplete(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "PUT")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := c.Param("id")
	taskComplete(params)
	return c.JSON(http.StatusOK, params)
}

// UndoTask undo the complete task route
func UndoTask(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "PUT")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := c.Param("id")
	undoTask(params)
	return c.JSON(http.StatusOK, params)
}

// DeleteTask delete one task route
func DeleteTask(c echo.Context) error {
	c.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "DELETE")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := c.Param("id")
	deleteOneTask(params)
	return c.JSON(http.StatusOK, params)
	// json.NewEncoder(w).Encode("Task not found")

}

// DeleteAllTask delete all tasks route
func DeleteAllTask(c echo.Context) error {
	c.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	count := deleteAllTask()
	return c.JSON(http.StatusOK, count)
	// json.NewEncoder(w).Encode("Task not found")

}

// get all task from the DB and return it
func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// task complete method, update task's status to true
func taskComplete(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// task undo method, update task's status to false
func undoTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}
