// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sns.html
    /// </summary>
    [OutputType]
    public sealed class DetectorModelSns
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sns.html#cfn-iotevents-detectormodel-sns-payload
        /// </summary>
        public readonly Outputs.DetectorModelPayload? Payload;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sns.html#cfn-iotevents-detectormodel-sns-targetarn
        /// </summary>
        public readonly string TargetArn;

        [OutputConstructor]
        private DetectorModelSns(
            Outputs.DetectorModelPayload? payload,

            string targetArn)
        {
            Payload = payload;
            TargetArn = targetArn;
        }
    }
}
