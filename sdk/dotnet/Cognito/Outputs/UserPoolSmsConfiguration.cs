// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolSmsConfiguration
    {
        public readonly string? ExternalId;
        public readonly string? SnsCallerArn;

        [OutputConstructor]
        private UserPoolSmsConfiguration(
            string? externalId,

            string? snsCallerArn)
        {
            ExternalId = externalId;
            SnsCallerArn = snsCallerArn;
        }
    }
}
