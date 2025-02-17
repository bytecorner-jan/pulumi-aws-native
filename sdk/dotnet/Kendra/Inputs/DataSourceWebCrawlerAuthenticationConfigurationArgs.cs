// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceWebCrawlerAuthenticationConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("basicAuthentication")]
        private InputList<Inputs.DataSourceWebCrawlerBasicAuthenticationArgs>? _basicAuthentication;
        public InputList<Inputs.DataSourceWebCrawlerBasicAuthenticationArgs> BasicAuthentication
        {
            get => _basicAuthentication ?? (_basicAuthentication = new InputList<Inputs.DataSourceWebCrawlerBasicAuthenticationArgs>());
            set => _basicAuthentication = value;
        }

        public DataSourceWebCrawlerAuthenticationConfigurationArgs()
        {
        }
    }
}
