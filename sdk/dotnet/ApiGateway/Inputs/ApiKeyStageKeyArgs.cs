// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Inputs
{

    public sealed class ApiKeyStageKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of a RestApi resource that includes the stage with which you want to associate the API key.
        /// </summary>
        [Input("restApiId")]
        public Input<string>? RestApiId { get; set; }

        /// <summary>
        /// The name of the stage with which to associate the API key. The stage must be included in the RestApi resource that you specified in the RestApiId property. 
        /// </summary>
        [Input("stageName")]
        public Input<string>? StageName { get; set; }

        public ApiKeyStageKeyArgs()
        {
        }
    }
}
