// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const IPSetIPAddressVersion = {
    Ipv4: "IPV4",
    Ipv6: "IPV6",
} as const;

/**
 * Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.
 */
export type IPSetIPAddressVersion = (typeof IPSetIPAddressVersion)[keyof typeof IPSetIPAddressVersion];

export const IPSetScope = {
    Cloudfront: "CLOUDFRONT",
    Regional: "REGIONAL",
} as const;

/**
 * Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.
 */
export type IPSetScope = (typeof IPSetScope)[keyof typeof IPSetScope];

export const LoggingConfigurationConditionActionConditionPropertiesAction = {
    Allow: "ALLOW",
    Block: "BLOCK",
    Count: "COUNT",
} as const;

/**
 * Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
 */
export type LoggingConfigurationConditionActionConditionPropertiesAction = (typeof LoggingConfigurationConditionActionConditionPropertiesAction)[keyof typeof LoggingConfigurationConditionActionConditionPropertiesAction];

export const LoggingConfigurationFieldToMatchJsonBodyPropertiesInvalidFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
    EvaluateAsString: "EVALUATE_AS_STRING",
} as const;

/**
 * What AWS WAF should do if it fails to completely parse the JSON body.
 */
export type LoggingConfigurationFieldToMatchJsonBodyPropertiesInvalidFallbackBehavior = (typeof LoggingConfigurationFieldToMatchJsonBodyPropertiesInvalidFallbackBehavior)[keyof typeof LoggingConfigurationFieldToMatchJsonBodyPropertiesInvalidFallbackBehavior];

export const LoggingConfigurationFieldToMatchJsonBodyPropertiesMatchScope = {
    All: "ALL",
    Key: "KEY",
    Value: "VALUE",
} as const;

/**
 * The parts of the JSON to match against using the MatchPattern. If you specify All, AWS WAF matches against keys and values. 
 */
export type LoggingConfigurationFieldToMatchJsonBodyPropertiesMatchScope = (typeof LoggingConfigurationFieldToMatchJsonBodyPropertiesMatchScope)[keyof typeof LoggingConfigurationFieldToMatchJsonBodyPropertiesMatchScope];

export const LoggingConfigurationFilterBehavior = {
    Keep: "KEEP",
    Drop: "DROP",
} as const;

/**
 * How to handle logs that satisfy the filter's conditions and requirement. 
 */
export type LoggingConfigurationFilterBehavior = (typeof LoggingConfigurationFilterBehavior)[keyof typeof LoggingConfigurationFilterBehavior];

export const LoggingConfigurationFilterRequirement = {
    MeetsAll: "MEETS_ALL",
    MeetsAny: "MEETS_ANY",
} as const;

/**
 * Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
 */
export type LoggingConfigurationFilterRequirement = (typeof LoggingConfigurationFilterRequirement)[keyof typeof LoggingConfigurationFilterRequirement];

export const LoggingConfigurationLoggingFilterPropertiesDefaultBehavior = {
    Keep: "KEEP",
    Drop: "DROP",
} as const;

/**
 * Default handling for logs that don't match any of the specified filtering conditions.
 */
export type LoggingConfigurationLoggingFilterPropertiesDefaultBehavior = (typeof LoggingConfigurationLoggingFilterPropertiesDefaultBehavior)[keyof typeof LoggingConfigurationLoggingFilterPropertiesDefaultBehavior];

export const RegexPatternSetScope = {
    Cloudfront: "CLOUDFRONT",
    Regional: "REGIONAL",
} as const;

/**
 * Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
 */
export type RegexPatternSetScope = (typeof RegexPatternSetScope)[keyof typeof RegexPatternSetScope];

export const RuleGroupBodyParsingFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
    EvaluateAsString: "EVALUATE_AS_STRING",
} as const;

/**
 * The inspection behavior to fall back to if the JSON in the request body is invalid.
 */
export type RuleGroupBodyParsingFallbackBehavior = (typeof RuleGroupBodyParsingFallbackBehavior)[keyof typeof RuleGroupBodyParsingFallbackBehavior];

export const RuleGroupForwardedIPConfigurationFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

export type RuleGroupForwardedIPConfigurationFallbackBehavior = (typeof RuleGroupForwardedIPConfigurationFallbackBehavior)[keyof typeof RuleGroupForwardedIPConfigurationFallbackBehavior];

export const RuleGroupIPSetForwardedIPConfigurationFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

export type RuleGroupIPSetForwardedIPConfigurationFallbackBehavior = (typeof RuleGroupIPSetForwardedIPConfigurationFallbackBehavior)[keyof typeof RuleGroupIPSetForwardedIPConfigurationFallbackBehavior];

export const RuleGroupIPSetForwardedIPConfigurationPosition = {
    First: "FIRST",
    Last: "LAST",
    Any: "ANY",
} as const;

export type RuleGroupIPSetForwardedIPConfigurationPosition = (typeof RuleGroupIPSetForwardedIPConfigurationPosition)[keyof typeof RuleGroupIPSetForwardedIPConfigurationPosition];

export const RuleGroupJsonMatchScope = {
    All: "ALL",
    Key: "KEY",
    Value: "VALUE",
} as const;

/**
 * The parts of the JSON to match against using the MatchPattern.
 */
export type RuleGroupJsonMatchScope = (typeof RuleGroupJsonMatchScope)[keyof typeof RuleGroupJsonMatchScope];

export const RuleGroupLabelMatchScope = {
    Label: "LABEL",
    Namespace: "NAMESPACE",
} as const;

export type RuleGroupLabelMatchScope = (typeof RuleGroupLabelMatchScope)[keyof typeof RuleGroupLabelMatchScope];

export const RuleGroupPositionalConstraint = {
    Exactly: "EXACTLY",
    StartsWith: "STARTS_WITH",
    EndsWith: "ENDS_WITH",
    Contains: "CONTAINS",
    ContainsWord: "CONTAINS_WORD",
} as const;

/**
 * Position of the evaluation in the FieldToMatch of request.
 */
export type RuleGroupPositionalConstraint = (typeof RuleGroupPositionalConstraint)[keyof typeof RuleGroupPositionalConstraint];

export const RuleGroupRateBasedStatementAggregateKeyType = {
    Ip: "IP",
    ForwardedIp: "FORWARDED_IP",
} as const;

export type RuleGroupRateBasedStatementAggregateKeyType = (typeof RuleGroupRateBasedStatementAggregateKeyType)[keyof typeof RuleGroupRateBasedStatementAggregateKeyType];

export const RuleGroupScope = {
    Cloudfront: "CLOUDFRONT",
    Regional: "REGIONAL",
} as const;

/**
 * Use CLOUDFRONT for CloudFront RuleGroup, use REGIONAL for Application Load Balancer and API Gateway.
 */
export type RuleGroupScope = (typeof RuleGroupScope)[keyof typeof RuleGroupScope];

export const RuleGroupSizeConstraintStatementComparisonOperator = {
    Eq: "EQ",
    Ne: "NE",
    Le: "LE",
    Lt: "LT",
    Ge: "GE",
    Gt: "GT",
} as const;

export type RuleGroupSizeConstraintStatementComparisonOperator = (typeof RuleGroupSizeConstraintStatementComparisonOperator)[keyof typeof RuleGroupSizeConstraintStatementComparisonOperator];

export const RuleGroupTextTransformationType = {
    None: "NONE",
    CompressWhiteSpace: "COMPRESS_WHITE_SPACE",
    HtmlEntityDecode: "HTML_ENTITY_DECODE",
    Lowercase: "LOWERCASE",
    CmdLine: "CMD_LINE",
    UrlDecode: "URL_DECODE",
    Base64Decode: "BASE64_DECODE",
    HexDecode: "HEX_DECODE",
    Md5: "MD5",
    ReplaceComments: "REPLACE_COMMENTS",
    EscapeSeqDecode: "ESCAPE_SEQ_DECODE",
    SqlHexDecode: "SQL_HEX_DECODE",
    CssDecode: "CSS_DECODE",
    JsDecode: "JS_DECODE",
    NormalizePath: "NORMALIZE_PATH",
    NormalizePathWin: "NORMALIZE_PATH_WIN",
    RemoveNulls: "REMOVE_NULLS",
    ReplaceNulls: "REPLACE_NULLS",
    Base64DecodeExt: "BASE64_DECODE_EXT",
    UrlDecodeUni: "URL_DECODE_UNI",
    Utf8ToUnicode: "UTF8_TO_UNICODE",
} as const;

/**
 * Type of text transformation.
 */
export type RuleGroupTextTransformationType = (typeof RuleGroupTextTransformationType)[keyof typeof RuleGroupTextTransformationType];

export const WebACLBodyParsingFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
    EvaluateAsString: "EVALUATE_AS_STRING",
} as const;

/**
 * The inspection behavior to fall back to if the JSON in the request body is invalid.
 */
export type WebACLBodyParsingFallbackBehavior = (typeof WebACLBodyParsingFallbackBehavior)[keyof typeof WebACLBodyParsingFallbackBehavior];

export const WebACLForwardedIPConfigurationFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

export type WebACLForwardedIPConfigurationFallbackBehavior = (typeof WebACLForwardedIPConfigurationFallbackBehavior)[keyof typeof WebACLForwardedIPConfigurationFallbackBehavior];

export const WebACLIPSetForwardedIPConfigurationFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

export type WebACLIPSetForwardedIPConfigurationFallbackBehavior = (typeof WebACLIPSetForwardedIPConfigurationFallbackBehavior)[keyof typeof WebACLIPSetForwardedIPConfigurationFallbackBehavior];

export const WebACLIPSetForwardedIPConfigurationPosition = {
    First: "FIRST",
    Last: "LAST",
    Any: "ANY",
} as const;

export type WebACLIPSetForwardedIPConfigurationPosition = (typeof WebACLIPSetForwardedIPConfigurationPosition)[keyof typeof WebACLIPSetForwardedIPConfigurationPosition];

export const WebACLJsonMatchScope = {
    All: "ALL",
    Key: "KEY",
    Value: "VALUE",
} as const;

/**
 * The parts of the JSON to match against using the MatchPattern.
 */
export type WebACLJsonMatchScope = (typeof WebACLJsonMatchScope)[keyof typeof WebACLJsonMatchScope];

export const WebACLLabelMatchScope = {
    Label: "LABEL",
    Namespace: "NAMESPACE",
} as const;

export type WebACLLabelMatchScope = (typeof WebACLLabelMatchScope)[keyof typeof WebACLLabelMatchScope];

export const WebACLPositionalConstraint = {
    Exactly: "EXACTLY",
    StartsWith: "STARTS_WITH",
    EndsWith: "ENDS_WITH",
    Contains: "CONTAINS",
    ContainsWord: "CONTAINS_WORD",
} as const;

/**
 * Position of the evaluation in the FieldToMatch of request.
 */
export type WebACLPositionalConstraint = (typeof WebACLPositionalConstraint)[keyof typeof WebACLPositionalConstraint];

export const WebACLRateBasedStatementAggregateKeyType = {
    Ip: "IP",
    ForwardedIp: "FORWARDED_IP",
} as const;

export type WebACLRateBasedStatementAggregateKeyType = (typeof WebACLRateBasedStatementAggregateKeyType)[keyof typeof WebACLRateBasedStatementAggregateKeyType];

export const WebACLScope = {
    Cloudfront: "CLOUDFRONT",
    Regional: "REGIONAL",
} as const;

/**
 * Use CLOUDFRONT for CloudFront WebACL, use REGIONAL for Application Load Balancer and API Gateway.
 */
export type WebACLScope = (typeof WebACLScope)[keyof typeof WebACLScope];

export const WebACLSizeConstraintStatementComparisonOperator = {
    Eq: "EQ",
    Ne: "NE",
    Le: "LE",
    Lt: "LT",
    Ge: "GE",
    Gt: "GT",
} as const;

export type WebACLSizeConstraintStatementComparisonOperator = (typeof WebACLSizeConstraintStatementComparisonOperator)[keyof typeof WebACLSizeConstraintStatementComparisonOperator];

export const WebACLTextTransformationType = {
    None: "NONE",
    CompressWhiteSpace: "COMPRESS_WHITE_SPACE",
    HtmlEntityDecode: "HTML_ENTITY_DECODE",
    Lowercase: "LOWERCASE",
    CmdLine: "CMD_LINE",
    UrlDecode: "URL_DECODE",
    Base64Decode: "BASE64_DECODE",
    HexDecode: "HEX_DECODE",
    Md5: "MD5",
    ReplaceComments: "REPLACE_COMMENTS",
    EscapeSeqDecode: "ESCAPE_SEQ_DECODE",
    SqlHexDecode: "SQL_HEX_DECODE",
    CssDecode: "CSS_DECODE",
    JsDecode: "JS_DECODE",
    NormalizePath: "NORMALIZE_PATH",
    NormalizePathWin: "NORMALIZE_PATH_WIN",
    RemoveNulls: "REMOVE_NULLS",
    ReplaceNulls: "REPLACE_NULLS",
    Base64DecodeExt: "BASE64_DECODE_EXT",
    UrlDecodeUni: "URL_DECODE_UNI",
    Utf8ToUnicode: "UTF8_TO_UNICODE",
} as const;

/**
 * Type of text transformation.
 */
export type WebACLTextTransformationType = (typeof WebACLTextTransformationType)[keyof typeof WebACLTextTransformationType];
