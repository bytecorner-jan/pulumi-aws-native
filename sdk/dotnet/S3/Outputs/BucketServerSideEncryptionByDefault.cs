// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
    /// </summary>
    [OutputType]
    public sealed class BucketServerSideEncryptionByDefault
    {
        /// <summary>
        /// "KMSMasterKeyID" can only be used when you set the value of SSEAlgorithm as aws:kms.
        /// </summary>
        public readonly string? KMSMasterKeyID;
        public readonly Pulumi.AwsNative.S3.BucketServerSideEncryptionByDefaultSSEAlgorithm SSEAlgorithm;

        [OutputConstructor]
        private BucketServerSideEncryptionByDefault(
            string? kMSMasterKeyID,

            Pulumi.AwsNative.S3.BucketServerSideEncryptionByDefaultSSEAlgorithm sSEAlgorithm)
        {
            KMSMasterKeyID = kMSMasterKeyID;
            SSEAlgorithm = sSEAlgorithm;
        }
    }
}
