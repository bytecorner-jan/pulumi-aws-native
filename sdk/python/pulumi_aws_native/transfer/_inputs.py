# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ServerEndpointDetailsArgs',
    'ServerIdentityProviderDetailsArgs',
    'ServerProtocolDetailsArgs',
    'ServerProtocolArgs',
    'ServerTagArgs',
    'ServerWorkflowDetailsArgs',
    'ServerWorkflowDetailArgs',
    'UserHomeDirectoryMapEntryArgs',
    'UserPosixProfileArgs',
    'UserSshPublicKeyArgs',
    'UserTagArgs',
]

@pulumi.input_type
class ServerEndpointDetailsArgs:
    def __init__(__self__, *,
                 address_allocation_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        if address_allocation_ids is not None:
            pulumi.set(__self__, "address_allocation_ids", address_allocation_ids)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if vpc_endpoint_id is not None:
            pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="addressAllocationIds")
    def address_allocation_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "address_allocation_ids")

    @address_allocation_ids.setter
    def address_allocation_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "address_allocation_ids", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_endpoint_id")

    @vpc_endpoint_id.setter
    def vpc_endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_endpoint_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class ServerIdentityProviderDetailsArgs:
    def __init__(__self__, *,
                 directory_id: Optional[pulumi.Input[str]] = None,
                 invocation_role: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        if directory_id is not None:
            pulumi.set(__self__, "directory_id", directory_id)
        if invocation_role is not None:
            pulumi.set(__self__, "invocation_role", invocation_role)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "directory_id")

    @directory_id.setter
    def directory_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_id", value)

    @property
    @pulumi.getter(name="invocationRole")
    def invocation_role(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "invocation_role")

    @invocation_role.setter
    def invocation_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invocation_role", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class ServerProtocolDetailsArgs:
    def __init__(__self__, *,
                 passive_ip: Optional[pulumi.Input[str]] = None):
        if passive_ip is not None:
            pulumi.set(__self__, "passive_ip", passive_ip)

    @property
    @pulumi.getter(name="passiveIp")
    def passive_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "passive_ip")

    @passive_ip.setter
    def passive_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passive_ip", value)


@pulumi.input_type
class ServerProtocolArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class ServerTagArgs:
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
class ServerWorkflowDetailsArgs:
    def __init__(__self__, *,
                 on_upload: pulumi.Input[Sequence[pulumi.Input['ServerWorkflowDetailArgs']]]):
        pulumi.set(__self__, "on_upload", on_upload)

    @property
    @pulumi.getter(name="onUpload")
    def on_upload(self) -> pulumi.Input[Sequence[pulumi.Input['ServerWorkflowDetailArgs']]]:
        return pulumi.get(self, "on_upload")

    @on_upload.setter
    def on_upload(self, value: pulumi.Input[Sequence[pulumi.Input['ServerWorkflowDetailArgs']]]):
        pulumi.set(self, "on_upload", value)


@pulumi.input_type
class ServerWorkflowDetailArgs:
    def __init__(__self__, *,
                 execution_role: pulumi.Input[str],
                 workflow_id: pulumi.Input[str]):
        pulumi.set(__self__, "execution_role", execution_role)
        pulumi.set(__self__, "workflow_id", workflow_id)

    @property
    @pulumi.getter(name="executionRole")
    def execution_role(self) -> pulumi.Input[str]:
        return pulumi.get(self, "execution_role")

    @execution_role.setter
    def execution_role(self, value: pulumi.Input[str]):
        pulumi.set(self, "execution_role", value)

    @property
    @pulumi.getter(name="workflowId")
    def workflow_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workflow_id")

    @workflow_id.setter
    def workflow_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workflow_id", value)


@pulumi.input_type
class UserHomeDirectoryMapEntryArgs:
    def __init__(__self__, *,
                 entry: pulumi.Input[str],
                 target: pulumi.Input[str]):
        pulumi.set(__self__, "entry", entry)
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def entry(self) -> pulumi.Input[str]:
        return pulumi.get(self, "entry")

    @entry.setter
    def entry(self, value: pulumi.Input[str]):
        pulumi.set(self, "entry", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input[str]):
        pulumi.set(self, "target", value)


@pulumi.input_type
class UserPosixProfileArgs:
    def __init__(__self__, *,
                 gid: pulumi.Input[float],
                 uid: pulumi.Input[float],
                 secondary_gids: Optional[pulumi.Input[Sequence[pulumi.Input[float]]]] = None):
        pulumi.set(__self__, "gid", gid)
        pulumi.set(__self__, "uid", uid)
        if secondary_gids is not None:
            pulumi.set(__self__, "secondary_gids", secondary_gids)

    @property
    @pulumi.getter
    def gid(self) -> pulumi.Input[float]:
        return pulumi.get(self, "gid")

    @gid.setter
    def gid(self, value: pulumi.Input[float]):
        pulumi.set(self, "gid", value)

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Input[float]:
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: pulumi.Input[float]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter(name="secondaryGids")
    def secondary_gids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[float]]]]:
        return pulumi.get(self, "secondary_gids")

    @secondary_gids.setter
    def secondary_gids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[float]]]]):
        pulumi.set(self, "secondary_gids", value)


@pulumi.input_type
class UserSshPublicKeyArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class UserTagArgs:
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


