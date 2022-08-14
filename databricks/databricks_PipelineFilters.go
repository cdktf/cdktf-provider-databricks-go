// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type PipelineFilters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#exclude Pipeline#exclude}.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#include Pipeline#include}.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

