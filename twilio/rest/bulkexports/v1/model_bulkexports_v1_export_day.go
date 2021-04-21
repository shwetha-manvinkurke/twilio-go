/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// BulkexportsV1ExportDay struct for BulkexportsV1ExportDay
type BulkexportsV1ExportDay struct {
	// The date when resource is created
	CreateDate *string `json:"create_date,omitempty"`
	// The date of the data in the file
	Day *string `json:"day,omitempty"`
	// The friendly name specified when creating the job
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
	// Size of the file in bytes
	Size *int32 `json:"size,omitempty"`
}
