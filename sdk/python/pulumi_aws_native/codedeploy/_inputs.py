# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ApplicationTagArgs',
    'DeploymentConfigMinimumHealthyHostsArgs',
    'DeploymentConfigTimeBasedCanaryArgs',
    'DeploymentConfigTimeBasedLinearArgs',
    'DeploymentConfigTrafficRoutingConfigArgs',
    'DeploymentGroupAlarmConfigurationArgs',
    'DeploymentGroupAlarmArgs',
    'DeploymentGroupAutoRollbackConfigurationArgs',
    'DeploymentGroupBlueGreenDeploymentConfigurationArgs',
    'DeploymentGroupBlueInstanceTerminationOptionArgs',
    'DeploymentGroupDeploymentReadyOptionArgs',
    'DeploymentGroupDeploymentStyleArgs',
    'DeploymentGroupDeploymentArgs',
    'DeploymentGroupEC2TagFilterArgs',
    'DeploymentGroupEC2TagSetListObjectArgs',
    'DeploymentGroupEC2TagSetArgs',
    'DeploymentGroupECSServiceArgs',
    'DeploymentGroupELBInfoArgs',
    'DeploymentGroupGitHubLocationArgs',
    'DeploymentGroupGreenFleetProvisioningOptionArgs',
    'DeploymentGroupLoadBalancerInfoArgs',
    'DeploymentGroupOnPremisesTagSetListObjectArgs',
    'DeploymentGroupOnPremisesTagSetArgs',
    'DeploymentGroupRevisionLocationArgs',
    'DeploymentGroupS3LocationArgs',
    'DeploymentGroupTagFilterArgs',
    'DeploymentGroupTargetGroupInfoArgs',
    'DeploymentGroupTriggerConfigArgs',
]

@pulumi.input_type
class ApplicationTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DeploymentConfigMinimumHealthyHostsArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 value: pulumi.Input[int]):
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[int]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[int]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DeploymentConfigTimeBasedCanaryArgs:
    def __init__(__self__, *,
                 canary_interval: pulumi.Input[int],
                 canary_percentage: pulumi.Input[int]):
        pulumi.set(__self__, "canary_interval", canary_interval)
        pulumi.set(__self__, "canary_percentage", canary_percentage)

    @property
    @pulumi.getter(name="canaryInterval")
    def canary_interval(self) -> pulumi.Input[int]:
        return pulumi.get(self, "canary_interval")

    @canary_interval.setter
    def canary_interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "canary_interval", value)

    @property
    @pulumi.getter(name="canaryPercentage")
    def canary_percentage(self) -> pulumi.Input[int]:
        return pulumi.get(self, "canary_percentage")

    @canary_percentage.setter
    def canary_percentage(self, value: pulumi.Input[int]):
        pulumi.set(self, "canary_percentage", value)


@pulumi.input_type
class DeploymentConfigTimeBasedLinearArgs:
    def __init__(__self__, *,
                 linear_interval: pulumi.Input[int],
                 linear_percentage: pulumi.Input[int]):
        pulumi.set(__self__, "linear_interval", linear_interval)
        pulumi.set(__self__, "linear_percentage", linear_percentage)

    @property
    @pulumi.getter(name="linearInterval")
    def linear_interval(self) -> pulumi.Input[int]:
        return pulumi.get(self, "linear_interval")

    @linear_interval.setter
    def linear_interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "linear_interval", value)

    @property
    @pulumi.getter(name="linearPercentage")
    def linear_percentage(self) -> pulumi.Input[int]:
        return pulumi.get(self, "linear_percentage")

    @linear_percentage.setter
    def linear_percentage(self, value: pulumi.Input[int]):
        pulumi.set(self, "linear_percentage", value)


@pulumi.input_type
class DeploymentConfigTrafficRoutingConfigArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 time_based_canary: Optional[pulumi.Input['DeploymentConfigTimeBasedCanaryArgs']] = None,
                 time_based_linear: Optional[pulumi.Input['DeploymentConfigTimeBasedLinearArgs']] = None):
        pulumi.set(__self__, "type", type)
        if time_based_canary is not None:
            pulumi.set(__self__, "time_based_canary", time_based_canary)
        if time_based_linear is not None:
            pulumi.set(__self__, "time_based_linear", time_based_linear)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="timeBasedCanary")
    def time_based_canary(self) -> Optional[pulumi.Input['DeploymentConfigTimeBasedCanaryArgs']]:
        return pulumi.get(self, "time_based_canary")

    @time_based_canary.setter
    def time_based_canary(self, value: Optional[pulumi.Input['DeploymentConfigTimeBasedCanaryArgs']]):
        pulumi.set(self, "time_based_canary", value)

    @property
    @pulumi.getter(name="timeBasedLinear")
    def time_based_linear(self) -> Optional[pulumi.Input['DeploymentConfigTimeBasedLinearArgs']]:
        return pulumi.get(self, "time_based_linear")

    @time_based_linear.setter
    def time_based_linear(self, value: Optional[pulumi.Input['DeploymentConfigTimeBasedLinearArgs']]):
        pulumi.set(self, "time_based_linear", value)


@pulumi.input_type
class DeploymentGroupAlarmConfigurationArgs:
    def __init__(__self__, *,
                 alarms: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupAlarmArgs']]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ignore_poll_alarm_failure: Optional[pulumi.Input[bool]] = None):
        if alarms is not None:
            pulumi.set(__self__, "alarms", alarms)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if ignore_poll_alarm_failure is not None:
            pulumi.set(__self__, "ignore_poll_alarm_failure", ignore_poll_alarm_failure)

    @property
    @pulumi.getter
    def alarms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupAlarmArgs']]]]:
        return pulumi.get(self, "alarms")

    @alarms.setter
    def alarms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupAlarmArgs']]]]):
        pulumi.set(self, "alarms", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="ignorePollAlarmFailure")
    def ignore_poll_alarm_failure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "ignore_poll_alarm_failure")

    @ignore_poll_alarm_failure.setter
    def ignore_poll_alarm_failure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_poll_alarm_failure", value)


@pulumi.input_type
class DeploymentGroupAlarmArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DeploymentGroupAutoRollbackConfigurationArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if events is not None:
            pulumi.set(__self__, "events", events)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def events(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "events")

    @events.setter
    def events(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "events", value)


@pulumi.input_type
class DeploymentGroupBlueGreenDeploymentConfigurationArgs:
    def __init__(__self__, *,
                 deployment_ready_option: Optional[pulumi.Input['DeploymentGroupDeploymentReadyOptionArgs']] = None,
                 green_fleet_provisioning_option: Optional[pulumi.Input['DeploymentGroupGreenFleetProvisioningOptionArgs']] = None,
                 terminate_blue_instances_on_deployment_success: Optional[pulumi.Input['DeploymentGroupBlueInstanceTerminationOptionArgs']] = None):
        if deployment_ready_option is not None:
            pulumi.set(__self__, "deployment_ready_option", deployment_ready_option)
        if green_fleet_provisioning_option is not None:
            pulumi.set(__self__, "green_fleet_provisioning_option", green_fleet_provisioning_option)
        if terminate_blue_instances_on_deployment_success is not None:
            pulumi.set(__self__, "terminate_blue_instances_on_deployment_success", terminate_blue_instances_on_deployment_success)

    @property
    @pulumi.getter(name="deploymentReadyOption")
    def deployment_ready_option(self) -> Optional[pulumi.Input['DeploymentGroupDeploymentReadyOptionArgs']]:
        return pulumi.get(self, "deployment_ready_option")

    @deployment_ready_option.setter
    def deployment_ready_option(self, value: Optional[pulumi.Input['DeploymentGroupDeploymentReadyOptionArgs']]):
        pulumi.set(self, "deployment_ready_option", value)

    @property
    @pulumi.getter(name="greenFleetProvisioningOption")
    def green_fleet_provisioning_option(self) -> Optional[pulumi.Input['DeploymentGroupGreenFleetProvisioningOptionArgs']]:
        return pulumi.get(self, "green_fleet_provisioning_option")

    @green_fleet_provisioning_option.setter
    def green_fleet_provisioning_option(self, value: Optional[pulumi.Input['DeploymentGroupGreenFleetProvisioningOptionArgs']]):
        pulumi.set(self, "green_fleet_provisioning_option", value)

    @property
    @pulumi.getter(name="terminateBlueInstancesOnDeploymentSuccess")
    def terminate_blue_instances_on_deployment_success(self) -> Optional[pulumi.Input['DeploymentGroupBlueInstanceTerminationOptionArgs']]:
        return pulumi.get(self, "terminate_blue_instances_on_deployment_success")

    @terminate_blue_instances_on_deployment_success.setter
    def terminate_blue_instances_on_deployment_success(self, value: Optional[pulumi.Input['DeploymentGroupBlueInstanceTerminationOptionArgs']]):
        pulumi.set(self, "terminate_blue_instances_on_deployment_success", value)


@pulumi.input_type
class DeploymentGroupBlueInstanceTerminationOptionArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 termination_wait_time_in_minutes: Optional[pulumi.Input[int]] = None):
        if action is not None:
            pulumi.set(__self__, "action", action)
        if termination_wait_time_in_minutes is not None:
            pulumi.set(__self__, "termination_wait_time_in_minutes", termination_wait_time_in_minutes)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="terminationWaitTimeInMinutes")
    def termination_wait_time_in_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "termination_wait_time_in_minutes")

    @termination_wait_time_in_minutes.setter
    def termination_wait_time_in_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "termination_wait_time_in_minutes", value)


@pulumi.input_type
class DeploymentGroupDeploymentReadyOptionArgs:
    def __init__(__self__, *,
                 action_on_timeout: Optional[pulumi.Input[str]] = None,
                 wait_time_in_minutes: Optional[pulumi.Input[int]] = None):
        if action_on_timeout is not None:
            pulumi.set(__self__, "action_on_timeout", action_on_timeout)
        if wait_time_in_minutes is not None:
            pulumi.set(__self__, "wait_time_in_minutes", wait_time_in_minutes)

    @property
    @pulumi.getter(name="actionOnTimeout")
    def action_on_timeout(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action_on_timeout")

    @action_on_timeout.setter
    def action_on_timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_on_timeout", value)

    @property
    @pulumi.getter(name="waitTimeInMinutes")
    def wait_time_in_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "wait_time_in_minutes")

    @wait_time_in_minutes.setter
    def wait_time_in_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_time_in_minutes", value)


@pulumi.input_type
class DeploymentGroupDeploymentStyleArgs:
    def __init__(__self__, *,
                 deployment_option: Optional[pulumi.Input[str]] = None,
                 deployment_type: Optional[pulumi.Input[str]] = None):
        if deployment_option is not None:
            pulumi.set(__self__, "deployment_option", deployment_option)
        if deployment_type is not None:
            pulumi.set(__self__, "deployment_type", deployment_type)

    @property
    @pulumi.getter(name="deploymentOption")
    def deployment_option(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "deployment_option")

    @deployment_option.setter
    def deployment_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deployment_option", value)

    @property
    @pulumi.getter(name="deploymentType")
    def deployment_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "deployment_type")

    @deployment_type.setter
    def deployment_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deployment_type", value)


@pulumi.input_type
class DeploymentGroupDeploymentArgs:
    def __init__(__self__, *,
                 revision: pulumi.Input['DeploymentGroupRevisionLocationArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 ignore_application_stop_failures: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "revision", revision)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ignore_application_stop_failures is not None:
            pulumi.set(__self__, "ignore_application_stop_failures", ignore_application_stop_failures)

    @property
    @pulumi.getter
    def revision(self) -> pulumi.Input['DeploymentGroupRevisionLocationArgs']:
        return pulumi.get(self, "revision")

    @revision.setter
    def revision(self, value: pulumi.Input['DeploymentGroupRevisionLocationArgs']):
        pulumi.set(self, "revision", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ignoreApplicationStopFailures")
    def ignore_application_stop_failures(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "ignore_application_stop_failures")

    @ignore_application_stop_failures.setter
    def ignore_application_stop_failures(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_application_stop_failures", value)


@pulumi.input_type
class DeploymentGroupEC2TagFilterArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DeploymentGroupEC2TagSetListObjectArgs:
    def __init__(__self__, *,
                 ec2_tag_group: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupEC2TagFilterArgs']]]] = None):
        if ec2_tag_group is not None:
            pulumi.set(__self__, "ec2_tag_group", ec2_tag_group)

    @property
    @pulumi.getter(name="ec2TagGroup")
    def ec2_tag_group(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupEC2TagFilterArgs']]]]:
        return pulumi.get(self, "ec2_tag_group")

    @ec2_tag_group.setter
    def ec2_tag_group(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupEC2TagFilterArgs']]]]):
        pulumi.set(self, "ec2_tag_group", value)


@pulumi.input_type
class DeploymentGroupEC2TagSetArgs:
    def __init__(__self__, *,
                 ec2_tag_set_list: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupEC2TagSetListObjectArgs']]]] = None):
        if ec2_tag_set_list is not None:
            pulumi.set(__self__, "ec2_tag_set_list", ec2_tag_set_list)

    @property
    @pulumi.getter(name="ec2TagSetList")
    def ec2_tag_set_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupEC2TagSetListObjectArgs']]]]:
        return pulumi.get(self, "ec2_tag_set_list")

    @ec2_tag_set_list.setter
    def ec2_tag_set_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupEC2TagSetListObjectArgs']]]]):
        pulumi.set(self, "ec2_tag_set_list", value)


@pulumi.input_type
class DeploymentGroupECSServiceArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class DeploymentGroupELBInfoArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DeploymentGroupGitHubLocationArgs:
    def __init__(__self__, *,
                 commit_id: pulumi.Input[str],
                 repository: pulumi.Input[str]):
        pulumi.set(__self__, "commit_id", commit_id)
        pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter(name="commitId")
    def commit_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "commit_id")

    @commit_id.setter
    def commit_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "commit_id", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)


@pulumi.input_type
class DeploymentGroupGreenFleetProvisioningOptionArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None):
        if action is not None:
            pulumi.set(__self__, "action", action)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)


@pulumi.input_type
class DeploymentGroupLoadBalancerInfoArgs:
    def __init__(__self__, *,
                 elb_info_list: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupELBInfoArgs']]]] = None,
                 target_group_info_list: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupTargetGroupInfoArgs']]]] = None):
        if elb_info_list is not None:
            pulumi.set(__self__, "elb_info_list", elb_info_list)
        if target_group_info_list is not None:
            pulumi.set(__self__, "target_group_info_list", target_group_info_list)

    @property
    @pulumi.getter(name="elbInfoList")
    def elb_info_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupELBInfoArgs']]]]:
        return pulumi.get(self, "elb_info_list")

    @elb_info_list.setter
    def elb_info_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupELBInfoArgs']]]]):
        pulumi.set(self, "elb_info_list", value)

    @property
    @pulumi.getter(name="targetGroupInfoList")
    def target_group_info_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupTargetGroupInfoArgs']]]]:
        return pulumi.get(self, "target_group_info_list")

    @target_group_info_list.setter
    def target_group_info_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupTargetGroupInfoArgs']]]]):
        pulumi.set(self, "target_group_info_list", value)


@pulumi.input_type
class DeploymentGroupOnPremisesTagSetListObjectArgs:
    def __init__(__self__, *,
                 on_premises_tag_group: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupTagFilterArgs']]]] = None):
        if on_premises_tag_group is not None:
            pulumi.set(__self__, "on_premises_tag_group", on_premises_tag_group)

    @property
    @pulumi.getter(name="onPremisesTagGroup")
    def on_premises_tag_group(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupTagFilterArgs']]]]:
        return pulumi.get(self, "on_premises_tag_group")

    @on_premises_tag_group.setter
    def on_premises_tag_group(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupTagFilterArgs']]]]):
        pulumi.set(self, "on_premises_tag_group", value)


@pulumi.input_type
class DeploymentGroupOnPremisesTagSetArgs:
    def __init__(__self__, *,
                 on_premises_tag_set_list: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupOnPremisesTagSetListObjectArgs']]]] = None):
        if on_premises_tag_set_list is not None:
            pulumi.set(__self__, "on_premises_tag_set_list", on_premises_tag_set_list)

    @property
    @pulumi.getter(name="onPremisesTagSetList")
    def on_premises_tag_set_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupOnPremisesTagSetListObjectArgs']]]]:
        return pulumi.get(self, "on_premises_tag_set_list")

    @on_premises_tag_set_list.setter
    def on_premises_tag_set_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentGroupOnPremisesTagSetListObjectArgs']]]]):
        pulumi.set(self, "on_premises_tag_set_list", value)


@pulumi.input_type
class DeploymentGroupRevisionLocationArgs:
    def __init__(__self__, *,
                 git_hub_location: Optional[pulumi.Input['DeploymentGroupGitHubLocationArgs']] = None,
                 revision_type: Optional[pulumi.Input[str]] = None,
                 s3_location: Optional[pulumi.Input['DeploymentGroupS3LocationArgs']] = None):
        if git_hub_location is not None:
            pulumi.set(__self__, "git_hub_location", git_hub_location)
        if revision_type is not None:
            pulumi.set(__self__, "revision_type", revision_type)
        if s3_location is not None:
            pulumi.set(__self__, "s3_location", s3_location)

    @property
    @pulumi.getter(name="gitHubLocation")
    def git_hub_location(self) -> Optional[pulumi.Input['DeploymentGroupGitHubLocationArgs']]:
        return pulumi.get(self, "git_hub_location")

    @git_hub_location.setter
    def git_hub_location(self, value: Optional[pulumi.Input['DeploymentGroupGitHubLocationArgs']]):
        pulumi.set(self, "git_hub_location", value)

    @property
    @pulumi.getter(name="revisionType")
    def revision_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "revision_type")

    @revision_type.setter
    def revision_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "revision_type", value)

    @property
    @pulumi.getter(name="s3Location")
    def s3_location(self) -> Optional[pulumi.Input['DeploymentGroupS3LocationArgs']]:
        return pulumi.get(self, "s3_location")

    @s3_location.setter
    def s3_location(self, value: Optional[pulumi.Input['DeploymentGroupS3LocationArgs']]):
        pulumi.set(self, "s3_location", value)


@pulumi.input_type
class DeploymentGroupS3LocationArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 key: pulumi.Input[str],
                 bundle_type: Optional[pulumi.Input[str]] = None,
                 e_tag: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)
        if bundle_type is not None:
            pulumi.set(__self__, "bundle_type", bundle_type)
        if e_tag is not None:
            pulumi.set(__self__, "e_tag", e_tag)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="bundleType")
    def bundle_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bundle_type")

    @bundle_type.setter
    def bundle_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bundle_type", value)

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "e_tag")

    @e_tag.setter
    def e_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "e_tag", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class DeploymentGroupTagFilterArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DeploymentGroupTargetGroupInfoArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DeploymentGroupTriggerConfigArgs:
    def __init__(__self__, *,
                 trigger_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 trigger_name: Optional[pulumi.Input[str]] = None,
                 trigger_target_arn: Optional[pulumi.Input[str]] = None):
        if trigger_events is not None:
            pulumi.set(__self__, "trigger_events", trigger_events)
        if trigger_name is not None:
            pulumi.set(__self__, "trigger_name", trigger_name)
        if trigger_target_arn is not None:
            pulumi.set(__self__, "trigger_target_arn", trigger_target_arn)

    @property
    @pulumi.getter(name="triggerEvents")
    def trigger_events(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "trigger_events")

    @trigger_events.setter
    def trigger_events(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "trigger_events", value)

    @property
    @pulumi.getter(name="triggerName")
    def trigger_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trigger_name")

    @trigger_name.setter
    def trigger_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_name", value)

    @property
    @pulumi.getter(name="triggerTargetArn")
    def trigger_target_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trigger_target_arn")

    @trigger_target_arn.setter
    def trigger_target_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_target_arn", value)


