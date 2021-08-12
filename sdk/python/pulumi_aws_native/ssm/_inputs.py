# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AssociationInstanceAssociationOutputLocationArgs',
    'AssociationS3OutputLocationArgs',
    'AssociationTargetArgs',
    'DocumentAttachmentsSourceArgs',
    'DocumentDocumentRequiresArgs',
    'ResourceDataSyncAwsOrganizationsSourceArgs',
    'ResourceDataSyncS3DestinationArgs',
    'ResourceDataSyncSyncSourceArgs',
]

@pulumi.input_type
class AssociationInstanceAssociationOutputLocationArgs:
    def __init__(__self__, *,
                 s3_location: Optional[pulumi.Input['AssociationS3OutputLocationArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html
        :param pulumi.Input['AssociationS3OutputLocationArgs'] s3_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html#cfn-ssm-association-instanceassociationoutputlocation-s3location
        """
        if s3_location is not None:
            pulumi.set(__self__, "s3_location", s3_location)

    @property
    @pulumi.getter(name="s3Location")
    def s3_location(self) -> Optional[pulumi.Input['AssociationS3OutputLocationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html#cfn-ssm-association-instanceassociationoutputlocation-s3location
        """
        return pulumi.get(self, "s3_location")

    @s3_location.setter
    def s3_location(self, value: Optional[pulumi.Input['AssociationS3OutputLocationArgs']]):
        pulumi.set(self, "s3_location", value)


@pulumi.input_type
class AssociationS3OutputLocationArgs:
    def __init__(__self__, *,
                 output_s3_bucket_name: Optional[pulumi.Input[str]] = None,
                 output_s3_key_prefix: Optional[pulumi.Input[str]] = None,
                 output_s3_region: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html
        :param pulumi.Input[str] output_s3_bucket_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3bucketname
        :param pulumi.Input[str] output_s3_key_prefix: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3keyprefix
        :param pulumi.Input[str] output_s3_region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3region
        """
        if output_s3_bucket_name is not None:
            pulumi.set(__self__, "output_s3_bucket_name", output_s3_bucket_name)
        if output_s3_key_prefix is not None:
            pulumi.set(__self__, "output_s3_key_prefix", output_s3_key_prefix)
        if output_s3_region is not None:
            pulumi.set(__self__, "output_s3_region", output_s3_region)

    @property
    @pulumi.getter(name="outputS3BucketName")
    def output_s3_bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3bucketname
        """
        return pulumi.get(self, "output_s3_bucket_name")

    @output_s3_bucket_name.setter
    def output_s3_bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_s3_bucket_name", value)

    @property
    @pulumi.getter(name="outputS3KeyPrefix")
    def output_s3_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3keyprefix
        """
        return pulumi.get(self, "output_s3_key_prefix")

    @output_s3_key_prefix.setter
    def output_s3_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_s3_key_prefix", value)

    @property
    @pulumi.getter(name="outputS3Region")
    def output_s3_region(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3region
        """
        return pulumi.get(self, "output_s3_region")

    @output_s3_region.setter
    def output_s3_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_s3_region", value)


@pulumi.input_type
class AssociationTargetArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 values: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html
        :param pulumi.Input[str] key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-key
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-values
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-values
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class DocumentAttachmentsSourceArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html
        :param pulumi.Input[str] key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-key
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-values
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-values
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class DocumentDocumentRequiresArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-name
        :param pulumi.Input[str] version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-version
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-version
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class ResourceDataSyncAwsOrganizationsSourceArgs:
    def __init__(__self__, *,
                 organization_source_type: pulumi.Input[str],
                 organizational_units: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html
        :param pulumi.Input[str] organization_source_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationsourcetype
        :param pulumi.Input[Sequence[pulumi.Input[str]]] organizational_units: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationalunits
        """
        pulumi.set(__self__, "organization_source_type", organization_source_type)
        if organizational_units is not None:
            pulumi.set(__self__, "organizational_units", organizational_units)

    @property
    @pulumi.getter(name="organizationSourceType")
    def organization_source_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationsourcetype
        """
        return pulumi.get(self, "organization_source_type")

    @organization_source_type.setter
    def organization_source_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_source_type", value)

    @property
    @pulumi.getter(name="organizationalUnits")
    def organizational_units(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationalunits
        """
        return pulumi.get(self, "organizational_units")

    @organizational_units.setter
    def organizational_units(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "organizational_units", value)


@pulumi.input_type
class ResourceDataSyncS3DestinationArgs:
    def __init__(__self__, *,
                 bucket_name: pulumi.Input[str],
                 bucket_region: pulumi.Input[str],
                 sync_format: pulumi.Input[str],
                 bucket_prefix: Optional[pulumi.Input[str]] = None,
                 k_ms_key_arn: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html
        :param pulumi.Input[str] bucket_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketname
        :param pulumi.Input[str] bucket_region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketregion
        :param pulumi.Input[str] sync_format: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-syncformat
        :param pulumi.Input[str] bucket_prefix: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketprefix
        :param pulumi.Input[str] k_ms_key_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-kmskeyarn
        """
        pulumi.set(__self__, "bucket_name", bucket_name)
        pulumi.set(__self__, "bucket_region", bucket_region)
        pulumi.set(__self__, "sync_format", sync_format)
        if bucket_prefix is not None:
            pulumi.set(__self__, "bucket_prefix", bucket_prefix)
        if k_ms_key_arn is not None:
            pulumi.set(__self__, "k_ms_key_arn", k_ms_key_arn)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketname
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="bucketRegion")
    def bucket_region(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketregion
        """
        return pulumi.get(self, "bucket_region")

    @bucket_region.setter
    def bucket_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_region", value)

    @property
    @pulumi.getter(name="syncFormat")
    def sync_format(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-syncformat
        """
        return pulumi.get(self, "sync_format")

    @sync_format.setter
    def sync_format(self, value: pulumi.Input[str]):
        pulumi.set(self, "sync_format", value)

    @property
    @pulumi.getter(name="bucketPrefix")
    def bucket_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketprefix
        """
        return pulumi.get(self, "bucket_prefix")

    @bucket_prefix.setter
    def bucket_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_prefix", value)

    @property
    @pulumi.getter(name="kMSKeyArn")
    def k_ms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-kmskeyarn
        """
        return pulumi.get(self, "k_ms_key_arn")

    @k_ms_key_arn.setter
    def k_ms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "k_ms_key_arn", value)


@pulumi.input_type
class ResourceDataSyncSyncSourceArgs:
    def __init__(__self__, *,
                 source_regions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 source_type: pulumi.Input[str],
                 aws_organizations_source: Optional[pulumi.Input['ResourceDataSyncAwsOrganizationsSourceArgs']] = None,
                 include_future_regions: Optional[pulumi.Input[bool]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source_regions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourceregions
        :param pulumi.Input[str] source_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourcetype
        :param pulumi.Input['ResourceDataSyncAwsOrganizationsSourceArgs'] aws_organizations_source: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-awsorganizationssource
        :param pulumi.Input[bool] include_future_regions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-includefutureregions
        """
        pulumi.set(__self__, "source_regions", source_regions)
        pulumi.set(__self__, "source_type", source_type)
        if aws_organizations_source is not None:
            pulumi.set(__self__, "aws_organizations_source", aws_organizations_source)
        if include_future_regions is not None:
            pulumi.set(__self__, "include_future_regions", include_future_regions)

    @property
    @pulumi.getter(name="sourceRegions")
    def source_regions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourceregions
        """
        return pulumi.get(self, "source_regions")

    @source_regions.setter
    def source_regions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "source_regions", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourcetype
        """
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter(name="awsOrganizationsSource")
    def aws_organizations_source(self) -> Optional[pulumi.Input['ResourceDataSyncAwsOrganizationsSourceArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-awsorganizationssource
        """
        return pulumi.get(self, "aws_organizations_source")

    @aws_organizations_source.setter
    def aws_organizations_source(self, value: Optional[pulumi.Input['ResourceDataSyncAwsOrganizationsSourceArgs']]):
        pulumi.set(self, "aws_organizations_source", value)

    @property
    @pulumi.getter(name="includeFutureRegions")
    def include_future_regions(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-includefutureregions
        """
        return pulumi.get(self, "include_future_regions")

    @include_future_regions.setter
    def include_future_regions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_future_regions", value)


