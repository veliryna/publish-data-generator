package handlers

import (
	"context"
	"data-generator/reports"
	"fmt"
)

func HandleMuteCommand(ctx context.Context, client reports.ReportServiceClient, id int32) {
	resp, err := client.MuteDevice(ctx, &reports.DeviceControlRequest{DeviceId: id})
	if err != nil {
		fmt.Printf("Error muting device %d: %v\n", id, err)
		return
	}
	fmt.Println(resp.Message)
}

func HandleUnmuteCommand(ctx context.Context, client reports.ReportServiceClient, id int32) {
	resp, err := client.UnmuteDevice(ctx, &reports.DeviceControlRequest{DeviceId: id})
	if err != nil {
		fmt.Printf("Error unmuting device %d: %v\n", id, err)
		return
	}
	fmt.Println(resp.Message)
}

func HandleListCommand(ctx context.Context, client reports.ReportServiceClient) {
	resp, err := client.GetDeviceStatuses(ctx, &reports.EmptyRequest{})
	if err != nil {
		fmt.Printf("Error listing statuses: %v\n", err)
		return
	}
	fmt.Println("Device Status List:")
	for _, s := range resp.Statuses {
		status := "unmuted"
		if s.Muted {
			status = "muted"
		}
		fmt.Printf("Device %d: %s\n", s.DeviceId, status)
	}
}
