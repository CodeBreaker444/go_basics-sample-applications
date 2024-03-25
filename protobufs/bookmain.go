package main

import (
    "fmt"
    "io/ioutil"
    "log"

    "google.golang.org/protobuf/proto"
    pb "protobufs/protos" // Replace with your package path
)

func main() {
    // Create a new book instance
    bk := &pb.Book{
        Person: []*pb.Data{
            {Name: "Alice", Email: "alice@example.com", Id: 1},
            {Name: "Bob", Email: "bob@example.com", Id: 2},
        },
    }

    // Serialize the data to a byte slice
    data, err := proto.Marshal(bk)
    if err != nil {
        log.Fatalf("Failed to marshal data: %v", err)
    }

    // Write serialized data to a file
    if err := ioutil.WriteFile("book.bin", data, 0644); err != nil {
        log.Fatalf("Failed to write data to file: %v", err)
    }

    // Read the data back from the file
    newData, err := ioutil.ReadFile("book.bin")
    if err != nil {
        log.Fatalf("Failed to read data from file: %v", err)
    }

    // Deserialize the data
    newBk := &pb.Book{}
    if err := proto.Unmarshal(newData, newBk); err != nil {
        log.Fatalf("Failed to unmarshal data: %v", err)
    }

    // Print the deserialized data
    fmt.Println("Deserialized data:", newBk)
}
