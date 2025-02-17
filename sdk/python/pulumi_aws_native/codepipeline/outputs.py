# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'CustomActionTypeArtifactDetails',
    'CustomActionTypeConfigurationProperties',
    'CustomActionTypeSettings',
    'CustomActionTypeTag',
    'PipelineActionDeclaration',
    'PipelineActionTypeId',
    'PipelineArtifactStore',
    'PipelineArtifactStoreMap',
    'PipelineBlockerDeclaration',
    'PipelineEncryptionKey',
    'PipelineInputArtifact',
    'PipelineOutputArtifact',
    'PipelineStageDeclaration',
    'PipelineStageTransition',
    'PipelineTag',
    'WebhookAuthConfiguration',
    'WebhookFilterRule',
]

@pulumi.output_type
class CustomActionTypeArtifactDetails(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maximumCount":
            suggest = "maximum_count"
        elif key == "minimumCount":
            suggest = "minimum_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomActionTypeArtifactDetails. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomActionTypeArtifactDetails.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomActionTypeArtifactDetails.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 maximum_count: int,
                 minimum_count: int):
        pulumi.set(__self__, "maximum_count", maximum_count)
        pulumi.set(__self__, "minimum_count", minimum_count)

    @property
    @pulumi.getter(name="maximumCount")
    def maximum_count(self) -> int:
        return pulumi.get(self, "maximum_count")

    @property
    @pulumi.getter(name="minimumCount")
    def minimum_count(self) -> int:
        return pulumi.get(self, "minimum_count")


@pulumi.output_type
class CustomActionTypeConfigurationProperties(dict):
    def __init__(__self__, *,
                 key: bool,
                 name: str,
                 required: bool,
                 secret: bool,
                 description: Optional[str] = None,
                 queryable: Optional[bool] = None,
                 type: Optional[str] = None):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "required", required)
        pulumi.set(__self__, "secret", secret)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if queryable is not None:
            pulumi.set(__self__, "queryable", queryable)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def key(self) -> bool:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def required(self) -> bool:
        return pulumi.get(self, "required")

    @property
    @pulumi.getter
    def secret(self) -> bool:
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def queryable(self) -> Optional[bool]:
        return pulumi.get(self, "queryable")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


@pulumi.output_type
class CustomActionTypeSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "entityUrlTemplate":
            suggest = "entity_url_template"
        elif key == "executionUrlTemplate":
            suggest = "execution_url_template"
        elif key == "revisionUrlTemplate":
            suggest = "revision_url_template"
        elif key == "thirdPartyConfigurationUrl":
            suggest = "third_party_configuration_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomActionTypeSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomActionTypeSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomActionTypeSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 entity_url_template: Optional[str] = None,
                 execution_url_template: Optional[str] = None,
                 revision_url_template: Optional[str] = None,
                 third_party_configuration_url: Optional[str] = None):
        if entity_url_template is not None:
            pulumi.set(__self__, "entity_url_template", entity_url_template)
        if execution_url_template is not None:
            pulumi.set(__self__, "execution_url_template", execution_url_template)
        if revision_url_template is not None:
            pulumi.set(__self__, "revision_url_template", revision_url_template)
        if third_party_configuration_url is not None:
            pulumi.set(__self__, "third_party_configuration_url", third_party_configuration_url)

    @property
    @pulumi.getter(name="entityUrlTemplate")
    def entity_url_template(self) -> Optional[str]:
        return pulumi.get(self, "entity_url_template")

    @property
    @pulumi.getter(name="executionUrlTemplate")
    def execution_url_template(self) -> Optional[str]:
        return pulumi.get(self, "execution_url_template")

    @property
    @pulumi.getter(name="revisionUrlTemplate")
    def revision_url_template(self) -> Optional[str]:
        return pulumi.get(self, "revision_url_template")

    @property
    @pulumi.getter(name="thirdPartyConfigurationUrl")
    def third_party_configuration_url(self) -> Optional[str]:
        return pulumi.get(self, "third_party_configuration_url")


@pulumi.output_type
class CustomActionTypeTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class PipelineActionDeclaration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionTypeId":
            suggest = "action_type_id"
        elif key == "inputArtifacts":
            suggest = "input_artifacts"
        elif key == "outputArtifacts":
            suggest = "output_artifacts"
        elif key == "roleArn":
            suggest = "role_arn"
        elif key == "runOrder":
            suggest = "run_order"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PipelineActionDeclaration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PipelineActionDeclaration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PipelineActionDeclaration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action_type_id: 'outputs.PipelineActionTypeId',
                 name: str,
                 configuration: Optional[Any] = None,
                 input_artifacts: Optional[Sequence['outputs.PipelineInputArtifact']] = None,
                 namespace: Optional[str] = None,
                 output_artifacts: Optional[Sequence['outputs.PipelineOutputArtifact']] = None,
                 region: Optional[str] = None,
                 role_arn: Optional[str] = None,
                 run_order: Optional[int] = None):
        pulumi.set(__self__, "action_type_id", action_type_id)
        pulumi.set(__self__, "name", name)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if input_artifacts is not None:
            pulumi.set(__self__, "input_artifacts", input_artifacts)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if output_artifacts is not None:
            pulumi.set(__self__, "output_artifacts", output_artifacts)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if run_order is not None:
            pulumi.set(__self__, "run_order", run_order)

    @property
    @pulumi.getter(name="actionTypeId")
    def action_type_id(self) -> 'outputs.PipelineActionTypeId':
        return pulumi.get(self, "action_type_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def configuration(self) -> Optional[Any]:
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="inputArtifacts")
    def input_artifacts(self) -> Optional[Sequence['outputs.PipelineInputArtifact']]:
        return pulumi.get(self, "input_artifacts")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="outputArtifacts")
    def output_artifacts(self) -> Optional[Sequence['outputs.PipelineOutputArtifact']]:
        return pulumi.get(self, "output_artifacts")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="runOrder")
    def run_order(self) -> Optional[int]:
        return pulumi.get(self, "run_order")


@pulumi.output_type
class PipelineActionTypeId(dict):
    def __init__(__self__, *,
                 category: str,
                 owner: str,
                 provider: str,
                 version: str):
        pulumi.set(__self__, "category", category)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "provider", provider)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def category(self) -> str:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def owner(self) -> str:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def provider(self) -> str:
        return pulumi.get(self, "provider")

    @property
    @pulumi.getter
    def version(self) -> str:
        return pulumi.get(self, "version")


@pulumi.output_type
class PipelineArtifactStore(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionKey":
            suggest = "encryption_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PipelineArtifactStore. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PipelineArtifactStore.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PipelineArtifactStore.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 location: str,
                 type: str,
                 encryption_key: Optional['outputs.PipelineEncryptionKey'] = None):
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "type", type)
        if encryption_key is not None:
            pulumi.set(__self__, "encryption_key", encryption_key)

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional['outputs.PipelineEncryptionKey']:
        return pulumi.get(self, "encryption_key")


@pulumi.output_type
class PipelineArtifactStoreMap(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "artifactStore":
            suggest = "artifact_store"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PipelineArtifactStoreMap. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PipelineArtifactStoreMap.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PipelineArtifactStoreMap.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 artifact_store: 'outputs.PipelineArtifactStore',
                 region: str):
        pulumi.set(__self__, "artifact_store", artifact_store)
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="artifactStore")
    def artifact_store(self) -> 'outputs.PipelineArtifactStore':
        return pulumi.get(self, "artifact_store")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


@pulumi.output_type
class PipelineBlockerDeclaration(dict):
    def __init__(__self__, *,
                 name: str,
                 type: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class PipelineEncryptionKey(dict):
    def __init__(__self__, *,
                 id: str,
                 type: str):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class PipelineInputArtifact(dict):
    def __init__(__self__, *,
                 name: str):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


@pulumi.output_type
class PipelineOutputArtifact(dict):
    def __init__(__self__, *,
                 name: str):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


@pulumi.output_type
class PipelineStageDeclaration(dict):
    def __init__(__self__, *,
                 actions: Sequence['outputs.PipelineActionDeclaration'],
                 name: str,
                 blockers: Optional[Sequence['outputs.PipelineBlockerDeclaration']] = None):
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "name", name)
        if blockers is not None:
            pulumi.set(__self__, "blockers", blockers)

    @property
    @pulumi.getter
    def actions(self) -> Sequence['outputs.PipelineActionDeclaration']:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def blockers(self) -> Optional[Sequence['outputs.PipelineBlockerDeclaration']]:
        return pulumi.get(self, "blockers")


@pulumi.output_type
class PipelineStageTransition(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "stageName":
            suggest = "stage_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PipelineStageTransition. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PipelineStageTransition.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PipelineStageTransition.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 reason: str,
                 stage_name: str):
        pulumi.set(__self__, "reason", reason)
        pulumi.set(__self__, "stage_name", stage_name)

    @property
    @pulumi.getter
    def reason(self) -> str:
        return pulumi.get(self, "reason")

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> str:
        return pulumi.get(self, "stage_name")


@pulumi.output_type
class PipelineTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class WebhookAuthConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowedIPRange":
            suggest = "allowed_ip_range"
        elif key == "secretToken":
            suggest = "secret_token"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WebhookAuthConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WebhookAuthConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WebhookAuthConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allowed_ip_range: Optional[str] = None,
                 secret_token: Optional[str] = None):
        if allowed_ip_range is not None:
            pulumi.set(__self__, "allowed_ip_range", allowed_ip_range)
        if secret_token is not None:
            pulumi.set(__self__, "secret_token", secret_token)

    @property
    @pulumi.getter(name="allowedIPRange")
    def allowed_ip_range(self) -> Optional[str]:
        return pulumi.get(self, "allowed_ip_range")

    @property
    @pulumi.getter(name="secretToken")
    def secret_token(self) -> Optional[str]:
        return pulumi.get(self, "secret_token")


@pulumi.output_type
class WebhookFilterRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "jsonPath":
            suggest = "json_path"
        elif key == "matchEquals":
            suggest = "match_equals"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WebhookFilterRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WebhookFilterRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WebhookFilterRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 json_path: str,
                 match_equals: Optional[str] = None):
        pulumi.set(__self__, "json_path", json_path)
        if match_equals is not None:
            pulumi.set(__self__, "match_equals", match_equals)

    @property
    @pulumi.getter(name="jsonPath")
    def json_path(self) -> str:
        return pulumi.get(self, "json_path")

    @property
    @pulumi.getter(name="matchEquals")
    def match_equals(self) -> Optional[str]:
        return pulumi.get(self, "match_equals")


