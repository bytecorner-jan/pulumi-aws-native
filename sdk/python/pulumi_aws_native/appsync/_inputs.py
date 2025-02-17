# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DataSourceAuthorizationConfigArgs',
    'DataSourceAwsIamConfigArgs',
    'DataSourceDeltaSyncConfigArgs',
    'DataSourceDynamoDBConfigArgs',
    'DataSourceElasticsearchConfigArgs',
    'DataSourceHttpConfigArgs',
    'DataSourceLambdaConfigArgs',
    'DataSourceOpenSearchServiceConfigArgs',
    'DataSourceRdsHttpEndpointConfigArgs',
    'DataSourceRelationalDatabaseConfigArgs',
    'FunctionConfigurationLambdaConflictHandlerConfigArgs',
    'FunctionConfigurationSyncConfigArgs',
    'GraphQLApiAdditionalAuthenticationProvidersArgs',
    'GraphQLApiLambdaAuthorizerConfigArgs',
    'GraphQLApiLogConfigArgs',
    'GraphQLApiOpenIDConnectConfigArgs',
    'GraphQLApiTagsArgs',
    'GraphQLApiUserPoolConfigArgs',
    'ResolverCachingConfigArgs',
    'ResolverLambdaConflictHandlerConfigArgs',
    'ResolverPipelineConfigArgs',
    'ResolverSyncConfigArgs',
]

@pulumi.input_type
class DataSourceAuthorizationConfigArgs:
    def __init__(__self__, *,
                 authorization_type: pulumi.Input[str],
                 aws_iam_config: Optional[pulumi.Input['DataSourceAwsIamConfigArgs']] = None):
        pulumi.set(__self__, "authorization_type", authorization_type)
        if aws_iam_config is not None:
            pulumi.set(__self__, "aws_iam_config", aws_iam_config)

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "authorization_type")

    @authorization_type.setter
    def authorization_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorization_type", value)

    @property
    @pulumi.getter(name="awsIamConfig")
    def aws_iam_config(self) -> Optional[pulumi.Input['DataSourceAwsIamConfigArgs']]:
        return pulumi.get(self, "aws_iam_config")

    @aws_iam_config.setter
    def aws_iam_config(self, value: Optional[pulumi.Input['DataSourceAwsIamConfigArgs']]):
        pulumi.set(self, "aws_iam_config", value)


@pulumi.input_type
class DataSourceAwsIamConfigArgs:
    def __init__(__self__, *,
                 signing_region: Optional[pulumi.Input[str]] = None,
                 signing_service_name: Optional[pulumi.Input[str]] = None):
        if signing_region is not None:
            pulumi.set(__self__, "signing_region", signing_region)
        if signing_service_name is not None:
            pulumi.set(__self__, "signing_service_name", signing_service_name)

    @property
    @pulumi.getter(name="signingRegion")
    def signing_region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "signing_region")

    @signing_region.setter
    def signing_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signing_region", value)

    @property
    @pulumi.getter(name="signingServiceName")
    def signing_service_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "signing_service_name")

    @signing_service_name.setter
    def signing_service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signing_service_name", value)


@pulumi.input_type
class DataSourceDeltaSyncConfigArgs:
    def __init__(__self__, *,
                 base_table_ttl: pulumi.Input[str],
                 delta_sync_table_name: pulumi.Input[str],
                 delta_sync_table_ttl: pulumi.Input[str]):
        pulumi.set(__self__, "base_table_ttl", base_table_ttl)
        pulumi.set(__self__, "delta_sync_table_name", delta_sync_table_name)
        pulumi.set(__self__, "delta_sync_table_ttl", delta_sync_table_ttl)

    @property
    @pulumi.getter(name="baseTableTTL")
    def base_table_ttl(self) -> pulumi.Input[str]:
        return pulumi.get(self, "base_table_ttl")

    @base_table_ttl.setter
    def base_table_ttl(self, value: pulumi.Input[str]):
        pulumi.set(self, "base_table_ttl", value)

    @property
    @pulumi.getter(name="deltaSyncTableName")
    def delta_sync_table_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "delta_sync_table_name")

    @delta_sync_table_name.setter
    def delta_sync_table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "delta_sync_table_name", value)

    @property
    @pulumi.getter(name="deltaSyncTableTTL")
    def delta_sync_table_ttl(self) -> pulumi.Input[str]:
        return pulumi.get(self, "delta_sync_table_ttl")

    @delta_sync_table_ttl.setter
    def delta_sync_table_ttl(self, value: pulumi.Input[str]):
        pulumi.set(self, "delta_sync_table_ttl", value)


@pulumi.input_type
class DataSourceDynamoDBConfigArgs:
    def __init__(__self__, *,
                 aws_region: pulumi.Input[str],
                 table_name: pulumi.Input[str],
                 delta_sync_config: Optional[pulumi.Input['DataSourceDeltaSyncConfigArgs']] = None,
                 use_caller_credentials: Optional[pulumi.Input[bool]] = None,
                 versioned: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "aws_region", aws_region)
        pulumi.set(__self__, "table_name", table_name)
        if delta_sync_config is not None:
            pulumi.set(__self__, "delta_sync_config", delta_sync_config)
        if use_caller_credentials is not None:
            pulumi.set(__self__, "use_caller_credentials", use_caller_credentials)
        if versioned is not None:
            pulumi.set(__self__, "versioned", versioned)

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "aws_region")

    @aws_region.setter
    def aws_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "aws_region", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_name", value)

    @property
    @pulumi.getter(name="deltaSyncConfig")
    def delta_sync_config(self) -> Optional[pulumi.Input['DataSourceDeltaSyncConfigArgs']]:
        return pulumi.get(self, "delta_sync_config")

    @delta_sync_config.setter
    def delta_sync_config(self, value: Optional[pulumi.Input['DataSourceDeltaSyncConfigArgs']]):
        pulumi.set(self, "delta_sync_config", value)

    @property
    @pulumi.getter(name="useCallerCredentials")
    def use_caller_credentials(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_caller_credentials")

    @use_caller_credentials.setter
    def use_caller_credentials(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_caller_credentials", value)

    @property
    @pulumi.getter
    def versioned(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "versioned")

    @versioned.setter
    def versioned(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "versioned", value)


@pulumi.input_type
class DataSourceElasticsearchConfigArgs:
    def __init__(__self__, *,
                 aws_region: pulumi.Input[str],
                 endpoint: pulumi.Input[str]):
        pulumi.set(__self__, "aws_region", aws_region)
        pulumi.set(__self__, "endpoint", endpoint)

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "aws_region")

    @aws_region.setter
    def aws_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "aws_region", value)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)


@pulumi.input_type
class DataSourceHttpConfigArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[str],
                 authorization_config: Optional[pulumi.Input['DataSourceAuthorizationConfigArgs']] = None):
        pulumi.set(__self__, "endpoint", endpoint)
        if authorization_config is not None:
            pulumi.set(__self__, "authorization_config", authorization_config)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="authorizationConfig")
    def authorization_config(self) -> Optional[pulumi.Input['DataSourceAuthorizationConfigArgs']]:
        return pulumi.get(self, "authorization_config")

    @authorization_config.setter
    def authorization_config(self, value: Optional[pulumi.Input['DataSourceAuthorizationConfigArgs']]):
        pulumi.set(self, "authorization_config", value)


@pulumi.input_type
class DataSourceLambdaConfigArgs:
    def __init__(__self__, *,
                 lambda_function_arn: pulumi.Input[str]):
        pulumi.set(__self__, "lambda_function_arn", lambda_function_arn)

    @property
    @pulumi.getter(name="lambdaFunctionArn")
    def lambda_function_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "lambda_function_arn")

    @lambda_function_arn.setter
    def lambda_function_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "lambda_function_arn", value)


@pulumi.input_type
class DataSourceOpenSearchServiceConfigArgs:
    def __init__(__self__, *,
                 aws_region: pulumi.Input[str],
                 endpoint: pulumi.Input[str]):
        pulumi.set(__self__, "aws_region", aws_region)
        pulumi.set(__self__, "endpoint", endpoint)

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "aws_region")

    @aws_region.setter
    def aws_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "aws_region", value)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)


@pulumi.input_type
class DataSourceRdsHttpEndpointConfigArgs:
    def __init__(__self__, *,
                 aws_region: pulumi.Input[str],
                 aws_secret_store_arn: pulumi.Input[str],
                 db_cluster_identifier: pulumi.Input[str],
                 database_name: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "aws_region", aws_region)
        pulumi.set(__self__, "aws_secret_store_arn", aws_secret_store_arn)
        pulumi.set(__self__, "db_cluster_identifier", db_cluster_identifier)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "aws_region")

    @aws_region.setter
    def aws_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "aws_region", value)

    @property
    @pulumi.getter(name="awsSecretStoreArn")
    def aws_secret_store_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "aws_secret_store_arn")

    @aws_secret_store_arn.setter
    def aws_secret_store_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "aws_secret_store_arn", value)

    @property
    @pulumi.getter(name="dbClusterIdentifier")
    def db_cluster_identifier(self) -> pulumi.Input[str]:
        return pulumi.get(self, "db_cluster_identifier")

    @db_cluster_identifier.setter
    def db_cluster_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_cluster_identifier", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)


@pulumi.input_type
class DataSourceRelationalDatabaseConfigArgs:
    def __init__(__self__, *,
                 relational_database_source_type: pulumi.Input[str],
                 rds_http_endpoint_config: Optional[pulumi.Input['DataSourceRdsHttpEndpointConfigArgs']] = None):
        pulumi.set(__self__, "relational_database_source_type", relational_database_source_type)
        if rds_http_endpoint_config is not None:
            pulumi.set(__self__, "rds_http_endpoint_config", rds_http_endpoint_config)

    @property
    @pulumi.getter(name="relationalDatabaseSourceType")
    def relational_database_source_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "relational_database_source_type")

    @relational_database_source_type.setter
    def relational_database_source_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "relational_database_source_type", value)

    @property
    @pulumi.getter(name="rdsHttpEndpointConfig")
    def rds_http_endpoint_config(self) -> Optional[pulumi.Input['DataSourceRdsHttpEndpointConfigArgs']]:
        return pulumi.get(self, "rds_http_endpoint_config")

    @rds_http_endpoint_config.setter
    def rds_http_endpoint_config(self, value: Optional[pulumi.Input['DataSourceRdsHttpEndpointConfigArgs']]):
        pulumi.set(self, "rds_http_endpoint_config", value)


@pulumi.input_type
class FunctionConfigurationLambdaConflictHandlerConfigArgs:
    def __init__(__self__, *,
                 lambda_conflict_handler_arn: Optional[pulumi.Input[str]] = None):
        if lambda_conflict_handler_arn is not None:
            pulumi.set(__self__, "lambda_conflict_handler_arn", lambda_conflict_handler_arn)

    @property
    @pulumi.getter(name="lambdaConflictHandlerArn")
    def lambda_conflict_handler_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "lambda_conflict_handler_arn")

    @lambda_conflict_handler_arn.setter
    def lambda_conflict_handler_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lambda_conflict_handler_arn", value)


@pulumi.input_type
class FunctionConfigurationSyncConfigArgs:
    def __init__(__self__, *,
                 conflict_detection: pulumi.Input[str],
                 conflict_handler: Optional[pulumi.Input[str]] = None,
                 lambda_conflict_handler_config: Optional[pulumi.Input['FunctionConfigurationLambdaConflictHandlerConfigArgs']] = None):
        pulumi.set(__self__, "conflict_detection", conflict_detection)
        if conflict_handler is not None:
            pulumi.set(__self__, "conflict_handler", conflict_handler)
        if lambda_conflict_handler_config is not None:
            pulumi.set(__self__, "lambda_conflict_handler_config", lambda_conflict_handler_config)

    @property
    @pulumi.getter(name="conflictDetection")
    def conflict_detection(self) -> pulumi.Input[str]:
        return pulumi.get(self, "conflict_detection")

    @conflict_detection.setter
    def conflict_detection(self, value: pulumi.Input[str]):
        pulumi.set(self, "conflict_detection", value)

    @property
    @pulumi.getter(name="conflictHandler")
    def conflict_handler(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "conflict_handler")

    @conflict_handler.setter
    def conflict_handler(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "conflict_handler", value)

    @property
    @pulumi.getter(name="lambdaConflictHandlerConfig")
    def lambda_conflict_handler_config(self) -> Optional[pulumi.Input['FunctionConfigurationLambdaConflictHandlerConfigArgs']]:
        return pulumi.get(self, "lambda_conflict_handler_config")

    @lambda_conflict_handler_config.setter
    def lambda_conflict_handler_config(self, value: Optional[pulumi.Input['FunctionConfigurationLambdaConflictHandlerConfigArgs']]):
        pulumi.set(self, "lambda_conflict_handler_config", value)


@pulumi.input_type
class GraphQLApiAdditionalAuthenticationProvidersArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class GraphQLApiLambdaAuthorizerConfigArgs:
    def __init__(__self__, *,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[float]] = None,
                 authorizer_uri: Optional[pulumi.Input[str]] = None,
                 identity_validation_expression: Optional[pulumi.Input[str]] = None):
        if authorizer_result_ttl_in_seconds is not None:
            pulumi.set(__self__, "authorizer_result_ttl_in_seconds", authorizer_result_ttl_in_seconds)
        if authorizer_uri is not None:
            pulumi.set(__self__, "authorizer_uri", authorizer_uri)
        if identity_validation_expression is not None:
            pulumi.set(__self__, "identity_validation_expression", identity_validation_expression)

    @property
    @pulumi.getter(name="authorizerResultTtlInSeconds")
    def authorizer_result_ttl_in_seconds(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "authorizer_result_ttl_in_seconds")

    @authorizer_result_ttl_in_seconds.setter
    def authorizer_result_ttl_in_seconds(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "authorizer_result_ttl_in_seconds", value)

    @property
    @pulumi.getter(name="authorizerUri")
    def authorizer_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authorizer_uri")

    @authorizer_uri.setter
    def authorizer_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_uri", value)

    @property
    @pulumi.getter(name="identityValidationExpression")
    def identity_validation_expression(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "identity_validation_expression")

    @identity_validation_expression.setter
    def identity_validation_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_validation_expression", value)


@pulumi.input_type
class GraphQLApiLogConfigArgs:
    def __init__(__self__, *,
                 cloud_watch_logs_role_arn: Optional[pulumi.Input[str]] = None,
                 exclude_verbose_content: Optional[pulumi.Input[bool]] = None,
                 field_log_level: Optional[pulumi.Input[str]] = None):
        if cloud_watch_logs_role_arn is not None:
            pulumi.set(__self__, "cloud_watch_logs_role_arn", cloud_watch_logs_role_arn)
        if exclude_verbose_content is not None:
            pulumi.set(__self__, "exclude_verbose_content", exclude_verbose_content)
        if field_log_level is not None:
            pulumi.set(__self__, "field_log_level", field_log_level)

    @property
    @pulumi.getter(name="cloudWatchLogsRoleArn")
    def cloud_watch_logs_role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cloud_watch_logs_role_arn")

    @cloud_watch_logs_role_arn.setter
    def cloud_watch_logs_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_watch_logs_role_arn", value)

    @property
    @pulumi.getter(name="excludeVerboseContent")
    def exclude_verbose_content(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "exclude_verbose_content")

    @exclude_verbose_content.setter
    def exclude_verbose_content(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_verbose_content", value)

    @property
    @pulumi.getter(name="fieldLogLevel")
    def field_log_level(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "field_log_level")

    @field_log_level.setter
    def field_log_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_log_level", value)


@pulumi.input_type
class GraphQLApiOpenIDConnectConfigArgs:
    def __init__(__self__, *,
                 auth_ttl: Optional[pulumi.Input[float]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 iat_ttl: Optional[pulumi.Input[float]] = None,
                 issuer: Optional[pulumi.Input[str]] = None):
        if auth_ttl is not None:
            pulumi.set(__self__, "auth_ttl", auth_ttl)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if iat_ttl is not None:
            pulumi.set(__self__, "iat_ttl", iat_ttl)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)

    @property
    @pulumi.getter(name="authTTL")
    def auth_ttl(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "auth_ttl")

    @auth_ttl.setter
    def auth_ttl(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "auth_ttl", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="iatTTL")
    def iat_ttl(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "iat_ttl")

    @iat_ttl.setter
    def iat_ttl(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "iat_ttl", value)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)


@pulumi.input_type
class GraphQLApiTagsArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class GraphQLApiUserPoolConfigArgs:
    def __init__(__self__, *,
                 app_id_client_regex: Optional[pulumi.Input[str]] = None,
                 aws_region: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 user_pool_id: Optional[pulumi.Input[str]] = None):
        if app_id_client_regex is not None:
            pulumi.set(__self__, "app_id_client_regex", app_id_client_regex)
        if aws_region is not None:
            pulumi.set(__self__, "aws_region", aws_region)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if user_pool_id is not None:
            pulumi.set(__self__, "user_pool_id", user_pool_id)

    @property
    @pulumi.getter(name="appIdClientRegex")
    def app_id_client_regex(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "app_id_client_regex")

    @app_id_client_regex.setter
    def app_id_client_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id_client_regex", value)

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "aws_region")

    @aws_region.setter
    def aws_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_region", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="userPoolId")
    def user_pool_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_pool_id")

    @user_pool_id.setter
    def user_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_pool_id", value)


@pulumi.input_type
class ResolverCachingConfigArgs:
    def __init__(__self__, *,
                 caching_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[float]] = None):
        if caching_keys is not None:
            pulumi.set(__self__, "caching_keys", caching_keys)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter(name="cachingKeys")
    def caching_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "caching_keys")

    @caching_keys.setter
    def caching_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "caching_keys", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class ResolverLambdaConflictHandlerConfigArgs:
    def __init__(__self__, *,
                 lambda_conflict_handler_arn: Optional[pulumi.Input[str]] = None):
        if lambda_conflict_handler_arn is not None:
            pulumi.set(__self__, "lambda_conflict_handler_arn", lambda_conflict_handler_arn)

    @property
    @pulumi.getter(name="lambdaConflictHandlerArn")
    def lambda_conflict_handler_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "lambda_conflict_handler_arn")

    @lambda_conflict_handler_arn.setter
    def lambda_conflict_handler_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lambda_conflict_handler_arn", value)


@pulumi.input_type
class ResolverPipelineConfigArgs:
    def __init__(__self__, *,
                 functions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if functions is not None:
            pulumi.set(__self__, "functions", functions)

    @property
    @pulumi.getter
    def functions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "functions")

    @functions.setter
    def functions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "functions", value)


@pulumi.input_type
class ResolverSyncConfigArgs:
    def __init__(__self__, *,
                 conflict_detection: pulumi.Input[str],
                 conflict_handler: Optional[pulumi.Input[str]] = None,
                 lambda_conflict_handler_config: Optional[pulumi.Input['ResolverLambdaConflictHandlerConfigArgs']] = None):
        pulumi.set(__self__, "conflict_detection", conflict_detection)
        if conflict_handler is not None:
            pulumi.set(__self__, "conflict_handler", conflict_handler)
        if lambda_conflict_handler_config is not None:
            pulumi.set(__self__, "lambda_conflict_handler_config", lambda_conflict_handler_config)

    @property
    @pulumi.getter(name="conflictDetection")
    def conflict_detection(self) -> pulumi.Input[str]:
        return pulumi.get(self, "conflict_detection")

    @conflict_detection.setter
    def conflict_detection(self, value: pulumi.Input[str]):
        pulumi.set(self, "conflict_detection", value)

    @property
    @pulumi.getter(name="conflictHandler")
    def conflict_handler(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "conflict_handler")

    @conflict_handler.setter
    def conflict_handler(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "conflict_handler", value)

    @property
    @pulumi.getter(name="lambdaConflictHandlerConfig")
    def lambda_conflict_handler_config(self) -> Optional[pulumi.Input['ResolverLambdaConflictHandlerConfigArgs']]:
        return pulumi.get(self, "lambda_conflict_handler_config")

    @lambda_conflict_handler_config.setter
    def lambda_conflict_handler_config(self, value: Optional[pulumi.Input['ResolverLambdaConflictHandlerConfigArgs']]):
        pulumi.set(self, "lambda_conflict_handler_config", value)


