// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the AWS native package. By default, resources use package-wide configuration settings, however an explicit `Provider` instance may be created and passed during resource construction to achieve fine-grained programmatic control over provider settings. See the [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.
	AccessKey pulumi.StringPtrOutput `pulumi:"accessKey"`
	// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
	Profile pulumi.StringPtrOutput `pulumi:"profile"`
	// The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	SecretKey pulumi.StringPtrOutput `pulumi:"secretKey"`
	// The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.
	SharedCredentialsFile pulumi.StringPtrOutput `pulumi:"sharedCredentialsFile"`
	// Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.
	Token pulumi.StringPtrOutput `pulumi:"token"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKey == nil {
		args.AccessKey = pulumi.StringPtr(getEnvOrDefault("", nil, "AWS_ACCESS_KEY_ID").(string))
	}
	if args.Profile == nil {
		args.Profile = pulumi.StringPtr(getEnvOrDefault("", nil, "AWS_PROFILE").(string))
	}
	if args.Region == nil {
		args.Region = pulumi.String(getEnvOrDefault("", nil, "AWS_REGION", "AWS_DEFAULT_REGION").(string))
	}
	if args.SecretKey == nil {
		args.SecretKey = pulumi.StringPtr(getEnvOrDefault("", nil, "AWS_SECRET_ACCESS_KEY").(string))
	}
	if args.SharedCredentialsFile == nil {
		args.SharedCredentialsFile = pulumi.StringPtr(getEnvOrDefault("", nil, "AWS_SHARED_CREDENTIALS_FILE").(string))
	}
	if args.SkipCredentialsValidation == nil {
		args.SkipCredentialsValidation = pulumi.BoolPtr(true)
	}
	if args.SkipGetEc2Platforms == nil {
		args.SkipGetEc2Platforms = pulumi.BoolPtr(true)
	}
	if args.SkipMetadataApiCheck == nil {
		args.SkipMetadataApiCheck = pulumi.BoolPtr(true)
	}
	if args.SkipRegionValidation == nil {
		args.SkipRegionValidation = pulumi.BoolPtr(true)
	}
	if args.Token == nil {
		args.Token = pulumi.StringPtr(getEnvOrDefault("", nil, "AWS_SESSION_TOKEN").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:aws-native", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.
	AccessKey *string `pulumi:"accessKey"`
	// List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.
	AllowedAccountIds []string `pulumi:"allowedAccountIds"`
	// Configuration for retrieving temporary credentials from the STS service.
	AssumeRole *ProviderAssumeRole `pulumi:"assumeRole"`
	// Configuration block with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.
	DefaultTags *ProviderDefaultTags `pulumi:"defaultTags"`
	// Configuration block for customizing service endpoints.
	Endpoints []ProviderEndpoint `pulumi:"endpoints"`
	// List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.
	ForbiddenAccountIds []string `pulumi:"forbiddenAccountIds"`
	// Configuration block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.
	IgnoreTags *ProviderIgnoreTags `pulumi:"ignoreTags"`
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
	Insecure *bool `pulumi:"insecure"`
	// The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
	MaxRetries *int `pulumi:"maxRetries"`
	// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
	Profile *string `pulumi:"profile"`
	// The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.
	Region string `pulumi:"region"`
	// Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.
	S3ForcePathStyle *bool `pulumi:"s3ForcePathStyle"`
	// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	SecretKey *string `pulumi:"secretKey"`
	// The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.
	SharedCredentialsFile *string `pulumi:"sharedCredentialsFile"`
	// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.
	SkipCredentialsValidation *bool `pulumi:"skipCredentialsValidation"`
	// Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.
	SkipGetEc2Platforms *bool `pulumi:"skipGetEc2Platforms"`
	// Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.
	SkipMetadataApiCheck *bool `pulumi:"skipMetadataApiCheck"`
	// Skip static validation of region name. Used by users of alternative AWS-like APIs or users with access to regions that are not public.
	SkipRegionValidation *bool `pulumi:"skipRegionValidation"`
	// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
	SkipRequestingAccountId *bool `pulumi:"skipRequestingAccountId"`
	// Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.
	AccessKey pulumi.StringPtrInput
	// List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.
	AllowedAccountIds pulumi.StringArrayInput
	// Configuration for retrieving temporary credentials from the STS service.
	AssumeRole ProviderAssumeRolePtrInput
	// Configuration block with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.
	DefaultTags ProviderDefaultTagsPtrInput
	// Configuration block for customizing service endpoints.
	Endpoints ProviderEndpointArrayInput
	// List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.
	ForbiddenAccountIds pulumi.StringArrayInput
	// Configuration block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.
	IgnoreTags ProviderIgnoreTagsPtrInput
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
	Insecure pulumi.BoolPtrInput
	// The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
	MaxRetries pulumi.IntPtrInput
	// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
	Profile pulumi.StringPtrInput
	// The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.
	Region pulumi.StringInput
	// Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.
	S3ForcePathStyle pulumi.BoolPtrInput
	// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	SecretKey pulumi.StringPtrInput
	// The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.
	SharedCredentialsFile pulumi.StringPtrInput
	// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.
	SkipCredentialsValidation pulumi.BoolPtrInput
	// Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.
	SkipGetEc2Platforms pulumi.BoolPtrInput
	// Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.
	SkipMetadataApiCheck pulumi.BoolPtrInput
	// Skip static validation of region name. Used by users of alternative AWS-like APIs or users with access to regions that are not public.
	SkipRegionValidation pulumi.BoolPtrInput
	// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
	SkipRequestingAccountId pulumi.BoolPtrInput
	// Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.
	Token pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
}
