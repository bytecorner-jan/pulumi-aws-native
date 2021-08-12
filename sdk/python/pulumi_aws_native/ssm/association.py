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

__all__ = ['AssociationArgs', 'Association']

@pulumi.input_type
class AssociationArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 apply_only_at_cron_interval: Optional[pulumi.Input[bool]] = None,
                 association_name: Optional[pulumi.Input[str]] = None,
                 automation_target_parameter_name: Optional[pulumi.Input[str]] = None,
                 calendar_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 compliance_severity: Optional[pulumi.Input[str]] = None,
                 document_version: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_concurrency: Optional[pulumi.Input[str]] = None,
                 max_errors: Optional[pulumi.Input[str]] = None,
                 output_location: Optional[pulumi.Input['AssociationInstanceAssociationOutputLocationArgs']] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[Any, str]]]]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None,
                 sync_compliance: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input['AssociationTargetArgs']]]] = None,
                 wait_for_success_timeout_seconds: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Association resource.
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
        :param pulumi.Input[bool] apply_only_at_cron_interval: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-applyonlyatcroninterval
        :param pulumi.Input[str] association_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-associationname
        :param pulumi.Input[str] automation_target_parameter_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-automationtargetparametername
        :param pulumi.Input[Sequence[pulumi.Input[str]]] calendar_names: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-calendarnames
        :param pulumi.Input[str] compliance_severity: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-complianceseverity
        :param pulumi.Input[str] document_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
        :param pulumi.Input[str] instance_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
        :param pulumi.Input[str] max_concurrency: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxconcurrency
        :param pulumi.Input[str] max_errors: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxerrors
        :param pulumi.Input['AssociationInstanceAssociationOutputLocationArgs'] output_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-outputlocation
        :param pulumi.Input[Mapping[str, pulumi.Input[Union[Any, str]]]] parameters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
        :param pulumi.Input[str] schedule_expression: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
        :param pulumi.Input[str] sync_compliance: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-synccompliance
        :param pulumi.Input[Sequence[pulumi.Input['AssociationTargetArgs']]] targets: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
        :param pulumi.Input[int] wait_for_success_timeout_seconds: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-waitforsuccesstimeoutseconds
        """
        pulumi.set(__self__, "name", name)
        if apply_only_at_cron_interval is not None:
            pulumi.set(__self__, "apply_only_at_cron_interval", apply_only_at_cron_interval)
        if association_name is not None:
            pulumi.set(__self__, "association_name", association_name)
        if automation_target_parameter_name is not None:
            pulumi.set(__self__, "automation_target_parameter_name", automation_target_parameter_name)
        if calendar_names is not None:
            pulumi.set(__self__, "calendar_names", calendar_names)
        if compliance_severity is not None:
            pulumi.set(__self__, "compliance_severity", compliance_severity)
        if document_version is not None:
            pulumi.set(__self__, "document_version", document_version)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if max_concurrency is not None:
            pulumi.set(__self__, "max_concurrency", max_concurrency)
        if max_errors is not None:
            pulumi.set(__self__, "max_errors", max_errors)
        if output_location is not None:
            pulumi.set(__self__, "output_location", output_location)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if schedule_expression is not None:
            pulumi.set(__self__, "schedule_expression", schedule_expression)
        if sync_compliance is not None:
            pulumi.set(__self__, "sync_compliance", sync_compliance)
        if targets is not None:
            pulumi.set(__self__, "targets", targets)
        if wait_for_success_timeout_seconds is not None:
            pulumi.set(__self__, "wait_for_success_timeout_seconds", wait_for_success_timeout_seconds)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="applyOnlyAtCronInterval")
    def apply_only_at_cron_interval(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-applyonlyatcroninterval
        """
        return pulumi.get(self, "apply_only_at_cron_interval")

    @apply_only_at_cron_interval.setter
    def apply_only_at_cron_interval(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "apply_only_at_cron_interval", value)

    @property
    @pulumi.getter(name="associationName")
    def association_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-associationname
        """
        return pulumi.get(self, "association_name")

    @association_name.setter
    def association_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "association_name", value)

    @property
    @pulumi.getter(name="automationTargetParameterName")
    def automation_target_parameter_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-automationtargetparametername
        """
        return pulumi.get(self, "automation_target_parameter_name")

    @automation_target_parameter_name.setter
    def automation_target_parameter_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "automation_target_parameter_name", value)

    @property
    @pulumi.getter(name="calendarNames")
    def calendar_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-calendarnames
        """
        return pulumi.get(self, "calendar_names")

    @calendar_names.setter
    def calendar_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "calendar_names", value)

    @property
    @pulumi.getter(name="complianceSeverity")
    def compliance_severity(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-complianceseverity
        """
        return pulumi.get(self, "compliance_severity")

    @compliance_severity.setter
    def compliance_severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compliance_severity", value)

    @property
    @pulumi.getter(name="documentVersion")
    def document_version(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
        """
        return pulumi.get(self, "document_version")

    @document_version.setter
    def document_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "document_version", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxConcurrency")
    def max_concurrency(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxconcurrency
        """
        return pulumi.get(self, "max_concurrency")

    @max_concurrency.setter
    def max_concurrency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_concurrency", value)

    @property
    @pulumi.getter(name="maxErrors")
    def max_errors(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxerrors
        """
        return pulumi.get(self, "max_errors")

    @max_errors.setter
    def max_errors(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_errors", value)

    @property
    @pulumi.getter(name="outputLocation")
    def output_location(self) -> Optional[pulumi.Input['AssociationInstanceAssociationOutputLocationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-outputlocation
        """
        return pulumi.get(self, "output_location")

    @output_location.setter
    def output_location(self, value: Optional[pulumi.Input['AssociationInstanceAssociationOutputLocationArgs']]):
        pulumi.set(self, "output_location", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[Any, str]]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[Any, str]]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
        """
        return pulumi.get(self, "schedule_expression")

    @schedule_expression.setter
    def schedule_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_expression", value)

    @property
    @pulumi.getter(name="syncCompliance")
    def sync_compliance(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-synccompliance
        """
        return pulumi.get(self, "sync_compliance")

    @sync_compliance.setter
    def sync_compliance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sync_compliance", value)

    @property
    @pulumi.getter
    def targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AssociationTargetArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AssociationTargetArgs']]]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter(name="waitForSuccessTimeoutSeconds")
    def wait_for_success_timeout_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-waitforsuccesstimeoutseconds
        """
        return pulumi.get(self, "wait_for_success_timeout_seconds")

    @wait_for_success_timeout_seconds.setter
    def wait_for_success_timeout_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_for_success_timeout_seconds", value)


class Association(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apply_only_at_cron_interval: Optional[pulumi.Input[bool]] = None,
                 association_name: Optional[pulumi.Input[str]] = None,
                 automation_target_parameter_name: Optional[pulumi.Input[str]] = None,
                 calendar_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 compliance_severity: Optional[pulumi.Input[str]] = None,
                 document_version: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_concurrency: Optional[pulumi.Input[str]] = None,
                 max_errors: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_location: Optional[pulumi.Input[pulumi.InputType['AssociationInstanceAssociationOutputLocationArgs']]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[Any, str]]]]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None,
                 sync_compliance: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AssociationTargetArgs']]]]] = None,
                 wait_for_success_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_only_at_cron_interval: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-applyonlyatcroninterval
        :param pulumi.Input[str] association_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-associationname
        :param pulumi.Input[str] automation_target_parameter_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-automationtargetparametername
        :param pulumi.Input[Sequence[pulumi.Input[str]]] calendar_names: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-calendarnames
        :param pulumi.Input[str] compliance_severity: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-complianceseverity
        :param pulumi.Input[str] document_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
        :param pulumi.Input[str] instance_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
        :param pulumi.Input[str] max_concurrency: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxconcurrency
        :param pulumi.Input[str] max_errors: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxerrors
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
        :param pulumi.Input[pulumi.InputType['AssociationInstanceAssociationOutputLocationArgs']] output_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-outputlocation
        :param pulumi.Input[Mapping[str, pulumi.Input[Union[Any, str]]]] parameters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
        :param pulumi.Input[str] schedule_expression: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
        :param pulumi.Input[str] sync_compliance: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-synccompliance
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AssociationTargetArgs']]]] targets: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
        :param pulumi.Input[int] wait_for_success_timeout_seconds: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-waitforsuccesstimeoutseconds
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html

        :param str resource_name: The name of the resource.
        :param AssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apply_only_at_cron_interval: Optional[pulumi.Input[bool]] = None,
                 association_name: Optional[pulumi.Input[str]] = None,
                 automation_target_parameter_name: Optional[pulumi.Input[str]] = None,
                 calendar_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 compliance_severity: Optional[pulumi.Input[str]] = None,
                 document_version: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_concurrency: Optional[pulumi.Input[str]] = None,
                 max_errors: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_location: Optional[pulumi.Input[pulumi.InputType['AssociationInstanceAssociationOutputLocationArgs']]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[Any, str]]]]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None,
                 sync_compliance: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AssociationTargetArgs']]]]] = None,
                 wait_for_success_timeout_seconds: Optional[pulumi.Input[int]] = None,
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
            __props__ = AssociationArgs.__new__(AssociationArgs)

            __props__.__dict__["apply_only_at_cron_interval"] = apply_only_at_cron_interval
            __props__.__dict__["association_name"] = association_name
            __props__.__dict__["automation_target_parameter_name"] = automation_target_parameter_name
            __props__.__dict__["calendar_names"] = calendar_names
            __props__.__dict__["compliance_severity"] = compliance_severity
            __props__.__dict__["document_version"] = document_version
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["max_concurrency"] = max_concurrency
            __props__.__dict__["max_errors"] = max_errors
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["output_location"] = output_location
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["schedule_expression"] = schedule_expression
            __props__.__dict__["sync_compliance"] = sync_compliance
            __props__.__dict__["targets"] = targets
            __props__.__dict__["wait_for_success_timeout_seconds"] = wait_for_success_timeout_seconds
            __props__.__dict__["association_id"] = None
        super(Association, __self__).__init__(
            'aws-native:SSM:Association',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Association':
        """
        Get an existing Association resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AssociationArgs.__new__(AssociationArgs)

        __props__.__dict__["apply_only_at_cron_interval"] = None
        __props__.__dict__["association_id"] = None
        __props__.__dict__["association_name"] = None
        __props__.__dict__["automation_target_parameter_name"] = None
        __props__.__dict__["calendar_names"] = None
        __props__.__dict__["compliance_severity"] = None
        __props__.__dict__["document_version"] = None
        __props__.__dict__["instance_id"] = None
        __props__.__dict__["max_concurrency"] = None
        __props__.__dict__["max_errors"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["output_location"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["schedule_expression"] = None
        __props__.__dict__["sync_compliance"] = None
        __props__.__dict__["targets"] = None
        __props__.__dict__["wait_for_success_timeout_seconds"] = None
        return Association(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applyOnlyAtCronInterval")
    def apply_only_at_cron_interval(self) -> pulumi.Output[Optional[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-applyonlyatcroninterval
        """
        return pulumi.get(self, "apply_only_at_cron_interval")

    @property
    @pulumi.getter(name="associationId")
    def association_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "association_id")

    @property
    @pulumi.getter(name="associationName")
    def association_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-associationname
        """
        return pulumi.get(self, "association_name")

    @property
    @pulumi.getter(name="automationTargetParameterName")
    def automation_target_parameter_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-automationtargetparametername
        """
        return pulumi.get(self, "automation_target_parameter_name")

    @property
    @pulumi.getter(name="calendarNames")
    def calendar_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-calendarnames
        """
        return pulumi.get(self, "calendar_names")

    @property
    @pulumi.getter(name="complianceSeverity")
    def compliance_severity(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-complianceseverity
        """
        return pulumi.get(self, "compliance_severity")

    @property
    @pulumi.getter(name="documentVersion")
    def document_version(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
        """
        return pulumi.get(self, "document_version")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="maxConcurrency")
    def max_concurrency(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxconcurrency
        """
        return pulumi.get(self, "max_concurrency")

    @property
    @pulumi.getter(name="maxErrors")
    def max_errors(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxerrors
        """
        return pulumi.get(self, "max_errors")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputLocation")
    def output_location(self) -> pulumi.Output[Optional['outputs.AssociationInstanceAssociationOutputLocation']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-outputlocation
        """
        return pulumi.get(self, "output_location")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
        """
        return pulumi.get(self, "schedule_expression")

    @property
    @pulumi.getter(name="syncCompliance")
    def sync_compliance(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-synccompliance
        """
        return pulumi.get(self, "sync_compliance")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[Optional[Sequence['outputs.AssociationTarget']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="waitForSuccessTimeoutSeconds")
    def wait_for_success_timeout_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-waitforsuccesstimeoutseconds
        """
        return pulumi.get(self, "wait_for_success_timeout_seconds")

