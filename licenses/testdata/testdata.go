// Copyright 2019 Google Inc. All Rights Reserved.
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

package testdata

import (
	// This import should be detected by library_test.go. It has to be a
	// package that isn't in the standard library and has a separate license
	// file to the one covering the Trillian repository, so that it's detected
	// as being an external dependency.
	_ "github.com/Shubhranshu153/go-licenses/v2/licenses/testdata/direct"

	// This import should be ignored, since it's an internal dependency.
	_ "github.com/Shubhranshu153/go-licenses/v2/licenses/testdata/internal"

	// This import should be ignored, since it's an standard library package.
	_ "strings"
)
