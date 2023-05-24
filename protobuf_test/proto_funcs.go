package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/Exsusiai/protobuf_test/proto_path"
	"github.com/golang/protobuf/proto"
)

// The function to serialize data
func serializeData(pages []*proto_path.Page) ([]byte, error) {
	var data []byte
	
	for _, p := range pages {
		// Serialize the page
		serializedPage, err := proto.Marshal(p)
		if err != nil {
			return nil, err
		}

		// Create and serialize frame header
		header := &proto_path.FrameHeader{
			FrameId: p.GetPageId(),
			Timestamp: 0,
			FrameLength: int32(len(serializedPage)),
		}
		serializedHeader, err := proto.Marshal(header)
		if err != nil {
			return nil, err
		}

		// Append the serialized header and page to the data
		data = append(data, serializedHeader...)
		data = append(data, serializedPage...)
	}
	return data, nil
}

// The function to write data to a .proto file
func writeDataToFile(data []byte, filename string) error {
	return ioutil.WriteFile(filename, data, 0644)
}

// The function to read data from a .pb file
func retrieveMessage(id int, filename string) error {
	// open the file
	file, err := os.Open(filename)
	if err != nil {
			return err
	}
	defer file.Close()

	//Assuming that a message will not exceed 1024 bytes, if it exceeds, you need to adjust this value
	// buf := make([]byte, 1024)
	
	for {
			// read frame header
			headerBuf := make([]byte, 1024) 
			n, err := file.Read(headerBuf)
			if err == io.EOF {
					break
			}
			if err != nil {
					fmt.Println("Error reading file for header: ", err)
					return err
			}

			var frameHeader proto_path.FrameHeader
			err = proto.Unmarshal(headerBuf[:n], &frameHeader)
			if err != nil {
					fmt.Println("Error unmarshalling header: ", err)
					return err
			}

			//check if the message id matches
			if int(frameHeader.GetFrameId()) == id {
					// if it matches, read the message
					msgBuf := make([]byte, frameHeader.GetFrameLength()) // Allocate buffer according to FrameLength
					n, err := file.Read(msgBuf)
					if err != nil {
							fmt.Println("Error reading file for message: ", err)		
							return err
					}

					var page proto_path.Page
					err = proto.Unmarshal(msgBuf[:n], &page)
					if err != nil {
							fmt.Println("Error unmarshalling page: ", err)
							return err
					}

					// print the message
					fmt.Printf("Page ID: %v\n", page.GetPageId())
					fmt.Printf("Page Title: %v\n", page.GetPageTitle())
					fmt.Printf("Title Class: %v\n", page.GetTitleClass())
					fmt.Printf("Revision Text Length: %v\n", page.GetRevisionTextLength())
					fmt.Printf("Revision Text Lines: %v\n", page.GetRevisionTextLines())
					fmt.Printf("Revision Datetime: %v\n", page.GetRevisionDatetime())
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