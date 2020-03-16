package external_mysql

import (
	"github.com/kubeflow/manifests/tests"
	"testing"
)

func TestKustomize(t *testing.T) {
	testCase := &tests.KustomizeTestCase{
		Package: "../../../../../pipeline/api-service/overlays/external-mysql",
		Expected: "test_data/expected",
	}

	tests.RunTestCase(t, testCase)
}