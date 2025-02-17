// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// First level statement that contains conditions, such as ByteMatch, SizeConstraint, etc
    /// </summary>
    [OutputType]
    public sealed class RuleGroupStatement
    {
        public readonly Outputs.RuleGroupAndStatement? AndStatement;
        public readonly Outputs.RuleGroupByteMatchStatement? ByteMatchStatement;
        public readonly Outputs.RuleGroupGeoMatchStatement? GeoMatchStatement;
        public readonly Outputs.RuleGroupIPSetReferenceStatement? IPSetReferenceStatement;
        public readonly Outputs.RuleGroupLabelMatchStatement? LabelMatchStatement;
        public readonly Outputs.RuleGroupNotStatement? NotStatement;
        public readonly Outputs.RuleGroupOrStatement? OrStatement;
        public readonly Outputs.RuleGroupRateBasedStatement? RateBasedStatement;
        public readonly Outputs.RuleGroupRegexPatternSetReferenceStatement? RegexPatternSetReferenceStatement;
        public readonly Outputs.RuleGroupSizeConstraintStatement? SizeConstraintStatement;
        public readonly Outputs.RuleGroupSqliMatchStatement? SqliMatchStatement;
        public readonly Outputs.RuleGroupXssMatchStatement? XssMatchStatement;

        [OutputConstructor]
        private RuleGroupStatement(
            Outputs.RuleGroupAndStatement? andStatement,

            Outputs.RuleGroupByteMatchStatement? byteMatchStatement,

            Outputs.RuleGroupGeoMatchStatement? geoMatchStatement,

            Outputs.RuleGroupIPSetReferenceStatement? iPSetReferenceStatement,

            Outputs.RuleGroupLabelMatchStatement? labelMatchStatement,

            Outputs.RuleGroupNotStatement? notStatement,

            Outputs.RuleGroupOrStatement? orStatement,

            Outputs.RuleGroupRateBasedStatement? rateBasedStatement,

            Outputs.RuleGroupRegexPatternSetReferenceStatement? regexPatternSetReferenceStatement,

            Outputs.RuleGroupSizeConstraintStatement? sizeConstraintStatement,

            Outputs.RuleGroupSqliMatchStatement? sqliMatchStatement,

            Outputs.RuleGroupXssMatchStatement? xssMatchStatement)
        {
            AndStatement = andStatement;
            ByteMatchStatement = byteMatchStatement;
            GeoMatchStatement = geoMatchStatement;
            IPSetReferenceStatement = iPSetReferenceStatement;
            LabelMatchStatement = labelMatchStatement;
            NotStatement = notStatement;
            OrStatement = orStatement;
            RateBasedStatement = rateBasedStatement;
            RegexPatternSetReferenceStatement = regexPatternSetReferenceStatement;
            SizeConstraintStatement = sizeConstraintStatement;
            SqliMatchStatement = sqliMatchStatement;
            XssMatchStatement = xssMatchStatement;
        }
    }
}
