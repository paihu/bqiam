# bqiam

[![Actions Status: golangci-lint](https://github.com/hirosassa/bqiam/workflows/golangci-lint/badge.svg)](https://github.com/hirosassa/bqiam/actions?query=workflow%3A"golangci-lint")
[![Apache-2.0](https://img.shields.io/github/license/hirosassa/bqiam)](LICENSE)

## What is this?

This tool provides easier permission management for BigQuery.

Currently supports;

- list the user's permissions for each BigQuery Datasets
- permit users to each BigQuery Datasets access role  

## Usage

Prepare configuration file as following format (currently support only the file name is `.bqiam.toml` on your `$HOME`):

```
// .bqiam.toml
BigqueryProjects = ["project-id-A", "project-id-B", ...]
CacheFile = "path/to/cache-file.toml"
```

Next, fetch bigquery dataset metadata and store it to cache file (take about 30-60 sec.).

```bash
$ bqiam cache
dataset meta data are cached to path/to/cache-file.toml
```

List datasets the user is able to access.
```bash
$ bqiam dataset abc@sample.com
sample-prj sample-ds1 OWNER
sample-prj sample-ds2 READER
...
```

Grant the user(s) a role to access the dataset(s).

```bash
$ bqiam permit READER -p bq-project-id -u user1@email.com -u user2@email.com -d dataset1 -d dataset2
Permit user1@email.com to dataset1 access as READER
Permit user2@email.com to dataset1 access as READER
...

```