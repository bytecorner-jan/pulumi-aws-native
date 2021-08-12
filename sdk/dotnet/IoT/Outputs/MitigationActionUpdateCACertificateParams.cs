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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-updatecacertificateparams.html
    /// </summary>
    [OutputType]
    public sealed class MitigationActionUpdateCACertificateParams
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-updatecacertificateparams.html#cfn-iot-mitigationaction-updatecacertificateparams-action
        /// </summary>
        public readonly string Action;

        [OutputConstructor]
        private MitigationActionUpdateCACertificateParams(string action)
        {
            Action = action;
        }
    }
}
