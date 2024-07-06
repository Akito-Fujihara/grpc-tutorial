package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Akito-Fujihara/grpc-tutorial/pb"
	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Akito",
		Email:       "akito@gmail.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "Hello, I'm Akito",
		},
		Birthday: &pb.Date{
			Year:  1995,
			Month: 1,
			Day:   1,
		},
	}

	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalf("failed to marshal employee: %v", err)
	}

	if err := ioutil.WriteFile("employee.bin", binData, 0666); err != nil {
		log.Fatalf("failed to write employee.bin: %v", err)
	}

	inData, err := ioutil.ReadFile("employee.bin")
	if err != nil {
		log.Fatalf("failed to read employee.bin: %v", err)
	}

	employee = &pb.Employee{}
	if err := proto.Unmarshal(inData, employee); err != nil {
		log.Fatalf("failed to unmarshal employee: %v", err)
	}

	m := jsonpb.Marshaler{}
	jsonData, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalf("failed to marshal employee: %v", err)
	}
	// fmt.Println(jsonData)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(jsonData, readEmployee); err != nil {
		log.Fatalf("failed to unmarshal employee: %v", err)
	}
	fmt.Println(readEmployee)
}
