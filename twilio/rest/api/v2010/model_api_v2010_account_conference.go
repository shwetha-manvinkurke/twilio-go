/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountConference struct for ApiV2010AccountConference
type ApiV2010AccountConference struct {
	// The SID of the Account that created this resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The API version used to create this conference
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// The call SID that caused the conference to end
	CallSidEndingConference *string `json:"CallSidEndingConference,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was created
	DateCreated *string `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated *string `json:"DateUpdated,omitempty"`
	// A string that you assigned to describe this conference room
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The reason why a conference ended.
	ReasonConferenceEnded *string `json:"ReasonConferenceEnded,omitempty"`
	// A string that represents the Twilio Region where the conference was mixed
	Region *string `json:"Region,omitempty"`
	// The unique string that identifies this resource
	Sid *string `json:"Sid,omitempty"`
	// The status of this conference
	Status *string `json:"Status,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"SubresourceUris,omitempty"`
	// The URI of this resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
}
