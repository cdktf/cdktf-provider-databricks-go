// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineLibrary struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/pipeline#file Pipeline#file}
	File *PipelineLibraryFile `field:"optional" json:"file" yaml:"file"`
	// glob block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/pipeline#glob Pipeline#glob}
	Glob *PipelineLibraryGlob `field:"optional" json:"glob" yaml:"glob"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/pipeline#jar Pipeline#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/pipeline#maven Pipeline#maven}
	Maven *PipelineLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// notebook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/pipeline#notebook Pipeline#notebook}
	Notebook *PipelineLibraryNotebook `field:"optional" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/pipeline#whl Pipeline#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

