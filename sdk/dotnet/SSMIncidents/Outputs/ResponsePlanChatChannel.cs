// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSMIncidents.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-chatchannel.html
    /// </summary>
    [OutputType]
    public sealed class ResponsePlanChatChannel
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-chatchannel.html#cfn-ssmincidents-responseplan-chatchannel-chatbotsns
        /// </summary>
        public readonly ImmutableArray<string> ChatbotSns;

        [OutputConstructor]
        private ResponsePlanChatChannel(ImmutableArray<string> chatbotSns)
        {
            ChatbotSns = chatbotSns;
        }
    }
}
