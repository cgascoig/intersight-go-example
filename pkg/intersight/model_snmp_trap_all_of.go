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

// SnmpTrapAllOf Definition of the list of properties defined in 'snmp.Trap', excluding properties defined in parent classes.
type SnmpTrapAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// SNMP community group used for sending SNMP trap to other devices. Applicable only for SNMP v2c.
	Community *string `json:"Community,omitempty"`
	// Address to which the SNMP trap information is sent.
	Destination *string `json:"Destination,omitempty"`
	// Enables/disables the trap on the server If enabled, trap is active on the server.
	Enabled *bool `json:"Enabled,omitempty"`
	// Port used by the server to communicate with the trap destination. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269).
	Port *int64 `json:"Port,omitempty"`
	// Type of trap which decides whether to receive a notification when a trap is received at the destination. * `Trap` - Do not receive notifications when trap is sent to the destination. * `Inform` - Receive notifications when trap is sent to the destination. This option is valid only for V2 users.
	Type *string `json:"Type,omitempty"`
	// SNMP user for the trap. Applicable only to SNMPv3.
	User *string `json:"User,omitempty"`
	// SNMP version used for the trap. * `V3` - SNMP v3 trap version notifications. * `V2` - SNMP v2 trap version notifications.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpTrapAllOf SnmpTrapAllOf

// NewSnmpTrapAllOf instantiates a new SnmpTrapAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpTrapAllOf(classId string, objectType string) *SnmpTrapAllOf {
	this := SnmpTrapAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var port int64 = 162
	this.Port = &port
	var type_ string = "Trap"
	this.Type = &type_
	var version string = "V3"
	this.Version = &version
	return &this
}

// NewSnmpTrapAllOfWithDefaults instantiates a new SnmpTrapAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpTrapAllOfWithDefaults() *SnmpTrapAllOf {
	this := SnmpTrapAllOf{}
	var classId string = "snmp.Trap"
	this.ClassId = classId
	var objectType string = "snmp.Trap"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var port int64 = 162
	this.Port = &port
	var type_ string = "Trap"
	this.Type = &type_
	var version string = "V3"
	this.Version = &version
	return &this
}

// GetClassId returns the ClassId field value
func (o *SnmpTrapAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SnmpTrapAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SnmpTrapAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SnmpTrapAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SnmpTrapAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SnmpTrapAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommunity returns the Community field value if set, zero value otherwise.
func (o *SnmpTrapAllOf) GetCommunity() string {
	if o == nil || o.Community == nil {
		var ret string
		return ret
	}
	return *o.Community
}

// GetCommunityOk returns a tuple with the Community field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapAllOf) GetCommunityOk() (*string, bool) {
	if o == nil || o.Community == nil {
		return nil, false
	}
	return o.Community, true
}

// HasCommunity returns a boolean if a field has been set.
func (o *SnmpTrapAllOf) HasCommunity() bool {
	if o != nil && o.Community != nil {
		return true
	}

	return false
}

// SetCommunity gets a reference to the given string and assigns it to the Community field.
func (o *SnmpTrapAllOf) SetCommunity(v string) {
	o.Community = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *SnmpTrapAllOf) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapAllOf) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *SnmpTrapAllOf) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *SnmpTrapAllOf) SetDestination(v string) {
	o.Destination = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SnmpTrapAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SnmpTrapAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SnmpTrapAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SnmpTrapAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SnmpTrapAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SnmpTrapAllOf) SetPort(v int64) {
	o.Port = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SnmpTrapAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SnmpTrapAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SnmpTrapAllOf) SetType(v string) {
	o.Type = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SnmpTrapAllOf) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapAllOf) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SnmpTrapAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *SnmpTrapAllOf) SetUser(v string) {
	o.User = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SnmpTrapAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SnmpTrapAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SnmpTrapAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o SnmpTrapAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Community != nil {
		toSerialize["Community"] = o.Community
	}
	if o.Destination != nil {
		toSerialize["Destination"] = o.Destination
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SnmpTrapAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSnmpTrapAllOf := _SnmpTrapAllOf{}

	if err = json.Unmarshal(bytes, &varSnmpTrapAllOf); err == nil {
		*o = SnmpTrapAllOf(varSnmpTrapAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Community")
		delete(additionalProperties, "Destination")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "User")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpTrapAllOf struct {
	value *SnmpTrapAllOf
	isSet bool
}

func (v NullableSnmpTrapAllOf) Get() *SnmpTrapAllOf {
	return v.value
}

func (v *NullableSnmpTrapAllOf) Set(val *SnmpTrapAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpTrapAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpTrapAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpTrapAllOf(val *SnmpTrapAllOf) *NullableSnmpTrapAllOf {
	return &NullableSnmpTrapAllOf{value: val, isSet: true}
}

func (v NullableSnmpTrapAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpTrapAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
