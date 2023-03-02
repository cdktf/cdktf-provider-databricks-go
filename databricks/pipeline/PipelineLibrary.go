package pipeline


type PipelineLibrary struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#file Pipeline#file}
	File *PipelineLibraryFile `field:"optional" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#jar Pipeline#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// maven block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#maven Pipeline#maven}
	Maven *PipelineLibraryMaven `field:"optional" json:"maven" yaml:"maven"`
	// notebook block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#notebook Pipeline#notebook}
	Notebook *PipelineLibraryNotebook `field:"optional" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#whl Pipeline#whl}.
	Whl *string `field:"optional" json:"whl" yaml:"whl"`
}

