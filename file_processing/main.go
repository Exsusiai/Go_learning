package main

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Page struct {
    PageId              int
    PageTitle           string
    TitleClass          int
    RevisionTextLength  int
    RevisionTextLines   int
    RevisionDatetime    int64
}

func main() {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    if err = client.Connect(context.Background()); err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.Background())

    collection := client.Database("test_cols_csv").Collection("c1")
    
    file, err := os.Open("/home/jason/LD_works/Go/test_csv/c1.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.LazyQuotes = true  // Make the reader more lenient, even if there are irregular quotes

    _, err = reader.Read()  // skip the first line
    if err != nil {
        log.Fatal(err)
    }
    
    for {
			multi_string := false
			record, err := reader.Read()
			if err != nil {
					if err == io.EOF {
							break
					} else if strings.Contains(err.Error(), "wrong number of fields") {
							// log.Printf("Skipping line: %s\n", err)
							multi_string = true
							// continue
					} else {
							log.Fatal(err)
					}
			}
			pageId, _ := strconv.Atoi(record[0])
			pageTitle := record[1]
			titleClass, _ := strconv.Atoi(record[2])
			revisionTextLength, _ := strconv.Atoi(record[3])
			revisionTextLines, _ := strconv.Atoi(record[4])
			revisionDatetime, _ := strconv.ParseInt(record[5], 10, 64)

			if multi_string {//this is to process the multi strings combine with comma
				i := len(record)-1
				pageId, _ = strconv.Atoi(record[0])
				pageTitle = strings.Join(record[1:i-3], ",")
				titleClass, _ = strconv.Atoi(record[i-3])
				revisionTextLength, _ = strconv.Atoi(record[i-2])
				revisionTextLines, _ = strconv.Atoi(record[i-1])
				revisionDatetime, _ = strconv.ParseInt(record[i], 10, 64)
			}

			page := Page{
					PageId:              pageId,
					PageTitle:           pageTitle,
					TitleClass:          titleClass,
					RevisionTextLength:  revisionTextLength,
					RevisionTextLines:   revisionTextLines,
					RevisionDatetime:    revisionDatetime,
			}

			_, err = collection.InsertOne(context.Background(), page)
			if err != nil {
					log.Fatal(err)
			}
    }
}
