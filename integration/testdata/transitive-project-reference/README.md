# Weather Forecast Web API

This app is based on the Web Api sample app that is provided by running `dotnet
new webapi`. A set of interdependent project files has been added. WebApi
depends on ProjectReferenceA, and ProjectReferenceA depends on
ProjectReferenceB.  This is a minimal app to test for the buggy behaviour
uncovered in https://github.com/paketo-buildpacks/dotnet-core/issues/670.
