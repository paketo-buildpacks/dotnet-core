# Generating this fixture

1. Install .NET SDK (6.0.301)
1. `dotnet publish $(pwd)/../source-app --configuration Release --runtime ubuntu.18.04-x64 --self-contained true --output $(pwd)`

