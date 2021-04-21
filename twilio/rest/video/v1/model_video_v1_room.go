/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VideoV1Room struct for VideoV1Room
type VideoV1Room struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The duration of the room in seconds
	Duration *int32 `json:"duration,omitempty"`
	// Enable Twilio's Network Traversal TURN service
	EnableTurn *bool `json:"enable_turn,omitempty"`
	// The UTC end time of the room in UTC ISO 8601 format
	EndTime *time.Time `json:"end_time,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The maximum number of published tracks allowed in the room at the same time
	MaxConcurrentPublishedTracks *int32 `json:"max_concurrent_published_tracks,omitempty"`
	// The maximum number of concurrent Participants allowed in the room
	MaxParticipants *int32 `json:"max_participants,omitempty"`
	// The region for the media server in Group Rooms
	MediaRegion *string `json:"media_region,omitempty"`
	// Whether to start recording when Participants connect
	RecordParticipantsOnConnect *bool `json:"record_participants_on_connect,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the room
	Status *string `json:"status,omitempty"`
	// The URL to send status information to your application
	StatusCallback *string `json:"status_callback,omitempty"`
	// The HTTP method we use to call status_callback
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
	// The type of room
	Type *string `json:"type,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// An array of the video codecs that are supported when publishing a track in the room
	VideoCodecs *[]string `json:"video_codecs,omitempty"`
}
