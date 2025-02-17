// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CE.Outputs
{

    [OutputType]
    public sealed class AnomalySubscriptionSubscriber
    {
        public readonly string Address;
        public readonly Pulumi.AwsNative.CE.AnomalySubscriptionSubscriberStatus? Status;
        public readonly Pulumi.AwsNative.CE.AnomalySubscriptionSubscriberType Type;

        [OutputConstructor]
        private AnomalySubscriptionSubscriber(
            string address,

            Pulumi.AwsNative.CE.AnomalySubscriptionSubscriberStatus? status,

            Pulumi.AwsNative.CE.AnomalySubscriptionSubscriberType type)
        {
            Address = address;
            Status = status;
            Type = type;
        }
    }
}
