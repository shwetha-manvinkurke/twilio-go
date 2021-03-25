/*
 * Twilio - Serverless
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

// ServerlessV1ServiceAsset struct for ServerlessV1ServiceAsset
type ServerlessV1ServiceAsset struct {
	// The SID of the Account that created the Asset resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The ISO 8601 date and time in GMT when the Asset resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the Asset resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The string that you assigned to describe the Asset resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URLs of the Asset resource's nested resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The SID of the Service that the Asset resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The unique string that identifies the Asset resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the Asset resource
	Url *string `json:"Url,omitempty"`
}
