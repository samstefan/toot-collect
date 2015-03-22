## Toot collect

A small go application that polls for tweets from specified accounts and exposes them via a rest API.

Requirements:
  - `mongoDB`
  - `go`
  - `make`

## Getting started

Copy the `properties-example.json` file to `properties,json` and fill out appropriately.

- Build the project with `$ make build`
- Run the app with `$ make run`

## End points

Get users latests tweet: `/tweets/:screenName`