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
	"reflect"
	"strings"
)

// PolicyConfigContext Configuration related state and results.
type PolicyConfigContext struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Out-of-sync, Validating, Configuring, Failed.
	ConfigState *string `json:"ConfigState,omitempty"`
	// System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind.
	ControlAction *string `json:"ControlAction,omitempty"`
	// Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error).
	ErrorState *string `json:"ErrorState,omitempty"`
	// Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed.
	OperState            *string `json:"OperState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyConfigContext PolicyConfigContext

// NewPolicyConfigContext instantiates a new PolicyConfigContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyConfigContext(classId string, objectType string) *PolicyConfigContext {
	this := PolicyConfigContext{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyConfigContextWithDefaults instantiates a new PolicyConfigContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyConfigContextWithDefaults() *PolicyConfigContext {
	this := PolicyConfigContext{}
	var classId string = "policy.ConfigContext"
	this.ClassId = classId
	var objectType string = "policy.ConfigContext"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyConfigContext) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigContext) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyConfigContext) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyConfigContext) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigContext) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyConfigContext) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *PolicyConfigContext) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigContext) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *PolicyConfigContext) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *PolicyConfigContext) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetControlAction returns the ControlAction field value if set, zero value otherwise.
func (o *PolicyConfigContext) GetControlAction() string {
	if o == nil || o.ControlAction == nil {
		var ret string
		return ret
	}
	return *o.ControlAction
}

// GetControlActionOk returns a tuple with the ControlAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigContext) GetControlActionOk() (*string, bool) {
	if o == nil || o.ControlAction == nil {
		return nil, false
	}
	return o.ControlAction, true
}

// HasControlAction returns a boolean if a field has been set.
func (o *PolicyConfigContext) HasControlAction() bool {
	if o != nil && o.ControlAction != nil {
		return true
	}

	return false
}

// SetControlAction gets a reference to the given string and assigns it to the ControlAction field.
func (o *PolicyConfigContext) SetControlAction(v string) {
	o.ControlAction = &v
}

// GetErrorState returns the ErrorState field value if set, zero value otherwise.
func (o *PolicyConfigContext) GetErrorState() string {
	if o == nil || o.ErrorState == nil {
		var ret string
		return ret
	}
	return *o.ErrorState
}

// GetErrorStateOk returns a tuple with the ErrorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigContext) GetErrorStateOk() (*string, bool) {
	if o == nil || o.ErrorState == nil {
		return nil, false
	}
	return o.ErrorState, true
}

// HasErrorState returns a boolean if a field has been set.
func (o *PolicyConfigContext) HasErrorState() bool {
	if o != nil && o.ErrorState != nil {
		return true
	}

	return false
}

// SetErrorState gets a reference to the given string and assigns it to the ErrorState field.
func (o *PolicyConfigContext) SetErrorState(v string) {
	o.ErrorState = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *PolicyConfigContext) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigContext) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *PolicyConfigContext) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *PolicyConfigContext) SetOperState(v string) {
	o.OperState = &v
}

func (o PolicyConfigContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.ControlAction != nil {
		toSerialize["ControlAction"] = o.ControlAction
	}
	if o.ErrorState != nil {
		toSerialize["ErrorState"] = o.ErrorState
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyConfigContext) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyConfigContextWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Out-of-sync, Validating, Configuring, Failed.
		ConfigState *string `json:"ConfigState,omitempty"`
		// System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind.
		ControlAction *string `json:"ControlAction,omitempty"`
		// Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error).
		ErrorState *string `json:"ErrorState,omitempty"`
		// Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed.
		OperState *string `json:"OperState,omitempty"`
	}

	varPolicyConfigContextWithoutEmbeddedStruct := PolicyConfigContextWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyConfigContextWithoutEmbeddedStruct)
	if err == nil {
		varPolicyConfigContext := _PolicyConfigContext{}
		varPolicyConfigContext.ClassId = varPolicyConfigContextWithoutEmbeddedStruct.ClassId
		varPolicyConfigContext.ObjectType = varPolicyConfigContextWithoutEmbeddedStruct.ObjectType
		varPolicyConfigContext.ConfigState = varPolicyConfigContextWithoutEmbeddedStruct.ConfigState
		varPolicyConfigContext.ControlAction = varPolicyConfigContextWithoutEmbeddedStruct.ControlAction
		varPolicyConfigContext.ErrorState = varPolicyConfigContextWithoutEmbeddedStruct.ErrorState
		varPolicyConfigContext.OperState = varPolicyConfigContextWithoutEmbeddedStruct.OperState
		*o = PolicyConfigContext(varPolicyConfigContext)
	} else {
		return err
	}

	varPolicyConfigContext := _PolicyConfigContext{}

	err = json.Unmarshal(bytes, &varPolicyConfigContext)
	if err == nil {
		o.MoBaseComplexType = varPolicyConfigContext.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "ControlAction")
		delete(additionalProperties, "ErrorState")
		delete(additionalProperties, "OperState")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyConfigContext struct {
	value *PolicyConfigContext
	isSet bool
}

func (v NullablePolicyConfigContext) Get() *PolicyConfigContext {
	return v.value
}

func (v *NullablePolicyConfigContext) Set(val *PolicyConfigContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyConfigContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyConfigContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyConfigContext(val *PolicyConfigContext) *NullablePolicyConfigContext {
	return &NullablePolicyConfigContext{value: val, isSet: true}
}

func (v NullablePolicyConfigContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyConfigContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
