// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Outputs
{

    [OutputType]
    public sealed class RecordSetGeoLocation
    {
        public readonly string? ContinentCode;
        public readonly string? CountryCode;
        public readonly string? SubdivisionCode;

        [OutputConstructor]
        private RecordSetGeoLocation(
            string? continentCode,

            string? countryCode,

            string? subdivisionCode)
        {
            ContinentCode = continentCode;
            CountryCode = countryCode;
            SubdivisionCode = subdivisionCode;
        }
    }
}
