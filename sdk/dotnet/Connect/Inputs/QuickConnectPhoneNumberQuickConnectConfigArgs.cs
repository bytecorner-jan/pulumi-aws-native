// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// The phone configuration. This is required only if QuickConnectType is PHONE_NUMBER.
    /// </summary>
    public sealed class QuickConnectPhoneNumberQuickConnectConfigArgs : Pulumi.ResourceArgs
    {
        [Input("phoneNumber", required: true)]
        public Input<string> PhoneNumber { get; set; } = null!;

        public QuickConnectPhoneNumberQuickConnectConfigArgs()
        {
        }
    }
}
