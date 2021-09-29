/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-21T16:37:30Z.
 *
 * API version: 1.0.9-4403
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// MoAggregateTransformAllOf struct for MoAggregateTransformAllOf
type MoAggregateTransformAllOf struct {
	// The results of the aggregation query.
	Results              []map[string]interface{} `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MoAggregateTransformAllOf MoAggregateTransformAllOf

// NewMoAggregateTransformAllOf instantiates a new MoAggregateTransformAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoAggregateTransformAllOf() *MoAggregateTransformAllOf {
	this := MoAggregateTransformAllOf{}
	return &this
}

// NewMoAggregateTransformAllOfWithDefaults instantiates a new MoAggregateTransformAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoAggregateTransformAllOfWithDefaults() *MoAggregateTransformAllOf {
	this := MoAggregateTransformAllOf{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoAggregateTransformAllOf) GetResults() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoAggregateTransformAllOf) GetResultsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MoAggregateTransformAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []map[string]interface{} and assigns it to the Results field.
func (o *MoAggregateTransformAllOf) SetResults(v []map[string]interface{}) {
	o.Results = v
}

func (o MoAggregateTransformAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MoAggregateTransformAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMoAggregateTransformAllOf := _MoAggregateTransformAllOf{}

	if err = json.Unmarshal(bytes, &varMoAggregateTransformAllOf); err == nil {
		*o = MoAggregateTransformAllOf(varMoAggregateTransformAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMoAggregateTransformAllOf struct {
	value *MoAggregateTransformAllOf
	isSet bool
}

func (v NullableMoAggregateTransformAllOf) Get() *MoAggregateTransformAllOf {
	return v.value
}

func (v *NullableMoAggregateTransformAllOf) Set(val *MoAggregateTransformAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMoAggregateTransformAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMoAggregateTransformAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoAggregateTransformAllOf(val *MoAggregateTransformAllOf) *NullableMoAggregateTransformAllOf {
	return &NullableMoAggregateTransformAllOf{value: val, isSet: true}
}

func (v NullableMoAggregateTransformAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoAggregateTransformAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
