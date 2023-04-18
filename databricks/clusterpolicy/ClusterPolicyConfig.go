package clusterpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Policy definition JSON document expressed in Databricks Policy Definition Language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/cluster_policy#definition ClusterPolicy#definition}
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Cluster policy name. This must be unique. Length must be between 1 and 100 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/cluster_policy#name ClusterPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/cluster_policy#id ClusterPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Max number of clusters per user that can be active using this policy.
	//
	// If not present, there is no max limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/cluster_policy#max_clusters_per_user ClusterPolicy#max_clusters_per_user}
	MaxClustersPerUser *float64 `field:"optional" json:"maxClustersPerUser" yaml:"maxClustersPerUser"`
}

