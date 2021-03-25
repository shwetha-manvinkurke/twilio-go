/*
 * Twilio - Accounts
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

// AccountsV1CredentialCredentialPublicKey struct for AccountsV1CredentialCredentialPublicKey
type AccountsV1CredentialCredentialPublicKey struct {
	// The SID of the Account that created the Credential that the PublicKey resource belongs to
	AccountSid *string `json:"AccountSid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The URI for this resource, relative to `https://accounts.twilio.com`
	Url *string `json:"Url,omitempty"`
}
