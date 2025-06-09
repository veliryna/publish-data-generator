package lib

import (
	"data-generator/reports"
	"log"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

var kafkaEventTopic string = "event-based-messages"

func startEventSimulation(p *kafka.Producer, device_id int32, event_id int32) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	switch event_id {
	case 1:
		bp := &BloodPressure{
			systolic:  float64(rand.Intn(40) + 90),
			diastolic: float64(rand.Intn(20) + 60),
			phase:     "normal",
		}

		for t := range ticker.C {
			values, active := bp.UpdateBloodPressure()
			if active {
				sendEventReport(p, device_id, event_id, t.UnixMilli(), values)
			}
			if bp.phase == "normal" && active {
				break
			}
		}

	case 2:
		bo := &BloodOxygen{
			value: float64(rand.Int31n(6) + 95),
			phase: "normal",
		}
		for t := range ticker.C {
			values, active := bo.UpdateBloodOxygen()
			if active {
				sendEventReport(p, device_id, event_id, t.UnixMilli(), values)
			}
			if bo.phase == "normal" && values != nil {
				break
			}
		}

	case 3:
		temp := &Temperature{
			value: float64(rand.Int31n(25) + 360),
			phase: "normal",
		}

		for t := range ticker.C {
			values, active := temp.UpdateTemperature()
			if active {
				sendEventReport(p, device_id, event_id, t.UnixMilli(), values)
			}
			if temp.phase == "normal" && values != nil {
				break
			}
		}
	}
}

func sendEventReport(p *kafka.Producer, device_id, event_id int32, timestamp int64, data []int32) {
	report := &reports.EventBasedReport{
		DeviceId:  device_id,
		Timestamp: timestamp,
		EventId:   event_id,
		EventData: data,
	}

	bytes, err := proto.Marshal(report)
	if err != nil {
		log.Printf("Device %d proto marshal event error: %v\n", report.DeviceId, err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaEventTopic, Partition: kafka.PartitionAny},
		Value:          bytes,
	}, nil)
	if err != nil {
		log.Printf("Device %d failed to send event report: %v\n", report.DeviceId, err)
	} else {
		log.Printf("Device %d sent event report at %d\n", report.DeviceId, report.Timestamp)
	}
}
