// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT1Click.Inputs
{

    public sealed class ProjectPlacementTemplateArgs : Pulumi.ResourceArgs
    {
        [Input("defaultAttributes")]
        public Input<object>? DefaultAttributes { get; set; }

        [Input("deviceTemplates")]
        public Input<object>? DeviceTemplates { get; set; }

        public ProjectPlacementTemplateArgs()
        {
        }
    }
}
