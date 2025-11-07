// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterpolicy


type ClusterPolicyLibraries struct {
	// cran block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/cluster_policy#cran ClusterPolicy#cran}
	Cran *ClusterPolicyLibrariesCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/cluster_policy#egg ClusterPolicy#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/cluster_policy#jar ClusterPolicy#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/cluster_policy#maven ClusterPolicy#maven}
	Maven *ClusterPolicyLibrariesMaven `field:"optional" json:"maven" yaml:"maven"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/cluster_policy#provider_config ClusterPolicy#provider_config}
	ProviderConfig *ClusterPolicyLibrariesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// pypi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/cluster_policy#pypi ClusterPolicy#pypi}
	Pypi *ClusterPolicyLibrariesPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/cluster_policy#requirements ClusterPolicy#requirements}.
	Requirements *string `field:"optional" json:"requirements" yaml:"requirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/cluster_policy#whl ClusterPolicy#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

