// Copyright 2019 Tetrate
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

package binary

import "github.com/tetratelabs/getenvoy/pkg/manifest"

// Runner wraps the Envoy Run interface
type Runner interface {
	Run(key *manifest.Key, args []string) error
	RunPath(path string, args []string) error
}

// Fetcher wraps the Envoy Fetch interface
type Fetcher interface {
	Fetch(key *manifest.Key, binaryLocation string) error
}

// Runtime wraps the Run and Fetch interfaces
type Runtime interface {
	Runner
	Fetcher
}
