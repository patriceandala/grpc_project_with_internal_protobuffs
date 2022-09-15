package main

import (
	"fmt"
	"github.com/patriceandala/grpc_project_with_internal_protobuffs/protocol_buffers"
)

func main() {
	user1 := protocol_buffers.User{
		Id:        1,
		FirstName: "Patrice",
		LastName:  "Andala",
		Email:     "patriceandala@gmail.com",
	}

	user2 := protocol_buffers.User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@gmail.com",
	}

	message := protocol_buffers.Message{
		MessageId:  1,
		SenderId:   string(user1.Id),
		ReceiverId: string(user2.Id),
		Message:    "Hello there",
	}

	fmt.Println(message)

}
