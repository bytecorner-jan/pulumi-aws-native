// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Outputs
{

    [OutputType]
    public sealed class PipelineInputArtifact
    {
        public readonly string Name;

        [OutputConstructor]
        private PipelineInputArtifact(string name)
        {
            Name = name;
        }
    }
}
