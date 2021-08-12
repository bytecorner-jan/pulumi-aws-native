# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['DataSourceArgs', 'DataSource']

@pulumi.input_type
class DataSourceArgs:
    def __init__(__self__, *,
                 alternate_data_source_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceDataSourceParametersArgs']]]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input['DataSourceDataSourceCredentialsArgs']] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 data_source_parameters: Optional[pulumi.Input['DataSourceDataSourceParametersArgs']] = None,
                 error_info: Optional[pulumi.Input['DataSourceDataSourceErrorInfoArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceResourcePermissionArgs']]]] = None,
                 ssl_properties: Optional[pulumi.Input['DataSourceSslPropertiesArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_connection_properties: Optional[pulumi.Input['DataSourceVpcConnectionPropertiesArgs']] = None):
        """
        The set of arguments for constructing a DataSource resource.
        :param pulumi.Input[Sequence[pulumi.Input['DataSourceDataSourceParametersArgs']]] alternate_data_source_parameters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-alternatedatasourceparameters
        :param pulumi.Input[str] aws_account_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-awsaccountid
        :param pulumi.Input['DataSourceDataSourceCredentialsArgs'] credentials: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-credentials
        :param pulumi.Input[str] data_source_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceid
        :param pulumi.Input['DataSourceDataSourceParametersArgs'] data_source_parameters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceparameters
        :param pulumi.Input['DataSourceDataSourceErrorInfoArgs'] error_info: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-errorinfo
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-name
        :param pulumi.Input[Sequence[pulumi.Input['DataSourceResourcePermissionArgs']]] permissions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-permissions
        :param pulumi.Input['DataSourceSslPropertiesArgs'] ssl_properties: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-sslproperties
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-tags
        :param pulumi.Input[str] type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-type
        :param pulumi.Input['DataSourceVpcConnectionPropertiesArgs'] vpc_connection_properties: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-vpcconnectionproperties
        """
        if alternate_data_source_parameters is not None:
            pulumi.set(__self__, "alternate_data_source_parameters", alternate_data_source_parameters)
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if data_source_id is not None:
            pulumi.set(__self__, "data_source_id", data_source_id)
        if data_source_parameters is not None:
            pulumi.set(__self__, "data_source_parameters", data_source_parameters)
        if error_info is not None:
            pulumi.set(__self__, "error_info", error_info)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if ssl_properties is not None:
            pulumi.set(__self__, "ssl_properties", ssl_properties)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vpc_connection_properties is not None:
            pulumi.set(__self__, "vpc_connection_properties", vpc_connection_properties)

    @property
    @pulumi.getter(name="alternateDataSourceParameters")
    def alternate_data_source_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceDataSourceParametersArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-alternatedatasourceparameters
        """
        return pulumi.get(self, "alternate_data_source_parameters")

    @alternate_data_source_parameters.setter
    def alternate_data_source_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceDataSourceParametersArgs']]]]):
        pulumi.set(self, "alternate_data_source_parameters", value)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-awsaccountid
        """
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input['DataSourceDataSourceCredentialsArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-credentials
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input['DataSourceDataSourceCredentialsArgs']]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceid
        """
        return pulumi.get(self, "data_source_id")

    @data_source_id.setter
    def data_source_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_source_id", value)

    @property
    @pulumi.getter(name="dataSourceParameters")
    def data_source_parameters(self) -> Optional[pulumi.Input['DataSourceDataSourceParametersArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceparameters
        """
        return pulumi.get(self, "data_source_parameters")

    @data_source_parameters.setter
    def data_source_parameters(self, value: Optional[pulumi.Input['DataSourceDataSourceParametersArgs']]):
        pulumi.set(self, "data_source_parameters", value)

    @property
    @pulumi.getter(name="errorInfo")
    def error_info(self) -> Optional[pulumi.Input['DataSourceDataSourceErrorInfoArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-errorinfo
        """
        return pulumi.get(self, "error_info")

    @error_info.setter
    def error_info(self, value: Optional[pulumi.Input['DataSourceDataSourceErrorInfoArgs']]):
        pulumi.set(self, "error_info", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceResourcePermissionArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-permissions
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceResourcePermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="sslProperties")
    def ssl_properties(self) -> Optional[pulumi.Input['DataSourceSslPropertiesArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-sslproperties
        """
        return pulumi.get(self, "ssl_properties")

    @ssl_properties.setter
    def ssl_properties(self, value: Optional[pulumi.Input['DataSourceSslPropertiesArgs']]):
        pulumi.set(self, "ssl_properties", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcConnectionProperties")
    def vpc_connection_properties(self) -> Optional[pulumi.Input['DataSourceVpcConnectionPropertiesArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-vpcconnectionproperties
        """
        return pulumi.get(self, "vpc_connection_properties")

    @vpc_connection_properties.setter
    def vpc_connection_properties(self, value: Optional[pulumi.Input['DataSourceVpcConnectionPropertiesArgs']]):
        pulumi.set(self, "vpc_connection_properties", value)


class DataSource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alternate_data_source_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceDataSourceParametersArgs']]]]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input[pulumi.InputType['DataSourceDataSourceCredentialsArgs']]] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 data_source_parameters: Optional[pulumi.Input[pulumi.InputType['DataSourceDataSourceParametersArgs']]] = None,
                 error_info: Optional[pulumi.Input[pulumi.InputType['DataSourceDataSourceErrorInfoArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceResourcePermissionArgs']]]]] = None,
                 ssl_properties: Optional[pulumi.Input[pulumi.InputType['DataSourceSslPropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_connection_properties: Optional[pulumi.Input[pulumi.InputType['DataSourceVpcConnectionPropertiesArgs']]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceDataSourceParametersArgs']]]] alternate_data_source_parameters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-alternatedatasourceparameters
        :param pulumi.Input[str] aws_account_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-awsaccountid
        :param pulumi.Input[pulumi.InputType['DataSourceDataSourceCredentialsArgs']] credentials: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-credentials
        :param pulumi.Input[str] data_source_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceid
        :param pulumi.Input[pulumi.InputType['DataSourceDataSourceParametersArgs']] data_source_parameters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceparameters
        :param pulumi.Input[pulumi.InputType['DataSourceDataSourceErrorInfoArgs']] error_info: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-errorinfo
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-name
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceResourcePermissionArgs']]]] permissions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-permissions
        :param pulumi.Input[pulumi.InputType['DataSourceSslPropertiesArgs']] ssl_properties: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-sslproperties
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-tags
        :param pulumi.Input[str] type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-type
        :param pulumi.Input[pulumi.InputType['DataSourceVpcConnectionPropertiesArgs']] vpc_connection_properties: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-vpcconnectionproperties
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DataSourceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html

        :param str resource_name: The name of the resource.
        :param DataSourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alternate_data_source_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceDataSourceParametersArgs']]]]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input[pulumi.InputType['DataSourceDataSourceCredentialsArgs']]] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 data_source_parameters: Optional[pulumi.Input[pulumi.InputType['DataSourceDataSourceParametersArgs']]] = None,
                 error_info: Optional[pulumi.Input[pulumi.InputType['DataSourceDataSourceErrorInfoArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceResourcePermissionArgs']]]]] = None,
                 ssl_properties: Optional[pulumi.Input[pulumi.InputType['DataSourceSslPropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_connection_properties: Optional[pulumi.Input[pulumi.InputType['DataSourceVpcConnectionPropertiesArgs']]] = None,
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
            __props__ = DataSourceArgs.__new__(DataSourceArgs)

            __props__.__dict__["alternate_data_source_parameters"] = alternate_data_source_parameters
            __props__.__dict__["aws_account_id"] = aws_account_id
            __props__.__dict__["credentials"] = credentials
            __props__.__dict__["data_source_id"] = data_source_id
            __props__.__dict__["data_source_parameters"] = data_source_parameters
            __props__.__dict__["error_info"] = error_info
            __props__.__dict__["name"] = name
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["ssl_properties"] = ssl_properties
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["vpc_connection_properties"] = vpc_connection_properties
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["last_updated_time"] = None
            __props__.__dict__["status"] = None
        super(DataSource, __self__).__init__(
            'aws-native:QuickSight:DataSource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataSource':
        """
        Get an existing DataSource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataSourceArgs.__new__(DataSourceArgs)

        __props__.__dict__["alternate_data_source_parameters"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_account_id"] = None
        __props__.__dict__["created_time"] = None
        __props__.__dict__["credentials"] = None
        __props__.__dict__["data_source_id"] = None
        __props__.__dict__["data_source_parameters"] = None
        __props__.__dict__["error_info"] = None
        __props__.__dict__["last_updated_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["permissions"] = None
        __props__.__dict__["ssl_properties"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["vpc_connection_properties"] = None
        return DataSource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alternateDataSourceParameters")
    def alternate_data_source_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.DataSourceDataSourceParameters']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-alternatedatasourceparameters
        """
        return pulumi.get(self, "alternate_data_source_parameters")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-awsaccountid
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional['outputs.DataSourceDataSourceCredentials']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-credentials
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceid
        """
        return pulumi.get(self, "data_source_id")

    @property
    @pulumi.getter(name="dataSourceParameters")
    def data_source_parameters(self) -> pulumi.Output[Optional['outputs.DataSourceDataSourceParameters']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceparameters
        """
        return pulumi.get(self, "data_source_parameters")

    @property
    @pulumi.getter(name="errorInfo")
    def error_info(self) -> pulumi.Output[Optional['outputs.DataSourceDataSourceErrorInfo']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-errorinfo
        """
        return pulumi.get(self, "error_info")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence['outputs.DataSourceResourcePermission']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-permissions
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="sslProperties")
    def ssl_properties(self) -> pulumi.Output[Optional['outputs.DataSourceSslProperties']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-sslproperties
        """
        return pulumi.get(self, "ssl_properties")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcConnectionProperties")
    def vpc_connection_properties(self) -> pulumi.Output[Optional['outputs.DataSourceVpcConnectionProperties']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-vpcconnectionproperties
        """
        return pulumi.get(self, "vpc_connection_properties")

