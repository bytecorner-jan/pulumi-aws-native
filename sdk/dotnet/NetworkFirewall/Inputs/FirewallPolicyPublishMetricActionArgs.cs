// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-publishmetricaction.html
    /// </summary>
    public sealed class FirewallPolicyPublishMetricActionArgs : Pulumi.ResourceArgs
    {
        [Input("dimensions", required: true)]
        private InputList<Inputs.FirewallPolicyDimensionArgs>? _dimensions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-publishmetricaction.html#cfn-networkfirewall-firewallpolicy-publishmetricaction-dimensions
        /// </summary>
        public InputList<Inputs.FirewallPolicyDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.FirewallPolicyDimensionArgs>());
            set => _dimensions = value;
        }

        public FirewallPolicyPublishMetricActionArgs()
        {
        }
    }
}
