// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream.Inputs
{

    public sealed class DirectoryConfigServiceAccountCredentialsArgs : Pulumi.ResourceArgs
    {
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("accountPassword", required: true)]
        public Input<string> AccountPassword { get; set; } = null!;

        public DirectoryConfigServiceAccountCredentialsArgs()
        {
        }
    }
}
