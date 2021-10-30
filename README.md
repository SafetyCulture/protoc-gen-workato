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

```proto

```

Additional or default limits can be configured within the configuration file given to protoc-gen-workato.

The format for `key` in both the configuration and proto file is a pipe separated string of values for the workato descriptors. These map to the `descriptors` list supplied to the config base on their order. For example given a configuration of

```yaml
descriptors:
  - api_class
  - user_id
  - bucket # Must be the final descriptor
default_limits:
  # Rate limit applied to the `TasksComments` bucket
  - key: "public_api||TasksComments"
    value:
      unit: minute
      requests_per_unit: 20
```

`public_api||TasksComments` maps to `api_class=public_api,user_id:"",bucket:"TaskComments"`.

### Example Configuration

```yaml
descriptors:
  - api_class
  - user_id
  - bucket # Must be the final descriptor
default_limits:
  # All APIs have a limit of 800 RPM
  - key: ""
    value:
      unit: minute
      requests_per_unit: 800
  # Private APIs have no limits by default
  - key: "private_api"
    value:
      unlimited: true
  
  # This customer pays us lots of money, so we given them an increased workato
  - key: "|user_abc122"
    value:
      unlimited: true

  # Rate limit applied to the `TasksComments` bucket
  - key: "public_api||TasksComments"
    value:
      unit: minute
      requests_per_unit: 20
```

A complete example can be found in `protos/`.

## Development

This repo uses [buf](https://buf.build) to build Protocol Buffers.

To generate the image for fixtures run `buf build -o fixtures/image.bin`.  
To generate the annotations Go package run `buf generate`.
