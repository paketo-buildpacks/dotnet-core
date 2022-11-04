# multi-framework=solution
This is a .NET solution that contains .NET 3.1 and .NET 6 dependent packages.
It's common for .NET developers to write modules that import other modules.
It's also common for them to upgrade their app _without_ upgrading the
frameworks used by its dependencies. This app tests that use-case.

This sample app is borrowed from @KieranJeffreySmart's
[exploration](https://github.com/KieranJeffreySmart/dotnet-execute/blob/0e010ff4c4afc16a391cf145b20dc80bf1b91c1c/integration/testdata/solution_7)
into .NET 7 support.

NOTE: This app can _only_ build in online mode. See discussion on [this
RFC](https://github.com/paketo-buildpacks/rfcs/pull/185) for more context on
how enabling this type of project in offline mode may be possible.
