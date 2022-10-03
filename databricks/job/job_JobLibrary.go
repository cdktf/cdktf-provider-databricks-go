package job


type JobLibrary struct {
	// cran block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#cran Job#cran}
	Cran *JobLibraryCran `field:"optional" json:"cran" yaml:"cran"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#egg Job#egg}.
	Egg *string `field:"optional" json:"egg" yaml:"egg"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#jar Job#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#maven Job#maven}
	Maven *JobLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// pypi block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#pypi Job#pypi}
	Pypi *JobLibraryPypi `field:"optional" json:"pypi" yaml:"pypi"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#whl Job#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

