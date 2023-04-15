package pipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#allow_duplicate_names Pipeline#allow_duplicate_names}.
	AllowDuplicateNames interface{} `field:"optional" json:"allowDuplicateNames" yaml:"allowDuplicateNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#catalog Pipeline#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#channel Pipeline#channel}.
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#cluster Pipeline#cluster}
	Cluster interface{} `field:"optional" json:"cluster" yaml:"cluster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#configuration Pipeline#configuration}.
	Configuration *map[string]*string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#continuous Pipeline#continuous}.
	Continuous interface{} `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#development Pipeline#development}.
	Development interface{} `field:"optional" json:"development" yaml:"development"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#edition Pipeline#edition}.
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#filters Pipeline#filters}
	Filters *PipelineFilters `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#id Pipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#library Pipeline#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#name Pipeline#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#notification Pipeline#notification}
	Notification interface{} `field:"optional" json:"notification" yaml:"notification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#photon Pipeline#photon}.
	Photon interface{} `field:"optional" json:"photon" yaml:"photon"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#storage Pipeline#storage}.
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#target Pipeline#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#timeouts Pipeline#timeouts}
	Timeouts *PipelineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

