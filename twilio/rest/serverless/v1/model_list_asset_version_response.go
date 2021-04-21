/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListAssetVersionResponse struct for ListAssetVersionResponse
type ListAssetVersionResponse struct {
	AssetVersions []ServerlessV1ServiceAssetAssetVersion `json:"asset_versions,omitempty"`
	Meta          ListServiceResponseMeta                `json:"meta,omitempty"`
}
