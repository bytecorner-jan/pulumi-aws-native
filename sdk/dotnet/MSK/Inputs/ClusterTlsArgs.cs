// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MSK.Inputs
{

    public sealed class ClusterTlsArgs : Pulumi.ResourceArgs
    {
        [Input("certificateAuthorityArnList")]
        private InputList<string>? _certificateAuthorityArnList;
        public InputList<string> CertificateAuthorityArnList
        {
            get => _certificateAuthorityArnList ?? (_certificateAuthorityArnList = new InputList<string>());
            set => _certificateAuthorityArnList = value;
        }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ClusterTlsArgs()
        {
        }
    }
}
