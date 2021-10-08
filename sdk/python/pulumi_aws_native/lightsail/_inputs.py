# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'DiskAddOnArgs',
    'DiskAutoSnapshotAddOnArgs',
    'DiskTagArgs',
    'InstanceAddOnArgs',
    'InstanceAutoSnapshotAddOnArgs',
    'InstanceDiskArgs',
    'InstanceHardwareArgs',
    'InstanceLocationArgs',
    'InstanceMonthlyTransferArgs',
    'InstanceNetworkingArgs',
    'InstancePortArgs',
    'InstanceStateArgs',
    'InstanceTagArgs',
]

@pulumi.input_type
class DiskAddOnArgs:
    def __init__(__self__, *,
                 add_on_type: pulumi.Input[str],
                 auto_snapshot_add_on_request: Optional[pulumi.Input['DiskAutoSnapshotAddOnArgs']] = None,
                 status: Optional[pulumi.Input['DiskAddOnStatus']] = None):
        """
        A addon associate with a resource.
        :param pulumi.Input[str] add_on_type: The add-on type
        :param pulumi.Input['DiskAddOnStatus'] status: Status of the Addon
        """
        pulumi.set(__self__, "add_on_type", add_on_type)
        if auto_snapshot_add_on_request is not None:
            pulumi.set(__self__, "auto_snapshot_add_on_request", auto_snapshot_add_on_request)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="addOnType")
    def add_on_type(self) -> pulumi.Input[str]:
        """
        The add-on type
        """
        return pulumi.get(self, "add_on_type")

    @add_on_type.setter
    def add_on_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "add_on_type", value)

    @property
    @pulumi.getter(name="autoSnapshotAddOnRequest")
    def auto_snapshot_add_on_request(self) -> Optional[pulumi.Input['DiskAutoSnapshotAddOnArgs']]:
        return pulumi.get(self, "auto_snapshot_add_on_request")

    @auto_snapshot_add_on_request.setter
    def auto_snapshot_add_on_request(self, value: Optional[pulumi.Input['DiskAutoSnapshotAddOnArgs']]):
        pulumi.set(self, "auto_snapshot_add_on_request", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['DiskAddOnStatus']]:
        """
        Status of the Addon
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['DiskAddOnStatus']]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class DiskAutoSnapshotAddOnArgs:
    def __init__(__self__, *,
                 snapshot_time_of_day: Optional[pulumi.Input[str]] = None):
        """
        An object that represents additional parameters when enabling or modifying the automatic snapshot add-on
        :param pulumi.Input[str] snapshot_time_of_day: The daily time when an automatic snapshot will be created.
        """
        if snapshot_time_of_day is not None:
            pulumi.set(__self__, "snapshot_time_of_day", snapshot_time_of_day)

    @property
    @pulumi.getter(name="snapshotTimeOfDay")
    def snapshot_time_of_day(self) -> Optional[pulumi.Input[str]]:
        """
        The daily time when an automatic snapshot will be created.
        """
        return pulumi.get(self, "snapshot_time_of_day")

    @snapshot_time_of_day.setter
    def snapshot_time_of_day(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_time_of_day", value)


@pulumi.input_type
class DiskTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: Optional[pulumi.Input[str]] = None):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class InstanceAddOnArgs:
    def __init__(__self__, *,
                 add_on_type: pulumi.Input[str],
                 auto_snapshot_add_on_request: Optional[pulumi.Input['InstanceAutoSnapshotAddOnArgs']] = None,
                 status: Optional[pulumi.Input['InstanceAddOnStatus']] = None):
        """
        A addon associate with a resource.
        :param pulumi.Input[str] add_on_type: The add-on type
        :param pulumi.Input['InstanceAddOnStatus'] status: Status of the Addon
        """
        pulumi.set(__self__, "add_on_type", add_on_type)
        if auto_snapshot_add_on_request is not None:
            pulumi.set(__self__, "auto_snapshot_add_on_request", auto_snapshot_add_on_request)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="addOnType")
    def add_on_type(self) -> pulumi.Input[str]:
        """
        The add-on type
        """
        return pulumi.get(self, "add_on_type")

    @add_on_type.setter
    def add_on_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "add_on_type", value)

    @property
    @pulumi.getter(name="autoSnapshotAddOnRequest")
    def auto_snapshot_add_on_request(self) -> Optional[pulumi.Input['InstanceAutoSnapshotAddOnArgs']]:
        return pulumi.get(self, "auto_snapshot_add_on_request")

    @auto_snapshot_add_on_request.setter
    def auto_snapshot_add_on_request(self, value: Optional[pulumi.Input['InstanceAutoSnapshotAddOnArgs']]):
        pulumi.set(self, "auto_snapshot_add_on_request", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['InstanceAddOnStatus']]:
        """
        Status of the Addon
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['InstanceAddOnStatus']]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class InstanceAutoSnapshotAddOnArgs:
    def __init__(__self__, *,
                 snapshot_time_of_day: Optional[pulumi.Input[str]] = None):
        """
        An object that represents additional parameters when enabling or modifying the automatic snapshot add-on
        :param pulumi.Input[str] snapshot_time_of_day: The daily time when an automatic snapshot will be created.
        """
        if snapshot_time_of_day is not None:
            pulumi.set(__self__, "snapshot_time_of_day", snapshot_time_of_day)

    @property
    @pulumi.getter(name="snapshotTimeOfDay")
    def snapshot_time_of_day(self) -> Optional[pulumi.Input[str]]:
        """
        The daily time when an automatic snapshot will be created.
        """
        return pulumi.get(self, "snapshot_time_of_day")

    @snapshot_time_of_day.setter
    def snapshot_time_of_day(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_time_of_day", value)


@pulumi.input_type
class InstanceDiskArgs:
    def __init__(__self__, *,
                 disk_name: pulumi.Input[str],
                 path: pulumi.Input[str],
                 attached_to: Optional[pulumi.Input[str]] = None,
                 attachment_state: Optional[pulumi.Input[str]] = None,
                 i_ops: Optional[pulumi.Input[int]] = None,
                 is_system_disk: Optional[pulumi.Input[bool]] = None,
                 size_in_gb: Optional[pulumi.Input[str]] = None):
        """
        Disk associated with the Instance.
        :param pulumi.Input[str] disk_name: The names to use for your new Lightsail disk.
        :param pulumi.Input[str] path: Path of the disk attached to the instance.
        :param pulumi.Input[str] attached_to: Instance attached to the disk.
        :param pulumi.Input[str] attachment_state: Attachment state of the disk.
        :param pulumi.Input[int] i_ops: IOPS of disk.
        :param pulumi.Input[bool] is_system_disk: Is the Attached disk is the system disk of the Instance.
        :param pulumi.Input[str] size_in_gb: Size of the disk attached to the Instance.
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
    def disk_name(self) -> pulumi.Input[str]:
        """
        The names to use for your new Lightsail disk.
        """
        return pulumi.get(self, "disk_name")

    @disk_name.setter
    def disk_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_name", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        Path of the disk attached to the instance.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="attachedTo")
    def attached_to(self) -> Optional[pulumi.Input[str]]:
        """
        Instance attached to the disk.
        """
        return pulumi.get(self, "attached_to")

    @attached_to.setter
    def attached_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attached_to", value)

    @property
    @pulumi.getter(name="attachmentState")
    def attachment_state(self) -> Optional[pulumi.Input[str]]:
        """
        Attachment state of the disk.
        """
        return pulumi.get(self, "attachment_state")

    @attachment_state.setter
    def attachment_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attachment_state", value)

    @property
    @pulumi.getter(name="iOPS")
    def i_ops(self) -> Optional[pulumi.Input[int]]:
        """
        IOPS of disk.
        """
        return pulumi.get(self, "i_ops")

    @i_ops.setter
    def i_ops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "i_ops", value)

    @property
    @pulumi.getter(name="isSystemDisk")
    def is_system_disk(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the Attached disk is the system disk of the Instance.
        """
        return pulumi.get(self, "is_system_disk")

    @is_system_disk.setter
    def is_system_disk(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_system_disk", value)

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> Optional[pulumi.Input[str]]:
        """
        Size of the disk attached to the Instance.
        """
        return pulumi.get(self, "size_in_gb")

    @size_in_gb.setter
    def size_in_gb(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "size_in_gb", value)


@pulumi.input_type
class InstanceHardwareArgs:
    def __init__(__self__, *,
                 cpu_count: Optional[pulumi.Input[int]] = None,
                 disks: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceDiskArgs']]]] = None,
                 ram_size_in_gb: Optional[pulumi.Input[int]] = None):
        """
        Hardware of the Instance.
        :param pulumi.Input[int] cpu_count: CPU count of the Instance.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceDiskArgs']]] disks: Disks attached to the Instance.
        :param pulumi.Input[int] ram_size_in_gb: RAM Size of the Instance.
        """
        if cpu_count is not None:
            pulumi.set(__self__, "cpu_count", cpu_count)
        if disks is not None:
            pulumi.set(__self__, "disks", disks)
        if ram_size_in_gb is not None:
            pulumi.set(__self__, "ram_size_in_gb", ram_size_in_gb)

    @property
    @pulumi.getter(name="cpuCount")
    def cpu_count(self) -> Optional[pulumi.Input[int]]:
        """
        CPU count of the Instance.
        """
        return pulumi.get(self, "cpu_count")

    @cpu_count.setter
    def cpu_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cpu_count", value)

    @property
    @pulumi.getter
    def disks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceDiskArgs']]]]:
        """
        Disks attached to the Instance.
        """
        return pulumi.get(self, "disks")

    @disks.setter
    def disks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceDiskArgs']]]]):
        pulumi.set(self, "disks", value)

    @property
    @pulumi.getter(name="ramSizeInGb")
    def ram_size_in_gb(self) -> Optional[pulumi.Input[int]]:
        """
        RAM Size of the Instance.
        """
        return pulumi.get(self, "ram_size_in_gb")

    @ram_size_in_gb.setter
    def ram_size_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ram_size_in_gb", value)


@pulumi.input_type
class InstanceLocationArgs:
    def __init__(__self__, *,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None):
        """
        Location of a resource.
        :param pulumi.Input[str] availability_zone: The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
        :param pulumi.Input[str] region_name: The Region Name in which to create your instance.
        """
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Region Name in which to create your instance.
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_name", value)


@pulumi.input_type
class InstanceMonthlyTransferArgs:
    def __init__(__self__, *,
                 gb_per_month_allocated: Optional[pulumi.Input[str]] = None):
        """
        Monthly Transfer of the Instance.
        :param pulumi.Input[str] gb_per_month_allocated: GbPerMonthAllocated of the Instance.
        """
        if gb_per_month_allocated is not None:
            pulumi.set(__self__, "gb_per_month_allocated", gb_per_month_allocated)

    @property
    @pulumi.getter(name="gbPerMonthAllocated")
    def gb_per_month_allocated(self) -> Optional[pulumi.Input[str]]:
        """
        GbPerMonthAllocated of the Instance.
        """
        return pulumi.get(self, "gb_per_month_allocated")

    @gb_per_month_allocated.setter
    def gb_per_month_allocated(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gb_per_month_allocated", value)


@pulumi.input_type
class InstanceNetworkingArgs:
    def __init__(__self__, *,
                 ports: pulumi.Input[Sequence[pulumi.Input['InstancePortArgs']]],
                 monthly_transfer: Optional[pulumi.Input['InstanceMonthlyTransferArgs']] = None):
        """
        Networking of the Instance.
        :param pulumi.Input[Sequence[pulumi.Input['InstancePortArgs']]] ports: Ports to the Instance.
        """
        pulumi.set(__self__, "ports", ports)
        if monthly_transfer is not None:
            pulumi.set(__self__, "monthly_transfer", monthly_transfer)

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Input[Sequence[pulumi.Input['InstancePortArgs']]]:
        """
        Ports to the Instance.
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: pulumi.Input[Sequence[pulumi.Input['InstancePortArgs']]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter(name="monthlyTransfer")
    def monthly_transfer(self) -> Optional[pulumi.Input['InstanceMonthlyTransferArgs']]:
        return pulumi.get(self, "monthly_transfer")

    @monthly_transfer.setter
    def monthly_transfer(self, value: Optional[pulumi.Input['InstanceMonthlyTransferArgs']]):
        pulumi.set(self, "monthly_transfer", value)


@pulumi.input_type
class InstancePortArgs:
    def __init__(__self__, *,
                 access_direction: Optional[pulumi.Input[str]] = None,
                 access_from: Optional[pulumi.Input[str]] = None,
                 access_type: Optional[pulumi.Input[str]] = None,
                 cidr_list_aliases: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 from_port: Optional[pulumi.Input[int]] = None,
                 ipv6_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 to_port: Optional[pulumi.Input[int]] = None):
        """
        Port of the Instance.
        :param pulumi.Input[str] access_direction: Access Direction for Protocol of the Instance(inbound/outbound).
        :param pulumi.Input[str] access_from: Access From Protocol of the Instance.
        :param pulumi.Input[str] access_type: Access Type Protocol of the Instance.
        :param pulumi.Input[str] common_name: CommonName for Protocol of the Instance.
        :param pulumi.Input[int] from_port: From Port of the Instance.
        :param pulumi.Input[str] protocol: Port Protocol of the Instance.
        :param pulumi.Input[int] to_port: To Port of the Instance.
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
    def access_direction(self) -> Optional[pulumi.Input[str]]:
        """
        Access Direction for Protocol of the Instance(inbound/outbound).
        """
        return pulumi.get(self, "access_direction")

    @access_direction.setter
    def access_direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_direction", value)

    @property
    @pulumi.getter(name="accessFrom")
    def access_from(self) -> Optional[pulumi.Input[str]]:
        """
        Access From Protocol of the Instance.
        """
        return pulumi.get(self, "access_from")

    @access_from.setter
    def access_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_from", value)

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> Optional[pulumi.Input[str]]:
        """
        Access Type Protocol of the Instance.
        """
        return pulumi.get(self, "access_type")

    @access_type.setter
    def access_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_type", value)

    @property
    @pulumi.getter(name="cidrListAliases")
    def cidr_list_aliases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "cidr_list_aliases")

    @cidr_list_aliases.setter
    def cidr_list_aliases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidr_list_aliases", value)

    @property
    @pulumi.getter
    def cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "cidrs")

    @cidrs.setter
    def cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidrs", value)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[pulumi.Input[str]]:
        """
        CommonName for Protocol of the Instance.
        """
        return pulumi.get(self, "common_name")

    @common_name.setter
    def common_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "common_name", value)

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> Optional[pulumi.Input[int]]:
        """
        From Port of the Instance.
        """
        return pulumi.get(self, "from_port")

    @from_port.setter
    def from_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "from_port", value)

    @property
    @pulumi.getter(name="ipv6Cidrs")
    def ipv6_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "ipv6_cidrs")

    @ipv6_cidrs.setter
    def ipv6_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ipv6_cidrs", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Port Protocol of the Instance.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> Optional[pulumi.Input[int]]:
        """
        To Port of the Instance.
        """
        return pulumi.get(self, "to_port")

    @to_port.setter
    def to_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "to_port", value)


@pulumi.input_type
class InstanceStateArgs:
    def __init__(__self__, *,
                 code: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Current State of the Instance.
        :param pulumi.Input[int] code: Status code of the Instance.
        :param pulumi.Input[str] name: Status code of the Instance.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input[int]]:
        """
        Status code of the Instance.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Status code of the Instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class InstanceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: Optional[pulumi.Input[str]] = None):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


