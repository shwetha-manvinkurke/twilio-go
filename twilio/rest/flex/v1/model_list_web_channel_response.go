/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListWebChannelResponse struct for ListWebChannelResponse
type ListWebChannelResponse struct {
	FlexChatChannels []FlexV1WebChannel      `json:"FlexChatChannels,omitempty"`
	Meta             ListChannelResponseMeta `json:"Meta,omitempty"`
}
