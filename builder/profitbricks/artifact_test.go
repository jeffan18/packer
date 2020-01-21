package profitbricks

import (
	"testing"

	"github.com/hashicorp/packer/packer"
)

func TestArtifact_Impl(t *testing.T) {
	var raw interface{}
	raw = &Artifact{}
	if _, ok := raw.(packer.Artifact); !ok {
		t.Fatalf("Artifact should be artifact")
	}
}

func TestArtifactString(t *testing.T) {
	generatedData := make(map[string]interface{})
	a := &Artifact{"packer-foobar", generatedData}
	expected := "A snapshot was created: 'packer-foobar'"

	if a.String() != expected {
		t.Fatalf("artifact string should match: %v", expected)
	}
}

func TestArtifactState_StateData(t *testing.T) {
	expectedData := "this is the data"
	artifact := &Artifact{
		StateData: map[string]interface{}{"state_data": expectedData},
	}

	// Valid state
	result := artifact.State("state_data")
	if result != expectedData {
		t.Fatalf("Bad: State data was %s instead of %s", result, expectedData)
	}

	// Invalid state
	result = artifact.State("invalid_key")
	if result != nil {
		t.Fatalf("Bad: State should be nil for invalid state data name")
	}
}
