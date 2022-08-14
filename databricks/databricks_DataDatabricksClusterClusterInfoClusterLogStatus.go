// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksClusterClusterInfoClusterLogStatus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#last_attempted DataDatabricksCluster#last_attempted}.
	LastAttempted *float64 `field:"optional" json:"lastAttempted" yaml:"lastAttempted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#last_exception DataDatabricksCluster#last_exception}.
	LastException *string `field:"optional" json:"lastException" yaml:"lastException"`
}

