/*
Alerting

OpenAPI schema for alerting endpoints

API version: 0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerting

import (
	"encoding/json"
)

// checks if the ActionsInnerAlertsFilterTimeframe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionsInnerAlertsFilterTimeframe{}

// ActionsInnerAlertsFilterTimeframe Defines a period that limits whether the action runs.
type ActionsInnerAlertsFilterTimeframe struct {
	Days  []int32                                 `json:"days,omitempty"`
	Hours *ActionsInnerAlertsFilterTimeframeHours `json:"hours,omitempty"`
	// The ISO time zone for the `hours` values. Values such as `UTC` and `UTC+1` also work but lack built-in daylight savings time support and are not recommended.
	Timezone *string `json:"timezone,omitempty"`
}

// NewActionsInnerAlertsFilterTimeframe instantiates a new ActionsInnerAlertsFilterTimeframe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsInnerAlertsFilterTimeframe() *ActionsInnerAlertsFilterTimeframe {
	this := ActionsInnerAlertsFilterTimeframe{}
	return &this
}

// NewActionsInnerAlertsFilterTimeframeWithDefaults instantiates a new ActionsInnerAlertsFilterTimeframe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsInnerAlertsFilterTimeframeWithDefaults() *ActionsInnerAlertsFilterTimeframe {
	this := ActionsInnerAlertsFilterTimeframe{}
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *ActionsInnerAlertsFilterTimeframe) GetDays() []int32 {
	if o == nil || IsNil(o.Days) {
		var ret []int32
		return ret
	}
	return o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsInnerAlertsFilterTimeframe) GetDaysOk() ([]int32, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *ActionsInnerAlertsFilterTimeframe) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given []int32 and assigns it to the Days field.
func (o *ActionsInnerAlertsFilterTimeframe) SetDays(v []int32) {
	o.Days = v
}

// GetHours returns the Hours field value if set, zero value otherwise.
func (o *ActionsInnerAlertsFilterTimeframe) GetHours() ActionsInnerAlertsFilterTimeframeHours {
	if o == nil || IsNil(o.Hours) {
		var ret ActionsInnerAlertsFilterTimeframeHours
		return ret
	}
	return *o.Hours
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsInnerAlertsFilterTimeframe) GetHoursOk() (*ActionsInnerAlertsFilterTimeframeHours, bool) {
	if o == nil || IsNil(o.Hours) {
		return nil, false
	}
	return o.Hours, true
}

// HasHours returns a boolean if a field has been set.
func (o *ActionsInnerAlertsFilterTimeframe) HasHours() bool {
	if o != nil && !IsNil(o.Hours) {
		return true
	}

	return false
}

// SetHours gets a reference to the given ActionsInnerAlertsFilterTimeframeHours and assigns it to the Hours field.
func (o *ActionsInnerAlertsFilterTimeframe) SetHours(v ActionsInnerAlertsFilterTimeframeHours) {
	o.Hours = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ActionsInnerAlertsFilterTimeframe) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsInnerAlertsFilterTimeframe) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ActionsInnerAlertsFilterTimeframe) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ActionsInnerAlertsFilterTimeframe) SetTimezone(v string) {
	o.Timezone = &v
}

func (o ActionsInnerAlertsFilterTimeframe) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionsInnerAlertsFilterTimeframe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	if !IsNil(o.Hours) {
		toSerialize["hours"] = o.Hours
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	return toSerialize, nil
}

type NullableActionsInnerAlertsFilterTimeframe struct {
	value *ActionsInnerAlertsFilterTimeframe
	isSet bool
}

func (v NullableActionsInnerAlertsFilterTimeframe) Get() *ActionsInnerAlertsFilterTimeframe {
	return v.value
}

func (v *NullableActionsInnerAlertsFilterTimeframe) Set(val *ActionsInnerAlertsFilterTimeframe) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsInnerAlertsFilterTimeframe) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsInnerAlertsFilterTimeframe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsInnerAlertsFilterTimeframe(val *ActionsInnerAlertsFilterTimeframe) *NullableActionsInnerAlertsFilterTimeframe {
	return &NullableActionsInnerAlertsFilterTimeframe{value: val, isSet: true}
}

func (v NullableActionsInnerAlertsFilterTimeframe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsInnerAlertsFilterTimeframe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
