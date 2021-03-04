# Support Building Optimized Apps

## Proposal

Addition of an  environment variable to
determine the value of the `--self-contained` flag when running
`dotnet_publish`.

## Motivation

Some .NET Core applications make use of "app trimming" to reduce their size
during builds. Application  which make use of this functionality (e.g. Blazor
Web Assembly apps) [fail to
build](https://github.com/paketo-buildpacks/dotnet-publish/issues/145) when it
is disabled. This behaviour can be enabled either by setting
`--self-contained=true` explicitly or omitting the flag completely when running
the `dotnet publish` command. The Dotnet Buildpack currently sets
`--self-contained=false` by default, and the lack of existing configuration
options forces users into customizing the buildpack to suit their needs.

## Implementation (Optional)

It is proposed that the buildpack uses the `BP_DOTNET_PUBLISH_SELF_CONTAINED`
environment variable to decide on the value for the `--self-contained` flag
during the build phase.

## Source Material (Optional)

- [App Trimmming (.NET article)](https://devblogs.microsoft.com/dotnet/app-trimming-in-net-5/)
- [Example issue](https://github.com/paketo-buildpacks/dotnet-publish/issues/145)

## Alternatives

- Add a more general `BP_DOTNET_PUBLISH_ARGS` environment variable which could
  enable support for the entire breadth of flags available for `dotnet
  publish`.
  Pros: Allows for greater customization, akin to `LD_FLAGS` in Go.
  Cons: May introduce unneccessary complexity.

- Attempt to automatically detect the necessary value for the `--self-contained` flag.
  Pros: Users need not concern themselves with any extra configuration.
  Cons: There is not yet an obvious way to determine which apps use "app trimming".


## Unresolved Questions and Bikeshedding (Optional)

{{Write about any arbitrary decisions that need to be made (syntax, colors, formatting, minor UX decisions), and any questions for the proposal that have not been answered.}}

