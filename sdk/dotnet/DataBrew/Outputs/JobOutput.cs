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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html
    /// </summary>
    [OutputType]
    public sealed class JobOutput
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-compressionformat
        /// </summary>
        public readonly string? CompressionFormat;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-format
        /// </summary>
        public readonly string? Format;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-formatoptions
        /// </summary>
        public readonly Outputs.JobOutputFormatOptions? FormatOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-location
        /// </summary>
        public readonly Outputs.JobS3Location Location;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-overwrite
        /// </summary>
        public readonly bool? Overwrite;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-partitioncolumns
        /// </summary>
        public readonly ImmutableArray<string> PartitionColumns;

        [OutputConstructor]
        private JobOutput(
            string? compressionFormat,

            string? format,

            Outputs.JobOutputFormatOptions? formatOptions,

            Outputs.JobS3Location location,

            bool? overwrite,

            ImmutableArray<string> partitionColumns)
        {
            CompressionFormat = compressionFormat;
            Format = format;
            FormatOptions = formatOptions;
            Location = location;
            Overwrite = overwrite;
            PartitionColumns = partitionColumns;
        }
    }
}
