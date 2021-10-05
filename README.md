# protoc-gen-workato

This is a Rate Limit generator for Google Protocol Buffers compiler `protoc`. The plugin generates a Lua filter to bucket requests based on their paths, and a descriptor file for [envoyproxy/workato](https://github.com/envoyproxy/workato).

## Installation

```bash
go get -u github.com/SafetyCulture/protoc-gen-workato/cmd/protoc-gen-workato
```

## Usage

The plugin is invoked by passing the --workato_out, and --workato_opt options to the protoc compiler. The option has the following format:

```bash
--doc_opt=workato/config.yaml
```

Annotations for rate limits can be applied at the service or method level. Here is an example of what that looks like:

```proto
service TasksService {
  option (s12.protobuf.workato.api_limit) = {
    limits: {
      key: "public_api",
      value: {
        unit: "minute"
        requests_per_unit: 100
      }
    }
    limits: {
      key: "private_api",
      value: {
        unit: "minute"
        requests_per_unit: 400
      }
    }
  };

  // CreateTask is used to create a new task.
  rpc CreateTask(CreateTaskRequest) returns (CreateTaskResponse) {
    option (s12.protobuf.workato.limit) = {
      limits: {
        key: "public_api",
        value: {
          unit: "minute"
          requests_per_unit: 10
        }
      }
      limits: {
        key: "private_api",
        value: {
          unit: "minute"
          requests_per_unit: 20
        }
      }
    };
  }

  // GetTask returns a task by id.
  rpc GetTask(GetTaskRequest) returns (GetTaskResponse) {}

  rpc AddComment(AddCommentRequest) returns (AddCommentResponse) {
    option (s12.protobuf.workato.limit) = {
      bucket: "TaskComments" // Custom bucket, so AddComment and UpdateComment can share a workato
    };
  }
  rpc UpdateComment(AddCommentRequest) returns (AddCommentResponse) {
    option (s12.protobuf.workato.limit) = {
      bucket: "TaskComments" // Custom bucket, so AddComment and UpdateComment can share a workato
    };
  }
}
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
