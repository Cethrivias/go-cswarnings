# C# Warnings parser and visualizer

A simple tool built to help me understand the state of one of the projects I have to work on right now

This is the first somewhat complete project in Go that I've ever done.

Most of the code is just copy-paste from echarts examples

# Usage

Generate a dotnet build output

```shell
dotnet build > build.log
```

pipe the result into the app

```shell
cat build.log | go run .
```

Or pipe build's output directly into the app.

It will generate **charts.html** file with a few useful graphs to visualize the state of your solution
