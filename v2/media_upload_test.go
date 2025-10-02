package twitter

import (
	"context"
	"strings"
	"testing"
)

func TestUploadMedia(t *testing.T) {
	type args struct {
		ctx context.Context
		req MediaUploadRequest
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "missing media content",
			args: args{
				ctx: context.Background(),
				req: MediaUploadRequest{
					MediaType:     "image/png",
					MediaCategory: MediaCategoryTweetImage,
				},
			},
			wantErr:        true,
			wantErrMessage: "media content is required",
		},
		{
			name: "missing media type",
			args: args{
				ctx: context.Background(),
				req: MediaUploadRequest{
					Media:         strings.NewReader("fake image data"),
					MediaCategory: MediaCategoryTweetImage,
				},
			},
			wantErr:        true,
			wantErrMessage: "media type is required",
		},
		{
			name: "missing media category",
			args: args{
				ctx: context.Background(),
				req: MediaUploadRequest{
					Media:     strings.NewReader("fake image data"),
					MediaType: "image/png",
				},
			},
			wantErr:        true,
			wantErrMessage: "media category is required",
		},
		{
			name: "valid request structure",
			args: args{
				ctx: context.Background(),
				req: MediaUploadRequest{
					Media:            []byte("fake image data"),
					MediaType:        "image/png",
					MediaCategory:    MediaCategoryTweetImage,
					AdditionalOwners: []string{"user123"},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For validation-only tests, we can test the validation directly
			if err := tt.args.req.validate(); err != nil {
				if !tt.wantErr {
					t.Errorf("MediaUploadRequest.validate() error = %v, wantErr %v", err, tt.wantErr)
				} else if !strings.Contains(err.Error(), tt.wantErrMessage) {
					t.Errorf("MediaUploadRequest.validate() error = %v, wanted error containing %v", err, tt.wantErrMessage)
				}
				return
			}

			if tt.wantErr {
				t.Errorf("MediaUploadRequest.validate() expected error containing %v, but got none", tt.wantErrMessage)
			}

			// Note: We're not testing the actual HTTP call here since that would require a mock server
			// This test focuses on input validation and basic structure verification
		})
	}
}

func TestMediaUploadRequest_validate(t *testing.T) {
	tests := []struct {
		name    string
		req     MediaUploadRequest
		wantErr bool
	}{
		{
			name: "valid request",
			req: MediaUploadRequest{
				Media:         []byte("test"),
				MediaType:     "image/png",
				MediaCategory: MediaCategoryTweetImage,
			},
			wantErr: false,
		},
		{
			name: "missing media",
			req: MediaUploadRequest{
				MediaType:     "image/png",
				MediaCategory: MediaCategoryTweetImage,
			},
			wantErr: true,
		},
		{
			name: "missing media type",
			req: MediaUploadRequest{
				Media:         []byte("test"),
				MediaCategory: MediaCategoryTweetImage,
			},
			wantErr: true,
		},
		{
			name: "missing media category",
			req: MediaUploadRequest{
				Media:     []byte("test"),
				MediaType: "image/png",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.validate(); (err != nil) != tt.wantErr {
				t.Errorf("MediaUploadRequest.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}