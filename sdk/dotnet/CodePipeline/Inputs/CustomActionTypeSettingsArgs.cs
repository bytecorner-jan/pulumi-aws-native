// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Inputs
{

    public sealed class CustomActionTypeSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("entityUrlTemplate")]
        public Input<string>? EntityUrlTemplate { get; set; }

        [Input("executionUrlTemplate")]
        public Input<string>? ExecutionUrlTemplate { get; set; }

        [Input("revisionUrlTemplate")]
        public Input<string>? RevisionUrlTemplate { get; set; }

        [Input("thirdPartyConfigurationUrl")]
        public Input<string>? ThirdPartyConfigurationUrl { get; set; }

        public CustomActionTypeSettingsArgs()
        {
        }
    }
}
