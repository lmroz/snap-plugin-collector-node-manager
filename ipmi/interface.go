// +build linux

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ipmi

// Abstract type for ipmi backend.
type IpmiAL interface {
	BatchExecRaw(requests []IpmiRequest, nSim int) ([]IpmiResponse, error)
}

// Defines request parameter passed to abstraction layer.
type IpmiRequest struct {
	Data    []byte
	Channel int16
	Slave   uint8
}

// Defines response data.
type IpmiResponse struct {
	Data []byte
	IsValid uint
}

// Vendor exposed structure. Defines request content and response format.
// List of submetrics for given format should be concatenated with MetricsRoot
// to specify full metric name.
type RequestDescription struct {
	Request     IpmiRequest
	MetricsRoot string
	Format      ParserFormat
}

// Defines interface that all response formats must implement.
// GetMetrics() should return all available submetrics for given format.
// Main metric value should have label "" (empty string).
// Validate() should check response correctness. Nil is returned when response
// is correct.
// Parse() extracts submetrics from binary data.
type ParserFormat interface {
	GetMetrics() []string
	Validate(response IpmiResponse) error
	Parse(response IpmiResponse) map[string]uint16
}
