// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-sizeconstraintstatement.html
    /// </summary>
    public sealed class RuleGroupSizeConstraintStatementArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-sizeconstraintstatement.html#cfn-wafv2-rulegroup-sizeconstraintstatement-comparisonoperator
        /// </summary>
        [Input("comparisonOperator", required: true)]
        public Input<string> ComparisonOperator { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-sizeconstraintstatement.html#cfn-wafv2-rulegroup-sizeconstraintstatement-fieldtomatch
        /// </summary>
        [Input("fieldToMatch", required: true)]
        public Input<Inputs.RuleGroupFieldToMatchArgs> FieldToMatch { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-sizeconstraintstatement.html#cfn-wafv2-rulegroup-sizeconstraintstatement-size
        /// </summary>
        [Input("size", required: true)]
        public Input<double> Size { get; set; } = null!;

        [Input("textTransformations", required: true)]
        private InputList<Inputs.RuleGroupTextTransformationArgs>? _textTransformations;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-sizeconstraintstatement.html#cfn-wafv2-rulegroup-sizeconstraintstatement-texttransformations
        /// </summary>
        public InputList<Inputs.RuleGroupTextTransformationArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.RuleGroupTextTransformationArgs>());
            set => _textTransformations = value;
        }

        public RuleGroupSizeConstraintStatementArgs()
        {
        }
    }
}
