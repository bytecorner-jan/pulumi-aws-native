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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-teradataparameters.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceTeradataParameters
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-teradataparameters.html#cfn-quicksight-datasource-teradataparameters-database
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-teradataparameters.html#cfn-quicksight-datasource-teradataparameters-host
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-teradataparameters.html#cfn-quicksight-datasource-teradataparameters-port
        /// </summary>
        public readonly double Port;

        [OutputConstructor]
        private DataSourceTeradataParameters(
            string database,

            string host,

            double port)
        {
            Database = database;
            Host = host;
            Port = port;
        }
    }
}
