# protoc-gen-flat

`protoc-gen-flat` is a plugin for the Protocol Buffers compiler (`protoc`) that generates flat golang structs and constants code from `.proto` files.


## Why?
Most of the time, I develop backend services with gRPC and need to store those messages in memory. However, storing generated Go structs has more overhead compared to flat Go structs and constants because Protobuf was not designed for this purpose. Therefore, I need to write corresponding structs for proto messages manually. This tool automates that process for me.


## Installation

To install `protoc-gen-flat`, use the following command:

```sh
go install github.com/selam/protoc-gen-flat@latest
```

Make sure your `GOPATH` is set and `$GOPATH/bin` is added to your `PATH`.

## Usage

To generate flatbuffers code from your `.proto` files, run the following command:

```sh
protoc --flat_out=. yourfile.proto
```

## Example

Given a `example.proto` file:

```proto
syntax = "proto3";

message Example {
    string name = 1;
    int32 id = 2;
}
```

Run the following command to generate the flatbuffers code:

```sh
protoc --flat_out=. --flag_opt=suffix=_flat example.proto
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or feedback, please contact [yourname@example.com](mailto:yourname@example.com).
