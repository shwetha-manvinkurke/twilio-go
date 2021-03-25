/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListByocTrunkResponse struct for ListByocTrunkResponse
type ListByocTrunkResponse struct {
	ByocTrunks []VoiceV1ByocTrunk        `json:"ByocTrunks,omitempty"`
	Meta       ListByocTrunkResponseMeta `json:"Meta,omitempty"`
}
