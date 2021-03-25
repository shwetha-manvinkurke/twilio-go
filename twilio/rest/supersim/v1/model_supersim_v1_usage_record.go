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

// SupersimV1UsageRecord struct for SupersimV1UsageRecord
type SupersimV1UsageRecord struct {
	// The SID of the Account that incurred the usage.
	AccountSid *string `json:"AccountSid,omitempty"`
	// Total data downloaded in bytes, aggregated by the query parameters.
	DataDownload *int32 `json:"DataDownload,omitempty"`
	// Total of data_upload and data_download.
	DataTotal *int32 `json:"DataTotal,omitempty"`
	// Total data uploaded in bytes, aggregated by the query parameters.
	DataUpload *int32 `json:"DataUpload,omitempty"`
	// SID of the Fleet resource on which the usage occurred.
	FleetSid *string `json:"FleetSid,omitempty"`
	// Alpha-2 ISO Country Code of the country the usage occurred in.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// SID of the Network resource on which the usage occurred.
	NetworkSid *string `json:"NetworkSid,omitempty"`
	// The time period for which the usage is reported.
	Period *map[string]interface{} `json:"Period,omitempty"`
	// SID of a Sim resource to which the UsageRecord belongs.
	SimSid *string `json:"SimSid,omitempty"`
}
