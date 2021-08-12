// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceColumnConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-changedetectingcolumns
        /// </summary>
        public readonly ImmutableArray<string> ChangeDetectingColumns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-documentdatacolumnname
        /// </summary>
        public readonly string DocumentDataColumnName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-documentidcolumnname
        /// </summary>
        public readonly string DocumentIdColumnName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-documenttitlecolumnname
        /// </summary>
        public readonly string? DocumentTitleColumnName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-fieldmappings
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceDataSourceToIndexFieldMapping> FieldMappings;

        [OutputConstructor]
        private DataSourceColumnConfiguration(
            ImmutableArray<string> changeDetectingColumns,

            string documentDataColumnName,

            string documentIdColumnName,

            string? documentTitleColumnName,

            ImmutableArray<Outputs.DataSourceDataSourceToIndexFieldMapping> fieldMappings)
        {
            ChangeDetectingColumns = changeDetectingColumns;
            DocumentDataColumnName = documentDataColumnName;
            DocumentIdColumnName = documentIdColumnName;
            DocumentTitleColumnName = documentTitleColumnName;
            FieldMappings = fieldMappings;
        }
    }
}
