package main

import (
	"context"
	//"flag"
	"encoding/base64"
	"fmt"
	"log"
	"net"

	"gitlab.com/secops/development/aws/terrascan/cmd"
	"gitlab.com/secops/development/aws/terrascan/helpers"
	pb "gitlab.com/secops/development/aws/terrascan/proto"
	//"gitlab.com/secops/development/aws/terrascan/resource"

	"google.golang.org/grpc"
	//"github.com/spf13/viper"
)

type Message struct {
	Body          string
	MD5           string
	MessageId     string
	ReceiptHandle string
}

type server struct {
	pb.UnimplementedEventEmitterServer
}

func (s *server) SendEvent(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())

	event := in.GetEvent()
	log.Printf("Received: %v", event)

	go func(e string) {
		defer cancel()
		if err := cmd.Init(e); err != nil {
			helpers.Response = base64.StdEncoding.EncodeToString([]byte(err.Error()))

			log.Printf("Init() Error: %v\n", err)
		}
		log.Println("Done Init()")
	}(event)
	select {
	case <-ctx.Done():
		log.Println("Done()")
		return &pb.MessageResponse{
			Response: helpers.Response,
		}, nil
	}
}

var (
	port = 8080
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterEventEmitterServer(s, &server{})
	log.Printf("Listening on: %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to server: %v\n", err)
	}
}

/*
func Init() {

	log.Println("Terrascan Started")

	// target resource block
	targetResource := resource.Resource{
		Region: region,
	}


		* Cli Configuration
		// retreive input variables
		// resource-name: string of resources in terraformer format, ex...s3,ec2
		// filters: string representations of filters, ex..vpc=vpc_id1:vpc_id2:vpc_id3
		flag.StringVar(&targetResource.Name, "resource-name", "None", "Resource Name Flag Set")
		flag.StringVar(&targetResource.Filters, "filters", "None", "Resource Name Flag Set")
		flag.Parse()


	if err := cmd.Setup(&targetResource); err != nil {
		log.Println(err)
	}

}
*/
