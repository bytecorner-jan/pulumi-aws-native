// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MWAA
{
    /// <summary>
    /// Resource schema for AWS::MWAA::Environment
    /// </summary>
    [AwsNativeResourceType("aws-native:mwaa:Environment")]
    public partial class Environment : Pulumi.CustomResource
    {
        /// <summary>
        /// Key/value pairs representing Airflow configuration variables.
        ///     Keys are prefixed by their section:
        /// 
        ///     [core]
        ///     dags_folder={AIRFLOW_HOME}/dags
        /// 
        ///     Would be represented as
        /// 
        ///     "core.dags_folder": "{AIRFLOW_HOME}/dags"
        /// </summary>
        [Output("airflowConfigurationOptions")]
        public Output<object?> AirflowConfigurationOptions { get; private set; } = null!;

        [Output("airflowVersion")]
        public Output<string?> AirflowVersion { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("dagS3Path")]
        public Output<string?> DagS3Path { get; private set; } = null!;

        [Output("environmentClass")]
        public Output<string?> EnvironmentClass { get; private set; } = null!;

        [Output("executionRoleArn")]
        public Output<string?> ExecutionRoleArn { get; private set; } = null!;

        [Output("kmsKey")]
        public Output<string?> KmsKey { get; private set; } = null!;

        [Output("loggingConfiguration")]
        public Output<Outputs.EnvironmentLoggingConfiguration?> LoggingConfiguration { get; private set; } = null!;

        [Output("maxWorkers")]
        public Output<int?> MaxWorkers { get; private set; } = null!;

        [Output("minWorkers")]
        public Output<int?> MinWorkers { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("networkConfiguration")]
        public Output<Outputs.EnvironmentNetworkConfiguration?> NetworkConfiguration { get; private set; } = null!;

        [Output("pluginsS3ObjectVersion")]
        public Output<string?> PluginsS3ObjectVersion { get; private set; } = null!;

        [Output("pluginsS3Path")]
        public Output<string?> PluginsS3Path { get; private set; } = null!;

        [Output("requirementsS3ObjectVersion")]
        public Output<string?> RequirementsS3ObjectVersion { get; private set; } = null!;

        [Output("requirementsS3Path")]
        public Output<string?> RequirementsS3Path { get; private set; } = null!;

        [Output("schedulers")]
        public Output<int?> Schedulers { get; private set; } = null!;

        [Output("sourceBucketArn")]
        public Output<string?> SourceBucketArn { get; private set; } = null!;

        /// <summary>
        /// A map of tags for the environment.
        /// </summary>
        [Output("tags")]
        public Output<object?> Tags { get; private set; } = null!;

        [Output("webserverAccessMode")]
        public Output<Pulumi.AwsNative.MWAA.EnvironmentWebserverAccessMode?> WebserverAccessMode { get; private set; } = null!;

        [Output("webserverUrl")]
        public Output<string> WebserverUrl { get; private set; } = null!;

        [Output("weeklyMaintenanceWindowStart")]
        public Output<string?> WeeklyMaintenanceWindowStart { get; private set; } = null!;


        /// <summary>
        /// Create a Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Environment(string name, EnvironmentArgs args, CustomResourceOptions? options = null)
            : base("aws-native:mwaa:Environment", name, args ?? new EnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Environment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:mwaa:Environment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Environment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Environment(name, id, options);
        }
    }

    public sealed class EnvironmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key/value pairs representing Airflow configuration variables.
        ///     Keys are prefixed by their section:
        /// 
        ///     [core]
        ///     dags_folder={AIRFLOW_HOME}/dags
        /// 
        ///     Would be represented as
        /// 
        ///     "core.dags_folder": "{AIRFLOW_HOME}/dags"
        /// </summary>
        [Input("airflowConfigurationOptions")]
        public Input<object>? AirflowConfigurationOptions { get; set; }

        [Input("airflowVersion")]
        public Input<string>? AirflowVersion { get; set; }

        [Input("dagS3Path")]
        public Input<string>? DagS3Path { get; set; }

        [Input("environmentClass")]
        public Input<string>? EnvironmentClass { get; set; }

        [Input("executionRoleArn")]
        public Input<string>? ExecutionRoleArn { get; set; }

        [Input("kmsKey")]
        public Input<string>? KmsKey { get; set; }

        [Input("loggingConfiguration")]
        public Input<Inputs.EnvironmentLoggingConfigurationArgs>? LoggingConfiguration { get; set; }

        [Input("maxWorkers")]
        public Input<int>? MaxWorkers { get; set; }

        [Input("minWorkers")]
        public Input<int>? MinWorkers { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("networkConfiguration")]
        public Input<Inputs.EnvironmentNetworkConfigurationArgs>? NetworkConfiguration { get; set; }

        [Input("pluginsS3ObjectVersion")]
        public Input<string>? PluginsS3ObjectVersion { get; set; }

        [Input("pluginsS3Path")]
        public Input<string>? PluginsS3Path { get; set; }

        [Input("requirementsS3ObjectVersion")]
        public Input<string>? RequirementsS3ObjectVersion { get; set; }

        [Input("requirementsS3Path")]
        public Input<string>? RequirementsS3Path { get; set; }

        [Input("schedulers")]
        public Input<int>? Schedulers { get; set; }

        [Input("sourceBucketArn")]
        public Input<string>? SourceBucketArn { get; set; }

        /// <summary>
        /// A map of tags for the environment.
        /// </summary>
        [Input("tags")]
        public Input<object>? Tags { get; set; }

        [Input("webserverAccessMode")]
        public Input<Pulumi.AwsNative.MWAA.EnvironmentWebserverAccessMode>? WebserverAccessMode { get; set; }

        [Input("weeklyMaintenanceWindowStart")]
        public Input<string>? WeeklyMaintenanceWindowStart { get; set; }

        public EnvironmentArgs()
        {
        }
    }
}
