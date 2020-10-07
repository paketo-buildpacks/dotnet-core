# .NET Core Paketo Buildpack
## `gcr.io/paketo-buildpacks/dotnet-core`

The .NET Core Paketo Buildpack provides a set of collaborating buildpacks that
enable the building of a Dotnet Core-based application. These buildpacks include:
- [Dotnet Core Runtime CNB](https://github.com/paketo-buildpacks/dotnet-core-runtime)
- [Dotnet Core ASPNet CNB](https://github.com/paketo-buildpacks/dotnet-core-aspnet)
- [Dotnet Core SDK CNB](https://github.com/paketo-buildpacks/dotnet-core-sdk)
- [Dotnet Core Build CNB](https://github.com/paketo-buildpacks/dotnet-core-build)
- [Dotnet Core Execute CNB](https://github.com/paketo-buildpacks/dotnet-core-execute)

The buildpack supports building both runtime-dependent and self-contained
applications. Additionally, the buildpack supports a number of combinations of
runtimes, SDKs, and ASPNet library configurations.

The buildpack supports the inclusion of
[NodeJS](https://nodejs.org) for applications that may require it. This support
is provided by the [Node Engine
CNB](https://github.com/paketo-buildpacks/node-engine).

Usage examples can be found in the
[`samples` repository under the `dotnet-core` directory](https://github.com/paketo-buildpacks/samples/tree/main/dotnet-core).
