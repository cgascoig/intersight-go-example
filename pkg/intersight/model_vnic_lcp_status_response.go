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
	"fmt"
)

// VnicLcpStatusResponse - The response body of a HTTP GET request for the 'vnic.LcpStatus' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'vnic.LcpStatus' resources.
type VnicLcpStatusResponse struct {
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
	VnicLcpStatusList    *VnicLcpStatusList
}

// MoAggregateTransformAsVnicLcpStatusResponse is a convenience function that returns MoAggregateTransform wrapped in VnicLcpStatusResponse
func MoAggregateTransformAsVnicLcpStatusResponse(v *MoAggregateTransform) VnicLcpStatusResponse {
	return VnicLcpStatusResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsVnicLcpStatusResponse is a convenience function that returns MoDocumentCount wrapped in VnicLcpStatusResponse
func MoDocumentCountAsVnicLcpStatusResponse(v *MoDocumentCount) VnicLcpStatusResponse {
	return VnicLcpStatusResponse{MoDocumentCount: v}
}

// MoTagSummaryAsVnicLcpStatusResponse is a convenience function that returns MoTagSummary wrapped in VnicLcpStatusResponse
func MoTagSummaryAsVnicLcpStatusResponse(v *MoTagSummary) VnicLcpStatusResponse {
	return VnicLcpStatusResponse{MoTagSummary: v}
}

// VnicLcpStatusListAsVnicLcpStatusResponse is a convenience function that returns VnicLcpStatusList wrapped in VnicLcpStatusResponse
func VnicLcpStatusListAsVnicLcpStatusResponse(v *VnicLcpStatusList) VnicLcpStatusResponse {
	return VnicLcpStatusResponse{VnicLcpStatusList: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VnicLcpStatusResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal VnicLcpStatusResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal VnicLcpStatusResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("Failed to unmarshal VnicLcpStatusResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vnic.LcpStatus.List'
	if jsonDict["ObjectType"] == "vnic.LcpStatus.List" {
		// try to unmarshal JSON data into VnicLcpStatusList
		err = json.Unmarshal(data, &dst.VnicLcpStatusList)
		if err == nil {
			return nil // data stored in dst.VnicLcpStatusList, return on the first match
		} else {
			dst.VnicLcpStatusList = nil
			return fmt.Errorf("Failed to unmarshal VnicLcpStatusResponse as VnicLcpStatusList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VnicLcpStatusResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.VnicLcpStatusList != nil {
		return json.Marshal(&src.VnicLcpStatusList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VnicLcpStatusResponse) GetActualInstance() interface{} {
	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	if obj.VnicLcpStatusList != nil {
		return obj.VnicLcpStatusList
	}

	// all schemas are nil
	return nil
}

type NullableVnicLcpStatusResponse struct {
	value *VnicLcpStatusResponse
	isSet bool
}

func (v NullableVnicLcpStatusResponse) Get() *VnicLcpStatusResponse {
	return v.value
}

func (v *NullableVnicLcpStatusResponse) Set(val *VnicLcpStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicLcpStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicLcpStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicLcpStatusResponse(val *VnicLcpStatusResponse) *NullableVnicLcpStatusResponse {
	return &NullableVnicLcpStatusResponse{value: val, isSet: true}
}

func (v NullableVnicLcpStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicLcpStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
