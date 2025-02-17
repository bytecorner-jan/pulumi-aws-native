// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ACMPCA.Inputs
{

    /// <summary>
    /// Structure that contains X.509 OtherName information.
    /// </summary>
    public sealed class CertificateAuthorityOtherNameArgs : Pulumi.ResourceArgs
    {
        [Input("typeId", required: true)]
        public Input<string> TypeId { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public CertificateAuthorityOtherNameArgs()
        {
        }
    }
}
