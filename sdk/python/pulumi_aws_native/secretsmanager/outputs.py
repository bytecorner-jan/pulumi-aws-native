# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'RotationScheduleHostedRotationLambda',
    'RotationScheduleRotationRules',
    'SecretGenerateSecretString',
    'SecretReplicaRegion',
    'SecretTag',
]

@pulumi.output_type
class RotationScheduleHostedRotationLambda(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "rotationType":
            suggest = "rotation_type"
        elif key == "kmsKeyArn":
            suggest = "kms_key_arn"
        elif key == "masterSecretArn":
            suggest = "master_secret_arn"
        elif key == "masterSecretKmsKeyArn":
            suggest = "master_secret_kms_key_arn"
        elif key == "rotationLambdaName":
            suggest = "rotation_lambda_name"
        elif key == "vpcSecurityGroupIds":
            suggest = "vpc_security_group_ids"
        elif key == "vpcSubnetIds":
            suggest = "vpc_subnet_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RotationScheduleHostedRotationLambda. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RotationScheduleHostedRotationLambda.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RotationScheduleHostedRotationLambda.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 rotation_type: str,
                 kms_key_arn: Optional[str] = None,
                 master_secret_arn: Optional[str] = None,
                 master_secret_kms_key_arn: Optional[str] = None,
                 rotation_lambda_name: Optional[str] = None,
                 vpc_security_group_ids: Optional[str] = None,
                 vpc_subnet_ids: Optional[str] = None):
        pulumi.set(__self__, "rotation_type", rotation_type)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if master_secret_arn is not None:
            pulumi.set(__self__, "master_secret_arn", master_secret_arn)
        if master_secret_kms_key_arn is not None:
            pulumi.set(__self__, "master_secret_kms_key_arn", master_secret_kms_key_arn)
        if rotation_lambda_name is not None:
            pulumi.set(__self__, "rotation_lambda_name", rotation_lambda_name)
        if vpc_security_group_ids is not None:
            pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)
        if vpc_subnet_ids is not None:
            pulumi.set(__self__, "vpc_subnet_ids", vpc_subnet_ids)

    @property
    @pulumi.getter(name="rotationType")
    def rotation_type(self) -> str:
        return pulumi.get(self, "rotation_type")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[str]:
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter(name="masterSecretArn")
    def master_secret_arn(self) -> Optional[str]:
        return pulumi.get(self, "master_secret_arn")

    @property
    @pulumi.getter(name="masterSecretKmsKeyArn")
    def master_secret_kms_key_arn(self) -> Optional[str]:
        return pulumi.get(self, "master_secret_kms_key_arn")

    @property
    @pulumi.getter(name="rotationLambdaName")
    def rotation_lambda_name(self) -> Optional[str]:
        return pulumi.get(self, "rotation_lambda_name")

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> Optional[str]:
        return pulumi.get(self, "vpc_security_group_ids")

    @property
    @pulumi.getter(name="vpcSubnetIds")
    def vpc_subnet_ids(self) -> Optional[str]:
        return pulumi.get(self, "vpc_subnet_ids")


@pulumi.output_type
class RotationScheduleRotationRules(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "automaticallyAfterDays":
            suggest = "automatically_after_days"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RotationScheduleRotationRules. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RotationScheduleRotationRules.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RotationScheduleRotationRules.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 automatically_after_days: Optional[int] = None):
        if automatically_after_days is not None:
            pulumi.set(__self__, "automatically_after_days", automatically_after_days)

    @property
    @pulumi.getter(name="automaticallyAfterDays")
    def automatically_after_days(self) -> Optional[int]:
        return pulumi.get(self, "automatically_after_days")


@pulumi.output_type
class SecretGenerateSecretString(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "excludeCharacters":
            suggest = "exclude_characters"
        elif key == "excludeLowercase":
            suggest = "exclude_lowercase"
        elif key == "excludeNumbers":
            suggest = "exclude_numbers"
        elif key == "excludePunctuation":
            suggest = "exclude_punctuation"
        elif key == "excludeUppercase":
            suggest = "exclude_uppercase"
        elif key == "generateStringKey":
            suggest = "generate_string_key"
        elif key == "includeSpace":
            suggest = "include_space"
        elif key == "passwordLength":
            suggest = "password_length"
        elif key == "requireEachIncludedType":
            suggest = "require_each_included_type"
        elif key == "secretStringTemplate":
            suggest = "secret_string_template"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SecretGenerateSecretString. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SecretGenerateSecretString.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SecretGenerateSecretString.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 exclude_characters: Optional[str] = None,
                 exclude_lowercase: Optional[bool] = None,
                 exclude_numbers: Optional[bool] = None,
                 exclude_punctuation: Optional[bool] = None,
                 exclude_uppercase: Optional[bool] = None,
                 generate_string_key: Optional[str] = None,
                 include_space: Optional[bool] = None,
                 password_length: Optional[int] = None,
                 require_each_included_type: Optional[bool] = None,
                 secret_string_template: Optional[str] = None):
        if exclude_characters is not None:
            pulumi.set(__self__, "exclude_characters", exclude_characters)
        if exclude_lowercase is not None:
            pulumi.set(__self__, "exclude_lowercase", exclude_lowercase)
        if exclude_numbers is not None:
            pulumi.set(__self__, "exclude_numbers", exclude_numbers)
        if exclude_punctuation is not None:
            pulumi.set(__self__, "exclude_punctuation", exclude_punctuation)
        if exclude_uppercase is not None:
            pulumi.set(__self__, "exclude_uppercase", exclude_uppercase)
        if generate_string_key is not None:
            pulumi.set(__self__, "generate_string_key", generate_string_key)
        if include_space is not None:
            pulumi.set(__self__, "include_space", include_space)
        if password_length is not None:
            pulumi.set(__self__, "password_length", password_length)
        if require_each_included_type is not None:
            pulumi.set(__self__, "require_each_included_type", require_each_included_type)
        if secret_string_template is not None:
            pulumi.set(__self__, "secret_string_template", secret_string_template)

    @property
    @pulumi.getter(name="excludeCharacters")
    def exclude_characters(self) -> Optional[str]:
        return pulumi.get(self, "exclude_characters")

    @property
    @pulumi.getter(name="excludeLowercase")
    def exclude_lowercase(self) -> Optional[bool]:
        return pulumi.get(self, "exclude_lowercase")

    @property
    @pulumi.getter(name="excludeNumbers")
    def exclude_numbers(self) -> Optional[bool]:
        return pulumi.get(self, "exclude_numbers")

    @property
    @pulumi.getter(name="excludePunctuation")
    def exclude_punctuation(self) -> Optional[bool]:
        return pulumi.get(self, "exclude_punctuation")

    @property
    @pulumi.getter(name="excludeUppercase")
    def exclude_uppercase(self) -> Optional[bool]:
        return pulumi.get(self, "exclude_uppercase")

    @property
    @pulumi.getter(name="generateStringKey")
    def generate_string_key(self) -> Optional[str]:
        return pulumi.get(self, "generate_string_key")

    @property
    @pulumi.getter(name="includeSpace")
    def include_space(self) -> Optional[bool]:
        return pulumi.get(self, "include_space")

    @property
    @pulumi.getter(name="passwordLength")
    def password_length(self) -> Optional[int]:
        return pulumi.get(self, "password_length")

    @property
    @pulumi.getter(name="requireEachIncludedType")
    def require_each_included_type(self) -> Optional[bool]:
        return pulumi.get(self, "require_each_included_type")

    @property
    @pulumi.getter(name="secretStringTemplate")
    def secret_string_template(self) -> Optional[str]:
        return pulumi.get(self, "secret_string_template")


@pulumi.output_type
class SecretReplicaRegion(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "kmsKeyId":
            suggest = "kms_key_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SecretReplicaRegion. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SecretReplicaRegion.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SecretReplicaRegion.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 region: str,
                 kms_key_id: Optional[str] = None):
        pulumi.set(__self__, "region", region)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        return pulumi.get(self, "kms_key_id")


@pulumi.output_type
class SecretTag(dict):
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


