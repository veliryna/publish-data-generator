package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"data-generator/client/handlers"
	"data-generator/reports"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcPort = ":50051"

func main() {
	interactiveMode := flag.Bool("i", false, "use grps command line interface")
	conn, err := grpc.NewClient(grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := reports.NewReportServiceClient(conn)
	ctx := context.Background()

	go handlers.CreatePublishingServer()

	timeStream, err := client.StreamTimeBasedReports(ctx, &reports.EmptyRequest{})
	if err != nil {
		log.Fatalf("Failed to open stream: %v", err)
	}

	eventStream, err := client.StreamEventBasedReports(ctx, &reports.EmptyRequest{})
	if err != nil {
		log.Fatalf("Failed to open stream: %v", err)
	}

	go handlers.SendTimeBasedReports(timeStream)

	go handlers.SendEventBasedReports(eventStream)

	if *interactiveMode {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nDevice control terminal:")
		for {
			fmt.Print("\n> ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			args := strings.Split(input, " ")

			switch args[0] {
			case "exit":
				fmt.Println("Exiting client...")
				return

			case "mute":
				if len(args) != 2 {
					fmt.Println("Usage: mute <device_id>")
					continue
				}
				id, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println("Invalid device ID")
					continue
				}
				handlers.HandleMuteCommand(ctx, client, int32(id))

			case "unmute":
				if len(args) != 2 {
					fmt.Println("Usage: unmute <device_id>")
					continue
				}
				id, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println("Invalid device ID")
					continue
				}
				handlers.HandleUnmuteCommand(ctx, client, int32(id))

			case "list":
				handlers.HandleListCommand(ctx, client)

			default:
				fmt.Println("Use mute id, unmute id, list, exit commands.")
			}
		}
	}
	select {}
}
