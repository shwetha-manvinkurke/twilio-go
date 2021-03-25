/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// InsightsV1VideoRoomSummaryVideoParticipantSummary struct for InsightsV1VideoRoomSummaryVideoParticipantSummary
type InsightsV1VideoRoomSummaryVideoParticipantSummary struct {
	// Account SID associated with the room.
	AccountSid *string `json:"AccountSid,omitempty"`
	// Codecs detected from the participant.
	Codecs *[]string `json:"Codecs,omitempty"`
	// Amount of time in seconds the participant was in the room.
	DurationSec *int32 `json:"DurationSec,omitempty"`
	// Name of the edge location the participant connected to.
	EdgeLocation *string `json:"EdgeLocation,omitempty"`
	// Reason the participant left the room.
	EndReason *string `json:"EndReason,omitempty"`
	// Errors encountered by the participant.
	ErrorCode *int32 `json:"ErrorCode,omitempty"`
	// Twilio error code dictionary link.
	ErrorCodeUrl *string `json:"ErrorCodeUrl,omitempty"`
	// When the participant joined the room.
	JoinTime *time.Time `json:"JoinTime,omitempty"`
	// When the participant left the room
	LeaveTime *time.Time `json:"LeaveTime,omitempty"`
	// Twilio media region the participant connected to.
	MediaRegion *string `json:"MediaRegion,omitempty"`
	// The application-defined string that uniquely identifies the participant within a Room.
	ParticipantIdentity *string `json:"ParticipantIdentity,omitempty"`
	// Unique identifier for the participant.
	ParticipantSid *string `json:"ParticipantSid,omitempty"`
	// Object containing information about the participant's data from the room.
	Properties *map[string]interface{} `json:"Properties,omitempty"`
	// Object containing information about the SDK name and version.
	PublisherInfo *map[string]interface{} `json:"PublisherInfo,omitempty"`
	// Unique identifier for the room.
	RoomSid *string `json:"RoomSid,omitempty"`
	// Status of the room.
	Status *string `json:"Status,omitempty"`
	// URL of the participant resource.
	Url *string `json:"Url,omitempty"`
}
