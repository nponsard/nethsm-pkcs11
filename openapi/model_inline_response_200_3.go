/*
 * NitroHSM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	RealName string `json:"realName"`
	Role UserRole `json:"role"`
}

// NewInlineResponse2003 instantiates a new InlineResponse2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003(realName string, role UserRole, ) *InlineResponse2003 {
	this := InlineResponse2003{}
	this.RealName = realName
	this.Role = role
	return &this
}

// NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003WithDefaults() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// GetRealName returns the RealName field value
func (o *InlineResponse2003) GetRealName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetRealNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RealName, true
}

// SetRealName sets field value
func (o *InlineResponse2003) SetRealName(v string) {
	o.RealName = v
}

// GetRole returns the Role field value
func (o *InlineResponse2003) GetRole() UserRole {
	if o == nil  {
		var ret UserRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetRoleOk() (*UserRole, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InlineResponse2003) SetRole(v UserRole) {
	o.Role = v
}

func (o InlineResponse2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["realName"] = o.RealName
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003 struct {
	value *InlineResponse2003
	isSet bool
}

func (v NullableInlineResponse2003) Get() *InlineResponse2003 {
	return v.value
}

func (v *NullableInlineResponse2003) Set(val *InlineResponse2003) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003(val *InlineResponse2003) *NullableInlineResponse2003 {
	return &NullableInlineResponse2003{value: val, isSet: true}
}

func (v NullableInlineResponse2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


