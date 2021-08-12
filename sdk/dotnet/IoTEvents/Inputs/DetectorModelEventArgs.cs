// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-event.html
    /// </summary>
    public sealed class DetectorModelEventArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.DetectorModelActionArgs>? _actions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-event.html#cfn-iotevents-detectormodel-event-actions
        /// </summary>
        public InputList<Inputs.DetectorModelActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.DetectorModelActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-event.html#cfn-iotevents-detectormodel-event-condition
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-event.html#cfn-iotevents-detectormodel-event-eventname
        /// </summary>
        [Input("eventName", required: true)]
        public Input<string> EventName { get; set; } = null!;

        public DetectorModelEventArgs()
        {
        }
    }
}
