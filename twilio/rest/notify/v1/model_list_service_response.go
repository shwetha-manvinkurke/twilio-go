/*
 * Twilio - Notify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceResponse struct for ListServiceResponse
type ListServiceResponse struct {
	Meta     ListCredentialResponseMeta `json:"meta,omitempty"`
	Services []NotifyV1Service          `json:"services,omitempty"`
}
