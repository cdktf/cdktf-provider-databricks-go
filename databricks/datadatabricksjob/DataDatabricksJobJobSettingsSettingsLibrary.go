package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsLibrary struct {
	// cran block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/data-sources/job#cran DataDatabricksJob#cran}
	Cran *DataDatabricksJobJobSettingsSettingsLibraryCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/data-sources/job#egg DataDatabricksJob#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/data-sources/job#jar DataDatabricksJob#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/data-sources/job#maven DataDatabricksJob#maven}
	Maven *DataDatabricksJobJobSettingsSettingsLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// pypi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/data-sources/job#pypi DataDatabricksJob#pypi}
	Pypi *DataDatabricksJobJobSettingsSettingsLibraryPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/data-sources/job#whl DataDatabricksJob#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

