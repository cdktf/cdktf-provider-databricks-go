package sqlalert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlAlertConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_alert#name SqlAlert#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_alert#options SqlAlert#options}
	Options *SqlAlertOptions `field:"required" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_alert#query_id SqlAlert#query_id}.
	QueryId *string `field:"required" json:"queryId" yaml:"queryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_alert#id SqlAlert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_alert#parent SqlAlert#parent}.
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_alert#rearm SqlAlert#rearm}.
	Rearm *float64 `field:"optional" json:"rearm" yaml:"rearm"`
}
