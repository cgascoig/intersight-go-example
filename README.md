# Intersight Go SDK Example

This repo contains an extremely simple example of how to use the Intersight Go SDK. It uses the unofficial SDK version published at [github.com/cgascoig/intersight-go-sdk](https://github.com/cgascoig/intersight-go-sdk) as it should be easier to import without needing to be vendored in. 

# Usage

## Building

```
go build -o "intersight-go-example" ./
```

Note: The SDK is large and the first build can take a while (e.g. 13mins on a 2018 MacBook Pro). Subsequent builds should be quicker.

If you have GNU make 4.0 or later, you can also use the Makefile by running `make` or `gmake`

## Running the example

To run, first set some environment variables for your Intersight API Key ID and the path to your Intersight API Key file:

```
export IS_KEY_ID="123456789012345678901234/123456789012345678901234/123456789012345678901234"
export IS_KEY_FILE="/path/to/intersight-api-key.pem"
```

Finally, run the example:

```
./intersight-go-example
```