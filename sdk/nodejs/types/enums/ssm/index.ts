// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AssociationComplianceSeverity = {
    Critical: "CRITICAL",
    High: "HIGH",
    Medium: "MEDIUM",
    Low: "LOW",
    Unspecified: "UNSPECIFIED",
} as const;

export type AssociationComplianceSeverity = (typeof AssociationComplianceSeverity)[keyof typeof AssociationComplianceSeverity];

export const AssociationSyncCompliance = {
    Auto: "AUTO",
    Manual: "MANUAL",
} as const;

export type AssociationSyncCompliance = (typeof AssociationSyncCompliance)[keyof typeof AssociationSyncCompliance];

export const DocumentAttachmentsSourceKey = {
    SourceUrl: "SourceUrl",
    S3FileUrl: "S3FileUrl",
    AttachmentReference: "AttachmentReference",
} as const;

/**
 * The key of a key-value pair that identifies the location of an attachment to a document.
 */
export type DocumentAttachmentsSourceKey = (typeof DocumentAttachmentsSourceKey)[keyof typeof DocumentAttachmentsSourceKey];

export const DocumentFormat = {
    Yaml: "YAML",
    Json: "JSON",
    Text: "TEXT",
} as const;

/**
 * Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.
 */
export type DocumentFormat = (typeof DocumentFormat)[keyof typeof DocumentFormat];

export const DocumentType = {
    ApplicationConfiguration: "ApplicationConfiguration",
    ApplicationConfigurationSchema: "ApplicationConfigurationSchema",
    Automation: "Automation",
    AutomationChangeTemplate: "Automation.ChangeTemplate",
    ChangeCalendar: "ChangeCalendar",
    CloudFormation: "CloudFormation",
    Command: "Command",
    DeploymentStrategy: "DeploymentStrategy",
    Package: "Package",
    Policy: "Policy",
    ProblemAnalysis: "ProblemAnalysis",
    ProblemAnalysisTemplate: "ProblemAnalysisTemplate",
    Session: "Session",
} as const;

/**
 * The type of document to create.
 */
export type DocumentType = (typeof DocumentType)[keyof typeof DocumentType];
