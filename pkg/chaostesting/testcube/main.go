// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/matrixorigin/matrixone/pkg/chaostesting"
)

func run() (err error) {
	defer he(&err)

	scope := NewScope()

	var execute fz.Execute
	scope.Assign(&execute)
	ce(execute())

	return
}

func main() {
	run()
}
