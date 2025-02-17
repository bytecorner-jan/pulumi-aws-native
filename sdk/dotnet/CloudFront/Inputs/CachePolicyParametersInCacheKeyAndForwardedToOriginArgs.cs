// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class CachePolicyParametersInCacheKeyAndForwardedToOriginArgs : Pulumi.ResourceArgs
    {
        [Input("cookiesConfig", required: true)]
        public Input<Inputs.CachePolicyCookiesConfigArgs> CookiesConfig { get; set; } = null!;

        [Input("enableAcceptEncodingBrotli")]
        public Input<bool>? EnableAcceptEncodingBrotli { get; set; }

        [Input("enableAcceptEncodingGzip", required: true)]
        public Input<bool> EnableAcceptEncodingGzip { get; set; } = null!;

        [Input("headersConfig", required: true)]
        public Input<Inputs.CachePolicyHeadersConfigArgs> HeadersConfig { get; set; } = null!;

        [Input("queryStringsConfig", required: true)]
        public Input<Inputs.CachePolicyQueryStringsConfigArgs> QueryStringsConfig { get; set; } = null!;

        public CachePolicyParametersInCacheKeyAndForwardedToOriginArgs()
        {
        }
    }
}
