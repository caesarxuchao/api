/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package unversioned

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TODO: We need to remove the GroupVersion in types.go. We use the name GroupAndVersion here temporarily.
type GroupAndVersion struct {
	Group   string
	Version string
}

func (gv *GroupAndVersion) String() string {
	// special case of "v1" for backward compatibility
	if gv.Group == "" {
		return gv.Version
	} else {
		return gv.Group + "/" + gv.Version
	}
}

func ParseGroupVersion(gv string) (GroupAndVersion, error) {
	s := strings.Split(gv, "/")
	// "v1" is the only special case. Otherwise GroupVersion is expected to contain
	// one "/" dividing the string into two parts.
	if len(s) == 1 && gv == "v1" {
		return GroupAndVersion{"", "v1"}, nil
	} else if len(s) == 2 {
		return GroupAndVersion{s[0], s[1]}, nil
	} else {
		return GroupAndVersion{}, fmt.Errorf("Unexpected GroupVersion string: %v", gv)
	}
}

// MarshalJSON implements the json.Marshaller interface.
func (gv GroupAndVersion) MarshalJSON() ([]byte, error) {
	s := gv.String()
	if strings.Count(s, "/") > 1 {
		return []byte{}, fmt.Errorf("illegal GroupVersion %v: contains more than one /", s)
	}
	return json.Marshal(s)
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (gv *GroupAndVersion) UnmarshalJSON(value []byte) error {
	var s string
	if err := json.Unmarshal(value, &s); err != nil {
		return err
	}
	parsed, err := ParseGroupVersion(s)
	if err != nil {
		return err
	}
	gv.Group = parsed.Group
	gv.Version = parsed.Version
	return nil
}
