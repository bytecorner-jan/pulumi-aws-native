# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ExperimentTemplateActionMap',
    'ExperimentTemplateStopCondition',
    'ExperimentTemplateTargetMap',
]

@pulumi.output_type
class ExperimentTemplateActionMap(dict):
    """
    The actions for the experiment.
    """
    def __init__(__self__):
        """
        The actions for the experiment.
        """
        pass


@pulumi.output_type
class ExperimentTemplateStopCondition(dict):
    def __init__(__self__, *,
                 source: str,
                 value: Optional[str] = None):
        pulumi.set(__self__, "source", source)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def source(self) -> str:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


@pulumi.output_type
class ExperimentTemplateTargetMap(dict):
    """
    The targets for the experiment.
    """
    def __init__(__self__):
        """
        The targets for the experiment.
        """
        pass


