# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['LaunchProfileArgs', 'LaunchProfile']

@pulumi.input_type
class LaunchProfileArgs:
    def __init__(__self__, *,
                 ec2_subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 launch_profile_protocol_versions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: pulumi.Input[str],
                 stream_configuration: pulumi.Input['LaunchProfileStreamConfigurationArgs'],
                 studio_component_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 studio_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input['LaunchProfileTagsArgs']] = None):
        """
        The set of arguments for constructing a LaunchProfile resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ec2_subnet_ids: <p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from. These subnets must support the specified instance types. </p>
        :param pulumi.Input[Sequence[pulumi.Input[str]]] launch_profile_protocol_versions: <p>The version number of the protocol that is used by the launch profile. The only valid version is "2021-03-31".</p>
        :param pulumi.Input[str] name: <p>The name for the launch profile.</p>
        :param pulumi.Input[Sequence[pulumi.Input[str]]] studio_component_ids: <p>Unique identifiers for a collection of studio components that can be used with this launch profile.</p>
        :param pulumi.Input[str] studio_id: <p>The studioId. </p>
        :param pulumi.Input[str] description: <p>The description.</p>
        """
        pulumi.set(__self__, "ec2_subnet_ids", ec2_subnet_ids)
        pulumi.set(__self__, "launch_profile_protocol_versions", launch_profile_protocol_versions)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "stream_configuration", stream_configuration)
        pulumi.set(__self__, "studio_component_ids", studio_component_ids)
        pulumi.set(__self__, "studio_id", studio_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="ec2SubnetIds")
    def ec2_subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        <p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from. These subnets must support the specified instance types. </p>
        """
        return pulumi.get(self, "ec2_subnet_ids")

    @ec2_subnet_ids.setter
    def ec2_subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "ec2_subnet_ids", value)

    @property
    @pulumi.getter(name="launchProfileProtocolVersions")
    def launch_profile_protocol_versions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        <p>The version number of the protocol that is used by the launch profile. The only valid version is "2021-03-31".</p>
        """
        return pulumi.get(self, "launch_profile_protocol_versions")

    @launch_profile_protocol_versions.setter
    def launch_profile_protocol_versions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "launch_profile_protocol_versions", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        <p>The name for the launch profile.</p>
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="streamConfiguration")
    def stream_configuration(self) -> pulumi.Input['LaunchProfileStreamConfigurationArgs']:
        return pulumi.get(self, "stream_configuration")

    @stream_configuration.setter
    def stream_configuration(self, value: pulumi.Input['LaunchProfileStreamConfigurationArgs']):
        pulumi.set(self, "stream_configuration", value)

    @property
    @pulumi.getter(name="studioComponentIds")
    def studio_component_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        <p>Unique identifiers for a collection of studio components that can be used with this launch profile.</p>
        """
        return pulumi.get(self, "studio_component_ids")

    @studio_component_ids.setter
    def studio_component_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "studio_component_ids", value)

    @property
    @pulumi.getter(name="studioId")
    def studio_id(self) -> pulumi.Input[str]:
        """
        <p>The studioId. </p>
        """
        return pulumi.get(self, "studio_id")

    @studio_id.setter
    def studio_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "studio_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        <p>The description.</p>
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input['LaunchProfileTagsArgs']]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input['LaunchProfileTagsArgs']]):
        pulumi.set(self, "tags", value)


class LaunchProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ec2_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 launch_profile_protocol_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_configuration: Optional[pulumi.Input[pulumi.InputType['LaunchProfileStreamConfigurationArgs']]] = None,
                 studio_component_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 studio_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['LaunchProfileTagsArgs']]] = None,
                 __props__=None):
        """
        Represents a launch profile which delegates access to a collection of studio components to studio users

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: <p>The description.</p>
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ec2_subnet_ids: <p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from. These subnets must support the specified instance types. </p>
        :param pulumi.Input[Sequence[pulumi.Input[str]]] launch_profile_protocol_versions: <p>The version number of the protocol that is used by the launch profile. The only valid version is "2021-03-31".</p>
        :param pulumi.Input[str] name: <p>The name for the launch profile.</p>
        :param pulumi.Input[Sequence[pulumi.Input[str]]] studio_component_ids: <p>Unique identifiers for a collection of studio components that can be used with this launch profile.</p>
        :param pulumi.Input[str] studio_id: <p>The studioId. </p>
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LaunchProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a launch profile which delegates access to a collection of studio components to studio users

        :param str resource_name: The name of the resource.
        :param LaunchProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LaunchProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ec2_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 launch_profile_protocol_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_configuration: Optional[pulumi.Input[pulumi.InputType['LaunchProfileStreamConfigurationArgs']]] = None,
                 studio_component_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 studio_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['LaunchProfileTagsArgs']]] = None,
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
            __props__ = LaunchProfileArgs.__new__(LaunchProfileArgs)

            __props__.__dict__["description"] = description
            if ec2_subnet_ids is None and not opts.urn:
                raise TypeError("Missing required property 'ec2_subnet_ids'")
            __props__.__dict__["ec2_subnet_ids"] = ec2_subnet_ids
            if launch_profile_protocol_versions is None and not opts.urn:
                raise TypeError("Missing required property 'launch_profile_protocol_versions'")
            __props__.__dict__["launch_profile_protocol_versions"] = launch_profile_protocol_versions
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if stream_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'stream_configuration'")
            __props__.__dict__["stream_configuration"] = stream_configuration
            if studio_component_ids is None and not opts.urn:
                raise TypeError("Missing required property 'studio_component_ids'")
            __props__.__dict__["studio_component_ids"] = studio_component_ids
            if studio_id is None and not opts.urn:
                raise TypeError("Missing required property 'studio_id'")
            __props__.__dict__["studio_id"] = studio_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["launch_profile_id"] = None
        super(LaunchProfile, __self__).__init__(
            'aws-native:nimblestudio:LaunchProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LaunchProfile':
        """
        Get an existing LaunchProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LaunchProfileArgs.__new__(LaunchProfileArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["ec2_subnet_ids"] = None
        __props__.__dict__["launch_profile_id"] = None
        __props__.__dict__["launch_profile_protocol_versions"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["stream_configuration"] = None
        __props__.__dict__["studio_component_ids"] = None
        __props__.__dict__["studio_id"] = None
        __props__.__dict__["tags"] = None
        return LaunchProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        <p>The description.</p>
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ec2SubnetIds")
    def ec2_subnet_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        <p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from. These subnets must support the specified instance types. </p>
        """
        return pulumi.get(self, "ec2_subnet_ids")

    @property
    @pulumi.getter(name="launchProfileId")
    def launch_profile_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "launch_profile_id")

    @property
    @pulumi.getter(name="launchProfileProtocolVersions")
    def launch_profile_protocol_versions(self) -> pulumi.Output[Sequence[str]]:
        """
        <p>The version number of the protocol that is used by the launch profile. The only valid version is "2021-03-31".</p>
        """
        return pulumi.get(self, "launch_profile_protocol_versions")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        <p>The name for the launch profile.</p>
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="streamConfiguration")
    def stream_configuration(self) -> pulumi.Output['outputs.LaunchProfileStreamConfiguration']:
        return pulumi.get(self, "stream_configuration")

    @property
    @pulumi.getter(name="studioComponentIds")
    def studio_component_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        <p>Unique identifiers for a collection of studio components that can be used with this launch profile.</p>
        """
        return pulumi.get(self, "studio_component_ids")

    @property
    @pulumi.getter(name="studioId")
    def studio_id(self) -> pulumi.Output[str]:
        """
        <p>The studioId. </p>
        """
        return pulumi.get(self, "studio_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional['outputs.LaunchProfileTags']]:
        return pulumi.get(self, "tags")

