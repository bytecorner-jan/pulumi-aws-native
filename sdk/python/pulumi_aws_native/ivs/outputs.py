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
    'RecordingConfigurationDestinationConfiguration',
    'RecordingConfigurationS3DestinationConfiguration',
]

@pulumi.output_type
class RecordingConfigurationDestinationConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-destinationconfiguration.html
    """
    def __init__(__self__, *,
                 s3: 'outputs.RecordingConfigurationS3DestinationConfiguration'):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-destinationconfiguration.html
        :param 'RecordingConfigurationS3DestinationConfiguration' s3: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-destinationconfiguration.html#cfn-ivs-recordingconfiguration-destinationconfiguration-s3
        """
        pulumi.set(__self__, "s3", s3)

    @property
    @pulumi.getter
    def s3(self) -> 'outputs.RecordingConfigurationS3DestinationConfiguration':
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-destinationconfiguration.html#cfn-ivs-recordingconfiguration-destinationconfiguration-s3
        """
        return pulumi.get(self, "s3")


@pulumi.output_type
class RecordingConfigurationS3DestinationConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-s3destinationconfiguration.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bucketName":
            suggest = "bucket_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RecordingConfigurationS3DestinationConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RecordingConfigurationS3DestinationConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RecordingConfigurationS3DestinationConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bucket_name: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-s3destinationconfiguration.html
        :param str bucket_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-s3destinationconfiguration.html#cfn-ivs-recordingconfiguration-s3destinationconfiguration-bucketname
        """
        pulumi.set(__self__, "bucket_name", bucket_name)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-s3destinationconfiguration.html#cfn-ivs-recordingconfiguration-s3destinationconfiguration-bucketname
        """
        return pulumi.get(self, "bucket_name")


