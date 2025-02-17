// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecretsManager.Outputs
{

    [OutputType]
    public sealed class SecretGenerateSecretString
    {
        public readonly string? ExcludeCharacters;
        public readonly bool? ExcludeLowercase;
        public readonly bool? ExcludeNumbers;
        public readonly bool? ExcludePunctuation;
        public readonly bool? ExcludeUppercase;
        public readonly string? GenerateStringKey;
        public readonly bool? IncludeSpace;
        public readonly int? PasswordLength;
        public readonly bool? RequireEachIncludedType;
        public readonly string? SecretStringTemplate;

        [OutputConstructor]
        private SecretGenerateSecretString(
            string? excludeCharacters,

            bool? excludeLowercase,

            bool? excludeNumbers,

            bool? excludePunctuation,

            bool? excludeUppercase,

            string? generateStringKey,

            bool? includeSpace,

            int? passwordLength,

            bool? requireEachIncludedType,

            string? secretStringTemplate)
        {
            ExcludeCharacters = excludeCharacters;
            ExcludeLowercase = excludeLowercase;
            ExcludeNumbers = excludeNumbers;
            ExcludePunctuation = excludePunctuation;
            ExcludeUppercase = excludeUppercase;
            GenerateStringKey = generateStringKey;
            IncludeSpace = includeSpace;
            PasswordLength = passwordLength;
            RequireEachIncludedType = requireEachIncludedType;
            SecretStringTemplate = secretStringTemplate;
        }
    }
}
