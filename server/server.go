package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"data-generator/reports"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

type ReportServiceServer struct {
	reports.UnimplementedReportServiceServer
	Mu     sync.RWMutex
	Status map[int32]bool
}

func NewReportServiceServer() *ReportServiceServer {
	return &ReportServiceServer{
		Status: make(map[int32]bool),
	}
}

func (s *ReportServiceServer) MuteDevice(ctx context.Context, req *reports.DeviceControlRequest) (*reports.DeviceControlResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	if req.DeviceId == 0 {
		for key := range s.Status {
			s.Status[key] = true
		}
		log.Printf("Muted all registered devices")
		return &reports.DeviceControlResponse{Message: fmt.Sprintf("All devices muted")}, nil
	}
	s.Status[req.DeviceId] = true
	log.Printf("Muted device %d", req.DeviceId)
	return &reports.DeviceControlResponse{Message: fmt.Sprintf("Device %d muted", req.DeviceId)}, nil
}

func (s *ReportServiceServer) UnmuteDevice(ctx context.Context, req *reports.DeviceControlRequest) (*reports.DeviceControlResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	if req.DeviceId == 0 {
		for key := range s.Status {
			s.Status[key] = false
		}
		log.Printf("Unmuted all registered devices")
		return &reports.DeviceControlResponse{Message: fmt.Sprintf("All devices unmuted")}, nil
	}
	s.Status[req.DeviceId] = false
	log.Printf("Unmuted device %d", req.DeviceId)
	return &reports.DeviceControlResponse{Message: fmt.Sprintf("Device %d unmuted", req.DeviceId)}, nil
}

func (s *ReportServiceServer) GetDeviceStatuses(ctx context.Context, _ *reports.EmptyRequest) (*reports.DeviceStatusResponse, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	var statuses []*reports.DeviceStatus
	for id, muted := range s.Status {
		statuses = append(statuses, &reports.DeviceStatus{
			DeviceId: id,
			Muted:    muted,
		})
	}

	return &reports.DeviceStatusResponse{
		Statuses: statuses,
	}, nil
}

func (s *ReportServiceServer) IsMuted(id int32) bool {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	return s.Status[id]
}

func (s *ReportServiceServer) RegisterDevice(id int32) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	if _, exists := s.Status[id]; !exists {
		s.Status[id] = false // devices are unmuted by default
	}
}

func (s *ReportServiceServer) StreamTimeBasedReports(req *reports.EmptyRequest, stream reports.ReportService_StreamTimeBasedReportsServer) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":          "grpc-streaming-consumer",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return fmt.Errorf("failed to create consumer: %w", err)
	}
	defer consumer.Close()

	var topic []string
	topic = append(topic, "time-based-messages")
	if err := consumer.SubscribeTopics(topic, nil); err != nil {
		return fmt.Errorf("failed to subscribe: %w", err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Kafka read error: %v", err)
			continue
		}

		var report reports.TimeBasedReport
		if err := proto.Unmarshal(msg.Value, &report); err != nil {
			log.Printf("Unmarshal error: %v", err)
			continue
		}

		if err := stream.Send(&report); err != nil {
			return fmt.Errorf("stream send error: %w", err)
		}
	}
}

func (s *ReportServiceServer) StreamEventBasedReports(req *reports.EmptyRequest, stream reports.ReportService_StreamEventBasedReportsServer) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":          "grpc-streaming-consumer",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return fmt.Errorf("failed to create consumer: %w", err)
	}
	defer consumer.Close()

	var topic []string
	topic = append(topic, "event-based-messages")
	if err := consumer.SubscribeTopics(topic, nil); err != nil {
		return fmt.Errorf("failed to subscribe: %w", err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Kafka read error: %v", err)
			continue
		}

		var report reports.EventBasedReport
		if err := proto.Unmarshal(msg.Value, &report); err != nil {
			log.Printf("Unmarshal error: %v", err)
			continue
		}

		if err := stream.Send(&report); err != nil {
			return fmt.Errorf("stream send error: %w", err)
		}
	}
}
