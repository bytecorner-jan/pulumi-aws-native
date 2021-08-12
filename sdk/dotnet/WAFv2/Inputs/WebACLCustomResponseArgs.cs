// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponse.html
    /// </summary>
    public sealed class WebACLCustomResponseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponse.html#cfn-wafv2-webacl-customresponse-customresponsebodykey
        /// </summary>
        [Input("customResponseBodyKey")]
        public Input<string>? CustomResponseBodyKey { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponse.html#cfn-wafv2-webacl-customresponse-responsecode
        /// </summary>
        [Input("responseCode", required: true)]
        public Input<int> ResponseCode { get; set; } = null!;

        [Input("responseHeaders")]
        private InputList<Inputs.WebACLCustomHTTPHeaderArgs>? _responseHeaders;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponse.html#cfn-wafv2-webacl-customresponse-responseheaders
        /// </summary>
        public InputList<Inputs.WebACLCustomHTTPHeaderArgs> ResponseHeaders
        {
            get => _responseHeaders ?? (_responseHeaders = new InputList<Inputs.WebACLCustomHTTPHeaderArgs>());
            set => _responseHeaders = value;
        }

        public WebACLCustomResponseArgs()
        {
        }
    }
}
