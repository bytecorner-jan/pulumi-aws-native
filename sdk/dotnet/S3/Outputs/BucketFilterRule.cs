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
    /// Specifies the Amazon S3 object key name to filter on and whether to filter on the suffix or prefix of the key name.
    /// </summary>
    [OutputType]
    public sealed class BucketFilterRule
    {
        public readonly string Name;
        public readonly string Value;

        [OutputConstructor]
        private BucketFilterRule(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
