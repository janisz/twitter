package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func main() {
	// This example demonstrates how to upload media to X/Twitter

	// You need to provide your Bearer Token for authentication
	token := os.Getenv("TWITTER_BEARER_TOKEN")
	if token == "" {
		log.Fatal("TWITTER_BEARER_TOKEN environment variable is required")
	}

	// Create the client
	client := &twitter.Client{
		Authorizer: authorize{Token: token},
		Client:     nil, // Use default HTTP client
		Host:       "https://api.twitter.com",
	}

	// Example 1: Upload an image
	// In a real scenario, you would read from a file
	imageData := []byte("fake image data for example")

	uploadReq := twitter.MediaUploadRequest{
		Media:         imageData,
		MediaType:     "image/png",
		MediaCategory: twitter.MediaCategoryTweetImage,
	}

	ctx := context.Background()
	resp, err := client.UploadMedia(ctx, uploadReq)
	if err != nil {
		log.Fatalf("Failed to upload media: %v", err)
	}

	fmt.Printf("Media uploaded successfully!\n")
	fmt.Printf("Media ID: %d\n", resp.MediaID)
	fmt.Printf("Media ID String: %s\n", resp.MediaIDString)
	fmt.Printf("Media Key: %s\n", resp.MediaKey)
	fmt.Printf("Size: %d bytes\n", resp.Size)
	fmt.Printf("Expires after: %d seconds\n", resp.ExpiresAfter)

	if resp.ProcessingInfo != nil {
		fmt.Printf("Processing state: %s\n", resp.ProcessingInfo.State)
		if resp.ProcessingInfo.CheckAfterSecs > 0 {
			fmt.Printf("Check after: %d seconds\n", resp.ProcessingInfo.CheckAfterSecs)
		}
	}

	// Example 2: Upload with additional owners
	fmt.Println("\n--- Upload with additional owners ---")

	uploadReqWithOwners := twitter.MediaUploadRequest{
		Media:            []byte("another fake image"),
		MediaType:        "image/jpeg",
		MediaCategory:    twitter.MediaCategoryTweetImage,
		AdditionalOwners: []string{"user123", "user456"},
	}

	resp2, err := client.UploadMedia(ctx, uploadReqWithOwners)
	if err != nil {
		log.Fatalf("Failed to upload media with additional owners: %v", err)
	}

	fmt.Printf("Media with additional owners uploaded!\n")
	fmt.Printf("Media ID: %d\n", resp2.MediaID)
	fmt.Printf("Media ID String: %s\n", resp2.MediaIDString)
	fmt.Printf("Media Key: %s\n", resp2.MediaKey)

	// Example 3: Upload subtitles
	fmt.Println("\n--- Upload subtitles ---")

	subtitleData := []byte("1\n00:00:01,000 --> 00:00:04,000\nHello World!")

	subtitleReq := twitter.MediaUploadRequest{
		Media:         subtitleData,
		MediaType:     "text/srt",
		MediaCategory: twitter.MediaCategorySubtitles,
	}

	resp3, err := client.UploadMedia(ctx, subtitleReq)
	if err != nil {
		log.Fatalf("Failed to upload subtitles: %v", err)
	}

	fmt.Printf("Subtitles uploaded successfully!\n")
	fmt.Printf("Media ID: %d\n", resp3.MediaID)
	fmt.Printf("Media ID String: %s\n", resp3.MediaIDString)
	fmt.Printf("Media Key: %s\n", resp3.MediaKey)

	fmt.Println("\nTip: Use the Media ID or Media Key when creating tweets with attached media")
}