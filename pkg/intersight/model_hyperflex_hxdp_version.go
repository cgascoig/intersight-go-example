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

// HyperflexHxdpVersion A HyperFlex Data Platform version.
type HyperflexHxdpVersion struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The HyperFlex Data Platform version.
	Version              *string                          `json:"Version,omitempty"`
	AppCatalog           *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxdpVersion HyperflexHxdpVersion

// NewHyperflexHxdpVersion instantiates a new HyperflexHxdpVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxdpVersion(classId string, objectType string) *HyperflexHxdpVersion {
	this := HyperflexHxdpVersion{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxdpVersionWithDefaults instantiates a new HyperflexHxdpVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxdpVersionWithDefaults() *HyperflexHxdpVersion {
	this := HyperflexHxdpVersion{}
	var classId string = "hyperflex.HxdpVersion"
	this.ClassId = classId
	var objectType string = "hyperflex.HxdpVersion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxdpVersion) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxdpVersion) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxdpVersion) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxdpVersion) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxdpVersion) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxdpVersion) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexHxdpVersion) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxdpVersion) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexHxdpVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexHxdpVersion) SetVersion(v string) {
	o.Version = &v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HyperflexHxdpVersion) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxdpVersion) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexHxdpVersion) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexHxdpVersion) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

func (o HyperflexHxdpVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxdpVersion) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxdpVersionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The HyperFlex Data Platform version.
		Version    *string                          `json:"Version,omitempty"`
		AppCatalog *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	}

	varHyperflexHxdpVersionWithoutEmbeddedStruct := HyperflexHxdpVersionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxdpVersionWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxdpVersion := _HyperflexHxdpVersion{}
		varHyperflexHxdpVersion.ClassId = varHyperflexHxdpVersionWithoutEmbeddedStruct.ClassId
		varHyperflexHxdpVersion.ObjectType = varHyperflexHxdpVersionWithoutEmbeddedStruct.ObjectType
		varHyperflexHxdpVersion.Version = varHyperflexHxdpVersionWithoutEmbeddedStruct.Version
		varHyperflexHxdpVersion.AppCatalog = varHyperflexHxdpVersionWithoutEmbeddedStruct.AppCatalog
		*o = HyperflexHxdpVersion(varHyperflexHxdpVersion)
	} else {
		return err
	}

	varHyperflexHxdpVersion := _HyperflexHxdpVersion{}

	err = json.Unmarshal(bytes, &varHyperflexHxdpVersion)
	if err == nil {
		o.MoBaseMo = varHyperflexHxdpVersion.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "AppCatalog")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHyperflexHxdpVersion struct {
	value *HyperflexHxdpVersion
	isSet bool
}

func (v NullableHyperflexHxdpVersion) Get() *HyperflexHxdpVersion {
	return v.value
}

func (v *NullableHyperflexHxdpVersion) Set(val *HyperflexHxdpVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxdpVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxdpVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxdpVersion(val *HyperflexHxdpVersion) *NullableHyperflexHxdpVersion {
	return &NullableHyperflexHxdpVersion{value: val, isSet: true}
}

func (v NullableHyperflexHxdpVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxdpVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
