// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ACMPCA.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html
    /// </summary>
    [OutputType]
    public sealed class CertificateSubject
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-commonname
        /// </summary>
        public readonly string? CommonName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-country
        /// </summary>
        public readonly string? Country;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-distinguishednamequalifier
        /// </summary>
        public readonly string? DistinguishedNameQualifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-generationqualifier
        /// </summary>
        public readonly string? GenerationQualifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-givenname
        /// </summary>
        public readonly string? GivenName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-initials
        /// </summary>
        public readonly string? Initials;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-locality
        /// </summary>
        public readonly string? Locality;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-organization
        /// </summary>
        public readonly string? Organization;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-organizationalunit
        /// </summary>
        public readonly string? OrganizationalUnit;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-pseudonym
        /// </summary>
        public readonly string? Pseudonym;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-serialnumber
        /// </summary>
        public readonly string? SerialNumber;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-state
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-surname
        /// </summary>
        public readonly string? Surname;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-subject.html#cfn-acmpca-certificate-subject-title
        /// </summary>
        public readonly string? Title;

        [OutputConstructor]
        private CertificateSubject(
            string? commonName,

            string? country,

            string? distinguishedNameQualifier,

            string? generationQualifier,

            string? givenName,

            string? initials,

            string? locality,

            string? organization,

            string? organizationalUnit,

            string? pseudonym,

            string? serialNumber,

            string? state,

            string? surname,

            string? title)
        {
            CommonName = commonName;
            Country = country;
            DistinguishedNameQualifier = distinguishedNameQualifier;
            GenerationQualifier = generationQualifier;
            GivenName = givenName;
            Initials = initials;
            Locality = locality;
            Organization = organization;
            OrganizationalUnit = organizationalUnit;
            Pseudonym = pseudonym;
            SerialNumber = serialNumber;
            State = state;
            Surname = surname;
            Title = title;
        }
    }
}
