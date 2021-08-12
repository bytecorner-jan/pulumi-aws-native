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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html
    /// </summary>
    [OutputType]
    public sealed class DataSetRelationalTable
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-catalog
        /// </summary>
        public readonly string? Catalog;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-datasourcearn
        /// </summary>
        public readonly string DataSourceArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-inputcolumns
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSetInputColumn> InputColumns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-schema
        /// </summary>
        public readonly string? Schema;

        [OutputConstructor]
        private DataSetRelationalTable(
            string? catalog,

            string dataSourceArn,

            ImmutableArray<Outputs.DataSetInputColumn> inputColumns,

            string name,

            string? schema)
        {
            Catalog = catalog;
            DataSourceArn = dataSourceArn;
            InputColumns = inputColumns;
            Name = name;
            Schema = schema;
        }
    }
}
