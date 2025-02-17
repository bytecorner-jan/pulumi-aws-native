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
    'FileSystemAuditLogConfiguration',
    'FileSystemLustreConfiguration',
    'FileSystemSelfManagedActiveDirectoryConfiguration',
    'FileSystemTag',
    'FileSystemWindowsConfiguration',
]

@pulumi.output_type
class FileSystemAuditLogConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fileAccessAuditLogLevel":
            suggest = "file_access_audit_log_level"
        elif key == "fileShareAccessAuditLogLevel":
            suggest = "file_share_access_audit_log_level"
        elif key == "auditLogDestination":
            suggest = "audit_log_destination"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FileSystemAuditLogConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FileSystemAuditLogConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FileSystemAuditLogConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 file_access_audit_log_level: str,
                 file_share_access_audit_log_level: str,
                 audit_log_destination: Optional[str] = None):
        pulumi.set(__self__, "file_access_audit_log_level", file_access_audit_log_level)
        pulumi.set(__self__, "file_share_access_audit_log_level", file_share_access_audit_log_level)
        if audit_log_destination is not None:
            pulumi.set(__self__, "audit_log_destination", audit_log_destination)

    @property
    @pulumi.getter(name="fileAccessAuditLogLevel")
    def file_access_audit_log_level(self) -> str:
        return pulumi.get(self, "file_access_audit_log_level")

    @property
    @pulumi.getter(name="fileShareAccessAuditLogLevel")
    def file_share_access_audit_log_level(self) -> str:
        return pulumi.get(self, "file_share_access_audit_log_level")

    @property
    @pulumi.getter(name="auditLogDestination")
    def audit_log_destination(self) -> Optional[str]:
        return pulumi.get(self, "audit_log_destination")


@pulumi.output_type
class FileSystemLustreConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "autoImportPolicy":
            suggest = "auto_import_policy"
        elif key == "automaticBackupRetentionDays":
            suggest = "automatic_backup_retention_days"
        elif key == "copyTagsToBackups":
            suggest = "copy_tags_to_backups"
        elif key == "dailyAutomaticBackupStartTime":
            suggest = "daily_automatic_backup_start_time"
        elif key == "dataCompressionType":
            suggest = "data_compression_type"
        elif key == "deploymentType":
            suggest = "deployment_type"
        elif key == "driveCacheType":
            suggest = "drive_cache_type"
        elif key == "exportPath":
            suggest = "export_path"
        elif key == "importPath":
            suggest = "import_path"
        elif key == "importedFileChunkSize":
            suggest = "imported_file_chunk_size"
        elif key == "perUnitStorageThroughput":
            suggest = "per_unit_storage_throughput"
        elif key == "weeklyMaintenanceStartTime":
            suggest = "weekly_maintenance_start_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FileSystemLustreConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FileSystemLustreConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FileSystemLustreConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 auto_import_policy: Optional[str] = None,
                 automatic_backup_retention_days: Optional[int] = None,
                 copy_tags_to_backups: Optional[bool] = None,
                 daily_automatic_backup_start_time: Optional[str] = None,
                 data_compression_type: Optional[str] = None,
                 deployment_type: Optional[str] = None,
                 drive_cache_type: Optional[str] = None,
                 export_path: Optional[str] = None,
                 import_path: Optional[str] = None,
                 imported_file_chunk_size: Optional[int] = None,
                 per_unit_storage_throughput: Optional[int] = None,
                 weekly_maintenance_start_time: Optional[str] = None):
        if auto_import_policy is not None:
            pulumi.set(__self__, "auto_import_policy", auto_import_policy)
        if automatic_backup_retention_days is not None:
            pulumi.set(__self__, "automatic_backup_retention_days", automatic_backup_retention_days)
        if copy_tags_to_backups is not None:
            pulumi.set(__self__, "copy_tags_to_backups", copy_tags_to_backups)
        if daily_automatic_backup_start_time is not None:
            pulumi.set(__self__, "daily_automatic_backup_start_time", daily_automatic_backup_start_time)
        if data_compression_type is not None:
            pulumi.set(__self__, "data_compression_type", data_compression_type)
        if deployment_type is not None:
            pulumi.set(__self__, "deployment_type", deployment_type)
        if drive_cache_type is not None:
            pulumi.set(__self__, "drive_cache_type", drive_cache_type)
        if export_path is not None:
            pulumi.set(__self__, "export_path", export_path)
        if import_path is not None:
            pulumi.set(__self__, "import_path", import_path)
        if imported_file_chunk_size is not None:
            pulumi.set(__self__, "imported_file_chunk_size", imported_file_chunk_size)
        if per_unit_storage_throughput is not None:
            pulumi.set(__self__, "per_unit_storage_throughput", per_unit_storage_throughput)
        if weekly_maintenance_start_time is not None:
            pulumi.set(__self__, "weekly_maintenance_start_time", weekly_maintenance_start_time)

    @property
    @pulumi.getter(name="autoImportPolicy")
    def auto_import_policy(self) -> Optional[str]:
        return pulumi.get(self, "auto_import_policy")

    @property
    @pulumi.getter(name="automaticBackupRetentionDays")
    def automatic_backup_retention_days(self) -> Optional[int]:
        return pulumi.get(self, "automatic_backup_retention_days")

    @property
    @pulumi.getter(name="copyTagsToBackups")
    def copy_tags_to_backups(self) -> Optional[bool]:
        return pulumi.get(self, "copy_tags_to_backups")

    @property
    @pulumi.getter(name="dailyAutomaticBackupStartTime")
    def daily_automatic_backup_start_time(self) -> Optional[str]:
        return pulumi.get(self, "daily_automatic_backup_start_time")

    @property
    @pulumi.getter(name="dataCompressionType")
    def data_compression_type(self) -> Optional[str]:
        return pulumi.get(self, "data_compression_type")

    @property
    @pulumi.getter(name="deploymentType")
    def deployment_type(self) -> Optional[str]:
        return pulumi.get(self, "deployment_type")

    @property
    @pulumi.getter(name="driveCacheType")
    def drive_cache_type(self) -> Optional[str]:
        return pulumi.get(self, "drive_cache_type")

    @property
    @pulumi.getter(name="exportPath")
    def export_path(self) -> Optional[str]:
        return pulumi.get(self, "export_path")

    @property
    @pulumi.getter(name="importPath")
    def import_path(self) -> Optional[str]:
        return pulumi.get(self, "import_path")

    @property
    @pulumi.getter(name="importedFileChunkSize")
    def imported_file_chunk_size(self) -> Optional[int]:
        return pulumi.get(self, "imported_file_chunk_size")

    @property
    @pulumi.getter(name="perUnitStorageThroughput")
    def per_unit_storage_throughput(self) -> Optional[int]:
        return pulumi.get(self, "per_unit_storage_throughput")

    @property
    @pulumi.getter(name="weeklyMaintenanceStartTime")
    def weekly_maintenance_start_time(self) -> Optional[str]:
        return pulumi.get(self, "weekly_maintenance_start_time")


@pulumi.output_type
class FileSystemSelfManagedActiveDirectoryConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dnsIps":
            suggest = "dns_ips"
        elif key == "domainName":
            suggest = "domain_name"
        elif key == "fileSystemAdministratorsGroup":
            suggest = "file_system_administrators_group"
        elif key == "organizationalUnitDistinguishedName":
            suggest = "organizational_unit_distinguished_name"
        elif key == "userName":
            suggest = "user_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FileSystemSelfManagedActiveDirectoryConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FileSystemSelfManagedActiveDirectoryConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FileSystemSelfManagedActiveDirectoryConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dns_ips: Optional[Sequence[str]] = None,
                 domain_name: Optional[str] = None,
                 file_system_administrators_group: Optional[str] = None,
                 organizational_unit_distinguished_name: Optional[str] = None,
                 password: Optional[str] = None,
                 user_name: Optional[str] = None):
        if dns_ips is not None:
            pulumi.set(__self__, "dns_ips", dns_ips)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if file_system_administrators_group is not None:
            pulumi.set(__self__, "file_system_administrators_group", file_system_administrators_group)
        if organizational_unit_distinguished_name is not None:
            pulumi.set(__self__, "organizational_unit_distinguished_name", organizational_unit_distinguished_name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="dnsIps")
    def dns_ips(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "dns_ips")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[str]:
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="fileSystemAdministratorsGroup")
    def file_system_administrators_group(self) -> Optional[str]:
        return pulumi.get(self, "file_system_administrators_group")

    @property
    @pulumi.getter(name="organizationalUnitDistinguishedName")
    def organizational_unit_distinguished_name(self) -> Optional[str]:
        return pulumi.get(self, "organizational_unit_distinguished_name")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[str]:
        return pulumi.get(self, "user_name")


@pulumi.output_type
class FileSystemTag(dict):
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
class FileSystemWindowsConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "throughputCapacity":
            suggest = "throughput_capacity"
        elif key == "activeDirectoryId":
            suggest = "active_directory_id"
        elif key == "auditLogConfiguration":
            suggest = "audit_log_configuration"
        elif key == "automaticBackupRetentionDays":
            suggest = "automatic_backup_retention_days"
        elif key == "copyTagsToBackups":
            suggest = "copy_tags_to_backups"
        elif key == "dailyAutomaticBackupStartTime":
            suggest = "daily_automatic_backup_start_time"
        elif key == "deploymentType":
            suggest = "deployment_type"
        elif key == "preferredSubnetId":
            suggest = "preferred_subnet_id"
        elif key == "selfManagedActiveDirectoryConfiguration":
            suggest = "self_managed_active_directory_configuration"
        elif key == "weeklyMaintenanceStartTime":
            suggest = "weekly_maintenance_start_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FileSystemWindowsConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FileSystemWindowsConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FileSystemWindowsConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 throughput_capacity: int,
                 active_directory_id: Optional[str] = None,
                 aliases: Optional[Sequence[str]] = None,
                 audit_log_configuration: Optional['outputs.FileSystemAuditLogConfiguration'] = None,
                 automatic_backup_retention_days: Optional[int] = None,
                 copy_tags_to_backups: Optional[bool] = None,
                 daily_automatic_backup_start_time: Optional[str] = None,
                 deployment_type: Optional[str] = None,
                 preferred_subnet_id: Optional[str] = None,
                 self_managed_active_directory_configuration: Optional['outputs.FileSystemSelfManagedActiveDirectoryConfiguration'] = None,
                 weekly_maintenance_start_time: Optional[str] = None):
        pulumi.set(__self__, "throughput_capacity", throughput_capacity)
        if active_directory_id is not None:
            pulumi.set(__self__, "active_directory_id", active_directory_id)
        if aliases is not None:
            pulumi.set(__self__, "aliases", aliases)
        if audit_log_configuration is not None:
            pulumi.set(__self__, "audit_log_configuration", audit_log_configuration)
        if automatic_backup_retention_days is not None:
            pulumi.set(__self__, "automatic_backup_retention_days", automatic_backup_retention_days)
        if copy_tags_to_backups is not None:
            pulumi.set(__self__, "copy_tags_to_backups", copy_tags_to_backups)
        if daily_automatic_backup_start_time is not None:
            pulumi.set(__self__, "daily_automatic_backup_start_time", daily_automatic_backup_start_time)
        if deployment_type is not None:
            pulumi.set(__self__, "deployment_type", deployment_type)
        if preferred_subnet_id is not None:
            pulumi.set(__self__, "preferred_subnet_id", preferred_subnet_id)
        if self_managed_active_directory_configuration is not None:
            pulumi.set(__self__, "self_managed_active_directory_configuration", self_managed_active_directory_configuration)
        if weekly_maintenance_start_time is not None:
            pulumi.set(__self__, "weekly_maintenance_start_time", weekly_maintenance_start_time)

    @property
    @pulumi.getter(name="throughputCapacity")
    def throughput_capacity(self) -> int:
        return pulumi.get(self, "throughput_capacity")

    @property
    @pulumi.getter(name="activeDirectoryId")
    def active_directory_id(self) -> Optional[str]:
        return pulumi.get(self, "active_directory_id")

    @property
    @pulumi.getter
    def aliases(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "aliases")

    @property
    @pulumi.getter(name="auditLogConfiguration")
    def audit_log_configuration(self) -> Optional['outputs.FileSystemAuditLogConfiguration']:
        return pulumi.get(self, "audit_log_configuration")

    @property
    @pulumi.getter(name="automaticBackupRetentionDays")
    def automatic_backup_retention_days(self) -> Optional[int]:
        return pulumi.get(self, "automatic_backup_retention_days")

    @property
    @pulumi.getter(name="copyTagsToBackups")
    def copy_tags_to_backups(self) -> Optional[bool]:
        return pulumi.get(self, "copy_tags_to_backups")

    @property
    @pulumi.getter(name="dailyAutomaticBackupStartTime")
    def daily_automatic_backup_start_time(self) -> Optional[str]:
        return pulumi.get(self, "daily_automatic_backup_start_time")

    @property
    @pulumi.getter(name="deploymentType")
    def deployment_type(self) -> Optional[str]:
        return pulumi.get(self, "deployment_type")

    @property
    @pulumi.getter(name="preferredSubnetId")
    def preferred_subnet_id(self) -> Optional[str]:
        return pulumi.get(self, "preferred_subnet_id")

    @property
    @pulumi.getter(name="selfManagedActiveDirectoryConfiguration")
    def self_managed_active_directory_configuration(self) -> Optional['outputs.FileSystemSelfManagedActiveDirectoryConfiguration']:
        return pulumi.get(self, "self_managed_active_directory_configuration")

    @property
    @pulumi.getter(name="weeklyMaintenanceStartTime")
    def weekly_maintenance_start_time(self) -> Optional[str]:
        return pulumi.get(self, "weekly_maintenance_start_time")


