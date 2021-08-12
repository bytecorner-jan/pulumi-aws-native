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

__all__ = ['DetectorArgs', 'Detector']

@pulumi.input_type
class DetectorArgs:
    def __init__(__self__, *,
                 detector_id: pulumi.Input[str],
                 event_type: pulumi.Input['DetectorEventTypeArgs'],
                 rules: pulumi.Input[Sequence[pulumi.Input['DetectorRuleArgs']]],
                 associated_models: Optional[pulumi.Input[Sequence[pulumi.Input['DetectorModelArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detector_version_status: Optional[pulumi.Input[str]] = None,
                 rule_execution_mode: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Detector resource.
        :param pulumi.Input[str] detector_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorid
        :param pulumi.Input['DetectorEventTypeArgs'] event_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-eventtype
        :param pulumi.Input[Sequence[pulumi.Input['DetectorRuleArgs']]] rules: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-rules
        :param pulumi.Input[Sequence[pulumi.Input['DetectorModelArgs']]] associated_models: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-associatedmodels
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-description
        :param pulumi.Input[str] detector_version_status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorversionstatus
        :param pulumi.Input[str] rule_execution_mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-ruleexecutionmode
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-tags
        """
        pulumi.set(__self__, "detector_id", detector_id)
        pulumi.set(__self__, "event_type", event_type)
        pulumi.set(__self__, "rules", rules)
        if associated_models is not None:
            pulumi.set(__self__, "associated_models", associated_models)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if detector_version_status is not None:
            pulumi.set(__self__, "detector_version_status", detector_version_status)
        if rule_execution_mode is not None:
            pulumi.set(__self__, "rule_execution_mode", rule_execution_mode)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorid
        """
        return pulumi.get(self, "detector_id")

    @detector_id.setter
    def detector_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "detector_id", value)

    @property
    @pulumi.getter(name="eventType")
    def event_type(self) -> pulumi.Input['DetectorEventTypeArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-eventtype
        """
        return pulumi.get(self, "event_type")

    @event_type.setter
    def event_type(self, value: pulumi.Input['DetectorEventTypeArgs']):
        pulumi.set(self, "event_type", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['DetectorRuleArgs']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-rules
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['DetectorRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="associatedModels")
    def associated_models(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DetectorModelArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-associatedmodels
        """
        return pulumi.get(self, "associated_models")

    @associated_models.setter
    def associated_models(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DetectorModelArgs']]]]):
        pulumi.set(self, "associated_models", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="detectorVersionStatus")
    def detector_version_status(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorversionstatus
        """
        return pulumi.get(self, "detector_version_status")

    @detector_version_status.setter
    def detector_version_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "detector_version_status", value)

    @property
    @pulumi.getter(name="ruleExecutionMode")
    def rule_execution_mode(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-ruleexecutionmode
        """
        return pulumi.get(self, "rule_execution_mode")

    @rule_execution_mode.setter
    def rule_execution_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_execution_mode", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class Detector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associated_models: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorModelArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 detector_version_status: Optional[pulumi.Input[str]] = None,
                 event_type: Optional[pulumi.Input[pulumi.InputType['DetectorEventTypeArgs']]] = None,
                 rule_execution_mode: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorRuleArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorModelArgs']]]] associated_models: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-associatedmodels
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-description
        :param pulumi.Input[str] detector_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorid
        :param pulumi.Input[str] detector_version_status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorversionstatus
        :param pulumi.Input[pulumi.InputType['DetectorEventTypeArgs']] event_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-eventtype
        :param pulumi.Input[str] rule_execution_mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-ruleexecutionmode
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorRuleArgs']]]] rules: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-rules
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DetectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html

        :param str resource_name: The name of the resource.
        :param DetectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DetectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associated_models: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorModelArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 detector_version_status: Optional[pulumi.Input[str]] = None,
                 event_type: Optional[pulumi.Input[pulumi.InputType['DetectorEventTypeArgs']]] = None,
                 rule_execution_mode: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorRuleArgs']]]]] = None,
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
            __props__ = DetectorArgs.__new__(DetectorArgs)

            __props__.__dict__["associated_models"] = associated_models
            __props__.__dict__["description"] = description
            if detector_id is None and not opts.urn:
                raise TypeError("Missing required property 'detector_id'")
            __props__.__dict__["detector_id"] = detector_id
            __props__.__dict__["detector_version_status"] = detector_version_status
            if event_type is None and not opts.urn:
                raise TypeError("Missing required property 'event_type'")
            __props__.__dict__["event_type"] = event_type
            __props__.__dict__["rule_execution_mode"] = rule_execution_mode
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["detector_version_id"] = None
            __props__.__dict__["event_type_arn"] = None
            __props__.__dict__["event_type_created_time"] = None
            __props__.__dict__["event_type_last_updated_time"] = None
            __props__.__dict__["last_updated_time"] = None
        super(Detector, __self__).__init__(
            'aws-native:FraudDetector:Detector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Detector':
        """
        Get an existing Detector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DetectorArgs.__new__(DetectorArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["associated_models"] = None
        __props__.__dict__["created_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["detector_id"] = None
        __props__.__dict__["detector_version_id"] = None
        __props__.__dict__["detector_version_status"] = None
        __props__.__dict__["event_type"] = None
        __props__.__dict__["event_type_arn"] = None
        __props__.__dict__["event_type_created_time"] = None
        __props__.__dict__["event_type_last_updated_time"] = None
        __props__.__dict__["last_updated_time"] = None
        __props__.__dict__["rule_execution_mode"] = None
        __props__.__dict__["rules"] = None
        __props__.__dict__["tags"] = None
        return Detector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="associatedModels")
    def associated_models(self) -> pulumi.Output[Optional[Sequence['outputs.DetectorModel']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-associatedmodels
        """
        return pulumi.get(self, "associated_models")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorid
        """
        return pulumi.get(self, "detector_id")

    @property
    @pulumi.getter(name="detectorVersionId")
    def detector_version_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "detector_version_id")

    @property
    @pulumi.getter(name="detectorVersionStatus")
    def detector_version_status(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorversionstatus
        """
        return pulumi.get(self, "detector_version_status")

    @property
    @pulumi.getter(name="eventType")
    def event_type(self) -> pulumi.Output['outputs.DetectorEventType']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-eventtype
        """
        return pulumi.get(self, "event_type")

    @property
    @pulumi.getter(name="eventTypeArn")
    def event_type_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "event_type_arn")

    @property
    @pulumi.getter(name="eventTypeCreatedTime")
    def event_type_created_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "event_type_created_time")

    @property
    @pulumi.getter(name="eventTypeLastUpdatedTime")
    def event_type_last_updated_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "event_type_last_updated_time")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter(name="ruleExecutionMode")
    def rule_execution_mode(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-ruleexecutionmode
        """
        return pulumi.get(self, "rule_execution_mode")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.DetectorRule']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-rules
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-tags
        """
        return pulumi.get(self, "tags")

