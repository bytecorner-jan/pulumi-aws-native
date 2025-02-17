// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    [OutputType]
    public sealed class VirtualGatewayListenerTlsCertificate
    {
        public readonly Outputs.VirtualGatewayListenerTlsAcmCertificate? ACM;
        public readonly Outputs.VirtualGatewayListenerTlsFileCertificate? File;
        public readonly Outputs.VirtualGatewayListenerTlsSdsCertificate? SDS;

        [OutputConstructor]
        private VirtualGatewayListenerTlsCertificate(
            Outputs.VirtualGatewayListenerTlsAcmCertificate? aCM,

            Outputs.VirtualGatewayListenerTlsFileCertificate? file,

            Outputs.VirtualGatewayListenerTlsSdsCertificate? sDS)
        {
            ACM = aCM;
            File = file;
            SDS = sDS;
        }
    }
}
