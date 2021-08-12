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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-lambda.html
    /// </summary>
    [OutputType]
    public sealed class DetectorModelLambda
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-lambda.html#cfn-iotevents-detectormodel-lambda-functionarn
        /// </summary>
        public readonly string FunctionArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-lambda.html#cfn-iotevents-detectormodel-lambda-payload
        /// </summary>
        public readonly Outputs.DetectorModelPayload? Payload;

        [OutputConstructor]
        private DetectorModelLambda(
            string functionArn,

            Outputs.DetectorModelPayload? payload)
        {
            FunctionArn = functionArn;
            Payload = payload;
        }
    }
}
