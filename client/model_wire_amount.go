/*
 * WIRE API
 *
 * Moov WIRE implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// WireAmount Amount is an amount up to a penny less than $10 billion
type WireAmount struct {
	// Amount 12 numeric, right-justified with leading zeros, an implied decimal point and no commas; e.g., $12,345.67 becomes 000001234567 Can be all zeros for subtype 90
	Amount string `json:"amount"`
}
