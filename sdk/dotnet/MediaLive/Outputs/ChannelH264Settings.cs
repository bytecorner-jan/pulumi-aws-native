// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelH264Settings
    {
        public readonly string? AdaptiveQuantization;
        public readonly string? AfdSignaling;
        public readonly int? Bitrate;
        public readonly int? BufFillPct;
        public readonly int? BufSize;
        public readonly string? ColorMetadata;
        public readonly Outputs.ChannelH264ColorSpaceSettings? ColorSpaceSettings;
        public readonly string? EntropyEncoding;
        public readonly Outputs.ChannelH264FilterSettings? FilterSettings;
        public readonly string? FixedAfd;
        public readonly string? FlickerAq;
        public readonly string? ForceFieldPictures;
        public readonly string? FramerateControl;
        public readonly int? FramerateDenominator;
        public readonly int? FramerateNumerator;
        public readonly string? GopBReference;
        public readonly int? GopClosedCadence;
        public readonly int? GopNumBFrames;
        public readonly double? GopSize;
        public readonly string? GopSizeUnits;
        public readonly string? Level;
        public readonly string? LookAheadRateControl;
        public readonly int? MaxBitrate;
        public readonly int? MinIInterval;
        public readonly int? NumRefFrames;
        public readonly string? ParControl;
        public readonly int? ParDenominator;
        public readonly int? ParNumerator;
        public readonly string? Profile;
        public readonly string? QualityLevel;
        public readonly int? QvbrQualityLevel;
        public readonly string? RateControlMode;
        public readonly string? ScanType;
        public readonly string? SceneChangeDetect;
        public readonly int? Slices;
        public readonly int? Softness;
        public readonly string? SpatialAq;
        public readonly string? SubgopLength;
        public readonly string? Syntax;
        public readonly string? TemporalAq;
        public readonly string? TimecodeInsertion;

        [OutputConstructor]
        private ChannelH264Settings(
            string? adaptiveQuantization,

            string? afdSignaling,

            int? bitrate,

            int? bufFillPct,

            int? bufSize,

            string? colorMetadata,

            Outputs.ChannelH264ColorSpaceSettings? colorSpaceSettings,

            string? entropyEncoding,

            Outputs.ChannelH264FilterSettings? filterSettings,

            string? fixedAfd,

            string? flickerAq,

            string? forceFieldPictures,

            string? framerateControl,

            int? framerateDenominator,

            int? framerateNumerator,

            string? gopBReference,

            int? gopClosedCadence,

            int? gopNumBFrames,

            double? gopSize,

            string? gopSizeUnits,

            string? level,

            string? lookAheadRateControl,

            int? maxBitrate,

            int? minIInterval,

            int? numRefFrames,

            string? parControl,

            int? parDenominator,

            int? parNumerator,

            string? profile,

            string? qualityLevel,

            int? qvbrQualityLevel,

            string? rateControlMode,

            string? scanType,

            string? sceneChangeDetect,

            int? slices,

            int? softness,

            string? spatialAq,

            string? subgopLength,

            string? syntax,

            string? temporalAq,

            string? timecodeInsertion)
        {
            AdaptiveQuantization = adaptiveQuantization;
            AfdSignaling = afdSignaling;
            Bitrate = bitrate;
            BufFillPct = bufFillPct;
            BufSize = bufSize;
            ColorMetadata = colorMetadata;
            ColorSpaceSettings = colorSpaceSettings;
            EntropyEncoding = entropyEncoding;
            FilterSettings = filterSettings;
            FixedAfd = fixedAfd;
            FlickerAq = flickerAq;
            ForceFieldPictures = forceFieldPictures;
            FramerateControl = framerateControl;
            FramerateDenominator = framerateDenominator;
            FramerateNumerator = framerateNumerator;
            GopBReference = gopBReference;
            GopClosedCadence = gopClosedCadence;
            GopNumBFrames = gopNumBFrames;
            GopSize = gopSize;
            GopSizeUnits = gopSizeUnits;
            Level = level;
            LookAheadRateControl = lookAheadRateControl;
            MaxBitrate = maxBitrate;
            MinIInterval = minIInterval;
            NumRefFrames = numRefFrames;
            ParControl = parControl;
            ParDenominator = parDenominator;
            ParNumerator = parNumerator;
            Profile = profile;
            QualityLevel = qualityLevel;
            QvbrQualityLevel = qvbrQualityLevel;
            RateControlMode = rateControlMode;
            ScanType = scanType;
            SceneChangeDetect = sceneChangeDetect;
            Slices = slices;
            Softness = softness;
            SpatialAq = spatialAq;
            SubgopLength = subgopLength;
            Syntax = syntax;
            TemporalAq = temporalAq;
            TimecodeInsertion = timecodeInsertion;
        }
    }
}
