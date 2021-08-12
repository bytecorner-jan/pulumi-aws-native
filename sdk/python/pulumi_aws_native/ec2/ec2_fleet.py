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

__all__ = ['EC2FleetArgs', 'EC2Fleet']

@pulumi.input_type
class EC2FleetArgs:
    def __init__(__self__, *,
                 launch_template_configs: pulumi.Input[Sequence[pulumi.Input['EC2FleetFleetLaunchTemplateConfigRequestArgs']]],
                 target_capacity_specification: pulumi.Input['EC2FleetTargetCapacitySpecificationRequestArgs'],
                 context: Optional[pulumi.Input[str]] = None,
                 excess_capacity_termination_policy: Optional[pulumi.Input[str]] = None,
                 on_demand_options: Optional[pulumi.Input['EC2FleetOnDemandOptionsRequestArgs']] = None,
                 replace_unhealthy_instances: Optional[pulumi.Input[bool]] = None,
                 spot_options: Optional[pulumi.Input['EC2FleetSpotOptionsRequestArgs']] = None,
                 tag_specifications: Optional[pulumi.Input[Sequence[pulumi.Input['EC2FleetTagSpecificationArgs']]]] = None,
                 terminate_instances_with_expiration: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EC2Fleet resource.
        :param pulumi.Input[Sequence[pulumi.Input['EC2FleetFleetLaunchTemplateConfigRequestArgs']]] launch_template_configs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-launchtemplateconfigs
        :param pulumi.Input['EC2FleetTargetCapacitySpecificationRequestArgs'] target_capacity_specification: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-targetcapacityspecification
        :param pulumi.Input[str] context: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-context
        :param pulumi.Input[str] excess_capacity_termination_policy: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-excesscapacityterminationpolicy
        :param pulumi.Input['EC2FleetOnDemandOptionsRequestArgs'] on_demand_options: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-ondemandoptions
        :param pulumi.Input[bool] replace_unhealthy_instances: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-replaceunhealthyinstances
        :param pulumi.Input['EC2FleetSpotOptionsRequestArgs'] spot_options: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-spotoptions
        :param pulumi.Input[Sequence[pulumi.Input['EC2FleetTagSpecificationArgs']]] tag_specifications: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-tagspecifications
        :param pulumi.Input[bool] terminate_instances_with_expiration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-terminateinstanceswithexpiration
        :param pulumi.Input[str] type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-type
        :param pulumi.Input[str] valid_from: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validfrom
        :param pulumi.Input[str] valid_until: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validuntil
        """
        pulumi.set(__self__, "launch_template_configs", launch_template_configs)
        pulumi.set(__self__, "target_capacity_specification", target_capacity_specification)
        if context is not None:
            pulumi.set(__self__, "context", context)
        if excess_capacity_termination_policy is not None:
            pulumi.set(__self__, "excess_capacity_termination_policy", excess_capacity_termination_policy)
        if on_demand_options is not None:
            pulumi.set(__self__, "on_demand_options", on_demand_options)
        if replace_unhealthy_instances is not None:
            pulumi.set(__self__, "replace_unhealthy_instances", replace_unhealthy_instances)
        if spot_options is not None:
            pulumi.set(__self__, "spot_options", spot_options)
        if tag_specifications is not None:
            pulumi.set(__self__, "tag_specifications", tag_specifications)
        if terminate_instances_with_expiration is not None:
            pulumi.set(__self__, "terminate_instances_with_expiration", terminate_instances_with_expiration)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if valid_from is not None:
            pulumi.set(__self__, "valid_from", valid_from)
        if valid_until is not None:
            pulumi.set(__self__, "valid_until", valid_until)

    @property
    @pulumi.getter(name="launchTemplateConfigs")
    def launch_template_configs(self) -> pulumi.Input[Sequence[pulumi.Input['EC2FleetFleetLaunchTemplateConfigRequestArgs']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-launchtemplateconfigs
        """
        return pulumi.get(self, "launch_template_configs")

    @launch_template_configs.setter
    def launch_template_configs(self, value: pulumi.Input[Sequence[pulumi.Input['EC2FleetFleetLaunchTemplateConfigRequestArgs']]]):
        pulumi.set(self, "launch_template_configs", value)

    @property
    @pulumi.getter(name="targetCapacitySpecification")
    def target_capacity_specification(self) -> pulumi.Input['EC2FleetTargetCapacitySpecificationRequestArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-targetcapacityspecification
        """
        return pulumi.get(self, "target_capacity_specification")

    @target_capacity_specification.setter
    def target_capacity_specification(self, value: pulumi.Input['EC2FleetTargetCapacitySpecificationRequestArgs']):
        pulumi.set(self, "target_capacity_specification", value)

    @property
    @pulumi.getter
    def context(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-context
        """
        return pulumi.get(self, "context")

    @context.setter
    def context(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "context", value)

    @property
    @pulumi.getter(name="excessCapacityTerminationPolicy")
    def excess_capacity_termination_policy(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-excesscapacityterminationpolicy
        """
        return pulumi.get(self, "excess_capacity_termination_policy")

    @excess_capacity_termination_policy.setter
    def excess_capacity_termination_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "excess_capacity_termination_policy", value)

    @property
    @pulumi.getter(name="onDemandOptions")
    def on_demand_options(self) -> Optional[pulumi.Input['EC2FleetOnDemandOptionsRequestArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-ondemandoptions
        """
        return pulumi.get(self, "on_demand_options")

    @on_demand_options.setter
    def on_demand_options(self, value: Optional[pulumi.Input['EC2FleetOnDemandOptionsRequestArgs']]):
        pulumi.set(self, "on_demand_options", value)

    @property
    @pulumi.getter(name="replaceUnhealthyInstances")
    def replace_unhealthy_instances(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-replaceunhealthyinstances
        """
        return pulumi.get(self, "replace_unhealthy_instances")

    @replace_unhealthy_instances.setter
    def replace_unhealthy_instances(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "replace_unhealthy_instances", value)

    @property
    @pulumi.getter(name="spotOptions")
    def spot_options(self) -> Optional[pulumi.Input['EC2FleetSpotOptionsRequestArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-spotoptions
        """
        return pulumi.get(self, "spot_options")

    @spot_options.setter
    def spot_options(self, value: Optional[pulumi.Input['EC2FleetSpotOptionsRequestArgs']]):
        pulumi.set(self, "spot_options", value)

    @property
    @pulumi.getter(name="tagSpecifications")
    def tag_specifications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EC2FleetTagSpecificationArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-tagspecifications
        """
        return pulumi.get(self, "tag_specifications")

    @tag_specifications.setter
    def tag_specifications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EC2FleetTagSpecificationArgs']]]]):
        pulumi.set(self, "tag_specifications", value)

    @property
    @pulumi.getter(name="terminateInstancesWithExpiration")
    def terminate_instances_with_expiration(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-terminateinstanceswithexpiration
        """
        return pulumi.get(self, "terminate_instances_with_expiration")

    @terminate_instances_with_expiration.setter
    def terminate_instances_with_expiration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "terminate_instances_with_expiration", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validfrom
        """
        return pulumi.get(self, "valid_from")

    @valid_from.setter
    def valid_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_from", value)

    @property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validuntil
        """
        return pulumi.get(self, "valid_until")

    @valid_until.setter
    def valid_until(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_until", value)


class EC2Fleet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 excess_capacity_termination_policy: Optional[pulumi.Input[str]] = None,
                 launch_template_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EC2FleetFleetLaunchTemplateConfigRequestArgs']]]]] = None,
                 on_demand_options: Optional[pulumi.Input[pulumi.InputType['EC2FleetOnDemandOptionsRequestArgs']]] = None,
                 replace_unhealthy_instances: Optional[pulumi.Input[bool]] = None,
                 spot_options: Optional[pulumi.Input[pulumi.InputType['EC2FleetSpotOptionsRequestArgs']]] = None,
                 tag_specifications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EC2FleetTagSpecificationArgs']]]]] = None,
                 target_capacity_specification: Optional[pulumi.Input[pulumi.InputType['EC2FleetTargetCapacitySpecificationRequestArgs']]] = None,
                 terminate_instances_with_expiration: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-context
        :param pulumi.Input[str] excess_capacity_termination_policy: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-excesscapacityterminationpolicy
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EC2FleetFleetLaunchTemplateConfigRequestArgs']]]] launch_template_configs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-launchtemplateconfigs
        :param pulumi.Input[pulumi.InputType['EC2FleetOnDemandOptionsRequestArgs']] on_demand_options: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-ondemandoptions
        :param pulumi.Input[bool] replace_unhealthy_instances: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-replaceunhealthyinstances
        :param pulumi.Input[pulumi.InputType['EC2FleetSpotOptionsRequestArgs']] spot_options: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-spotoptions
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EC2FleetTagSpecificationArgs']]]] tag_specifications: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-tagspecifications
        :param pulumi.Input[pulumi.InputType['EC2FleetTargetCapacitySpecificationRequestArgs']] target_capacity_specification: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-targetcapacityspecification
        :param pulumi.Input[bool] terminate_instances_with_expiration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-terminateinstanceswithexpiration
        :param pulumi.Input[str] type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-type
        :param pulumi.Input[str] valid_from: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validfrom
        :param pulumi.Input[str] valid_until: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validuntil
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EC2FleetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html

        :param str resource_name: The name of the resource.
        :param EC2FleetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EC2FleetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 excess_capacity_termination_policy: Optional[pulumi.Input[str]] = None,
                 launch_template_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EC2FleetFleetLaunchTemplateConfigRequestArgs']]]]] = None,
                 on_demand_options: Optional[pulumi.Input[pulumi.InputType['EC2FleetOnDemandOptionsRequestArgs']]] = None,
                 replace_unhealthy_instances: Optional[pulumi.Input[bool]] = None,
                 spot_options: Optional[pulumi.Input[pulumi.InputType['EC2FleetSpotOptionsRequestArgs']]] = None,
                 tag_specifications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EC2FleetTagSpecificationArgs']]]]] = None,
                 target_capacity_specification: Optional[pulumi.Input[pulumi.InputType['EC2FleetTargetCapacitySpecificationRequestArgs']]] = None,
                 terminate_instances_with_expiration: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None,
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
            __props__ = EC2FleetArgs.__new__(EC2FleetArgs)

            __props__.__dict__["context"] = context
            __props__.__dict__["excess_capacity_termination_policy"] = excess_capacity_termination_policy
            if launch_template_configs is None and not opts.urn:
                raise TypeError("Missing required property 'launch_template_configs'")
            __props__.__dict__["launch_template_configs"] = launch_template_configs
            __props__.__dict__["on_demand_options"] = on_demand_options
            __props__.__dict__["replace_unhealthy_instances"] = replace_unhealthy_instances
            __props__.__dict__["spot_options"] = spot_options
            __props__.__dict__["tag_specifications"] = tag_specifications
            if target_capacity_specification is None and not opts.urn:
                raise TypeError("Missing required property 'target_capacity_specification'")
            __props__.__dict__["target_capacity_specification"] = target_capacity_specification
            __props__.__dict__["terminate_instances_with_expiration"] = terminate_instances_with_expiration
            __props__.__dict__["type"] = type
            __props__.__dict__["valid_from"] = valid_from
            __props__.__dict__["valid_until"] = valid_until
            __props__.__dict__["fleet_id"] = None
        super(EC2Fleet, __self__).__init__(
            'aws-native:EC2:EC2Fleet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EC2Fleet':
        """
        Get an existing EC2Fleet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EC2FleetArgs.__new__(EC2FleetArgs)

        __props__.__dict__["context"] = None
        __props__.__dict__["excess_capacity_termination_policy"] = None
        __props__.__dict__["fleet_id"] = None
        __props__.__dict__["launch_template_configs"] = None
        __props__.__dict__["on_demand_options"] = None
        __props__.__dict__["replace_unhealthy_instances"] = None
        __props__.__dict__["spot_options"] = None
        __props__.__dict__["tag_specifications"] = None
        __props__.__dict__["target_capacity_specification"] = None
        __props__.__dict__["terminate_instances_with_expiration"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["valid_from"] = None
        __props__.__dict__["valid_until"] = None
        return EC2Fleet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def context(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-context
        """
        return pulumi.get(self, "context")

    @property
    @pulumi.getter(name="excessCapacityTerminationPolicy")
    def excess_capacity_termination_policy(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-excesscapacityterminationpolicy
        """
        return pulumi.get(self, "excess_capacity_termination_policy")

    @property
    @pulumi.getter(name="fleetId")
    def fleet_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fleet_id")

    @property
    @pulumi.getter(name="launchTemplateConfigs")
    def launch_template_configs(self) -> pulumi.Output[Sequence['outputs.EC2FleetFleetLaunchTemplateConfigRequest']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-launchtemplateconfigs
        """
        return pulumi.get(self, "launch_template_configs")

    @property
    @pulumi.getter(name="onDemandOptions")
    def on_demand_options(self) -> pulumi.Output[Optional['outputs.EC2FleetOnDemandOptionsRequest']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-ondemandoptions
        """
        return pulumi.get(self, "on_demand_options")

    @property
    @pulumi.getter(name="replaceUnhealthyInstances")
    def replace_unhealthy_instances(self) -> pulumi.Output[Optional[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-replaceunhealthyinstances
        """
        return pulumi.get(self, "replace_unhealthy_instances")

    @property
    @pulumi.getter(name="spotOptions")
    def spot_options(self) -> pulumi.Output[Optional['outputs.EC2FleetSpotOptionsRequest']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-spotoptions
        """
        return pulumi.get(self, "spot_options")

    @property
    @pulumi.getter(name="tagSpecifications")
    def tag_specifications(self) -> pulumi.Output[Optional[Sequence['outputs.EC2FleetTagSpecification']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-tagspecifications
        """
        return pulumi.get(self, "tag_specifications")

    @property
    @pulumi.getter(name="targetCapacitySpecification")
    def target_capacity_specification(self) -> pulumi.Output['outputs.EC2FleetTargetCapacitySpecificationRequest']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-targetcapacityspecification
        """
        return pulumi.get(self, "target_capacity_specification")

    @property
    @pulumi.getter(name="terminateInstancesWithExpiration")
    def terminate_instances_with_expiration(self) -> pulumi.Output[Optional[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-terminateinstanceswithexpiration
        """
        return pulumi.get(self, "terminate_instances_with_expiration")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validfrom
        """
        return pulumi.get(self, "valid_from")

    @property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validuntil
        """
        return pulumi.get(self, "valid_until")

