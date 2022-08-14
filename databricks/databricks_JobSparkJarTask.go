// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobSparkJarTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#jar_uri Job#jar_uri}.
	JarUri *string `field:"optional" json:"jarUri" yaml:"jarUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#main_class_name Job#main_class_name}.
	MainClassName *string `field:"optional" json:"mainClassName" yaml:"mainClassName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#parameters Job#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

