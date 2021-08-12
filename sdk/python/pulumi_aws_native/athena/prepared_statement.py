# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PreparedStatementArgs', 'PreparedStatement']

@pulumi.input_type
class PreparedStatementArgs:
    def __init__(__self__, *,
                 query_statement: pulumi.Input[str],
                 statement_name: pulumi.Input[str],
                 work_group: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PreparedStatement resource.
        :param pulumi.Input[str] query_statement: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-querystatement
        :param pulumi.Input[str] statement_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-statementname
        :param pulumi.Input[str] work_group: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-workgroup
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-description
        """
        pulumi.set(__self__, "query_statement", query_statement)
        pulumi.set(__self__, "statement_name", statement_name)
        pulumi.set(__self__, "work_group", work_group)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="queryStatement")
    def query_statement(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-querystatement
        """
        return pulumi.get(self, "query_statement")

    @query_statement.setter
    def query_statement(self, value: pulumi.Input[str]):
        pulumi.set(self, "query_statement", value)

    @property
    @pulumi.getter(name="statementName")
    def statement_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-statementname
        """
        return pulumi.get(self, "statement_name")

    @statement_name.setter
    def statement_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "statement_name", value)

    @property
    @pulumi.getter(name="workGroup")
    def work_group(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-workgroup
        """
        return pulumi.get(self, "work_group")

    @work_group.setter
    def work_group(self, value: pulumi.Input[str]):
        pulumi.set(self, "work_group", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


class PreparedStatement(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 query_statement: Optional[pulumi.Input[str]] = None,
                 statement_name: Optional[pulumi.Input[str]] = None,
                 work_group: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-description
        :param pulumi.Input[str] query_statement: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-querystatement
        :param pulumi.Input[str] statement_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-statementname
        :param pulumi.Input[str] work_group: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-workgroup
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PreparedStatementArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html

        :param str resource_name: The name of the resource.
        :param PreparedStatementArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PreparedStatementArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 query_statement: Optional[pulumi.Input[str]] = None,
                 statement_name: Optional[pulumi.Input[str]] = None,
                 work_group: Optional[pulumi.Input[str]] = None,
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
            __props__ = PreparedStatementArgs.__new__(PreparedStatementArgs)

            __props__.__dict__["description"] = description
            if query_statement is None and not opts.urn:
                raise TypeError("Missing required property 'query_statement'")
            __props__.__dict__["query_statement"] = query_statement
            if statement_name is None and not opts.urn:
                raise TypeError("Missing required property 'statement_name'")
            __props__.__dict__["statement_name"] = statement_name
            if work_group is None and not opts.urn:
                raise TypeError("Missing required property 'work_group'")
            __props__.__dict__["work_group"] = work_group
        super(PreparedStatement, __self__).__init__(
            'aws-native:Athena:PreparedStatement',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PreparedStatement':
        """
        Get an existing PreparedStatement resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PreparedStatementArgs.__new__(PreparedStatementArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["query_statement"] = None
        __props__.__dict__["statement_name"] = None
        __props__.__dict__["work_group"] = None
        return PreparedStatement(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="queryStatement")
    def query_statement(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-querystatement
        """
        return pulumi.get(self, "query_statement")

    @property
    @pulumi.getter(name="statementName")
    def statement_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-statementname
        """
        return pulumi.get(self, "statement_name")

    @property
    @pulumi.getter(name="workGroup")
    def work_group(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-workgroup
        """
        return pulumi.get(self, "work_group")

