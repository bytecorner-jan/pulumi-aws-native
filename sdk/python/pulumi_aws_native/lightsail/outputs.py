# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'DatabaseRelationalDatabaseParameter',
    'DatabaseTag',
    'DiskAddOn',
    'DiskAutoSnapshotAddOn',
    'DiskLocation',
    'DiskTag',
    'InstanceAddOn',
    'InstanceAutoSnapshotAddOn',
    'InstanceDisk',
    'InstanceHardware',
    'InstanceLocation',
    'InstanceMonthlyTransfer',
    'InstanceNetworking',
    'InstancePort',
    'InstanceState',
    'InstanceTag',
]

@pulumi.output_type
class DatabaseRelationalDatabaseParameter(dict):
    """
    Describes the parameters of the database.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowedValues":
            suggest = "allowed_values"
        elif key == "applyMethod":
            suggest = "apply_method"
        elif key == "applyType":
            suggest = "apply_type"
        elif key == "dataType":
            suggest = "data_type"
        elif key == "isModifiable":
            suggest = "is_modifiable"
        elif key == "parameterName":
            suggest = "parameter_name"
        elif key == "parameterValue":
            suggest = "parameter_value"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DatabaseRelationalDatabaseParameter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DatabaseRelationalDatabaseParameter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DatabaseRelationalDatabaseParameter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allowed_values: Optional[str] = None,
                 apply_method: Optional[str] = None,
                 apply_type: Optional[str] = None,
                 data_type: Optional[str] = None,
                 description: Optional[str] = None,
                 is_modifiable: Optional[bool] = None,
                 parameter_name: Optional[str] = None,
                 parameter_value: Optional[str] = None):
        """
        Describes the parameters of the database.
        :param str allowed_values: Specifies the valid range of values for the parameter.
        :param str apply_method: Indicates when parameter updates are applied. Can be immediate or pending-reboot.
        :param str apply_type: Specifies the engine-specific parameter type.
        :param str data_type: Specifies the valid data type for the parameter.
        :param str description: Provides a description of the parameter.
        :param bool is_modifiable: A Boolean value indicating whether the parameter can be modified.
        :param str parameter_name: Specifies the name of the parameter.
        :param str parameter_value: Specifies the value of the parameter.
        """
        if allowed_values is not None:
            pulumi.set(__self__, "allowed_values", allowed_values)
        if apply_method is not None:
            pulumi.set(__self__, "apply_method", apply_method)
        if apply_type is not None:
            pulumi.set(__self__, "apply_type", apply_type)
        if data_type is not None:
            pulumi.set(__self__, "data_type", data_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_modifiable is not None:
            pulumi.set(__self__, "is_modifiable", is_modifiable)
        if parameter_name is not None:
            pulumi.set(__self__, "parameter_name", parameter_name)
        if parameter_value is not None:
            pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="allowedValues")
    def allowed_values(self) -> Optional[str]:
        """
        Specifies the valid range of values for the parameter.
        """
        return pulumi.get(self, "allowed_values")

    @property
    @pulumi.getter(name="applyMethod")
    def apply_method(self) -> Optional[str]:
        """
        Indicates when parameter updates are applied. Can be immediate or pending-reboot.
        """
        return pulumi.get(self, "apply_method")

    @property
    @pulumi.getter(name="applyType")
    def apply_type(self) -> Optional[str]:
        """
        Specifies the engine-specific parameter type.
        """
        return pulumi.get(self, "apply_type")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional[str]:
        """
        Specifies the valid data type for the parameter.
        """
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Provides a description of the parameter.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isModifiable")
    def is_modifiable(self) -> Optional[bool]:
        """
        A Boolean value indicating whether the parameter can be modified.
        """
        return pulumi.get(self, "is_modifiable")

    @property
    @pulumi.getter(name="parameterName")
    def parameter_name(self) -> Optional[str]:
        """
        Specifies the name of the parameter.
        """
        return pulumi.get(self, "parameter_name")

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> Optional[str]:
        """
        Specifies the value of the parameter.
        """
        return pulumi.get(self, "parameter_value")


@pulumi.output_type
class DatabaseTag(dict):
    """
    A key-value pair to associate with a resource.
    """
    def __init__(__self__, *,
                 key: str,
                 value: Optional[str] = None):
        """
        A key-value pair to associate with a resource.
        :param str key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param str value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class DiskAddOn(dict):
    """
    A addon associate with a resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "addOnType":
            suggest = "add_on_type"
        elif key == "autoSnapshotAddOnRequest":
            suggest = "auto_snapshot_add_on_request"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DiskAddOn. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DiskAddOn.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DiskAddOn.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 add_on_type: str,
                 auto_snapshot_add_on_request: Optional['outputs.DiskAutoSnapshotAddOn'] = None,
                 status: Optional['DiskAddOnStatus'] = None):
        """
        A addon associate with a resource.
        :param str add_on_type: The add-on type
        :param 'DiskAddOnStatus' status: Status of the Addon
        """
        pulumi.set(__self__, "add_on_type", add_on_type)
        if auto_snapshot_add_on_request is not None:
            pulumi.set(__self__, "auto_snapshot_add_on_request", auto_snapshot_add_on_request)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="addOnType")
    def add_on_type(self) -> str:
        """
        The add-on type
        """
        return pulumi.get(self, "add_on_type")

    @property
    @pulumi.getter(name="autoSnapshotAddOnRequest")
    def auto_snapshot_add_on_request(self) -> Optional['outputs.DiskAutoSnapshotAddOn']:
        return pulumi.get(self, "auto_snapshot_add_on_request")

    @property
    @pulumi.getter
    def status(self) -> Optional['DiskAddOnStatus']:
        """
        Status of the Addon
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class DiskAutoSnapshotAddOn(dict):
    """
    An object that represents additional parameters when enabling or modifying the automatic snapshot add-on
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "snapshotTimeOfDay":
            suggest = "snapshot_time_of_day"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DiskAutoSnapshotAddOn. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DiskAutoSnapshotAddOn.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DiskAutoSnapshotAddOn.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 snapshot_time_of_day: Optional[str] = None):
        """
        An object that represents additional parameters when enabling or modifying the automatic snapshot add-on
        :param str snapshot_time_of_day: The daily time when an automatic snapshot will be created.
        """
        if snapshot_time_of_day is not None:
            pulumi.set(__self__, "snapshot_time_of_day", snapshot_time_of_day)

    @property
    @pulumi.getter(name="snapshotTimeOfDay")
    def snapshot_time_of_day(self) -> Optional[str]:
        """
        The daily time when an automatic snapshot will be created.
        """
        return pulumi.get(self, "snapshot_time_of_day")


@pulumi.output_type
class DiskLocation(dict):
    """
    Location of a resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "availabilityZone":
            suggest = "availability_zone"
        elif key == "regionName":
            suggest = "region_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DiskLocation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DiskLocation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DiskLocation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 availability_zone: Optional[str] = None,
                 region_name: Optional[str] = None):
        """
        Location of a resource.
        :param str availability_zone: The Availability Zone in which to create your disk. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
        :param str region_name: The Region Name in which to create your disk.
        """
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        """
        The Availability Zone in which to create your disk. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[str]:
        """
        The Region Name in which to create your disk.
        """
        return pulumi.get(self, "region_name")


@pulumi.output_type
class DiskTag(dict):
    """
    A key-value pair to associate with a resource.
    """
    def __init__(__self__, *,
                 key: str,
                 value: Optional[str] = None):
        """
        A key-value pair to associate with a resource.
        :param str key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param str value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class InstanceAddOn(dict):
    """
    A addon associate with a resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "addOnType":
            suggest = "add_on_type"
        elif key == "autoSnapshotAddOnRequest":
            suggest = "auto_snapshot_add_on_request"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceAddOn. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceAddOn.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceAddOn.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 add_on_type: str,
                 auto_snapshot_add_on_request: Optional['outputs.InstanceAutoSnapshotAddOn'] = None,
                 status: Optional['InstanceAddOnStatus'] = None):
        """
        A addon associate with a resource.
        :param str add_on_type: The add-on type
        :param 'InstanceAddOnStatus' status: Status of the Addon
        """
        pulumi.set(__self__, "add_on_type", add_on_type)
        if auto_snapshot_add_on_request is not None:
            pulumi.set(__self__, "auto_snapshot_add_on_request", auto_snapshot_add_on_request)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="addOnType")
    def add_on_type(self) -> str:
        """
        The add-on type
        """
        return pulumi.get(self, "add_on_type")

    @property
    @pulumi.getter(name="autoSnapshotAddOnRequest")
    def auto_snapshot_add_on_request(self) -> Optional['outputs.InstanceAutoSnapshotAddOn']:
        return pulumi.get(self, "auto_snapshot_add_on_request")

    @property
    @pulumi.getter
    def status(self) -> Optional['InstanceAddOnStatus']:
        """
        Status of the Addon
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class InstanceAutoSnapshotAddOn(dict):
    """
    An object that represents additional parameters when enabling or modifying the automatic snapshot add-on
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "snapshotTimeOfDay":
            suggest = "snapshot_time_of_day"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceAutoSnapshotAddOn. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceAutoSnapshotAddOn.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceAutoSnapshotAddOn.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 snapshot_time_of_day: Optional[str] = None):
        """
        An object that represents additional parameters when enabling or modifying the automatic snapshot add-on
        :param str snapshot_time_of_day: The daily time when an automatic snapshot will be created.
        """
        if snapshot_time_of_day is not None:
            pulumi.set(__self__, "snapshot_time_of_day", snapshot_time_of_day)

    @property
    @pulumi.getter(name="snapshotTimeOfDay")
    def snapshot_time_of_day(self) -> Optional[str]:
        """
        The daily time when an automatic snapshot will be created.
        """
        return pulumi.get(self, "snapshot_time_of_day")


@pulumi.output_type
class InstanceDisk(dict):
    """
    Disk associated with the Instance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "diskName":
            suggest = "disk_name"
        elif key == "attachedTo":
            suggest = "attached_to"
        elif key == "attachmentState":
            suggest = "attachment_state"
        elif key == "iOPS":
            suggest = "i_ops"
        elif key == "isSystemDisk":
            suggest = "is_system_disk"
        elif key == "sizeInGb":
            suggest = "size_in_gb"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceDisk. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceDisk.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceDisk.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 disk_name: str,
                 path: str,
                 attached_to: Optional[str] = None,
                 attachment_state: Optional[str] = None,
                 i_ops: Optional[int] = None,
                 is_system_disk: Optional[bool] = None,
                 size_in_gb: Optional[str] = None):
        """
        Disk associated with the Instance.
        :param str disk_name: The names to use for your new Lightsail disk.
        :param str path: Path of the disk attached to the instance.
        :param str attached_to: Instance attached to the disk.
        :param str attachment_state: Attachment state of the disk.
        :param int i_ops: IOPS of disk.
        :param bool is_system_disk: Is the Attached disk is the system disk of the Instance.
        :param str size_in_gb: Size of the disk attached to the Instance.
        """
        pulumi.set(__self__, "disk_name", disk_name)
        pulumi.set(__self__, "path", path)
        if attached_to is not None:
            pulumi.set(__self__, "attached_to", attached_to)
        if attachment_state is not None:
            pulumi.set(__self__, "attachment_state", attachment_state)
        if i_ops is not None:
            pulumi.set(__self__, "i_ops", i_ops)
        if is_system_disk is not None:
            pulumi.set(__self__, "is_system_disk", is_system_disk)
        if size_in_gb is not None:
            pulumi.set(__self__, "size_in_gb", size_in_gb)

    @property
    @pulumi.getter(name="diskName")
    def disk_name(self) -> str:
        """
        The names to use for your new Lightsail disk.
        """
        return pulumi.get(self, "disk_name")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        Path of the disk attached to the instance.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="attachedTo")
    def attached_to(self) -> Optional[str]:
        """
        Instance attached to the disk.
        """
        return pulumi.get(self, "attached_to")

    @property
    @pulumi.getter(name="attachmentState")
    def attachment_state(self) -> Optional[str]:
        """
        Attachment state of the disk.
        """
        return pulumi.get(self, "attachment_state")

    @property
    @pulumi.getter(name="iOPS")
    def i_ops(self) -> Optional[int]:
        """
        IOPS of disk.
        """
        return pulumi.get(self, "i_ops")

    @property
    @pulumi.getter(name="isSystemDisk")
    def is_system_disk(self) -> Optional[bool]:
        """
        Is the Attached disk is the system disk of the Instance.
        """
        return pulumi.get(self, "is_system_disk")

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> Optional[str]:
        """
        Size of the disk attached to the Instance.
        """
        return pulumi.get(self, "size_in_gb")


@pulumi.output_type
class InstanceHardware(dict):
    """
    Hardware of the Instance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cpuCount":
            suggest = "cpu_count"
        elif key == "ramSizeInGb":
            suggest = "ram_size_in_gb"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceHardware. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceHardware.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceHardware.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cpu_count: Optional[int] = None,
                 disks: Optional[Sequence['outputs.InstanceDisk']] = None,
                 ram_size_in_gb: Optional[int] = None):
        """
        Hardware of the Instance.
        :param int cpu_count: CPU count of the Instance.
        :param Sequence['InstanceDisk'] disks: Disks attached to the Instance.
        :param int ram_size_in_gb: RAM Size of the Instance.
        """
        if cpu_count is not None:
            pulumi.set(__self__, "cpu_count", cpu_count)
        if disks is not None:
            pulumi.set(__self__, "disks", disks)
        if ram_size_in_gb is not None:
            pulumi.set(__self__, "ram_size_in_gb", ram_size_in_gb)

    @property
    @pulumi.getter(name="cpuCount")
    def cpu_count(self) -> Optional[int]:
        """
        CPU count of the Instance.
        """
        return pulumi.get(self, "cpu_count")

    @property
    @pulumi.getter
    def disks(self) -> Optional[Sequence['outputs.InstanceDisk']]:
        """
        Disks attached to the Instance.
        """
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter(name="ramSizeInGb")
    def ram_size_in_gb(self) -> Optional[int]:
        """
        RAM Size of the Instance.
        """
        return pulumi.get(self, "ram_size_in_gb")


@pulumi.output_type
class InstanceLocation(dict):
    """
    Location of a resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "availabilityZone":
            suggest = "availability_zone"
        elif key == "regionName":
            suggest = "region_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceLocation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceLocation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceLocation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 availability_zone: Optional[str] = None,
                 region_name: Optional[str] = None):
        """
        Location of a resource.
        :param str availability_zone: The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
        :param str region_name: The Region Name in which to create your instance.
        """
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        """
        The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[str]:
        """
        The Region Name in which to create your instance.
        """
        return pulumi.get(self, "region_name")


@pulumi.output_type
class InstanceMonthlyTransfer(dict):
    """
    Monthly Transfer of the Instance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "gbPerMonthAllocated":
            suggest = "gb_per_month_allocated"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceMonthlyTransfer. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceMonthlyTransfer.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceMonthlyTransfer.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 gb_per_month_allocated: Optional[str] = None):
        """
        Monthly Transfer of the Instance.
        :param str gb_per_month_allocated: GbPerMonthAllocated of the Instance.
        """
        if gb_per_month_allocated is not None:
            pulumi.set(__self__, "gb_per_month_allocated", gb_per_month_allocated)

    @property
    @pulumi.getter(name="gbPerMonthAllocated")
    def gb_per_month_allocated(self) -> Optional[str]:
        """
        GbPerMonthAllocated of the Instance.
        """
        return pulumi.get(self, "gb_per_month_allocated")


@pulumi.output_type
class InstanceNetworking(dict):
    """
    Networking of the Instance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "monthlyTransfer":
            suggest = "monthly_transfer"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceNetworking. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceNetworking.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceNetworking.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ports: Sequence['outputs.InstancePort'],
                 monthly_transfer: Optional['outputs.InstanceMonthlyTransfer'] = None):
        """
        Networking of the Instance.
        :param Sequence['InstancePort'] ports: Ports to the Instance.
        """
        pulumi.set(__self__, "ports", ports)
        if monthly_transfer is not None:
            pulumi.set(__self__, "monthly_transfer", monthly_transfer)

    @property
    @pulumi.getter
    def ports(self) -> Sequence['outputs.InstancePort']:
        """
        Ports to the Instance.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter(name="monthlyTransfer")
    def monthly_transfer(self) -> Optional['outputs.InstanceMonthlyTransfer']:
        return pulumi.get(self, "monthly_transfer")


@pulumi.output_type
class InstancePort(dict):
    """
    Port of the Instance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessDirection":
            suggest = "access_direction"
        elif key == "accessFrom":
            suggest = "access_from"
        elif key == "accessType":
            suggest = "access_type"
        elif key == "cidrListAliases":
            suggest = "cidr_list_aliases"
        elif key == "commonName":
            suggest = "common_name"
        elif key == "fromPort":
            suggest = "from_port"
        elif key == "ipv6Cidrs":
            suggest = "ipv6_cidrs"
        elif key == "toPort":
            suggest = "to_port"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstancePort. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstancePort.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstancePort.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_direction: Optional[str] = None,
                 access_from: Optional[str] = None,
                 access_type: Optional[str] = None,
                 cidr_list_aliases: Optional[Sequence[str]] = None,
                 cidrs: Optional[Sequence[str]] = None,
                 common_name: Optional[str] = None,
                 from_port: Optional[int] = None,
                 ipv6_cidrs: Optional[Sequence[str]] = None,
                 protocol: Optional[str] = None,
                 to_port: Optional[int] = None):
        """
        Port of the Instance.
        :param str access_direction: Access Direction for Protocol of the Instance(inbound/outbound).
        :param str access_from: Access From Protocol of the Instance.
        :param str access_type: Access Type Protocol of the Instance.
        :param str common_name: CommonName for Protocol of the Instance.
        :param int from_port: From Port of the Instance.
        :param str protocol: Port Protocol of the Instance.
        :param int to_port: To Port of the Instance.
        """
        if access_direction is not None:
            pulumi.set(__self__, "access_direction", access_direction)
        if access_from is not None:
            pulumi.set(__self__, "access_from", access_from)
        if access_type is not None:
            pulumi.set(__self__, "access_type", access_type)
        if cidr_list_aliases is not None:
            pulumi.set(__self__, "cidr_list_aliases", cidr_list_aliases)
        if cidrs is not None:
            pulumi.set(__self__, "cidrs", cidrs)
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if from_port is not None:
            pulumi.set(__self__, "from_port", from_port)
        if ipv6_cidrs is not None:
            pulumi.set(__self__, "ipv6_cidrs", ipv6_cidrs)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if to_port is not None:
            pulumi.set(__self__, "to_port", to_port)

    @property
    @pulumi.getter(name="accessDirection")
    def access_direction(self) -> Optional[str]:
        """
        Access Direction for Protocol of the Instance(inbound/outbound).
        """
        return pulumi.get(self, "access_direction")

    @property
    @pulumi.getter(name="accessFrom")
    def access_from(self) -> Optional[str]:
        """
        Access From Protocol of the Instance.
        """
        return pulumi.get(self, "access_from")

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> Optional[str]:
        """
        Access Type Protocol of the Instance.
        """
        return pulumi.get(self, "access_type")

    @property
    @pulumi.getter(name="cidrListAliases")
    def cidr_list_aliases(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "cidr_list_aliases")

    @property
    @pulumi.getter
    def cidrs(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "cidrs")

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[str]:
        """
        CommonName for Protocol of the Instance.
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> Optional[int]:
        """
        From Port of the Instance.
        """
        return pulumi.get(self, "from_port")

    @property
    @pulumi.getter(name="ipv6Cidrs")
    def ipv6_cidrs(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ipv6_cidrs")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        """
        Port Protocol of the Instance.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> Optional[int]:
        """
        To Port of the Instance.
        """
        return pulumi.get(self, "to_port")


@pulumi.output_type
class InstanceState(dict):
    """
    Current State of the Instance.
    """
    def __init__(__self__, *,
                 code: Optional[int] = None,
                 name: Optional[str] = None):
        """
        Current State of the Instance.
        :param int code: Status code of the Instance.
        :param str name: Status code of the Instance.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def code(self) -> Optional[int]:
        """
        Status code of the Instance.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Status code of the Instance.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class InstanceTag(dict):
    """
    A key-value pair to associate with a resource.
    """
    def __init__(__self__, *,
                 key: str,
                 value: Optional[str] = None):
        """
        A key-value pair to associate with a resource.
        :param str key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param str value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")


