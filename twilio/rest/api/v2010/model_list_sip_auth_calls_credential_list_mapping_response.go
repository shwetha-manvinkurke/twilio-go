/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSipAuthCallsCredentialListMappingResponse struct for ListSipAuthCallsCredentialListMappingResponse
type ListSipAuthCallsCredentialListMappingResponse struct {
	Contents        []ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping `json:"contents,omitempty"`
	End             int32                                                                             `json:"end,omitempty"`
	FirstPageUri    string                                                                            `json:"first_page_uri,omitempty"`
	NextPageUri     string                                                                            `json:"next_page_uri,omitempty"`
	Page            int32                                                                             `json:"page,omitempty"`
	PageSize        int32                                                                             `json:"page_size,omitempty"`
	PreviousPageUri string                                                                            `json:"previous_page_uri,omitempty"`
	Start           int32                                                                             `json:"start,omitempty"`
	Uri             string                                                                            `json:"uri,omitempty"`
}
