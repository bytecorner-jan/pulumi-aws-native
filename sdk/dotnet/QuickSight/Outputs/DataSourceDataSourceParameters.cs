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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceDataSourceParameters
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-amazonelasticsearchparameters
        /// </summary>
        public readonly Outputs.DataSourceAmazonElasticsearchParameters? AmazonElasticsearchParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-athenaparameters
        /// </summary>
        public readonly Outputs.DataSourceAthenaParameters? AthenaParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-auroraparameters
        /// </summary>
        public readonly Outputs.DataSourceAuroraParameters? AuroraParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-aurorapostgresqlparameters
        /// </summary>
        public readonly Outputs.DataSourceAuroraPostgreSqlParameters? AuroraPostgreSqlParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-mariadbparameters
        /// </summary>
        public readonly Outputs.DataSourceMariaDbParameters? MariaDbParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-mysqlparameters
        /// </summary>
        public readonly Outputs.DataSourceMySqlParameters? MySqlParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-oracleparameters
        /// </summary>
        public readonly Outputs.DataSourceOracleParameters? OracleParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-postgresqlparameters
        /// </summary>
        public readonly Outputs.DataSourcePostgreSqlParameters? PostgreSqlParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-prestoparameters
        /// </summary>
        public readonly Outputs.DataSourcePrestoParameters? PrestoParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-rdsparameters
        /// </summary>
        public readonly Outputs.DataSourceRdsParameters? RdsParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-redshiftparameters
        /// </summary>
        public readonly Outputs.DataSourceRedshiftParameters? RedshiftParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-s3parameters
        /// </summary>
        public readonly Outputs.DataSourceS3Parameters? S3Parameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-snowflakeparameters
        /// </summary>
        public readonly Outputs.DataSourceSnowflakeParameters? SnowflakeParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-sparkparameters
        /// </summary>
        public readonly Outputs.DataSourceSparkParameters? SparkParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-sqlserverparameters
        /// </summary>
        public readonly Outputs.DataSourceSqlServerParameters? SqlServerParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-teradataparameters
        /// </summary>
        public readonly Outputs.DataSourceTeradataParameters? TeradataParameters;

        [OutputConstructor]
        private DataSourceDataSourceParameters(
            Outputs.DataSourceAmazonElasticsearchParameters? amazonElasticsearchParameters,

            Outputs.DataSourceAthenaParameters? athenaParameters,

            Outputs.DataSourceAuroraParameters? auroraParameters,

            Outputs.DataSourceAuroraPostgreSqlParameters? auroraPostgreSqlParameters,

            Outputs.DataSourceMariaDbParameters? mariaDbParameters,

            Outputs.DataSourceMySqlParameters? mySqlParameters,

            Outputs.DataSourceOracleParameters? oracleParameters,

            Outputs.DataSourcePostgreSqlParameters? postgreSqlParameters,

            Outputs.DataSourcePrestoParameters? prestoParameters,

            Outputs.DataSourceRdsParameters? rdsParameters,

            Outputs.DataSourceRedshiftParameters? redshiftParameters,

            Outputs.DataSourceS3Parameters? s3Parameters,

            Outputs.DataSourceSnowflakeParameters? snowflakeParameters,

            Outputs.DataSourceSparkParameters? sparkParameters,

            Outputs.DataSourceSqlServerParameters? sqlServerParameters,

            Outputs.DataSourceTeradataParameters? teradataParameters)
        {
            AmazonElasticsearchParameters = amazonElasticsearchParameters;
            AthenaParameters = athenaParameters;
            AuroraParameters = auroraParameters;
            AuroraPostgreSqlParameters = auroraPostgreSqlParameters;
            MariaDbParameters = mariaDbParameters;
            MySqlParameters = mySqlParameters;
            OracleParameters = oracleParameters;
            PostgreSqlParameters = postgreSqlParameters;
            PrestoParameters = prestoParameters;
            RdsParameters = rdsParameters;
            RedshiftParameters = redshiftParameters;
            S3Parameters = s3Parameters;
            SnowflakeParameters = snowflakeParameters;
            SparkParameters = sparkParameters;
            SqlServerParameters = sqlServerParameters;
            TeradataParameters = teradataParameters;
        }
    }
}
