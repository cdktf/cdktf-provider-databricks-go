// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoSpecLibrary struct {
	// cran block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/cluster#cran DataDatabricksCluster#cran}
	Cran *DataDatabricksClusterClusterInfoSpecLibraryCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/cluster#egg DataDatabricksCluster#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/cluster#jar DataDatabricksCluster#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/cluster#maven DataDatabricksCluster#maven}
	Maven *DataDatabricksClusterClusterInfoSpecLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/cluster#provider_config DataDatabricksCluster#provider_config}
	ProviderConfig *DataDatabricksClusterClusterInfoSpecLibraryProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// pypi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/cluster#pypi DataDatabricksCluster#pypi}
	Pypi *DataDatabricksClusterClusterInfoSpecLibraryPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/cluster#requirements DataDatabricksCluster#requirements}.
	Requirements *string `field:"optional" json:"requirements" yaml:"requirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/cluster#whl DataDatabricksCluster#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

