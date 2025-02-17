// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ChannelLatencyMode = {
    Normal: "NORMAL",
    Low: "LOW",
} as const;

/**
 * Channel latency mode.
 */
export type ChannelLatencyMode = (typeof ChannelLatencyMode)[keyof typeof ChannelLatencyMode];

export const ChannelType = {
    Standard: "STANDARD",
    Basic: "BASIC",
} as const;

/**
 * Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
 */
export type ChannelType = (typeof ChannelType)[keyof typeof ChannelType];

export const RecordingConfigurationState = {
    Creating: "CREATING",
    CreateFailed: "CREATE_FAILED",
    Active: "ACTIVE",
} as const;

/**
 * Recording Configuration State.
 */
export type RecordingConfigurationState = (typeof RecordingConfigurationState)[keyof typeof RecordingConfigurationState];
