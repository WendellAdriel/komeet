<div align="center">
    <h1>Komeet ☄️</h1>
    <p><strong>API Template built with Go and Gin</strong></p>
    <p><img alt="Static Badge" src="https://img.shields.io/badge/IN%20DEVELOPMENT-FFA200"></p>
</div>

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

This will copy the configuration files from the sample ones.
Make sure to update the configuration files with the needed values.

## Building the application

```bash
make build
```

## Running the application

```bash
make run
```

## Creating Users

**Komeet** ships with a `create-user` command for creating users to the application (great for dev envs):

```bash
make create-user NAME="John Doe" EMAIL=johndoe@example.com PASSWORD=secret
```

## Setting a CI/CD Pipeline

For non-local envs, you'll need only the content from the `dist` folder to be added to your server/container.
These are the steps for setting a CI/CD pipeline for **Komeet** applications:

1 - Run the `make build` command to generate the `dist` folder.

2 - Place the contents of the `dist` folder in your server/container.

3 - Add the `config.json` and `secrets.json` files to the same folder from point 2.

4 - Run the `./komeet serve` command. 

## Credits

- [Wendell Adriel](https://github.com/WendellAdriel)
- [All Contributors](../../contributors)

## Contributing

Check the **[Contributing Guide](CONTRIBUTING.md)**.