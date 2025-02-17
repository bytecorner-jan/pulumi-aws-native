# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'LedgerTag',
    'StreamKinesisConfiguration',
    'StreamTag',
]

@pulumi.output_type
class LedgerTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class StreamKinesisConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "aggregationEnabled":
            suggest = "aggregation_enabled"
        elif key == "streamArn":
            suggest = "stream_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StreamKinesisConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StreamKinesisConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StreamKinesisConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 aggregation_enabled: Optional[bool] = None,
                 stream_arn: Optional[str] = None):
        if aggregation_enabled is not None:
            pulumi.set(__self__, "aggregation_enabled", aggregation_enabled)
        if stream_arn is not None:
            pulumi.set(__self__, "stream_arn", stream_arn)

    @property
    @pulumi.getter(name="aggregationEnabled")
    def aggregation_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "aggregation_enabled")

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> Optional[str]:
        return pulumi.get(self, "stream_arn")


@pulumi.output_type
class StreamTag(dict):
    """
    A key-value pair to associate with a resource.
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A key-value pair to associate with a resource.
        :param str key: The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        :param str value: The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        return pulumi.get(self, "value")


