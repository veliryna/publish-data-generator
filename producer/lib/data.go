package lib

import (
	"data-generator/reports"
	"data-generator/server"
	"log"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

var kafkaTimeTopic string = "time-based-messages"

func SimulateDeviceTimeData(p *kafka.Producer, id int32, server *server.ReportServiceServer) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var lastHeartbeat = time.Now()
	var lastBreath = time.Now()
	var runnerDistance int32 = 0

	heartrate := time.Duration(rand.Intn(600)+400) * time.Millisecond
	breathrate := time.Duration(rand.Intn(400)+1000) * time.Millisecond

	for t := range ticker.C {
		distanceIncrement := int32(rand.Intn(3) + 2)
		runnerDistance += distanceIncrement

		report := &reports.TimeBasedReport{
			DeviceId:  id,
			Timestamp: t.UnixMilli(),
			Distance:  runnerDistance,
		}

		for lastHeartbeat.Add(heartrate).Before(t) || lastHeartbeat.Add(heartrate).Equal(t) {
			lastHeartbeat = lastHeartbeat.Add(heartrate)
			report.Heartbeats = append(report.Heartbeats, lastHeartbeat.UnixMilli())
		}
		for lastBreath.Add(breathrate).Before(t) || lastBreath.Add(breathrate).Equal(t) {
			lastBreath = lastBreath.Add(breathrate)
			report.Breaths = append(report.Breaths, lastBreath.UnixMilli())
		}

		if server.IsMuted(id) {
			continue
		}

		data, err := proto.Marshal(report)
		if err != nil {
			log.Printf("Device %d proto marshal error: %v\n", id, err)
			continue
		}

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kafkaTimeTopic, Partition: kafka.PartitionAny},
			Value:          data,
		}, nil)
		if err != nil {
			log.Printf("Device %d: failed to send message: %v\n", id, err)
		} else {
			log.Printf("Device %d sent report at %d\n", id, report.Timestamp)
		}
	}
}

func SimulateDeviceEventData(p *kafka.Producer, amount int, server *server.ReportServiceServer) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		device_id := rand.Int31n(int32(amount)) + 1
		event_id := rand.Int31n(3) + 1
		startEventSimulation(p, device_id, event_id)
	}
}
