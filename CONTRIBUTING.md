# Contributing to NocoDB Go Client

Thank you for considering contributing to the NocoDB Go Client! This document
outlines the process for contributing to the project.

## How Can I Contribute?

### Reporting Bugs

If you find a bug in the code, please report it by creating an issue in the
GitHub repository. When filing a bug report, please include:

- A clear and descriptive title
- A detailed description of the issue
- Steps to reproduce the bug
- Expected behavior
- Actual behavior
- Any relevant logs or error messages
- Your environment (Go version, OS, etc.)

### Suggesting Enhancements

If you have an idea for a new feature or an enhancement to an existing feature,
please create an issue in the GitHub repository. When suggesting an enhancement,
please include:

- A clear and descriptive title
- A detailed description of the enhancement
- Why you think the enhancement would be useful
- Any relevant examples or use cases

### Pull Requests

If you want to contribute code to the project, please follow these steps:

1. Fork the repository
2. Create a new branch for your changes
3. Make your changes
4. Write tests for your changes
5. Run the tests to make sure they pass
6. Update the documentation if necessary
7. Submit a pull request

When submitting a pull request, please include:

- A clear and descriptive title
- A detailed description of the changes
- Any relevant issue numbers (e.g., "Fixes #123")

## Development Setup

To set up the project for development, follow these steps:

1. Clone the repository
2. Open the project in any code editor with support for Devcontainer
3. Run the tests to make sure everything is working

```bash
git clone https://github.com/eduardolat/nocodbgo.git
cd nocodbgo
code .
```

## What is Devcontainer?

Devcontainer is a feature of many modern code editors that allows you to create
a containerized development environment for your project. This is particularly
useful for development in environments that don't have Go installed.

More information can be found in the
[Devcontainer VS Code documentation](https://code.visualstudio.com/docs/devcontainers/containers)
and the [Devcontainers](https://containers.dev/) website.

## Coding Standards

Please follow these coding standards when contributing to the project:

- Write idiomatic Go code
- Write clear and concise comments
- Write tests for your code
- Use meaningful variable and function names
- Keep functions small and focused
- Use proper error handling

## License

By contributing to this project, you agree that your contributions will be
licensed under the project's [MIT License](LICENSE).
