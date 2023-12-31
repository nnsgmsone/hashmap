// Copyright 2022 Matrix Origin
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

package hashmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashmap(t *testing.T) {
	m := New[int, int](0)
	for i := 0; i < 1000; i++ {
		m.Set(i, i)
	}
	// test Len
	require.Equal(t, 1000, m.Len())
	// test Get
	for i := 0; i < 1000; i++ {
		v, ok := m.Get(i)
		require.Equal(t, true, ok)
		require.Equal(t, i, v)
	}
	// test Delete
	m.Del(0)
	_, ok := m.Get(0)
	require.Equal(t, false, ok)
}
