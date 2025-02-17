// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class UserPoolLambdaConfigArgs : Pulumi.ResourceArgs
    {
        [Input("createAuthChallenge")]
        public Input<string>? CreateAuthChallenge { get; set; }

        [Input("customEmailSender")]
        public Input<Inputs.UserPoolCustomEmailSenderArgs>? CustomEmailSender { get; set; }

        [Input("customMessage")]
        public Input<string>? CustomMessage { get; set; }

        [Input("customSMSSender")]
        public Input<Inputs.UserPoolCustomSMSSenderArgs>? CustomSMSSender { get; set; }

        [Input("defineAuthChallenge")]
        public Input<string>? DefineAuthChallenge { get; set; }

        [Input("kMSKeyID")]
        public Input<string>? KMSKeyID { get; set; }

        [Input("postAuthentication")]
        public Input<string>? PostAuthentication { get; set; }

        [Input("postConfirmation")]
        public Input<string>? PostConfirmation { get; set; }

        [Input("preAuthentication")]
        public Input<string>? PreAuthentication { get; set; }

        [Input("preSignUp")]
        public Input<string>? PreSignUp { get; set; }

        [Input("preTokenGeneration")]
        public Input<string>? PreTokenGeneration { get; set; }

        [Input("userMigration")]
        public Input<string>? UserMigration { get; set; }

        [Input("verifyAuthChallengeResponse")]
        public Input<string>? VerifyAuthChallengeResponse { get; set; }

        public UserPoolLambdaConfigArgs()
        {
        }
    }
}
