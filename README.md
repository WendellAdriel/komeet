<div align="center">
    <h1>Komeet ☄️</h1>
</div>

**Komeet** is a Full-Stack Web App Template using **Go** for the API part and **Vue 3** for the UI. 

## Using the Template

There are two ways of using this template:

### GitHub Template

Click the `Use this template` button in the GitHub repository page.

### Git Clone

```bash
git clone git@github.com:WendellAdriel/komeet.git my-app && cd my-app && rm -rf .git
```

## Configuration

```bash
make configure
```

This will copy the configuration files from the sample ones and install the Front-end dependencies.
Make sure to update the configuration files with the needed values.

### Creating Users

**Komeet** ships with a `create-user` command for creating users to the application (great for dev envs):

```bash
make create-user NAME="John Doe" EMAIL=johndoe@example.com PASSWORD=secret
```

## Building the application

```bash
make build
```

## Running the application

```bash
make run
```

## Credits

- [Wendell Adriel](https://github.com/WendellAdriel)
- [All Contributors](../../contributors)

## Contributing

Check the **[Contributing Guide](CONTRIBUTING.md)**.