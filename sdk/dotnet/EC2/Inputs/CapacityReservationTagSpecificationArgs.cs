// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class CapacityReservationTagSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        [Input("tags")]
        private InputList<Inputs.CapacityReservationTagArgs>? _tags;
        public InputList<Inputs.CapacityReservationTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CapacityReservationTagArgs>());
            set => _tags = value;
        }

        public CapacityReservationTagSpecificationArgs()
        {
        }
    }
}
