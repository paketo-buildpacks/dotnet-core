---
title: ".Net Core Buildpack"
weight: 301
menu:
  main:
    parent: "language-family-buildpacks"
---

The [.Net Core Paketo Buildpack](https://github.com/paketo-buildpacks/dotnet-core)
supports building several configurations of .Net Core applications.

To build your app locally with the buildpack using the `pack` CLI, run

{{< code/copyable >}}
git clone https://github.com/paketo-buildpacks/samples
cd samples/dotnet-core/aspnet
pack build my-app --buildpack gcr.io/paketo-buildpacks/dotnet-core \
  --builder paketobuildpacks/builder:base
{{< /code/copyable >}}

See
[samples](https://github.com/paketo-buildpacks/samples/tree/main/dotnet-core/aspnet)
for how to run the app.

**NOTE: Though the example above uses the Paketo Base builder, this buildpack is
also compatible with the Paketo Full builder.**

{{< table_of_contents >}}

## Supported Dependencies

The .Net Core Paketo Buildpack supports several versions of the .Net Core Framework.
For more details on the specific versions supported in a given buildpack
version, see the [release
notes](https://github.com/paketo-buildpacks/dotnet-core/releases).

## Application Types

The .Net Core Buildpack supports several types of application source code that
can be built into a container image. Developers can provide raw source code, or
built artifacts like Framework-Dependent Deployments/Executables or
Self-Contained Deployments when building their application.

### Source Applications

The .Net Core Build Buildpack is capable of building application source code
into Framework-Dependent Deployments (FDD) or Executables (FDE).  This is
achieved using the `dotnet publish` command. For .Net Core Framework 2.x
applications, [an FDD is
produced](https://docs.microsoft.com/en-us/dotnet/core/deploying/deploy-with-cli#framework-dependent-deployment)
as the default build artifact, while [an FDE is
produced](https://docs.microsoft.com/en-us/dotnet/core/deploying/deploy-with-cli#framework-dependent-executable)
when the application source is for .Net Core Framework 3.x.

### Framework-Dependent Deployments or Framework-Dependent Executables

When building an application that has already been published as a
Framework-Dependent Deployment or Framework-Dependent Executable, the buildpack
will include the required .Net Core Framework dependencies and set the start
command.

### Self-Contained Deployment

When building an application as a [Self-Contained
Deployment](https://docs.microsoft.com/en-us/dotnet/core/deploying/deploy-with-cli#self-contained-deployment)(SCD),
the buildpack will ensure the correct start command will be used to run your
app. No .Net Core Framework dependencies will be included in the built image as
they are already included in the SCD artifact.

## Specifying Runtime and ASP.Net Versions

The .Net Core Runtime and .Net Core ASP.Net Buildpacks allow you to specify a
version of the .Net Core Runtime and ASP.Net to use during deployment. This
version can be specified in several ways including through a
`runtimeconfig.json`, MSBuild Project file, or build-time environment
variables. When specifying a version of the .Net Core Runtime and ASP.Net, you
must choose a version that is available within these buildpacks. These versions
can be found in the [.Net Core Runtime release
notes](https://github.com/paketo-buildpacks/dotnet-core-runtime/releases) and
[.Net Core ASP.Net release
notes](https://github.com/paketo-buildpacks/dotnet-core-aspnet/releases).

.Net Core ASP.Net will only be included in the build process if your
application declares its Runtime Framework as either `Microsoft.AspNetCore.App`
or `Microsoft.AspNetCore.All`.

### Using `BP_DOTNET_FRAMEWORK_VERSION`

To configure the buildpack to use a certain version of the .Net Core Runtime
and ASP.Net when deploying your app, set the `$BP_DOTNET_FRAMEWORK` environment
variable at build time, either by passing a flag to the
[platform](https://buildpacks.io/docs/concepts/components/platform/) or by
adding it to your `project.toml`. See the Cloud Native Buildpacks
[documentation](https://buildpacks.io/docs/app-developer-guide/using-project-descriptor/)
to learn more about `project.toml` files.

**With a `pack build` flag**
{{< code/copyable >}}
pack build myapp --env BP_DOTNET_FRAMEWORK_VERSION=5.0.4
{{< /code/copyable >}}

**In a `project.toml` file**
{{< code/copyable >}}
[[ build.env ]]
  name = 'BP_DOTNET_FRAMEWORK_VERSION'
  value = '5.0.4'
{{< /code/copyable >}}

**Note**: If you specify a particular version using the above environment
variable, the buildpack **will not** run runtime version roll-forward logic. To
learn more about roll-forward logic, see the [Microsoft .Net Runtime
documentation](https://docs.microsoft.com/en-us/dotnet/core/versions/selection#framework-dependent-apps-roll-forward).

### Using runtimeconfig.json

If you are using a
[`runtimeconfig.json`](https://docs.microsoft.com/en-us/dotnet/core/run-time-config/)
file, you can specify the .Net Core Runtime version within that file. To
configure the buildpack to use .Net Core Runtime v2.1.14 when deploying your
app, include the values below in your `runtimeconfig.json` file:

{{< code/copyable >}}
{
  "runtimeOptions": {
    "framework": {
      "version": "2.1.14"
    }
  }
}
{{< /code/copyable >}}

### Using a Project file

If you are using a Project file (eg. `*.csproj`, `*.fsproj`, or `*.vbproj`), you can specify
the .Net Core Runtime version within that file. To configure the buildpack to
use .Net Core Runtime v2.1.14 when deploying your app, include the values below
in your Project file:

{{< code/copyable >}}
<Project>
  <PropertyGroup>
    <RuntimeFrameworkVersion>2.1.14</RuntimeFrameworkVersion>
  </PropertyGroup>
</Project>
{{< /code/copyable >}}

Alternatively, for applications that do not rely upon a specific .Net Core
Runtime patch version, you can specify the Target Framework and the buildpack
will choose the appropriate .Net Core Runtime version. To configure the
buildpack to use a .Net Core Runtime version in the 2.1 .Net Core Target Framework
when deploying your app, include the values below in your Project file:

{{< code/copyable >}}
<Project>
  <PropertyGroup>
    <TargetFramework>netcoreapp2.1</TargetFramework>
  </PropertyGroup>
</Project>
{{< /code/copyable >}}

For more details about specifying a .Net Core version using a Project file,
please review the [Microsoft
documentation](https://docs.microsoft.com/en-us/dotnet/core/versions/selection).

### .Net Core Framework Version Selection

The .Net Core Buildpack uses the same version selection policy that Microsoft
has put together for .Net Core Framework. If you would like to know more about
the policy please refer to this
[documentation](https://docs.microsoft.com/en-us/dotnet/core/versions/selection)
provided by Microsoft.

### Deprecated: Using buildpack.yml

Specifying the .Net Core Framework version through `buildpack.yml`
configuration will be deprecated in .Net Core Runtime and .Net Core ASPNET
Buildpacks v1.0.0.  To migrate from using `buildpack.yml`, please set the
`BP_DOTNET_FRAMEWORK_VERSION` environment variable.

## Specifying an SDK Version

By default, the .Net Core SDK Buildpack installs the latest available patch
version of the SDK that is compatible with the installed .Net Core runtime.
The available SDK versions for each buildpack release can be found in the
[release notes](https://github.com/paketo-buildpacks/dotnet-core-sdk/releases).

However, the .Net Core SDK version can be explicitly set by specifying a
version in a `buildpack.yml` file.

### Deprecated: Using buildpack.yml

Specifying the .Net Core SDK version through `buildpack.yml` configuration will
be deprecated in .Net Core SDK Buildpack v1.0.0.

Because versions of the .NET Core runtime and .NET Core SDK dependencies are so
tightly coupled, most users should instead use the
`BP_DOTNET_FRAMEWORK_VERSION` environment variable to specify which version of
the .NET Core runtime that the .NET Core Runtime Buildpack should install. The
.Net Core SDK buildpack will automatically install an SDK version that is
compatible with the selected .NET Core runtime version.

## Specifying a Custom Project Path

By default, the .Net Core Build Buildpack will consider the root directory of
your codebase to be the project directory. This directory should contain a C#,
F#, or Visual Basic Project file. If your project directory is not located at
the root of your source code you will need to set a custom project path.

### Using `BP_DOTNET_PROJECT_PATH`

You can specify a project path by setting the `$BP_DOTNET_PROJECT_PATH`
environment variable at build time, either by passing a flag to the
[platform](https://buildpacks.io/docs/concepts/components/platform/) or by
adding it to your `project.toml`. See the Cloud Native Buildpacks
[documentation](https://buildpacks.io/docs/app-developer-guide/using-project-descriptor/)
to learn more about `project.toml` files.

**With a `pack build` flag**
{{< code/copyable >}}
pack build my-app --env BP_DOTNET_PROJECT_PATH=./src/my-app
{{< /code/copyable >}}


**In a `project.toml` file**
{{< code/copyable >}}
[[ build.env ]]
  name = 'BP_DOTNET_PROJECT_PATH'
  value = './src/my-app'
{{< /code/copyable >}}
See the [Cloud Native Buildpacks
documentation](https://buildpacks.io/docs/app-developer-guide/using-project-descriptor/)
to learn more about `project.toml` files.

### Deprecated: Using buildpack.yml

Specifying the project path through `buildpack.yml` configuration will be
deprecated in Dotnet Publish Buildpack v1.0.0 & Dotnet Execute Buildpack
v1.0.0. To migrate from using `buildpack.yml`, please set the
`$BP_DOTNET_PROJECT_PATH` environment variable.

## Buildpack-Set Environment Variables

### DOTNET_ROOT

The `DOTNET_ROOT` environment variable specifies the path to the directory where .Net Runtimes and SDKs are installed.

* Set by: `dotnet-core-runtime`, `dotnet-core-sdk`, `dotnet-core-aspnet` buildpacks
* Phases: `build` and `launch`
* Value: path to the .Net root directory

### RUNTIME_VERSION

The `RUNTIME_VERSION` environment variable specifies the version of the .Net Core Runtime installed by the .Net Core Runtime Buildpack.

* Set by: `dotnet-core-runtime`
* Phases: `build`
* Value: installed version of the .Net Core Runtime

### SDK_LOCATION

The `SDK_LOCATION` environment variable specifies the file path location of the installed .Net Core SDK.

* Set by: `dotnet-core-sdk`
* Phases: `build`
* Value: path to the .Net Core SDK installation

### PATH

The `PATH` environment variable is modified to enable the `dotnet` CLI to be found during subsequent `build` and `launch` phases.

* Set by: `dotnet-core-sdk`
* Phases: `build` and `launch`
* Value: path the directory containing the `dotnet` executable

## Launch Process

The .Net Core Conf Buildpack will ensure that your application image is built
with a valid launch process command. These commands differ slightly depending
upon the type of built artifact produced during the build process.

For more information about which built artifact is produced for a Source
Application, see [this section]({{< relref "#application-types" >}}).

### Framework-Dependent Deployments and Source Applications

For Framework-Dependent Deployments (FDD), the `dotnet` CLI will be invoked to
start your application. The application will be given configuration to help it
bind to a port inside the container. The default port is 8080, but can be
overridden using the `$PORT` environment variable.

{{< code/copyable >}}
dotnet myapp.dll --urls http://0.0.0.0:${PORT:-8080}
{{< /code/copyable >}}

### Self-Contained Deployment and Framework-Dependent Executables

For Self-Contained Deployments and Framework-Dependent Executables, the
executable will be invoked directly to start your application. The application
will be given configuration to help it bind to a port inside the container. The
default port is 8080, but can be overridden using the `$PORT` environment
variable.

{{< code/copyable >}}
./myapp --urls http://0.0.0.0:${PORT:-8080}
{{< /code/copyable >}}

## Using CA Certificates
.Net Core Buildpack users can provide their own CA certificates and have them
included in the container root truststore at build-time and runtime by
following the instructions outlined in the [CA
Certificates](https://paketo.io/docs/buildpacks/configuration/#ca-certificates)
section of our configuration docs.

## Setting Custom Start Processes
.Net Core Buildpack users can set custom start processes for their app image by
following the instructions in the
[Procfiles](https://paketo.io/docs/buildpacks/configuration/#procfiles) section
of our configuration docs.

## Setting Environment Variables in the App Image
.Net Core Buildpack users can embed launch-time environment variables in their
app image by following the documentation for the [Environment Variables
Buildpack](https://github.com/paketo-buildpacks/environment-variables/blob/main/README.md).

## Adding Custom Labels to the App Image
.Net Core Buildpack users can add labels to their app image by following the
instructions in the [Applying Custom
Labels](https://paketo.io/docs/buildpacks/configuration/#applying-custom-labels)
section of our configuration docs.
