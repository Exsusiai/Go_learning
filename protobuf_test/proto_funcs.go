package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Exsusiai/protobuf_test/proto_path"
	"github.com/golang/protobuf/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func readDataFromDB(ctx context.Context, client *mongo.Client, _db_name string, _collection_name string ) ([]*proto_path.Page, error) {
// 	collection := client.Database(_db_name).Collection(_collection_name)
// 	cursor, err := collection.Find(ctx, bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var pages []*proto_path.Page

// 	for cursor.Next(ctx) {
// 		var p *proto_path.Page
// 		err := cursor.Decode(&p)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		pages = append(pages, p)
// 	}

// 	return pages, nil
// }

// // The function to serialize data
// func serializeData(pages []*proto_path.Page) ([]byte, error) {
// 	var data []byte

// 	for _, p := range pages {
// 		// Serialize the page
// 		serializedPage, err := proto.Marshal(p)
// 		if err != nil {
// 			return nil, err
// 		}

// 		var FrameId uint32 = uint32(p.GetPageId())
// 		var Timestamp uint64 = 0
// 		var FrameLength uint32 = uint32(len(serializedPage))

// 		FrameHeader_data := make([]byte, 16)

// 		binary.BigEndian.PutUint32(FrameHeader_data[0:4], FrameId)
// 		binary.BigEndian.PutUint64(FrameHeader_data[4:12], Timestamp)
// 		binary.BigEndian.PutUint32(FrameHeader_data[12:16], FrameLength)

// 		// Append the serialized page to the data
// 		data = append(data, FrameHeader_data...)
// 		data = append(data, serializedPage...)
// 	}
// 	return data, nil
// }

// // The function to write data to a .proto file
// func writeDataToFile(data []byte, filename string) error {
// 	return ioutil.WriteFile(filename, data, 0644)
// }

func MongoDB_To_bin(ctx context.Context, client *mongo.Client, _db_name string, _collection_name string, filename string) error{
	file, err := os.Create(filename)
	if err != nil {
    log.Fatal(err)
	}
	defer file.Close()
	
	collection := client.Database(_db_name).Collection(_collection_name)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var p *proto_path.Page
		err := cursor.Decode(&p)
		if err != nil {
			return err
		}

		serializedPage, err := proto.Marshal(p)
		if err != nil {
			return err
		}

		header := &proto_path.FrameHeader{
			FrameId:     uint32(p.PageId),
			Timestamp:   0,
			FrameLength: uint32(len(serializedPage)),
		}
		serializedHeader, err := proto.Marshal(header)
		if err != nil {
			return err
		}

		// println(len(serializedHeader))

		_, err = file.Write(serializedHeader)
		if err != nil {
			return err
		}
		_, err = file.Write(serializedPage)
		if err != nil {
			return err
		}
	}

	return nil
}

// The function to read data from a .pb file
func retrieveMessage(id int, filename string) error {
	// open the file
	file, err := os.Open(filename)
	if err != nil {
			return err
	}
	defer file.Close()

	for {
			// read frame header
			headerBuf := make([]byte, 10)
			_, err := file.Read(headerBuf)			
			if err == io.EOF {
					fmt.Println("End of file reached")
					break
			}

			if err != nil {
					fmt.Println("Error reading file for header: ", err)
					return err
			}

			var frameHeader proto_path.FrameHeader
			err = proto.Unmarshal(headerBuf, &frameHeader)
			if err != nil {
					fmt.Println("Error unmarshalling header: ", err)
					return err
			}


			//check if the message id matches
			if int(frameHeader.GetFrameId()) == id {
					fmt.Println("Frame ID matches.")
					// if it matches, read the message
					msgBuf := make([]byte, frameHeader.GetFrameLength()) // Allocate buffer according to FrameLength
					_, err := file.Read(msgBuf)
					if err != nil {
							fmt.Println("Error reading file for message: ", err)		
							return err
					}

					var page proto_path.Page
					err = proto.Unmarshal(msgBuf, &page)
					if err != nil {
							fmt.Println("Error unmarshalling page: ", err)
							return err
					}

					fmt.Println("Find the message with the given ID:")
					fmt.Println("------message------")
					// print the message
					fmt.Printf("Page ID: %v\n", page.GetPageId())
					fmt.Printf("Page Title: %v\n", page.GetPageTitle())
					fmt.Printf("Title Class: %v\n", page.GetTitleClass())
					fmt.Printf("Revision Text Length: %v\n", page.GetRevisionTextLength())
					fmt.Printf("Revision Text Lines: %v\n", page.GetRevisionTextLines())
					fmt.Printf("Revision Datetime: %v\n", page.GetRevisionDatetime())
					fmt.Println("------message------")
					return nil
			} else {
					// if it does not match, skip the message
					_, err = file.Seek(int64(frameHeader.GetFrameLength()), io.SeekCurrent)
					if err != nil {
							fmt.Println("Error skipping message: ", err)
							return err
					}
			}
	}

	fmt.Println("Did not find the message with the given ID.")
	return nil
}