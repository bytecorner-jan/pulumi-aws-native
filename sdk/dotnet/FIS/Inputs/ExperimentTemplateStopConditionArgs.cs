// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FIS.Inputs
{

    public sealed class ExperimentTemplateStopConditionArgs : Pulumi.ResourceArgs
    {
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ExperimentTemplateStopConditionArgs()
        {
        }
    }
}
