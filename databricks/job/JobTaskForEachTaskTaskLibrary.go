// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskLibrary struct {
	// cran block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#cran Job#cran}
	Cran *JobTaskForEachTaskTaskLibraryCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#egg Job#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#jar Job#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#maven Job#maven}
	Maven *JobTaskForEachTaskTaskLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#provider_config Job#provider_config}
	ProviderConfig *JobTaskForEachTaskTaskLibraryProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// pypi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#pypi Job#pypi}
	Pypi *JobTaskForEachTaskTaskLibraryPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#requirements Job#requirements}.
	Requirements *string `field:"optional" json:"requirements" yaml:"requirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#whl Job#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

