# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'HoursOfOperationConfigDay',
    'QuickConnectType',
    'UserPhoneType',
]


class HoursOfOperationConfigDay(str, Enum):
    """
    The day that the hours of operation applies to.
    """
    SUNDAY = "SUNDAY"
    MONDAY = "MONDAY"
    TUESDAY = "TUESDAY"
    WEDNESDAY = "WEDNESDAY"
    THURSDAY = "THURSDAY"
    FRIDAY = "FRIDAY"
    SATURDAY = "SATURDAY"


class QuickConnectType(str, Enum):
    """
    The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).
    """
    PHONE_NUMBER = "PHONE_NUMBER"
    QUEUE = "QUEUE"
    USER = "USER"


class UserPhoneType(str, Enum):
    """
    The phone type.
    """
    SOFT_PHONE = "SOFT_PHONE"
    DESK_PHONE = "DESK_PHONE"
