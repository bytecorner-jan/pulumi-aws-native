// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html
    /// </summary>
    public sealed class DataSourceSalesforceChatterFeedConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html#cfn-kendra-datasource-salesforcechatterfeedconfiguration-documentdatafieldname
        /// </summary>
        [Input("documentDataFieldName", required: true)]
        public Input<string> DocumentDataFieldName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html#cfn-kendra-datasource-salesforcechatterfeedconfiguration-documenttitlefieldname
        /// </summary>
        [Input("documentTitleFieldName")]
        public Input<string>? DocumentTitleFieldName { get; set; }

        [Input("fieldMappings")]
        private InputList<Inputs.DataSourceDataSourceToIndexFieldMappingArgs>? _fieldMappings;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html#cfn-kendra-datasource-salesforcechatterfeedconfiguration-fieldmappings
        /// </summary>
        public InputList<Inputs.DataSourceDataSourceToIndexFieldMappingArgs> FieldMappings
        {
            get => _fieldMappings ?? (_fieldMappings = new InputList<Inputs.DataSourceDataSourceToIndexFieldMappingArgs>());
            set => _fieldMappings = value;
        }

        [Input("includeFilterTypes")]
        private InputList<string>? _includeFilterTypes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html#cfn-kendra-datasource-salesforcechatterfeedconfiguration-includefiltertypes
        /// </summary>
        public InputList<string> IncludeFilterTypes
        {
            get => _includeFilterTypes ?? (_includeFilterTypes = new InputList<string>());
            set => _includeFilterTypes = value;
        }

        public DataSourceSalesforceChatterFeedConfigurationArgs()
        {
        }
    }
}
