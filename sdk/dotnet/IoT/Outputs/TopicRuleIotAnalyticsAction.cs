// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotanalyticsaction.html
    /// </summary>
    [OutputType]
    public sealed class TopicRuleIotAnalyticsAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotanalyticsaction.html#cfn-iot-topicrule-iotanalyticsaction-batchmode
        /// </summary>
        public readonly bool? BatchMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotanalyticsaction.html#cfn-iot-topicrule-iotanalyticsaction-channelname
        /// </summary>
        public readonly string ChannelName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotanalyticsaction.html#cfn-iot-topicrule-iotanalyticsaction-rolearn
        /// </summary>
        public readonly string RoleArn;

        [OutputConstructor]
        private TopicRuleIotAnalyticsAction(
            bool? batchMode,

            string channelName,

            string roleArn)
        {
            BatchMode = batchMode;
            ChannelName = channelName;
            RoleArn = roleArn;
        }
    }
}
