// Copyright 2014 - anova r&d bvba. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package noko

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagelinks(t *testing.T) {
	const header = `<https://apitest.nokotime.com/api/v2/users/?page=3&per_page=100>; rel="next",
  <https://apitest.nokotime.com/api/v2/users/?page=2&per_page=100>; rel="prev",
  <https://apitest.nokotime.com/api/v2/users/?page=1&per_page=100>; rel="first",
  <https://apitest.nokotime.com/api/v2/users/?page=50&per_page=100>; rel="last"`

	var links map[string]string = pagelinks(header)
	assert.Equal(t, "https://apitest.nokotime.com/api/v2/users/?page=3&per_page=100", links["next"])
	assert.Equal(t, "https://apitest.nokotime.com/api/v2/users/?page=2&per_page=100", links["prev"])
	assert.Equal(t, "https://apitest.nokotime.com/api/v2/users/?page=1&per_page=100", links["first"])
	assert.Equal(t, "https://apitest.nokotime.com/api/v2/users/?page=50&per_page=100", links["last"])
}
