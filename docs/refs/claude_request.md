# Messages


Send a structured list of input messages with text and/or image content, and the model will generate the next message in the conversation.

The Messages API can be used for either single queries or stateless multi-turn conversations.

Learn more about the Messages API in our user guide

`POST /v1/messages`

```sh
curl https://api.anthropic.com/v1/messages \
     --header "x-api-key: $ANTHROPIC_API_KEY" \
     --header "anthropic-version: 2023-06-01" \
     --header "content-type: application/json" \
     --data \
'{
    "model": "claude-sonnet-4-20250514",
    "max_tokens": 1024,
    "messages": [
        {"role": "user", "content": "Hello, world"}
    ]
}'
```

response:

```json
{
  "content": [
    {
      "text": "Hi! My name is Claude.",
      "type": "text"
    }
  ],
  "id": "msg_013Zva2CMHLNnXjNJJKqJ2EF",
  "model": "claude-sonnet-4-20250514",
  "role": "assistant",
  "stop_reason": "end_turn",
  "stop_sequence": null,
  "type": "message",
  "usage": {
    "input_tokens": 2095,
    "output_tokens": 503
  }
}
```

## OpenAPI

````yaml post /v1/messages
paths:
  path: /v1/messages
  method: post
  servers:
    - url: https://api.anthropic.com
  request:
    security: []
    parameters:
      path: {}
      query: {}
      header:
        anthropic-beta:
          schema:
            - type: array
              items:
                allOf:
                  - type: string
              required: false
              title: Anthropic-Beta
              description: >-
                Optional header to specify the beta version(s) you want to use.


                To use multiple betas, use a comma separated list like
                `beta1,beta2` or specify the header multiple times for each
                beta.
        anthropic-version:
          schema:
            - type: string
              required: true
              title: Anthropic-Version
              description: >-
                The version of the Anthropic API you want to use.


                Read more about versioning and our version history
                [here](https://docs.anthropic.com/en/api/versioning).
        x-api-key:
          schema:
            - type: string
              required: true
              title: X-Api-Key
              description: >-
                Your unique API key for authentication.


                This key is required in the header of all API requests, to
                authenticate your account and access Anthropic's services. Get
                your API key through the
                [Console](https://console.anthropic.com/settings/keys). Each key
                is scoped to a Workspace.
      cookie: {}
    body:
      application/json:
        schemaArray:
          - type: object
            properties:
              model:
                allOf:
                  - description: >-
                      The model that will complete your prompt.


                      See
                      [models](https://docs.anthropic.com/en/docs/models-overview)
                      for additional details and options.
                    examples:
                      - claude-sonnet-4-20250514
                    maxLength: 256
                    minLength: 1
                    title: Model
                    type: string
              messages:
                allOf:
                  - description: >-
                      Input messages.


                      Our models are trained to operate on alternating `user`
                      and `assistant` conversational turns. When creating a new
                      `Message`, you specify the prior conversational turns with
                      the `messages` parameter, and the model then generates the
                      next `Message` in the conversation. Consecutive `user` or
                      `assistant` turns in your request will be combined into a
                      single turn.


                      Each input message must be an object with a `role` and
                      `content`. You can specify a single `user`-role message,
                      or you can include multiple `user` and `assistant`
                      messages.


                      If the final message uses the `assistant` role, the
                      response content will continue immediately from the
                      content in that message. This can be used to constrain
                      part of the model's response.


                      Example with a single `user` message:


                      ```json

                      [{"role": "user", "content": "Hello, Claude"}]

                      ```


                      Example with multiple conversational turns:


                      ```json

                      [
                        {"role": "user", "content": "Hello there."},
                        {"role": "assistant", "content": "Hi, I'm Claude. How can I help you?"},
                        {"role": "user", "content": "Can you explain LLMs in plain English?"},
                      ]

                      ```


                      Example with a partially-filled response from Claude:


                      ```json

                      [
                        {"role": "user", "content": "What's the Greek name for Sun? (A) Sol (B) Helios (C) Sun"},
                        {"role": "assistant", "content": "The best answer is ("},
                      ]

                      ```


                      Each input message `content` may be either a single
                      `string` or an array of content blocks, where each block
                      has a specific `type`. Using a `string` for `content` is
                      shorthand for an array of one content block of type
                      `"text"`. The following input messages are equivalent:


                      ```json

                      {"role": "user", "content": "Hello, Claude"}

                      ```


                      ```json

                      {"role": "user", "content": [{"type": "text", "text":
                      "Hello, Claude"}]}

                      ```


                      Starting with Claude 3 models, you can also send image
                      content blocks:


                      ```json

                      {"role": "user", "content": [
                        {
                          "type": "image",
                          "source": {
                            "type": "base64",
                            "media_type": "image/jpeg",
                            "data": "/9j/4AAQSkZJRg...",
                          }
                        },
                        {"type": "text", "text": "What is in this image?"}
                      ]}

                      ```


                      We currently support the `base64` source type for images,
                      and the `image/jpeg`, `image/png`, `image/gif`, and
                      `image/webp` media types.


                      See
                      [examples](https://docs.anthropic.com/en/api/messages-examples#vision)
                      for more input examples.


                      Note that if you want to include a [system
                      prompt](https://docs.anthropic.com/en/docs/system-prompts),
                      you can use the top-level `system` parameter — there is no
                      `"system"` role for input messages in the Messages API.


                      There is a limit of 100,000 messages in a single request.
                    items:
                      $ref: '#/components/schemas/InputMessage'
                    title: Messages
                    type: array
              container:
                allOf:
                  - anyOf:
                      - type: string
                      - type: 'null'
                    description: Container identifier for reuse across requests.
                    title: Container
              max_tokens:
                allOf:
                  - description: >-
                      The maximum number of tokens to generate before stopping.


                      Note that our models may stop _before_ reaching this
                      maximum. This parameter only specifies the absolute
                      maximum number of tokens to generate.


                      Different models have different maximum values for this
                      parameter.  See
                      [models](https://docs.anthropic.com/en/docs/models-overview)
                      for details.
                    examples:
                      - 1024
                    minimum: 1
                    title: Max Tokens
                    type: integer
              mcp_servers:
                allOf:
                  - description: MCP servers to be utilized in this request
                    items:
                      $ref: '#/components/schemas/RequestMCPServerURLDefinition'
                    maxItems: 20
                    title: Mcp Servers
                    type: array
              metadata:
                allOf:
                  - $ref: '#/components/schemas/Metadata'
                    description: An object describing metadata about the request.
              service_tier:
                allOf:
                  - description: >-
                      Determines whether to use priority capacity (if available)
                      or standard capacity for this request.


                      Anthropic offers different levels of service for your API
                      requests. See
                      [service-tiers](https://docs.anthropic.com/en/api/service-tiers)
                      for details.
                    enum:
                      - auto
                      - standard_only
                    title: Service Tier
                    type: string
              stop_sequences:
                allOf:
                  - description: >-
                      Custom text sequences that will cause the model to stop
                      generating.


                      Our models will normally stop when they have naturally
                      completed their turn, which will result in a response
                      `stop_reason` of `"end_turn"`.


                      If you want the model to stop generating when it
                      encounters custom strings of text, you can use the
                      `stop_sequences` parameter. If the model encounters one of
                      the custom sequences, the response `stop_reason` value
                      will be `"stop_sequence"` and the response `stop_sequence`
                      value will contain the matched stop sequence.
                    items:
                      type: string
                    title: Stop Sequences
                    type: array
              stream:
                allOf:
                  - description: >-
                      Whether to incrementally stream the response using
                      server-sent events.


                      See
                      [streaming](https://docs.anthropic.com/en/api/messages-streaming)
                      for details.
                    title: Stream
                    type: boolean
              system:
                allOf:
                  - anyOf:
                      - type: string
                      - items:
                          $ref: '#/components/schemas/RequestTextBlock'
                        type: array
                    description: >-
                      System prompt.


                      A system prompt is a way of providing context and
                      instructions to Claude, such as specifying a particular
                      goal or role. See our [guide to system
                      prompts](https://docs.anthropic.com/en/docs/system-prompts).
                    examples:
                      - - text: Today's date is 2024-06-01.
                          type: text
                      - Today's date is 2023-01-01.
                    title: System
              temperature:
                allOf:
                  - description: >-
                      Amount of randomness injected into the response.


                      Defaults to `1.0`. Ranges from `0.0` to `1.0`. Use
                      `temperature` closer to `0.0` for analytical / multiple
                      choice, and closer to `1.0` for creative and generative
                      tasks.


                      Note that even with `temperature` of `0.0`, the results
                      will not be fully deterministic.
                    examples:
                      - 1
                    maximum: 1
                    minimum: 0
                    title: Temperature
                    type: number
              thinking:
                allOf:
                  - description: >-
                      Configuration for enabling Claude's extended thinking.


                      When enabled, responses include `thinking` content blocks
                      showing Claude's thinking process before the final answer.
                      Requires a minimum budget of 1,024 tokens and counts
                      towards your `max_tokens` limit.


                      See [extended
                      thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
                      for details.
                    discriminator:
                      mapping:
                        disabled: '#/components/schemas/ThinkingConfigDisabled'
                        enabled: '#/components/schemas/ThinkingConfigEnabled'
                      propertyName: type
                    oneOf:
                      - $ref: '#/components/schemas/ThinkingConfigEnabled'
                      - $ref: '#/components/schemas/ThinkingConfigDisabled'
              tool_choice:
                allOf:
                  - description: >-
                      How the model should use the provided tools. The model can
                      use a specific tool, any available tool, decide by itself,
                      or not use tools at all.
                    discriminator:
                      mapping:
                        any: '#/components/schemas/ToolChoiceAny'
                        auto: '#/components/schemas/ToolChoiceAuto'
                        none: '#/components/schemas/ToolChoiceNone'
                        tool: '#/components/schemas/ToolChoiceTool'
                      propertyName: type
                    oneOf:
                      - $ref: '#/components/schemas/ToolChoiceAuto'
                      - $ref: '#/components/schemas/ToolChoiceAny'
                      - $ref: '#/components/schemas/ToolChoiceTool'
                      - $ref: '#/components/schemas/ToolChoiceNone'
              tools:
                allOf:
                  - description: >-
                      Definitions of tools that the model may use.


                      If you include `tools` in your API request, the model may
                      return `tool_use` content blocks that represent the
                      model's use of those tools. You can then run those tools
                      using the tool input generated by the model and then
                      optionally return results back to the model using
                      `tool_result` content blocks.


                      There are two types of tools: **client tools** and
                      **server tools**. The behavior described below applies to
                      client tools. For [server
                      tools](https://docs.anthropic.com/en/docs/agents-and-tools/tool-use/overview\#server-tools),
                      see their individual documentation as each has its own
                      behavior (e.g., the [web search
                      tool](https://docs.anthropic.com/en/docs/agents-and-tools/tool-use/web-search-tool)).


                      Each tool definition includes:


                      * `name`: Name of the tool.

                      * `description`: Optional, but strongly-recommended
                      description of the tool.

                      * `input_schema`: [JSON
                      schema](https://json-schema.org/draft/2020-12) for the
                      tool `input` shape that the model will produce in
                      `tool_use` output content blocks.


                      For example, if you defined `tools` as:


                      ```json

                      [
                        {
                          "name": "get_stock_price",
                          "description": "Get the current stock price for a given ticker symbol.",
                          "input_schema": {
                            "type": "object",
                            "properties": {
                              "ticker": {
                                "type": "string",
                                "description": "The stock ticker symbol, e.g. AAPL for Apple Inc."
                              }
                            },
                            "required": ["ticker"]
                          }
                        }
                      ]

                      ```


                      And then asked the model "What's the S&P 500 at today?",
                      the model might produce `tool_use` content blocks in the
                      response like this:


                      ```json

                      [
                        {
                          "type": "tool_use",
                          "id": "toolu_01D7FLrfh4GYq7yT1ULFeyMV",
                          "name": "get_stock_price",
                          "input": { "ticker": "^GSPC" }
                        }
                      ]

                      ```


                      You might then run your `get_stock_price` tool with
                      `{"ticker": "^GSPC"}` as an input, and return the
                      following back to the model in a subsequent `user`
                      message:


                      ```json

                      [
                        {
                          "type": "tool_result",
                          "tool_use_id": "toolu_01D7FLrfh4GYq7yT1ULFeyMV",
                          "content": "259.75 USD"
                        }
                      ]

                      ```


                      Tools can be used for workflows that include running
                      client-side tools and functions, or more generally
                      whenever you want the model to produce a particular JSON
                      structure of output.


                      See our
                      [guide](https://docs.anthropic.com/en/docs/tool-use) for
                      more details.
                    examples:
                      - description: Get the current weather in a given location
                        input_schema:
                          properties:
                            location:
                              description: The city and state, e.g. San Francisco, CA
                              type: string
                            unit:
                              description: >-
                                Unit for the output - one of (celsius,
                                fahrenheit)
                              type: string
                          required:
                            - location
                          type: object
                        name: get_weather
                    items:
                      oneOf:
                        - $ref: '#/components/schemas/Tool'
                        - $ref: '#/components/schemas/BashTool_20241022'
                        - $ref: '#/components/schemas/BashTool_20250124'
                        - $ref: '#/components/schemas/CodeExecutionTool_20250522'
                        - $ref: '#/components/schemas/ComputerUseTool_20241022'
                        - $ref: '#/components/schemas/ComputerUseTool_20250124'
                        - $ref: '#/components/schemas/TextEditor_20241022'
                        - $ref: '#/components/schemas/TextEditor_20250124'
                        - $ref: '#/components/schemas/TextEditor_20250429'
                        - $ref: '#/components/schemas/WebSearchTool_20250305'
                    title: Tools
                    type: array
              top_k:
                allOf:
                  - description: >-
                      Only sample from the top K options for each subsequent
                      token.


                      Used to remove "long tail" low probability responses.
                      [Learn more technical details
                      here](https://towardsdatascience.com/how-to-sample-from-language-models-682bceb97277).


                      Recommended for advanced use cases only. You usually only
                      need to use `temperature`.
                    examples:
                      - 5
                    minimum: 0
                    title: Top K
                    type: integer
              top_p:
                allOf:
                  - description: >-
                      Use nucleus sampling.


                      In nucleus sampling, we compute the cumulative
                      distribution over all the options for each subsequent
                      token in decreasing probability order and cut it off once
                      it reaches a particular probability specified by `top_p`.
                      You should either alter `temperature` or `top_p`, but not
                      both.


                      Recommended for advanced use cases only. You usually only
                      need to use `temperature`.
                    examples:
                      - 0.7
                    maximum: 1
                    minimum: 0
                    title: Top P
                    type: number
            required: true
            title: CreateMessageParams
            requiredProperties:
              - model
              - messages
              - max_tokens
            additionalProperties: false
            example:
              max_tokens: 1024
              messages:
                - content: Hello, world
                  role: user
              model: claude-sonnet-4-20250514
        examples:
          example:
            value:
              max_tokens: 1024
              messages:
                - content: Hello, world
                  role: user
              model: claude-sonnet-4-20250514
    codeSamples:
      - lang: bash
        source: |-
          curl https://api.anthropic.com/v1/messages \
               --header "x-api-key: $ANTHROPIC_API_KEY" \
               --header "anthropic-version: 2023-06-01" \
               --header "content-type: application/json" \
               --data \
          '{
              "model": "claude-sonnet-4-20250514",
              "max_tokens": 1024,
              "messages": [
                  {"role": "user", "content": "Hello, world"}
              ]
          }'
      - lang: python
        source: |-
          import anthropic

          anthropic.Anthropic().messages.create(
              model="claude-sonnet-4-20250514",
              max_tokens=1024,
              messages=[
                  {"role": "user", "content": "Hello, world"}
              ]
          )
      - lang: javascript
        source: |-
          import { Anthropic } from '@anthropic-ai/sdk';

          const anthropic = new Anthropic();

          await anthropic.messages.create({
            model: "claude-sonnet-4-20250514",
            max_tokens: 1024,
            messages: [
              {"role": "user", "content": "Hello, world"}
            ]
          });
  response:
    '200':
      application/json:
        schemaArray:
          - type: object
            properties:
              id:
                allOf:
                  - description: |-
                      Unique object identifier.

                      The format and length of IDs may change over time.
                    examples:
                      - msg_013Zva2CMHLNnXjNJJKqJ2EF
                    title: Id
                    type: string
              type:
                allOf:
                  - const: message
                    default: message
                    description: |-
                      Object type.

                      For Messages, this is always `"message"`.
                    enum:
                      - message
                    title: Type
                    type: string
              role:
                allOf:
                  - const: assistant
                    default: assistant
                    description: |-
                      Conversational role of the generated message.

                      This will always be `"assistant"`.
                    enum:
                      - assistant
                    title: Role
                    type: string
              content:
                allOf:
                  - description: >-
                      Content generated by the model.


                      This is an array of content blocks, each of which has a
                      `type` that determines its shape.


                      Example:


                      ```json

                      [{"type": "text", "text": "Hi, I'm Claude."}]

                      ```


                      If the request input `messages` ended with an `assistant`
                      turn, then the response `content` will continue directly
                      from that last turn. You can use this to constrain the
                      model's output.


                      For example, if the input `messages` were:

                      ```json

                      [
                        {"role": "user", "content": "What's the Greek name for Sun? (A) Sol (B) Helios (C) Sun"},
                        {"role": "assistant", "content": "The best answer is ("}
                      ]

                      ```


                      Then the response `content` might be:


                      ```json

                      [{"type": "text", "text": "B)"}]

                      ```
                    examples:
                      - - text: Hi! My name is Claude.
                          type: text
                    items:
                      discriminator:
                        mapping:
                          code_execution_tool_result: >-
                            #/components/schemas/ResponseCodeExecutionToolResultBlock
                          container_upload: '#/components/schemas/ResponseContainerUploadBlock'
                          mcp_tool_result: '#/components/schemas/ResponseMCPToolResultBlock'
                          mcp_tool_use: '#/components/schemas/ResponseMCPToolUseBlock'
                          redacted_thinking: '#/components/schemas/ResponseRedactedThinkingBlock'
                          server_tool_use: '#/components/schemas/ResponseServerToolUseBlock'
                          text: '#/components/schemas/ResponseTextBlock'
                          thinking: '#/components/schemas/ResponseThinkingBlock'
                          tool_use: '#/components/schemas/ResponseToolUseBlock'
                          web_search_tool_result: >-
                            #/components/schemas/ResponseWebSearchToolResultBlock
                        propertyName: type
                      oneOf:
                        - $ref: '#/components/schemas/ResponseTextBlock'
                        - $ref: '#/components/schemas/ResponseThinkingBlock'
                        - $ref: '#/components/schemas/ResponseRedactedThinkingBlock'
                        - $ref: '#/components/schemas/ResponseToolUseBlock'
                        - $ref: '#/components/schemas/ResponseServerToolUseBlock'
                        - $ref: >-
                            #/components/schemas/ResponseWebSearchToolResultBlock
                        - $ref: >-
                            #/components/schemas/ResponseCodeExecutionToolResultBlock
                        - $ref: '#/components/schemas/ResponseMCPToolUseBlock'
                        - $ref: '#/components/schemas/ResponseMCPToolResultBlock'
                        - $ref: '#/components/schemas/ResponseContainerUploadBlock'
                    title: Content
                    type: array
              model:
                allOf:
                  - description: The model that handled the request.
                    examples:
                      - claude-sonnet-4-20250514
                    maxLength: 256
                    minLength: 1
                    title: Model
                    type: string
              stop_reason:
                allOf:
                  - anyOf:
                      - enum:
                          - end_turn
                          - max_tokens
                          - stop_sequence
                          - tool_use
                          - pause_turn
                          - refusal
                        type: string
                      - type: 'null'
                    description: >-
                      The reason that we stopped.


                      This may be one the following values:

                      * `"end_turn"`: the model reached a natural stopping point

                      * `"max_tokens"`: we exceeded the requested `max_tokens`
                      or the model's maximum

                      * `"stop_sequence"`: one of your provided custom
                      `stop_sequences` was generated

                      * `"tool_use"`: the model invoked one or more tools

                      * `"pause_turn"`: we paused a long-running turn. You may
                      provide the response back as-is in a subsequent request to
                      let the model continue.

                      * `"refusal"`: when streaming classifiers intervene to
                      handle potential policy violations


                      In non-streaming mode this value is always non-null. In
                      streaming mode, it is null in the `message_start` event
                      and non-null otherwise.
                    title: Stop Reason
              stop_sequence:
                allOf:
                  - anyOf:
                      - type: string
                      - type: 'null'
                    default: null
                    description: >-
                      Which custom stop sequence was generated, if any.


                      This value will be a non-null string if one of your custom
                      stop sequences was generated.
                    title: Stop Sequence
              usage:
                allOf:
                  - $ref: '#/components/schemas/Usage'
                    description: >-
                      Billing and rate-limit usage.


                      Anthropic's API bills and rate-limits by token counts, as
                      tokens represent the underlying cost to our systems.


                      Under the hood, the API transforms requests into a format
                      suitable for the model. The model's output then goes
                      through a parsing stage before becoming an API response.
                      As a result, the token counts in `usage` will not match
                      one-to-one with the exact visible content of an API
                      request or response.


                      For example, `output_tokens` will be non-zero, even for an
                      empty string response from Claude.


                      Total input tokens in a request is the summation of
                      `input_tokens`, `cache_creation_input_tokens`, and
                      `cache_read_input_tokens`.
                    examples:
                      - input_tokens: 2095
                        output_tokens: 503
              container:
                allOf:
                  - anyOf:
                      - $ref: '#/components/schemas/Container'
                      - type: 'null'
                    default: null
                    description: >-
                      Information about the container used in this request.


                      This will be non-null if a container tool (e.g. code
                      execution) was used.
            title: Message
            examples:
              - content: &ref_0
                  - text: Hi! My name is Claude.
                    type: text
                id: msg_013Zva2CMHLNnXjNJJKqJ2EF
                model: claude-sonnet-4-20250514
                role: assistant
                stop_reason: end_turn
                stop_sequence: null
                type: message
                usage: &ref_1
                  input_tokens: 2095
                  output_tokens: 503
            requiredProperties:
              - id
              - type
              - role
              - content
              - model
              - stop_reason
              - stop_sequence
              - usage
              - container
            example:
              content: *ref_0
              id: msg_013Zva2CMHLNnXjNJJKqJ2EF
              model: claude-sonnet-4-20250514
              role: assistant
              stop_reason: end_turn
              stop_sequence: null
              type: message
              usage: *ref_1
        examples:
          example:
            value:
              content:
                - text: Hi! My name is Claude.
                  type: text
              id: msg_013Zva2CMHLNnXjNJJKqJ2EF
              model: claude-sonnet-4-20250514
              role: assistant
              stop_reason: end_turn
              stop_sequence: null
              type: message
              usage:
                input_tokens: 2095
                output_tokens: 503
        description: Message object.
    4XX:
      application/json:
        schemaArray:
          - type: object
            properties:
              error:
                allOf:
                  - discriminator:
                      mapping:
                        api_error: '#/components/schemas/APIError'
                        authentication_error: '#/components/schemas/AuthenticationError'
                        billing_error: '#/components/schemas/BillingError'
                        invalid_request_error: '#/components/schemas/InvalidRequestError'
                        not_found_error: '#/components/schemas/NotFoundError'
                        overloaded_error: '#/components/schemas/OverloadedError'
                        permission_error: '#/components/schemas/PermissionError'
                        rate_limit_error: '#/components/schemas/RateLimitError'
                        timeout_error: '#/components/schemas/GatewayTimeoutError'
                      propertyName: type
                    oneOf:
                      - $ref: '#/components/schemas/InvalidRequestError'
                      - $ref: '#/components/schemas/AuthenticationError'
                      - $ref: '#/components/schemas/BillingError'
                      - $ref: '#/components/schemas/PermissionError'
                      - $ref: '#/components/schemas/NotFoundError'
                      - $ref: '#/components/schemas/RateLimitError'
                      - $ref: '#/components/schemas/GatewayTimeoutError'
                      - $ref: '#/components/schemas/APIError'
                      - $ref: '#/components/schemas/OverloadedError'
                    title: Error
              type:
                allOf:
                  - const: error
                    default: error
                    enum:
                      - error
                    title: Type
                    type: string
            title: ErrorResponse
            requiredProperties:
              - error
              - type
        examples:
          example:
            value:
              error:
                message: Invalid request
                type: invalid_request_error
              type: error
        description: >-
          Error response.


          See our [errors
          documentation](https://docs.anthropic.com/en/api/errors) for more
          details.
  deprecated: false
  type: path
components:
  schemas:
    APIError:
      properties:
        message:
          default: Internal server error
          title: Message
          type: string
        type:
          const: api_error
          default: api_error
          enum:
            - api_error
          title: Type
          type: string
      required:
        - message
        - type
      title: APIError
      type: object
    AuthenticationError:
      properties:
        message:
          default: Authentication error
          title: Message
          type: string
        type:
          const: authentication_error
          default: authentication_error
          enum:
            - authentication_error
          title: Type
          type: string
      required:
        - message
        - type
      title: AuthenticationError
      type: object
    Base64ImageSource:
      additionalProperties: false
      properties:
        data:
          format: byte
          title: Data
          type: string
        media_type:
          enum:
            - image/jpeg
            - image/png
            - image/gif
            - image/webp
          title: Media Type
          type: string
        type:
          const: base64
          enum:
            - base64
          title: Type
          type: string
      required:
        - data
        - media_type
        - type
      title: Base64ImageSource
      type: object
    Base64PDFSource:
      additionalProperties: false
      properties:
        data:
          format: byte
          title: Data
          type: string
        media_type:
          const: application/pdf
          enum:
            - application/pdf
          title: Media Type
          type: string
        type:
          const: base64
          enum:
            - base64
          title: Type
          type: string
      required:
        - data
        - media_type
        - type
      title: PDF (base64)
      type: object
    BashTool_20241022:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        name:
          const: bash
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          enum:
            - bash
          title: Name
          type: string
        type:
          const: bash_20241022
          enum:
            - bash_20241022
          title: Type
          type: string
      required:
        - name
        - type
      title: Bash tool (2024-10-22)
      type: object
    BashTool_20250124:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        name:
          const: bash
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          enum:
            - bash
          title: Name
          type: string
        type:
          const: bash_20250124
          enum:
            - bash_20250124
          title: Type
          type: string
      required:
        - name
        - type
      title: Bash tool (2025-01-24)
      type: object
    BillingError:
      properties:
        message:
          default: Billing error
          title: Message
          type: string
        type:
          const: billing_error
          default: billing_error
          enum:
            - billing_error
          title: Type
          type: string
      required:
        - message
        - type
      title: BillingError
      type: object
    CacheControlEphemeral:
      additionalProperties: false
      properties:
        ttl:
          description: |-
            The time-to-live for the cache control breakpoint.

            This may be one the following values:
            - `5m`: 5 minutes
            - `1h`: 1 hour

            Defaults to `5m`.
          enum:
            - 5m
            - 1h
          title: Ttl
          type: string
        type:
          const: ephemeral
          enum:
            - ephemeral
          title: Type
          type: string
      required:
        - type
      title: CacheControlEphemeral
      type: object
    CacheCreation:
      properties:
        ephemeral_1h_input_tokens:
          default: 0
          description: The number of input tokens used to create the 1 hour cache entry.
          minimum: 0
          title: Ephemeral 1H Input Tokens
          type: integer
        ephemeral_5m_input_tokens:
          default: 0
          description: The number of input tokens used to create the 5 minute cache entry.
          minimum: 0
          title: Ephemeral 5M Input Tokens
          type: integer
      required:
        - ephemeral_1h_input_tokens
        - ephemeral_5m_input_tokens
      title: CacheCreation
      type: object
    CodeExecutionToolResultErrorCode:
      enum:
        - invalid_tool_input
        - unavailable
        - too_many_requests
        - execution_time_exceeded
      title: CodeExecutionToolResultErrorCode
      type: string
    CodeExecutionTool_20250522:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        name:
          const: code_execution
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          enum:
            - code_execution
          title: Name
          type: string
        type:
          const: code_execution_20250522
          enum:
            - code_execution_20250522
          title: Type
          type: string
      required:
        - name
        - type
      title: Code execution tool (2025-05-22)
      type: object
    ComputerUseTool_20241022:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        display_height_px:
          description: The height of the display in pixels.
          minimum: 1
          title: Display Height Px
          type: integer
        display_number:
          anyOf:
            - minimum: 0
              type: integer
            - type: 'null'
          description: The X11 display number (e.g. 0, 1) for the display.
          title: Display Number
        display_width_px:
          description: The width of the display in pixels.
          minimum: 1
          title: Display Width Px
          type: integer
        name:
          const: computer
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          enum:
            - computer
          title: Name
          type: string
        type:
          const: computer_20241022
          enum:
            - computer_20241022
          title: Type
          type: string
      required:
        - display_height_px
        - display_width_px
        - name
        - type
      title: Computer use tool (2024-01-22)
      type: object
    ComputerUseTool_20250124:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        display_height_px:
          description: The height of the display in pixels.
          minimum: 1
          title: Display Height Px
          type: integer
        display_number:
          anyOf:
            - minimum: 0
              type: integer
            - type: 'null'
          description: The X11 display number (e.g. 0, 1) for the display.
          title: Display Number
        display_width_px:
          description: The width of the display in pixels.
          minimum: 1
          title: Display Width Px
          type: integer
        name:
          const: computer
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          enum:
            - computer
          title: Name
          type: string
        type:
          const: computer_20250124
          enum:
            - computer_20250124
          title: Type
          type: string
      required:
        - display_height_px
        - display_width_px
        - name
        - type
      title: Computer use tool (2025-01-24)
      type: object
    Container:
      description: >-
        Information about the container used in the request (for the code
        execution tool)
      properties:
        expires_at:
          description: The time at which the container will expire.
          format: date-time
          title: Expires At
          type: string
        id:
          description: Identifier for the container used in this request
          title: Id
          type: string
      required:
        - expires_at
        - id
      title: Container
      type: object
    ContentBlockSource:
      additionalProperties: false
      properties:
        content:
          anyOf:
            - type: string
            - items:
                discriminator:
                  mapping:
                    image: '#/components/schemas/RequestImageBlock'
                    text: '#/components/schemas/RequestTextBlock'
                  propertyName: type
                oneOf:
                  - $ref: '#/components/schemas/RequestTextBlock'
                  - $ref: '#/components/schemas/RequestImageBlock'
              type: array
          title: Content
        type:
          const: content
          enum:
            - content
          title: Type
          type: string
      required:
        - content
        - type
      title: Content block
      type: object
    FileDocumentSource:
      additionalProperties: false
      properties:
        file_id:
          title: File Id
          type: string
        type:
          const: file
          enum:
            - file
          title: Type
          type: string
      required:
        - file_id
        - type
      title: File document
      type: object
    FileImageSource:
      additionalProperties: false
      properties:
        file_id:
          title: File Id
          type: string
        type:
          const: file
          enum:
            - file
          title: Type
          type: string
      required:
        - file_id
        - type
      title: FileImageSource
      type: object
    GatewayTimeoutError:
      properties:
        message:
          default: Request timeout
          title: Message
          type: string
        type:
          const: timeout_error
          default: timeout_error
          enum:
            - timeout_error
          title: Type
          type: string
      required:
        - message
        - type
      title: GatewayTimeoutError
      type: object
    InputMessage:
      additionalProperties: false
      properties:
        content:
          anyOf:
            - type: string
            - items:
                discriminator:
                  mapping:
                    code_execution_tool_result: '#/components/schemas/RequestCodeExecutionToolResultBlock'
                    container_upload: '#/components/schemas/RequestContainerUploadBlock'
                    document: '#/components/schemas/RequestDocumentBlock'
                    image: '#/components/schemas/RequestImageBlock'
                    mcp_tool_result: '#/components/schemas/RequestMCPToolResultBlock'
                    mcp_tool_use: '#/components/schemas/RequestMCPToolUseBlock'
                    redacted_thinking: '#/components/schemas/RequestRedactedThinkingBlock'
                    search_result: '#/components/schemas/RequestSearchResultBlock'
                    server_tool_use: '#/components/schemas/RequestServerToolUseBlock'
                    text: '#/components/schemas/RequestTextBlock'
                    thinking: '#/components/schemas/RequestThinkingBlock'
                    tool_result: '#/components/schemas/RequestToolResultBlock'
                    tool_use: '#/components/schemas/RequestToolUseBlock'
                    web_search_tool_result: '#/components/schemas/RequestWebSearchToolResultBlock'
                  propertyName: type
                oneOf:
                  - $ref: '#/components/schemas/RequestTextBlock'
                    description: Regular text content.
                  - $ref: '#/components/schemas/RequestImageBlock'
                    description: >-
                      Image content specified directly as base64 data or as a
                      reference via a URL.
                  - $ref: '#/components/schemas/RequestDocumentBlock'
                    description: >-
                      Document content, either specified directly as base64
                      data, as text, or as a reference via a URL.
                  - $ref: '#/components/schemas/RequestSearchResultBlock'
                    description: >-
                      A search result block containing source, title, and
                      content from search operations.
                  - $ref: '#/components/schemas/RequestThinkingBlock'
                    description: A block specifying internal thinking by the model.
                  - $ref: '#/components/schemas/RequestRedactedThinkingBlock'
                    description: >-
                      A block specifying internal, redacted thinking by the
                      model.
                  - $ref: '#/components/schemas/RequestToolUseBlock'
                    description: A block indicating a tool use by the model.
                  - $ref: '#/components/schemas/RequestToolResultBlock'
                    description: A block specifying the results of a tool use by the model.
                  - $ref: '#/components/schemas/RequestServerToolUseBlock'
                  - $ref: '#/components/schemas/RequestWebSearchToolResultBlock'
                  - $ref: '#/components/schemas/RequestCodeExecutionToolResultBlock'
                  - $ref: '#/components/schemas/RequestMCPToolUseBlock'
                  - $ref: '#/components/schemas/RequestMCPToolResultBlock'
                  - $ref: '#/components/schemas/RequestContainerUploadBlock'
              type: array
          title: Content
        role:
          enum:
            - user
            - assistant
          title: Role
          type: string
      required:
        - content
        - role
      title: InputMessage
      type: object
    InputSchema:
      additionalProperties: true
      properties:
        properties:
          anyOf:
            - type: object
            - type: 'null'
          title: Properties
        required:
          anyOf:
            - items:
                type: string
              type: array
            - type: 'null'
          title: Required
        type:
          const: object
          enum:
            - object
          title: Type
          type: string
      required:
        - type
      title: InputSchema
      type: object
    InvalidRequestError:
      properties:
        message:
          default: Invalid request
          title: Message
          type: string
        type:
          const: invalid_request_error
          default: invalid_request_error
          enum:
            - invalid_request_error
          title: Type
          type: string
      required:
        - message
        - type
      title: InvalidRequestError
      type: object
    Metadata:
      additionalProperties: false
      properties:
        user_id:
          anyOf:
            - maxLength: 256
              type: string
            - type: 'null'
          description: >-
            An external identifier for the user who is associated with the
            request.


            This should be a uuid, hash value, or other opaque identifier.
            Anthropic may use this id to help detect abuse. Do not include any
            identifying information such as name, email address, or phone
            number.
          examples:
            - 13803d75-b4b5-4c3e-b2a2-6f21399b021b
          title: User Id
      title: Metadata
      type: object
    NotFoundError:
      properties:
        message:
          default: Not found
          title: Message
          type: string
        type:
          const: not_found_error
          default: not_found_error
          enum:
            - not_found_error
          title: Type
          type: string
      required:
        - message
        - type
      title: NotFoundError
      type: object
    OverloadedError:
      properties:
        message:
          default: Overloaded
          title: Message
          type: string
        type:
          const: overloaded_error
          default: overloaded_error
          enum:
            - overloaded_error
          title: Type
          type: string
      required:
        - message
        - type
      title: OverloadedError
      type: object
    PermissionError:
      properties:
        message:
          default: Permission denied
          title: Message
          type: string
        type:
          const: permission_error
          default: permission_error
          enum:
            - permission_error
          title: Type
          type: string
      required:
        - message
        - type
      title: PermissionError
      type: object
    PlainTextSource:
      additionalProperties: false
      properties:
        data:
          title: Data
          type: string
        media_type:
          const: text/plain
          enum:
            - text/plain
          title: Media Type
          type: string
        type:
          const: text
          enum:
            - text
          title: Type
          type: string
      required:
        - data
        - media_type
        - type
      title: Plain text
      type: object
    RateLimitError:
      properties:
        message:
          default: Rate limited
          title: Message
          type: string
        type:
          const: rate_limit_error
          default: rate_limit_error
          enum:
            - rate_limit_error
          title: Type
          type: string
      required:
        - message
        - type
      title: RateLimitError
      type: object
    RequestCharLocationCitation:
      additionalProperties: false
      properties:
        cited_text:
          title: Cited Text
          type: string
        document_index:
          minimum: 0
          title: Document Index
          type: integer
        document_title:
          anyOf:
            - maxLength: 255
              minLength: 1
              type: string
            - type: 'null'
          title: Document Title
        end_char_index:
          title: End Char Index
          type: integer
        start_char_index:
          minimum: 0
          title: Start Char Index
          type: integer
        type:
          const: char_location
          enum:
            - char_location
          title: Type
          type: string
      required:
        - cited_text
        - document_index
        - document_title
        - end_char_index
        - start_char_index
        - type
      title: Character location
      type: object
    RequestCitationsConfig:
      additionalProperties: false
      properties:
        enabled:
          title: Enabled
          type: boolean
      title: RequestCitationsConfig
      type: object
    RequestCodeExecutionOutputBlock:
      additionalProperties: false
      properties:
        file_id:
          title: File Id
          type: string
        type:
          const: code_execution_output
          enum:
            - code_execution_output
          title: Type
          type: string
      required:
        - file_id
        - type
      title: RequestCodeExecutionOutputBlock
      type: object
    RequestCodeExecutionResultBlock:
      additionalProperties: false
      properties:
        content:
          items:
            $ref: '#/components/schemas/RequestCodeExecutionOutputBlock'
          title: Content
          type: array
        return_code:
          title: Return Code
          type: integer
        stderr:
          title: Stderr
          type: string
        stdout:
          title: Stdout
          type: string
        type:
          const: code_execution_result
          enum:
            - code_execution_result
          title: Type
          type: string
      required:
        - content
        - return_code
        - stderr
        - stdout
        - type
      title: Code execution result
      type: object
    RequestCodeExecutionToolResultBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        content:
          anyOf:
            - $ref: '#/components/schemas/RequestCodeExecutionToolResultError'
            - $ref: '#/components/schemas/RequestCodeExecutionResultBlock'
          title: Content
        tool_use_id:
          pattern: ^srvtoolu_[a-zA-Z0-9_]+$
          title: Tool Use Id
          type: string
        type:
          const: code_execution_tool_result
          enum:
            - code_execution_tool_result
          title: Type
          type: string
      required:
        - content
        - tool_use_id
        - type
      title: Code execution tool result
      type: object
    RequestCodeExecutionToolResultError:
      additionalProperties: false
      properties:
        error_code:
          $ref: '#/components/schemas/CodeExecutionToolResultErrorCode'
        type:
          const: code_execution_tool_result_error
          enum:
            - code_execution_tool_result_error
          title: Type
          type: string
      required:
        - error_code
        - type
      title: Code execution tool error
      type: object
    RequestContainerUploadBlock:
      additionalProperties: false
      description: >-
        A content block that represents a file to be uploaded to the container

        Files uploaded via this block will be available in the container's input
        directory.
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        file_id:
          title: File Id
          type: string
        type:
          const: container_upload
          enum:
            - container_upload
          title: Type
          type: string
      required:
        - file_id
        - type
      title: Container upload
      type: object
    RequestContentBlockLocationCitation:
      additionalProperties: false
      properties:
        cited_text:
          title: Cited Text
          type: string
        document_index:
          minimum: 0
          title: Document Index
          type: integer
        document_title:
          anyOf:
            - maxLength: 255
              minLength: 1
              type: string
            - type: 'null'
          title: Document Title
        end_block_index:
          title: End Block Index
          type: integer
        start_block_index:
          minimum: 0
          title: Start Block Index
          type: integer
        type:
          const: content_block_location
          enum:
            - content_block_location
          title: Type
          type: string
      required:
        - cited_text
        - document_index
        - document_title
        - end_block_index
        - start_block_index
        - type
      title: Content block location
      type: object
    RequestDocumentBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        citations:
          $ref: '#/components/schemas/RequestCitationsConfig'
        context:
          anyOf:
            - minLength: 1
              type: string
            - type: 'null'
          title: Context
        source:
          discriminator:
            mapping:
              base64: '#/components/schemas/Base64PDFSource'
              content: '#/components/schemas/ContentBlockSource'
              file: '#/components/schemas/FileDocumentSource'
              text: '#/components/schemas/PlainTextSource'
              url: '#/components/schemas/URLPDFSource'
            propertyName: type
          oneOf:
            - $ref: '#/components/schemas/Base64PDFSource'
            - $ref: '#/components/schemas/PlainTextSource'
            - $ref: '#/components/schemas/ContentBlockSource'
            - $ref: '#/components/schemas/URLPDFSource'
            - $ref: '#/components/schemas/FileDocumentSource'
        title:
          anyOf:
            - maxLength: 500
              minLength: 1
              type: string
            - type: 'null'
          title: Title
        type:
          const: document
          enum:
            - document
          title: Type
          type: string
      required:
        - source
        - type
      title: Document
      type: object
    RequestImageBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        source:
          discriminator:
            mapping:
              base64: '#/components/schemas/Base64ImageSource'
              file: '#/components/schemas/FileImageSource'
              url: '#/components/schemas/URLImageSource'
            propertyName: type
          oneOf:
            - $ref: '#/components/schemas/Base64ImageSource'
            - $ref: '#/components/schemas/URLImageSource'
            - $ref: '#/components/schemas/FileImageSource'
          title: Source
        type:
          const: image
          enum:
            - image
          title: Type
          type: string
      required:
        - source
        - type
      title: Image
      type: object
    RequestMCPServerToolConfiguration:
      additionalProperties: false
      properties:
        allowed_tools:
          anyOf:
            - items:
                type: string
              type: array
            - type: 'null'
          title: Allowed Tools
        enabled:
          anyOf:
            - type: boolean
            - type: 'null'
          title: Enabled
      title: RequestMCPServerToolConfiguration
      type: object
    RequestMCPServerURLDefinition:
      additionalProperties: false
      properties:
        authorization_token:
          anyOf:
            - type: string
            - type: 'null'
          title: Authorization Token
        name:
          title: Name
          type: string
        tool_configuration:
          anyOf:
            - $ref: '#/components/schemas/RequestMCPServerToolConfiguration'
            - type: 'null'
        type:
          const: url
          enum:
            - url
          title: Type
          type: string
        url:
          title: Url
          type: string
      required:
        - name
        - type
        - url
      title: RequestMCPServerURLDefinition
      type: object
    RequestMCPToolResultBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        content:
          anyOf:
            - type: string
            - items:
                $ref: '#/components/schemas/RequestTextBlock'
              type: array
          title: Content
        is_error:
          title: Is Error
          type: boolean
        tool_use_id:
          pattern: ^[a-zA-Z0-9_-]+$
          title: Tool Use Id
          type: string
        type:
          const: mcp_tool_result
          enum:
            - mcp_tool_result
          title: Type
          type: string
      required:
        - tool_use_id
        - type
      title: MCP tool result
      type: object
    RequestMCPToolUseBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        id:
          pattern: ^[a-zA-Z0-9_-]+$
          title: Id
          type: string
        input:
          title: Input
          type: object
        name:
          title: Name
          type: string
        server_name:
          description: The name of the MCP server
          title: Server Name
          type: string
        type:
          const: mcp_tool_use
          enum:
            - mcp_tool_use
          title: Type
          type: string
      required:
        - id
        - input
        - name
        - server_name
        - type
      title: MCP tool use
      type: object
    RequestPageLocationCitation:
      additionalProperties: false
      properties:
        cited_text:
          title: Cited Text
          type: string
        document_index:
          minimum: 0
          title: Document Index
          type: integer
        document_title:
          anyOf:
            - maxLength: 255
              minLength: 1
              type: string
            - type: 'null'
          title: Document Title
        end_page_number:
          title: End Page Number
          type: integer
        start_page_number:
          minimum: 1
          title: Start Page Number
          type: integer
        type:
          const: page_location
          enum:
            - page_location
          title: Type
          type: string
      required:
        - cited_text
        - document_index
        - document_title
        - end_page_number
        - start_page_number
        - type
      title: Page location
      type: object
    RequestRedactedThinkingBlock:
      additionalProperties: false
      properties:
        data:
          title: Data
          type: string
        type:
          const: redacted_thinking
          enum:
            - redacted_thinking
          title: Type
          type: string
      required:
        - data
        - type
      title: Redacted thinking
      type: object
    RequestSearchResultBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        citations:
          $ref: '#/components/schemas/RequestCitationsConfig'
        content:
          items:
            $ref: '#/components/schemas/RequestTextBlock'
          title: Content
          type: array
        source:
          title: Source
          type: string
        title:
          title: Title
          type: string
        type:
          const: search_result
          enum:
            - search_result
          title: Type
          type: string
      required:
        - content
        - source
        - title
        - type
      title: RequestSearchResultBlock
      type: object
    RequestSearchResultLocationCitation:
      properties:
        cited_text:
          title: Cited Text
          type: string
        end_block_index:
          title: End Block Index
          type: integer
        search_result_index:
          minimum: 0
          title: Search Result Index
          type: integer
        source:
          title: Source
          type: string
        start_block_index:
          minimum: 0
          title: Start Block Index
          type: integer
        title:
          anyOf:
            - type: string
            - type: 'null'
          title: Title
        type:
          const: search_result_location
          default: search_result_location
          enum:
            - search_result_location
          title: Type
          type: string
      required:
        - cited_text
        - end_block_index
        - search_result_index
        - source
        - start_block_index
        - title
        - type
      title: RequestSearchResultLocationCitation
      type: object
    RequestServerToolUseBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        id:
          pattern: ^srvtoolu_[a-zA-Z0-9_]+$
          title: Id
          type: string
        input:
          title: Input
          type: object
        name:
          enum:
            - web_search
            - code_execution
          title: Name
          type: string
        type:
          const: server_tool_use
          enum:
            - server_tool_use
          title: Type
          type: string
      required:
        - id
        - input
        - name
        - type
      title: Server tool use
      type: object
    RequestTextBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        citations:
          anyOf:
            - items:
                discriminator:
                  mapping:
                    char_location: '#/components/schemas/RequestCharLocationCitation'
                    content_block_location: '#/components/schemas/RequestContentBlockLocationCitation'
                    page_location: '#/components/schemas/RequestPageLocationCitation'
                    search_result_location: '#/components/schemas/RequestSearchResultLocationCitation'
                    web_search_result_location: >-
                      #/components/schemas/RequestWebSearchResultLocationCitation
                  propertyName: type
                oneOf:
                  - $ref: '#/components/schemas/RequestCharLocationCitation'
                  - $ref: '#/components/schemas/RequestPageLocationCitation'
                  - $ref: '#/components/schemas/RequestContentBlockLocationCitation'
                  - $ref: >-
                      #/components/schemas/RequestWebSearchResultLocationCitation
                  - $ref: '#/components/schemas/RequestSearchResultLocationCitation'
              type: array
            - type: 'null'
          title: Citations
        text:
          minLength: 1
          title: Text
          type: string
        type:
          const: text
          enum:
            - text
          title: Type
          type: string
      required:
        - text
        - type
      title: Text
      type: object
    RequestThinkingBlock:
      additionalProperties: false
      properties:
        signature:
          title: Signature
          type: string
        thinking:
          title: Thinking
          type: string
        type:
          const: thinking
          enum:
            - thinking
          title: Type
          type: string
      required:
        - signature
        - thinking
        - type
      title: Thinking
      type: object
    RequestToolResultBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        content:
          anyOf:
            - type: string
            - items:
                discriminator:
                  mapping:
                    image: '#/components/schemas/RequestImageBlock'
                    search_result: '#/components/schemas/RequestSearchResultBlock'
                    text: '#/components/schemas/RequestTextBlock'
                  propertyName: type
                oneOf:
                  - $ref: '#/components/schemas/RequestTextBlock'
                  - $ref: '#/components/schemas/RequestImageBlock'
                  - $ref: '#/components/schemas/RequestSearchResultBlock'
              type: array
          title: Content
        is_error:
          title: Is Error
          type: boolean
        tool_use_id:
          pattern: ^[a-zA-Z0-9_-]+$
          title: Tool Use Id
          type: string
        type:
          const: tool_result
          enum:
            - tool_result
          title: Type
          type: string
      required:
        - tool_use_id
        - type
      title: Tool result
      type: object
    RequestToolUseBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        id:
          pattern: ^[a-zA-Z0-9_-]+$
          title: Id
          type: string
        input:
          title: Input
          type: object
        name:
          maxLength: 200
          minLength: 1
          title: Name
          type: string
        type:
          const: tool_use
          enum:
            - tool_use
          title: Type
          type: string
      required:
        - id
        - input
        - name
        - type
      title: Tool use
      type: object
    RequestWebSearchResultBlock:
      additionalProperties: false
      properties:
        encrypted_content:
          title: Encrypted Content
          type: string
        page_age:
          anyOf:
            - type: string
            - type: 'null'
          title: Page Age
        title:
          title: Title
          type: string
        type:
          const: web_search_result
          enum:
            - web_search_result
          title: Type
          type: string
        url:
          title: Url
          type: string
      required:
        - encrypted_content
        - title
        - type
        - url
      title: RequestWebSearchResultBlock
      type: object
    RequestWebSearchResultLocationCitation:
      additionalProperties: false
      properties:
        cited_text:
          title: Cited Text
          type: string
        encrypted_index:
          title: Encrypted Index
          type: string
        title:
          anyOf:
            - maxLength: 512
              minLength: 1
              type: string
            - type: 'null'
          title: Title
        type:
          const: web_search_result_location
          enum:
            - web_search_result_location
          title: Type
          type: string
        url:
          maxLength: 2048
          minLength: 1
          title: Url
          type: string
      required:
        - cited_text
        - encrypted_index
        - title
        - type
        - url
      title: RequestWebSearchResultLocationCitation
      type: object
    RequestWebSearchToolResultBlock:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        content:
          anyOf:
            - items:
                $ref: '#/components/schemas/RequestWebSearchResultBlock'
              type: array
            - $ref: '#/components/schemas/RequestWebSearchToolResultError'
          title: Content
        tool_use_id:
          pattern: ^srvtoolu_[a-zA-Z0-9_]+$
          title: Tool Use Id
          type: string
        type:
          const: web_search_tool_result
          enum:
            - web_search_tool_result
          title: Type
          type: string
      required:
        - content
        - tool_use_id
        - type
      title: Web search tool result
      type: object
    RequestWebSearchToolResultError:
      additionalProperties: false
      properties:
        error_code:
          $ref: '#/components/schemas/WebSearchToolResultErrorCode'
        type:
          const: web_search_tool_result_error
          enum:
            - web_search_tool_result_error
          title: Type
          type: string
      required:
        - error_code
        - type
      title: RequestWebSearchToolResultError
      type: object
    ResponseCharLocationCitation:
      properties:
        cited_text:
          title: Cited Text
          type: string
        document_index:
          minimum: 0
          title: Document Index
          type: integer
        document_title:
          anyOf:
            - type: string
            - type: 'null'
          title: Document Title
        end_char_index:
          title: End Char Index
          type: integer
        start_char_index:
          minimum: 0
          title: Start Char Index
          type: integer
        type:
          const: char_location
          default: char_location
          enum:
            - char_location
          title: Type
          type: string
      required:
        - cited_text
        - document_index
        - document_title
        - end_char_index
        - start_char_index
        - type
      title: Character location
      type: object
    ResponseCodeExecutionOutputBlock:
      properties:
        file_id:
          title: File Id
          type: string
        type:
          const: code_execution_output
          default: code_execution_output
          enum:
            - code_execution_output
          title: Type
          type: string
      required:
        - file_id
        - type
      title: ResponseCodeExecutionOutputBlock
      type: object
    ResponseCodeExecutionResultBlock:
      properties:
        content:
          items:
            $ref: '#/components/schemas/ResponseCodeExecutionOutputBlock'
          title: Content
          type: array
        return_code:
          title: Return Code
          type: integer
        stderr:
          title: Stderr
          type: string
        stdout:
          title: Stdout
          type: string
        type:
          const: code_execution_result
          default: code_execution_result
          enum:
            - code_execution_result
          title: Type
          type: string
      required:
        - content
        - return_code
        - stderr
        - stdout
        - type
      title: Code execution result
      type: object
    ResponseCodeExecutionToolResultBlock:
      properties:
        content:
          anyOf:
            - $ref: '#/components/schemas/ResponseCodeExecutionToolResultError'
            - $ref: '#/components/schemas/ResponseCodeExecutionResultBlock'
          title: Content
        tool_use_id:
          pattern: ^srvtoolu_[a-zA-Z0-9_]+$
          title: Tool Use Id
          type: string
        type:
          const: code_execution_tool_result
          default: code_execution_tool_result
          enum:
            - code_execution_tool_result
          title: Type
          type: string
      required:
        - content
        - tool_use_id
        - type
      title: Code execution tool result
      type: object
    ResponseCodeExecutionToolResultError:
      properties:
        error_code:
          $ref: '#/components/schemas/CodeExecutionToolResultErrorCode'
        type:
          const: code_execution_tool_result_error
          default: code_execution_tool_result_error
          enum:
            - code_execution_tool_result_error
          title: Type
          type: string
      required:
        - error_code
        - type
      title: Code execution tool error
      type: object
    ResponseContainerUploadBlock:
      description: Response model for a file uploaded to the container.
      properties:
        file_id:
          title: File Id
          type: string
        type:
          const: container_upload
          default: container_upload
          enum:
            - container_upload
          title: Type
          type: string
      required:
        - file_id
        - type
      title: Container upload
      type: object
    ResponseContentBlockLocationCitation:
      properties:
        cited_text:
          title: Cited Text
          type: string
        document_index:
          minimum: 0
          title: Document Index
          type: integer
        document_title:
          anyOf:
            - type: string
            - type: 'null'
          title: Document Title
        end_block_index:
          title: End Block Index
          type: integer
        start_block_index:
          minimum: 0
          title: Start Block Index
          type: integer
        type:
          const: content_block_location
          default: content_block_location
          enum:
            - content_block_location
          title: Type
          type: string
      required:
        - cited_text
        - document_index
        - document_title
        - end_block_index
        - start_block_index
        - type
      title: Content block location
      type: object
    ResponseMCPToolResultBlock:
      properties:
        content:
          anyOf:
            - type: string
            - items:
                $ref: '#/components/schemas/ResponseTextBlock'
              type: array
          title: Content
        is_error:
          default: false
          title: Is Error
          type: boolean
        tool_use_id:
          pattern: ^[a-zA-Z0-9_-]+$
          title: Tool Use Id
          type: string
        type:
          const: mcp_tool_result
          default: mcp_tool_result
          enum:
            - mcp_tool_result
          title: Type
          type: string
      required:
        - content
        - is_error
        - tool_use_id
        - type
      title: MCP tool result
      type: object
    ResponseMCPToolUseBlock:
      properties:
        id:
          pattern: ^[a-zA-Z0-9_-]+$
          title: Id
          type: string
        input:
          title: Input
          type: object
        name:
          description: The name of the MCP tool
          title: Name
          type: string
        server_name:
          description: The name of the MCP server
          title: Server Name
          type: string
        type:
          const: mcp_tool_use
          default: mcp_tool_use
          enum:
            - mcp_tool_use
          title: Type
          type: string
      required:
        - id
        - input
        - name
        - server_name
        - type
      title: MCP tool use
      type: object
    ResponsePageLocationCitation:
      properties:
        cited_text:
          title: Cited Text
          type: string
        document_index:
          minimum: 0
          title: Document Index
          type: integer
        document_title:
          anyOf:
            - type: string
            - type: 'null'
          title: Document Title
        end_page_number:
          title: End Page Number
          type: integer
        start_page_number:
          minimum: 1
          title: Start Page Number
          type: integer
        type:
          const: page_location
          default: page_location
          enum:
            - page_location
          title: Type
          type: string
      required:
        - cited_text
        - document_index
        - document_title
        - end_page_number
        - start_page_number
        - type
      title: Page location
      type: object
    ResponseRedactedThinkingBlock:
      properties:
        data:
          title: Data
          type: string
        type:
          const: redacted_thinking
          default: redacted_thinking
          enum:
            - redacted_thinking
          title: Type
          type: string
      required:
        - data
        - type
      title: Redacted thinking
      type: object
    ResponseSearchResultLocationCitation:
      properties:
        cited_text:
          title: Cited Text
          type: string
        end_block_index:
          title: End Block Index
          type: integer
        search_result_index:
          minimum: 0
          title: Search Result Index
          type: integer
        source:
          title: Source
          type: string
        start_block_index:
          minimum: 0
          title: Start Block Index
          type: integer
        title:
          anyOf:
            - type: string
            - type: 'null'
          title: Title
        type:
          const: search_result_location
          default: search_result_location
          enum:
            - search_result_location
          title: Type
          type: string
      required:
        - cited_text
        - end_block_index
        - search_result_index
        - source
        - start_block_index
        - title
        - type
      title: ResponseSearchResultLocationCitation
      type: object
    ResponseServerToolUseBlock:
      properties:
        id:
          pattern: ^srvtoolu_[a-zA-Z0-9_]+$
          title: Id
          type: string
        input:
          title: Input
          type: object
        name:
          enum:
            - web_search
            - code_execution
          title: Name
          type: string
        type:
          const: server_tool_use
          default: server_tool_use
          enum:
            - server_tool_use
          title: Type
          type: string
      required:
        - id
        - input
        - name
        - type
      title: Server tool use
      type: object
    ResponseTextBlock:
      properties:
        citations:
          anyOf:
            - items:
                discriminator:
                  mapping:
                    char_location: '#/components/schemas/ResponseCharLocationCitation'
                    content_block_location: '#/components/schemas/ResponseContentBlockLocationCitation'
                    page_location: '#/components/schemas/ResponsePageLocationCitation'
                    search_result_location: '#/components/schemas/ResponseSearchResultLocationCitation'
                    web_search_result_location: >-
                      #/components/schemas/ResponseWebSearchResultLocationCitation
                  propertyName: type
                oneOf:
                  - $ref: '#/components/schemas/ResponseCharLocationCitation'
                  - $ref: '#/components/schemas/ResponsePageLocationCitation'
                  - $ref: '#/components/schemas/ResponseContentBlockLocationCitation'
                  - $ref: >-
                      #/components/schemas/ResponseWebSearchResultLocationCitation
                  - $ref: '#/components/schemas/ResponseSearchResultLocationCitation'
              type: array
            - type: 'null'
          default: null
          description: >-
            Citations supporting the text block.


            The type of citation returned will depend on the type of document
            being cited. Citing a PDF results in `page_location`, plain text
            results in `char_location`, and content document results in
            `content_block_location`.
          title: Citations
        text:
          maxLength: 5000000
          minLength: 0
          title: Text
          type: string
        type:
          const: text
          default: text
          enum:
            - text
          title: Type
          type: string
      required:
        - citations
        - text
        - type
      title: Text
      type: object
    ResponseThinkingBlock:
      properties:
        signature:
          title: Signature
          type: string
        thinking:
          title: Thinking
          type: string
        type:
          const: thinking
          default: thinking
          enum:
            - thinking
          title: Type
          type: string
      required:
        - signature
        - thinking
        - type
      title: Thinking
      type: object
    ResponseToolUseBlock:
      properties:
        id:
          pattern: ^[a-zA-Z0-9_-]+$
          title: Id
          type: string
        input:
          title: Input
          type: object
        name:
          minLength: 1
          title: Name
          type: string
        type:
          const: tool_use
          default: tool_use
          enum:
            - tool_use
          title: Type
          type: string
      required:
        - id
        - input
        - name
        - type
      title: Tool use
      type: object
    ResponseWebSearchResultBlock:
      properties:
        encrypted_content:
          title: Encrypted Content
          type: string
        page_age:
          anyOf:
            - type: string
            - type: 'null'
          default: null
          title: Page Age
        title:
          title: Title
          type: string
        type:
          const: web_search_result
          default: web_search_result
          enum:
            - web_search_result
          title: Type
          type: string
        url:
          title: Url
          type: string
      required:
        - encrypted_content
        - page_age
        - title
        - type
        - url
      title: ResponseWebSearchResultBlock
      type: object
    ResponseWebSearchResultLocationCitation:
      properties:
        cited_text:
          title: Cited Text
          type: string
        encrypted_index:
          title: Encrypted Index
          type: string
        title:
          anyOf:
            - maxLength: 512
              type: string
            - type: 'null'
          title: Title
        type:
          const: web_search_result_location
          default: web_search_result_location
          enum:
            - web_search_result_location
          title: Type
          type: string
        url:
          title: Url
          type: string
      required:
        - cited_text
        - encrypted_index
        - title
        - type
        - url
      title: ResponseWebSearchResultLocationCitation
      type: object
    ResponseWebSearchToolResultBlock:
      properties:
        content:
          anyOf:
            - $ref: '#/components/schemas/ResponseWebSearchToolResultError'
            - items:
                $ref: '#/components/schemas/ResponseWebSearchResultBlock'
              type: array
          title: Content
        tool_use_id:
          pattern: ^srvtoolu_[a-zA-Z0-9_]+$
          title: Tool Use Id
          type: string
        type:
          const: web_search_tool_result
          default: web_search_tool_result
          enum:
            - web_search_tool_result
          title: Type
          type: string
      required:
        - content
        - tool_use_id
        - type
      title: Web search tool result
      type: object
    ResponseWebSearchToolResultError:
      properties:
        error_code:
          $ref: '#/components/schemas/WebSearchToolResultErrorCode'
        type:
          const: web_search_tool_result_error
          default: web_search_tool_result_error
          enum:
            - web_search_tool_result_error
          title: Type
          type: string
      required:
        - error_code
        - type
      title: ResponseWebSearchToolResultError
      type: object
    ServerToolUsage:
      properties:
        web_search_requests:
          default: 0
          description: The number of web search tool requests.
          examples:
            - 0
          minimum: 0
          title: Web Search Requests
          type: integer
      required:
        - web_search_requests
      title: ServerToolUsage
      type: object
    TextEditor_20241022:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        name:
          const: str_replace_editor
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          enum:
            - str_replace_editor
          title: Name
          type: string
        type:
          const: text_editor_20241022
          enum:
            - text_editor_20241022
          title: Type
          type: string
      required:
        - name
        - type
      title: Text editor tool (2024-10-22)
      type: object
    TextEditor_20250124:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        name:
          const: str_replace_editor
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          enum:
            - str_replace_editor
          title: Name
          type: string
        type:
          const: text_editor_20250124
          enum:
            - text_editor_20250124
          title: Type
          type: string
      required:
        - name
        - type
      title: Text editor tool (2025-01-24)
      type: object
    TextEditor_20250429:
      additionalProperties: false
      properties:
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        name:
          const: str_replace_based_edit_tool
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          enum:
            - str_replace_based_edit_tool
          title: Name
          type: string
        type:
          const: text_editor_20250429
          enum:
            - text_editor_20250429
          title: Type
          type: string
      required:
        - name
        - type
      title: Text editor tool (2025-04-29)
      type: object
    ThinkingConfigDisabled:
      additionalProperties: false
      properties:
        type:
          const: disabled
          enum:
            - disabled
          title: Type
          type: string
      required:
        - type
      title: Disabled
      type: object
    ThinkingConfigEnabled:
      additionalProperties: false
      properties:
        budget_tokens:
          description: >-
            Determines how many tokens Claude can use for its internal reasoning
            process. Larger budgets can enable more thorough analysis for
            complex problems, improving response quality.


            Must be ≥1024 and less than `max_tokens`.


            See [extended
            thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
            for details.
          minimum: 1024
          title: Budget Tokens
          type: integer
        type:
          const: enabled
          enum:
            - enabled
          title: Type
          type: string
      required:
        - budget_tokens
        - type
      title: Enabled
      type: object
    Tool:
      additionalProperties: false
      properties:
        type:
          anyOf:
            - type: 'null'
            - const: custom
              enum:
                - custom
              type: string
          title: Type
        description:
          description: >-
            Description of what this tool does.


            Tool descriptions should be as detailed as possible. The more
            information that the model has about what the tool is and how to use
            it, the better it will perform. You can use natural language
            descriptions to reinforce important aspects of the tool input JSON
            schema.
          examples:
            - Get the current weather in a given location
          title: Description
          type: string
        name:
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          maxLength: 128
          minLength: 1
          pattern: ^[a-zA-Z0-9_-]{1,128}$
          title: Name
          type: string
        input_schema:
          $ref: '#/components/schemas/InputSchema'
          description: >-
            [JSON schema](https://json-schema.org/draft/2020-12) for this tool's
            input.


            This defines the shape of the `input` that your tool accepts and
            that the model will produce.
          examples:
            - properties:
                location:
                  description: The city and state, e.g. San Francisco, CA
                  type: string
                unit:
                  description: Unit for the output - one of (celsius, fahrenheit)
                  type: string
              required:
                - location
              type: object
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
      required:
        - name
        - input_schema
      title: Custom tool
      type: object
    ToolChoiceAny:
      additionalProperties: false
      description: The model will use any available tools.
      properties:
        disable_parallel_tool_use:
          description: >-
            Whether to disable parallel tool use.


            Defaults to `false`. If set to `true`, the model will output exactly
            one tool use.
          title: Disable Parallel Tool Use
          type: boolean
        type:
          const: any
          enum:
            - any
          title: Type
          type: string
      required:
        - type
      title: Any
      type: object
    ToolChoiceAuto:
      additionalProperties: false
      description: The model will automatically decide whether to use tools.
      properties:
        disable_parallel_tool_use:
          description: >-
            Whether to disable parallel tool use.


            Defaults to `false`. If set to `true`, the model will output at most
            one tool use.
          title: Disable Parallel Tool Use
          type: boolean
        type:
          const: auto
          enum:
            - auto
          title: Type
          type: string
      required:
        - type
      title: Auto
      type: object
    ToolChoiceNone:
      additionalProperties: false
      description: The model will not be allowed to use tools.
      properties:
        type:
          const: none
          enum:
            - none
          title: Type
          type: string
      required:
        - type
      title: None
      type: object
    ToolChoiceTool:
      additionalProperties: false
      description: The model will use the specified tool with `tool_choice.name`.
      properties:
        disable_parallel_tool_use:
          description: >-
            Whether to disable parallel tool use.


            Defaults to `false`. If set to `true`, the model will output exactly
            one tool use.
          title: Disable Parallel Tool Use
          type: boolean
        name:
          description: The name of the tool to use.
          title: Name
          type: string
        type:
          const: tool
          enum:
            - tool
          title: Type
          type: string
      required:
        - name
        - type
      title: Tool
      type: object
    URLImageSource:
      additionalProperties: false
      properties:
        type:
          const: url
          enum:
            - url
          title: Type
          type: string
        url:
          title: Url
          type: string
      required:
        - type
        - url
      title: URLImageSource
      type: object
    URLPDFSource:
      additionalProperties: false
      properties:
        type:
          const: url
          enum:
            - url
          title: Type
          type: string
        url:
          title: Url
          type: string
      required:
        - type
        - url
      title: PDF (URL)
      type: object
    Usage:
      properties:
        cache_creation:
          anyOf:
            - $ref: '#/components/schemas/CacheCreation'
            - type: 'null'
          default: null
          description: Breakdown of cached tokens by TTL
        cache_creation_input_tokens:
          anyOf:
            - minimum: 0
              type: integer
            - type: 'null'
          default: null
          description: The number of input tokens used to create the cache entry.
          examples:
            - 2051
          title: Cache Creation Input Tokens
        cache_read_input_tokens:
          anyOf:
            - minimum: 0
              type: integer
            - type: 'null'
          default: null
          description: The number of input tokens read from the cache.
          examples:
            - 2051
          title: Cache Read Input Tokens
        input_tokens:
          description: The number of input tokens which were used.
          examples:
            - 2095
          minimum: 0
          title: Input Tokens
          type: integer
        output_tokens:
          description: The number of output tokens which were used.
          examples:
            - 503
          minimum: 0
          title: Output Tokens
          type: integer
        server_tool_use:
          anyOf:
            - $ref: '#/components/schemas/ServerToolUsage'
            - type: 'null'
          default: null
          description: The number of server tool requests.
        service_tier:
          anyOf:
            - enum:
                - standard
                - priority
                - batch
              type: string
            - type: 'null'
          default: null
          description: If the request used the priority, standard, or batch tier.
          title: Service Tier
      required:
        - cache_creation
        - cache_creation_input_tokens
        - cache_read_input_tokens
        - input_tokens
        - output_tokens
        - server_tool_use
        - service_tier
      title: Usage
      type: object
    UserLocation:
      additionalProperties: false
      properties:
        city:
          anyOf:
            - maxLength: 255
              minLength: 1
              type: string
            - type: 'null'
          description: The city of the user.
          examples:
            - New York
            - Tokyo
            - Los Angeles
          title: City
        country:
          anyOf:
            - maxLength: 2
              minLength: 2
              type: string
            - type: 'null'
          description: >-
            The two letter [ISO country
            code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the user.
          examples:
            - US
            - JP
            - GB
          title: Country
        region:
          anyOf:
            - maxLength: 255
              minLength: 1
              type: string
            - type: 'null'
          description: The region of the user.
          examples:
            - California
            - Ontario
            - Wales
          title: Region
        timezone:
          anyOf:
            - maxLength: 255
              minLength: 1
              type: string
            - type: 'null'
          description: The [IANA timezone](https://nodatime.org/TimeZones) of the user.
          examples:
            - America/New_York
            - Asia/Tokyo
            - Europe/London
          title: Timezone
        type:
          const: approximate
          enum:
            - approximate
          title: Type
          type: string
      required:
        - type
      title: UserLocation
      type: object
    WebSearchToolResultErrorCode:
      enum:
        - invalid_tool_input
        - unavailable
        - max_uses_exceeded
        - too_many_requests
        - query_too_long
      title: WebSearchToolResultErrorCode
      type: string
    WebSearchTool_20250305:
      additionalProperties: false
      properties:
        allowed_domains:
          anyOf:
            - items:
                type: string
              type: array
            - type: 'null'
          description: >-
            If provided, only these domains will be included in results. Cannot
            be used alongside `blocked_domains`.
          title: Allowed Domains
        blocked_domains:
          anyOf:
            - items:
                type: string
              type: array
            - type: 'null'
          description: >-
            If provided, these domains will never appear in results. Cannot be
            used alongside `allowed_domains`.
          title: Blocked Domains
        cache_control:
          anyOf:
            - discriminator:
                mapping:
                  ephemeral: '#/components/schemas/CacheControlEphemeral'
                propertyName: type
              oneOf:
                - $ref: '#/components/schemas/CacheControlEphemeral'
            - type: 'null'
          description: Create a cache control breakpoint at this content block.
          title: Cache Control
        max_uses:
          anyOf:
            - exclusiveMinimum: 0
              type: integer
            - type: 'null'
          description: Maximum number of times the tool can be used in the API request.
          title: Max Uses
        name:
          const: web_search
          description: >-
            Name of the tool.


            This is how the tool will be called by the model and in `tool_use`
            blocks.
          enum:
            - web_search
          title: Name
          type: string
        type:
          const: web_search_20250305
          enum:
            - web_search_20250305
          title: Type
          type: string
        user_location:
          anyOf:
            - $ref: '#/components/schemas/UserLocation'
            - type: 'null'
          description: >-
            Parameters for the user's location. Used to provide more relevant
            search results.
      required:
        - name
        - type
      title: Web search tool (2025-03-05)
      type: object

````

# Vision

> The Claude 3 and 4 families of models comes with new vision capabilities that allow Claude to understand and analyze images, opening up exciting possibilities for multimodal interaction.

This guide describes how to work with images in Claude, including best practices, code examples, and limitations to keep in mind.

***

## How to use vision

Use Claude’s vision capabilities via:

* [claude.ai](https://claude.ai/). Upload an image like you would a file, or drag and drop an image directly into the chat window.
* The [Console Workbench](https://console.anthropic.com/workbench/). If you select a model that accepts images (Claude 3 and 4 models only), a button to add images appears at the top right of every User message block.
* **API request**. See the examples in this guide.

***

## Before you upload

### Basics and Limits

You can include multiple images in a single request (up to 20 for [claude.ai](https://claude.ai/) and 100 for API requests). Claude will analyze all provided images when formulating its response. This can be helpful for comparing or contrasting images.

If you submit an image larger than 8000x8000 px, it will be rejected. If you submit more than 20 images in one API request, this limit is 2000x2000 px.

### Evaluate image size

For optimal performance, we recommend resizing images before uploading if they are too large. If your image’s long edge is more than 1568 pixels, or your image is more than \~1,600 tokens, it will first be scaled down, preserving aspect ratio, until it’s within the size limits.

If your input image is too large and needs to be resized, it will increase latency of [time-to-first-token](/en/docs/about-claude/glossary), without giving you any additional model performance. Very small images under 200 pixels on any given edge may degrade performance.

<Tip>
  To improve [time-to-first-token](/en/docs/about-claude/glossary), we recommend
  resizing images to no more than 1.15 megapixels (and within 1568 pixels in
  both dimensions).
</Tip>

Here is a table of maximum image sizes accepted by our API that will not be resized for common aspect ratios. With the Claude Sonnet 3.7 model, these images use approximately 1,600 tokens and around \$4.80/1K images.

| Aspect ratio | Image size   |
| ------------ | ------------ |
| 1:1          | 1092x1092 px |
| 3:4          | 951x1268 px  |
| 2:3          | 896x1344 px  |
| 9:16         | 819x1456 px  |
| 1:2          | 784x1568 px  |

### Calculate image costs

Each image you include in a request to Claude counts towards your token usage. To calculate the approximate cost, multiply the approximate number of image tokens by the [per-token price of the model](https://anthropic.com/pricing) you’re using.

If your image does not need to be resized, you can estimate the number of tokens used through this algorithm: `tokens = (width px * height px)/750`

Here are examples of approximate tokenization and costs for different image sizes within our API’s size constraints based on Claude Sonnet 3.7 per-token price of \$3 per million input tokens:

| Image size                    | # of Tokens | Cost / image | Cost / 1K images |
| ----------------------------- | ----------- | ------------ | ---------------- |
| 200x200 px(0.04 megapixels)   | \~54        | \~\$0.00016  | \~\$0.16         |
| 1000x1000 px(1 megapixel)     | \~1334      | \~\$0.004    | \~\$4.00         |
| 1092x1092 px(1.19 megapixels) | \~1590      | \~\$0.0048   | \~\$4.80         |

### Ensuring image quality

When providing images to Claude, keep the following in mind for best results:

* **Image format**: Use a supported image format: JPEG, PNG, GIF, or WebP.
* **Image clarity**: Ensure images are clear and not too blurry or pixelated.
* **Text**: If the image contains important text, make sure it’s legible and not too small. Avoid cropping out key visual context just to enlarge the text.

***

## Prompt examples

Many of the [prompting techniques](/en/docs/build-with-claude/prompt-engineering/overview) that work well for text-based interactions with Claude can also be applied to image-based prompts.

These examples demonstrate best practice prompt structures involving images.

<Tip>
  Just as with document-query placement, Claude works best when images come
  before text. Images placed after text or interpolated with text will still
  perform well, but if your use case allows it, we recommend an image-then-text
  structure.
</Tip>

### About the prompt examples

The following examples demonstrate how to use Claude's vision capabilities using various programming languages and approaches. You can provide images to Claude in three ways:

1. As a base64-encoded image in `image` content blocks
2. As a URL reference to an image hosted online
3. Using the Files API (upload once, use multiple times)

The base64 example prompts use these variables:

<CodeGroup>
  ```bash Shell
      # For URL-based images, you can use the URL directly in your JSON request

      # For base64-encoded images, you need to first encode the image
      # Example of how to encode an image to base64 in bash:
      BASE64_IMAGE_DATA=$(curl -s "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg" | base64)

      # The encoded data can now be used in your API calls
  ```

  ```Python Python
  import base64
  import httpx

  # For base64-encoded images
  image1_url = "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg"
  image1_media_type = "image/jpeg"
  image1_data = base64.standard_b64encode(httpx.get(image1_url).content).decode("utf-8")

  image2_url = "https://upload.wikimedia.org/wikipedia/commons/b/b5/Iridescent.green.sweat.bee1.jpg"
  image2_media_type = "image/jpeg"
  image2_data = base64.standard_b64encode(httpx.get(image2_url).content).decode("utf-8")

  # For URL-based images, you can use the URLs directly in your requests
  ```

  ```TypeScript TypeScript
  import axios from 'axios';

  // For base64-encoded images
  async function getBase64Image(url: string): Promise<string> {
    const response = await axios.get(url, { responseType: 'arraybuffer' });
    return Buffer.from(response.data, 'binary').toString('base64');
  }

  // Usage
  async function prepareImages() {
    const imageData = await getBase64Image('https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg');
    // Now you can use imageData in your API calls
  }

  // For URL-based images, you can use the URLs directly in your requests
  ```

  ```java Java
  import java.io.IOException;
  import java.util.Base64;
  import java.io.InputStream;
  import java.net.URL;

  public class ImageHandlingExample {

      public static void main(String[] args) throws IOException, InterruptedException {
          // For base64-encoded images
          String image1Url = "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg";
          String image1MediaType = "image/jpeg";
          String image1Data = downloadAndEncodeImage(image1Url);

          String image2Url = "https://upload.wikimedia.org/wikipedia/commons/b/b5/Iridescent.green.sweat.bee1.jpg";
          String image2MediaType = "image/jpeg";
          String image2Data = downloadAndEncodeImage(image2Url);

          // For URL-based images, you can use the URLs directly in your requests
      }

      private static String downloadAndEncodeImage(String imageUrl) throws IOException {
          try (InputStream inputStream = new URL(imageUrl).openStream()) {
              return Base64.getEncoder().encodeToString(inputStream.readAllBytes());
          }
      }

  }
  ```
</CodeGroup>

Below are examples of how to include images in a Messages API request using base64-encoded images and URL references:

### Base64-encoded image example

<CodeGroup>
  ```bash Shell
  curl https://api.anthropic.com/v1/messages \
    -H "x-api-key: $ANTHROPIC_API_KEY" \
    -H "anthropic-version: 2023-06-01" \
    -H "content-type: application/json" \
    -d '{
      "model": "claude-sonnet-4-20250514",
      "max_tokens": 1024,
      "messages": [
        {
          "role": "user",
          "content": [
            {
              "type": "image",
              "source": {
                "type": "base64",
                "media_type": "image/jpeg",
                "data": "'"$BASE64_IMAGE_DATA"'"
              }
            },
            {
              "type": "text",
              "text": "Describe this image."
            }
          ]
        }
      ]
    }'
  ```

  ```Python Python
  import anthropic

  client = anthropic.Anthropic()
  message = client.messages.create(
      model="claude-sonnet-4-20250514",
      max_tokens=1024,
      messages=[
          {
              "role": "user",
              "content": [
                  {
                      "type": "image",
                      "source": {
                          "type": "base64",
                          "media_type": image1_media_type,
                          "data": image1_data,
                      },
                  },
                  {
                      "type": "text",
                      "text": "Describe this image."
                  }
              ],
          }
      ],
  )
  print(message)
  ```

  ```TypeScript TypeScript
  import Anthropic from '@anthropic-ai/sdk';

  const anthropic = new Anthropic({
    apiKey: process.env.ANTHROPIC_API_KEY,
  });

  async function main() {
    const message = await anthropic.messages.create({
      model: "claude-sonnet-4-20250514",
      max_tokens: 1024,
      messages: [
        {
          role: "user",
          content: [
            {
              type: "image",
              source: {
                type: "base64",
                media_type: "image/jpeg",
                data: imageData, // Base64-encoded image data as string
              }
            },
            {
              type: "text",
              text: "Describe this image."
            }
          ]
        }
      ]
    });

    console.log(message);
  }

  main();
  ```

  ```Java Java
  import java.io.IOException;
  import java.util.List;

  import com.anthropic.client.AnthropicClient;
  import com.anthropic.client.okhttp.AnthropicOkHttpClient;
  import com.anthropic.models.messages.*;

  public class VisionExample {
      public static void main(String[] args) throws IOException, InterruptedException {
          AnthropicClient client = AnthropicOkHttpClient.fromEnv();
          String imageData = ""; // // Base64-encoded image data as string

          List<ContentBlockParam> contentBlockParams = List.of(
                  ContentBlockParam.ofImage(
                          ImageBlockParam.builder()
                                  .source(Base64ImageSource.builder()
                                          .data(imageData)
                                          .build())
                                  .build()
                  ),
                  ContentBlockParam.ofText(TextBlockParam.builder()
                          .text("Describe this image.")
                          .build())
          );
          Message message = client.messages().create(
                  MessageCreateParams.builder()
                          .model(Model.CLAUDE_3_7_SONNET_LATEST)
                          .maxTokens(1024)
                          .addUserMessageOfBlockParams(contentBlockParams)
                          .build()
          );

          System.out.println(message);
      }
  }
  ```
</CodeGroup>

### URL-based image example

<CodeGroup>
  ```bash Shell
  curl https://api.anthropic.com/v1/messages \
    -H "x-api-key: $ANTHROPIC_API_KEY" \
    -H "anthropic-version: 2023-06-01" \
    -H "content-type: application/json" \
    -d '{
      "model": "claude-sonnet-4-20250514",
      "max_tokens": 1024,
      "messages": [
        {
          "role": "user",
          "content": [
            {
              "type": "image",
              "source": {
                "type": "url",
                "url": "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg"
              }
            },
            {
              "type": "text",
              "text": "Describe this image."
            }
          ]
        }
      ]
    }'
  ```

  ```Python Python
  import anthropic

  client = anthropic.Anthropic()
  message = client.messages.create(
      model="claude-sonnet-4-20250514",
      max_tokens=1024,
      messages=[
          {
              "role": "user",
              "content": [
                  {
                      "type": "image",
                      "source": {
                          "type": "url",
                          "url": "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg",
                      },
                  },
                  {
                      "type": "text",
                      "text": "Describe this image."
                  }
              ],
          }
      ],
  )
  print(message)
  ```

  ```TypeScript TypeScript
  import Anthropic from '@anthropic-ai/sdk';

  const anthropic = new Anthropic({
    apiKey: process.env.ANTHROPIC_API_KEY,
  });

  async function main() {
    const message = await anthropic.messages.create({
      model: "claude-sonnet-4-20250514",
      max_tokens: 1024,
      messages: [
        {
          role: "user",
          content: [
            {
              type: "image",
              source: {
                type: "url",
                url: "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg"
              }
            },
            {
              type: "text",
              text: "Describe this image."
            }
          ]
        }
      ]
    });

    console.log(message);
  }

  main();
  ```

  ```Java Java
  import java.io.IOException;
  import java.util.List;

  import com.anthropic.client.AnthropicClient;
  import com.anthropic.client.okhttp.AnthropicOkHttpClient;
  import com.anthropic.models.messages.*;

  public class VisionExample {

      public static void main(String[] args) throws IOException, InterruptedException {
          AnthropicClient client = AnthropicOkHttpClient.fromEnv();

          List<ContentBlockParam> contentBlockParams = List.of(
                  ContentBlockParam.ofImage(
                          ImageBlockParam.builder()
                                  .source(UrlImageSource.builder()
                                          .url("https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg")
                                          .build())
                                  .build()
                  ),
                  ContentBlockParam.ofText(TextBlockParam.builder()
                          .text("Describe this image.")
                          .build())
          );
          Message message = client.messages().create(
                  MessageCreateParams.builder()
                          .model(Model.CLAUDE_3_7_SONNET_LATEST)
                          .maxTokens(1024)
                          .addUserMessageOfBlockParams(contentBlockParams)
                          .build()
          );
          System.out.println(message);
      }
  }
  ```
</CodeGroup>

### Files API image example

For images you'll use repeatedly or when you want to avoid encoding overhead, use the [Files API](/en/docs/build-with-claude/files):

<CodeGroup>
  ```bash Shell
  # First, upload your image to the Files API
  curl -X POST https://api.anthropic.com/v1/files \
    -H "x-api-key: $ANTHROPIC_API_KEY" \
    -H "anthropic-version: 2023-06-01" \
    -H "anthropic-beta: files-api-2025-04-14" \
    -F "file=@image.jpg"

  # Then use the returned file_id in your message
  curl https://api.anthropic.com/v1/messages \
    -H "x-api-key: $ANTHROPIC_API_KEY" \
    -H "anthropic-version: 2023-06-01" \
    -H "anthropic-beta: files-api-2025-04-14" \
    -H "content-type: application/json" \
    -d '{
      "model": "claude-sonnet-4-20250514",
      "max_tokens": 1024,
      "messages": [
        {
          "role": "user",
          "content": [
            {
              "type": "image",
              "source": {
                "type": "file",
                "file_id": "file_abc123"
              }
            },
            {
              "type": "text",
              "text": "Describe this image."
            }
          ]
        }
      ]
    }'
  ```

  ```python Python
  import anthropic

  client = anthropic.Anthropic()

  # Upload the image file
  with open("image.jpg", "rb") as f:
      file_upload = client.beta.files.upload(file=("image.jpg", f, "image/jpeg"))

  # Use the uploaded file in a message
  message = client.beta.messages.create(
      model="claude-sonnet-4-20250514",
      max_tokens=1024,
      betas=["files-api-2025-04-14"],
      messages=[
          {
              "role": "user",
              "content": [
                  {
                      "type": "image",
                      "source": {
                          "type": "file",
                          "file_id": file_upload.id
                      }
                  },
                  {
                      "type": "text",
                      "text": "Describe this image."
                  }
              ]
          }
      ],
  )

  print(message.content)
  ```

  ```typescript TypeScript
  import { Anthropic, toFile } from '@anthropic-ai/sdk';
  import fs from 'fs';

  const anthropic = new Anthropic();

  async function main() {
    // Upload the image file
    const fileUpload = await anthropic.beta.files.upload({
      file: toFile(fs.createReadStream('image.jpg'), undefined, { type: "image/jpeg" })
    }, {
      betas: ['files-api-2025-04-14']
    });

    // Use the uploaded file in a message
    const response = await anthropic.beta.messages.create({
      model: 'claude-sonnet-4-20250514',
      max_tokens: 1024,
      betas: ['files-api-2025-04-14'],
      messages: [
        {
          role: 'user',
          content: [
            {
              type: 'image',
              source: {
                type: 'file',
                file_id: fileUpload.id
              }
            },
            {
              type: 'text',
              text: 'Describe this image.'
            }
          ]
        }
      ]
    });

    console.log(response);
  }

  main();
  ```

  ```java Java
  import java.io.IOException;
  import java.nio.file.Files;
  import java.nio.file.Path;
  import java.util.List;

  import com.anthropic.client.AnthropicClient;
  import com.anthropic.client.okhttp.AnthropicOkHttpClient;
  import com.anthropic.models.File;
  import com.anthropic.models.files.FileUploadParams;
  import com.anthropic.models.messages.*;

  public class ImageFilesExample {
      public static void main(String[] args) throws IOException {
          AnthropicClient client = AnthropicOkHttpClient.fromEnv();

          // Upload the image file
          File file = client.beta().files().upload(FileUploadParams.builder()
                  .file(Files.newInputStream(Path.of("image.jpg")))
                  .build());

          // Use the uploaded file in a message
          ImageBlockParam imageParam = ImageBlockParam.builder()
                  .fileSource(file.id())
                  .build();

          MessageCreateParams params = MessageCreateParams.builder()
                  .model(Model.CLAUDE_3_7_SONNET_LATEST)
                  .maxTokens(1024)
                  .addUserMessageOfBlockParams(
                          List.of(
                                  ContentBlockParam.ofImage(imageParam),
                                  ContentBlockParam.ofText(
                                          TextBlockParam.builder()
                                                  .text("Describe this image.")
                                                  .build()
                                  )
                          )
                  )
                  .build();

          Message message = client.messages().create(params);
          System.out.println(message.content());
      }
  }
  ```
</CodeGroup>

See [Messages API examples](/en/api/messages) for more example code and parameter details.

<AccordionGroup>
  <Accordion title="Example: One image">
    It’s best to place images earlier in the prompt than questions about them or instructions for tasks that use them.

    Ask Claude to describe one image.

    | Role | Content                       |
    | ---- | ----------------------------- |
    | User | \[Image] Describe this image. |

    Here is the corresponding API call using the Claude Sonnet 3.7 model.

    <Tabs>
      <Tab title="Using Base64">
        ```Python Python
        message = client.messages.create(
            model="claude-sonnet-4-20250514",
            max_tokens=1024,
            messages=[
                {
                    "role": "user",
                    "content": [
                        {
                            "type": "image",
                            "source": {
                                "type": "base64",
                                "media_type": image1_media_type,
                                "data": image1_data,
                            },
                        },
                        {
                            "type": "text",
                            "text": "Describe this image."
                        }
                    ],
                }
            ],
        )
        ```
      </Tab>

      <Tab title="Using URL">
        ```Python Python
        message = client.messages.create(
            model="claude-sonnet-4-20250514",
            max_tokens=1024,
            messages=[
                {
                    "role": "user",
                    "content": [
                        {
                            "type": "image",
                            "source": {
                                "type": "url",
                                "url": "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg",
                            },
                        },
                        {
                            "type": "text",
                            "text": "Describe this image."
                        }
                    ],
                }
            ],
        )
        ```
      </Tab>
    </Tabs>
  </Accordion>

  <Accordion title="Example: Multiple images">
    In situations where there are multiple images, introduce each image with `Image 1:` and `Image 2:` and so on. You don’t need newlines between images or between images and the prompt.

    Ask Claude to describe the differences between multiple images.

    | Role | Content                                                                 |
    | ---- | ----------------------------------------------------------------------- |
    | User | Image 1: \[Image 1] Image 2: \[Image 2] How are these images different? |

    Here is the corresponding API call using the Claude Sonnet 3.7 model.

    <Tabs>
      <Tab title="Using Base64">
        ```Python Python
        message = client.messages.create(
            model="claude-sonnet-4-20250514",
            max_tokens=1024,
            messages=[
                {
                    "role": "user",
                    "content": [
                        {
                            "type": "text",
                            "text": "Image 1:"
                        },
                        {
                            "type": "image",
                            "source": {
                                "type": "base64",
                                "media_type": image1_media_type,
                                "data": image1_data,
                            },
                        },
                        {
                            "type": "text",
                            "text": "Image 2:"
                        },
                        {
                            "type": "image",
                            "source": {
                                "type": "base64",
                                "media_type": image2_media_type,
                                "data": image2_data,
                            },
                        },
                        {
                            "type": "text",
                            "text": "How are these images different?"
                        }
                    ],
                }
            ],
        )
        ```
      </Tab>

      <Tab title="Using URL">
        ```Python Python
        message = client.messages.create(
            model="claude-sonnet-4-20250514",
            max_tokens=1024,
            messages=[
                {
                    "role": "user",
                    "content": [
                        {
                            "type": "text",
                            "text": "Image 1:"
                        },
                        {
                            "type": "image",
                            "source": {
                                "type": "url",
                                "url": "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg",
                            },
                        },
                        {
                            "type": "text",
                            "text": "Image 2:"
                        },
                        {
                            "type": "image",
                            "source": {
                                "type": "url",
                                "url": "https://upload.wikimedia.org/wikipedia/commons/b/b5/Iridescent.green.sweat.bee1.jpg",
                            },
                        },
                        {
                            "type": "text",
                            "text": "How are these images different?"
                        }
                    ],
                }
            ],
        )
        ```
      </Tab>
    </Tabs>
  </Accordion>

  <Accordion title="Example: Multiple images with a system prompt">
    Ask Claude to describe the differences between multiple images, while giving it a system prompt for how to respond.

    | Content |                                                                         |
    | ------- | ----------------------------------------------------------------------- |
    | System  | Respond only in Spanish.                                                |
    | User    | Image 1: \[Image 1] Image 2: \[Image 2] How are these images different? |

    Here is the corresponding API call using the Claude Sonnet 3.7 model.

    <Tabs>
      <Tab title="Using Base64">
        ```Python Python
        message = client.messages.create(
            model="claude-sonnet-4-20250514",
            max_tokens=1024,
            system="Respond only in Spanish.",
            messages=[
                {
                    "role": "user",
                    "content": [
                        {
                            "type": "text",
                            "text": "Image 1:"
                        },
                        {
                            "type": "image",
                            "source": {
                                "type": "base64",
                                "media_type": image1_media_type,
                                "data": image1_data,
                            },
                        },
                        {
                            "type": "text",
                            "text": "Image 2:"
                        },
                        {
                            "type": "image",
                            "source": {
                                "type": "base64",
                                "media_type": image2_media_type,
                                "data": image2_data,
                            },
                        },
                        {
                            "type": "text",
                            "text": "How are these images different?"
                        }
                    ],
                }
            ],
        )
        ```
      </Tab>

      <Tab title="Using URL">
        ```Python Python
        message = client.messages.create(
            model="claude-sonnet-4-20250514",
            max_tokens=1024,
            system="Respond only in Spanish.",
            messages=[
                {
                    "role": "user",
                    "content": [
                        {
                            "type": "text",
                            "text": "Image 1:"
                        },
                        {
                            "type": "image",
                            "source": {
                                "type": "url",
                                "url": "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg",
                            },
                        },
                        {
                            "type": "text",
                            "text": "Image 2:"
                        },
                        {
                            "type": "image",
                            "source": {
                                "type": "url",
                                "url": "https://upload.wikimedia.org/wikipedia/commons/b/b5/Iridescent.green.sweat.bee1.jpg",
                            },
                        },
                        {
                            "type": "text",
                            "text": "How are these images different?"
                        }
                    ],
                }
            ],
        )
        ```
      </Tab>
    </Tabs>
  </Accordion>

  <Accordion title="Example: Four images across two conversation turns">
    Claude’s vision capabilities shine in multimodal conversations that mix images and text. You can have extended back-and-forth exchanges with Claude, adding new images or follow-up questions at any point. This enables powerful workflows for iterative image analysis, comparison, or combining visuals with other knowledge.

    Ask Claude to contrast two images, then ask a follow-up question comparing the first images to two new images.

    | Role      | Content                                                                            |
    | --------- | ---------------------------------------------------------------------------------- |
    | User      | Image 1: \[Image 1] Image 2: \[Image 2] How are these images different?            |
    | Assistant | \[Claude's response]                                                               |
    | User      | Image 1: \[Image 3] Image 2: \[Image 4] Are these images similar to the first two? |
    | Assistant | \[Claude's response]                                                               |

    When using the API, simply insert new images into the array of Messages in the `user` role as part of any standard [multiturn conversation](/en/api/messages-examples#multiple-conversational-turns) structure.
  </Accordion>
</AccordionGroup>

***

## Limitations

While Claude's image understanding capabilities are cutting-edge, there are some limitations to be aware of:

* **People identification**: Claude [cannot be used](https://www.anthropic.com/legal/aup) to identify (i.e., name) people in images and will refuse to do so.
* **Accuracy**: Claude may hallucinate or make mistakes when interpreting low-quality, rotated, or very small images under 200 pixels.
* **Spatial reasoning**: Claude's spatial reasoning abilities are limited. It may struggle with tasks requiring precise localization or layouts, like reading an analog clock face or describing exact positions of chess pieces.
* **Counting**: Claude can give approximate counts of objects in an image but may not always be precisely accurate, especially with large numbers of small objects.
* **AI generated images**: Claude does not know if an image is AI-generated and may be incorrect if asked. Do not rely on it to detect fake or synthetic images.
* **Inappropriate content**: Claude will not process inappropriate or explicit images that violate our [Acceptable Use Policy](https://www.anthropic.com/legal/aup).
* **Healthcare applications**: While Claude can analyze general medical images, it is not designed to interpret complex diagnostic scans such as CTs or MRIs. Claude's outputs should not be considered a substitute for professional medical advice or diagnosis.

Always carefully review and verify Claude's image interpretations, especially for high-stakes use cases. Do not use Claude for tasks requiring perfect precision or sensitive image analysis without human oversight.

***

## FAQ

<AccordionGroup>
  <Accordion title="What image file types does Claude support?">
    Claude currently supports JPEG, PNG, GIF, and WebP image formats, specifically:

    * `image/jpeg`
    * `image/png`
    * `image/gif`
    * `image/webp`
  </Accordion>

  {" "}

  <Accordion title="Can Claude read image URLs?">
    Yes, Claude can now process images from URLs with our URL image source blocks in the API.
    Simply use the "url" source type instead of "base64" in your API requests.
    Example:

    ```json
    {
      "type": "image",
      "source": {
        "type": "url",
        "url": "https://upload.wikimedia.org/wikipedia/commons/a/a7/Camponotus_flavomarginatus_ant.jpg"
      }
    }
    ```
  </Accordion>

  <Accordion title="Is there a limit to the image file size I can upload?">
    Yes, there are limits:

    * API: Maximum 5MB per image
    * claude.ai: Maximum 10MB per image

    Images larger than these limits will be rejected and return an error when using our API.
  </Accordion>

  <Accordion title="How many images can I include in one request?">
    The image limits are:

    * Messages API: Up to 100 images per request
    * claude.ai: Up to 20 images per turn

    Requests exceeding these limits will be rejected and return an error.
  </Accordion>

  {" "}

  <Accordion title="Does Claude read image metadata?">
    No, Claude does not parse or receive any metadata from images passed to it.
  </Accordion>

  {" "}

  <Accordion title="Can I delete images I've uploaded?">
    No. Image uploads are ephemeral and not stored beyond the duration of the API
    request. Uploaded images are automatically deleted after they have been
    processed.
  </Accordion>

  {" "}

  <Accordion title="Where can I find details on data privacy for image uploads?">
    Please refer to our privacy policy page for information on how we handle
    uploaded images and other data. We do not use uploaded images to train our
    models.
  </Accordion>

  <Accordion title="What if Claude's image interpretation seems wrong?">
    If Claude's image interpretation seems incorrect:

    1. Ensure the image is clear, high-quality, and correctly oriented.
    2. Try prompt engineering techniques to improve results.
    3. If the issue persists, flag the output in claude.ai (thumbs up/down) or contact our support team.

    Your feedback helps us improve!
  </Accordion>

  <Accordion title="Can Claude generate or edit images?">
    No, Claude is an image understanding model only. It can interpret and analyze images, but it cannot generate, produce, edit, manipulate, or create images.
  </Accordion>
</AccordionGroup>

***

## Dive deeper into vision

Ready to start building with images using Claude? Here are a few helpful resources:

* [Multimodal cookbook](https://github.com/anthropics/anthropic-cookbook/tree/main/multimodal): This cookbook has tips on [getting started with images](https://github.com/anthropics/anthropic-cookbook/blob/main/multimodal/getting%5Fstarted%5Fwith%5Fvision.ipynb) and [best practice techniques](https://github.com/anthropics/anthropic-cookbook/blob/main/multimodal/best%5Fpractices%5Ffor%5Fvision.ipynb) to ensure the highest quality performance with images. See how you can effectively prompt Claude with images to carry out tasks such as [interpreting and analyzing charts](https://github.com/anthropics/anthropic-cookbook/blob/main/multimodal/reading%5Fcharts%5Fgraphs%5Fpowerpoints.ipynb) or [extracting content from forms](https://github.com/anthropics/anthropic-cookbook/blob/main/multimodal/how%5Fto%5Ftranscribe%5Ftext.ipynb).
* [API reference](/en/api/messages): Visit our documentation for the Messages API, including example [API calls involving images](/en/api/messages-examples).

If you have any other questions, feel free to reach out to our [support team](https://support.anthropic.com/). You can also join our [developer community](https://www.anthropic.com/discord) to connect with other creators and get help from Anthropic experts.
