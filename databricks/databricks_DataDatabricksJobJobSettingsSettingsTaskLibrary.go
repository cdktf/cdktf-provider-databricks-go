// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsTaskLibrary struct {
	// cran block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#cran DataDatabricksJob#cran}
	Cran *DataDatabricksJobJobSettingsSettingsTaskLibraryCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#egg DataDatabricksJob#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#jar DataDatabricksJob#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#maven DataDatabricksJob#maven}
	Maven *DataDatabricksJobJobSettingsSettingsTaskLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// pypi block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#pypi DataDatabricksJob#pypi}
	Pypi *DataDatabricksJobJobSettingsSettingsTaskLibraryPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#whl DataDatabricksJob#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}
