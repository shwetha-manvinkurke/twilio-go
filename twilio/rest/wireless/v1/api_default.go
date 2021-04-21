/*
 * Twilio - Wireless
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
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://wireless.twilio.com",
	}
}

// CreateCommandParams Optional parameters for the method 'CreateCommand'
type CreateCommandParams struct {
	CallbackMethod           *string `json:"CallbackMethod,omitempty"`
	CallbackUrl              *string `json:"CallbackUrl,omitempty"`
	Command                  *string `json:"Command,omitempty"`
	CommandMode              *string `json:"CommandMode,omitempty"`
	DeliveryReceiptRequested *bool   `json:"DeliveryReceiptRequested,omitempty"`
	IncludeSid               *string `json:"IncludeSid,omitempty"`
	Sim                      *string `json:"Sim,omitempty"`
}

/*
* CreateCommand Method for CreateCommand
* Send a Command to a Sim.
* @param optional nil or *CreateCommandParams - Optional Parameters:
* @param "CallbackMethod" (string) - The HTTP method we use to call `callback_url`. Can be: `POST` or `GET`, and the default is `POST`.
* @param "CallbackUrl" (string) - The URL we call using the `callback_url` when the Command has finished sending, whether the command was delivered or it failed.
* @param "Command" (string) - The message body of the Command. Can be plain text in text mode or a Base64 encoded byte string in binary mode.
* @param "CommandMode" (string) - The mode to use when sending the SMS message. Can be: `text` or `binary`. The default SMS mode is `text`.
* @param "DeliveryReceiptRequested" (bool) - Whether to request delivery receipt from the recipient. For Commands that request delivery receipt, the Command state transitions to 'delivered' once the server has received a delivery receipt from the device. The default value is `true`.
* @param "IncludeSid" (string) - Whether to include the SID of the command in the message body. Can be: `none`, `start`, or `end`, and the default behavior is `none`. When sending a Command to a SIM in text mode, we can automatically include the SID of the Command in the message body, which could be used to ensure that the device does not process the same Command more than once.  A value of `start` will prepend the message with the Command SID, and `end` will append it to the end, separating the Command SID from the message body with a space. The length of the Command SID is included in the 160 character limit so the SMS body must be 128 characters or less before the Command SID is included.
* @param "Sim" (string) - The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
* @return WirelessV1Command
 */
func (c *DefaultApiService) CreateCommand(params *CreateCommandParams) (*WirelessV1Command, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Command != nil {
		data.Set("Command", *params.Command)
	}
	if params != nil && params.CommandMode != nil {
		data.Set("CommandMode", *params.CommandMode)
	}
	if params != nil && params.DeliveryReceiptRequested != nil {
		data.Set("DeliveryReceiptRequested", fmt.Sprint(*params.DeliveryReceiptRequested))
	}
	if params != nil && params.IncludeSid != nil {
		data.Set("IncludeSid", *params.IncludeSid)
	}
	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateRatePlanParams Optional parameters for the method 'CreateRatePlan'
type CreateRatePlanParams struct {
	DataEnabled                   *bool     `json:"DataEnabled,omitempty"`
	DataLimit                     *int32    `json:"DataLimit,omitempty"`
	DataMetering                  *string   `json:"DataMetering,omitempty"`
	FriendlyName                  *string   `json:"FriendlyName,omitempty"`
	InternationalRoaming          *[]string `json:"InternationalRoaming,omitempty"`
	InternationalRoamingDataLimit *int32    `json:"InternationalRoamingDataLimit,omitempty"`
	MessagingEnabled              *bool     `json:"MessagingEnabled,omitempty"`
	NationalRoamingDataLimit      *int32    `json:"NationalRoamingDataLimit,omitempty"`
	NationalRoamingEnabled        *bool     `json:"NationalRoamingEnabled,omitempty"`
	UniqueName                    *string   `json:"UniqueName,omitempty"`
	VoiceEnabled                  *bool     `json:"VoiceEnabled,omitempty"`
}

/*
* CreateRatePlan Method for CreateRatePlan
* @param optional nil or *CreateRatePlanParams - Optional Parameters:
* @param "DataEnabled" (bool) - Whether SIMs can use GPRS/3G/4G/LTE data connectivity.
* @param "DataLimit" (int32) - The total data usage (download and upload combined) in Megabytes that the Network allows during one month on the home network (T-Mobile USA). The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB and the default value is `1000`.
* @param "DataMetering" (string) - The model used to meter data usage. Can be: `payg` and `quota-1`, `quota-10`, and `quota-50`. Learn more about the available [data metering models](https://www.twilio.com/docs/wireless/api/rateplan-resource#payg-vs-quota-data-plans).
* @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It does not have to be unique.
* @param "InternationalRoaming" ([]string) - The list of services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States. Can be: `data`, `voice`, and `messaging`.
* @param "InternationalRoamingDataLimit" (int32) - The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States. Can be up to 2TB.
* @param "MessagingEnabled" (bool) - Whether SIMs can make, send, and receive SMS using [Commands](https://www.twilio.com/docs/wireless/api/command-resource).
* @param "NationalRoamingDataLimit" (int32) - The total data usage (download and upload combined) in Megabytes that the Network allows during one month on non-home networks in the United States. The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming) for more info.
* @param "NationalRoamingEnabled" (bool) - Whether SIMs can roam on networks other than the home network (T-Mobile USA) in the United States. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming).
* @param "UniqueName" (string) - An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
* @param "VoiceEnabled" (bool) - Whether SIMs can make and receive voice calls.
* @return WirelessV1RatePlan
 */
func (c *DefaultApiService) CreateRatePlan(params *CreateRatePlanParams) (*WirelessV1RatePlan, error) {
	path := "/v1/RatePlans"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DataEnabled != nil {
		data.Set("DataEnabled", fmt.Sprint(*params.DataEnabled))
	}
	if params != nil && params.DataLimit != nil {
		data.Set("DataLimit", fmt.Sprint(*params.DataLimit))
	}
	if params != nil && params.DataMetering != nil {
		data.Set("DataMetering", *params.DataMetering)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.InternationalRoaming != nil {
		data.Set("InternationalRoaming", strings.Join(*params.InternationalRoaming, ","))
	}
	if params != nil && params.InternationalRoamingDataLimit != nil {
		data.Set("InternationalRoamingDataLimit", fmt.Sprint(*params.InternationalRoamingDataLimit))
	}
	if params != nil && params.MessagingEnabled != nil {
		data.Set("MessagingEnabled", fmt.Sprint(*params.MessagingEnabled))
	}
	if params != nil && params.NationalRoamingDataLimit != nil {
		data.Set("NationalRoamingDataLimit", fmt.Sprint(*params.NationalRoamingDataLimit))
	}
	if params != nil && params.NationalRoamingEnabled != nil {
		data.Set("NationalRoamingEnabled", fmt.Sprint(*params.NationalRoamingEnabled))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.VoiceEnabled != nil {
		data.Set("VoiceEnabled", fmt.Sprint(*params.VoiceEnabled))
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1RatePlan{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* DeleteCommand Method for DeleteCommand
* Delete a Command instance from your account.
* @param Sid The SID of the Command resource to delete.
 */
func (c *DefaultApiService) DeleteCommand(Sid string) error {
	path := "/v1/Commands/{Sid}"
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
* DeleteRatePlan Method for DeleteRatePlan
* @param Sid The SID of the RatePlan resource to delete.
 */
func (c *DefaultApiService) DeleteRatePlan(Sid string) error {
	path := "/v1/RatePlans/{Sid}"
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
* DeleteSim Method for DeleteSim
* Delete a Sim resource on your Account.
* @param Sid The SID or the `unique_name` of the Sim resource to delete.
 */
func (c *DefaultApiService) DeleteSim(Sid string) error {
	path := "/v1/Sims/{Sid}"
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
* FetchCommand Method for FetchCommand
* Fetch a Command instance from your account.
* @param Sid The SID of the Command resource to fetch.
* @return WirelessV1Command
 */
func (c *DefaultApiService) FetchCommand(Sid string) (*WirelessV1Command, error) {
	path := "/v1/Commands/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchRatePlan Method for FetchRatePlan
* @param Sid The SID of the RatePlan resource to fetch.
* @return WirelessV1RatePlan
 */
func (c *DefaultApiService) FetchRatePlan(Sid string) (*WirelessV1RatePlan, error) {
	path := "/v1/RatePlans/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1RatePlan{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchSim Method for FetchSim
* Fetch a Sim resource on your Account.
* @param Sid The SID or the `unique_name` of the Sim resource to fetch.
* @return WirelessV1Sim
 */
func (c *DefaultApiService) FetchSim(Sid string) (*WirelessV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListAccountUsageRecordParams Optional parameters for the method 'ListAccountUsageRecord'
type ListAccountUsageRecordParams struct {
	End         *time.Time `json:"End,omitempty"`
	Start       *time.Time `json:"Start,omitempty"`
	Granularity *string    `json:"Granularity,omitempty"`
	PageSize    *int32     `json:"PageSize,omitempty"`
}

/*
* ListAccountUsageRecord Method for ListAccountUsageRecord
* @param optional nil or *ListAccountUsageRecordParams - Optional Parameters:
* @param "End" (time.Time) - Only include usage that has occurred on or before this date. Format is [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html).
* @param "Start" (time.Time) - Only include usage that has occurred on or after this date. Format is [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html).
* @param "Granularity" (string) - How to summarize the usage by time. Can be: `daily`, `hourly`, or `all`. A value of `all` returns one Usage Record that describes the usage for the entire period.
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListAccountUsageRecordResponse
 */
func (c *DefaultApiService) ListAccountUsageRecord(params *ListAccountUsageRecordParams) (*ListAccountUsageRecordResponse, error) {
	path := "/v1/UsageRecords"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.End != nil {
		data.Set("End", fmt.Sprint((*params.End).Format(time.RFC3339)))
	}
	if params != nil && params.Start != nil {
		data.Set("Start", fmt.Sprint((*params.Start).Format(time.RFC3339)))
	}
	if params != nil && params.Granularity != nil {
		data.Set("Granularity", *params.Granularity)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAccountUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListCommandParams Optional parameters for the method 'ListCommand'
type ListCommandParams struct {
	Sim       *string `json:"Sim,omitempty"`
	Status    *string `json:"Status,omitempty"`
	Direction *string `json:"Direction,omitempty"`
	Transport *string `json:"Transport,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty"`
}

/*
* ListCommand Method for ListCommand
* Retrieve a list of Commands from your account.
* @param optional nil or *ListCommandParams - Optional Parameters:
* @param "Sim" (string) - The `sid` or `unique_name` of the [Sim resources](https://www.twilio.com/docs/wireless/api/sim-resource) to read.
* @param "Status" (string) - The status of the resources to read. Can be: `queued`, `sent`, `delivered`, `received`, or `failed`.
* @param "Direction" (string) - Only return Commands with this direction value.
* @param "Transport" (string) - Only return Commands with this transport value. Can be: `sms` or `ip`.
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListCommandResponse
 */
func (c *DefaultApiService) ListCommand(params *ListCommandParams) (*ListCommandResponse, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.Transport != nil {
		data.Set("Transport", *params.Transport)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListDataSessionParams Optional parameters for the method 'ListDataSession'
type ListDataSessionParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListDataSession Method for ListDataSession
* @param SimSid The SID of the [Sim resource](https://www.twilio.com/docs/wireless/api/sim-resource) with the Data Sessions to read.
* @param optional nil or *ListDataSessionParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListDataSessionResponse
 */
func (c *DefaultApiService) ListDataSession(SimSid string, params *ListDataSessionParams) (*ListDataSessionResponse, error) {
	path := "/v1/Sims/{SimSid}/DataSessions"
	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

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

	ps := &ListDataSessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListRatePlanParams Optional parameters for the method 'ListRatePlan'
type ListRatePlanParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListRatePlan Method for ListRatePlan
* @param optional nil or *ListRatePlanParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListRatePlanResponse
 */
func (c *DefaultApiService) ListRatePlan(params *ListRatePlanParams) (*ListRatePlanResponse, error) {
	path := "/v1/RatePlans"

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

	ps := &ListRatePlanResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListSimParams Optional parameters for the method 'ListSim'
type ListSimParams struct {
	Status              *string `json:"Status,omitempty"`
	Iccid               *string `json:"Iccid,omitempty"`
	RatePlan            *string `json:"RatePlan,omitempty"`
	EId                 *string `json:"EId,omitempty"`
	SimRegistrationCode *string `json:"SimRegistrationCode,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty"`
}

/*
* ListSim Method for ListSim
* Retrieve a list of Sim resources on your Account.
* @param optional nil or *ListSimParams - Optional Parameters:
* @param "Status" (string) - Only return Sim resources with this status.
* @param "Iccid" (string) - Only return Sim resources with this ICCID. This will return a list with a maximum size of 1.
* @param "RatePlan" (string) - The SID or unique name of a [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource). Only return Sim resources assigned to this RatePlan resource.
* @param "EId" (string) - Deprecated.
* @param "SimRegistrationCode" (string) - Only return Sim resources with this registration code. This will return a list with a maximum size of 1.
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListSimResponse
 */
func (c *DefaultApiService) ListSim(params *ListSimParams) (*ListSimResponse, error) {
	path := "/v1/Sims"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Iccid != nil {
		data.Set("Iccid", *params.Iccid)
	}
	if params != nil && params.RatePlan != nil {
		data.Set("RatePlan", *params.RatePlan)
	}
	if params != nil && params.EId != nil {
		data.Set("EId", *params.EId)
	}
	if params != nil && params.SimRegistrationCode != nil {
		data.Set("SimRegistrationCode", *params.SimRegistrationCode)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSimResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListUsageRecordParams Optional parameters for the method 'ListUsageRecord'
type ListUsageRecordParams struct {
	End         *time.Time `json:"End,omitempty"`
	Start       *time.Time `json:"Start,omitempty"`
	Granularity *string    `json:"Granularity,omitempty"`
	PageSize    *int32     `json:"PageSize,omitempty"`
}

/*
* ListUsageRecord Method for ListUsageRecord
* @param SimSid The SID of the [Sim resource](https://www.twilio.com/docs/wireless/api/sim-resource)  to read the usage from.
* @param optional nil or *ListUsageRecordParams - Optional Parameters:
* @param "End" (time.Time) - Only include usage that occurred on or before this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is the current time.
* @param "Start" (time.Time) - Only include usage that has occurred on or after this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is one month before the `end` parameter value.
* @param "Granularity" (string) - How to summarize the usage by time. Can be: `daily`, `hourly`, or `all`. The default is `all`. A value of `all` returns one Usage Record that describes the usage for the entire period.
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListUsageRecordResponse
 */
func (c *DefaultApiService) ListUsageRecord(SimSid string, params *ListUsageRecordParams) (*ListUsageRecordResponse, error) {
	path := "/v1/Sims/{SimSid}/UsageRecords"
	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.End != nil {
		data.Set("End", fmt.Sprint((*params.End).Format(time.RFC3339)))
	}
	if params != nil && params.Start != nil {
		data.Set("Start", fmt.Sprint((*params.Start).Format(time.RFC3339)))
	}
	if params != nil && params.Granularity != nil {
		data.Set("Granularity", *params.Granularity)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateRatePlanParams Optional parameters for the method 'UpdateRatePlan'
type UpdateRatePlanParams struct {
	FriendlyName *string `json:"FriendlyName,omitempty"`
	UniqueName   *string `json:"UniqueName,omitempty"`
}

/*
* UpdateRatePlan Method for UpdateRatePlan
* @param Sid The SID of the RatePlan resource to update.
* @param optional nil or *UpdateRatePlanParams - Optional Parameters:
* @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It does not have to be unique.
* @param "UniqueName" (string) - An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
* @return WirelessV1RatePlan
 */
func (c *DefaultApiService) UpdateRatePlan(Sid string, params *UpdateRatePlanParams) (*WirelessV1RatePlan, error) {
	path := "/v1/RatePlans/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1RatePlan{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateSimParams Optional parameters for the method 'UpdateSim'
type UpdateSimParams struct {
	AccountSid             *string `json:"AccountSid,omitempty"`
	CallbackMethod         *string `json:"CallbackMethod,omitempty"`
	CallbackUrl            *string `json:"CallbackUrl,omitempty"`
	CommandsCallbackMethod *string `json:"CommandsCallbackMethod,omitempty"`
	CommandsCallbackUrl    *string `json:"CommandsCallbackUrl,omitempty"`
	FriendlyName           *string `json:"FriendlyName,omitempty"`
	RatePlan               *string `json:"RatePlan,omitempty"`
	ResetStatus            *string `json:"ResetStatus,omitempty"`
	SmsFallbackMethod      *string `json:"SmsFallbackMethod,omitempty"`
	SmsFallbackUrl         *string `json:"SmsFallbackUrl,omitempty"`
	SmsMethod              *string `json:"SmsMethod,omitempty"`
	SmsUrl                 *string `json:"SmsUrl,omitempty"`
	Status                 *string `json:"Status,omitempty"`
	UniqueName             *string `json:"UniqueName,omitempty"`
	VoiceFallbackMethod    *string `json:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackUrl       *string `json:"VoiceFallbackUrl,omitempty"`
	VoiceMethod            *string `json:"VoiceMethod,omitempty"`
	VoiceUrl               *string `json:"VoiceUrl,omitempty"`
}

/*
* UpdateSim Method for UpdateSim
* Updates the given properties of a Sim resource on your Account.
* @param Sid The SID or the `unique_name` of the Sim resource to update.
* @param optional nil or *UpdateSimParams - Optional Parameters:
* @param "AccountSid" (string) - The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a [Subaccount](https://www.twilio.com/docs/iam/api/subaccounts) of the requesting Account. Only valid when the Sim resource's status is `new`. For more information, see the [Move SIMs between Subaccounts documentation](https://www.twilio.com/docs/wireless/api/sim-resource#move-sims-between-subaccounts).
* @param "CallbackMethod" (string) - The HTTP method we should use to call `callback_url`. Can be: `POST` or `GET`. The default is `POST`.
* @param "CallbackUrl" (string) - The URL we should call using the `callback_url` when the SIM has finished updating. When the SIM transitions from `new` to `ready` or from any status to `deactivated`, we call this URL when the status changes to an intermediate status (`ready` or `deactivated`) and again when the status changes to its final status (`active` or `canceled`).
* @param "CommandsCallbackMethod" (string) - The HTTP method we should use to call `commands_callback_url`. Can be: `POST` or `GET`. The default is `POST`.
* @param "CommandsCallbackUrl" (string) - The URL we should call using the `commands_callback_method` when the SIM sends a [Command](https://www.twilio.com/docs/wireless/api/command-resource). Your server should respond with an HTTP status code in the 200 range; any response body is ignored.
* @param "FriendlyName" (string) - A descriptive string that you create to describe the Sim resource. It does not need to be unique.
* @param "RatePlan" (string) - The SID or unique name of the [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource) to which the Sim resource should be assigned.
* @param "ResetStatus" (string) - Initiate a connectivity reset on the SIM. Set to `resetting` to initiate a connectivity reset on the SIM. No other value is valid.
* @param "SmsFallbackMethod" (string) - The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`. Default is `POST`.
* @param "SmsFallbackUrl" (string) - The URL we should call using the `sms_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `sms_url`.
* @param "SmsMethod" (string) - The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`. Default is `POST`.
* @param "SmsUrl" (string) - The URL we should call using the `sms_method` when the SIM-connected device sends an SMS message that is not a [Command](https://www.twilio.com/docs/wireless/api/command-resource).
* @param "Status" (string) - The new status of the Sim resource. Can be: `ready`, `active`, `suspended`, or `deactivated`.
* @param "UniqueName" (string) - An application-defined string that uniquely identifies the resource. It can be used in place of the `sid` in the URL path to address the resource.
* @param "VoiceFallbackMethod" (string) - The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
* @param "VoiceFallbackUrl" (string) - The URL we should call using the `voice_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `voice_url`.
* @param "VoiceMethod" (string) - The HTTP method we should use when we call `voice_url`. Can be: `GET` or `POST`.
* @param "VoiceUrl" (string) - The URL we should call using the `voice_method` when the SIM-connected device makes a voice call.
* @return WirelessV1Sim
 */
func (c *DefaultApiService) UpdateSim(Sid string, params *UpdateSimParams) (*WirelessV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.CommandsCallbackMethod != nil {
		data.Set("CommandsCallbackMethod", *params.CommandsCallbackMethod)
	}
	if params != nil && params.CommandsCallbackUrl != nil {
		data.Set("CommandsCallbackUrl", *params.CommandsCallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.RatePlan != nil {
		data.Set("RatePlan", *params.RatePlan)
	}
	if params != nil && params.ResetStatus != nil {
		data.Set("ResetStatus", *params.ResetStatus)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
