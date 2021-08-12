# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GeofenceCollectionArgs', 'GeofenceCollection']

@pulumi.input_type
class GeofenceCollectionArgs:
    def __init__(__self__, *,
                 collection_name: pulumi.Input[str],
                 pricing_plan: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 pricing_plan_data_source: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GeofenceCollection resource.
        :param pulumi.Input[str] collection_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-collectionname
        :param pulumi.Input[str] pricing_plan: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplan
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-description
        :param pulumi.Input[str] kms_key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-kmskeyid
        :param pulumi.Input[str] pricing_plan_data_source: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplandatasource
        """
        pulumi.set(__self__, "collection_name", collection_name)
        pulumi.set(__self__, "pricing_plan", pricing_plan)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if pricing_plan_data_source is not None:
            pulumi.set(__self__, "pricing_plan_data_source", pricing_plan_data_source)

    @property
    @pulumi.getter(name="collectionName")
    def collection_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-collectionname
        """
        return pulumi.get(self, "collection_name")

    @collection_name.setter
    def collection_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "collection_name", value)

    @property
    @pulumi.getter(name="pricingPlan")
    def pricing_plan(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplan
        """
        return pulumi.get(self, "pricing_plan")

    @pricing_plan.setter
    def pricing_plan(self, value: pulumi.Input[str]):
        pulumi.set(self, "pricing_plan", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-kmskeyid
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="pricingPlanDataSource")
    def pricing_plan_data_source(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplandatasource
        """
        return pulumi.get(self, "pricing_plan_data_source")

    @pricing_plan_data_source.setter
    def pricing_plan_data_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pricing_plan_data_source", value)


class GeofenceCollection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 pricing_plan: Optional[pulumi.Input[str]] = None,
                 pricing_plan_data_source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collection_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-collectionname
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-description
        :param pulumi.Input[str] kms_key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-kmskeyid
        :param pulumi.Input[str] pricing_plan: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplan
        :param pulumi.Input[str] pricing_plan_data_source: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplandatasource
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GeofenceCollectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html

        :param str resource_name: The name of the resource.
        :param GeofenceCollectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GeofenceCollectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 pricing_plan: Optional[pulumi.Input[str]] = None,
                 pricing_plan_data_source: Optional[pulumi.Input[str]] = None,
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
            __props__ = GeofenceCollectionArgs.__new__(GeofenceCollectionArgs)

            if collection_name is None and not opts.urn:
                raise TypeError("Missing required property 'collection_name'")
            __props__.__dict__["collection_name"] = collection_name
            __props__.__dict__["description"] = description
            __props__.__dict__["kms_key_id"] = kms_key_id
            if pricing_plan is None and not opts.urn:
                raise TypeError("Missing required property 'pricing_plan'")
            __props__.__dict__["pricing_plan"] = pricing_plan
            __props__.__dict__["pricing_plan_data_source"] = pricing_plan_data_source
            __props__.__dict__["collection_arn"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        super(GeofenceCollection, __self__).__init__(
            'aws-native:Location:GeofenceCollection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GeofenceCollection':
        """
        Get an existing GeofenceCollection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GeofenceCollectionArgs.__new__(GeofenceCollectionArgs)

        __props__.__dict__["collection_arn"] = None
        __props__.__dict__["collection_name"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["kms_key_id"] = None
        __props__.__dict__["pricing_plan"] = None
        __props__.__dict__["pricing_plan_data_source"] = None
        __props__.__dict__["update_time"] = None
        return GeofenceCollection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="collectionArn")
    def collection_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "collection_arn")

    @property
    @pulumi.getter(name="collectionName")
    def collection_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-collectionname
        """
        return pulumi.get(self, "collection_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-kmskeyid
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="pricingPlan")
    def pricing_plan(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplan
        """
        return pulumi.get(self, "pricing_plan")

    @property
    @pulumi.getter(name="pricingPlanDataSource")
    def pricing_plan_data_source(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplandatasource
        """
        return pulumi.get(self, "pricing_plan_data_source")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "update_time")

