// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    [OutputType]
    public sealed class FunctionCode
    {
        /// <summary>
        /// ImageUri.
        /// </summary>
        public readonly string? ImageUri;
        /// <summary>
        /// An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.
        /// </summary>
        public readonly string? S3Bucket;
        /// <summary>
        /// The Amazon S3 key of the deployment package.
        /// </summary>
        public readonly string? S3Key;
        /// <summary>
        /// For versioned objects, the version of the deployment package object to use.
        /// </summary>
        public readonly string? S3ObjectVersion;
        /// <summary>
        /// The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..
        /// </summary>
        public readonly AssetOrArchive? ZipFile;

        [OutputConstructor]
        private FunctionCode(
            string? imageUri,

            string? s3Bucket,

            string? s3Key,

            string? s3ObjectVersion,

            AssetOrArchive? zipFile)
        {
            ImageUri = imageUri;
            S3Bucket = s3Bucket;
            S3Key = s3Key;
            S3ObjectVersion = s3ObjectVersion;
            ZipFile = zipFile;
        }
    }
}
