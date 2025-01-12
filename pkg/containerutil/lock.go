/*
   Copyright The containerd Authors.

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

package containerutil

import (
	"path/filepath"

	"github.com/containerd/nerdctl/v2/pkg/store"
)

func Lock(stateDir string) (store.Store, error) {
	stor, err := store.New(filepath.Join(stateDir, "oplock"), 0, 0)
	if err != nil {
		return nil, err
	}

	err = stor.Lock()
	if err != nil {
		return nil, err
	}

	return stor, nil
}
