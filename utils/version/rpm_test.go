// Copyright (c) 2022 Institute of Software, Chinese Academy of Sciences (ISCAS)
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

package version

import (
	"strings"
	"testing"
)

func TestGetInstalledAgentVersion(t *testing.T) {
	version, err := GetInstalledAgentVersion()
	if err != nil {
		t.Fatal(err)
	}
	if version == "" {
		t.Fatal("Can't get agent version'")
	}
	if strings.Contains(version, "\n") {
		t.Errorf("the agent version contains '\\n'")
	}
	t.Logf("version is: %v", version)
}
