/*
 * Twilio - Supersim
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

// SupersimV1Fleet struct for SupersimV1Fleet
type SupersimV1Fleet struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands
	CommandsEnabled *bool `json:"CommandsEnabled,omitempty"`
	// A string representing the HTTP method to use when making a request to `commands_url`
	CommandsMethod *string `json:"CommandsMethod,omitempty"`
	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number
	CommandsUrl *string `json:"CommandsUrl,omitempty"`
	// Defines whether SIMs in the Fleet are capable of using data connectivity
	DataEnabled *bool `json:"DataEnabled,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume
	DataLimit *int32 `json:"DataLimit,omitempty"`
	// The model by which a SIM is metered and billed
	DataMetering *string `json:"DataMetering,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The SID of the Network Access Profile of the Fleet
	NetworkAccessProfileSid *string `json:"NetworkAccessProfileSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands
	SmsCommandsEnabled *bool `json:"SmsCommandsEnabled,omitempty"`
	// A string representing the HTTP method to use when making a request to `sms_commands_url`
	SmsCommandsMethod *string `json:"SmsCommandsMethod,omitempty"`
	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number
	SmsCommandsUrl *string `json:"SmsCommandsUrl,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"UniqueName,omitempty"`
	// The absolute URL of the Fleet resource
	Url *string `json:"Url,omitempty"`
}
