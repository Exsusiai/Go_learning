package main

import (
	"encoding/binary"
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

		var FrameId uint32 = uint32(p.GetPageId())
		var Timestamp uint64 = 0
		var FrameLength uint32 = uint32(len(serializedPage))

		FrameHeader_data := make([]byte, 16)

		binary.BigEndian.PutUint32(FrameHeader_data[0:4], FrameId)
		binary.BigEndian.PutUint64(FrameHeader_data[4:12], Timestamp)
		binary.BigEndian.PutUint32(FrameHeader_data[12:16], FrameLength)		

		// Append the serialized page to the data
		data = append(data, FrameHeader_data...)
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

	for {
			// read frame header
			headerBuf := make([]byte, 16) 
			_, err := file.Read(headerBuf)
			if err == io.EOF {
					fmt.Println("End of file reached")
					break
			}

			if err != nil {
					fmt.Println("Error reading file for header: ", err)
					return err
			}

			frameId := uint32(binary.BigEndian.Uint32(headerBuf[0:4]))
			// timestamp := uint64(binary.BigEndian.Uint64(headerBuf[4:12]))
			frameLength := uint32(binary.BigEndian.Uint32(headerBuf[12:16]))

			// fmt.Printf("FrameId: %v\n", frameId)
			// fmt.Printf("Timestamp: %v\n", timestamp)
			// fmt.Printf("FrameLength: %v\n", frameLength)

			//check if the message id matches
			if int(frameId) == id {
					fmt.Println("Frame ID matches.")
					// if it matches, read the message
					msgBuf := make([]byte, frameLength) // Allocate buffer according to FrameLength
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
					_, err = file.Seek(int64(frameLength), io.SeekCurrent)
					if err != nil {
							fmt.Println("Error skipping message: ", err)
							return err
					}
			}
	}

	fmt.Println("Did not find the message with the given ID.")
	return nil
}