// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3ObjectLambda.Outputs
{

    [OutputType]
    public sealed class PolicyStatusProperties
    {
        /// <summary>
        /// Specifies whether the Object lambda Access Point Policy is Public or not. Object lambda Access Points are private by default.
        /// </summary>
        public readonly bool? IsPublic;

        [OutputConstructor]
        private PolicyStatusProperties(bool? isPublic)
        {
            IsPublic = isPublic;
        }
    }
}
