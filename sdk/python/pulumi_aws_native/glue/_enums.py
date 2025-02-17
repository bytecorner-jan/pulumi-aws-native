# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'SchemaCompatibility',
    'SchemaDataFormat',
]


class SchemaCompatibility(str, Enum):
    """
    Compatibility setting for the schema.
    """
    NONE = "NONE"
    DISABLED = "DISABLED"
    BACKWARD = "BACKWARD"
    BACKWARD_ALL = "BACKWARD_ALL"
    FORWARD = "FORWARD"
    FORWARD_ALL = "FORWARD_ALL"
    FULL = "FULL"
    FULL_ALL = "FULL_ALL"


class SchemaDataFormat(str, Enum):
    """
    Data format name to use for the schema. Accepted values: 'AVRO', 'JSON'
    """
    AVRO = "AVRO"
    JSON = "JSON"
