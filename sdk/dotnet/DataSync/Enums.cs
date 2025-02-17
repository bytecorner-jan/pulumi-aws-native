// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.DataSync
{
    /// <summary>
    /// The service endpoints that the agent will connect to.
    /// </summary>
    [EnumType]
    public readonly struct AgentEndpointType : IEquatable<AgentEndpointType>
    {
        private readonly string _value;

        private AgentEndpointType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AgentEndpointType Fips { get; } = new AgentEndpointType("FIPS");
        public static AgentEndpointType Public { get; } = new AgentEndpointType("PUBLIC");
        public static AgentEndpointType PrivateLink { get; } = new AgentEndpointType("PRIVATE_LINK");

        public static bool operator ==(AgentEndpointType left, AgentEndpointType right) => left.Equals(right);
        public static bool operator !=(AgentEndpointType left, AgentEndpointType right) => !left.Equals(right);

        public static explicit operator string(AgentEndpointType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AgentEndpointType other && Equals(other);
        public bool Equals(AgentEndpointType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The authentication mode used to determine identity of user.
    /// </summary>
    [EnumType]
    public readonly struct LocationHDFSAuthenticationType : IEquatable<LocationHDFSAuthenticationType>
    {
        private readonly string _value;

        private LocationHDFSAuthenticationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LocationHDFSAuthenticationType Simple { get; } = new LocationHDFSAuthenticationType("SIMPLE");
        public static LocationHDFSAuthenticationType Kerberos { get; } = new LocationHDFSAuthenticationType("KERBEROS");

        public static bool operator ==(LocationHDFSAuthenticationType left, LocationHDFSAuthenticationType right) => left.Equals(right);
        public static bool operator !=(LocationHDFSAuthenticationType left, LocationHDFSAuthenticationType right) => !left.Equals(right);

        public static explicit operator string(LocationHDFSAuthenticationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LocationHDFSAuthenticationType other && Equals(other);
        public bool Equals(LocationHDFSAuthenticationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Configuration for Data Transfer Protection.
    /// </summary>
    [EnumType]
    public readonly struct LocationHDFSQopConfigurationDataTransferProtection : IEquatable<LocationHDFSQopConfigurationDataTransferProtection>
    {
        private readonly string _value;

        private LocationHDFSQopConfigurationDataTransferProtection(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LocationHDFSQopConfigurationDataTransferProtection Authentication { get; } = new LocationHDFSQopConfigurationDataTransferProtection("AUTHENTICATION");
        public static LocationHDFSQopConfigurationDataTransferProtection Integrity { get; } = new LocationHDFSQopConfigurationDataTransferProtection("INTEGRITY");
        public static LocationHDFSQopConfigurationDataTransferProtection Privacy { get; } = new LocationHDFSQopConfigurationDataTransferProtection("PRIVACY");
        public static LocationHDFSQopConfigurationDataTransferProtection Disabled { get; } = new LocationHDFSQopConfigurationDataTransferProtection("DISABLED");

        public static bool operator ==(LocationHDFSQopConfigurationDataTransferProtection left, LocationHDFSQopConfigurationDataTransferProtection right) => left.Equals(right);
        public static bool operator !=(LocationHDFSQopConfigurationDataTransferProtection left, LocationHDFSQopConfigurationDataTransferProtection right) => !left.Equals(right);

        public static explicit operator string(LocationHDFSQopConfigurationDataTransferProtection value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LocationHDFSQopConfigurationDataTransferProtection other && Equals(other);
        public bool Equals(LocationHDFSQopConfigurationDataTransferProtection other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Configuration for RPC Protection.
    /// </summary>
    [EnumType]
    public readonly struct LocationHDFSQopConfigurationRpcProtection : IEquatable<LocationHDFSQopConfigurationRpcProtection>
    {
        private readonly string _value;

        private LocationHDFSQopConfigurationRpcProtection(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LocationHDFSQopConfigurationRpcProtection Authentication { get; } = new LocationHDFSQopConfigurationRpcProtection("AUTHENTICATION");
        public static LocationHDFSQopConfigurationRpcProtection Integrity { get; } = new LocationHDFSQopConfigurationRpcProtection("INTEGRITY");
        public static LocationHDFSQopConfigurationRpcProtection Privacy { get; } = new LocationHDFSQopConfigurationRpcProtection("PRIVACY");
        public static LocationHDFSQopConfigurationRpcProtection Disabled { get; } = new LocationHDFSQopConfigurationRpcProtection("DISABLED");

        public static bool operator ==(LocationHDFSQopConfigurationRpcProtection left, LocationHDFSQopConfigurationRpcProtection right) => left.Equals(right);
        public static bool operator !=(LocationHDFSQopConfigurationRpcProtection left, LocationHDFSQopConfigurationRpcProtection right) => !left.Equals(right);

        public static explicit operator string(LocationHDFSQopConfigurationRpcProtection value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LocationHDFSQopConfigurationRpcProtection other && Equals(other);
        public bool Equals(LocationHDFSQopConfigurationRpcProtection other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The specific NFS version that you want DataSync to use to mount your NFS share.
    /// </summary>
    [EnumType]
    public readonly struct LocationNFSMountOptionsVersion : IEquatable<LocationNFSMountOptionsVersion>
    {
        private readonly string _value;

        private LocationNFSMountOptionsVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LocationNFSMountOptionsVersion Automatic { get; } = new LocationNFSMountOptionsVersion("AUTOMATIC");
        public static LocationNFSMountOptionsVersion Nfs3 { get; } = new LocationNFSMountOptionsVersion("NFS3");
        public static LocationNFSMountOptionsVersion Nfs40 { get; } = new LocationNFSMountOptionsVersion("NFS4_0");
        public static LocationNFSMountOptionsVersion Nfs41 { get; } = new LocationNFSMountOptionsVersion("NFS4_1");

        public static bool operator ==(LocationNFSMountOptionsVersion left, LocationNFSMountOptionsVersion right) => left.Equals(right);
        public static bool operator !=(LocationNFSMountOptionsVersion left, LocationNFSMountOptionsVersion right) => !left.Equals(right);

        public static explicit operator string(LocationNFSMountOptionsVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LocationNFSMountOptionsVersion other && Equals(other);
        public bool Equals(LocationNFSMountOptionsVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The protocol that the object storage server uses to communicate.
    /// </summary>
    [EnumType]
    public readonly struct LocationObjectStorageServerProtocol : IEquatable<LocationObjectStorageServerProtocol>
    {
        private readonly string _value;

        private LocationObjectStorageServerProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LocationObjectStorageServerProtocol Https { get; } = new LocationObjectStorageServerProtocol("HTTPS");
        public static LocationObjectStorageServerProtocol Http { get; } = new LocationObjectStorageServerProtocol("HTTP");

        public static bool operator ==(LocationObjectStorageServerProtocol left, LocationObjectStorageServerProtocol right) => left.Equals(right);
        public static bool operator !=(LocationObjectStorageServerProtocol left, LocationObjectStorageServerProtocol right) => !left.Equals(right);

        public static explicit operator string(LocationObjectStorageServerProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LocationObjectStorageServerProtocol other && Equals(other);
        public bool Equals(LocationObjectStorageServerProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The Amazon S3 storage class you want to store your files in when this location is used as a task destination.
    /// </summary>
    [EnumType]
    public readonly struct LocationS3S3StorageClass : IEquatable<LocationS3S3StorageClass>
    {
        private readonly string _value;

        private LocationS3S3StorageClass(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LocationS3S3StorageClass Standard { get; } = new LocationS3S3StorageClass("STANDARD");
        public static LocationS3S3StorageClass StandardIa { get; } = new LocationS3S3StorageClass("STANDARD_IA");
        public static LocationS3S3StorageClass OnezoneIa { get; } = new LocationS3S3StorageClass("ONEZONE_IA");
        public static LocationS3S3StorageClass IntelligentTiering { get; } = new LocationS3S3StorageClass("INTELLIGENT_TIERING");
        public static LocationS3S3StorageClass Glacier { get; } = new LocationS3S3StorageClass("GLACIER");
        public static LocationS3S3StorageClass DeepArchive { get; } = new LocationS3S3StorageClass("DEEP_ARCHIVE");

        public static bool operator ==(LocationS3S3StorageClass left, LocationS3S3StorageClass right) => left.Equals(right);
        public static bool operator !=(LocationS3S3StorageClass left, LocationS3S3StorageClass right) => !left.Equals(right);

        public static explicit operator string(LocationS3S3StorageClass value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LocationS3S3StorageClass other && Equals(other);
        public bool Equals(LocationS3S3StorageClass other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The specific SMB version that you want DataSync to use to mount your SMB share.
    /// </summary>
    [EnumType]
    public readonly struct LocationSMBMountOptionsVersion : IEquatable<LocationSMBMountOptionsVersion>
    {
        private readonly string _value;

        private LocationSMBMountOptionsVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LocationSMBMountOptionsVersion Automatic { get; } = new LocationSMBMountOptionsVersion("AUTOMATIC");
        public static LocationSMBMountOptionsVersion Smb2 { get; } = new LocationSMBMountOptionsVersion("SMB2");
        public static LocationSMBMountOptionsVersion Smb3 { get; } = new LocationSMBMountOptionsVersion("SMB3");

        public static bool operator ==(LocationSMBMountOptionsVersion left, LocationSMBMountOptionsVersion right) => left.Equals(right);
        public static bool operator !=(LocationSMBMountOptionsVersion left, LocationSMBMountOptionsVersion right) => !left.Equals(right);

        public static explicit operator string(LocationSMBMountOptionsVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LocationSMBMountOptionsVersion other && Equals(other);
        public bool Equals(LocationSMBMountOptionsVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.
    /// </summary>
    [EnumType]
    public readonly struct TaskFilterRuleFilterType : IEquatable<TaskFilterRuleFilterType>
    {
        private readonly string _value;

        private TaskFilterRuleFilterType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskFilterRuleFilterType SimplePattern { get; } = new TaskFilterRuleFilterType("SIMPLE_PATTERN");

        public static bool operator ==(TaskFilterRuleFilterType left, TaskFilterRuleFilterType right) => left.Equals(right);
        public static bool operator !=(TaskFilterRuleFilterType left, TaskFilterRuleFilterType right) => !left.Equals(right);

        public static explicit operator string(TaskFilterRuleFilterType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskFilterRuleFilterType other && Equals(other);
        public bool Equals(TaskFilterRuleFilterType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A file metadata value that shows the last time a file was accessed (that is, when the file was read or written to).
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsAtime : IEquatable<TaskOptionsAtime>
    {
        private readonly string _value;

        private TaskOptionsAtime(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsAtime None { get; } = new TaskOptionsAtime("NONE");
        public static TaskOptionsAtime BestEffort { get; } = new TaskOptionsAtime("BEST_EFFORT");

        public static bool operator ==(TaskOptionsAtime left, TaskOptionsAtime right) => left.Equals(right);
        public static bool operator !=(TaskOptionsAtime left, TaskOptionsAtime right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsAtime value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsAtime other && Equals(other);
        public bool Equals(TaskOptionsAtime other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The group ID (GID) of the file's owners.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsGid : IEquatable<TaskOptionsGid>
    {
        private readonly string _value;

        private TaskOptionsGid(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsGid None { get; } = new TaskOptionsGid("NONE");
        public static TaskOptionsGid IntValue { get; } = new TaskOptionsGid("INT_VALUE");
        public static TaskOptionsGid Name { get; } = new TaskOptionsGid("NAME");
        public static TaskOptionsGid Both { get; } = new TaskOptionsGid("BOTH");

        public static bool operator ==(TaskOptionsGid left, TaskOptionsGid right) => left.Equals(right);
        public static bool operator !=(TaskOptionsGid left, TaskOptionsGid right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsGid value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsGid other && Equals(other);
        public bool Equals(TaskOptionsGid other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that determines the types of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsLogLevel : IEquatable<TaskOptionsLogLevel>
    {
        private readonly string _value;

        private TaskOptionsLogLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsLogLevel Off { get; } = new TaskOptionsLogLevel("OFF");
        public static TaskOptionsLogLevel Basic { get; } = new TaskOptionsLogLevel("BASIC");
        public static TaskOptionsLogLevel Transfer { get; } = new TaskOptionsLogLevel("TRANSFER");

        public static bool operator ==(TaskOptionsLogLevel left, TaskOptionsLogLevel right) => left.Equals(right);
        public static bool operator !=(TaskOptionsLogLevel left, TaskOptionsLogLevel right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsLogLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsLogLevel other && Equals(other);
        public bool Equals(TaskOptionsLogLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that indicates the last time that a file was modified (that is, a file was written to) before the PREPARING phase.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsMtime : IEquatable<TaskOptionsMtime>
    {
        private readonly string _value;

        private TaskOptionsMtime(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsMtime None { get; } = new TaskOptionsMtime("NONE");
        public static TaskOptionsMtime Preserve { get; } = new TaskOptionsMtime("PRESERVE");

        public static bool operator ==(TaskOptionsMtime left, TaskOptionsMtime right) => left.Equals(right);
        public static bool operator !=(TaskOptionsMtime left, TaskOptionsMtime right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsMtime value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsMtime other && Equals(other);
        public bool Equals(TaskOptionsMtime other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that determines whether files at the destination should be overwritten or preserved when copying files.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsOverwriteMode : IEquatable<TaskOptionsOverwriteMode>
    {
        private readonly string _value;

        private TaskOptionsOverwriteMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsOverwriteMode Always { get; } = new TaskOptionsOverwriteMode("ALWAYS");
        public static TaskOptionsOverwriteMode Never { get; } = new TaskOptionsOverwriteMode("NEVER");

        public static bool operator ==(TaskOptionsOverwriteMode left, TaskOptionsOverwriteMode right) => left.Equals(right);
        public static bool operator !=(TaskOptionsOverwriteMode left, TaskOptionsOverwriteMode right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsOverwriteMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsOverwriteMode other && Equals(other);
        public bool Equals(TaskOptionsOverwriteMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsPosixPermissions : IEquatable<TaskOptionsPosixPermissions>
    {
        private readonly string _value;

        private TaskOptionsPosixPermissions(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsPosixPermissions None { get; } = new TaskOptionsPosixPermissions("NONE");
        public static TaskOptionsPosixPermissions Preserve { get; } = new TaskOptionsPosixPermissions("PRESERVE");

        public static bool operator ==(TaskOptionsPosixPermissions left, TaskOptionsPosixPermissions right) => left.Equals(right);
        public static bool operator !=(TaskOptionsPosixPermissions left, TaskOptionsPosixPermissions right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsPosixPermissions value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsPosixPermissions other && Equals(other);
        public bool Equals(TaskOptionsPosixPermissions other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that specifies whether files in the destination that don't exist in the source file system should be preserved.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsPreserveDeletedFiles : IEquatable<TaskOptionsPreserveDeletedFiles>
    {
        private readonly string _value;

        private TaskOptionsPreserveDeletedFiles(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsPreserveDeletedFiles Preserve { get; } = new TaskOptionsPreserveDeletedFiles("PRESERVE");
        public static TaskOptionsPreserveDeletedFiles Remove { get; } = new TaskOptionsPreserveDeletedFiles("REMOVE");

        public static bool operator ==(TaskOptionsPreserveDeletedFiles left, TaskOptionsPreserveDeletedFiles right) => left.Equals(right);
        public static bool operator !=(TaskOptionsPreserveDeletedFiles left, TaskOptionsPreserveDeletedFiles right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsPreserveDeletedFiles value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsPreserveDeletedFiles other && Equals(other);
        public bool Equals(TaskOptionsPreserveDeletedFiles other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that determines whether AWS DataSync should preserve the metadata of block and character devices in the source file system, and recreate the files with that device name and metadata on the destination.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsPreserveDevices : IEquatable<TaskOptionsPreserveDevices>
    {
        private readonly string _value;

        private TaskOptionsPreserveDevices(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsPreserveDevices None { get; } = new TaskOptionsPreserveDevices("NONE");
        public static TaskOptionsPreserveDevices Preserve { get; } = new TaskOptionsPreserveDevices("PRESERVE");

        public static bool operator ==(TaskOptionsPreserveDevices left, TaskOptionsPreserveDevices right) => left.Equals(right);
        public static bool operator !=(TaskOptionsPreserveDevices left, TaskOptionsPreserveDevices right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsPreserveDevices value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsPreserveDevices other && Equals(other);
        public bool Equals(TaskOptionsPreserveDevices other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that determines which components of the SMB security descriptor are copied during transfer.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsSecurityDescriptorCopyFlags : IEquatable<TaskOptionsSecurityDescriptorCopyFlags>
    {
        private readonly string _value;

        private TaskOptionsSecurityDescriptorCopyFlags(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsSecurityDescriptorCopyFlags None { get; } = new TaskOptionsSecurityDescriptorCopyFlags("NONE");
        public static TaskOptionsSecurityDescriptorCopyFlags OwnerDacl { get; } = new TaskOptionsSecurityDescriptorCopyFlags("OWNER_DACL");
        public static TaskOptionsSecurityDescriptorCopyFlags OwnerDaclSacl { get; } = new TaskOptionsSecurityDescriptorCopyFlags("OWNER_DACL_SACL");

        public static bool operator ==(TaskOptionsSecurityDescriptorCopyFlags left, TaskOptionsSecurityDescriptorCopyFlags right) => left.Equals(right);
        public static bool operator !=(TaskOptionsSecurityDescriptorCopyFlags left, TaskOptionsSecurityDescriptorCopyFlags right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsSecurityDescriptorCopyFlags value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsSecurityDescriptorCopyFlags other && Equals(other);
        public bool Equals(TaskOptionsSecurityDescriptorCopyFlags other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that determines whether tasks should be queued before executing the tasks.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsTaskQueueing : IEquatable<TaskOptionsTaskQueueing>
    {
        private readonly string _value;

        private TaskOptionsTaskQueueing(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsTaskQueueing Enabled { get; } = new TaskOptionsTaskQueueing("ENABLED");
        public static TaskOptionsTaskQueueing Disabled { get; } = new TaskOptionsTaskQueueing("DISABLED");

        public static bool operator ==(TaskOptionsTaskQueueing left, TaskOptionsTaskQueueing right) => left.Equals(right);
        public static bool operator !=(TaskOptionsTaskQueueing left, TaskOptionsTaskQueueing right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsTaskQueueing value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsTaskQueueing other && Equals(other);
        public bool Equals(TaskOptionsTaskQueueing other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsTransferMode : IEquatable<TaskOptionsTransferMode>
    {
        private readonly string _value;

        private TaskOptionsTransferMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsTransferMode Changed { get; } = new TaskOptionsTransferMode("CHANGED");
        public static TaskOptionsTransferMode All { get; } = new TaskOptionsTransferMode("ALL");

        public static bool operator ==(TaskOptionsTransferMode left, TaskOptionsTransferMode right) => left.Equals(right);
        public static bool operator !=(TaskOptionsTransferMode left, TaskOptionsTransferMode right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsTransferMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsTransferMode other && Equals(other);
        public bool Equals(TaskOptionsTransferMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The user ID (UID) of the file's owner.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsUid : IEquatable<TaskOptionsUid>
    {
        private readonly string _value;

        private TaskOptionsUid(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsUid None { get; } = new TaskOptionsUid("NONE");
        public static TaskOptionsUid IntValue { get; } = new TaskOptionsUid("INT_VALUE");
        public static TaskOptionsUid Name { get; } = new TaskOptionsUid("NAME");
        public static TaskOptionsUid Both { get; } = new TaskOptionsUid("BOTH");

        public static bool operator ==(TaskOptionsUid left, TaskOptionsUid right) => left.Equals(right);
        public static bool operator !=(TaskOptionsUid left, TaskOptionsUid right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsUid value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsUid other && Equals(other);
        public bool Equals(TaskOptionsUid other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value that determines whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred.
    /// </summary>
    [EnumType]
    public readonly struct TaskOptionsVerifyMode : IEquatable<TaskOptionsVerifyMode>
    {
        private readonly string _value;

        private TaskOptionsVerifyMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskOptionsVerifyMode PointInTimeConsistent { get; } = new TaskOptionsVerifyMode("POINT_IN_TIME_CONSISTENT");
        public static TaskOptionsVerifyMode OnlyFilesTransferred { get; } = new TaskOptionsVerifyMode("ONLY_FILES_TRANSFERRED");
        public static TaskOptionsVerifyMode None { get; } = new TaskOptionsVerifyMode("NONE");

        public static bool operator ==(TaskOptionsVerifyMode left, TaskOptionsVerifyMode right) => left.Equals(right);
        public static bool operator !=(TaskOptionsVerifyMode left, TaskOptionsVerifyMode right) => !left.Equals(right);

        public static explicit operator string(TaskOptionsVerifyMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskOptionsVerifyMode other && Equals(other);
        public bool Equals(TaskOptionsVerifyMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the task that was described.
    /// </summary>
    [EnumType]
    public readonly struct TaskStatus : IEquatable<TaskStatus>
    {
        private readonly string _value;

        private TaskStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskStatus Available { get; } = new TaskStatus("AVAILABLE");
        public static TaskStatus Creating { get; } = new TaskStatus("CREATING");
        public static TaskStatus Queued { get; } = new TaskStatus("QUEUED");
        public static TaskStatus Running { get; } = new TaskStatus("RUNNING");
        public static TaskStatus Unavailable { get; } = new TaskStatus("UNAVAILABLE");

        public static bool operator ==(TaskStatus left, TaskStatus right) => left.Equals(right);
        public static bool operator !=(TaskStatus left, TaskStatus right) => !left.Equals(right);

        public static explicit operator string(TaskStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskStatus other && Equals(other);
        public bool Equals(TaskStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
