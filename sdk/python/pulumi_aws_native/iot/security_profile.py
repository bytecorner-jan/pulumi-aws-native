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

__all__ = ['SecurityProfileArgs', 'SecurityProfile']

@pulumi.input_type
class SecurityProfileArgs:
    def __init__(__self__, *,
                 additional_metrics_to_retain_v2: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityProfileMetricToRetainArgs']]]] = None,
                 alert_targets: Optional[pulumi.Input[Mapping[str, pulumi.Input['SecurityProfileAlertTargetArgs']]]] = None,
                 behaviors: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityProfileBehaviorArgs']]]] = None,
                 security_profile_description: Optional[pulumi.Input[str]] = None,
                 security_profile_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 target_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SecurityProfile resource.
        :param pulumi.Input[Sequence[pulumi.Input['SecurityProfileMetricToRetainArgs']]] additional_metrics_to_retain_v2: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-additionalmetricstoretainv2
        :param pulumi.Input[Mapping[str, pulumi.Input['SecurityProfileAlertTargetArgs']]] alert_targets: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-alerttargets
        :param pulumi.Input[Sequence[pulumi.Input['SecurityProfileBehaviorArgs']]] behaviors: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-behaviors
        :param pulumi.Input[str] security_profile_description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-securityprofiledescription
        :param pulumi.Input[str] security_profile_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-securityprofilename
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-tags
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_arns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-targetarns
        """
        if additional_metrics_to_retain_v2 is not None:
            pulumi.set(__self__, "additional_metrics_to_retain_v2", additional_metrics_to_retain_v2)
        if alert_targets is not None:
            pulumi.set(__self__, "alert_targets", alert_targets)
        if behaviors is not None:
            pulumi.set(__self__, "behaviors", behaviors)
        if security_profile_description is not None:
            pulumi.set(__self__, "security_profile_description", security_profile_description)
        if security_profile_name is not None:
            pulumi.set(__self__, "security_profile_name", security_profile_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_arns is not None:
            pulumi.set(__self__, "target_arns", target_arns)

    @property
    @pulumi.getter(name="additionalMetricsToRetainV2")
    def additional_metrics_to_retain_v2(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecurityProfileMetricToRetainArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-additionalmetricstoretainv2
        """
        return pulumi.get(self, "additional_metrics_to_retain_v2")

    @additional_metrics_to_retain_v2.setter
    def additional_metrics_to_retain_v2(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityProfileMetricToRetainArgs']]]]):
        pulumi.set(self, "additional_metrics_to_retain_v2", value)

    @property
    @pulumi.getter(name="alertTargets")
    def alert_targets(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['SecurityProfileAlertTargetArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-alerttargets
        """
        return pulumi.get(self, "alert_targets")

    @alert_targets.setter
    def alert_targets(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['SecurityProfileAlertTargetArgs']]]]):
        pulumi.set(self, "alert_targets", value)

    @property
    @pulumi.getter
    def behaviors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecurityProfileBehaviorArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-behaviors
        """
        return pulumi.get(self, "behaviors")

    @behaviors.setter
    def behaviors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityProfileBehaviorArgs']]]]):
        pulumi.set(self, "behaviors", value)

    @property
    @pulumi.getter(name="securityProfileDescription")
    def security_profile_description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-securityprofiledescription
        """
        return pulumi.get(self, "security_profile_description")

    @security_profile_description.setter
    def security_profile_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_profile_description", value)

    @property
    @pulumi.getter(name="securityProfileName")
    def security_profile_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-securityprofilename
        """
        return pulumi.get(self, "security_profile_name")

    @security_profile_name.setter
    def security_profile_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_profile_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="targetArns")
    def target_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-targetarns
        """
        return pulumi.get(self, "target_arns")

    @target_arns.setter
    def target_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_arns", value)


class SecurityProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_metrics_to_retain_v2: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityProfileMetricToRetainArgs']]]]] = None,
                 alert_targets: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SecurityProfileAlertTargetArgs']]]]] = None,
                 behaviors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityProfileBehaviorArgs']]]]] = None,
                 security_profile_description: Optional[pulumi.Input[str]] = None,
                 security_profile_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 target_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityProfileMetricToRetainArgs']]]] additional_metrics_to_retain_v2: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-additionalmetricstoretainv2
        :param pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SecurityProfileAlertTargetArgs']]]] alert_targets: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-alerttargets
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityProfileBehaviorArgs']]]] behaviors: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-behaviors
        :param pulumi.Input[str] security_profile_description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-securityprofiledescription
        :param pulumi.Input[str] security_profile_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-securityprofilename
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-tags
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_arns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-targetarns
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecurityProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html

        :param str resource_name: The name of the resource.
        :param SecurityProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_metrics_to_retain_v2: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityProfileMetricToRetainArgs']]]]] = None,
                 alert_targets: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SecurityProfileAlertTargetArgs']]]]] = None,
                 behaviors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityProfileBehaviorArgs']]]]] = None,
                 security_profile_description: Optional[pulumi.Input[str]] = None,
                 security_profile_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 target_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = SecurityProfileArgs.__new__(SecurityProfileArgs)

            __props__.__dict__["additional_metrics_to_retain_v2"] = additional_metrics_to_retain_v2
            __props__.__dict__["alert_targets"] = alert_targets
            __props__.__dict__["behaviors"] = behaviors
            __props__.__dict__["security_profile_description"] = security_profile_description
            __props__.__dict__["security_profile_name"] = security_profile_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["target_arns"] = target_arns
            __props__.__dict__["security_profile_arn"] = None
        super(SecurityProfile, __self__).__init__(
            'aws-native:IoT:SecurityProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SecurityProfile':
        """
        Get an existing SecurityProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SecurityProfileArgs.__new__(SecurityProfileArgs)

        __props__.__dict__["additional_metrics_to_retain_v2"] = None
        __props__.__dict__["alert_targets"] = None
        __props__.__dict__["behaviors"] = None
        __props__.__dict__["security_profile_arn"] = None
        __props__.__dict__["security_profile_description"] = None
        __props__.__dict__["security_profile_name"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["target_arns"] = None
        return SecurityProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalMetricsToRetainV2")
    def additional_metrics_to_retain_v2(self) -> pulumi.Output[Optional[Sequence['outputs.SecurityProfileMetricToRetain']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-additionalmetricstoretainv2
        """
        return pulumi.get(self, "additional_metrics_to_retain_v2")

    @property
    @pulumi.getter(name="alertTargets")
    def alert_targets(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.SecurityProfileAlertTarget']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-alerttargets
        """
        return pulumi.get(self, "alert_targets")

    @property
    @pulumi.getter
    def behaviors(self) -> pulumi.Output[Optional[Sequence['outputs.SecurityProfileBehavior']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-behaviors
        """
        return pulumi.get(self, "behaviors")

    @property
    @pulumi.getter(name="securityProfileArn")
    def security_profile_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "security_profile_arn")

    @property
    @pulumi.getter(name="securityProfileDescription")
    def security_profile_description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-securityprofiledescription
        """
        return pulumi.get(self, "security_profile_description")

    @property
    @pulumi.getter(name="securityProfileName")
    def security_profile_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-securityprofilename
        """
        return pulumi.get(self, "security_profile_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetArns")
    def target_arns(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html#cfn-iot-securityprofile-targetarns
        """
        return pulumi.get(self, "target_arns")

