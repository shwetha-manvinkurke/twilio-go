/*
 * Twilio - Ip_messaging
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

// IpMessagingV2Service struct for IpMessagingV2Service
type IpMessagingV2Service struct {
	AccountSid                   *string                 `json:"AccountSid,omitempty"`
	ConsumptionReportInterval    *int32                  `json:"ConsumptionReportInterval,omitempty"`
	DateCreated                  *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated                  *time.Time              `json:"DateUpdated,omitempty"`
	DefaultChannelCreatorRoleSid *string                 `json:"DefaultChannelCreatorRoleSid,omitempty"`
	DefaultChannelRoleSid        *string                 `json:"DefaultChannelRoleSid,omitempty"`
	DefaultServiceRoleSid        *string                 `json:"DefaultServiceRoleSid,omitempty"`
	FriendlyName                 *string                 `json:"FriendlyName,omitempty"`
	Limits                       *map[string]interface{} `json:"Limits,omitempty"`
	Links                        *map[string]interface{} `json:"Links,omitempty"`
	Media                        *map[string]interface{} `json:"Media,omitempty"`
	Notifications                *map[string]interface{} `json:"Notifications,omitempty"`
	PostWebhookRetryCount        *int32                  `json:"PostWebhookRetryCount,omitempty"`
	PostWebhookUrl               *string                 `json:"PostWebhookUrl,omitempty"`
	PreWebhookRetryCount         *int32                  `json:"PreWebhookRetryCount,omitempty"`
	PreWebhookUrl                *string                 `json:"PreWebhookUrl,omitempty"`
	ReachabilityEnabled          *bool                   `json:"ReachabilityEnabled,omitempty"`
	ReadStatusEnabled            *bool                   `json:"ReadStatusEnabled,omitempty"`
	Sid                          *string                 `json:"Sid,omitempty"`
	TypingIndicatorTimeout       *int32                  `json:"TypingIndicatorTimeout,omitempty"`
	Url                          *string                 `json:"Url,omitempty"`
	WebhookFilters               *[]string               `json:"WebhookFilters,omitempty"`
	WebhookMethod                *string                 `json:"WebhookMethod,omitempty"`
}
