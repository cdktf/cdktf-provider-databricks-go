// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type ClusterLibrary struct {
	// cran block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#cran Cluster#cran}
	Cran *ClusterLibraryCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#egg Cluster#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#jar Cluster#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#maven Cluster#maven}
	Maven *ClusterLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// pypi block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#pypi Cluster#pypi}
	Pypi *ClusterLibraryPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#whl Cluster#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}
