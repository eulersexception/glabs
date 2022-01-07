# glabs

## What is glabs?

glabs is a tool to manage repositories that maps the functionality provided by GitHub Classrooms on GitLab and its organizations.

## How to run glabs?

Clone this repository locally and run

```
go run main.go
```

This will create a local database named "default_db" in the project root and start the application. The database is initially filled with dummy data which will be dropped everytime the application is closed. To work with own data some changes must be applied to `main.go`. Currently Teams and Students can not be created by view elements and data input via file upload is not supported.

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
