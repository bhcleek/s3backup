// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiReadCertDeltaCrlResponse struct for PkiReadCertDeltaCrlResponse
type PkiReadCertDeltaCrlResponse struct {
	// Issuing CA Chain
	CaChain []string `json:"ca_chain,omitempty"`

	// Certificate
	Certificate string `json:"certificate,omitempty"`

	// ID of the issuer
	IssuerId string `json:"issuer_id,omitempty"`

	// Revocation time
	RevocationTime string `json:"revocation_time,omitempty"`

	// Revocation time RFC 3339 formatted
	RevocationTimeRfc3339 string `json:"revocation_time_rfc3339,omitempty"`
}

// NewPkiReadCertDeltaCrlResponseWithDefaults instantiates a new PkiReadCertDeltaCrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiReadCertDeltaCrlResponseWithDefaults() *PkiReadCertDeltaCrlResponse {
	var this PkiReadCertDeltaCrlResponse

	return &this
}
