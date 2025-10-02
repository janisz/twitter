package twitter

import (
	"fmt"
)

// MediaCategory represents the category of media being uploaded
type MediaCategory string

const (
	// MediaCategoryTweetImage for tweet image uploads
	MediaCategoryTweetImage MediaCategory = "tweet_image"
	// MediaCategoryDMImage for direct message image uploads
	MediaCategoryDMImage MediaCategory = "dm_image"
	// MediaCategorySubtitles for subtitle uploads
	MediaCategorySubtitles MediaCategory = "subtitles"
)

// MediaUploadRequest contains the parameters for uploading media
type MediaUploadRequest struct {
	// Media is the file content to upload
	Media []byte
	// MediaType is the MIME type of the media (e.g., "image/png", "image/jpeg", "video/mp4")
	MediaType string
	// MediaCategory specifies the category of the media
	MediaCategory MediaCategory
	// AdditionalOwners is an optional array of user IDs to grant access to the media
	AdditionalOwners []string
}

func (r MediaUploadRequest) validate() error {
	if len(r.Media) == 0 {
		return fmt.Errorf("media upload: media content is required: %w", ErrParameter)
	}
	if len(r.MediaType) == 0 {
		return fmt.Errorf("media upload: media type is required: %w", ErrParameter)
	}
	if len(r.MediaCategory) == 0 {
		return fmt.Errorf("media upload: media category is required: %w", ErrParameter)
	}
	return nil
}

// MediaUploadData contains the response data from media upload
type MediaUploadData struct {
	MediaID        int64                   `json:"media_id"`
	MediaIDString  string                  `json:"media_id_string"`
	MediaKey       string                  `json:"media_key"`
	Size           int64                   `json:"size"`
	ExpiresAfter   int64                   `json:"expires_after_secs"`
	ProcessingInfo *MediaProcessingInfo    `json:"processing_info,omitempty"`
}

// MediaProcessingInfo contains information about media processing status
type MediaProcessingInfo struct {
	State           string `json:"state"`
	CheckAfterSecs  int    `json:"check_after_secs,omitempty"`
	ProgressPercent int    `json:"progress_percent,omitempty"`
	Error           *MediaProcessingError `json:"error,omitempty"`
}

// MediaProcessingError contains error information from media processing
type MediaProcessingError struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// MediaUploadInitRequest is the request for the INIT command
type MediaUploadInitRequest struct {
	Command       string `json:"command"`
	MediaType     string `json:"media_type"`
	TotalBytes    int64  `json:"total_bytes"`
	MediaCategory string `json:"media_category,omitempty"`
}

// MediaUploadInitResponse is the response from the INIT command
type MediaUploadInitResponse struct {
	MediaID       int64  `json:"media_id"`
	MediaIDString string `json:"media_id_string"`
	ExpiresAfter  int64  `json:"expires_after_secs"`
}

// MediaUploadAppendRequest is the request for the APPEND command
type MediaUploadAppendRequest struct {
	Command      string `json:"command"`
	MediaID      string `json:"media_id"`
	SegmentIndex int    `json:"segment_index"`
	Media        []byte `json:"-"` // Will be sent as multipart data
}

// MediaUploadFinalizeRequest is the request for the FINALIZE command
type MediaUploadFinalizeRequest struct {
	Command string `json:"command"`
	MediaID string `json:"media_id"`
}

// MediaUploadResponse is the response from the media upload endpoint
type MediaUploadResponse struct {
	*MediaUploadData
	RateLimit *RateLimit
}
