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

// PrivateKey struct for PrivateKey
type PrivateKey struct {
	Purpose KeyPurpose `json:"purpose"`
	Algorithm KeyAlgorithm `json:"algorithm"`
	Key PrivateKeyKey `json:"key"`
}

// NewPrivateKey instantiates a new PrivateKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateKey(purpose KeyPurpose, algorithm KeyAlgorithm, key PrivateKeyKey, ) *PrivateKey {
	this := PrivateKey{}
	this.Purpose = purpose
	this.Algorithm = algorithm
	this.Key = key
	return &this
}

// NewPrivateKeyWithDefaults instantiates a new PrivateKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateKeyWithDefaults() *PrivateKey {
	this := PrivateKey{}
	return &this
}

// GetPurpose returns the Purpose field value
func (o *PrivateKey) GetPurpose() KeyPurpose {
	if o == nil  {
		var ret KeyPurpose
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *PrivateKey) GetPurposeOk() (*KeyPurpose, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *PrivateKey) SetPurpose(v KeyPurpose) {
	o.Purpose = v
}

// GetAlgorithm returns the Algorithm field value
func (o *PrivateKey) GetAlgorithm() KeyAlgorithm {
	if o == nil  {
		var ret KeyAlgorithm
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *PrivateKey) GetAlgorithmOk() (*KeyAlgorithm, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *PrivateKey) SetAlgorithm(v KeyAlgorithm) {
	o.Algorithm = v
}

// GetKey returns the Key field value
func (o *PrivateKey) GetKey() PrivateKeyKey {
	if o == nil  {
		var ret PrivateKeyKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PrivateKey) GetKeyOk() (*PrivateKeyKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PrivateKey) SetKey(v PrivateKeyKey) {
	o.Key = v
}

func (o PrivateKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["purpose"] = o.Purpose
	}
	if true {
		toSerialize["algorithm"] = o.Algorithm
	}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateKey struct {
	value *PrivateKey
	isSet bool
}

func (v NullablePrivateKey) Get() *PrivateKey {
	return v.value
}

func (v *NullablePrivateKey) Set(val *PrivateKey) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateKey) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateKey(val *PrivateKey) *NullablePrivateKey {
	return &NullablePrivateKey{value: val, isSet: true}
}

func (v NullablePrivateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


