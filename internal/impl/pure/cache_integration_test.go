// Copyright 2025 Redpanda Data, Inc.

package pure

import (
	"testing"

	"github.com/redpanda-data/benthos/v4/public/service/integration"
)

func TestIntegrationMultilevelCache(t *testing.T) {
	integration.CheckSkip(t)

	t.Parallel()

	template := `
cache_resources:
  - label: testcache
    multilevel: [ first, second ]
  - label: first
    memory: {}
  - label: second
    memory: {}
`
	suite := integration.CacheTests(
		integration.CacheTestOpenClose(),
		integration.CacheTestMissingKey(),
		integration.CacheTestDoubleAdd(),
		integration.CacheTestDelete(),
		integration.CacheTestGetAndSet(50),
	)
	suite.Run(t, template)
}
