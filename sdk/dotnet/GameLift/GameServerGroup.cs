// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html
    /// </summary>
    [AwsNativeResourceType("aws-native:GameLift:GameServerGroup")]
    public partial class GameServerGroup : Pulumi.CustomResource
    {
        [Output("autoScalingGroupArn")]
        public Output<string> AutoScalingGroupArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-autoscalingpolicy
        /// </summary>
        [Output("autoScalingPolicy")]
        public Output<Outputs.GameServerGroupAutoScalingPolicy?> AutoScalingPolicy { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-balancingstrategy
        /// </summary>
        [Output("balancingStrategy")]
        public Output<string?> BalancingStrategy { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-deleteoption
        /// </summary>
        [Output("deleteOption")]
        public Output<string?> DeleteOption { get; private set; } = null!;

        [Output("gameServerGroupArn")]
        public Output<string> GameServerGroupArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameservergroupname
        /// </summary>
        [Output("gameServerGroupName")]
        public Output<string> GameServerGroupName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameserverprotectionpolicy
        /// </summary>
        [Output("gameServerProtectionPolicy")]
        public Output<string?> GameServerProtectionPolicy { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-instancedefinitions
        /// </summary>
        [Output("instanceDefinitions")]
        public Output<ImmutableArray<Outputs.GameServerGroupInstanceDefinition>> InstanceDefinitions { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-launchtemplate
        /// </summary>
        [Output("launchTemplate")]
        public Output<Outputs.GameServerGroupLaunchTemplate> LaunchTemplate { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-maxsize
        /// </summary>
        [Output("maxSize")]
        public Output<double?> MaxSize { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-minsize
        /// </summary>
        [Output("minSize")]
        public Output<double?> MinSize { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-rolearn
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-vpcsubnets
        /// </summary>
        [Output("vpcSubnets")]
        public Output<ImmutableArray<string>> VpcSubnets { get; private set; } = null!;


        /// <summary>
        /// Create a GameServerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GameServerGroup(string name, GameServerGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:GameLift:GameServerGroup", name, args ?? new GameServerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GameServerGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:GameLift:GameServerGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GameServerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GameServerGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GameServerGroup(name, id, options);
        }
    }

    public sealed class GameServerGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-autoscalingpolicy
        /// </summary>
        [Input("autoScalingPolicy")]
        public Input<Inputs.GameServerGroupAutoScalingPolicyArgs>? AutoScalingPolicy { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-balancingstrategy
        /// </summary>
        [Input("balancingStrategy")]
        public Input<string>? BalancingStrategy { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-deleteoption
        /// </summary>
        [Input("deleteOption")]
        public Input<string>? DeleteOption { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameservergroupname
        /// </summary>
        [Input("gameServerGroupName", required: true)]
        public Input<string> GameServerGroupName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameserverprotectionpolicy
        /// </summary>
        [Input("gameServerProtectionPolicy")]
        public Input<string>? GameServerProtectionPolicy { get; set; }

        [Input("instanceDefinitions", required: true)]
        private InputList<Inputs.GameServerGroupInstanceDefinitionArgs>? _instanceDefinitions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-instancedefinitions
        /// </summary>
        public InputList<Inputs.GameServerGroupInstanceDefinitionArgs> InstanceDefinitions
        {
            get => _instanceDefinitions ?? (_instanceDefinitions = new InputList<Inputs.GameServerGroupInstanceDefinitionArgs>());
            set => _instanceDefinitions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-launchtemplate
        /// </summary>
        [Input("launchTemplate", required: true)]
        public Input<Inputs.GameServerGroupLaunchTemplateArgs> LaunchTemplate { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-maxsize
        /// </summary>
        [Input("maxSize")]
        public Input<double>? MaxSize { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-minsize
        /// </summary>
        [Input("minSize")]
        public Input<double>? MinSize { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-rolearn
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("vpcSubnets")]
        private InputList<string>? _vpcSubnets;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-vpcsubnets
        /// </summary>
        public InputList<string> VpcSubnets
        {
            get => _vpcSubnets ?? (_vpcSubnets = new InputList<string>());
            set => _vpcSubnets = value;
        }

        public GameServerGroupArgs()
        {
        }
    }
}
