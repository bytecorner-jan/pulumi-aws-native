# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['InstanceAccessControlAttributeConfigurationArgs', 'InstanceAccessControlAttributeConfiguration']

@pulumi.input_type
class InstanceAccessControlAttributeConfigurationArgs:
    def __init__(__self__, *,
                 instance_arn: pulumi.Input[str],
                 access_control_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs']]]] = None):
        """
        The set of arguments for constructing a InstanceAccessControlAttributeConfiguration resource.
        :param pulumi.Input[str] instance_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-instancearn
        :param pulumi.Input[Sequence[pulumi.Input['InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs']]] access_control_attributes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributes
        """
        pulumi.set(__self__, "instance_arn", instance_arn)
        if access_control_attributes is not None:
            pulumi.set(__self__, "access_control_attributes", access_control_attributes)

    @property
    @pulumi.getter(name="instanceArn")
    def instance_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-instancearn
        """
        return pulumi.get(self, "instance_arn")

    @instance_arn.setter
    def instance_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_arn", value)

    @property
    @pulumi.getter(name="accessControlAttributes")
    def access_control_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributes
        """
        return pulumi.get(self, "access_control_attributes")

    @access_control_attributes.setter
    def access_control_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs']]]]):
        pulumi.set(self, "access_control_attributes", value)


class InstanceAccessControlAttributeConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_control_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs']]]]] = None,
                 instance_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs']]]] access_control_attributes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributes
        :param pulumi.Input[str] instance_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-instancearn
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceAccessControlAttributeConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html

        :param str resource_name: The name of the resource.
        :param InstanceAccessControlAttributeConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceAccessControlAttributeConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_control_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs']]]]] = None,
                 instance_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceAccessControlAttributeConfigurationArgs.__new__(InstanceAccessControlAttributeConfigurationArgs)

            __props__.__dict__["access_control_attributes"] = access_control_attributes
            if instance_arn is None and not opts.urn:
                raise TypeError("Missing required property 'instance_arn'")
            __props__.__dict__["instance_arn"] = instance_arn
        super(InstanceAccessControlAttributeConfiguration, __self__).__init__(
            'aws-native:SSO:InstanceAccessControlAttributeConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'InstanceAccessControlAttributeConfiguration':
        """
        Get an existing InstanceAccessControlAttributeConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InstanceAccessControlAttributeConfigurationArgs.__new__(InstanceAccessControlAttributeConfigurationArgs)

        __props__.__dict__["access_control_attributes"] = None
        __props__.__dict__["instance_arn"] = None
        return InstanceAccessControlAttributeConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessControlAttributes")
    def access_control_attributes(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceAccessControlAttributeConfigurationAccessControlAttribute']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributes
        """
        return pulumi.get(self, "access_control_attributes")

    @property
    @pulumi.getter(name="instanceArn")
    def instance_arn(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-instancearn
        """
        return pulumi.get(self, "instance_arn")

