package stubs

// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"github.com/googlecloudplatform/threat-automation/clients"
)

// EmailStub provides a stub for the Email client.
type EmailStub struct {
	StubbedSend    *clients.EmailResponse
	StubbedSendErr error
}

// Send to send email
func (e *EmailStub) Send(subject, from, body string, to []string) (*clients.EmailResponse, error) {
	return e.StubbedSend, e.StubbedSendErr
}
