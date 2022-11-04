# .NET Core Paketo Buildpack

## Documentation
Check out the [Paketo Buildpack for .NET Core
docs](https://paketo.io/docs/buildpacks/language-family-buildpacks/dotnet-core/)
for more information on how to use this buildpack. To edit the docs content published on the above page,
make changes to [this document](https://github.com/paketo-buildpacks/paketo-website/blob/main/content/docs/howto/dotnet-core.md).

## Distribution
The buildpack is distributed as an OCI image in the following places:
- `gcr.io/paketo-buildpacks/dotnet-core`
- `paketobuildpacks/dotnet-core`

The buildpack is also available as a `.cnb` or `.tgz` archive, attached to each Github release.

## About
The Paketo Buildpack for .NET Core provides a set of collaborating buildpacks that
enable the building of a .NET Core application. These buildpacks include:
- [Paketo Buildpack for ASP.NET Runtime](https://github.com/paketo-buildpacks/dotnet-core-aspnet-runtime)
- [Paketo Buildpack for .NET Core SDK](https://github.com/paketo-buildpacks/dotnet-core-sdk)
- [Paketo Buildpack for .NET Publish](https://github.com/paketo-buildpacks/dotnet-publish)
- [Paketo Buildpack for .NET Execute](https://github.com/paketo-buildpacks/dotnet-execute)

The buildpack supports building both runtime-dependent and self-contained
applications. Additionally, the buildpack supports a number of combinations of
runtimes, SDKs, and ASPNet library configurations.

The buildpack supports the inclusion of
[NodeJS](https://nodejs.org) for applications that may require it. This support
is provided by the [Paketo Buildpack for Node Engine](https://github.com/paketo-buildpacks/node-engine).

Usage examples can be found in the
[`samples` repository under the `dotnet-core` directory](https://github.com/paketo-buildpacks/samples/tree/main/dotnet-core).

### Builder compatibility
The Paketo Buildpack for .NET Core buildpack is compatible with the following builder(s):
- [Paketo Bionic Full Builder](https://github.com/paketo-buildpacks/full-builder)
- [Paketo Bionic Base Builder](https://github.com/paketo-buildpacks/base-builder)
- [Paketo Jammy Full Builder](https://github.com/paketo-buildpacks/builder-jammy-full)
- [Paketo Jammy Base Builder](https://github.com/paketo-buildpacks/builder-jammy-base)

This buildpack also includes the following utility buildpacks:
- [Paketo Buildpack for Procfile](https://github.com/paketo-buildpacks/procfile)
- [Paketo Buildpack for Environment Variables](https://github.com/paketo-buildpacks/environment-variables)
- [Paketo Buildpack for Image Labels](https://github.com/paketo-buildpacks/image-labels)
- [Paketo Buildpack for CA Certificates](https://github.com/paketo-buildpacks/ca-certificates)
