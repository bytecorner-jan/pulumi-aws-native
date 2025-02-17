// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker
{
    /// <summary>
    /// AWS::RoboMaker::Robot resource creates an AWS RoboMaker fleet.
    /// </summary>
    [AwsNativeResourceType("aws-native:robomaker:Robot")]
    public partial class Robot : Pulumi.CustomResource
    {
        /// <summary>
        /// The target architecture of the robot.
        /// </summary>
        [Output("architecture")]
        public Output<Pulumi.AwsNative.RoboMaker.RobotArchitecture> Architecture { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the fleet.
        /// </summary>
        [Output("fleet")]
        public Output<string?> Fleet { get; private set; } = null!;

        /// <summary>
        /// The Greengrass group id.
        /// </summary>
        [Output("greengrassGroupId")]
        public Output<string> GreengrassGroupId { get; private set; } = null!;

        /// <summary>
        /// The name for the robot.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("tags")]
        public Output<Outputs.RobotTags?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Robot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Robot(string name, RobotArgs args, CustomResourceOptions? options = null)
            : base("aws-native:robomaker:Robot", name, args ?? new RobotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Robot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:robomaker:Robot", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Robot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Robot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Robot(name, id, options);
        }
    }

    public sealed class RobotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target architecture of the robot.
        /// </summary>
        [Input("architecture", required: true)]
        public Input<Pulumi.AwsNative.RoboMaker.RobotArchitecture> Architecture { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the fleet.
        /// </summary>
        [Input("fleet")]
        public Input<string>? Fleet { get; set; }

        /// <summary>
        /// The Greengrass group id.
        /// </summary>
        [Input("greengrassGroupId", required: true)]
        public Input<string> GreengrassGroupId { get; set; } = null!;

        /// <summary>
        /// The name for the robot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        public Input<Inputs.RobotTagsArgs>? Tags { get; set; }

        public RobotArgs()
        {
        }
    }
}
