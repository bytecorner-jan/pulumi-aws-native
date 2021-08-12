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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-mariadbparameters.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceMariaDbParameters
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-mariadbparameters.html#cfn-quicksight-datasource-mariadbparameters-database
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-mariadbparameters.html#cfn-quicksight-datasource-mariadbparameters-host
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-mariadbparameters.html#cfn-quicksight-datasource-mariadbparameters-port
        /// </summary>
        public readonly double Port;

        [OutputConstructor]
        private DataSourceMariaDbParameters(
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
