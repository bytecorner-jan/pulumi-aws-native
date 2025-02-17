// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Route53::HealthCheck.
type HealthCheck struct {
	pulumi.CustomResourceState

	// A complex type that contains information about the health check.
	HealthCheckConfig HealthCheckConfigPropertiesOutput `pulumi:"healthCheckConfig"`
	HealthCheckId     pulumi.StringOutput               `pulumi:"healthCheckId"`
	// An array of key-value pairs to apply to this resource.
	HealthCheckTags HealthCheckTagArrayOutput `pulumi:"healthCheckTags"`
}

// NewHealthCheck registers a new resource with the given unique name, arguments, and options.
func NewHealthCheck(ctx *pulumi.Context,
	name string, args *HealthCheckArgs, opts ...pulumi.ResourceOption) (*HealthCheck, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HealthCheckConfig == nil {
		return nil, errors.New("invalid value for required argument 'HealthCheckConfig'")
	}
	var resource HealthCheck
	err := ctx.RegisterResource("aws-native:route53:HealthCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHealthCheck gets an existing HealthCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHealthCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthCheckState, opts ...pulumi.ResourceOption) (*HealthCheck, error) {
	var resource HealthCheck
	err := ctx.ReadResource("aws-native:route53:HealthCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HealthCheck resources.
type healthCheckState struct {
}

type HealthCheckState struct {
}

func (HealthCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckState)(nil)).Elem()
}

type healthCheckArgs struct {
	// A complex type that contains information about the health check.
	HealthCheckConfig HealthCheckConfigProperties `pulumi:"healthCheckConfig"`
	// An array of key-value pairs to apply to this resource.
	HealthCheckTags []HealthCheckTag `pulumi:"healthCheckTags"`
}

// The set of arguments for constructing a HealthCheck resource.
type HealthCheckArgs struct {
	// A complex type that contains information about the health check.
	HealthCheckConfig HealthCheckConfigPropertiesInput
	// An array of key-value pairs to apply to this resource.
	HealthCheckTags HealthCheckTagArrayInput
}

func (HealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckArgs)(nil)).Elem()
}

type HealthCheckInput interface {
	pulumi.Input

	ToHealthCheckOutput() HealthCheckOutput
	ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput
}

func (*HealthCheck) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthCheck)(nil))
}

func (i *HealthCheck) ToHealthCheckOutput() HealthCheckOutput {
	return i.ToHealthCheckOutputWithContext(context.Background())
}

func (i *HealthCheck) ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthCheckOutput)
}

type HealthCheckOutput struct{ *pulumi.OutputState }

func (HealthCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthCheck)(nil))
}

func (o HealthCheckOutput) ToHealthCheckOutput() HealthCheckOutput {
	return o
}

func (o HealthCheckOutput) ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HealthCheckInput)(nil)).Elem(), &HealthCheck{})
	pulumi.RegisterOutputType(HealthCheckOutput{})
}
