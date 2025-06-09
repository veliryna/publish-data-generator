package main

import (
	"flag"
	"log"
	"net"
	"os"

	"data-generator/producer/lib"
	"data-generator/reports"
	"data-generator/server"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = ":50051"

var kafkaBroker = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")

func main() {
	var runnersAmount int
	flag.IntVar(&runnersAmount, "runners_amount", 500, "Number of simulated runners")
	flag.Parse()

	srv := server.NewReportServiceServer()
	go startGrpcServer(srv)

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            kafkaBroker,
		"queue.buffering.max.messages": 100000,
		"batch.num.messages":           10000,
		"linger.ms":                    5,
		"compression.type":             "snappy",
		"acks":                         "1",
	})
	if err != nil {
		log.Printf("Failed to create Kafka producer")
		return
	}

	defer p.Flush(5000)
	defer p.Close()

	go lib.SimulateDeviceEventData(p, runnersAmount, srv)
	for i := 1; i <= runnersAmount; i++ {
		device_id := int32(i)
		srv.RegisterDevice(device_id)
		go lib.SimulateDeviceTimeData(p, device_id, srv)
	}

	// this line allows the goroutines to run indefinitely until stopped by a signal or an error
	select {}
}

func startGrpcServer(s reports.ReportServiceServer) {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", grpcPort, err)
	}

	grpcServer := grpc.NewServer()
	reports.RegisterReportServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	log.Printf("gRPC server listening on %s\n", grpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
