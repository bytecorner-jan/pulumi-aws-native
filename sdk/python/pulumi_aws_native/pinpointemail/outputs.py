# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'ConfigurationSetDeliveryOptions',
    'ConfigurationSetEventDestinationCloudWatchDestination',
    'ConfigurationSetEventDestinationDimensionConfiguration',
    'ConfigurationSetEventDestinationEventDestination',
    'ConfigurationSetEventDestinationKinesisFirehoseDestination',
    'ConfigurationSetEventDestinationPinpointDestination',
    'ConfigurationSetEventDestinationSnsDestination',
    'ConfigurationSetReputationOptions',
    'ConfigurationSetSendingOptions',
    'ConfigurationSetTags',
    'ConfigurationSetTrackingOptions',
    'DedicatedIpPoolTags',
    'IdentityMailFromAttributes',
    'IdentityTags',
]

@pulumi.output_type
class ConfigurationSetDeliveryOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sendingPoolName":
            suggest = "sending_pool_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetDeliveryOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetDeliveryOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetDeliveryOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 sending_pool_name: Optional[str] = None):
        if sending_pool_name is not None:
            pulumi.set(__self__, "sending_pool_name", sending_pool_name)

    @property
    @pulumi.getter(name="sendingPoolName")
    def sending_pool_name(self) -> Optional[str]:
        return pulumi.get(self, "sending_pool_name")


@pulumi.output_type
class ConfigurationSetEventDestinationCloudWatchDestination(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dimensionConfigurations":
            suggest = "dimension_configurations"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetEventDestinationCloudWatchDestination. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetEventDestinationCloudWatchDestination.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetEventDestinationCloudWatchDestination.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dimension_configurations: Optional[Sequence['outputs.ConfigurationSetEventDestinationDimensionConfiguration']] = None):
        if dimension_configurations is not None:
            pulumi.set(__self__, "dimension_configurations", dimension_configurations)

    @property
    @pulumi.getter(name="dimensionConfigurations")
    def dimension_configurations(self) -> Optional[Sequence['outputs.ConfigurationSetEventDestinationDimensionConfiguration']]:
        return pulumi.get(self, "dimension_configurations")


@pulumi.output_type
class ConfigurationSetEventDestinationDimensionConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "defaultDimensionValue":
            suggest = "default_dimension_value"
        elif key == "dimensionName":
            suggest = "dimension_name"
        elif key == "dimensionValueSource":
            suggest = "dimension_value_source"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetEventDestinationDimensionConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetEventDestinationDimensionConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetEventDestinationDimensionConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 default_dimension_value: str,
                 dimension_name: str,
                 dimension_value_source: str):
        pulumi.set(__self__, "default_dimension_value", default_dimension_value)
        pulumi.set(__self__, "dimension_name", dimension_name)
        pulumi.set(__self__, "dimension_value_source", dimension_value_source)

    @property
    @pulumi.getter(name="defaultDimensionValue")
    def default_dimension_value(self) -> str:
        return pulumi.get(self, "default_dimension_value")

    @property
    @pulumi.getter(name="dimensionName")
    def dimension_name(self) -> str:
        return pulumi.get(self, "dimension_name")

    @property
    @pulumi.getter(name="dimensionValueSource")
    def dimension_value_source(self) -> str:
        return pulumi.get(self, "dimension_value_source")


@pulumi.output_type
class ConfigurationSetEventDestinationEventDestination(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "matchingEventTypes":
            suggest = "matching_event_types"
        elif key == "cloudWatchDestination":
            suggest = "cloud_watch_destination"
        elif key == "kinesisFirehoseDestination":
            suggest = "kinesis_firehose_destination"
        elif key == "pinpointDestination":
            suggest = "pinpoint_destination"
        elif key == "snsDestination":
            suggest = "sns_destination"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetEventDestinationEventDestination. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetEventDestinationEventDestination.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetEventDestinationEventDestination.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 matching_event_types: Sequence[str],
                 cloud_watch_destination: Optional['outputs.ConfigurationSetEventDestinationCloudWatchDestination'] = None,
                 enabled: Optional[bool] = None,
                 kinesis_firehose_destination: Optional['outputs.ConfigurationSetEventDestinationKinesisFirehoseDestination'] = None,
                 pinpoint_destination: Optional['outputs.ConfigurationSetEventDestinationPinpointDestination'] = None,
                 sns_destination: Optional['outputs.ConfigurationSetEventDestinationSnsDestination'] = None):
        pulumi.set(__self__, "matching_event_types", matching_event_types)
        if cloud_watch_destination is not None:
            pulumi.set(__self__, "cloud_watch_destination", cloud_watch_destination)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if kinesis_firehose_destination is not None:
            pulumi.set(__self__, "kinesis_firehose_destination", kinesis_firehose_destination)
        if pinpoint_destination is not None:
            pulumi.set(__self__, "pinpoint_destination", pinpoint_destination)
        if sns_destination is not None:
            pulumi.set(__self__, "sns_destination", sns_destination)

    @property
    @pulumi.getter(name="matchingEventTypes")
    def matching_event_types(self) -> Sequence[str]:
        return pulumi.get(self, "matching_event_types")

    @property
    @pulumi.getter(name="cloudWatchDestination")
    def cloud_watch_destination(self) -> Optional['outputs.ConfigurationSetEventDestinationCloudWatchDestination']:
        return pulumi.get(self, "cloud_watch_destination")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="kinesisFirehoseDestination")
    def kinesis_firehose_destination(self) -> Optional['outputs.ConfigurationSetEventDestinationKinesisFirehoseDestination']:
        return pulumi.get(self, "kinesis_firehose_destination")

    @property
    @pulumi.getter(name="pinpointDestination")
    def pinpoint_destination(self) -> Optional['outputs.ConfigurationSetEventDestinationPinpointDestination']:
        return pulumi.get(self, "pinpoint_destination")

    @property
    @pulumi.getter(name="snsDestination")
    def sns_destination(self) -> Optional['outputs.ConfigurationSetEventDestinationSnsDestination']:
        return pulumi.get(self, "sns_destination")


@pulumi.output_type
class ConfigurationSetEventDestinationKinesisFirehoseDestination(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deliveryStreamArn":
            suggest = "delivery_stream_arn"
        elif key == "iamRoleArn":
            suggest = "iam_role_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetEventDestinationKinesisFirehoseDestination. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetEventDestinationKinesisFirehoseDestination.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetEventDestinationKinesisFirehoseDestination.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 delivery_stream_arn: str,
                 iam_role_arn: str):
        pulumi.set(__self__, "delivery_stream_arn", delivery_stream_arn)
        pulumi.set(__self__, "iam_role_arn", iam_role_arn)

    @property
    @pulumi.getter(name="deliveryStreamArn")
    def delivery_stream_arn(self) -> str:
        return pulumi.get(self, "delivery_stream_arn")

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> str:
        return pulumi.get(self, "iam_role_arn")


@pulumi.output_type
class ConfigurationSetEventDestinationPinpointDestination(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "applicationArn":
            suggest = "application_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetEventDestinationPinpointDestination. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetEventDestinationPinpointDestination.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetEventDestinationPinpointDestination.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 application_arn: Optional[str] = None):
        if application_arn is not None:
            pulumi.set(__self__, "application_arn", application_arn)

    @property
    @pulumi.getter(name="applicationArn")
    def application_arn(self) -> Optional[str]:
        return pulumi.get(self, "application_arn")


@pulumi.output_type
class ConfigurationSetEventDestinationSnsDestination(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "topicArn":
            suggest = "topic_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetEventDestinationSnsDestination. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetEventDestinationSnsDestination.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetEventDestinationSnsDestination.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 topic_arn: str):
        pulumi.set(__self__, "topic_arn", topic_arn)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> str:
        return pulumi.get(self, "topic_arn")


@pulumi.output_type
class ConfigurationSetReputationOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "reputationMetricsEnabled":
            suggest = "reputation_metrics_enabled"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetReputationOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetReputationOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetReputationOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 reputation_metrics_enabled: Optional[bool] = None):
        if reputation_metrics_enabled is not None:
            pulumi.set(__self__, "reputation_metrics_enabled", reputation_metrics_enabled)

    @property
    @pulumi.getter(name="reputationMetricsEnabled")
    def reputation_metrics_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "reputation_metrics_enabled")


@pulumi.output_type
class ConfigurationSetSendingOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sendingEnabled":
            suggest = "sending_enabled"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetSendingOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetSendingOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetSendingOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 sending_enabled: Optional[bool] = None):
        if sending_enabled is not None:
            pulumi.set(__self__, "sending_enabled", sending_enabled)

    @property
    @pulumi.getter(name="sendingEnabled")
    def sending_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "sending_enabled")


@pulumi.output_type
class ConfigurationSetTags(dict):
    def __init__(__self__, *,
                 key: Optional[str] = None,
                 value: Optional[str] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


@pulumi.output_type
class ConfigurationSetTrackingOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "customRedirectDomain":
            suggest = "custom_redirect_domain"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationSetTrackingOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationSetTrackingOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationSetTrackingOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 custom_redirect_domain: Optional[str] = None):
        if custom_redirect_domain is not None:
            pulumi.set(__self__, "custom_redirect_domain", custom_redirect_domain)

    @property
    @pulumi.getter(name="customRedirectDomain")
    def custom_redirect_domain(self) -> Optional[str]:
        return pulumi.get(self, "custom_redirect_domain")


@pulumi.output_type
class DedicatedIpPoolTags(dict):
    def __init__(__self__, *,
                 key: Optional[str] = None,
                 value: Optional[str] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


@pulumi.output_type
class IdentityMailFromAttributes(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "behaviorOnMxFailure":
            suggest = "behavior_on_mx_failure"
        elif key == "mailFromDomain":
            suggest = "mail_from_domain"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IdentityMailFromAttributes. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IdentityMailFromAttributes.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IdentityMailFromAttributes.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 behavior_on_mx_failure: Optional[str] = None,
                 mail_from_domain: Optional[str] = None):
        if behavior_on_mx_failure is not None:
            pulumi.set(__self__, "behavior_on_mx_failure", behavior_on_mx_failure)
        if mail_from_domain is not None:
            pulumi.set(__self__, "mail_from_domain", mail_from_domain)

    @property
    @pulumi.getter(name="behaviorOnMxFailure")
    def behavior_on_mx_failure(self) -> Optional[str]:
        return pulumi.get(self, "behavior_on_mx_failure")

    @property
    @pulumi.getter(name="mailFromDomain")
    def mail_from_domain(self) -> Optional[str]:
        return pulumi.get(self, "mail_from_domain")


@pulumi.output_type
class IdentityTags(dict):
    def __init__(__self__, *,
                 key: Optional[str] = None,
                 value: Optional[str] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


