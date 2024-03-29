// Generated by swift-openapi-generator, do not modify.
@_spi(Generated) import OpenAPIRuntime
#if os(Linux)
  @preconcurrency import struct Foundation.Data
  @preconcurrency import struct Foundation.Date
  @preconcurrency import struct Foundation.URL
#else
  import struct Foundation.Data
  import struct Foundation.Date
  import struct Foundation.URL
#endif
import HTTPTypes

/// Used by clients to notify the GitHubLightServer of the current status for PRs requiring reviews or merge, and local branches requiring pulls
struct Client: APIProtocol {
  /// The underlying HTTP client.
  private let client: UniversalClient
  /// Creates a new client.
  /// - Parameters:
  ///   - serverURL: The server URL that the client connects to. Any server
  ///   URLs defined in the OpenAPI document are available as static methods
  ///   on the ``Servers`` type.
  ///   - configuration: A set of configuration values for the client.
  ///   - transport: A transport that performs HTTP operations.
  ///   - middlewares: A list of middlewares to call before the transport.
  init(
    serverURL: Foundation.URL,
    configuration: Configuration = .init(),
    transport: any ClientTransport,
    middlewares: [any ClientMiddleware] = [])
  {
    client = .init(
      serverURL: serverURL,
      configuration: configuration,
      transport: transport,
      middlewares: middlewares)
  }

  private var converter: Converter {
    client.converter
  }

  /// Returns current lights for software implementations
  ///
  /// - Remark: HTTP `GET /lights`.
  /// - Remark: Generated from `#/paths//lights/get`.
  func get_sol_lights(_ input: Operations.get_sol_lights.Input) async throws -> Operations.get_sol_lights.Output {
    try await client.send(
      input: input,
      forOperation: Operations.get_sol_lights.id,
      serializer: { input in
        let path = try converter.renderedPath(
          template: "/lights",
          parameters: [])
        var request: HTTPTypes.HTTPRequest = .init(
          soar_path: path,
          method: .get)
        suppressMutabilityWarning(&request)
        converter.setAcceptHeader(
          in: &request.headerFields,
          contentTypes: input.headers.accept)
        return (request, nil)
      },
      deserializer: { response, responseBody in
        switch response.status.code {
          case 200:
            let contentType = converter.extractContentTypeIfPresent(in: response.headerFields)
            let body: Operations.get_sol_lights.Output.Ok.Body
            let chosenContentType = try converter.bestContentType(
              received: contentType,
              options: [
                "application/json",
              ])
            switch chosenContentType {
              case "application/json":
                body = try await converter.getResponseBodyAsJSON(
                  Components.Schemas.LightColor.self,
                  from: responseBody,
                  transforming: { value in
                    .json(value)
                  })
              default:
                preconditionFailure("bestContentType chose an invalid content type.")
            }
            return .ok(.init(body: body))
          default:
            let contentType = converter.extractContentTypeIfPresent(in: response.headerFields)
            let body: Operations.get_sol_lights.Output.Default.Body
            let chosenContentType = try converter.bestContentType(
              received: contentType,
              options: [
                "application/json",
              ])
            switch chosenContentType {
              case "application/json":
                body = try await converter.getResponseBodyAsJSON(
                  Components.Schemas._Error.self,
                  from: responseBody,
                  transforming: { value in
                    .json(value)
                  })
              default:
                preconditionFailure("bestContentType chose an invalid content type.")
            }
            return .default(
              statusCode: response.status.code,
              .init(body: body))
        }
      })
  }

  /// Returns current status
  ///
  /// - Remark: HTTP `GET /status`.
  /// - Remark: Generated from `#/paths//status/get`.
  func get_sol_status(_ input: Operations.get_sol_status.Input) async throws -> Operations.get_sol_status.Output {
    try await client.send(
      input: input,
      forOperation: Operations.get_sol_status.id,
      serializer: { input in
        let path = try converter.renderedPath(
          template: "/status",
          parameters: [])
        var request: HTTPTypes.HTTPRequest = .init(
          soar_path: path,
          method: .get)
        suppressMutabilityWarning(&request)
        converter.setAcceptHeader(
          in: &request.headerFields,
          contentTypes: input.headers.accept)
        return (request, nil)
      },
      deserializer: { response, responseBody in
        switch response.status.code {
          case 200:
            let contentType = converter.extractContentTypeIfPresent(in: response.headerFields)
            let body: Operations.get_sol_status.Output.Ok.Body
            let chosenContentType = try converter.bestContentType(
              received: contentType,
              options: [
                "application/json",
              ])
            switch chosenContentType {
              case "application/json":
                body = try await converter.getResponseBodyAsJSON(
                  Components.Schemas.Status.self,
                  from: responseBody,
                  transforming: { value in
                    .json(value)
                  })
              default:
                preconditionFailure("bestContentType chose an invalid content type.")
            }
            return .ok(.init(body: body))
          default:
            let contentType = converter.extractContentTypeIfPresent(in: response.headerFields)
            let body: Operations.get_sol_status.Output.Default.Body
            let chosenContentType = try converter.bestContentType(
              received: contentType,
              options: [
                "application/json",
              ])
            switch chosenContentType {
              case "application/json":
                body = try await converter.getResponseBodyAsJSON(
                  Components.Schemas._Error.self,
                  from: responseBody,
                  transforming: { value in
                    .json(value)
                  })
              default:
                preconditionFailure("bestContentType chose an invalid content type.")
            }
            return .default(
              statusCode: response.status.code,
              .init(body: body))
        }
      })
  }

  /// Resets the system
  ///
  /// - Remark: HTTP `GET /reset`.
  /// - Remark: Generated from `#/paths//reset/get`.
  func get_sol_reset(_ input: Operations.get_sol_reset.Input) async throws -> Operations.get_sol_reset.Output {
    try await client.send(
      input: input,
      forOperation: Operations.get_sol_reset.id,
      serializer: { input in
        let path = try converter.renderedPath(
          template: "/reset",
          parameters: [])
        var request: HTTPTypes.HTTPRequest = .init(
          soar_path: path,
          method: .get)
        suppressMutabilityWarning(&request)
        converter.setAcceptHeader(
          in: &request.headerFields,
          contentTypes: input.headers.accept)
        return (request, nil)
      },
      deserializer: { response, responseBody in
        switch response.status.code {
          case 200:
            let contentType = converter.extractContentTypeIfPresent(in: response.headerFields)
            let body: Operations.get_sol_reset.Output.Ok.Body
            let chosenContentType = try converter.bestContentType(
              received: contentType,
              options: [
                "application/json",
              ])
            switch chosenContentType {
              case "application/json":
                body = try await converter.getResponseBodyAsJSON(
                  Components.Schemas.Result.self,
                  from: responseBody,
                  transforming: { value in
                    .json(value)
                  })
              default:
                preconditionFailure("bestContentType chose an invalid content type.")
            }
            return .ok(.init(body: body))
          default:
            let contentType = converter.extractContentTypeIfPresent(in: response.headerFields)
            let body: Operations.get_sol_reset.Output.Default.Body
            let chosenContentType = try converter.bestContentType(
              received: contentType,
              options: [
                "application/json",
              ])
            switch chosenContentType {
              case "application/json":
                body = try await converter.getResponseBodyAsJSON(
                  Components.Schemas._Error.self,
                  from: responseBody,
                  transforming: { value in
                    .json(value)
                  })
              default:
                preconditionFailure("bestContentType chose an invalid content type.")
            }
            return .default(
              statusCode: response.status.code,
              .init(body: body))
        }
      })
  }

  /// Reports the current status for monitored tuple
  ///
  /// - Remark: HTTP `POST /report`.
  /// - Remark: Generated from `#/paths//report/post`.
  func post_sol_report(_ input: Operations.post_sol_report.Input) async throws -> Operations.post_sol_report.Output {
    try await client.send(
      input: input,
      forOperation: Operations.post_sol_report.id,
      serializer: { input in
        let path = try converter.renderedPath(
          template: "/report",
          parameters: [])
        var request: HTTPTypes.HTTPRequest = .init(
          soar_path: path,
          method: .post)
        suppressMutabilityWarning(&request)
        converter.setAcceptHeader(
          in: &request.headerFields,
          contentTypes: input.headers.accept)
        let body: OpenAPIRuntime.HTTPBody? = switch input.body {
          case let .json(value):
            try converter.setRequiredRequestBodyAsJSON(
              value,
              headerFields: &request.headerFields,
              contentType: "application/json; charset=utf-8")
        }
        return (request, body)
      },
      deserializer: { response, responseBody in
        switch response.status.code {
          case 200:
            let contentType = converter.extractContentTypeIfPresent(in: response.headerFields)
            let body: Operations.post_sol_report.Output.Ok.Body
            let chosenContentType = try converter.bestContentType(
              received: contentType,
              options: [
                "application/json",
              ])
            switch chosenContentType {
              case "application/json":
                body = try await converter.getResponseBodyAsJSON(
                  Components.Schemas.Result.self,
                  from: responseBody,
                  transforming: { value in
                    .json(value)
                  })
              default:
                preconditionFailure("bestContentType chose an invalid content type.")
            }
            return .ok(.init(body: body))
          default:
            let contentType = converter.extractContentTypeIfPresent(in: response.headerFields)
            let body: Operations.post_sol_report.Output.Default.Body
            let chosenContentType = try converter.bestContentType(
              received: contentType,
              options: [
                "application/json",
              ])
            switch chosenContentType {
              case "application/json":
                body = try await converter.getResponseBodyAsJSON(
                  Components.Schemas._Error.self,
                  from: responseBody,
                  transforming: { value in
                    .json(value)
                  })
              default:
                preconditionFailure("bestContentType chose an invalid content type.")
            }
            return .default(
              statusCode: response.status.code,
              .init(body: body))
        }
      })
  }
}
