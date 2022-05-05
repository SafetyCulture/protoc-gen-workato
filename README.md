# protoc-gen-workato

This is a [Workato Connector](https://docs.workato.com/developing-connectors/sdk.html) generator for Google Protocol Buffers compiler `protoc`. The plugin generates a Connector file based on your publicly tagged methods.

## Installation

```bash
go install github.com/SafetyCulture/protoc-gen-workato@latest
```

## Usage

The plugin is invoked by passing the --workato_out, and --workato_opt options to the protoc compiler. The option has the following format:

```bash
--doc_opt=workato/config.yaml
```

Generation of the actions in workato relies on the usage of [`google.api.http`](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L46) and [`grpc.gateway.protoc_gen_openapiv2.options`](https://github.com/grpc-ecosystem/grpc-gateway/blob/master/protoc-gen-openapiv2/options/annotations.proto) annotations.

A complete example can be found in `proto/`.

## Development

This repo uses [buf](https://buf.build) to build Protocol Buffers.
```bash
brew tap bufbuild/buf
brew install buf
```
To generate the image for fixtures run `buf build -o fixtures/image.bin`.
To generate the annotations Go package run `buf generate`.
Example `buf generate s12/protobuf/workato//annotations.proto`.

## Testing
```bash
go test -v -run TestGenerateWorkatoConnector
```

After you are happy with the results, to update the snapshot run this command:
```bash
UPDATE_SNAPSHOTS=true go test -v -run TestGenerateWorkatoConnector
```
