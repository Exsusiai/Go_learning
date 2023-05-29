package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const(
	CSV_PATH = "/home/jason/LD_works/Go/test_csv/c3.csv"
	DB_NAME = "test_cols_csv"
	COLLETION_NAME = "c1"
	OUTPUT_FILE_NAME = "page_output_test_5.bin"
	SEARCH_ID = 3734554
	PROGRAM_RUN_MODE = 3 // 1: csv_To_MongoDB, 2: MongoDB_To_bin, 3: retrieve Message from .bin
)

func main() {
	//start the timer
	start := time.Now()

	switch PROGRAM_RUN_MODE {
		case 1:
		// STEP 1：Read the CSV file and write to MongoDB
		err := csvToMongo(CSV_PATH, DB_NAME, COLLETION_NAME)
		if err != nil {
				log.Fatalf("Error in csvToMongo: %v", err)
		}
		fmt.Println("Data has been successfully written to MongoDB")

		case 2:
		// STEP 2：Connect to MongoDB
		ctx := context.Background()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to MongoDB!")

		// // STEP 3：Read data from MongoDB, serialize, and write to a .pb file
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
		// err = writeDataToFile(data, OUTPUT_FILE_NAME, DB_NAME, COLLETION_NAME)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Println("Data has been successfully written to a .pb file")

		// STEP 3 Version_2：Read data from MongoDB, serialize, and write to a .pb file
		err = MongoDB_To_bin(ctx, client, DB_NAME, COLLETION_NAME, OUTPUT_FILE_NAME)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Data has been successfully written to a .pb file")

		case 3:
		// STEP 4：Read data from a .pb file and search for a specific page
		err := retrieveMessage(SEARCH_ID, OUTPUT_FILE_NAME)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Retrieved finished")
	}

	//calculate the elapsed time
	elapsed := time.Since(start)
	fmt.Println("\n---------------------------------")
	fmt.Printf("We are done : )\n")
	fmt.Printf("Program running time: %s\n", elapsed)
	fmt.Println("---------------------------------")
}


