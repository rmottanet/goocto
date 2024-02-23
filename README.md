# GoOcto

## ctrl+s :v:

Welcome to GoOcto, a CLI (Command Line Interface) developed in Go with the purpose of simplifying interactions with GitHub directly from the command line. Currently, GoOcto offers basic functionalities to create, edit, and delete repositories on GitHub. This project is under constant development, with plans to expand its features in the future.


## Installation and Requirements

Before you start using **GoOcto**, follow these steps:

1. **GitHub API Token**: You need to create a personal access token in your GitHub account. Access your settings at Settings > Developer settings > Personal access tokens. Make sure to grant appropriate permissions for the functionalities you intend to use.

2. **Environment File Configuration**: Rename the `.env.example` file in the repository to `.env`. Insert the GitHub API token into the `.env` file as an environment variable.

Once you meet the above requirements, you can proceed with the installation of **GoOcto**. Follow these steps:

1. Clone the **GoOcto** repository to your local machine:

```bash
git clone https://github.com/rmottanet/goocto.git
```

2. Navigate to the **GoOcto** directory:

```bash
cd goocto/cli/cmd
```

3. Compile the source code of **GoOcto**:

```bash
go build -o ../../goocto
```

4. After successful compilation, you will have the executable binary of **GoOcto** ready for use.

With **GoOcto** installed and configured, you can start interacting with your GitHub repositories directly from the command line.


## Basic Usage

After successful installation, you can use **GoOcto** to perform the following operations:

- **Create a new repository:**
  ```bash
  goocto new <repository-name>
  ```

- **Delete a repository:**
  ```bash
  goocto del <owner> <repository-name>
  ```

For more details on each command, you can use the `--help` option. For example:

```bash
goocto new --help
```

## Contribution

Contributions to the GoOcto project are welcome! If you have ideas for improvements, feature requests, or bug reports, feel free to open an issue or submit a pull request.

---

Thank you for considering **GoOcto** for your needs. While it may be a modest tool compared to others, I hope it proves useful in simplifying your interactions with GitHub. If you have any feedback or suggestions, don't hesitate to reach out. Happy coding!

<br />
<br />
<p align="center">
<a href="https://gitlab.com/rmottanet"><img src="https://img.shields.io/badge/Gitlab--_.svg?style=social&logo=gitlab" alt="GitLab"></a>
<a href="https://github.com/rmottanet"><img src="https://img.shields.io/badge/Github--_.svg?style=social&logo=github" alt="GitHub"></a>
<a href="https://instagram.com/rmottanet/"><img src="https://img.shields.io/badge/Instagram--_.svg?style=social&logo=instagram" alt="Instagram"></a>
<a href="https://www.linkedin.com/in/rmottanet/"><img src="https://img.shields.io/badge/Linkedin--_.svg?style=social&logo=linkedin" alt="Linkedin"></a>
</p>
<br />
