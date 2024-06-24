package handler

import (
	"context"
	"fmt"

	// "kerala/witchcraft/routes"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	err := routes.Init()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	http.HandleFunc("/", HelloServer)
	_, err := MongoConnection()
	if err != nil {
		log.Fatal(err)
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "connection Success")
		})
	}
	log.Fatal(http.ListenAndServe("kerala-34uxbqzd8-alshifas-projects", nil))
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Andrews")
}

func MongoConnection() (*mongo.Client, error) {
	opts := options.Client()
	opts.ApplyURI("mongodb+srv://andreamose988:E6xYOfUn0erPPQYa@cluster0.4wntr2x.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	} else {
		fmt.Println("connection success")
	}

	err = client.Ping(context.TODO(), readpref.PrimaryPreferred())
	if err != nil {
		return nil, err
	}

	return client, nil
}
