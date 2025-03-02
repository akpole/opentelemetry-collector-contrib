// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/translator"

import "context"

// HostnameProvider gets a hostname
type HostnameProvider interface {
	// Hostname gets the hostname from the machine.
	Hostname(ctx context.Context) (string, error)
}

var _ HostnameProvider = (*noHostProvider)(nil)

type noHostProvider struct{}

func (*noHostProvider) Hostname(context.Context) (string, error) {
	return "", nil
}
