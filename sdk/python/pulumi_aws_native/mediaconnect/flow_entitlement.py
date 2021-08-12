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

__all__ = ['FlowEntitlementArgs', 'FlowEntitlement']

@pulumi.input_type
class FlowEntitlementArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 flow_arn: pulumi.Input[str],
                 name: pulumi.Input[str],
                 subscribers: pulumi.Input[Sequence[pulumi.Input[str]]],
                 data_transfer_subscriber_fee_percent: Optional[pulumi.Input[int]] = None,
                 encryption: Optional[pulumi.Input['FlowEntitlementEncryptionArgs']] = None,
                 entitlement_status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FlowEntitlement resource.
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-description
        :param pulumi.Input[str] flow_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-flowarn
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subscribers: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-subscribers
        :param pulumi.Input[int] data_transfer_subscriber_fee_percent: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-datatransfersubscriberfeepercent
        :param pulumi.Input['FlowEntitlementEncryptionArgs'] encryption: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-encryption
        :param pulumi.Input[str] entitlement_status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-entitlementstatus
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "flow_arn", flow_arn)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "subscribers", subscribers)
        if data_transfer_subscriber_fee_percent is not None:
            pulumi.set(__self__, "data_transfer_subscriber_fee_percent", data_transfer_subscriber_fee_percent)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if entitlement_status is not None:
            pulumi.set(__self__, "entitlement_status", entitlement_status)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="flowArn")
    def flow_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-flowarn
        """
        return pulumi.get(self, "flow_arn")

    @flow_arn.setter
    def flow_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "flow_arn", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def subscribers(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-subscribers
        """
        return pulumi.get(self, "subscribers")

    @subscribers.setter
    def subscribers(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subscribers", value)

    @property
    @pulumi.getter(name="dataTransferSubscriberFeePercent")
    def data_transfer_subscriber_fee_percent(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-datatransfersubscriberfeepercent
        """
        return pulumi.get(self, "data_transfer_subscriber_fee_percent")

    @data_transfer_subscriber_fee_percent.setter
    def data_transfer_subscriber_fee_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "data_transfer_subscriber_fee_percent", value)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input['FlowEntitlementEncryptionArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-encryption
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input['FlowEntitlementEncryptionArgs']]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter(name="entitlementStatus")
    def entitlement_status(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-entitlementstatus
        """
        return pulumi.get(self, "entitlement_status")

    @entitlement_status.setter
    def entitlement_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entitlement_status", value)


class FlowEntitlement(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_transfer_subscriber_fee_percent: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[pulumi.InputType['FlowEntitlementEncryptionArgs']]] = None,
                 entitlement_status: Optional[pulumi.Input[str]] = None,
                 flow_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subscribers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] data_transfer_subscriber_fee_percent: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-datatransfersubscriberfeepercent
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-description
        :param pulumi.Input[pulumi.InputType['FlowEntitlementEncryptionArgs']] encryption: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-encryption
        :param pulumi.Input[str] entitlement_status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-entitlementstatus
        :param pulumi.Input[str] flow_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-flowarn
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subscribers: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-subscribers
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlowEntitlementArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html

        :param str resource_name: The name of the resource.
        :param FlowEntitlementArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlowEntitlementArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_transfer_subscriber_fee_percent: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[pulumi.InputType['FlowEntitlementEncryptionArgs']]] = None,
                 entitlement_status: Optional[pulumi.Input[str]] = None,
                 flow_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subscribers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = FlowEntitlementArgs.__new__(FlowEntitlementArgs)

            __props__.__dict__["data_transfer_subscriber_fee_percent"] = data_transfer_subscriber_fee_percent
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["encryption"] = encryption
            __props__.__dict__["entitlement_status"] = entitlement_status
            if flow_arn is None and not opts.urn:
                raise TypeError("Missing required property 'flow_arn'")
            __props__.__dict__["flow_arn"] = flow_arn
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if subscribers is None and not opts.urn:
                raise TypeError("Missing required property 'subscribers'")
            __props__.__dict__["subscribers"] = subscribers
            __props__.__dict__["entitlement_arn"] = None
        super(FlowEntitlement, __self__).__init__(
            'aws-native:MediaConnect:FlowEntitlement',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FlowEntitlement':
        """
        Get an existing FlowEntitlement resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FlowEntitlementArgs.__new__(FlowEntitlementArgs)

        __props__.__dict__["data_transfer_subscriber_fee_percent"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["encryption"] = None
        __props__.__dict__["entitlement_arn"] = None
        __props__.__dict__["entitlement_status"] = None
        __props__.__dict__["flow_arn"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["subscribers"] = None
        return FlowEntitlement(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataTransferSubscriberFeePercent")
    def data_transfer_subscriber_fee_percent(self) -> pulumi.Output[Optional[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-datatransfersubscriberfeepercent
        """
        return pulumi.get(self, "data_transfer_subscriber_fee_percent")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output[Optional['outputs.FlowEntitlementEncryption']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-encryption
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="entitlementArn")
    def entitlement_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "entitlement_arn")

    @property
    @pulumi.getter(name="entitlementStatus")
    def entitlement_status(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-entitlementstatus
        """
        return pulumi.get(self, "entitlement_status")

    @property
    @pulumi.getter(name="flowArn")
    def flow_arn(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-flowarn
        """
        return pulumi.get(self, "flow_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def subscribers(self) -> pulumi.Output[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-subscribers
        """
        return pulumi.get(self, "subscribers")

