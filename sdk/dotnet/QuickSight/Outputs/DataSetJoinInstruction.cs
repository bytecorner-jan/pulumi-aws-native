// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html
    /// </summary>
    [OutputType]
    public sealed class DataSetJoinInstruction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-leftjoinkeyproperties
        /// </summary>
        public readonly Outputs.DataSetJoinKeyProperties? LeftJoinKeyProperties;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-leftoperand
        /// </summary>
        public readonly string LeftOperand;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-onclause
        /// </summary>
        public readonly string OnClause;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-rightjoinkeyproperties
        /// </summary>
        public readonly Outputs.DataSetJoinKeyProperties? RightJoinKeyProperties;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-rightoperand
        /// </summary>
        public readonly string RightOperand;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DataSetJoinInstruction(
            Outputs.DataSetJoinKeyProperties? leftJoinKeyProperties,

            string leftOperand,

            string onClause,

            Outputs.DataSetJoinKeyProperties? rightJoinKeyProperties,

            string rightOperand,

            string type)
        {
            LeftJoinKeyProperties = leftJoinKeyProperties;
            LeftOperand = leftOperand;
            OnClause = onClause;
            RightJoinKeyProperties = rightJoinKeyProperties;
            RightOperand = rightOperand;
            Type = type;
        }
    }
}
