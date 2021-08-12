# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceAccessControlAttributeConfigurationAccessControlAttributeValueArgs',
    'InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs',
]

@pulumi.input_type
class InstanceAccessControlAttributeConfigurationAccessControlAttributeValueArgs:
    def __init__(__self__, *,
                 source: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue.html
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue-source
        """
        pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue-source
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "source", value)


@pulumi.input_type
class InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input['InstanceAccessControlAttributeConfigurationAccessControlAttributeValueArgs']):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute.html
        :param pulumi.Input[str] key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute-key
        :param pulumi.Input['InstanceAccessControlAttributeConfigurationAccessControlAttributeValueArgs'] value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute-value
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute-key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input['InstanceAccessControlAttributeConfigurationAccessControlAttributeValueArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute-value
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input['InstanceAccessControlAttributeConfigurationAccessControlAttributeValueArgs']):
        pulumi.set(self, "value", value)


