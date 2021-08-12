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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceSalesforceStandardKnowledgeArticleTypeConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration.html#cfn-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration-documentdatafieldname
        /// </summary>
        public readonly string DocumentDataFieldName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration.html#cfn-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration-documenttitlefieldname
        /// </summary>
        public readonly string? DocumentTitleFieldName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration.html#cfn-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration-fieldmappings
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceDataSourceToIndexFieldMapping> FieldMappings;

        [OutputConstructor]
        private DataSourceSalesforceStandardKnowledgeArticleTypeConfiguration(
            string documentDataFieldName,

            string? documentTitleFieldName,

            ImmutableArray<Outputs.DataSourceDataSourceToIndexFieldMapping> fieldMappings)
        {
            DocumentDataFieldName = documentDataFieldName;
            DocumentTitleFieldName = documentTitleFieldName;
            FieldMappings = fieldMappings;
        }
    }
}
