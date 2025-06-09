#!/bin/bash

echo "Waiting for Kafka to be ready..."
kafka-topics --bootstrap-server $KAFKA_BOOTSTRAP_SERVERS --list

./producer --runners_amount=100 &

sleep 2

./client