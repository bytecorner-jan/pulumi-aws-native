// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    [OutputType]
    public sealed class DatasetDatabaseInputDefinition
    {
        /// <summary>
        /// Database table name
        /// </summary>
        public readonly string? DatabaseTableName;
        /// <summary>
        /// Glue connection name
        /// </summary>
        public readonly string? GlueConnectionName;
        public readonly Outputs.DatasetS3Location? TempDirectory;

        [OutputConstructor]
        private DatasetDatabaseInputDefinition(
            string? databaseTableName,

            string? glueConnectionName,

            Outputs.DatasetS3Location? tempDirectory)
        {
            DatabaseTableName = databaseTableName;
            GlueConnectionName = glueConnectionName;
            TempDirectory = tempDirectory;
        }
    }
}
