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

__all__ = ['StateMachineArgs', 'StateMachine']

@pulumi.input_type
class StateMachineArgs:
    def __init__(__self__, *,
                 role_arn: pulumi.Input[str],
                 definition: Optional[pulumi.Input['StateMachineDefinitionArgs']] = None,
                 definition_s3_location: Optional[pulumi.Input['StateMachineS3LocationArgs']] = None,
                 definition_string: Optional[pulumi.Input[str]] = None,
                 definition_substitutions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 logging_configuration: Optional[pulumi.Input['StateMachineLoggingConfigurationArgs']] = None,
                 state_machine_name: Optional[pulumi.Input[str]] = None,
                 state_machine_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['StateMachineTagsEntryArgs']]]] = None,
                 tracing_configuration: Optional[pulumi.Input['StateMachineTracingConfigurationArgs']] = None):
        """
        The set of arguments for constructing a StateMachine resource.
        :param pulumi.Input[str] role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
        :param pulumi.Input['StateMachineDefinitionArgs'] definition: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definition
        :param pulumi.Input['StateMachineS3LocationArgs'] definition_s3_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitions3location
        :param pulumi.Input[str] definition_string: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] definition_substitutions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionsubstitutions
        :param pulumi.Input['StateMachineLoggingConfigurationArgs'] logging_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-loggingconfiguration
        :param pulumi.Input[str] state_machine_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinename
        :param pulumi.Input[str] state_machine_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinetype
        :param pulumi.Input[Sequence[pulumi.Input['StateMachineTagsEntryArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tags
        :param pulumi.Input['StateMachineTracingConfigurationArgs'] tracing_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tracingconfiguration
        """
        pulumi.set(__self__, "role_arn", role_arn)
        if definition is not None:
            pulumi.set(__self__, "definition", definition)
        if definition_s3_location is not None:
            pulumi.set(__self__, "definition_s3_location", definition_s3_location)
        if definition_string is not None:
            pulumi.set(__self__, "definition_string", definition_string)
        if definition_substitutions is not None:
            pulumi.set(__self__, "definition_substitutions", definition_substitutions)
        if logging_configuration is not None:
            pulumi.set(__self__, "logging_configuration", logging_configuration)
        if state_machine_name is not None:
            pulumi.set(__self__, "state_machine_name", state_machine_name)
        if state_machine_type is not None:
            pulumi.set(__self__, "state_machine_type", state_machine_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tracing_configuration is not None:
            pulumi.set(__self__, "tracing_configuration", tracing_configuration)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def definition(self) -> Optional[pulumi.Input['StateMachineDefinitionArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definition
        """
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: Optional[pulumi.Input['StateMachineDefinitionArgs']]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter(name="definitionS3Location")
    def definition_s3_location(self) -> Optional[pulumi.Input['StateMachineS3LocationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitions3location
        """
        return pulumi.get(self, "definition_s3_location")

    @definition_s3_location.setter
    def definition_s3_location(self, value: Optional[pulumi.Input['StateMachineS3LocationArgs']]):
        pulumi.set(self, "definition_s3_location", value)

    @property
    @pulumi.getter(name="definitionString")
    def definition_string(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
        """
        return pulumi.get(self, "definition_string")

    @definition_string.setter
    def definition_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "definition_string", value)

    @property
    @pulumi.getter(name="definitionSubstitutions")
    def definition_substitutions(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionsubstitutions
        """
        return pulumi.get(self, "definition_substitutions")

    @definition_substitutions.setter
    def definition_substitutions(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "definition_substitutions", value)

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> Optional[pulumi.Input['StateMachineLoggingConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-loggingconfiguration
        """
        return pulumi.get(self, "logging_configuration")

    @logging_configuration.setter
    def logging_configuration(self, value: Optional[pulumi.Input['StateMachineLoggingConfigurationArgs']]):
        pulumi.set(self, "logging_configuration", value)

    @property
    @pulumi.getter(name="stateMachineName")
    def state_machine_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinename
        """
        return pulumi.get(self, "state_machine_name")

    @state_machine_name.setter
    def state_machine_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_machine_name", value)

    @property
    @pulumi.getter(name="stateMachineType")
    def state_machine_type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinetype
        """
        return pulumi.get(self, "state_machine_type")

    @state_machine_type.setter
    def state_machine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_machine_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StateMachineTagsEntryArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StateMachineTagsEntryArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tracingConfiguration")
    def tracing_configuration(self) -> Optional[pulumi.Input['StateMachineTracingConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tracingconfiguration
        """
        return pulumi.get(self, "tracing_configuration")

    @tracing_configuration.setter
    def tracing_configuration(self, value: Optional[pulumi.Input['StateMachineTracingConfigurationArgs']]):
        pulumi.set(self, "tracing_configuration", value)


class StateMachine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['StateMachineDefinitionArgs']]] = None,
                 definition_s3_location: Optional[pulumi.Input[pulumi.InputType['StateMachineS3LocationArgs']]] = None,
                 definition_string: Optional[pulumi.Input[str]] = None,
                 definition_substitutions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 logging_configuration: Optional[pulumi.Input[pulumi.InputType['StateMachineLoggingConfigurationArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 state_machine_name: Optional[pulumi.Input[str]] = None,
                 state_machine_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StateMachineTagsEntryArgs']]]]] = None,
                 tracing_configuration: Optional[pulumi.Input[pulumi.InputType['StateMachineTracingConfigurationArgs']]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['StateMachineDefinitionArgs']] definition: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definition
        :param pulumi.Input[pulumi.InputType['StateMachineS3LocationArgs']] definition_s3_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitions3location
        :param pulumi.Input[str] definition_string: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] definition_substitutions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionsubstitutions
        :param pulumi.Input[pulumi.InputType['StateMachineLoggingConfigurationArgs']] logging_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-loggingconfiguration
        :param pulumi.Input[str] role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
        :param pulumi.Input[str] state_machine_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinename
        :param pulumi.Input[str] state_machine_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinetype
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StateMachineTagsEntryArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tags
        :param pulumi.Input[pulumi.InputType['StateMachineTracingConfigurationArgs']] tracing_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tracingconfiguration
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StateMachineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html

        :param str resource_name: The name of the resource.
        :param StateMachineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StateMachineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['StateMachineDefinitionArgs']]] = None,
                 definition_s3_location: Optional[pulumi.Input[pulumi.InputType['StateMachineS3LocationArgs']]] = None,
                 definition_string: Optional[pulumi.Input[str]] = None,
                 definition_substitutions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 logging_configuration: Optional[pulumi.Input[pulumi.InputType['StateMachineLoggingConfigurationArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 state_machine_name: Optional[pulumi.Input[str]] = None,
                 state_machine_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StateMachineTagsEntryArgs']]]]] = None,
                 tracing_configuration: Optional[pulumi.Input[pulumi.InputType['StateMachineTracingConfigurationArgs']]] = None,
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
            __props__ = StateMachineArgs.__new__(StateMachineArgs)

            __props__.__dict__["definition"] = definition
            __props__.__dict__["definition_s3_location"] = definition_s3_location
            __props__.__dict__["definition_string"] = definition_string
            __props__.__dict__["definition_substitutions"] = definition_substitutions
            __props__.__dict__["logging_configuration"] = logging_configuration
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["state_machine_name"] = state_machine_name
            __props__.__dict__["state_machine_type"] = state_machine_type
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tracing_configuration"] = tracing_configuration
            __props__.__dict__["arn"] = None
            __props__.__dict__["name"] = None
        super(StateMachine, __self__).__init__(
            'aws-native:StepFunctions:StateMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'StateMachine':
        """
        Get an existing StateMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StateMachineArgs.__new__(StateMachineArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["definition"] = None
        __props__.__dict__["definition_s3_location"] = None
        __props__.__dict__["definition_string"] = None
        __props__.__dict__["definition_substitutions"] = None
        __props__.__dict__["logging_configuration"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["state_machine_name"] = None
        __props__.__dict__["state_machine_type"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["tracing_configuration"] = None
        return StateMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[Optional['outputs.StateMachineDefinition']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definition
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter(name="definitionS3Location")
    def definition_s3_location(self) -> pulumi.Output[Optional['outputs.StateMachineS3Location']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitions3location
        """
        return pulumi.get(self, "definition_s3_location")

    @property
    @pulumi.getter(name="definitionString")
    def definition_string(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
        """
        return pulumi.get(self, "definition_string")

    @property
    @pulumi.getter(name="definitionSubstitutions")
    def definition_substitutions(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionsubstitutions
        """
        return pulumi.get(self, "definition_substitutions")

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> pulumi.Output[Optional['outputs.StateMachineLoggingConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-loggingconfiguration
        """
        return pulumi.get(self, "logging_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="stateMachineName")
    def state_machine_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinename
        """
        return pulumi.get(self, "state_machine_name")

    @property
    @pulumi.getter(name="stateMachineType")
    def state_machine_type(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinetype
        """
        return pulumi.get(self, "state_machine_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.StateMachineTagsEntry']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tracingConfiguration")
    def tracing_configuration(self) -> pulumi.Output[Optional['outputs.StateMachineTracingConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tracingconfiguration
        """
        return pulumi.get(self, "tracing_configuration")

