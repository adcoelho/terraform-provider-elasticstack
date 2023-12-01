/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the IndicatorPropertiesTimesliceMetricParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicatorPropertiesTimesliceMetricParams{}

// IndicatorPropertiesTimesliceMetricParams An object containing the indicator parameters.
type IndicatorPropertiesTimesliceMetricParams struct {
	// The index or index pattern to use
	Index string `json:"index"`
	// the KQL query to filter the documents with.
	Filter *string `json:"filter,omitempty"`
	// The timestamp field used in the source indice.
	TimestampField string                                         `json:"timestampField"`
	Metric         IndicatorPropertiesTimesliceMetricParamsMetric `json:"metric"`
}

// NewIndicatorPropertiesTimesliceMetricParams instantiates a new IndicatorPropertiesTimesliceMetricParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorPropertiesTimesliceMetricParams(index string, timestampField string, metric IndicatorPropertiesTimesliceMetricParamsMetric) *IndicatorPropertiesTimesliceMetricParams {
	this := IndicatorPropertiesTimesliceMetricParams{}
	this.Index = index
	this.TimestampField = timestampField
	this.Metric = metric
	return &this
}

// NewIndicatorPropertiesTimesliceMetricParamsWithDefaults instantiates a new IndicatorPropertiesTimesliceMetricParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorPropertiesTimesliceMetricParamsWithDefaults() *IndicatorPropertiesTimesliceMetricParams {
	this := IndicatorPropertiesTimesliceMetricParams{}
	return &this
}

// GetIndex returns the Index field value
func (o *IndicatorPropertiesTimesliceMetricParams) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesTimesliceMetricParams) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *IndicatorPropertiesTimesliceMetricParams) SetIndex(v string) {
	o.Index = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *IndicatorPropertiesTimesliceMetricParams) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesTimesliceMetricParams) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *IndicatorPropertiesTimesliceMetricParams) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *IndicatorPropertiesTimesliceMetricParams) SetFilter(v string) {
	o.Filter = &v
}

// GetTimestampField returns the TimestampField field value
func (o *IndicatorPropertiesTimesliceMetricParams) GetTimestampField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimestampField
}

// GetTimestampFieldOk returns a tuple with the TimestampField field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesTimesliceMetricParams) GetTimestampFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampField, true
}

// SetTimestampField sets field value
func (o *IndicatorPropertiesTimesliceMetricParams) SetTimestampField(v string) {
	o.TimestampField = v
}

// GetMetric returns the Metric field value
func (o *IndicatorPropertiesTimesliceMetricParams) GetMetric() IndicatorPropertiesTimesliceMetricParamsMetric {
	if o == nil {
		var ret IndicatorPropertiesTimesliceMetricParamsMetric
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesTimesliceMetricParams) GetMetricOk() (*IndicatorPropertiesTimesliceMetricParamsMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *IndicatorPropertiesTimesliceMetricParams) SetMetric(v IndicatorPropertiesTimesliceMetricParamsMetric) {
	o.Metric = v
}

func (o IndicatorPropertiesTimesliceMetricParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicatorPropertiesTimesliceMetricParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["timestampField"] = o.TimestampField
	toSerialize["metric"] = o.Metric
	return toSerialize, nil
}

type NullableIndicatorPropertiesTimesliceMetricParams struct {
	value *IndicatorPropertiesTimesliceMetricParams
	isSet bool
}

func (v NullableIndicatorPropertiesTimesliceMetricParams) Get() *IndicatorPropertiesTimesliceMetricParams {
	return v.value
}

func (v *NullableIndicatorPropertiesTimesliceMetricParams) Set(val *IndicatorPropertiesTimesliceMetricParams) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorPropertiesTimesliceMetricParams) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorPropertiesTimesliceMetricParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorPropertiesTimesliceMetricParams(val *IndicatorPropertiesTimesliceMetricParams) *NullableIndicatorPropertiesTimesliceMetricParams {
	return &NullableIndicatorPropertiesTimesliceMetricParams{value: val, isSet: true}
}

func (v NullableIndicatorPropertiesTimesliceMetricParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorPropertiesTimesliceMetricParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
