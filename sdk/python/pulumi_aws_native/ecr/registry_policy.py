# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RegistryPolicyArgs', 'RegistryPolicy']

@pulumi.input_type
class RegistryPolicyArgs:
    def __init__(__self__, *,
                 policy_text: pulumi.Input[Union[Any, str]]):
        """
        The set of arguments for constructing a RegistryPolicy resource.
        :param pulumi.Input[Union[Any, str]] policy_text: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html#cfn-ecr-registrypolicy-policytext
        """
        pulumi.set(__self__, "policy_text", policy_text)

    @property
    @pulumi.getter(name="policyText")
    def policy_text(self) -> pulumi.Input[Union[Any, str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html#cfn-ecr-registrypolicy-policytext
        """
        return pulumi.get(self, "policy_text")

    @policy_text.setter
    def policy_text(self, value: pulumi.Input[Union[Any, str]]):
        pulumi.set(self, "policy_text", value)


class RegistryPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_text: Optional[pulumi.Input[Union[Any, str]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[Any, str]] policy_text: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html#cfn-ecr-registrypolicy-policytext
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegistryPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html

        :param str resource_name: The name of the resource.
        :param RegistryPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_text: Optional[pulumi.Input[Union[Any, str]]] = None,
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
            __props__ = RegistryPolicyArgs.__new__(RegistryPolicyArgs)

            if policy_text is None and not opts.urn:
                raise TypeError("Missing required property 'policy_text'")
            __props__.__dict__["policy_text"] = policy_text
            __props__.__dict__["registry_id"] = None
        super(RegistryPolicy, __self__).__init__(
            'aws-native:ECR:RegistryPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RegistryPolicy':
        """
        Get an existing RegistryPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegistryPolicyArgs.__new__(RegistryPolicyArgs)

        __props__.__dict__["policy_text"] = None
        __props__.__dict__["registry_id"] = None
        return RegistryPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policyText")
    def policy_text(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html#cfn-ecr-registrypolicy-policytext
        """
        return pulumi.get(self, "policy_text")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "registry_id")

