# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BackupSelectionArgs', 'BackupSelection']

@pulumi.input_type
class BackupSelectionArgs:
    def __init__(__self__, *,
                 backup_plan_id: pulumi.Input[str],
                 backup_selection: pulumi.Input['BackupSelectionResourceTypeArgs']):
        """
        The set of arguments for constructing a BackupSelection resource.
        """
        pulumi.set(__self__, "backup_plan_id", backup_plan_id)
        pulumi.set(__self__, "backup_selection", backup_selection)

    @property
    @pulumi.getter(name="backupPlanId")
    def backup_plan_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "backup_plan_id")

    @backup_plan_id.setter
    def backup_plan_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_plan_id", value)

    @property
    @pulumi.getter(name="backupSelection")
    def backup_selection(self) -> pulumi.Input['BackupSelectionResourceTypeArgs']:
        return pulumi.get(self, "backup_selection")

    @backup_selection.setter
    def backup_selection(self, value: pulumi.Input['BackupSelectionResourceTypeArgs']):
        pulumi.set(self, "backup_selection", value)


class BackupSelection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_plan_id: Optional[pulumi.Input[str]] = None,
                 backup_selection: Optional[pulumi.Input[pulumi.InputType['BackupSelectionResourceTypeArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Backup::BackupSelection

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupSelectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Backup::BackupSelection

        :param str resource_name: The name of the resource.
        :param BackupSelectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupSelectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_plan_id: Optional[pulumi.Input[str]] = None,
                 backup_selection: Optional[pulumi.Input[pulumi.InputType['BackupSelectionResourceTypeArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupSelectionArgs.__new__(BackupSelectionArgs)

            if backup_plan_id is None and not opts.urn:
                raise TypeError("Missing required property 'backup_plan_id'")
            __props__.__dict__["backup_plan_id"] = backup_plan_id
            if backup_selection is None and not opts.urn:
                raise TypeError("Missing required property 'backup_selection'")
            __props__.__dict__["backup_selection"] = backup_selection
            __props__.__dict__["selection_id"] = None
        super(BackupSelection, __self__).__init__(
            'aws-native:backup:BackupSelection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BackupSelection':
        """
        Get an existing BackupSelection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BackupSelectionArgs.__new__(BackupSelectionArgs)

        __props__.__dict__["backup_plan_id"] = None
        __props__.__dict__["backup_selection"] = None
        __props__.__dict__["selection_id"] = None
        return BackupSelection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupPlanId")
    def backup_plan_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "backup_plan_id")

    @property
    @pulumi.getter(name="backupSelection")
    def backup_selection(self) -> pulumi.Output['outputs.BackupSelectionResourceType']:
        return pulumi.get(self, "backup_selection")

    @property
    @pulumi.getter(name="selectionId")
    def selection_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "selection_id")

