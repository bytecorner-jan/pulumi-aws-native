// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    public sealed class TriggerConditionArgs : Pulumi.ResourceArgs
    {
        [Input("crawlState")]
        public Input<string>? CrawlState { get; set; }

        [Input("crawlerName")]
        public Input<string>? CrawlerName { get; set; }

        [Input("jobName")]
        public Input<string>? JobName { get; set; }

        [Input("logicalOperator")]
        public Input<string>? LogicalOperator { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        public TriggerConditionArgs()
        {
        }
    }
}
