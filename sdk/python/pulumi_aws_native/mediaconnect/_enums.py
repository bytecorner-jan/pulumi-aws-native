# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'FlowEncryptionAlgorithm',
    'FlowEncryptionKeyType',
    'FlowEntitlementEncryptionAlgorithm',
    'FlowEntitlementEncryptionKeyType',
    'FlowEntitlementEntitlementStatus',
    'FlowFailoverConfigState',
    'FlowOutputEncryptionAlgorithm',
    'FlowOutputEncryptionKeyType',
    'FlowOutputProtocol',
    'FlowSourceEncryptionAlgorithm',
    'FlowSourceEncryptionKeyType',
    'FlowSourceProtocol',
]


class FlowEncryptionAlgorithm(str, Enum):
    """
    The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
    """
    AES128 = "aes128"
    AES192 = "aes192"
    AES256 = "aes256"


class FlowEncryptionKeyType(str, Enum):
    """
    The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
    """
    SPEKE = "speke"
    STATIC_KEY = "static-key"
    SRT_PASSWORD = "srt-password"


class FlowEntitlementEncryptionAlgorithm(str, Enum):
    """
    The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
    """
    AES128 = "aes128"
    AES192 = "aes192"
    AES256 = "aes256"


class FlowEntitlementEncryptionKeyType(str, Enum):
    """
    The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
    """
    SPEKE = "speke"
    STATIC_KEY = "static-key"


class FlowEntitlementEntitlementStatus(str, Enum):
    """
     An indication of whether the entitlement is enabled.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class FlowFailoverConfigState(str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class FlowOutputEncryptionAlgorithm(str, Enum):
    """
    The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
    """
    AES128 = "aes128"
    AES192 = "aes192"
    AES256 = "aes256"


class FlowOutputEncryptionKeyType(str, Enum):
    """
    The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
    """
    STATIC_KEY = "static-key"
    SRT_PASSWORD = "srt-password"


class FlowOutputProtocol(str, Enum):
    """
    The protocol that is used by the source or output.
    """
    ZIXI_PUSH = "zixi-push"
    RTP_FEC = "rtp-fec"
    RTP = "rtp"
    ZIXI_PULL = "zixi-pull"
    RIST = "rist"
    SRT_LISTENER = "srt-listener"


class FlowSourceEncryptionAlgorithm(str, Enum):
    """
    The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
    """
    AES128 = "aes128"
    AES192 = "aes192"
    AES256 = "aes256"


class FlowSourceEncryptionKeyType(str, Enum):
    """
    The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
    """
    SPEKE = "speke"
    STATIC_KEY = "static-key"


class FlowSourceProtocol(str, Enum):
    """
    The protocol that is used by the source or output.
    """
    ZIXI_PUSH = "zixi-push"
    RTP_FEC = "rtp-fec"
    RTP = "rtp"
    RIST = "rist"
    SRT_LISTENER = "srt-listener"
