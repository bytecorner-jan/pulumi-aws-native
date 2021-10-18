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
    'AliasProvisionedConcurrencyConfigurationArgs',
    'AliasRoutingConfigurationArgs',
    'AliasVersionWeightArgs',
    'CodeSigningConfigAllowedPublishersArgs',
    'CodeSigningConfigCodeSigningPoliciesArgs',
    'EventInvokeConfigDestinationConfigArgs',
    'EventInvokeConfigOnFailureArgs',
    'EventInvokeConfigOnSuccessArgs',
    'EventSourceMappingDestinationConfigArgs',
    'EventSourceMappingEndpointsArgs',
    'EventSourceMappingOnFailureArgs',
    'EventSourceMappingSelfManagedEventSourceArgs',
    'EventSourceMappingSourceAccessConfigurationArgs',
    'FunctionCodeArgs',
    'FunctionDeadLetterConfigArgs',
    'FunctionEnvironmentArgs',
    'FunctionFileSystemConfigArgs',
    'FunctionImageConfigArgs',
    'FunctionTagArgs',
    'FunctionTracingConfigArgs',
    'FunctionVpcConfigArgs',
    'LayerVersionContentArgs',
    'VersionProvisionedConcurrencyConfigurationArgs',
]

@pulumi.input_type
class AliasProvisionedConcurrencyConfigurationArgs:
    def __init__(__self__, *,
                 provisioned_concurrent_executions: pulumi.Input[int]):
        pulumi.set(__self__, "provisioned_concurrent_executions", provisioned_concurrent_executions)

    @property
    @pulumi.getter(name="provisionedConcurrentExecutions")
    def provisioned_concurrent_executions(self) -> pulumi.Input[int]:
        return pulumi.get(self, "provisioned_concurrent_executions")

    @provisioned_concurrent_executions.setter
    def provisioned_concurrent_executions(self, value: pulumi.Input[int]):
        pulumi.set(self, "provisioned_concurrent_executions", value)


@pulumi.input_type
class AliasRoutingConfigurationArgs:
    def __init__(__self__, *,
                 additional_version_weights: pulumi.Input[Sequence[pulumi.Input['AliasVersionWeightArgs']]]):
        pulumi.set(__self__, "additional_version_weights", additional_version_weights)

    @property
    @pulumi.getter(name="additionalVersionWeights")
    def additional_version_weights(self) -> pulumi.Input[Sequence[pulumi.Input['AliasVersionWeightArgs']]]:
        return pulumi.get(self, "additional_version_weights")

    @additional_version_weights.setter
    def additional_version_weights(self, value: pulumi.Input[Sequence[pulumi.Input['AliasVersionWeightArgs']]]):
        pulumi.set(self, "additional_version_weights", value)


@pulumi.input_type
class AliasVersionWeightArgs:
    def __init__(__self__, *,
                 function_version: pulumi.Input[str],
                 function_weight: pulumi.Input[float]):
        pulumi.set(__self__, "function_version", function_version)
        pulumi.set(__self__, "function_weight", function_weight)

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> pulumi.Input[str]:
        return pulumi.get(self, "function_version")

    @function_version.setter
    def function_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_version", value)

    @property
    @pulumi.getter(name="functionWeight")
    def function_weight(self) -> pulumi.Input[float]:
        return pulumi.get(self, "function_weight")

    @function_weight.setter
    def function_weight(self, value: pulumi.Input[float]):
        pulumi.set(self, "function_weight", value)


@pulumi.input_type
class CodeSigningConfigAllowedPublishersArgs:
    def __init__(__self__, *,
                 signing_profile_version_arns: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list
        :param pulumi.Input[Sequence[pulumi.Input[str]]] signing_profile_version_arns: List of Signing profile version Arns
        """
        pulumi.set(__self__, "signing_profile_version_arns", signing_profile_version_arns)

    @property
    @pulumi.getter(name="signingProfileVersionArns")
    def signing_profile_version_arns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of Signing profile version Arns
        """
        return pulumi.get(self, "signing_profile_version_arns")

    @signing_profile_version_arns.setter
    def signing_profile_version_arns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "signing_profile_version_arns", value)


@pulumi.input_type
class CodeSigningConfigCodeSigningPoliciesArgs:
    def __init__(__self__, *,
                 untrusted_artifact_on_deployment: pulumi.Input['CodeSigningConfigCodeSigningPoliciesUntrustedArtifactOnDeployment']):
        """
        Policies to control how to act if a signature is invalid
        :param pulumi.Input['CodeSigningConfigCodeSigningPoliciesUntrustedArtifactOnDeployment'] untrusted_artifact_on_deployment: Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided
        """
        pulumi.set(__self__, "untrusted_artifact_on_deployment", untrusted_artifact_on_deployment)

    @property
    @pulumi.getter(name="untrustedArtifactOnDeployment")
    def untrusted_artifact_on_deployment(self) -> pulumi.Input['CodeSigningConfigCodeSigningPoliciesUntrustedArtifactOnDeployment']:
        """
        Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided
        """
        return pulumi.get(self, "untrusted_artifact_on_deployment")

    @untrusted_artifact_on_deployment.setter
    def untrusted_artifact_on_deployment(self, value: pulumi.Input['CodeSigningConfigCodeSigningPoliciesUntrustedArtifactOnDeployment']):
        pulumi.set(self, "untrusted_artifact_on_deployment", value)


@pulumi.input_type
class EventInvokeConfigDestinationConfigArgs:
    def __init__(__self__, *,
                 on_failure: Optional[pulumi.Input['EventInvokeConfigOnFailureArgs']] = None,
                 on_success: Optional[pulumi.Input['EventInvokeConfigOnSuccessArgs']] = None):
        if on_failure is not None:
            pulumi.set(__self__, "on_failure", on_failure)
        if on_success is not None:
            pulumi.set(__self__, "on_success", on_success)

    @property
    @pulumi.getter(name="onFailure")
    def on_failure(self) -> Optional[pulumi.Input['EventInvokeConfigOnFailureArgs']]:
        return pulumi.get(self, "on_failure")

    @on_failure.setter
    def on_failure(self, value: Optional[pulumi.Input['EventInvokeConfigOnFailureArgs']]):
        pulumi.set(self, "on_failure", value)

    @property
    @pulumi.getter(name="onSuccess")
    def on_success(self) -> Optional[pulumi.Input['EventInvokeConfigOnSuccessArgs']]:
        return pulumi.get(self, "on_success")

    @on_success.setter
    def on_success(self, value: Optional[pulumi.Input['EventInvokeConfigOnSuccessArgs']]):
        pulumi.set(self, "on_success", value)


@pulumi.input_type
class EventInvokeConfigOnFailureArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str]):
        pulumi.set(__self__, "destination", destination)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination", value)


@pulumi.input_type
class EventInvokeConfigOnSuccessArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str]):
        pulumi.set(__self__, "destination", destination)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination", value)


@pulumi.input_type
class EventSourceMappingDestinationConfigArgs:
    def __init__(__self__, *,
                 on_failure: Optional[pulumi.Input['EventSourceMappingOnFailureArgs']] = None):
        """
        (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
        :param pulumi.Input['EventSourceMappingOnFailureArgs'] on_failure: The destination configuration for failed invocations.
        """
        if on_failure is not None:
            pulumi.set(__self__, "on_failure", on_failure)

    @property
    @pulumi.getter(name="onFailure")
    def on_failure(self) -> Optional[pulumi.Input['EventSourceMappingOnFailureArgs']]:
        """
        The destination configuration for failed invocations.
        """
        return pulumi.get(self, "on_failure")

    @on_failure.setter
    def on_failure(self, value: Optional[pulumi.Input['EventSourceMappingOnFailureArgs']]):
        pulumi.set(self, "on_failure", value)


@pulumi.input_type
class EventSourceMappingEndpointsArgs:
    def __init__(__self__, *,
                 kafka_bootstrap_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The endpoints used by AWS Lambda to access a self-managed event source.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] kafka_bootstrap_servers: A list of Kafka server endpoints.
        """
        if kafka_bootstrap_servers is not None:
            pulumi.set(__self__, "kafka_bootstrap_servers", kafka_bootstrap_servers)

    @property
    @pulumi.getter(name="kafkaBootstrapServers")
    def kafka_bootstrap_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of Kafka server endpoints.
        """
        return pulumi.get(self, "kafka_bootstrap_servers")

    @kafka_bootstrap_servers.setter
    def kafka_bootstrap_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "kafka_bootstrap_servers", value)


@pulumi.input_type
class EventSourceMappingOnFailureArgs:
    def __init__(__self__, *,
                 destination: Optional[pulumi.Input[str]] = None):
        """
        A destination for events that failed processing.
        :param pulumi.Input[str] destination: The Amazon Resource Name (ARN) of the destination resource.
        """
        if destination is not None:
            pulumi.set(__self__, "destination", destination)

    @property
    @pulumi.getter
    def destination(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the destination resource.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination", value)


@pulumi.input_type
class EventSourceMappingSelfManagedEventSourceArgs:
    def __init__(__self__, *,
                 endpoints: Optional[pulumi.Input['EventSourceMappingEndpointsArgs']] = None):
        """
        The configuration used by AWS Lambda to access a self-managed event source.
        :param pulumi.Input['EventSourceMappingEndpointsArgs'] endpoints: The endpoints for a self-managed event source.
        """
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input['EventSourceMappingEndpointsArgs']]:
        """
        The endpoints for a self-managed event source.
        """
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input['EventSourceMappingEndpointsArgs']]):
        pulumi.set(self, "endpoints", value)


@pulumi.input_type
class EventSourceMappingSourceAccessConfigurationArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input['EventSourceMappingSourceAccessConfigurationType']] = None,
                 u_ri: Optional[pulumi.Input[str]] = None):
        """
        The configuration used by AWS Lambda to access event source
        :param pulumi.Input['EventSourceMappingSourceAccessConfigurationType'] type: The type of source access configuration.
        :param pulumi.Input[str] u_ri: The URI for the source access configuration resource.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)
        if u_ri is not None:
            pulumi.set(__self__, "u_ri", u_ri)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['EventSourceMappingSourceAccessConfigurationType']]:
        """
        The type of source access configuration.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['EventSourceMappingSourceAccessConfigurationType']]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="uRI")
    def u_ri(self) -> Optional[pulumi.Input[str]]:
        """
        The URI for the source access configuration resource.
        """
        return pulumi.get(self, "u_ri")

    @u_ri.setter
    def u_ri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "u_ri", value)


@pulumi.input_type
class FunctionCodeArgs:
    def __init__(__self__, *,
                 image_uri: Optional[pulumi.Input[str]] = None,
                 s3_bucket: Optional[pulumi.Input[str]] = None,
                 s3_key: Optional[pulumi.Input[str]] = None,
                 s3_object_version: Optional[pulumi.Input[str]] = None,
                 zip_file: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]] = None):
        """
        :param pulumi.Input[str] image_uri: ImageUri.
        :param pulumi.Input[str] s3_bucket: An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.
        :param pulumi.Input[str] s3_key: The Amazon S3 key of the deployment package.
        :param pulumi.Input[str] s3_object_version: For versioned objects, the version of the deployment package object to use.
        :param pulumi.Input[Union[pulumi.Asset, pulumi.Archive]] zip_file: The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..
        """
        if image_uri is not None:
            pulumi.set(__self__, "image_uri", image_uri)
        if s3_bucket is not None:
            pulumi.set(__self__, "s3_bucket", s3_bucket)
        if s3_key is not None:
            pulumi.set(__self__, "s3_key", s3_key)
        if s3_object_version is not None:
            pulumi.set(__self__, "s3_object_version", s3_object_version)
        if zip_file is not None:
            pulumi.set(__self__, "zip_file", zip_file)

    @property
    @pulumi.getter(name="imageUri")
    def image_uri(self) -> Optional[pulumi.Input[str]]:
        """
        ImageUri.
        """
        return pulumi.get(self, "image_uri")

    @image_uri.setter
    def image_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_uri", value)

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.
        """
        return pulumi.get(self, "s3_bucket")

    @s3_bucket.setter
    def s3_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_bucket", value)

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon S3 key of the deployment package.
        """
        return pulumi.get(self, "s3_key")

    @s3_key.setter
    def s3_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_key", value)

    @property
    @pulumi.getter(name="s3ObjectVersion")
    def s3_object_version(self) -> Optional[pulumi.Input[str]]:
        """
        For versioned objects, the version of the deployment package object to use.
        """
        return pulumi.get(self, "s3_object_version")

    @s3_object_version.setter
    def s3_object_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_object_version", value)

    @property
    @pulumi.getter(name="zipFile")
    def zip_file(self) -> Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]:
        """
        The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..
        """
        return pulumi.get(self, "zip_file")

    @zip_file.setter
    def zip_file(self, value: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]):
        pulumi.set(self, "zip_file", value)


@pulumi.input_type
class FunctionDeadLetterConfigArgs:
    def __init__(__self__, *,
                 target_arn: Optional[pulumi.Input[str]] = None):
        """
        The dead-letter queue for failed asynchronous invocations.
        :param pulumi.Input[str] target_arn: The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.
        """
        if target_arn is not None:
            pulumi.set(__self__, "target_arn", target_arn)

    @property
    @pulumi.getter(name="targetArn")
    def target_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.
        """
        return pulumi.get(self, "target_arn")

    @target_arn.setter
    def target_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_arn", value)


@pulumi.input_type
class FunctionEnvironmentArgs:
    def __init__(__self__, *,
                 variables: Optional[Any] = None):
        """
        A function's environment variable settings.
        :param Any variables: Environment variable key-value pairs.
        """
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter
    def variables(self) -> Optional[Any]:
        """
        Environment variable key-value pairs.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[Any]):
        pulumi.set(self, "variables", value)


@pulumi.input_type
class FunctionFileSystemConfigArgs:
    def __init__(__self__, *,
                 arn: pulumi.Input[str],
                 local_mount_path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.
        :param pulumi.Input[str] local_mount_path: The path where the function can access the file system, starting with /mnt/.
        """
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "local_mount_path", local_mount_path)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="localMountPath")
    def local_mount_path(self) -> pulumi.Input[str]:
        """
        The path where the function can access the file system, starting with /mnt/.
        """
        return pulumi.get(self, "local_mount_path")

    @local_mount_path.setter
    def local_mount_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_mount_path", value)


@pulumi.input_type
class FunctionImageConfigArgs:
    def __init__(__self__, *,
                 command: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 entry_point: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 working_directory: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] command: Command.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] entry_point: EntryPoint.
        :param pulumi.Input[str] working_directory: WorkingDirectory.
        """
        if command is not None:
            pulumi.set(__self__, "command", command)
        if entry_point is not None:
            pulumi.set(__self__, "entry_point", entry_point)
        if working_directory is not None:
            pulumi.set(__self__, "working_directory", working_directory)

    @property
    @pulumi.getter
    def command(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Command.
        """
        return pulumi.get(self, "command")

    @command.setter
    def command(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "command", value)

    @property
    @pulumi.getter(name="entryPoint")
    def entry_point(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        EntryPoint.
        """
        return pulumi.get(self, "entry_point")

    @entry_point.setter
    def entry_point(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "entry_point", value)

    @property
    @pulumi.getter(name="workingDirectory")
    def working_directory(self) -> Optional[pulumi.Input[str]]:
        """
        WorkingDirectory.
        """
        return pulumi.get(self, "working_directory")

    @working_directory.setter
    def working_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "working_directory", value)


@pulumi.input_type
class FunctionTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: Optional[pulumi.Input[str]] = None):
        """
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
class FunctionTracingConfigArgs:
    def __init__(__self__, *,
                 mode: Optional[pulumi.Input['FunctionTracingConfigMode']] = None):
        """
        The function's AWS X-Ray tracing configuration. To sample and record incoming requests, set Mode to Active.
        :param pulumi.Input['FunctionTracingConfigMode'] mode: The tracing mode.
        """
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input['FunctionTracingConfigMode']]:
        """
        The tracing mode.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input['FunctionTracingConfigMode']]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class FunctionVpcConfigArgs:
    def __init__(__self__, *,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The VPC security groups and subnets that are attached to a Lambda function. When you connect a function to a VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the function's VPC configuration. The function can only access resources and the internet through that VPC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: A list of VPC security groups IDs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: A list of VPC subnet IDs.
        """
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of VPC security groups IDs.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of VPC subnet IDs.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)


@pulumi.input_type
class LayerVersionContentArgs:
    def __init__(__self__, *,
                 s3_bucket: pulumi.Input[str],
                 s3_key: pulumi.Input[str],
                 s3_object_version: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        pulumi.set(__self__, "s3_key", s3_key)
        if s3_object_version is not None:
            pulumi.set(__self__, "s3_object_version", s3_object_version)

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> pulumi.Input[str]:
        return pulumi.get(self, "s3_bucket")

    @s3_bucket.setter
    def s3_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_bucket", value)

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "s3_key")

    @s3_key.setter
    def s3_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_key", value)

    @property
    @pulumi.getter(name="s3ObjectVersion")
    def s3_object_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "s3_object_version")

    @s3_object_version.setter
    def s3_object_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_object_version", value)


@pulumi.input_type
class VersionProvisionedConcurrencyConfigurationArgs:
    def __init__(__self__, *,
                 provisioned_concurrent_executions: pulumi.Input[int]):
        pulumi.set(__self__, "provisioned_concurrent_executions", provisioned_concurrent_executions)

    @property
    @pulumi.getter(name="provisionedConcurrentExecutions")
    def provisioned_concurrent_executions(self) -> pulumi.Input[int]:
        return pulumi.get(self, "provisioned_concurrent_executions")

    @provisioned_concurrent_executions.setter
    def provisioned_concurrent_executions(self, value: pulumi.Input[int]):
        pulumi.set(self, "provisioned_concurrent_executions", value)


