# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs

__all__ = ['ReplicaKeyArgs', 'ReplicaKey']

@pulumi.input_type
class ReplicaKeyArgs:
    def __init__(__self__, *,
                 key_policy: pulumi.Input[Union[Any, str]],
                 primary_key_arn: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 pending_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a ReplicaKey resource.
        :param pulumi.Input[Union[Any, str]] key_policy: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-keypolicy
        :param pulumi.Input[str] primary_key_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-primarykeyarn
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-description
        :param pulumi.Input[bool] enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-enabled
        :param pulumi.Input[int] pending_window_in_days: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-pendingwindowindays
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-tags
        """
        pulumi.set(__self__, "key_policy", key_policy)
        pulumi.set(__self__, "primary_key_arn", primary_key_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if pending_window_in_days is not None:
            pulumi.set(__self__, "pending_window_in_days", pending_window_in_days)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="keyPolicy")
    def key_policy(self) -> pulumi.Input[Union[Any, str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-keypolicy
        """
        return pulumi.get(self, "key_policy")

    @key_policy.setter
    def key_policy(self, value: pulumi.Input[Union[Any, str]]):
        pulumi.set(self, "key_policy", value)

    @property
    @pulumi.getter(name="primaryKeyArn")
    def primary_key_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-primarykeyarn
        """
        return pulumi.get(self, "primary_key_arn")

    @primary_key_arn.setter
    def primary_key_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "primary_key_arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-enabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="pendingWindowInDays")
    def pending_window_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-pendingwindowindays
        """
        return pulumi.get(self, "pending_window_in_days")

    @pending_window_in_days.setter
    def pending_window_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pending_window_in_days", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class ReplicaKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key_policy: Optional[pulumi.Input[Union[Any, str]]] = None,
                 pending_window_in_days: Optional[pulumi.Input[int]] = None,
                 primary_key_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-description
        :param pulumi.Input[bool] enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-enabled
        :param pulumi.Input[Union[Any, str]] key_policy: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-keypolicy
        :param pulumi.Input[int] pending_window_in_days: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-pendingwindowindays
        :param pulumi.Input[str] primary_key_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-primarykeyarn
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicaKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html

        :param str resource_name: The name of the resource.
        :param ReplicaKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicaKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key_policy: Optional[pulumi.Input[Union[Any, str]]] = None,
                 pending_window_in_days: Optional[pulumi.Input[int]] = None,
                 primary_key_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
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
            __props__ = ReplicaKeyArgs.__new__(ReplicaKeyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            if key_policy is None and not opts.urn:
                raise TypeError("Missing required property 'key_policy'")
            __props__.__dict__["key_policy"] = key_policy
            __props__.__dict__["pending_window_in_days"] = pending_window_in_days
            if primary_key_arn is None and not opts.urn:
                raise TypeError("Missing required property 'primary_key_arn'")
            __props__.__dict__["primary_key_arn"] = primary_key_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["key_id"] = None
        super(ReplicaKey, __self__).__init__(
            'aws-native:KMS:ReplicaKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ReplicaKey':
        """
        Get an existing ReplicaKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReplicaKeyArgs.__new__(ReplicaKeyArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["key_id"] = None
        __props__.__dict__["key_policy"] = None
        __props__.__dict__["pending_window_in_days"] = None
        __props__.__dict__["primary_key_arn"] = None
        __props__.__dict__["tags"] = None
        return ReplicaKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-enabled
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyPolicy")
    def key_policy(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-keypolicy
        """
        return pulumi.get(self, "key_policy")

    @property
    @pulumi.getter(name="pendingWindowInDays")
    def pending_window_in_days(self) -> pulumi.Output[Optional[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-pendingwindowindays
        """
        return pulumi.get(self, "pending_window_in_days")

    @property
    @pulumi.getter(name="primaryKeyArn")
    def primary_key_arn(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-primarykeyarn
        """
        return pulumi.get(self, "primary_key_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html#cfn-kms-replicakey-tags
        """
        return pulumi.get(self, "tags")

