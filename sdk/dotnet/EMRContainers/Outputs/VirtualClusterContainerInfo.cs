// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMRContainers.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-containerinfo.html
    /// </summary>
    [OutputType]
    public sealed class VirtualClusterContainerInfo
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-containerinfo.html#cfn-emrcontainers-virtualcluster-containerinfo-eksinfo
        /// </summary>
        public readonly Outputs.VirtualClusterEksInfo EksInfo;

        [OutputConstructor]
        private VirtualClusterContainerInfo(Outputs.VirtualClusterEksInfo eksInfo)
        {
            EksInfo = eksInfo;
        }
    }
}
