// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class ModelContainerDefinition
    {
        public readonly string? ContainerHostname;
        public readonly object? Environment;
        public readonly string? Image;
        public readonly Outputs.ModelImageConfig? ImageConfig;
        public readonly string? Mode;
        public readonly string? ModelDataUrl;
        public readonly string? ModelPackageName;
        public readonly Outputs.ModelMultiModelConfig? MultiModelConfig;

        [OutputConstructor]
        private ModelContainerDefinition(
            string? containerHostname,

            object? environment,

            string? image,

            Outputs.ModelImageConfig? imageConfig,

            string? mode,

            string? modelDataUrl,

            string? modelPackageName,

            Outputs.ModelMultiModelConfig? multiModelConfig)
        {
            ContainerHostname = containerHostname;
            Environment = environment;
            Image = image;
            ImageConfig = imageConfig;
            Mode = mode;
            ModelDataUrl = modelDataUrl;
            ModelPackageName = modelPackageName;
            MultiModelConfig = multiModelConfig;
        }
    }
}
