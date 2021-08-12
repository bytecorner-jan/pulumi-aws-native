// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FIS.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html
    /// </summary>
    public sealed class ExperimentTemplateExperimentTemplateTargetArgs : Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.ExperimentTemplateExperimentTemplateTargetFilterArgs>? _filters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-filters
        /// </summary>
        public InputList<Inputs.ExperimentTemplateExperimentTemplateTargetFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.ExperimentTemplateExperimentTemplateTargetFilterArgs>());
            set => _filters = value;
        }

        [Input("resourceArns")]
        private InputList<string>? _resourceArns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-resourcearns
        /// </summary>
        public InputList<string> ResourceArns
        {
            get => _resourceArns ?? (_resourceArns = new InputList<string>());
            set => _resourceArns = value;
        }

        [Input("resourceTags")]
        private InputMap<string>? _resourceTags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-resourcetags
        /// </summary>
        public InputMap<string> ResourceTags
        {
            get => _resourceTags ?? (_resourceTags = new InputMap<string>());
            set => _resourceTags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-resourcetype
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-selectionmode
        /// </summary>
        [Input("selectionMode", required: true)]
        public Input<string> SelectionMode { get; set; } = null!;

        public ExperimentTemplateExperimentTemplateTargetArgs()
        {
        }
    }
}
