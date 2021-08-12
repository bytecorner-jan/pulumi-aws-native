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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databasetableoutputoptions.html
    /// </summary>
    [OutputType]
    public sealed class JobDatabaseTableOutputOptions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databasetableoutputoptions.html#cfn-databrew-job-databasetableoutputoptions-tablename
        /// </summary>
        public readonly string TableName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databasetableoutputoptions.html#cfn-databrew-job-databasetableoutputoptions-tempdirectory
        /// </summary>
        public readonly Outputs.JobS3Location? TempDirectory;

        [OutputConstructor]
        private JobDatabaseTableOutputOptions(
            string tableName,

            Outputs.JobS3Location? tempDirectory)
        {
            TableName = tableName;
            TempDirectory = tempDirectory;
        }
    }
}
