// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskLibrary struct {
	// cran block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#cran Job#cran}
	Cran *JobTaskLibraryCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#egg Job#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#jar Job#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#maven Job#maven}
	Maven *JobTaskLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// pypi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#pypi Job#pypi}
	Pypi *JobTaskLibraryPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#whl Job#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

