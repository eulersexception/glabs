# glabs

## What is glabs?

glabs is a tool to manage repositories that maps the functionality provided by GitHub Classrooms on GitLab and its organizations.

## How to run glabs?

Clone this repository locally and run

```
go run main.go "tmp/test"
```

which will create a local database in the folder "tmp/test" in the project root and start the application.

## Unit tests

There are Unit tests to validate logic for creating and manipulating data. To execute tests follow these steps.

1. Change to model folder in project

```
cd /path/to/glabs/model
```

2. Run tests

```
go test
```

Yet the application is not fully functional and so is the test coverage.
