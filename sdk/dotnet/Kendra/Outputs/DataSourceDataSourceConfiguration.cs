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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceDataSourceConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-confluenceconfiguration
        /// </summary>
        public readonly Outputs.DataSourceConfluenceConfiguration? ConfluenceConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-databaseconfiguration
        /// </summary>
        public readonly Outputs.DataSourceDatabaseConfiguration? DatabaseConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-googledriveconfiguration
        /// </summary>
        public readonly Outputs.DataSourceGoogleDriveConfiguration? GoogleDriveConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-onedriveconfiguration
        /// </summary>
        public readonly Outputs.DataSourceOneDriveConfiguration? OneDriveConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-s3configuration
        /// </summary>
        public readonly Outputs.DataSourceS3DataSourceConfiguration? S3Configuration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-salesforceconfiguration
        /// </summary>
        public readonly Outputs.DataSourceSalesforceConfiguration? SalesforceConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-servicenowconfiguration
        /// </summary>
        public readonly Outputs.DataSourceServiceNowConfiguration? ServiceNowConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-sharepointconfiguration
        /// </summary>
        public readonly Outputs.DataSourceSharePointConfiguration? SharePointConfiguration;

        [OutputConstructor]
        private DataSourceDataSourceConfiguration(
            Outputs.DataSourceConfluenceConfiguration? confluenceConfiguration,

            Outputs.DataSourceDatabaseConfiguration? databaseConfiguration,

            Outputs.DataSourceGoogleDriveConfiguration? googleDriveConfiguration,

            Outputs.DataSourceOneDriveConfiguration? oneDriveConfiguration,

            Outputs.DataSourceS3DataSourceConfiguration? s3Configuration,

            Outputs.DataSourceSalesforceConfiguration? salesforceConfiguration,

            Outputs.DataSourceServiceNowConfiguration? serviceNowConfiguration,

            Outputs.DataSourceSharePointConfiguration? sharePointConfiguration)
        {
            ConfluenceConfiguration = confluenceConfiguration;
            DatabaseConfiguration = databaseConfiguration;
            GoogleDriveConfiguration = googleDriveConfiguration;
            OneDriveConfiguration = oneDriveConfiguration;
            S3Configuration = s3Configuration;
            SalesforceConfiguration = salesforceConfiguration;
            ServiceNowConfiguration = serviceNowConfiguration;
            SharePointConfiguration = sharePointConfiguration;
        }
    }
}
