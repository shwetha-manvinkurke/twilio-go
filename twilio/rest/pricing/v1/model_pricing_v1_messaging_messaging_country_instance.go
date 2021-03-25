/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV1MessagingMessagingCountryInstance struct for PricingV1MessagingMessagingCountryInstance
type PricingV1MessagingMessagingCountryInstance struct {
	// The name of the country
	Country *string `json:"Country,omitempty"`
	// The list of InboundPrice records
	InboundSmsPrices *[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices `json:"InboundSmsPrices,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The list of OutboundSMSPrice records
	OutboundSmsPrices *[]PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices `json:"OutboundSmsPrices,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"PriceUnit,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
}
