// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MountConfig struct {
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
	// abfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#abfs Mount#abfs}
	Abfs *MountAbfs `field:"optional" json:"abfs" yaml:"abfs"`
	// adl block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#adl Mount#adl}
	Adl *MountAdl `field:"optional" json:"adl" yaml:"adl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#cluster_id Mount#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#encryption_type Mount#encryption_type}.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#extra_configs Mount#extra_configs}.
	ExtraConfigs *map[string]*string `field:"optional" json:"extraConfigs" yaml:"extraConfigs"`
	// gs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#gs Mount#gs}
	Gs *MountGs `field:"optional" json:"gs" yaml:"gs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#id Mount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#name Mount#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#resource_id Mount#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#s3 Mount#s3}
	S3 *MountS3 `field:"optional" json:"s3" yaml:"s3"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#timeouts Mount#timeouts}
	Timeouts *MountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#uri Mount#uri}.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
	// wasb block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#wasb Mount#wasb}
	Wasb *MountWasb `field:"optional" json:"wasb" yaml:"wasb"`
}

