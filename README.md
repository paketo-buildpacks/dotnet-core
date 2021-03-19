# .NET Core Paketo Buildpack
## `gcr.io/paketo-buildpacks/dotnet-core`

The .NET Core Paketo Buildpack provides a set of collaborating buildpacks that
enable the building of a Dotnet Core-based application. These buildpacks include:
- [Dotnet Core Runtime CNB](https://github.com/paketo-buildpacks/dotnet-core-runtime)
- [Dotnet Core ASPNet CNB](https://github.com/paketo-buildpacks/dotnet-core-aspnet)
- [Dotnet Core SDK CNB](https://github.com/paketo-buildpacks/dotnet-core-sdk)
- [Dotnet Publish CNB](https://github.com/paketo-buildpacks/dotnet-publish)
- [Dotnet Execute CNB](https://github.com/paketo-buildpacks/dotnet-execute)

The buildpack supports building both runtime-dependent and self-contained
applications. Additionally, the buildpack supports a number of combinations of
runtimes, SDKs, and ASPNet library configurations.

The buildpack supports the inclusion of
[NodeJS](https://nodejs.org) for applications that may require it. This support
is provided by the [Node Engine
CNB](https://github.com/paketo-buildpacks/node-engine).

Usage examples can be found in the
[`samples` repository under the `dotnet-core` directory](https://github.com/paketo-buildpacks/samples/tree/main/dotnet-core).

#### The .Net Core buildpack is compatible with the following builder(s):
- [Paketo Full Builder](https://github.com/paketo-buildpacks/full-builder)
- [Paketo Base Builder](https://github.com/paketo-buildpacks/base-builder)

This buildpack also includes the following utility buildpacks:
- [Procfile CNB](https://github.com/paketo-buildpacks/procfile)
- [Environment Variables CNB](https://github.com/paketo-buildpacks/environment-variables)
- [Image Labels CNB](https://github.com/paketo-buildpacks/image-labels)
- [CA Certificates CNB](https://github.com/paketo-buildpacks/ca-certificates)

Check out the [Paketo .NET Core CNB
docs](https://paketo.io/docs/buildpacks/language-family-buildpacks/dotnet-core/)
for more information.
