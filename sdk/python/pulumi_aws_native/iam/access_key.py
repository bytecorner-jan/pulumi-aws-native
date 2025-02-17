# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccessKeyArgs', 'AccessKey']

@pulumi.input_type
class AccessKeyArgs:
    def __init__(__self__, *,
                 user_name: pulumi.Input[str],
                 serial: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AccessKey resource.
        """
        pulumi.set(__self__, "user_name", user_name)
        if serial is not None:
            pulumi.set(__self__, "serial", serial)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter
    def serial(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "serial")

    @serial.setter
    def serial(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "serial", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


warnings.warn("""AccessKey is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class AccessKey(pulumi.CustomResource):
    warnings.warn("""AccessKey is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 serial: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::IAM::AccessKey

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::IAM::AccessKey

        :param str resource_name: The name of the resource.
        :param AccessKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 serial: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""AccessKey is deprecated: AccessKey is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessKeyArgs.__new__(AccessKeyArgs)

            __props__.__dict__["serial"] = serial
            __props__.__dict__["status"] = status
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
            __props__.__dict__["secret_access_key"] = None
        super(AccessKey, __self__).__init__(
            'aws-native:iam:AccessKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AccessKey':
        """
        Get an existing AccessKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AccessKeyArgs.__new__(AccessKeyArgs)

        __props__.__dict__["secret_access_key"] = None
        __props__.__dict__["serial"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["user_name"] = None
        return AccessKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="secretAccessKey")
    def secret_access_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret_access_key")

    @property
    @pulumi.getter
    def serial(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "serial")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "user_name")

