// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    /// <summary>
    /// Resource Type definition for AWS::GameLift::GameSessionQueue
    /// </summary>
    [Obsolete(@"GameSessionQueue is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:gamelift:GameSessionQueue")]
    public partial class GameSessionQueue : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("customEventData")]
        public Output<string?> CustomEventData { get; private set; } = null!;

        [Output("destinations")]
        public Output<ImmutableArray<Outputs.GameSessionQueueDestination>> Destinations { get; private set; } = null!;

        [Output("filterConfiguration")]
        public Output<Outputs.GameSessionQueueFilterConfiguration?> FilterConfiguration { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("notificationTarget")]
        public Output<string?> NotificationTarget { get; private set; } = null!;

        [Output("playerLatencyPolicies")]
        public Output<ImmutableArray<Outputs.GameSessionQueuePlayerLatencyPolicy>> PlayerLatencyPolicies { get; private set; } = null!;

        [Output("priorityConfiguration")]
        public Output<Outputs.GameSessionQueuePriorityConfiguration?> PriorityConfiguration { get; private set; } = null!;

        [Output("timeoutInSeconds")]
        public Output<int?> TimeoutInSeconds { get; private set; } = null!;


        /// <summary>
        /// Create a GameSessionQueue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GameSessionQueue(string name, GameSessionQueueArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:gamelift:GameSessionQueue", name, args ?? new GameSessionQueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GameSessionQueue(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:gamelift:GameSessionQueue", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GameSessionQueue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GameSessionQueue Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GameSessionQueue(name, id, options);
        }
    }

    public sealed class GameSessionQueueArgs : Pulumi.ResourceArgs
    {
        [Input("customEventData")]
        public Input<string>? CustomEventData { get; set; }

        [Input("destinations")]
        private InputList<Inputs.GameSessionQueueDestinationArgs>? _destinations;
        public InputList<Inputs.GameSessionQueueDestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.GameSessionQueueDestinationArgs>());
            set => _destinations = value;
        }

        [Input("filterConfiguration")]
        public Input<Inputs.GameSessionQueueFilterConfigurationArgs>? FilterConfiguration { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationTarget")]
        public Input<string>? NotificationTarget { get; set; }

        [Input("playerLatencyPolicies")]
        private InputList<Inputs.GameSessionQueuePlayerLatencyPolicyArgs>? _playerLatencyPolicies;
        public InputList<Inputs.GameSessionQueuePlayerLatencyPolicyArgs> PlayerLatencyPolicies
        {
            get => _playerLatencyPolicies ?? (_playerLatencyPolicies = new InputList<Inputs.GameSessionQueuePlayerLatencyPolicyArgs>());
            set => _playerLatencyPolicies = value;
        }

        [Input("priorityConfiguration")]
        public Input<Inputs.GameSessionQueuePriorityConfigurationArgs>? PriorityConfiguration { get; set; }

        [Input("timeoutInSeconds")]
        public Input<int>? TimeoutInSeconds { get; set; }

        public GameSessionQueueArgs()
        {
        }
    }
}
