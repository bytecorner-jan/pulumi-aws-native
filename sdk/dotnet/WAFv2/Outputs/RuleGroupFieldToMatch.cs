// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Field of the request to match.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupFieldToMatch
    {
        /// <summary>
        /// All query arguments of a web request.
        /// </summary>
        public readonly object? AllQueryArguments;
        /// <summary>
        /// The body of a web request. This immediately follows the request headers.
        /// </summary>
        public readonly object? Body;
        public readonly Outputs.RuleGroupJsonBody? JsonBody;
        /// <summary>
        /// The HTTP method of a web request. The method indicates the type of operation that the request is asking the origin to perform.
        /// </summary>
        public readonly object? Method;
        /// <summary>
        /// The query string of a web request. This is the part of a URL that appears after a ? character, if any.
        /// </summary>
        public readonly object? QueryString;
        public readonly Outputs.RuleGroupFieldToMatchSingleHeaderProperties? SingleHeader;
        /// <summary>
        /// One query argument in a web request, identified by name, for example UserName or SalesRegion. The name can be up to 30 characters long and isn't case sensitive.
        /// </summary>
        public readonly Outputs.RuleGroupFieldToMatchSingleQueryArgumentProperties? SingleQueryArgument;
        /// <summary>
        /// The path component of the URI of a web request. This is the part of a web request that identifies a resource, for example, /images/daily-ad.jpg.
        /// </summary>
        public readonly object? UriPath;

        [OutputConstructor]
        private RuleGroupFieldToMatch(
            object? allQueryArguments,

            object? body,

            Outputs.RuleGroupJsonBody? jsonBody,

            object? method,

            object? queryString,

            Outputs.RuleGroupFieldToMatchSingleHeaderProperties? singleHeader,

            Outputs.RuleGroupFieldToMatchSingleQueryArgumentProperties? singleQueryArgument,

            object? uriPath)
        {
            AllQueryArguments = allQueryArguments;
            Body = body;
            JsonBody = jsonBody;
            Method = method;
            QueryString = queryString;
            SingleHeader = singleHeader;
            SingleQueryArgument = singleQueryArgument;
            UriPath = uriPath;
        }
    }
}
