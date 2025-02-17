// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks.Outputs
{

    [OutputType]
    public sealed class AppEnvironmentVariable
    {
        public readonly string Key;
        public readonly bool? Secure;
        public readonly string Value;

        [OutputConstructor]
        private AppEnvironmentVariable(
            string key,

            bool? secure,

            string value)
        {
            Key = key;
            Secure = secure;
            Value = value;
        }
    }
}
