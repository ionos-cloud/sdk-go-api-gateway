/*
 * IONOS Cloud - API Gateway
 *
 * API Gateway is an application that acts as a \"front door\" for backend services and APIs, handling client requests and routing them to the appropriate backend.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// GatewayEnsure struct for GatewayEnsure
type GatewayEnsure struct {
	// The ID (UUID) of the Gateway.
	Id *string `json:"id"`
	// Metadata
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	Properties *Gateway                `json:"properties"`
}

// NewGatewayEnsure instantiates a new GatewayEnsure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayEnsure(id string, properties Gateway) *GatewayEnsure {
	this := GatewayEnsure{}

	this.Id = &id
	this.Properties = &properties

	return &this
}

// NewGatewayEnsureWithDefaults instantiates a new GatewayEnsure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayEnsureWithDefaults() *GatewayEnsure {
	this := GatewayEnsure{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GatewayEnsure) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewayEnsure) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *GatewayEnsure) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *GatewayEnsure) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *GatewayEnsure) GetMetadata() *map[string]interface{} {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewayEnsure) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *GatewayEnsure) SetMetadata(v map[string]interface{}) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *GatewayEnsure) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for Gateway will be returned
func (o *GatewayEnsure) GetProperties() *Gateway {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewayEnsure) GetPropertiesOk() (*Gateway, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *GatewayEnsure) SetProperties(v Gateway) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *GatewayEnsure) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o GatewayEnsure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableGatewayEnsure struct {
	value *GatewayEnsure
	isSet bool
}

func (v NullableGatewayEnsure) Get() *GatewayEnsure {
	return v.value
}

func (v *NullableGatewayEnsure) Set(val *GatewayEnsure) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayEnsure) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayEnsure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayEnsure(val *GatewayEnsure) *NullableGatewayEnsure {
	return &NullableGatewayEnsure{value: val, isSet: true}
}

func (v NullableGatewayEnsure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayEnsure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
