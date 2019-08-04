/*
Copyright © 2019 wdstar

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

package client

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HMACAuth provides a HMAC auth info writer
func HMACAuth(keyID, keySecret string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		// TODO: Enabling HMAC Authentication
		// https://help.veracode.com/reader/LMv_dtSHyb7iIxAQznC~9w/hn2qc_7fz3zFYV~e4ulRaQ
		hmacStr := "********"
		return r.SetHeaderParam("Authorization", hmacStr)
	})
}
