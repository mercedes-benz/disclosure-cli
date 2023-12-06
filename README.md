<!--
SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH

SPDX-License-Identifier: MIT
-->

# Disclosure-CLI

## Introduction
The Disclosure-CLI provides an easy way to access the public api of the FOSS Disclosure Portal.
It is the recommended tool for external suppliers who do not have access to the Disclosure Portal and need access to the project data.

With the Disclosure-CLI external suppliers can:
- Create new project versions
- Access policy rules
- Upload SBOM files
- Set a reference to the corresponding source code
- Get general information about the project

## Table of Contents
* [Download and build](#download-and-build)
* [How-To](#how-to)
* [Guided example](#guided-example)
* [Contributing](#contributing)
* [Code of Conduct](#code-of-conduct)
* [License](#license)
* [Provider Information](#provider-information)

## Download and build
We distribute the Disclosure-CLI as executables, container image and github action.

### Disclosure-CLI executables
You can download the binaries for the Disclosure-CLI in the GitHub releases page. 

You can also download the source code and build it yourself. 
```
git clone https://github.com/mercedes-benz/disclosure-cli.git
cd disclosure-cli
go build -o disclosure-cli
```

### Disclosure-CLI as container image
You can pull a Disclosure-CLI image from https://github.com/orgs/mercedes-benz/packages or build it yourself with the Dockerfile provided in the repository. 

```
docker pull ghcr.io/mercedes-benz/disclosure-cli:0.96.4-amd64
```
or
```
git clone https://github.com/mercedes-benz/disclosure-cli.git
cd disclosure-cli
docker build . -t disclosure-cli
```

Run the image to get project information
```
docker run disclosure-cli project details -H HOST -u PROJECT_UUID -t TOKEN
```

Run the image to upload a sbom file to a project version
```
docker run -v $(pwd)/sbom.spdx.json:/sbom.spdx.json disclosure-cli version sbomUpload sbom.spdx.json -H HOST -u PROJECT_UUID  -t TOKEN -v VERSION
```

### Disclosure-CLI as GitHub Action
The definition of the action can be found in the file action.yml in the repository. A working example ([sbom-upload.yml](./.github/workflows/sbom-upload.yml)) can be found in the .github/workflows folder. 


## How-To

The recommended way to use the Disclosure-CLI as executable is with a config file, but you can also set environment variables or use flags with your commands instead. 
The config file needs to be in the same folder as the Disclosure-CLI, needs to be named `config.yml` and must have the following structure:

``` yml
projecttoken: "project-token"
projectuuid: "project-uuid"
projectversion: "1.0"
host: "PUBLIC API"
```

The environment variables mus be named as follows:
``` 
INPUT_TOKEN
INPUT_PROJECT_UUID
INPUT_PROJECT_VERSION
INPUT_HOST
```

### How to use the Disclosure-CLI

```
./disclosure-cli

A Disclosure Portal client for the disclosure public api.
Manage your projects with this client.

Usage:
  disclosure-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  project     Execute a project command
  sbom        Execute a project version sbom command
  sha256      Generates a sha256 hash
  version     Execute a project version command

Flags:
  -c, --configFile string   location of the config file (default "./config.yml")
  -h, --help                help for disclosure-cli
  -t, --token string        disclosure-cli project token
  -u, --uuid string         uuid of the project
  -v, --version string      version of the project (default "1.0")

Use "disclosure-cli [command] --help" for more information about a command.
```

### Sample commands
```
// Retrieving project detail
./disclosure-cli project details -H HOST -u PROJECT_UUID -t TOKEN

// Upload sbom to a project version
./disclosure-cli version sbomUpload FILENAME -H HOST -u PROJECT_UUID -t TOKEN -v VERSION
```

### Available Commands

```
project details       Returning the project details
project policyrules   Returning the project policy rules
project sbomCheck     On demand check for SBOM files
project schema        Returning the project schema
project status        Returning the project status

version ccs           CCS status
version ccsAdd        Add reference (url) to complete corresponding source code
version create        Create version
version details       Returning the project version details
version list          Returning the project version list
version sbomDetails   Details of SBOM
version sbomNotice    Get third party notice information for a SBOM as html / json / text 
version sbomStatus    Status information of SBOM
version sbomUpload    Uploads SBOM file to a project version
version sboms         List of all uploaded SBOMS

sbom tag              Add tag to a sbom

```

Note 12-06-2023: 
2024 we will review the disclosure-cli to restructure the command structure and add new features. 
With the current release we establish a new "sbom" command with a new sub command to add tags to a sbom.
Existing sbom commands (sbomDetails, sbomNotice, sbomStatus, sbomUpload, sboms) will move to the new "sbom" command in 2024.  


### Help on commands
```
version sbomUpload -h
```


## Guided example
The next few steps will guide you through a Disclosure Portal project with the Disclosure-CLI.
In this guided example we will use a config file to access our project, but you can also use flags or environment variables. 

### Step 0 - Preparation
For the next few steps you need a project in the Disclosure Portal, and the unique identifier and token of your project.  
You can skip to the Disclosure-CLI subsection, if you already have this data.

**Disclosure Portal**
*Only project owners in the Disclosure Portal have the permissions to do these steps.*
a) Create a new Project in the Disclosure Portal. Take note of the Unique identifier of the project, you will need it in a moment.  
b) Create a token. You will need this token to access your project with the Disclosure-CLI.

**Disclosure-CLI**  
a) Create a new folder wherever you want. I will call this folder `disclosure-cli`, but you can name it differently.  
b) Download the source code for the Disclosure-CLI client and build it.
```
go build -o disclosure-cli
```
c) Move `disclosure-cli` into your created folder (a)  
d) Create a new config file named `config.yml` with the following attributes and values in the same folder
``` yml
projecttoken: "project-token"
projectuuid: "project-uuid"
projectversion: "1.0"
host: "" // host needs to be changed to public api
```
If you have just created a new project, your project on the Disclosure Portal won't have a project version yet.
We will create a version `1.0` in the following steps.

### Step 1 - Your first command: Get the project details

```
./disclosure-cli project details
```
```
{
    "name": "disclosure-cli-example",
    "uuid": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeee",
    "created": "2022-03-16T14:54:13.306719888Z",
    "updated": "2022-03-16T15:13:45.461701011Z",
    "schema": "EnterpriseIT",
    "description": "This is the example project for the cli client."
}
```

### Step 2 - Create a project version
Our project still does not have a project version. Let's create one.
```
./disclosure-cli version create "1.0" "First iteration of the disclosure-cli example"
```
```
{
    "success": true,
    "message": "version created",
    "key": "ab02a5b5-8923-41pb-9a9b-cd5836d66dbe"
}
```
Let's have a look at the project versions of our project. There should only be one ;)
```
./disclosure-cli version list
```
```
[
    "1.0"
]
```

### Step 3 - Upload the SBOM of your project
Next we will upload the SBOM and look at the metadata of the SBOM.
The SBOM is always related to a specific project version. In our case it is `1.0` as described in our `config.yml`.

```
./disclosure-cli version sbomUpload ./disclosure-cli/sbom.json
```
```
{
    "docIsValid": true,
    "validationFailedMessage": "",
    "hash": "77a50e85477db7580d3052403a8d1ebe4ad9a3b912ad18e3a3b4f0ccecf36c65",
    "fileUploaded": true,
    "id": "SPDXRef-DOCUMENT"
}
```
```
./disclosure-cli version sbomDetails
```
```
{
    "name": "Disclosure-CLI,
    "id": "SPDXRef-DOCUMENT",
    "version": "SPDX-2.2",
    "creators": "Tool: xxx",
    "created": "2022-03-16T14:54:13.306719888Z",
    "updated": "2022-03-16T15:13:45.461701011Z",
    "status": true
}
```

## Contributing

We welcome any contributions.
If you want to contribute to this project, please read the [contributing guide](CONTRIBUTING.md).

## Code of Conduct

Please read our [Code of Conduct](https://github.com/mercedes-benz/foss/blob/master/CODE_OF_CONDUCT.md) as it is our base for interaction.

## License

This project is licensed under the [MIT LICENSE](LICENSE).

## Provider Information

Please visit <https://github.com/mercedes-benz/foss/blob/master/PROVIDER_INFORMATION.md> for information on the provider.

Notice: Before you use the program in productive use, please take all necessary precautions,
e.g. testing and verifying the program with regard to your specific use.
The program was tested solely for our own use cases, which might differ from yours.
