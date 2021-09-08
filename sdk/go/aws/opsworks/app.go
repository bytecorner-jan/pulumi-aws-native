// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html
type App struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-appsource
	AppSource AppSourcePtrOutput `pulumi:"appSource"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-attributes
	Attributes pulumi.StringMapOutput `pulumi:"attributes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-datasources
	DataSources AppDataSourceArrayOutput `pulumi:"dataSources"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-domains
	Domains pulumi.StringArrayOutput `pulumi:"domains"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-enablessl
	EnableSsl pulumi.BoolPtrOutput `pulumi:"enableSsl"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-environment
	Environment AppEnvironmentVariableArrayOutput `pulumi:"environment"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-name
	Name pulumi.StringOutput `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-shortname
	Shortname pulumi.StringPtrOutput `pulumi:"shortname"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-sslconfiguration
	SslConfiguration AppSslConfigurationPtrOutput `pulumi:"sslConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-stackid
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource App
	err := ctx.RegisterResource("aws-native:opsworks:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("aws-native:opsworks:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
}

type AppState struct {
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-appsource
	AppSource *AppSource `pulumi:"appSource"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-attributes
	Attributes map[string]string `pulumi:"attributes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-datasources
	DataSources []AppDataSource `pulumi:"dataSources"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-domains
	Domains []string `pulumi:"domains"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-enablessl
	EnableSsl *bool `pulumi:"enableSsl"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-environment
	Environment []AppEnvironmentVariable `pulumi:"environment"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-name
	Name string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-shortname
	Shortname *string `pulumi:"shortname"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-sslconfiguration
	SslConfiguration *AppSslConfiguration `pulumi:"sslConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-stackid
	StackId string `pulumi:"stackId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-type
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-appsource
	AppSource AppSourcePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-attributes
	Attributes pulumi.StringMapInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-datasources
	DataSources AppDataSourceArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-domains
	Domains pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-enablessl
	EnableSsl pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-environment
	Environment AppEnvironmentVariableArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-name
	Name pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-shortname
	Shortname pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-sslconfiguration
	SslConfiguration AppSslConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-stackid
	StackId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-type
	Type pulumi.StringInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((*App)(nil))
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*App)(nil))
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppOutput{})
}
