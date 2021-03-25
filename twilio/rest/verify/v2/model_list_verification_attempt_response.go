/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListVerificationAttemptResponse struct for ListVerificationAttemptResponse
type ListVerificationAttemptResponse struct {
	Attempts []VerifyV2VerificationAttempt       `json:"Attempts,omitempty"`
	Meta     ListVerificationAttemptResponseMeta `json:"Meta,omitempty"`
}
