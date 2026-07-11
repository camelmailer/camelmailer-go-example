package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	camelmailer "github.com/camelmailer/camelmailer-go"
)

func main() {
	apiKey := os.Getenv("CAMELMAILER_API_KEY")
	if apiKey == "" {
		log.Fatal("set CAMELMAILER_API_KEY to a server API key")
	}

	var opts []camelmailer.Option
	// Self-hosted instance? Defaults to https://app.camelmailer.com.
	if baseURL := os.Getenv("CAMELMAILER_BASE_URL"); baseURL != "" {
		opts = append(opts, camelmailer.WithBaseURL(baseURL))
	}
	client := camelmailer.NewClient(apiKey, opts...)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 1. Send an email.
	sent, err := client.Emails.Send(ctx, &camelmailer.SendEmailRequest{
		From:     camelmailer.Address{Email: envOr("CAMELMAILER_FROM", "you@yourdomain.com")},
		To:       []camelmailer.Address{{Email: envOr("CAMELMAILER_TO", "delivered@example.com")}},
		Subject:  "Hello from CamelMailer",
		HTMLBody: "<strong>It works!</strong>",
	})
	if err != nil {
		var apiErr *camelmailer.APIError
		if errors.As(err, &apiErr) {
			log.Fatalf("send failed — %s: %s", apiErr.Code, apiErr.Message)
		}
		log.Fatal(err)
	}
	fmt.Println("Sent message", sent.MessageID)

	// 2. Read the server's stats.
	stats, err := client.Stats.Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server stats: %d sent, %d pending, %d bounced (%d total)\n",
		stats.Sent, stats.Pending, stats.Bounced, stats.Total)
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
