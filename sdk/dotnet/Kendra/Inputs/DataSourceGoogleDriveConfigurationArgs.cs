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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html
    /// </summary>
    public sealed class DataSourceGoogleDriveConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("excludeMimeTypes")]
        private InputList<string>? _excludeMimeTypes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-excludemimetypes
        /// </summary>
        public InputList<string> ExcludeMimeTypes
        {
            get => _excludeMimeTypes ?? (_excludeMimeTypes = new InputList<string>());
            set => _excludeMimeTypes = value;
        }

        [Input("excludeSharedDrives")]
        private InputList<string>? _excludeSharedDrives;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-excludeshareddrives
        /// </summary>
        public InputList<string> ExcludeSharedDrives
        {
            get => _excludeSharedDrives ?? (_excludeSharedDrives = new InputList<string>());
            set => _excludeSharedDrives = value;
        }

        [Input("excludeUserAccounts")]
        private InputList<string>? _excludeUserAccounts;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-excludeuseraccounts
        /// </summary>
        public InputList<string> ExcludeUserAccounts
        {
            get => _excludeUserAccounts ?? (_excludeUserAccounts = new InputList<string>());
            set => _excludeUserAccounts = value;
        }

        [Input("exclusionPatterns")]
        private InputList<string>? _exclusionPatterns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-exclusionpatterns
        /// </summary>
        public InputList<string> ExclusionPatterns
        {
            get => _exclusionPatterns ?? (_exclusionPatterns = new InputList<string>());
            set => _exclusionPatterns = value;
        }

        [Input("fieldMappings")]
        private InputList<Inputs.DataSourceDataSourceToIndexFieldMappingArgs>? _fieldMappings;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-fieldmappings
        /// </summary>
        public InputList<Inputs.DataSourceDataSourceToIndexFieldMappingArgs> FieldMappings
        {
            get => _fieldMappings ?? (_fieldMappings = new InputList<Inputs.DataSourceDataSourceToIndexFieldMappingArgs>());
            set => _fieldMappings = value;
        }

        [Input("inclusionPatterns")]
        private InputList<string>? _inclusionPatterns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-inclusionpatterns
        /// </summary>
        public InputList<string> InclusionPatterns
        {
            get => _inclusionPatterns ?? (_inclusionPatterns = new InputList<string>());
            set => _inclusionPatterns = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-secretarn
        /// </summary>
        [Input("secretArn", required: true)]
        public Input<string> SecretArn { get; set; } = null!;

        public DataSourceGoogleDriveConfigurationArgs()
        {
        }
    }
}
