# Paketo .NET Core Documentation

The documentation that is served at https://paketo.io/docs/buildpacks/language-family-buildpacks/dotnet-core/.

To preview edits to this markdown:
```
git clone git@github.com:paketo-buildpacks/dotnet-core.git
git clone git@github.com:paketo-buildpacks/paketo-website.git
cd paketo-website
go mod edit -replace github.com/paketo-buildpacks/dotnet-core/docs=../dotnet-core/docs
hugo server
```
