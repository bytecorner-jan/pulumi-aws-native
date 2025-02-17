// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws-native:appstream:DirectoryConfig":
		r = &DirectoryConfig{}
	case "aws-native:appstream:Fleet":
		r = &Fleet{}
	case "aws-native:appstream:ImageBuilder":
		r = &ImageBuilder{}
	case "aws-native:appstream:Stack":
		r = &Stack{}
	case "aws-native:appstream:StackFleetAssociation":
		r = &StackFleetAssociation{}
	case "aws-native:appstream:StackUserAssociation":
		r = &StackUserAssociation{}
	case "aws-native:appstream:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"aws-native",
		"appstream",
		&module{version},
	)
}
