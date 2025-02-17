// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.EKS
{
    /// <summary>
    /// Resolve parameter value conflicts
    /// </summary>
    [EnumType]
    public readonly struct AddonResolveConflicts : IEquatable<AddonResolveConflicts>
    {
        private readonly string _value;

        private AddonResolveConflicts(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AddonResolveConflicts None { get; } = new AddonResolveConflicts("NONE");
        public static AddonResolveConflicts Overwrite { get; } = new AddonResolveConflicts("OVERWRITE");

        public static bool operator ==(AddonResolveConflicts left, AddonResolveConflicts right) => left.Equals(right);
        public static bool operator !=(AddonResolveConflicts left, AddonResolveConflicts right) => !left.Equals(right);

        public static explicit operator string(AddonResolveConflicts value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AddonResolveConflicts other && Equals(other);
        public bool Equals(AddonResolveConflicts other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
