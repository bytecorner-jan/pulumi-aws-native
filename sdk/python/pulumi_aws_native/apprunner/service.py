# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 source_configuration: pulumi.Input['ServiceSourceConfigurationArgs'],
                 auto_scaling_configuration_arn: Optional[pulumi.Input[str]] = None,
                 encryption_configuration: Optional[pulumi.Input['ServiceEncryptionConfigurationArgs']] = None,
                 health_check_configuration: Optional[pulumi.Input['ServiceHealthCheckConfigurationArgs']] = None,
                 instance_configuration: Optional[pulumi.Input['ServiceInstanceConfigurationArgs']] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Service resource.
        :param pulumi.Input['ServiceSourceConfigurationArgs'] source_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-sourceconfiguration
        :param pulumi.Input[str] auto_scaling_configuration_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-autoscalingconfigurationarn
        :param pulumi.Input['ServiceEncryptionConfigurationArgs'] encryption_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-encryptionconfiguration
        :param pulumi.Input['ServiceHealthCheckConfigurationArgs'] health_check_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-healthcheckconfiguration
        :param pulumi.Input['ServiceInstanceConfigurationArgs'] instance_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-instanceconfiguration
        :param pulumi.Input[str] service_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-servicename
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-tags
        """
        pulumi.set(__self__, "source_configuration", source_configuration)
        if auto_scaling_configuration_arn is not None:
            pulumi.set(__self__, "auto_scaling_configuration_arn", auto_scaling_configuration_arn)
        if encryption_configuration is not None:
            pulumi.set(__self__, "encryption_configuration", encryption_configuration)
        if health_check_configuration is not None:
            pulumi.set(__self__, "health_check_configuration", health_check_configuration)
        if instance_configuration is not None:
            pulumi.set(__self__, "instance_configuration", instance_configuration)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="sourceConfiguration")
    def source_configuration(self) -> pulumi.Input['ServiceSourceConfigurationArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-sourceconfiguration
        """
        return pulumi.get(self, "source_configuration")

    @source_configuration.setter
    def source_configuration(self, value: pulumi.Input['ServiceSourceConfigurationArgs']):
        pulumi.set(self, "source_configuration", value)

    @property
    @pulumi.getter(name="autoScalingConfigurationArn")
    def auto_scaling_configuration_arn(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-autoscalingconfigurationarn
        """
        return pulumi.get(self, "auto_scaling_configuration_arn")

    @auto_scaling_configuration_arn.setter
    def auto_scaling_configuration_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_scaling_configuration_arn", value)

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> Optional[pulumi.Input['ServiceEncryptionConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-encryptionconfiguration
        """
        return pulumi.get(self, "encryption_configuration")

    @encryption_configuration.setter
    def encryption_configuration(self, value: Optional[pulumi.Input['ServiceEncryptionConfigurationArgs']]):
        pulumi.set(self, "encryption_configuration", value)

    @property
    @pulumi.getter(name="healthCheckConfiguration")
    def health_check_configuration(self) -> Optional[pulumi.Input['ServiceHealthCheckConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-healthcheckconfiguration
        """
        return pulumi.get(self, "health_check_configuration")

    @health_check_configuration.setter
    def health_check_configuration(self, value: Optional[pulumi.Input['ServiceHealthCheckConfigurationArgs']]):
        pulumi.set(self, "health_check_configuration", value)

    @property
    @pulumi.getter(name="instanceConfiguration")
    def instance_configuration(self) -> Optional[pulumi.Input['ServiceInstanceConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-instanceconfiguration
        """
        return pulumi.get(self, "instance_configuration")

    @instance_configuration.setter
    def instance_configuration(self, value: Optional[pulumi.Input['ServiceInstanceConfigurationArgs']]):
        pulumi.set(self, "instance_configuration", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-servicename
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class Service(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_configuration_arn: Optional[pulumi.Input[str]] = None,
                 encryption_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceEncryptionConfigurationArgs']]] = None,
                 health_check_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceHealthCheckConfigurationArgs']]] = None,
                 instance_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceInstanceConfigurationArgs']]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 source_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceSourceConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_configuration_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-autoscalingconfigurationarn
        :param pulumi.Input[pulumi.InputType['ServiceEncryptionConfigurationArgs']] encryption_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-encryptionconfiguration
        :param pulumi.Input[pulumi.InputType['ServiceHealthCheckConfigurationArgs']] health_check_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-healthcheckconfiguration
        :param pulumi.Input[pulumi.InputType['ServiceInstanceConfigurationArgs']] instance_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-instanceconfiguration
        :param pulumi.Input[str] service_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-servicename
        :param pulumi.Input[pulumi.InputType['ServiceSourceConfigurationArgs']] source_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-sourceconfiguration
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html

        :param str resource_name: The name of the resource.
        :param ServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_configuration_arn: Optional[pulumi.Input[str]] = None,
                 encryption_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceEncryptionConfigurationArgs']]] = None,
                 health_check_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceHealthCheckConfigurationArgs']]] = None,
                 instance_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceInstanceConfigurationArgs']]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 source_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceSourceConfigurationArgs']]] = None,
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
            __props__ = ServiceArgs.__new__(ServiceArgs)

            __props__.__dict__["auto_scaling_configuration_arn"] = auto_scaling_configuration_arn
            __props__.__dict__["encryption_configuration"] = encryption_configuration
            __props__.__dict__["health_check_configuration"] = health_check_configuration
            __props__.__dict__["instance_configuration"] = instance_configuration
            __props__.__dict__["service_name"] = service_name
            if source_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'source_configuration'")
            __props__.__dict__["source_configuration"] = source_configuration
            __props__.__dict__["tags"] = tags
            __props__.__dict__["service_arn"] = None
            __props__.__dict__["service_id"] = None
            __props__.__dict__["service_url"] = None
            __props__.__dict__["status"] = None
        super(Service, __self__).__init__(
            'aws-native:AppRunner:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceArgs.__new__(ServiceArgs)

        __props__.__dict__["auto_scaling_configuration_arn"] = None
        __props__.__dict__["encryption_configuration"] = None
        __props__.__dict__["health_check_configuration"] = None
        __props__.__dict__["instance_configuration"] = None
        __props__.__dict__["service_arn"] = None
        __props__.__dict__["service_id"] = None
        __props__.__dict__["service_name"] = None
        __props__.__dict__["service_url"] = None
        __props__.__dict__["source_configuration"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScalingConfigurationArn")
    def auto_scaling_configuration_arn(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-autoscalingconfigurationarn
        """
        return pulumi.get(self, "auto_scaling_configuration_arn")

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> pulumi.Output[Optional['outputs.ServiceEncryptionConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-encryptionconfiguration
        """
        return pulumi.get(self, "encryption_configuration")

    @property
    @pulumi.getter(name="healthCheckConfiguration")
    def health_check_configuration(self) -> pulumi.Output[Optional['outputs.ServiceHealthCheckConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-healthcheckconfiguration
        """
        return pulumi.get(self, "health_check_configuration")

    @property
    @pulumi.getter(name="instanceConfiguration")
    def instance_configuration(self) -> pulumi.Output[Optional['outputs.ServiceInstanceConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-instanceconfiguration
        """
        return pulumi.get(self, "instance_configuration")

    @property
    @pulumi.getter(name="serviceArn")
    def service_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_arn")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-servicename
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceUrl")
    def service_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_url")

    @property
    @pulumi.getter(name="sourceConfiguration")
    def source_configuration(self) -> pulumi.Output['outputs.ServiceSourceConfiguration']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-sourceconfiguration
        """
        return pulumi.get(self, "source_configuration")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-service.html#cfn-apprunner-service-tags
        """
        return pulumi.get(self, "tags")

