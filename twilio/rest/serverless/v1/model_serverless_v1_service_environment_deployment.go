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

// ServerlessV1ServiceEnvironmentDeployment struct for ServerlessV1ServiceEnvironmentDeployment
type ServerlessV1ServiceEnvironmentDeployment struct {
	// The SID of the Account that created the Deployment resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The SID of the Build for the deployment
	BuildSid *string `json:"BuildSid,omitempty"`
	// The ISO 8601 date and time in GMT when the Deployment resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the Deployment resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The SID of the Environment for the Deployment
	EnvironmentSid *string `json:"EnvironmentSid,omitempty"`
	// The SID of the Service that the Deployment resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The unique string that identifies the Deployment resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the Deployment resource
	Url *string `json:"Url,omitempty"`
}
