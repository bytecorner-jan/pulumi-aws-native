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
    'AddonTagArgs',
    'ClusterEncryptionConfigProviderPropertiesArgs',
    'ClusterEncryptionConfigArgs',
    'ClusterKubernetesNetworkConfigArgs',
    'ClusterLoggingArgs',
    'ClusterResourcesVpcConfigArgs',
    'ClusterTagArgs',
    'FargateProfileLabelArgs',
    'FargateProfileSelectorArgs',
    'FargateProfileTagArgs',
    'NodegroupLaunchTemplateSpecificationArgs',
    'NodegroupRemoteAccessArgs',
    'NodegroupScalingConfigArgs',
    'NodegroupTaintArgs',
    'NodegroupUpdateConfigArgs',
]

@pulumi.input_type
class AddonTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ClusterEncryptionConfigProviderPropertiesArgs:
    def __init__(__self__, *,
                 key_arn: Optional[pulumi.Input[str]] = None):
        """
        The encryption provider for the cluster.
        :param pulumi.Input[str] key_arn: Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.
        """
        if key_arn is not None:
            pulumi.set(__self__, "key_arn", key_arn)

    @property
    @pulumi.getter(name="keyArn")
    def key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.
        """
        return pulumi.get(self, "key_arn")

    @key_arn.setter
    def key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_arn", value)


@pulumi.input_type
class ClusterEncryptionConfigArgs:
    def __init__(__self__, *,
                 provider: Optional[pulumi.Input['ClusterEncryptionConfigProviderPropertiesArgs']] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The encryption configuration for the cluster
        :param pulumi.Input['ClusterEncryptionConfigProviderPropertiesArgs'] provider: The encryption provider for the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resources: Specifies the resources to be encrypted. The only supported value is "secrets".
        """
        if provider is not None:
            pulumi.set(__self__, "provider", provider)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)

    @property
    @pulumi.getter
    def provider(self) -> Optional[pulumi.Input['ClusterEncryptionConfigProviderPropertiesArgs']]:
        """
        The encryption provider for the cluster.
        """
        return pulumi.get(self, "provider")

    @provider.setter
    def provider(self, value: Optional[pulumi.Input['ClusterEncryptionConfigProviderPropertiesArgs']]):
        pulumi.set(self, "provider", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the resources to be encrypted. The only supported value is "secrets".
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resources", value)


@pulumi.input_type
class ClusterKubernetesNetworkConfigArgs:
    def __init__(__self__, *,
                 service_ipv4_cidr: Optional[pulumi.Input[str]] = None):
        """
        The Kubernetes network configuration for the cluster.
        :param pulumi.Input[str] service_ipv4_cidr: The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. 
        """
        if service_ipv4_cidr is not None:
            pulumi.set(__self__, "service_ipv4_cidr", service_ipv4_cidr)

    @property
    @pulumi.getter(name="serviceIpv4Cidr")
    def service_ipv4_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. 
        """
        return pulumi.get(self, "service_ipv4_cidr")

    @service_ipv4_cidr.setter
    def service_ipv4_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_ipv4_cidr", value)


@pulumi.input_type
class ClusterLoggingArgs:
    def __init__(__self__, *,
                 cluster_logging: Optional[pulumi.Input['ClusterLoggingArgs']] = None):
        """
        Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.
        :param pulumi.Input['ClusterLoggingArgs'] cluster_logging: The cluster control plane logging configuration for your cluster. 
        """
        if cluster_logging is not None:
            pulumi.set(__self__, "cluster_logging", cluster_logging)

    @property
    @pulumi.getter(name="clusterLogging")
    def cluster_logging(self) -> Optional[pulumi.Input['ClusterLoggingArgs']]:
        """
        The cluster control plane logging configuration for your cluster. 
        """
        return pulumi.get(self, "cluster_logging")

    @cluster_logging.setter
    def cluster_logging(self, value: Optional[pulumi.Input['ClusterLoggingArgs']]):
        pulumi.set(self, "cluster_logging", value)


@pulumi.input_type
class ClusterResourcesVpcConfigArgs:
    def __init__(__self__, *,
                 subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 endpoint_private_access: Optional[pulumi.Input[bool]] = None,
                 endpoint_public_access: Optional[pulumi.Input[bool]] = None,
                 public_access_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        An object representing the VPC configuration to use for an Amazon EKS cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.
        :param pulumi.Input[bool] endpoint_private_access: Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.
        :param pulumi.Input[bool] endpoint_public_access: Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_access_cidrs: The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.
        """
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if endpoint_private_access is not None:
            pulumi.set(__self__, "endpoint_private_access", endpoint_private_access)
        if endpoint_public_access is not None:
            pulumi.set(__self__, "endpoint_public_access", endpoint_public_access)
        if public_access_cidrs is not None:
            pulumi.set(__self__, "public_access_cidrs", public_access_cidrs)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="endpointPrivateAccess")
    def endpoint_private_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.
        """
        return pulumi.get(self, "endpoint_private_access")

    @endpoint_private_access.setter
    def endpoint_private_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "endpoint_private_access", value)

    @property
    @pulumi.getter(name="endpointPublicAccess")
    def endpoint_public_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.
        """
        return pulumi.get(self, "endpoint_public_access")

    @endpoint_public_access.setter
    def endpoint_public_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "endpoint_public_access", value)

    @property
    @pulumi.getter(name="publicAccessCidrs")
    def public_access_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.
        """
        return pulumi.get(self, "public_access_cidrs")

    @public_access_cidrs.setter
    def public_access_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "public_access_cidrs", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)


@pulumi.input_type
class ClusterTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
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
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class FargateProfileLabelArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a pod.
        :param pulumi.Input[str] key: The key name of the label.
        :param pulumi.Input[str] value: The value for the label. 
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the label.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the label. 
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class FargateProfileSelectorArgs:
    def __init__(__self__, *,
                 namespace: pulumi.Input[str],
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input['FargateProfileLabelArgs']]]] = None):
        pulumi.set(__self__, "namespace", namespace)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FargateProfileLabelArgs']]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FargateProfileLabelArgs']]]]):
        pulumi.set(self, "labels", value)


@pulumi.input_type
class FargateProfileTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class NodegroupLaunchTemplateSpecificationArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class NodegroupRemoteAccessArgs:
    def __init__(__self__, *,
                 ec2_ssh_key: pulumi.Input[str],
                 source_security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "ec2_ssh_key", ec2_ssh_key)
        if source_security_groups is not None:
            pulumi.set(__self__, "source_security_groups", source_security_groups)

    @property
    @pulumi.getter(name="ec2SshKey")
    def ec2_ssh_key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ec2_ssh_key")

    @ec2_ssh_key.setter
    def ec2_ssh_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "ec2_ssh_key", value)

    @property
    @pulumi.getter(name="sourceSecurityGroups")
    def source_security_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "source_security_groups")

    @source_security_groups.setter
    def source_security_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "source_security_groups", value)


@pulumi.input_type
class NodegroupScalingConfigArgs:
    def __init__(__self__, *,
                 desired_size: Optional[pulumi.Input[float]] = None,
                 max_size: Optional[pulumi.Input[float]] = None,
                 min_size: Optional[pulumi.Input[float]] = None):
        if desired_size is not None:
            pulumi.set(__self__, "desired_size", desired_size)
        if max_size is not None:
            pulumi.set(__self__, "max_size", max_size)
        if min_size is not None:
            pulumi.set(__self__, "min_size", min_size)

    @property
    @pulumi.getter(name="desiredSize")
    def desired_size(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "desired_size")

    @desired_size.setter
    def desired_size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "desired_size", value)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "min_size")

    @min_size.setter
    def min_size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "min_size", value)


@pulumi.input_type
class NodegroupTaintArgs:
    def __init__(__self__, *,
                 effect: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if effect is not None:
            pulumi.set(__self__, "effect", effect)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def effect(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "effect")

    @effect.setter
    def effect(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "effect", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class NodegroupUpdateConfigArgs:
    def __init__(__self__, *,
                 max_unavailable: Optional[pulumi.Input[float]] = None,
                 max_unavailable_percentage: Optional[pulumi.Input[float]] = None):
        if max_unavailable is not None:
            pulumi.set(__self__, "max_unavailable", max_unavailable)
        if max_unavailable_percentage is not None:
            pulumi.set(__self__, "max_unavailable_percentage", max_unavailable_percentage)

    @property
    @pulumi.getter(name="maxUnavailable")
    def max_unavailable(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "max_unavailable")

    @max_unavailable.setter
    def max_unavailable(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_unavailable", value)

    @property
    @pulumi.getter(name="maxUnavailablePercentage")
    def max_unavailable_percentage(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "max_unavailable_percentage")

    @max_unavailable_percentage.setter
    def max_unavailable_percentage(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_unavailable_percentage", value)


