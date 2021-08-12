# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ResourcePolicyArgs', 'ResourcePolicy']

@pulumi.input_type
class ResourcePolicyArgs:
    def __init__(__self__, *,
                 policy_document: pulumi.Input[str],
                 policy_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a ResourcePolicy resource.
        :param pulumi.Input[str] policy_document: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policydocument
        :param pulumi.Input[str] policy_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policyname
        """
        pulumi.set(__self__, "policy_document", policy_document)
        pulumi.set(__self__, "policy_name", policy_name)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policydocument
        """
        return pulumi.get(self, "policy_document")

    @policy_document.setter
    def policy_document(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_document", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policyname
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_name", value)


class ResourcePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_document: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_document: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policydocument
        :param pulumi.Input[str] policy_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policyname
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourcePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html

        :param str resource_name: The name of the resource.
        :param ResourcePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourcePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_document: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = ResourcePolicyArgs.__new__(ResourcePolicyArgs)

            if policy_document is None and not opts.urn:
                raise TypeError("Missing required property 'policy_document'")
            __props__.__dict__["policy_document"] = policy_document
            if policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_name'")
            __props__.__dict__["policy_name"] = policy_name
        super(ResourcePolicy, __self__).__init__(
            'aws-native:Logs:ResourcePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ResourcePolicy':
        """
        Get an existing ResourcePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ResourcePolicyArgs.__new__(ResourcePolicyArgs)

        __props__.__dict__["policy_document"] = None
        __props__.__dict__["policy_name"] = None
        return ResourcePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policydocument
        """
        return pulumi.get(self, "policy_document")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policyname
        """
        return pulumi.get(self, "policy_name")

