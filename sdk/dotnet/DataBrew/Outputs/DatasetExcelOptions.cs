// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html
    /// </summary>
    [OutputType]
    public sealed class DatasetExcelOptions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html#cfn-databrew-dataset-exceloptions-headerrow
        /// </summary>
        public readonly bool? HeaderRow;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html#cfn-databrew-dataset-exceloptions-sheetindexes
        /// </summary>
        public readonly ImmutableArray<int> SheetIndexes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html#cfn-databrew-dataset-exceloptions-sheetnames
        /// </summary>
        public readonly ImmutableArray<string> SheetNames;

        [OutputConstructor]
        private DatasetExcelOptions(
            bool? headerRow,

            ImmutableArray<int> sheetIndexes,

            ImmutableArray<string> sheetNames)
        {
            HeaderRow = headerRow;
            SheetIndexes = sheetIndexes;
            SheetNames = sheetNames;
        }
    }
}
