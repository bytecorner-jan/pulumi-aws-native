# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'AssociationInstanceAssociationOutputLocation',
    'AssociationS3OutputLocation',
    'AssociationTarget',
    'DocumentAttachmentsSource',
    'DocumentDocumentRequires',
    'ResourceDataSyncAwsOrganizationsSource',
    'ResourceDataSyncS3Destination',
    'ResourceDataSyncSyncSource',
]

@pulumi.output_type
class AssociationInstanceAssociationOutputLocation(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "s3Location":
            suggest = "s3_location"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AssociationInstanceAssociationOutputLocation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AssociationInstanceAssociationOutputLocation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AssociationInstanceAssociationOutputLocation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 s3_location: Optional['outputs.AssociationS3OutputLocation'] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html
        :param 'AssociationS3OutputLocation' s3_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html#cfn-ssm-association-instanceassociationoutputlocation-s3location
        """
        if s3_location is not None:
            pulumi.set(__self__, "s3_location", s3_location)

    @property
    @pulumi.getter(name="s3Location")
    def s3_location(self) -> Optional['outputs.AssociationS3OutputLocation']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html#cfn-ssm-association-instanceassociationoutputlocation-s3location
        """
        return pulumi.get(self, "s3_location")


@pulumi.output_type
class AssociationS3OutputLocation(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "outputS3BucketName":
            suggest = "output_s3_bucket_name"
        elif key == "outputS3KeyPrefix":
            suggest = "output_s3_key_prefix"
        elif key == "outputS3Region":
            suggest = "output_s3_region"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AssociationS3OutputLocation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AssociationS3OutputLocation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AssociationS3OutputLocation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 output_s3_bucket_name: Optional[str] = None,
                 output_s3_key_prefix: Optional[str] = None,
                 output_s3_region: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html
        :param str output_s3_bucket_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3bucketname
        :param str output_s3_key_prefix: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3keyprefix
        :param str output_s3_region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3region
        """
        if output_s3_bucket_name is not None:
            pulumi.set(__self__, "output_s3_bucket_name", output_s3_bucket_name)
        if output_s3_key_prefix is not None:
            pulumi.set(__self__, "output_s3_key_prefix", output_s3_key_prefix)
        if output_s3_region is not None:
            pulumi.set(__self__, "output_s3_region", output_s3_region)

    @property
    @pulumi.getter(name="outputS3BucketName")
    def output_s3_bucket_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3bucketname
        """
        return pulumi.get(self, "output_s3_bucket_name")

    @property
    @pulumi.getter(name="outputS3KeyPrefix")
    def output_s3_key_prefix(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3keyprefix
        """
        return pulumi.get(self, "output_s3_key_prefix")

    @property
    @pulumi.getter(name="outputS3Region")
    def output_s3_region(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3region
        """
        return pulumi.get(self, "output_s3_region")


@pulumi.output_type
class AssociationTarget(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html
    """
    def __init__(__self__, *,
                 key: str,
                 values: Sequence[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html
        :param str key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-key
        :param Sequence[str] values: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-values
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-values
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class DocumentAttachmentsSource(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html
    """
    def __init__(__self__, *,
                 key: Optional[str] = None,
                 name: Optional[str] = None,
                 values: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html
        :param str key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-key
        :param str name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-name
        :param Sequence[str] values: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-values
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-attachmentssource.html#cfn-ssm-document-attachmentssource-values
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class DocumentDocumentRequires(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html
    """
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 version: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html
        :param str name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-name
        :param str version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-version
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-version
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class ResourceDataSyncAwsOrganizationsSource(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "organizationSourceType":
            suggest = "organization_source_type"
        elif key == "organizationalUnits":
            suggest = "organizational_units"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceDataSyncAwsOrganizationsSource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceDataSyncAwsOrganizationsSource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceDataSyncAwsOrganizationsSource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 organization_source_type: str,
                 organizational_units: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html
        :param str organization_source_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationsourcetype
        :param Sequence[str] organizational_units: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationalunits
        """
        pulumi.set(__self__, "organization_source_type", organization_source_type)
        if organizational_units is not None:
            pulumi.set(__self__, "organizational_units", organizational_units)

    @property
    @pulumi.getter(name="organizationSourceType")
    def organization_source_type(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationsourcetype
        """
        return pulumi.get(self, "organization_source_type")

    @property
    @pulumi.getter(name="organizationalUnits")
    def organizational_units(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationalunits
        """
        return pulumi.get(self, "organizational_units")


@pulumi.output_type
class ResourceDataSyncS3Destination(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bucketName":
            suggest = "bucket_name"
        elif key == "bucketRegion":
            suggest = "bucket_region"
        elif key == "syncFormat":
            suggest = "sync_format"
        elif key == "bucketPrefix":
            suggest = "bucket_prefix"
        elif key == "kMSKeyArn":
            suggest = "k_ms_key_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceDataSyncS3Destination. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceDataSyncS3Destination.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceDataSyncS3Destination.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bucket_name: str,
                 bucket_region: str,
                 sync_format: str,
                 bucket_prefix: Optional[str] = None,
                 k_ms_key_arn: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html
        :param str bucket_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketname
        :param str bucket_region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketregion
        :param str sync_format: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-syncformat
        :param str bucket_prefix: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketprefix
        :param str k_ms_key_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-kmskeyarn
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
    def bucket_name(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketname
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="bucketRegion")
    def bucket_region(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketregion
        """
        return pulumi.get(self, "bucket_region")

    @property
    @pulumi.getter(name="syncFormat")
    def sync_format(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-syncformat
        """
        return pulumi.get(self, "sync_format")

    @property
    @pulumi.getter(name="bucketPrefix")
    def bucket_prefix(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-bucketprefix
        """
        return pulumi.get(self, "bucket_prefix")

    @property
    @pulumi.getter(name="kMSKeyArn")
    def k_ms_key_arn(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-s3destination.html#cfn-ssm-resourcedatasync-s3destination-kmskeyarn
        """
        return pulumi.get(self, "k_ms_key_arn")


@pulumi.output_type
class ResourceDataSyncSyncSource(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sourceRegions":
            suggest = "source_regions"
        elif key == "sourceType":
            suggest = "source_type"
        elif key == "awsOrganizationsSource":
            suggest = "aws_organizations_source"
        elif key == "includeFutureRegions":
            suggest = "include_future_regions"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceDataSyncSyncSource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceDataSyncSyncSource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceDataSyncSyncSource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 source_regions: Sequence[str],
                 source_type: str,
                 aws_organizations_source: Optional['outputs.ResourceDataSyncAwsOrganizationsSource'] = None,
                 include_future_regions: Optional[bool] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html
        :param Sequence[str] source_regions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourceregions
        :param str source_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourcetype
        :param 'ResourceDataSyncAwsOrganizationsSource' aws_organizations_source: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-awsorganizationssource
        :param bool include_future_regions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-includefutureregions
        """
        pulumi.set(__self__, "source_regions", source_regions)
        pulumi.set(__self__, "source_type", source_type)
        if aws_organizations_source is not None:
            pulumi.set(__self__, "aws_organizations_source", aws_organizations_source)
        if include_future_regions is not None:
            pulumi.set(__self__, "include_future_regions", include_future_regions)

    @property
    @pulumi.getter(name="sourceRegions")
    def source_regions(self) -> Sequence[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourceregions
        """
        return pulumi.get(self, "source_regions")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourcetype
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="awsOrganizationsSource")
    def aws_organizations_source(self) -> Optional['outputs.ResourceDataSyncAwsOrganizationsSource']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-awsorganizationssource
        """
        return pulumi.get(self, "aws_organizations_source")

    @property
    @pulumi.getter(name="includeFutureRegions")
    def include_future_regions(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-includefutureregions
        """
        return pulumi.get(self, "include_future_regions")


