// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build (nodejs || all) && !smoke

package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestNodejsTransformations(t *testing.T) {
	t.Parallel()

	for _, dir := range Dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:                    d,
				Dependencies:           []string{"@pulumi/pulumi"},
				Quick:                  true,
				ExtraRuntimeValidation: Validator("nodejs"),
			})
		})
	}
}
