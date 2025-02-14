// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskNewClusterLibrary struct {
	// cran block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/job#cran Job#cran}
	Cran *JobTaskNewClusterLibraryCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/job#egg Job#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/job#jar Job#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/job#maven Job#maven}
	Maven *JobTaskNewClusterLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// pypi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/job#pypi Job#pypi}
	Pypi *JobTaskNewClusterLibraryPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/job#requirements Job#requirements}.
	Requirements *string `field:"optional" json:"requirements" yaml:"requirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/resources/job#whl Job#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

