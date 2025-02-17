# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'MetricFilterMetricTransformationArgs',
]

@pulumi.input_type
class MetricFilterMetricTransformationArgs:
    def __init__(__self__, *,
                 metric_name: pulumi.Input[str],
                 metric_namespace: pulumi.Input[str],
                 metric_value: pulumi.Input[str],
                 default_value: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "metric_name", metric_name)
        pulumi.set(__self__, "metric_namespace", metric_namespace)
        pulumi.set(__self__, "metric_value", metric_value)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter(name="metricNamespace")
    def metric_namespace(self) -> pulumi.Input[str]:
        return pulumi.get(self, "metric_namespace")

    @metric_namespace.setter
    def metric_namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric_namespace", value)

    @property
    @pulumi.getter(name="metricValue")
    def metric_value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "metric_value")

    @metric_value.setter
    def metric_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric_value", value)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "default_value", value)


