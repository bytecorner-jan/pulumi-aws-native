# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ServiceCodeConfigurationConfigurationSource',
    'ServiceCodeConfigurationValuesRuntime',
    'ServiceHealthCheckConfigurationProtocol',
    'ServiceImageRepositoryImageRepositoryType',
    'ServiceSourceCodeVersionType',
]


class ServiceCodeConfigurationConfigurationSource(str, Enum):
    """
    Configuration Source
    """
    REPOSITORY = "REPOSITORY"
    API = "API"


class ServiceCodeConfigurationValuesRuntime(str, Enum):
    """
    Runtime
    """
    PYTHON3 = "PYTHON_3"
    NODEJS12 = "NODEJS_12"


class ServiceHealthCheckConfigurationProtocol(str, Enum):
    """
    Health Check Protocol
    """
    TCP = "TCP"
    HTTP = "HTTP"


class ServiceImageRepositoryImageRepositoryType(str, Enum):
    """
    Image Repository Type
    """
    ECR = "ECR"
    ECR_PUBLIC = "ECR_PUBLIC"


class ServiceSourceCodeVersionType(str, Enum):
    """
    Source Code Version Type
    """
    BRANCH = "BRANCH"
