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
    'ApplicationMaxAgeRule',
    'ApplicationMaxCountRule',
    'ApplicationResourceLifecycleConfig',
    'ApplicationVersionLifecycleConfig',
    'ApplicationVersionSourceBundle',
    'ConfigurationTemplateConfigurationOptionSetting',
    'ConfigurationTemplateSourceConfiguration',
    'EnvironmentOptionSetting',
    'EnvironmentTag',
    'EnvironmentTier',
]

@pulumi.output_type
class ApplicationMaxAgeRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deleteSourceFromS3":
            suggest = "delete_source_from_s3"
        elif key == "maxAgeInDays":
            suggest = "max_age_in_days"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationMaxAgeRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationMaxAgeRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationMaxAgeRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 delete_source_from_s3: Optional[bool] = None,
                 enabled: Optional[bool] = None,
                 max_age_in_days: Optional[int] = None):
        if delete_source_from_s3 is not None:
            pulumi.set(__self__, "delete_source_from_s3", delete_source_from_s3)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if max_age_in_days is not None:
            pulumi.set(__self__, "max_age_in_days", max_age_in_days)

    @property
    @pulumi.getter(name="deleteSourceFromS3")
    def delete_source_from_s3(self) -> Optional[bool]:
        return pulumi.get(self, "delete_source_from_s3")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="maxAgeInDays")
    def max_age_in_days(self) -> Optional[int]:
        return pulumi.get(self, "max_age_in_days")


@pulumi.output_type
class ApplicationMaxCountRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deleteSourceFromS3":
            suggest = "delete_source_from_s3"
        elif key == "maxCount":
            suggest = "max_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationMaxCountRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationMaxCountRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationMaxCountRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 delete_source_from_s3: Optional[bool] = None,
                 enabled: Optional[bool] = None,
                 max_count: Optional[int] = None):
        if delete_source_from_s3 is not None:
            pulumi.set(__self__, "delete_source_from_s3", delete_source_from_s3)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if max_count is not None:
            pulumi.set(__self__, "max_count", max_count)

    @property
    @pulumi.getter(name="deleteSourceFromS3")
    def delete_source_from_s3(self) -> Optional[bool]:
        return pulumi.get(self, "delete_source_from_s3")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> Optional[int]:
        return pulumi.get(self, "max_count")


@pulumi.output_type
class ApplicationResourceLifecycleConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "serviceRole":
            suggest = "service_role"
        elif key == "versionLifecycleConfig":
            suggest = "version_lifecycle_config"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationResourceLifecycleConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationResourceLifecycleConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationResourceLifecycleConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 service_role: Optional[str] = None,
                 version_lifecycle_config: Optional['outputs.ApplicationVersionLifecycleConfig'] = None):
        if service_role is not None:
            pulumi.set(__self__, "service_role", service_role)
        if version_lifecycle_config is not None:
            pulumi.set(__self__, "version_lifecycle_config", version_lifecycle_config)

    @property
    @pulumi.getter(name="serviceRole")
    def service_role(self) -> Optional[str]:
        return pulumi.get(self, "service_role")

    @property
    @pulumi.getter(name="versionLifecycleConfig")
    def version_lifecycle_config(self) -> Optional['outputs.ApplicationVersionLifecycleConfig']:
        return pulumi.get(self, "version_lifecycle_config")


@pulumi.output_type
class ApplicationVersionLifecycleConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxAgeRule":
            suggest = "max_age_rule"
        elif key == "maxCountRule":
            suggest = "max_count_rule"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationVersionLifecycleConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationVersionLifecycleConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationVersionLifecycleConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_age_rule: Optional['outputs.ApplicationMaxAgeRule'] = None,
                 max_count_rule: Optional['outputs.ApplicationMaxCountRule'] = None):
        if max_age_rule is not None:
            pulumi.set(__self__, "max_age_rule", max_age_rule)
        if max_count_rule is not None:
            pulumi.set(__self__, "max_count_rule", max_count_rule)

    @property
    @pulumi.getter(name="maxAgeRule")
    def max_age_rule(self) -> Optional['outputs.ApplicationMaxAgeRule']:
        return pulumi.get(self, "max_age_rule")

    @property
    @pulumi.getter(name="maxCountRule")
    def max_count_rule(self) -> Optional['outputs.ApplicationMaxCountRule']:
        return pulumi.get(self, "max_count_rule")


@pulumi.output_type
class ApplicationVersionSourceBundle(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "s3Bucket":
            suggest = "s3_bucket"
        elif key == "s3Key":
            suggest = "s3_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationVersionSourceBundle. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationVersionSourceBundle.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationVersionSourceBundle.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 s3_bucket: str,
                 s3_key: str):
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        pulumi.set(__self__, "s3_key", s3_key)

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> str:
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> str:
        return pulumi.get(self, "s3_key")


@pulumi.output_type
class ConfigurationTemplateConfigurationOptionSetting(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "optionName":
            suggest = "option_name"
        elif key == "resourceName":
            suggest = "resource_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationTemplateConfigurationOptionSetting. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationTemplateConfigurationOptionSetting.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationTemplateConfigurationOptionSetting.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 namespace: str,
                 option_name: str,
                 resource_name: Optional[str] = None,
                 value: Optional[str] = None):
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "option_name", option_name)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="optionName")
    def option_name(self) -> str:
        return pulumi.get(self, "option_name")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[str]:
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


@pulumi.output_type
class ConfigurationTemplateSourceConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "applicationName":
            suggest = "application_name"
        elif key == "templateName":
            suggest = "template_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationTemplateSourceConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationTemplateSourceConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationTemplateSourceConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 application_name: str,
                 template_name: str):
        pulumi.set(__self__, "application_name", application_name)
        pulumi.set(__self__, "template_name", template_name)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> str:
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> str:
        return pulumi.get(self, "template_name")


@pulumi.output_type
class EnvironmentOptionSetting(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "optionName":
            suggest = "option_name"
        elif key == "resourceName":
            suggest = "resource_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EnvironmentOptionSetting. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EnvironmentOptionSetting.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EnvironmentOptionSetting.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 namespace: str,
                 option_name: str,
                 resource_name: Optional[str] = None,
                 value: Optional[str] = None):
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "option_name", option_name)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="optionName")
    def option_name(self) -> str:
        return pulumi.get(self, "option_name")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[str]:
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


@pulumi.output_type
class EnvironmentTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class EnvironmentTier(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 type: Optional[str] = None,
                 version: Optional[str] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


