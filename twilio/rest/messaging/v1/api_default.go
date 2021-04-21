/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://messaging.twilio.com",
	}
}

// CreateAlphaSenderParams Optional parameters for the method 'CreateAlphaSender'
type CreateAlphaSenderParams struct {
	AlphaSender *string `json:"AlphaSender,omitempty"`
}

/*
* CreateAlphaSender Method for CreateAlphaSender
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
* @param optional nil or *CreateAlphaSenderParams - Optional Parameters:
* @param "AlphaSender" (string) - The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, and hyphen `-`. This value cannot contain only numbers.
* @return MessagingV1ServiceAlphaSender
 */
func (c *DefaultApiService) CreateAlphaSender(ServiceSid string, params *CreateAlphaSenderParams) (*MessagingV1ServiceAlphaSender, error) {
	path := "/v1/Services/{ServiceSid}/AlphaSenders"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AlphaSender != nil {
		data.Set("AlphaSender", *params.AlphaSender)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceAlphaSender{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateBrandRegistrationsParams Optional parameters for the method 'CreateBrandRegistrations'
type CreateBrandRegistrationsParams struct {
	A2pProfileBundleSid      *string `json:"A2pProfileBundleSid,omitempty"`
	CustomerProfileBundleSid *string `json:"CustomerProfileBundleSid,omitempty"`
}

/*
* CreateBrandRegistrations Method for CreateBrandRegistrations
* @param optional nil or *CreateBrandRegistrationsParams - Optional Parameters:
* @param "A2pProfileBundleSid" (string) - A2P Messaging Profile Bundle Sid.
* @param "CustomerProfileBundleSid" (string) - Customer Profile Bundle Sid.
* @return MessagingV1BrandRegistrations
 */
func (c *DefaultApiService) CreateBrandRegistrations(params *CreateBrandRegistrationsParams) (*MessagingV1BrandRegistrations, error) {
	path := "/v1/a2p/BrandRegistrations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.A2pProfileBundleSid != nil {
		data.Set("A2pProfileBundleSid", *params.A2pProfileBundleSid)
	}
	if params != nil && params.CustomerProfileBundleSid != nil {
		data.Set("CustomerProfileBundleSid", *params.CustomerProfileBundleSid)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandRegistrations{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateExternalCampaignParams Optional parameters for the method 'CreateExternalCampaign'
type CreateExternalCampaignParams struct {
	CampaignId          *string `json:"CampaignId,omitempty"`
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
}

/*
* CreateExternalCampaign Method for CreateExternalCampaign
* @param optional nil or *CreateExternalCampaignParams - Optional Parameters:
* @param "CampaignId" (string) - ID of the preregistered campaign.
* @param "MessagingServiceSid" (string) - The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) that the resource is associated with.
* @return MessagingV1ExternalCampaign
 */
func (c *DefaultApiService) CreateExternalCampaign(params *CreateExternalCampaignParams) (*MessagingV1ExternalCampaign, error) {
	path := "/v1/Services/PreregisteredUsa2p"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CampaignId != nil {
		data.Set("CampaignId", *params.CampaignId)
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ExternalCampaign{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreatePhoneNumberParams Optional parameters for the method 'CreatePhoneNumber'
type CreatePhoneNumberParams struct {
	PhoneNumberSid *string `json:"PhoneNumberSid,omitempty"`
}

/*
* CreatePhoneNumber Method for CreatePhoneNumber
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
* @param optional nil or *CreatePhoneNumberParams - Optional Parameters:
* @param "PhoneNumberSid" (string) - The SID of the Phone Number being added to the Service.
* @return MessagingV1ServicePhoneNumber
 */
func (c *DefaultApiService) CreatePhoneNumber(ServiceSid string, params *CreatePhoneNumberParams) (*MessagingV1ServicePhoneNumber, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PhoneNumberSid != nil {
		data.Set("PhoneNumberSid", *params.PhoneNumberSid)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServicePhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateServiceParams Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
	AreaCodeGeomatch          *bool   `json:"AreaCodeGeomatch,omitempty"`
	FallbackMethod            *string `json:"FallbackMethod,omitempty"`
	FallbackToLongCode        *bool   `json:"FallbackToLongCode,omitempty"`
	FallbackUrl               *string `json:"FallbackUrl,omitempty"`
	FriendlyName              *string `json:"FriendlyName,omitempty"`
	InboundMethod             *string `json:"InboundMethod,omitempty"`
	InboundRequestUrl         *string `json:"InboundRequestUrl,omitempty"`
	MmsConverter              *bool   `json:"MmsConverter,omitempty"`
	ScanMessageContent        *string `json:"ScanMessageContent,omitempty"`
	SmartEncoding             *bool   `json:"SmartEncoding,omitempty"`
	StatusCallback            *string `json:"StatusCallback,omitempty"`
	StickySender              *bool   `json:"StickySender,omitempty"`
	SynchronousValidation     *bool   `json:"SynchronousValidation,omitempty"`
	UseInboundWebhookOnNumber *bool   `json:"UseInboundWebhookOnNumber,omitempty"`
	ValidityPeriod            *int32  `json:"ValidityPeriod,omitempty"`
}

/*
* CreateService Method for CreateService
* @param optional nil or *CreateServiceParams - Optional Parameters:
* @param "AreaCodeGeomatch" (bool) - Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance.
* @param "FallbackMethod" (string) - The HTTP method we should use to call `fallback_url`. Can be: `GET` or `POST`.
* @param "FallbackToLongCode" (bool) - Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance.
* @param "FallbackUrl" (string) - The URL that we call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `fallback_url` defined for the Messaging Service.
* @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
* @param "InboundMethod" (string) - The HTTP method we should use to call `inbound_request_url`. Can be `GET` or `POST` and the default is `POST`.
* @param "InboundRequestUrl" (string) - The URL we call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled. All messages sent to the Twilio phone number or short code will not be logged and received on the Account. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `inbound_request_url` defined for the Messaging Service.
* @param "MmsConverter" (bool) - Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance.
* @param "ScanMessageContent" (string) - Reserved.
* @param "SmartEncoding" (bool) - Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance.
* @param "StatusCallback" (string) - The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
* @param "StickySender" (bool) - Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance.
* @param "SynchronousValidation" (bool) - Reserved.
* @param "UseInboundWebhookOnNumber" (bool) - A boolean value that indicates either the webhook url configured on the phone number will be used or `inbound_request_url`/`fallback_url` url will be called when a message is received from the phone number. If this field is enabled then the webhook url defined on the phone number will override the `inbound_request_url`/`fallback_url` defined for the Messaging Service.
* @param "ValidityPeriod" (int32) - How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`.
* @return MessagingV1Service
 */
func (c *DefaultApiService) CreateService(params *CreateServiceParams) (*MessagingV1Service, error) {
	path := "/v1/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AreaCodeGeomatch != nil {
		data.Set("AreaCodeGeomatch", fmt.Sprint(*params.AreaCodeGeomatch))
	}
	if params != nil && params.FallbackMethod != nil {
		data.Set("FallbackMethod", *params.FallbackMethod)
	}
	if params != nil && params.FallbackToLongCode != nil {
		data.Set("FallbackToLongCode", fmt.Sprint(*params.FallbackToLongCode))
	}
	if params != nil && params.FallbackUrl != nil {
		data.Set("FallbackUrl", *params.FallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.InboundMethod != nil {
		data.Set("InboundMethod", *params.InboundMethod)
	}
	if params != nil && params.InboundRequestUrl != nil {
		data.Set("InboundRequestUrl", *params.InboundRequestUrl)
	}
	if params != nil && params.MmsConverter != nil {
		data.Set("MmsConverter", fmt.Sprint(*params.MmsConverter))
	}
	if params != nil && params.ScanMessageContent != nil {
		data.Set("ScanMessageContent", *params.ScanMessageContent)
	}
	if params != nil && params.SmartEncoding != nil {
		data.Set("SmartEncoding", fmt.Sprint(*params.SmartEncoding))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StickySender != nil {
		data.Set("StickySender", fmt.Sprint(*params.StickySender))
	}
	if params != nil && params.SynchronousValidation != nil {
		data.Set("SynchronousValidation", fmt.Sprint(*params.SynchronousValidation))
	}
	if params != nil && params.UseInboundWebhookOnNumber != nil {
		data.Set("UseInboundWebhookOnNumber", fmt.Sprint(*params.UseInboundWebhookOnNumber))
	}
	if params != nil && params.ValidityPeriod != nil {
		data.Set("ValidityPeriod", fmt.Sprint(*params.ValidityPeriod))
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateShortCodeParams Optional parameters for the method 'CreateShortCode'
type CreateShortCodeParams struct {
	ShortCodeSid *string `json:"ShortCodeSid,omitempty"`
}

/*
* CreateShortCode Method for CreateShortCode
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
* @param optional nil or *CreateShortCodeParams - Optional Parameters:
* @param "ShortCodeSid" (string) - The SID of the ShortCode resource being added to the Service.
* @return MessagingV1ServiceShortCode
 */
func (c *DefaultApiService) CreateShortCode(ServiceSid string, params *CreateShortCodeParams) (*MessagingV1ServiceShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ShortCodeSid != nil {
		data.Set("ShortCodeSid", *params.ShortCodeSid)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateUsAppToPersonParams Optional parameters for the method 'CreateUsAppToPerson'
type CreateUsAppToPersonParams struct {
	BrandRegistrationSid *string   `json:"BrandRegistrationSid,omitempty"`
	Description          *string   `json:"Description,omitempty"`
	HasEmbeddedLinks     *bool     `json:"HasEmbeddedLinks,omitempty"`
	HasEmbeddedPhone     *bool     `json:"HasEmbeddedPhone,omitempty"`
	MessageSamples       *[]string `json:"MessageSamples,omitempty"`
	UsAppToPersonUsecase *string   `json:"UsAppToPersonUsecase,omitempty"`
}

/*
* CreateUsAppToPerson Method for CreateUsAppToPerson
* @param MessagingServiceSid The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to create the resources from.
* @param optional nil or *CreateUsAppToPersonParams - Optional Parameters:
* @param "BrandRegistrationSid" (string) - A2P Brand Registration SID
* @param "Description" (string) - A short description of what this SMS campaign does.
* @param "HasEmbeddedLinks" (bool) - Indicates that this SMS campaign will send messages that contain links.
* @param "HasEmbeddedPhone" (bool) - Indicates that this SMS campaign will send messages that contain phone numbers.
* @param "MessageSamples" ([]string) - Message samples, up to 5 sample messages, <=1024 chars each.
* @param "UsAppToPersonUsecase" (string) - A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..]
* @return MessagingV1ServiceUsAppToPerson
 */
func (c *DefaultApiService) CreateUsAppToPerson(MessagingServiceSid string, params *CreateUsAppToPersonParams) (*MessagingV1ServiceUsAppToPerson, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BrandRegistrationSid != nil {
		data.Set("BrandRegistrationSid", *params.BrandRegistrationSid)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.HasEmbeddedLinks != nil {
		data.Set("HasEmbeddedLinks", fmt.Sprint(*params.HasEmbeddedLinks))
	}
	if params != nil && params.HasEmbeddedPhone != nil {
		data.Set("HasEmbeddedPhone", fmt.Sprint(*params.HasEmbeddedPhone))
	}
	if params != nil && params.MessageSamples != nil {
		data.Set("MessageSamples", strings.Join(*params.MessageSamples, ","))
	}
	if params != nil && params.UsAppToPersonUsecase != nil {
		data.Set("UsAppToPersonUsecase", *params.UsAppToPersonUsecase)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceUsAppToPerson{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* DeleteAlphaSender Method for DeleteAlphaSender
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
* @param Sid The SID of the AlphaSender resource to delete.
 */
func (c *DefaultApiService) DeleteAlphaSender(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/AlphaSenders/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* DeletePhoneNumber Method for DeletePhoneNumber
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
* @param Sid The SID of the PhoneNumber resource to delete.
 */
func (c *DefaultApiService) DeletePhoneNumber(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* DeleteService Method for DeleteService
* @param Sid The SID of the Service resource to delete.
 */
func (c *DefaultApiService) DeleteService(Sid string) error {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* DeleteShortCode Method for DeleteShortCode
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
* @param Sid The SID of the ShortCode resource to delete.
 */
func (c *DefaultApiService) DeleteShortCode(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* DeleteUsAppToPerson Method for DeleteUsAppToPerson
* @param MessagingServiceSid The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to delete the resource from.
 */
func (c *DefaultApiService) DeleteUsAppToPerson(MessagingServiceSid string) error {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* FetchAlphaSender Method for FetchAlphaSender
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
* @param Sid The SID of the AlphaSender resource to fetch.
* @return MessagingV1ServiceAlphaSender
 */
func (c *DefaultApiService) FetchAlphaSender(ServiceSid string, Sid string) (*MessagingV1ServiceAlphaSender, error) {
	path := "/v1/Services/{ServiceSid}/AlphaSenders/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceAlphaSender{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchBrandRegistrations Method for FetchBrandRegistrations
* @param Sid The SID of the Brand Registration resource to fetch.
* @return MessagingV1BrandRegistrations
 */
func (c *DefaultApiService) FetchBrandRegistrations(Sid string) (*MessagingV1BrandRegistrations, error) {
	path := "/v1/a2p/BrandRegistrations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandRegistrations{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchDeactivationParams Optional parameters for the method 'FetchDeactivation'
type FetchDeactivationParams struct {
	Date *string `json:"Date,omitempty"`
}

/*
* FetchDeactivation Method for FetchDeactivation
* Fetch a list of all United States numbers that have been deactivated on a specific date.
* @param optional nil or *FetchDeactivationParams - Optional Parameters:
* @param "Date" (string) - The request will return a list of all United States Phone Numbers that were deactivated on the day specified by this parameter. This date should be specified in YYYY-MM-DD format.
 */
func (c *DefaultApiService) FetchDeactivation(params *FetchDeactivationParams) error {
	path := "/v1/Deactivations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Date != nil {
		data.Set("Date", fmt.Sprint(*params.Date))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* FetchPhoneNumber Method for FetchPhoneNumber
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
* @param Sid The SID of the PhoneNumber resource to fetch.
* @return MessagingV1ServicePhoneNumber
 */
func (c *DefaultApiService) FetchPhoneNumber(ServiceSid string, Sid string) (*MessagingV1ServicePhoneNumber, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServicePhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchService Method for FetchService
* @param Sid The SID of the Service resource to fetch.
* @return MessagingV1Service
 */
func (c *DefaultApiService) FetchService(Sid string) (*MessagingV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchShortCode Method for FetchShortCode
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
* @param Sid The SID of the ShortCode resource to fetch.
* @return MessagingV1ServiceShortCode
 */
func (c *DefaultApiService) FetchShortCode(ServiceSid string, Sid string) (*MessagingV1ServiceShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchUsAppToPerson Method for FetchUsAppToPerson
* @param MessagingServiceSid The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to fetch the resource from.
* @return MessagingV1ServiceUsAppToPerson
 */
func (c *DefaultApiService) FetchUsAppToPerson(MessagingServiceSid string) (*MessagingV1ServiceUsAppToPerson, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceUsAppToPerson{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchUsAppToPersonUsecase Method for FetchUsAppToPersonUsecase
* @param MessagingServiceSid The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to fetch the resource from.
* @return MessagingV1ServiceUsAppToPersonUsecase
 */
func (c *DefaultApiService) FetchUsAppToPersonUsecase(MessagingServiceSid string) (*MessagingV1ServiceUsAppToPersonUsecase, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/Usecases"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceUsAppToPersonUsecase{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchUsecase Method for FetchUsecase
* @return MessagingV1Usecase
 */
func (c *DefaultApiService) FetchUsecase() (*MessagingV1Usecase, error) {
	path := "/v1/Services/Usecases"

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1Usecase{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListAlphaSenderParams Optional parameters for the method 'ListAlphaSender'
type ListAlphaSenderParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListAlphaSender Method for ListAlphaSender
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.
* @param optional nil or *ListAlphaSenderParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListAlphaSenderResponse
 */
func (c *DefaultApiService) ListAlphaSender(ServiceSid string, params *ListAlphaSenderParams) (*ListAlphaSenderResponse, error) {
	path := "/v1/Services/{ServiceSid}/AlphaSenders"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAlphaSenderResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListBrandRegistrationsParams Optional parameters for the method 'ListBrandRegistrations'
type ListBrandRegistrationsParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListBrandRegistrations Method for ListBrandRegistrations
* @param optional nil or *ListBrandRegistrationsParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListBrandRegistrationsResponse
 */
func (c *DefaultApiService) ListBrandRegistrations(params *ListBrandRegistrationsParams) (*ListBrandRegistrationsResponse, error) {
	path := "/v1/a2p/BrandRegistrations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBrandRegistrationsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListPhoneNumberParams Optional parameters for the method 'ListPhoneNumber'
type ListPhoneNumberParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListPhoneNumber Method for ListPhoneNumber
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.
* @param optional nil or *ListPhoneNumberParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListPhoneNumberResponse
 */
func (c *DefaultApiService) ListPhoneNumber(ServiceSid string, params *ListPhoneNumberParams) (*ListPhoneNumberResponse, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListServiceParams Optional parameters for the method 'ListService'
type ListServiceParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListService Method for ListService
* @param optional nil or *ListServiceParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListServiceResponse
 */
func (c *DefaultApiService) ListService(params *ListServiceParams) (*ListServiceResponse, error) {
	path := "/v1/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListShortCodeParams Optional parameters for the method 'ListShortCode'
type ListShortCodeParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListShortCode Method for ListShortCode
* @param ServiceSid The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.
* @param optional nil or *ListShortCodeParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListShortCodeResponse
 */
func (c *DefaultApiService) ListShortCode(ServiceSid string, params *ListShortCodeParams) (*ListShortCodeResponse, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateServiceParams Optional parameters for the method 'UpdateService'
type UpdateServiceParams struct {
	AreaCodeGeomatch          *bool   `json:"AreaCodeGeomatch,omitempty"`
	FallbackMethod            *string `json:"FallbackMethod,omitempty"`
	FallbackToLongCode        *bool   `json:"FallbackToLongCode,omitempty"`
	FallbackUrl               *string `json:"FallbackUrl,omitempty"`
	FriendlyName              *string `json:"FriendlyName,omitempty"`
	InboundMethod             *string `json:"InboundMethod,omitempty"`
	InboundRequestUrl         *string `json:"InboundRequestUrl,omitempty"`
	MmsConverter              *bool   `json:"MmsConverter,omitempty"`
	ScanMessageContent        *string `json:"ScanMessageContent,omitempty"`
	SmartEncoding             *bool   `json:"SmartEncoding,omitempty"`
	StatusCallback            *string `json:"StatusCallback,omitempty"`
	StickySender              *bool   `json:"StickySender,omitempty"`
	SynchronousValidation     *bool   `json:"SynchronousValidation,omitempty"`
	UseInboundWebhookOnNumber *bool   `json:"UseInboundWebhookOnNumber,omitempty"`
	ValidityPeriod            *int32  `json:"ValidityPeriod,omitempty"`
}

/*
* UpdateService Method for UpdateService
* @param Sid The SID of the Service resource to update.
* @param optional nil or *UpdateServiceParams - Optional Parameters:
* @param "AreaCodeGeomatch" (bool) - Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance.
* @param "FallbackMethod" (string) - The HTTP method we should use to call `fallback_url`. Can be: `GET` or `POST`.
* @param "FallbackToLongCode" (bool) - Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance.
* @param "FallbackUrl" (string) - The URL that we call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `fallback_url` defined for the Messaging Service.
* @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
* @param "InboundMethod" (string) - The HTTP method we should use to call `inbound_request_url`. Can be `GET` or `POST` and the default is `POST`.
* @param "InboundRequestUrl" (string) - The URL we call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled. All messages sent to the Twilio phone number or short code will not be logged and received on the Account. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `inbound_request_url` defined for the Messaging Service.
* @param "MmsConverter" (bool) - Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance.
* @param "ScanMessageContent" (string) - Reserved.
* @param "SmartEncoding" (bool) - Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance.
* @param "StatusCallback" (string) - The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
* @param "StickySender" (bool) - Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance.
* @param "SynchronousValidation" (bool) - Reserved.
* @param "UseInboundWebhookOnNumber" (bool) - A boolean value that indicates either the webhook url configured on the phone number will be used or `inbound_request_url`/`fallback_url` url will be called when a message is received from the phone number. If this field is enabled then the webhook url defined on the phone number will override the `inbound_request_url`/`fallback_url` defined for the Messaging Service.
* @param "ValidityPeriod" (int32) - How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`.
* @return MessagingV1Service
 */
func (c *DefaultApiService) UpdateService(Sid string, params *UpdateServiceParams) (*MessagingV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AreaCodeGeomatch != nil {
		data.Set("AreaCodeGeomatch", fmt.Sprint(*params.AreaCodeGeomatch))
	}
	if params != nil && params.FallbackMethod != nil {
		data.Set("FallbackMethod", *params.FallbackMethod)
	}
	if params != nil && params.FallbackToLongCode != nil {
		data.Set("FallbackToLongCode", fmt.Sprint(*params.FallbackToLongCode))
	}
	if params != nil && params.FallbackUrl != nil {
		data.Set("FallbackUrl", *params.FallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.InboundMethod != nil {
		data.Set("InboundMethod", *params.InboundMethod)
	}
	if params != nil && params.InboundRequestUrl != nil {
		data.Set("InboundRequestUrl", *params.InboundRequestUrl)
	}
	if params != nil && params.MmsConverter != nil {
		data.Set("MmsConverter", fmt.Sprint(*params.MmsConverter))
	}
	if params != nil && params.ScanMessageContent != nil {
		data.Set("ScanMessageContent", *params.ScanMessageContent)
	}
	if params != nil && params.SmartEncoding != nil {
		data.Set("SmartEncoding", fmt.Sprint(*params.SmartEncoding))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StickySender != nil {
		data.Set("StickySender", fmt.Sprint(*params.StickySender))
	}
	if params != nil && params.SynchronousValidation != nil {
		data.Set("SynchronousValidation", fmt.Sprint(*params.SynchronousValidation))
	}
	if params != nil && params.UseInboundWebhookOnNumber != nil {
		data.Set("UseInboundWebhookOnNumber", fmt.Sprint(*params.UseInboundWebhookOnNumber))
	}
	if params != nil && params.ValidityPeriod != nil {
		data.Set("ValidityPeriod", fmt.Sprint(*params.ValidityPeriod))
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
