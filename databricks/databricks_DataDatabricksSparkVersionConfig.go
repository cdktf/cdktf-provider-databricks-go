// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksSparkVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#beta DataDatabricksSparkVersion#beta}.
	Beta interface{} `field:"optional" json:"beta" yaml:"beta"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#genomics DataDatabricksSparkVersion#genomics}.
	Genomics interface{} `field:"optional" json:"genomics" yaml:"genomics"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#gpu DataDatabricksSparkVersion#gpu}.
	Gpu interface{} `field:"optional" json:"gpu" yaml:"gpu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#graviton DataDatabricksSparkVersion#graviton}.
	Graviton interface{} `field:"optional" json:"graviton" yaml:"graviton"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#id DataDatabricksSparkVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#latest DataDatabricksSparkVersion#latest}.
	Latest interface{} `field:"optional" json:"latest" yaml:"latest"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#long_term_support DataDatabricksSparkVersion#long_term_support}.
	LongTermSupport interface{} `field:"optional" json:"longTermSupport" yaml:"longTermSupport"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#ml DataDatabricksSparkVersion#ml}.
	Ml interface{} `field:"optional" json:"ml" yaml:"ml"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#photon DataDatabricksSparkVersion#photon}.
	Photon interface{} `field:"optional" json:"photon" yaml:"photon"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#scala DataDatabricksSparkVersion#scala}.
	Scala *string `field:"optional" json:"scala" yaml:"scala"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/spark_version#spark_version DataDatabricksSparkVersion#spark_version}.
	SparkVersion *string `field:"optional" json:"sparkVersion" yaml:"sparkVersion"`
}
