// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSMContacts.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-targets.html
    /// </summary>
    [OutputType]
    public sealed class ContactTargets
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-targets.html#cfn-ssmcontacts-contact-targets-channeltargetinfo
        /// </summary>
        public readonly Outputs.ContactChannelTargetInfo? ChannelTargetInfo;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-targets.html#cfn-ssmcontacts-contact-targets-contacttargetinfo
        /// </summary>
        public readonly Outputs.ContactContactTargetInfo? ContactTargetInfo;

        [OutputConstructor]
        private ContactTargets(
            Outputs.ContactChannelTargetInfo? channelTargetInfo,

            Outputs.ContactContactTargetInfo? contactTargetInfo)
        {
            ChannelTargetInfo = channelTargetInfo;
            ContactTargetInfo = contactTargetInfo;
        }
    }
}
