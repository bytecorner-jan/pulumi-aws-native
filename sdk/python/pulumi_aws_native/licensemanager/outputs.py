# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'LicenseBorrowConfiguration',
    'LicenseConsumptionConfiguration',
    'LicenseEntitlement',
    'LicenseIssuerData',
    'LicenseMetadata',
    'LicenseProvisionalConfiguration',
    'LicenseValidityDateFormat',
]

@pulumi.output_type
class LicenseBorrowConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowEarlyCheckIn":
            suggest = "allow_early_check_in"
        elif key == "maxTimeToLiveInMinutes":
            suggest = "max_time_to_live_in_minutes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LicenseBorrowConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LicenseBorrowConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LicenseBorrowConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allow_early_check_in: bool,
                 max_time_to_live_in_minutes: int):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html
        :param bool allow_early_check_in: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html#cfn-licensemanager-license-borrowconfiguration-allowearlycheckin
        :param int max_time_to_live_in_minutes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html#cfn-licensemanager-license-borrowconfiguration-maxtimetoliveinminutes
        """
        pulumi.set(__self__, "allow_early_check_in", allow_early_check_in)
        pulumi.set(__self__, "max_time_to_live_in_minutes", max_time_to_live_in_minutes)

    @property
    @pulumi.getter(name="allowEarlyCheckIn")
    def allow_early_check_in(self) -> bool:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html#cfn-licensemanager-license-borrowconfiguration-allowearlycheckin
        """
        return pulumi.get(self, "allow_early_check_in")

    @property
    @pulumi.getter(name="maxTimeToLiveInMinutes")
    def max_time_to_live_in_minutes(self) -> int:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-borrowconfiguration.html#cfn-licensemanager-license-borrowconfiguration-maxtimetoliveinminutes
        """
        return pulumi.get(self, "max_time_to_live_in_minutes")


@pulumi.output_type
class LicenseConsumptionConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "borrowConfiguration":
            suggest = "borrow_configuration"
        elif key == "provisionalConfiguration":
            suggest = "provisional_configuration"
        elif key == "renewType":
            suggest = "renew_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LicenseConsumptionConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LicenseConsumptionConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LicenseConsumptionConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 borrow_configuration: Optional['outputs.LicenseBorrowConfiguration'] = None,
                 provisional_configuration: Optional['outputs.LicenseProvisionalConfiguration'] = None,
                 renew_type: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html
        :param 'LicenseBorrowConfiguration' borrow_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-borrowconfiguration
        :param 'LicenseProvisionalConfiguration' provisional_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-provisionalconfiguration
        :param str renew_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-renewtype
        """
        if borrow_configuration is not None:
            pulumi.set(__self__, "borrow_configuration", borrow_configuration)
        if provisional_configuration is not None:
            pulumi.set(__self__, "provisional_configuration", provisional_configuration)
        if renew_type is not None:
            pulumi.set(__self__, "renew_type", renew_type)

    @property
    @pulumi.getter(name="borrowConfiguration")
    def borrow_configuration(self) -> Optional['outputs.LicenseBorrowConfiguration']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-borrowconfiguration
        """
        return pulumi.get(self, "borrow_configuration")

    @property
    @pulumi.getter(name="provisionalConfiguration")
    def provisional_configuration(self) -> Optional['outputs.LicenseProvisionalConfiguration']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-provisionalconfiguration
        """
        return pulumi.get(self, "provisional_configuration")

    @property
    @pulumi.getter(name="renewType")
    def renew_type(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-renewtype
        """
        return pulumi.get(self, "renew_type")


@pulumi.output_type
class LicenseEntitlement(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowCheckIn":
            suggest = "allow_check_in"
        elif key == "maxCount":
            suggest = "max_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LicenseEntitlement. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LicenseEntitlement.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LicenseEntitlement.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 unit: str,
                 allow_check_in: Optional[bool] = None,
                 max_count: Optional[int] = None,
                 overage: Optional[bool] = None,
                 value: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html
        :param str name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-name
        :param str unit: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-unit
        :param bool allow_check_in: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-allowcheckin
        :param int max_count: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-maxcount
        :param bool overage: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-overage
        :param str value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-value
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "unit", unit)
        if allow_check_in is not None:
            pulumi.set(__self__, "allow_check_in", allow_check_in)
        if max_count is not None:
            pulumi.set(__self__, "max_count", max_count)
        if overage is not None:
            pulumi.set(__self__, "overage", overage)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def unit(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-unit
        """
        return pulumi.get(self, "unit")

    @property
    @pulumi.getter(name="allowCheckIn")
    def allow_check_in(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-allowcheckin
        """
        return pulumi.get(self, "allow_check_in")

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-maxcount
        """
        return pulumi.get(self, "max_count")

    @property
    @pulumi.getter
    def overage(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-overage
        """
        return pulumi.get(self, "overage")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-value
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class LicenseIssuerData(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-issuerdata.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "signKey":
            suggest = "sign_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LicenseIssuerData. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LicenseIssuerData.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LicenseIssuerData.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 sign_key: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-issuerdata.html
        :param str name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-issuerdata.html#cfn-licensemanager-license-issuerdata-name
        :param str sign_key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-issuerdata.html#cfn-licensemanager-license-issuerdata-signkey
        """
        pulumi.set(__self__, "name", name)
        if sign_key is not None:
            pulumi.set(__self__, "sign_key", sign_key)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-issuerdata.html#cfn-licensemanager-license-issuerdata-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="signKey")
    def sign_key(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-issuerdata.html#cfn-licensemanager-license-issuerdata-signkey
        """
        return pulumi.get(self, "sign_key")


@pulumi.output_type
class LicenseMetadata(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-metadata.html
    """
    def __init__(__self__, *,
                 name: str,
                 value: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-metadata.html
        :param str name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-metadata.html#cfn-licensemanager-license-metadata-name
        :param str value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-metadata.html#cfn-licensemanager-license-metadata-value
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-metadata.html#cfn-licensemanager-license-metadata-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-metadata.html#cfn-licensemanager-license-metadata-value
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class LicenseProvisionalConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-provisionalconfiguration.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxTimeToLiveInMinutes":
            suggest = "max_time_to_live_in_minutes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LicenseProvisionalConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LicenseProvisionalConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LicenseProvisionalConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_time_to_live_in_minutes: int):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-provisionalconfiguration.html
        :param int max_time_to_live_in_minutes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-provisionalconfiguration.html#cfn-licensemanager-license-provisionalconfiguration-maxtimetoliveinminutes
        """
        pulumi.set(__self__, "max_time_to_live_in_minutes", max_time_to_live_in_minutes)

    @property
    @pulumi.getter(name="maxTimeToLiveInMinutes")
    def max_time_to_live_in_minutes(self) -> int:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-provisionalconfiguration.html#cfn-licensemanager-license-provisionalconfiguration-maxtimetoliveinminutes
        """
        return pulumi.get(self, "max_time_to_live_in_minutes")


@pulumi.output_type
class LicenseValidityDateFormat(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html
    """
    def __init__(__self__, *,
                 begin: str,
                 end: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html
        :param str begin: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html#cfn-licensemanager-license-validitydateformat-begin
        :param str end: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html#cfn-licensemanager-license-validitydateformat-end
        """
        pulumi.set(__self__, "begin", begin)
        pulumi.set(__self__, "end", end)

    @property
    @pulumi.getter
    def begin(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html#cfn-licensemanager-license-validitydateformat-begin
        """
        return pulumi.get(self, "begin")

    @property
    @pulumi.getter
    def end(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html#cfn-licensemanager-license-validitydateformat-end
        """
        return pulumi.get(self, "end")


