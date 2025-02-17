// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class EndpointConfigProductionVariantArgs : Pulumi.ResourceArgs
    {
        [Input("acceleratorType")]
        public Input<string>? AcceleratorType { get; set; }

        [Input("initialInstanceCount", required: true)]
        public Input<int> InitialInstanceCount { get; set; } = null!;

        [Input("initialVariantWeight", required: true)]
        public Input<double> InitialVariantWeight { get; set; } = null!;

        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("modelName", required: true)]
        public Input<string> ModelName { get; set; } = null!;

        [Input("variantName", required: true)]
        public Input<string> VariantName { get; set; } = null!;

        public EndpointConfigProductionVariantArgs()
        {
        }
    }
}
