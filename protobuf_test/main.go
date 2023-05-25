package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Exsusiai/protobuf_test/proto_path"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const(
	CSV_PATH = "/home/jason/LD_works/Go/test_csv/c3.csv"
	DB_NAME = "test_cols_csv"
	COLLETION_NAME = "c3"
	OUTPUT_FILE_NAME = "page_output_test_2.bin"
	SEARCH_ID = 10164
)

func main() {
	//start the timer
	start := time.Now()

	// Read the CSV file and write to MongoDB
	// err := csvToMongo(CSV_PATH, DB_NAME, COLLETION_NAME)
	// if err != nil {
	// 		log.Fatalf("Error in csvToMongo: %v", err)
	// }

	// // Set up MongoDB connection

	// ctx := context.Background()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)

	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connected!")

	// // Read data from MongoDB
	// pages, err := readDataFromDB(ctx, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Size of 'pages': %d\n", len(pages))

	// // Serialize data
	// data, err := serializeData(pages)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Write data to a .proto file
	// err = writeDataToFile(data, OUTPUT_FILE_NAME)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Data has been successfully written to a .pb file")

	err := retrieveMessage(SEARCH_ID, OUTPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	//calculate the elapsed time
	elapsed := time.Since(start)
	fmt.Println("\n---------------------------------")
	fmt.Printf("We are done : )\n")
	fmt.Printf("Program running time: %s\n", elapsed)
	fmt.Println("---------------------------------")
}

func readDataFromDB(ctx context.Context, client *mongo.Client) ([]*proto_path.Page, error) {
	collection := client.Database(DB_NAME).Collection(COLLETION_NAME)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var pages []*proto_path.Page

	for cursor.Next(ctx) {
		var p *proto_path.Page
		err := cursor.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}
		pages = append(pages, p)

	}
	
	return pages, nil
}
