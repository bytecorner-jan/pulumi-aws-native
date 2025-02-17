// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Inputs
{

    /// <summary>
    /// Contains a time interval window used for data aggregate computations (for example, average, sum, count, and so on).
    /// </summary>
    public sealed class AssetModelMetricWindowArgs : Pulumi.ResourceArgs
    {
        [Input("tumbling")]
        public Input<Inputs.AssetModelTumblingWindowArgs>? Tumbling { get; set; }

        public AssetModelMetricWindowArgs()
        {
        }
    }
}
