// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AccessAnalyzer.Inputs
{

    public sealed class AnalyzerFilterArgs : Pulumi.ResourceArgs
    {
        [Input("contains")]
        private InputList<string>? _contains;
        public InputList<string> Contains
        {
            get => _contains ?? (_contains = new InputList<string>());
            set => _contains = value;
        }

        [Input("eq")]
        private InputList<string>? _eq;
        public InputList<string> Eq
        {
            get => _eq ?? (_eq = new InputList<string>());
            set => _eq = value;
        }

        [Input("exists")]
        public Input<bool>? Exists { get; set; }

        [Input("neq")]
        private InputList<string>? _neq;
        public InputList<string> Neq
        {
            get => _neq ?? (_neq = new InputList<string>());
            set => _neq = value;
        }

        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        public AnalyzerFilterArgs()
        {
        }
    }
}
