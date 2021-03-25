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

// InsightsV1CallEvent struct for InsightsV1CallEvent
type InsightsV1CallEvent struct {
	AccountSid  *string                 `json:"AccountSid,omitempty"`
	CallSid     *string                 `json:"CallSid,omitempty"`
	CarrierEdge *map[string]interface{} `json:"CarrierEdge,omitempty"`
	ClientEdge  *map[string]interface{} `json:"ClientEdge,omitempty"`
	Edge        *string                 `json:"Edge,omitempty"`
	Group       *string                 `json:"Group,omitempty"`
	Level       *string                 `json:"Level,omitempty"`
	Name        *string                 `json:"Name,omitempty"`
	SdkEdge     *map[string]interface{} `json:"SdkEdge,omitempty"`
	SipEdge     *map[string]interface{} `json:"SipEdge,omitempty"`
	Timestamp   *string                 `json:"Timestamp,omitempty"`
}
