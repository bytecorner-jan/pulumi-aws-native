# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DeploymentArgs', 'Deployment']

@pulumi.input_type
class DeploymentArgs:
    def __init__(__self__, *,
                 rest_api_id: pulumi.Input[str],
                 deployment_canary_settings: Optional[pulumi.Input['DeploymentCanarySettingsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 stage_description: Optional[pulumi.Input['DeploymentStageDescriptionArgs']] = None,
                 stage_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Deployment resource.
        """
        pulumi.set(__self__, "rest_api_id", rest_api_id)
        if deployment_canary_settings is not None:
            pulumi.set(__self__, "deployment_canary_settings", deployment_canary_settings)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if stage_description is not None:
            pulumi.set(__self__, "stage_description", stage_description)
        if stage_name is not None:
            pulumi.set(__self__, "stage_name", stage_name)

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rest_api_id")

    @rest_api_id.setter
    def rest_api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "rest_api_id", value)

    @property
    @pulumi.getter(name="deploymentCanarySettings")
    def deployment_canary_settings(self) -> Optional[pulumi.Input['DeploymentCanarySettingsArgs']]:
        return pulumi.get(self, "deployment_canary_settings")

    @deployment_canary_settings.setter
    def deployment_canary_settings(self, value: Optional[pulumi.Input['DeploymentCanarySettingsArgs']]):
        pulumi.set(self, "deployment_canary_settings", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="stageDescription")
    def stage_description(self) -> Optional[pulumi.Input['DeploymentStageDescriptionArgs']]:
        return pulumi.get(self, "stage_description")

    @stage_description.setter
    def stage_description(self, value: Optional[pulumi.Input['DeploymentStageDescriptionArgs']]):
        pulumi.set(self, "stage_description", value)

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "stage_name")

    @stage_name.setter
    def stage_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stage_name", value)


warnings.warn("""Deployment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Deployment(pulumi.CustomResource):
    warnings.warn("""Deployment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_canary_settings: Optional[pulumi.Input[pulumi.InputType['DeploymentCanarySettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
                 stage_description: Optional[pulumi.Input[pulumi.InputType['DeploymentStageDescriptionArgs']]] = None,
                 stage_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::ApiGateway::Deployment

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeploymentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::ApiGateway::Deployment

        :param str resource_name: The name of the resource.
        :param DeploymentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeploymentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_canary_settings: Optional[pulumi.Input[pulumi.InputType['DeploymentCanarySettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
                 stage_description: Optional[pulumi.Input[pulumi.InputType['DeploymentStageDescriptionArgs']]] = None,
                 stage_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""Deployment is deprecated: Deployment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeploymentArgs.__new__(DeploymentArgs)

            __props__.__dict__["deployment_canary_settings"] = deployment_canary_settings
            __props__.__dict__["description"] = description
            if rest_api_id is None and not opts.urn:
                raise TypeError("Missing required property 'rest_api_id'")
            __props__.__dict__["rest_api_id"] = rest_api_id
            __props__.__dict__["stage_description"] = stage_description
            __props__.__dict__["stage_name"] = stage_name
        super(Deployment, __self__).__init__(
            'aws-native:apigateway:Deployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Deployment':
        """
        Get an existing Deployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeploymentArgs.__new__(DeploymentArgs)

        __props__.__dict__["deployment_canary_settings"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["rest_api_id"] = None
        __props__.__dict__["stage_description"] = None
        __props__.__dict__["stage_name"] = None
        return Deployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deploymentCanarySettings")
    def deployment_canary_settings(self) -> pulumi.Output[Optional['outputs.DeploymentCanarySettings']]:
        return pulumi.get(self, "deployment_canary_settings")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "rest_api_id")

    @property
    @pulumi.getter(name="stageDescription")
    def stage_description(self) -> pulumi.Output[Optional['outputs.DeploymentStageDescription']]:
        return pulumi.get(self, "stage_description")

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "stage_name")

