// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// The criteria that determine when and how a job abort takes place.
    /// </summary>
    public sealed class AbortConfigPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("criteriaList", required: true)]
        private InputList<Inputs.JobTemplateAbortCriteriaArgs>? _criteriaList;
        public InputList<Inputs.JobTemplateAbortCriteriaArgs> CriteriaList
        {
            get => _criteriaList ?? (_criteriaList = new InputList<Inputs.JobTemplateAbortCriteriaArgs>());
            set => _criteriaList = value;
        }

        public AbortConfigPropertiesArgs()
        {
        }
    }
}
