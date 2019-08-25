package config

import (
	"testing"
)

func TestEmptyDeciderTypeFails(t *testing.T) {
	conf := &Decider{
		Type: "",
	}

	_, err := ToDecider(conf)
	if err == nil {
		t.Error("expected decider config with no type to fail")
	}
}

func TestNoMappingForDeciderFails(t *testing.T) {
	conf := &Decider{
		Type: "doesnotexist",
	}

	_, err := ToDecider(conf)
	if err == nil {
		t.Errorf("expected decider config with type %s to fail for no mapping", conf.Type)
	}
}

func TestDeciderKeepDuration(t *testing.T) {
	tests := []struct {
		Config     *Decider
		ShouldFail bool
	}{
		{
			Config: &Decider{
				Type: "keepAfterDuration",
				Options: map[string]interface{}{
					"duration": "1s",
				},
			},
			ShouldFail: false,
		},
		{
			Config: &Decider{
				Type:    "keepAfterDuration",
				Options: map[string]interface{}{},
			},
			ShouldFail: true,
		},
		{
			Config: &Decider{
				Type: "keepAfterDuration",
				Options: map[string]interface{}{
					"duration": "-7h",
				},
			},
			ShouldFail: true,
		},
		{
			Config: &Decider{
				Type: "keepAfterDuration",
				Options: map[string]interface{}{
					"duration": "2880h",
				},
			},
			ShouldFail: false,
		},
	}

	for i, test := range tests {
		_, err := ToDecider(test.Config)
		if err == nil && test.ShouldFail {
			t.Fatalf("expected test %d to fail with an error", i)
		} else if err != nil && !test.ShouldFail {
			t.Fatalf("unexpected error for test %d: %v", i, err)
		}
	}
}
