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

package splice

import (
	"bytes"

	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func String(arg interface{}, buf *bytes.Buffer) {
	buf.WriteString("splice")
}

func Prepare(_ *process.Process, _ interface{}) error {
	return nil
}

func Call(proc *process.Process, arg interface{}) (bool, error) {
	var err error

	n := arg.(*Argument)
	if len(proc.Reg.MergeReceivers) == 0 {
		proc.Reg.InputBatch = n.bat
		n.bat = nil
		return true, nil
	}
	for i := 0; i < len(proc.Reg.MergeReceivers); i++ {
		reg := proc.Reg.MergeReceivers[i]
		bat := <-reg.Ch
		if bat == nil {
			proc.Reg.MergeReceivers = append(proc.Reg.MergeReceivers[:i], proc.Reg.MergeReceivers[i+1:]...)
			i--
			continue
		}
		if len(bat.Zs) == 0 {
			i--
			continue
		}
		if n.bat == nil {
			n.bat = bat
		} else {
			n.bat, err = n.bat.Append(proc.Mp, bat)
			if err != nil {
				return false, err
			}
		}
		i--
	}
	proc.Reg.InputBatch = n.bat
	n.bat = nil
	return true, nil
}
