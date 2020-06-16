# Restructure to simplify buildpack interactions

## Proposal

Simplify the buildpack by rationalizing the inter-buildpack `build` and
`launch` requirements. This builds a clean separation between the concerns for
building an application and those concerns for running it. This simplified
structure would include the following buildpacks in addition to the existing
`icu` and `node-engine` buildpacks:

* dotnet-core-runtime: installs a runtime version, including the `dotnet`
  command line tool, making it available on the `$PATH`
  * provides: `dotnet-runtime`
  * requires: none

* dotnet-core-sdk: installs an SDK version
  * provides: `dotnet-sdk`
  * requires: none

* dotnet-core-aspnet: installs an ASPNet verison
  * provides: `dotnet-aspnet`
  * requires: none

* dotnet-publish: resolves versions of runtime, SDK, and ASPNet that are
  required to build and executes `dotnet publish`
  * provides: `dotnet-application`
  * requires: `dotnet-runtime`, `dotnet-sdk`, `dotnet-aspnet`, `icu`, and
    `node-engine` only at `build`

* dotnet-execute: sets the start command, requiring the necessary launch-time
  dependencies
  * provides: none
  * requires: `dotnet-runtime`, `dotnet-sdk`, `dotnet-aspnet`, `icu`, and
    `node-engine` only at `launch`

This structure would look like the following:

![Proposed Structure](/rfcs/assets/0001-proposed.png)

<div>
  <div><b>Legend</b></div>
  <div><code>launch</code> requirements are red</div>
  <div><code>build</code> requirements are blue</div>
</div>
<br/>

This graph of buildpacks would ultimately result in a language-family
`buildpack.toml` with the following order grouping:

```toml
[[order]] # Source

  [[order.group]]
    id = "paketo-buildpacks/dotnet-core-runtime"
    version = ""

  [[order.group]]
    id = "paketo-buildpacks/dotnet-core-aspnet"
    version = ""
    optional = true

  [[order.group]]
    id = "paketo-buildpacks/dotnet-core-sdk"
    version = ""

  [[order.group]]
    id = "paketo-buildpacks/icu"
    version = ""
    optional = true

  [[order.group]]
    id = "paketo-buildpacks/node-engine"
    version = ""
    optional = true

  [[order.group]]
    id = "paketo-buildpacks/dotnet-publish"
    version = ""

  [[order.group]]
    id = "paketo-buildpacks/dotnet-execute"
    version = ""

[[order]] # Framework-dependent deployment and Framework-dependent executable

  [[order.group]]
    id = "paketo-buildpacks/dotnet-core-runtime"
    version = ""

  [[order.group]]
    id = "paketo-buildpacks/dotnet-core-aspnet"
    version = ""
    optional = true

  [[order.group]]
    id = "paketo-buildpacks/icu"
    version = ""
    optional = true

  [[order.group]]
    id = "paketo-buildpacks/node-engine"
    version = ""
    optional = true

  [[order.group]]
    id = "paketo-buildpacks/dotnet-execute"
    version = ""

[[order]] # Self-contained deployment

  [[order.group]]
    id = "paketo-buildpacks/icu"
    version = ""
    optional = true

  [[order.group]]
    id = "paketo-buildpacks/node-engine"
    version = ""
    optional = true

  [[order.group]]
    id = "paketo-buildpacks/dotnet-execute"
    version = ""
```

## Motivation

The current Dotnet-Core buildpack has a high degree of coupling between the
interacting buildpacks. This is most evident when viewing a graph of the build
plan requirements between buildpacks in its current form.

![Original Structure](/rfcs/assets/0001-original.png)

<div>
  <div><b>Legend</b></div>
  <div><code>launch</code> requirements are red</div>
  <div><code>build</code> requirements are blue</div>
</div>
<br/>

In this graph, we can see that nearly every buildpack requires itself for
either `launch` or `build`. The pattern of providing and requiring a
buildpack's own dependency is one that we've started to view as an anti-pattern
in most cases. Specifically, we've identified that for most buildpacks that
just provide a dependency, the buildpack cannot determine easily that it should
be required during either the `build` or `launch` phase. So, instead of
optimistically requiring oneself, buildpacks should simply report that they
provide a dependency and expect that a downstream buildpack will require that
dependency in the correct phase for their use-case.

Additionally, many of the buildpacks require others for `build` or `launch`
when they have no good reason to. Examples of this include `dotnet-core-sdk`
requiring `dotnet-core-aspnet` or `dotnet-core-build` requiring nearly all of
the buildpacks for `launch` when it doesn't set a start command. This linking
of buildpacks is mostly required because the buildpacks share a `$DOTNET_ROOT`
directory where they provide dependencies. We should remove these requirements
and realign the buildpacks so that the right buildpacks require only what they
need for their own purposes. To solve the shared `$DOTNET_ROOT` directory, we
should create that directory in the working directory and symlink from it to
buildpack layers that provide these dependencies.

There are also other cases where buildpacks do not require dependencies they
need for `launch`. The `dotnet-core-conf` buildpack sets a start command, but
doesn't require `node-engine`, `dotnet-core-runtime`, `dotnet-core-aspnet`, or
`dotnet-core-sdk` during the `launch` phase. Instead the only way that these
dependencies appear in the launch image is because they are overzealously
required by other buildpacks in the group, as described above. Instead, the
`dotnet-core-conf` buildpack should directly require its required dependencies
for launch.

Simplifying the structure will also allow us to rationalize the purpose of each
buildpack. Moving behavior to the right parts of the buildpack family will help
future developers to navigate these codebases and users to understand how the
buildpack behaves. An example of one of these rationalizations would be to move
the responsibility for adding the `dotnet` CLI onto the `$PATH` from the
`dotnet-sdk` buildpack to `dotnet-runtime` where that dependency is first
introduced.

## Unresolved Questions and Bikeshedding

{{REMOVE THIS SECTION BEFORE RATIFICATION!}}
