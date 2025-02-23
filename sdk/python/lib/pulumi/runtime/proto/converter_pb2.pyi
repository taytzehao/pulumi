"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
Copyright 2016-2023, Pulumi Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import pulumi.codegen.hcl_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing_extensions.final
class ConvertStateRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    MAPPER_TARGET_FIELD_NUMBER: builtins.int
    mapper_target: builtins.str
    """the gRPC target of the mapper service."""
    def __init__(
        self,
        *,
        mapper_target: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["mapper_target", b"mapper_target"]) -> None: ...

global___ConvertStateRequest = ConvertStateRequest

@typing_extensions.final
class ResourceImport(google.protobuf.message.Message):
    """A ResourceImport specifies a resource to import."""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    TYPE_FIELD_NUMBER: builtins.int
    NAME_FIELD_NUMBER: builtins.int
    ID_FIELD_NUMBER: builtins.int
    VERSION_FIELD_NUMBER: builtins.int
    PLUGINDOWNLOADURL_FIELD_NUMBER: builtins.int
    type: builtins.str
    """the type token for the resource."""
    name: builtins.str
    """the name of the resource."""
    id: builtins.str
    """the ID of the resource."""
    version: builtins.str
    """the provider version to use for the resource, if any."""
    pluginDownloadURL: builtins.str
    """the provider PluginDownloadURL to use for the resource, if any."""
    def __init__(
        self,
        *,
        type: builtins.str = ...,
        name: builtins.str = ...,
        id: builtins.str = ...,
        version: builtins.str = ...,
        pluginDownloadURL: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["id", b"id", "name", b"name", "pluginDownloadURL", b"pluginDownloadURL", "type", b"type", "version", b"version"]) -> None: ...

global___ResourceImport = ResourceImport

@typing_extensions.final
class ConvertStateResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    RESOURCES_FIELD_NUMBER: builtins.int
    @property
    def resources(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___ResourceImport]:
        """a list of resources to import."""
    def __init__(
        self,
        *,
        resources: collections.abc.Iterable[global___ResourceImport] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["resources", b"resources"]) -> None: ...

global___ConvertStateResponse = ConvertStateResponse

@typing_extensions.final
class ConvertProgramRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SOURCE_DIRECTORY_FIELD_NUMBER: builtins.int
    TARGET_DIRECTORY_FIELD_NUMBER: builtins.int
    MAPPER_TARGET_FIELD_NUMBER: builtins.int
    source_directory: builtins.str
    """the source directory containing the program to convert from."""
    target_directory: builtins.str
    """a target directory to write the resulting PCL code and project file to."""
    mapper_target: builtins.str
    """the gRPC target of the mapper service."""
    def __init__(
        self,
        *,
        source_directory: builtins.str = ...,
        target_directory: builtins.str = ...,
        mapper_target: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["mapper_target", b"mapper_target", "source_directory", b"source_directory", "target_directory", b"target_directory"]) -> None: ...

global___ConvertProgramRequest = ConvertProgramRequest

@typing_extensions.final
class ConvertProgramResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DIAGNOSTICS_FIELD_NUMBER: builtins.int
    @property
    def diagnostics(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[pulumi.codegen.hcl_pb2.Diagnostic]:
        """any diagnostics from code generation."""
    def __init__(
        self,
        *,
        diagnostics: collections.abc.Iterable[pulumi.codegen.hcl_pb2.Diagnostic] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["diagnostics", b"diagnostics"]) -> None: ...

global___ConvertProgramResponse = ConvertProgramResponse
