// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonbody.html
    /// </summary>
    [OutputType]
    public sealed class WebACLJsonBody
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonbody.html#cfn-wafv2-webacl-jsonbody-invalidfallbackbehavior
        /// </summary>
        public readonly string? InvalidFallbackBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonbody.html#cfn-wafv2-webacl-jsonbody-matchpattern
        /// </summary>
        public readonly Outputs.WebACLJsonMatchPattern MatchPattern;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonbody.html#cfn-wafv2-webacl-jsonbody-matchscope
        /// </summary>
        public readonly string MatchScope;

        [OutputConstructor]
        private WebACLJsonBody(
            string? invalidFallbackBehavior,

            Outputs.WebACLJsonMatchPattern matchPattern,

            string matchScope)
        {
            InvalidFallbackBehavior = invalidFallbackBehavior;
            MatchPattern = matchPattern;
            MatchScope = matchScope;
        }
    }
}
