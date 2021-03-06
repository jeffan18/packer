package common

import (
	"testing"

	"github.com/hashicorp/packer/template/interpolate"
)

func TestDriverConfigPrepare(t *testing.T) {
	var c *DriverConfig

	// Test a default boot_wait
	c = new(DriverConfig)
	errs := c.Prepare(interpolate.NewContext())
	if len(errs) > 0 {
		t.Fatalf("bad: %#v", errs)
	}
	if c.FusionAppPath != "/Applications/VMware Fusion.app" {
		t.Fatalf("bad value: %s", c.FusionAppPath)
	}

	// Test with a good one
	c = new(DriverConfig)
	c.FusionAppPath = "foo"
	errs = c.Prepare(interpolate.NewContext())
	if len(errs) > 0 {
		t.Fatalf("bad: %#v", errs)
	}
	if c.FusionAppPath != "foo" {
		t.Fatalf("bad value: %s", c.FusionAppPath)
	}
}
