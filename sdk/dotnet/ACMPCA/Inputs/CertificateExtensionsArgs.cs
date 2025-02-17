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
    /// Structure that contains X.500 extensions for a Certificate.
    /// </summary>
    public sealed class CertificateExtensionsArgs : Pulumi.ResourceArgs
    {
        [Input("certificatePolicies")]
        private InputList<Inputs.CertificatePolicyInformationArgs>? _certificatePolicies;
        public InputList<Inputs.CertificatePolicyInformationArgs> CertificatePolicies
        {
            get => _certificatePolicies ?? (_certificatePolicies = new InputList<Inputs.CertificatePolicyInformationArgs>());
            set => _certificatePolicies = value;
        }

        [Input("extendedKeyUsage")]
        private InputList<Inputs.CertificateExtendedKeyUsageArgs>? _extendedKeyUsage;
        public InputList<Inputs.CertificateExtendedKeyUsageArgs> ExtendedKeyUsage
        {
            get => _extendedKeyUsage ?? (_extendedKeyUsage = new InputList<Inputs.CertificateExtendedKeyUsageArgs>());
            set => _extendedKeyUsage = value;
        }

        [Input("keyUsage")]
        public Input<Inputs.CertificateKeyUsageArgs>? KeyUsage { get; set; }

        [Input("subjectAlternativeNames")]
        private InputList<Inputs.CertificateGeneralNameArgs>? _subjectAlternativeNames;
        public InputList<Inputs.CertificateGeneralNameArgs> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new InputList<Inputs.CertificateGeneralNameArgs>());
            set => _subjectAlternativeNames = value;
        }

        public CertificateExtensionsArgs()
        {
        }
    }
}
