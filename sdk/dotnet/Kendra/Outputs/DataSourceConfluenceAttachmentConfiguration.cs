// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceattachmentconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceConfluenceAttachmentConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceattachmentconfiguration.html#cfn-kendra-datasource-confluenceattachmentconfiguration-attachmentfieldmappings
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceConfluenceAttachmentToIndexFieldMapping> AttachmentFieldMappings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceattachmentconfiguration.html#cfn-kendra-datasource-confluenceattachmentconfiguration-crawlattachments
        /// </summary>
        public readonly bool? CrawlAttachments;

        [OutputConstructor]
        private DataSourceConfluenceAttachmentConfiguration(
            ImmutableArray<Outputs.DataSourceConfluenceAttachmentToIndexFieldMapping> attachmentFieldMappings,

            bool? crawlAttachments)
        {
            AttachmentFieldMappings = attachmentFieldMappings;
            CrawlAttachments = crawlAttachments;
        }
    }
}
