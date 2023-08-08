# README

Docat provides a fast reference to specific documents within a project.

## Features

- Listing of documents in a project
- Quick access to README and other standardized documents

## Get standard

Make sure Go is installed on your computer. If not, download and install Go from [this](https://go.dev/doc/install) page.

Enter the following command to download the latest version of docat.

```bash
go install github.com/aqyuki/docat@latest
```

### For advanced users

If you need a specific version of docat, please specify any version of docat in the following tags and download the docat of your choice.

```bash
go install github.com/aqyuki/docat@<tags>
```

See GitHub tags for docat tags.

## Usage

docat can be used as follows

```bash
docat <subcommands> [args] [options]
```

## Sub Commands (main)

1. `list`

   Displays specific documents in the current directory. This includes documents in sub folders.If you want to refer to a document in an arbitrary folder, specify the path of the target folder as the first argument.

2. `cat`

   Displays the contents of the specified document.

See `docat help` for details on all commands, and You can also check `docat <subcommand> --help` for more detailed options for each command.

## About the documents displayed

Documents currently recognized as documents are those with the following filenames.Note that the presence or absence of an extension does not affect the identification of a document.

- README
- LICENSE
- CHANGELOG
- CONTRIBUTING
- CONTRIBUTOR
