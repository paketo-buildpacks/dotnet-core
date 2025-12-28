# Paketo Buildpack for .NET Core

## `docker.io/paketobuildpacks/dotnet-core`

The Paketo Buildpack for .NET Core provides a set of collaborating buildpacks that
enable the building of a .NET Core application. These buildpacks include:

- [ASP.NET Runtime CNB](https://github.com/paketo-buildpacks/dotnet-core-aspnet-runtime)
- [,NET Core SDK CNB](https://github.com/paketo-buildpacks/dotnet-core-sdk)
- [.NET Publish CNB](https://github.com/paketo-buildpacks/dotnet-publish)
- [.NET Execute CNB](https://github.com/paketo-buildpacks/dotnet-execute)
- [Visual Studio Debugger CNB](https://github.com/paketo-buildpacks/vsdbg)
- [ICU (International Components for Unicode) CNB](https://github.com/paketo-buildpacks/icu)

The buildpack supports building both runtime-dependent and self-contained
applications. Additionally, the buildpack supports a number of combinations of
runtimes, SDKs, and ASP.NET library configurations.

The buildpack supports the inclusion of
[NodeJS](https://nodejs.org) for applications that may require it. This support
is provided by the [Paketo Buildpack for Node Engine](https://github.com/paketo-buildpacks/node-engine).

Usage examples can be found in the
[`samples` repository under the `dotnet-core` directory](https://github.com/paketo-buildpacks/samples/tree/main/dotnet-core).

#### Builder compatibility

The Paketo Buildpack for .NET Core buildpack is compatible with the following builder(s):
- [Paketo Noble Builder](https://github.com/paketo-buildpacks/ubuntu-noble-builder)
- [Paketo Jammy Full Builder](https://github.com/paketo-buildpacks/builder-jammy-full)
- [Paketo Jammy Base Builder](https://github.com/paketo-buildpacks/builder-jammy-base)

This buildpack also includes the following utility buildpacks:
- [Procfile CNB](https://github.com/paketo-buildpacks/procfile)
- [Environment Variables CNB](https://github.com/paketo-buildpacks/environment-variables)
- [Image Labels CNB](https://github.com/paketo-buildpacks/image-labels)
- [CA Certificates CNB](https://github.com/paketo-buildpacks/ca-certificates)
- [Watchexec CNB](https://github.com/paketo-buildpacks/watchexec)

## Documentation

Check out the [Paketo Buildpack for .NET Core
docs](https://paketo.io/docs/buildpacks/language-family-buildpacks/dotnet-core/)
for more information on how to use this buildpack. To edit the docs content published on the above page,
make changes to [this document](https://github.com/paketo-buildpacks/paketo-website/blob/main/content/docs/howto/dotnet-core.md).
