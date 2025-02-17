// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the AWS::IoT::Certificate resource to declare an AWS IoT X.509 certificate.
type Certificate struct {
	pulumi.CustomResourceState

	Arn                       pulumi.StringOutput      `pulumi:"arn"`
	CACertificatePem          pulumi.StringPtrOutput   `pulumi:"cACertificatePem"`
	CertificateMode           CertificateModePtrOutput `pulumi:"certificateMode"`
	CertificatePem            pulumi.StringPtrOutput   `pulumi:"certificatePem"`
	CertificateSigningRequest pulumi.StringPtrOutput   `pulumi:"certificateSigningRequest"`
	Status                    CertificateStatusOutput  `pulumi:"status"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	var resource Certificate
	err := ctx.RegisterResource("aws-native:iot:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("aws-native:iot:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
}

type CertificateState struct {
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	CACertificatePem          *string           `pulumi:"cACertificatePem"`
	CertificateMode           *CertificateMode  `pulumi:"certificateMode"`
	CertificatePem            *string           `pulumi:"certificatePem"`
	CertificateSigningRequest *string           `pulumi:"certificateSigningRequest"`
	Status                    CertificateStatus `pulumi:"status"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	CACertificatePem          pulumi.StringPtrInput
	CertificateMode           CertificateModePtrInput
	CertificatePem            pulumi.StringPtrInput
	CertificateSigningRequest pulumi.StringPtrInput
	Status                    CertificateStatusInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterOutputType(CertificateOutput{})
}
