# Web CLI

Associate websites to directory paths to easily open them from the CLI

## Usage

Add a webiste:

```
web add home http://www.example.com
```

List the sites for the current directory:

```
web list
```

Open a site:

```
web home
```

Remove a site:

```
web remove home
```

If in a git repository `web git` will attempt to parse and visit the homepage.

## Installation

Clone and:

```
make
make install
```
