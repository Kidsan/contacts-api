# Contacts API

This is a basic REST API that manages contact resources.
The design of this API implements some domain driven design techniques to manage different resources and facilitate code cleanliness and resusability.
This application serves no purpose other than for me to practice some implementations and patterns.


# Interfaces

The api currently exposes either a REST or a GRPC interface. This can be configured via the `APPLICATION_LISTENER_TYPE` environment variable. The default is REST.

# TODO

- GraphQL interface in the http package
    - maybe refactor http package to separate the REST, GRPC, and GraphQL implementations
- Terrafrom provider for contacts api
- Pipelines
- OpenAPI doc
- More work on metrics
