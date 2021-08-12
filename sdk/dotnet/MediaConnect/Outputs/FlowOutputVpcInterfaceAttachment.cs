// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-vpcinterfaceattachment.html
    /// </summary>
    [OutputType]
    public sealed class FlowOutputVpcInterfaceAttachment
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-vpcinterfaceattachment.html#cfn-mediaconnect-flowoutput-vpcinterfaceattachment-vpcinterfacename
        /// </summary>
        public readonly string? VpcInterfaceName;

        [OutputConstructor]
        private FlowOutputVpcInterfaceAttachment(string? vpcInterfaceName)
        {
            VpcInterfaceName = vpcInterfaceName;
        }
    }
}
