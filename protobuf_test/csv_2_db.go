package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func csvToMongo(csvFilePath string, databaseName string, collectionName string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
			return fmt.Errorf("MongoDB client creation error: %v", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
			return fmt.Errorf("MongoDB connection error: %v", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(databaseName).Collection(collectionName)
	
	file, err := os.Open(csvFilePath)
	if err != nil {
			return fmt.Errorf("CSV file open error: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true  // Make the reader more lenient, even if there are irregular quotes

	_, err = reader.Read()  // skip the first line
	if err != nil {
			return fmt.Errorf("CSV read error: %v", err)
	}
	
	for {
			multi_string := false
			record, err := reader.Read()
			if err != nil {
					if err == io.EOF {
							break
					} else if strings.Contains(err.Error(), "wrong number of fields") {
							multi_string = true
					} else {
							return fmt.Errorf("CSV read error: %v", err)
					}
			}
			pageId, _ := strconv.Atoi(record[0])
			pageTitle := record[1]
			titleClass, _ := strconv.Atoi(record[2])
			revisionTextLength, _ := strconv.Atoi(record[3])
			revisionTextLines, _ := strconv.Atoi(record[4])
			revisionDatetime, _ := strconv.ParseInt(record[5], 10, 64)

			if multi_string {
					i := len(record)-1
					pageId, _ = strconv.Atoi(record[0])
					pageTitle = strings.Join(record[1:i-3], ",")
					titleClass, _ = strconv.Atoi(record[i-3])
					revisionTextLength, _ = strconv.Atoi(record[i-2])
					revisionTextLines, _ = strconv.Atoi(record[i-1])
					revisionDatetime, _ = strconv.ParseInt(record[i], 10, 64)
			}

			page := _Page{
							PageId:              pageId,
							PageTitle:           pageTitle,
							TitleClass:          titleClass,
							RevisionTextLength:  revisionTextLength,
							RevisionTextLines:   revisionTextLines,
							RevisionDatetime:    revisionDatetime,
			}

			_, err = collection.InsertOne(context.Background(), page)
			if err != nil {
					return fmt.Errorf("InsertOne error: %v", err)
			}
	}
	return nil
}