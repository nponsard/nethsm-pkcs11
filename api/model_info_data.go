/*
 * NetHSM
 *
 * All endpoints expect exactly the specified JSON. Additional properties will cause a Bad Request Error (400). All HTTP errors contain a JSON structure with an explanation of type string. All <a href=\"https://tools.ietf.org/html/rfc4648#section-4\">base64</a> encoded values are Big Endian.
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InfoData struct for InfoData
type InfoData struct {
	Vendor string `json:"vendor"`
	Product string `json:"product"`
	AdditionalProperties map[string]interface{}
}

type _InfoData InfoData

// NewInfoData instantiates a new InfoData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfoData(vendor string, product string, ) *InfoData {
	this := InfoData{}
	this.Vendor = vendor
	this.Product = product
	return &this
}

// NewInfoDataWithDefaults instantiates a new InfoData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoDataWithDefaults() *InfoData {
	this := InfoData{}
	return &this
}

// GetVendor returns the Vendor field value
func (o *InfoData) GetVendor() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *InfoData) GetVendorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *InfoData) SetVendor(v string) {
	o.Vendor = v
}

// GetProduct returns the Product field value
func (o *InfoData) GetProduct() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *InfoData) GetProductOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *InfoData) SetProduct(v string) {
	o.Product = v
}

func (o InfoData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vendor"] = o.Vendor
	}
	if true {
		toSerialize["product"] = o.Product
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InfoData) UnmarshalJSON(bytes []byte) (err error) {
	varInfoData := _InfoData{}

	if err = json.Unmarshal(bytes, &varInfoData); err == nil {
		*o = InfoData(varInfoData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "product")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInfoData struct {
	value *InfoData
	isSet bool
}

func (v NullableInfoData) Get() *InfoData {
	return v.value
}

func (v *NullableInfoData) Set(val *InfoData) {
	v.value = val
	v.isSet = true
}

func (v NullableInfoData) IsSet() bool {
	return v.isSet
}

func (v *NullableInfoData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfoData(val *InfoData) *NullableInfoData {
	return &NullableInfoData{value: val, isSet: true}
}

func (v NullableInfoData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfoData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

