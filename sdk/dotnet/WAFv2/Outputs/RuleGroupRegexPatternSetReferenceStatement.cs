// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    [OutputType]
    public sealed class RuleGroupRegexPatternSetReferenceStatement
    {
        public readonly string Arn;
        public readonly Outputs.RuleGroupFieldToMatch FieldToMatch;
        public readonly ImmutableArray<Outputs.RuleGroupTextTransformation> TextTransformations;

        [OutputConstructor]
        private RuleGroupRegexPatternSetReferenceStatement(
            string arn,

            Outputs.RuleGroupFieldToMatch fieldToMatch,

            ImmutableArray<Outputs.RuleGroupTextTransformation> textTransformations)
        {
            Arn = arn;
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
