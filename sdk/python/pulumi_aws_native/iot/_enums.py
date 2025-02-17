# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuthorizerStatus',
    'CertificateMode',
    'CertificateStatus',
    'CustomMetricMetricType',
    'DimensionType',
    'DomainConfigurationDomainType',
    'DomainConfigurationServerCertificateSummaryServerCertificateStatus',
    'DomainConfigurationServiceType',
    'DomainConfigurationStatus',
    'JobTemplateAction',
    'JobTemplateFailureType',
    'LoggingDefaultLogLevel',
    'MitigationActionEnableIoTLoggingParamsLogLevel',
    'MitigationActionReplaceDefaultPolicyVersionParamsTemplateName',
    'MitigationActionUpdateCACertificateParamsAction',
    'MitigationActionUpdateDeviceCertificateParamsAction',
    'ResourceSpecificLoggingLogLevel',
    'ResourceSpecificLoggingTargetType',
    'ScheduledAuditDayOfWeek',
    'ScheduledAuditFrequency',
    'SecurityProfileBehaviorCriteriaComparisonOperator',
    'SecurityProfileMachineLearningDetectionConfigConfidenceLevel',
    'SecurityProfileMetricDimensionOperator',
    'SecurityProfileStatisticalThresholdStatistic',
    'TopicRuleCannedAccessControlList',
    'TopicRuleDestinationStatus',
]


class AuthorizerStatus(str, Enum):
    ACTIVE = "ACTIVE"
    INACTIVE = "INACTIVE"


class CertificateMode(str, Enum):
    DEFAULT = "DEFAULT"
    SNI_ONLY = "SNI_ONLY"


class CertificateStatus(str, Enum):
    ACTIVE = "ACTIVE"
    INACTIVE = "INACTIVE"
    REVOKED = "REVOKED"
    PENDING_TRANSFER = "PENDING_TRANSFER"
    PENDING_ACTIVATION = "PENDING_ACTIVATION"


class CustomMetricMetricType(str, Enum):
    """
    The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
    """
    STRING_LIST = "string-list"
    IP_ADDRESS_LIST = "ip-address-list"
    NUMBER_LIST = "number-list"
    NUMBER = "number"


class DimensionType(str, Enum):
    """
    Specifies the type of the dimension.
    """
    TOPIC_FILTER = "TOPIC_FILTER"


class DomainConfigurationDomainType(str, Enum):
    ENDPOINT = "ENDPOINT"
    AWS_MANAGED = "AWS_MANAGED"
    CUSTOMER_MANAGED = "CUSTOMER_MANAGED"


class DomainConfigurationServerCertificateSummaryServerCertificateStatus(str, Enum):
    INVALID = "INVALID"
    VALID = "VALID"


class DomainConfigurationServiceType(str, Enum):
    DATA = "DATA"
    CREDENTIAL_PROVIDER = "CREDENTIAL_PROVIDER"
    JOBS = "JOBS"


class DomainConfigurationStatus(str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class JobTemplateAction(str, Enum):
    CANCEL = "CANCEL"


class JobTemplateFailureType(str, Enum):
    FAILED = "FAILED"
    REJECTED = "REJECTED"
    TIMED_OUT = "TIMED_OUT"
    ALL = "ALL"


class LoggingDefaultLogLevel(str, Enum):
    """
    The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
    """
    ERROR = "ERROR"
    WARN = "WARN"
    INFO = "INFO"
    DEBUG = "DEBUG"
    DISABLED = "DISABLED"


class MitigationActionEnableIoTLoggingParamsLogLevel(str, Enum):
    """
     Specifies which types of information are logged.
    """
    DEBUG = "DEBUG"
    INFO = "INFO"
    ERROR = "ERROR"
    WARN = "WARN"


class MitigationActionReplaceDefaultPolicyVersionParamsTemplateName(str, Enum):
    BLANK_POLICY = "BLANK_POLICY"


class MitigationActionUpdateCACertificateParamsAction(str, Enum):
    DEACTIVATE = "DEACTIVATE"


class MitigationActionUpdateDeviceCertificateParamsAction(str, Enum):
    DEACTIVATE = "DEACTIVATE"


class ResourceSpecificLoggingLogLevel(str, Enum):
    """
    The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
    """
    ERROR = "ERROR"
    WARN = "WARN"
    INFO = "INFO"
    DEBUG = "DEBUG"
    DISABLED = "DISABLED"


class ResourceSpecificLoggingTargetType(str, Enum):
    """
    The target type. Value must be THING_GROUP.
    """
    THING_GROUP = "THING_GROUP"


class ScheduledAuditDayOfWeek(str, Enum):
    """
    The day of the week on which the scheduled audit takes place. Can be one of SUN, MON, TUE,WED, THU, FRI, or SAT. This field is required if the frequency parameter is set to WEEKLY or BIWEEKLY.
    """
    SUN = "SUN"
    MON = "MON"
    TUE = "TUE"
    WED = "WED"
    THU = "THU"
    FRI = "FRI"
    SAT = "SAT"


class ScheduledAuditFrequency(str, Enum):
    """
    How often the scheduled audit takes place. Can be one of DAILY, WEEKLY, BIWEEKLY, or MONTHLY.
    """
    DAILY = "DAILY"
    WEEKLY = "WEEKLY"
    BIWEEKLY = "BIWEEKLY"
    MONTHLY = "MONTHLY"


class SecurityProfileBehaviorCriteriaComparisonOperator(str, Enum):
    """
    The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).
    """
    LESS_THAN = "less-than"
    LESS_THAN_EQUALS = "less-than-equals"
    GREATER_THAN = "greater-than"
    GREATER_THAN_EQUALS = "greater-than-equals"
    IN_CIDR_SET = "in-cidr-set"
    NOT_IN_CIDR_SET = "not-in-cidr-set"
    IN_PORT_SET = "in-port-set"
    NOT_IN_PORT_SET = "not-in-port-set"
    IN_SET = "in-set"
    NOT_IN_SET = "not-in-set"


class SecurityProfileMachineLearningDetectionConfigConfidenceLevel(str, Enum):
    """
    The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.
    """
    LOW = "LOW"
    MEDIUM = "MEDIUM"
    HIGH = "HIGH"


class SecurityProfileMetricDimensionOperator(str, Enum):
    """
    Defines how the dimensionValues of a dimension are interpreted.
    """
    IN_ = "IN"
    NOT_IN = "NOT_IN"


class SecurityProfileStatisticalThresholdStatistic(str, Enum):
    """
    The percentile which resolves to a threshold value by which compliance with a behavior is determined
    """
    AVERAGE = "Average"
    P0 = "p0"
    P01 = "p0.1"
    P001 = "p0.01"
    P1 = "p1"
    P10 = "p10"
    P50 = "p50"
    P90 = "p90"
    P99 = "p99"
    P999 = "p99.9"
    P9999 = "p99.99"
    P100 = "p100"


class TopicRuleCannedAccessControlList(str, Enum):
    PRIVATE = "private"
    PUBLIC_READ = "public-read"
    PUBLIC_READ_WRITE = "public-read-write"
    AWS_EXEC_READ = "aws-exec-read"
    AUTHENTICATED_READ = "authenticated-read"
    BUCKET_OWNER_READ = "bucket-owner-read"
    BUCKET_OWNER_FULL_CONTROL = "bucket-owner-full-control"
    LOG_DELIVERY_WRITE = "log-delivery-write"


class TopicRuleDestinationStatus(str, Enum):
    ENABLED = "ENABLED"
    IN_PROGRESS = "IN_PROGRESS"
    DISABLED = "DISABLED"
