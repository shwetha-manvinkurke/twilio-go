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

// PricingV1PhoneNumberPhoneNumberCountryInstance struct for PricingV1PhoneNumberPhoneNumberCountryInstance
type PricingV1PhoneNumberPhoneNumberCountryInstance struct {
	// The name of the country
	Country *string `json:"Country,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The list of PhoneNumberPrices records
	PhoneNumberPrices *[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices `json:"PhoneNumberPrices,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"PriceUnit,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
}
