// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FeatureEngineeringKafkaConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_kafka_config#auth_config FeatureEngineeringKafkaConfig#auth_config}.
	AuthConfig *FeatureEngineeringKafkaConfigAuthConfig `field:"required" json:"authConfig" yaml:"authConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_kafka_config#bootstrap_servers FeatureEngineeringKafkaConfig#bootstrap_servers}.
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_kafka_config#subscription_mode FeatureEngineeringKafkaConfig#subscription_mode}.
	SubscriptionMode *FeatureEngineeringKafkaConfigSubscriptionMode `field:"required" json:"subscriptionMode" yaml:"subscriptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_kafka_config#extra_options FeatureEngineeringKafkaConfig#extra_options}.
	ExtraOptions *map[string]*string `field:"optional" json:"extraOptions" yaml:"extraOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_kafka_config#key_schema FeatureEngineeringKafkaConfig#key_schema}.
	KeySchema *FeatureEngineeringKafkaConfigKeySchema `field:"optional" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_kafka_config#value_schema FeatureEngineeringKafkaConfig#value_schema}.
	ValueSchema *FeatureEngineeringKafkaConfigValueSchema `field:"optional" json:"valueSchema" yaml:"valueSchema"`
}

