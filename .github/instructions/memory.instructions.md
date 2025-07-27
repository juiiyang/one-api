---
applyTo: "**/*"
---


# Project Memory & Handover Instructions

## Claude Messages API: Universal Conversion

- All Claude Messages API requests (`/v1/messages`) are routed to adapters that implement `ConvertClaudeRequest(c, request)` and `DoResponse(c, resp, meta)`. Anthropic uses native passthrough; most others use OpenAI-compatible or custom conversion.
- Conversion state is tracked via context keys in `common/ctxkey/key.go` (e.g., `ClaudeMessagesConversion`, `ConvertedResponse`, `OriginalClaudeRequest`).
- Data mapping is bi-directional: Claude → OpenAI (system, messages, tools, tool_choice, etc.) and OpenAI → Claude (choices, tool_calls, finish_reason, usage, etc.). Gemini and some others use custom mapping.
- Full support for function calling, streaming, structured/multimodal content, and tool use. Billing, quota, and token counting are handled identically to ChatCompletion, including image token calculation and fallback strategies.
- All errors are wrapped with `github.com/Laisky/errors/v2` and surfaced with context. Malformed content is handled gracefully with fallbacks.
- New adapters should follow the Claude Messages pattern: interface method + universal conversion + context marking. Specialized adapters (e.g., DeepL, Palm, Ollama) are excluded from Claude Messages support.
- See `relay/controller/claude_messages.go`, `relay/adaptor/interface.go`, `relay/adaptor/openai_compatible/claude_messages.go`, `relay/adaptor/gemini/adaptor.go`, `common/ctxkey/key.go`, and `docs/arch/api_convert.md` for reference.

## Pricing & Billing Architecture (2025-07)

- **Pricing Unit Standardization:** All model pricing, quota, and billing calculations are now standardized to use "per 1M tokens" (1 million tokens) instead of "per 1K tokens". This is reflected in all code, comments, and documentation. Double-check all user-facing messages and documentation for consistency.
- **Centralized Model Pricing:** Each channel/adaptor now imports and uses a shared `ModelRatios` constant from its respective `constants.go` or subadaptor. Local, hardcoded pricing maps have been removed to avoid duplication and drift.
- **Model List Generation:** Supported model lists are always derived from the keys of the shared pricing maps, ensuring pricing and support are always in sync.
- **Default/Fallback Pricing:** All adaptors use a unified fallback (e.g., `5 * ratio.MilliTokensUsd`) for unknown models. If a model is missing from the shared map, it will use this fallback.
- **VertexAI Aggregation:** VertexAI pricing is now aggregated from all subadaptors (Claude, Imagen, Gemini, Veo) and includes VertexAI-specific models. Any omission in a subadaptor will propagate to VertexAI.
- **Critical Subtleties:**
  - If any model is missing from the shared pricing map, it may become unsupported or use fallback pricing.
  - Models with non-token-based pricing (e.g., per image/video) require special handling and may not fit the token-based pattern.
  - All documentation and UI must be kept in sync with the new pricing unit to avoid confusion.

## Gemini Adapter: Function Schema Cleaning

- Gemini API rejects OpenAI-style function schemas with unsupported fields (`additionalProperties`, `description`, `strict`).
- Recursive cleaning removes `additionalProperties` everywhere, and `description`/`strict` only at the top level. Cleaned parameters are type-asserted before assignment.
- Only remove `description`/`strict` at the top; nested objects may require them.

## General Project Practices

- **Error Handling:** Always use `github.com/Laisky/errors/v2` for error wrapping; never return bare errors.
- **Context Keys:** All context keys must be pre-defined in `common/ctxkey/key.go`.
- **Package Management:** Use package managers (npm, pip, etc.), never edit package files by hand.
- **Testing:** All bug fixes/features must be covered by unit tests. No temporary scripts. Unit tests must be updated to cover new issues and features.
- **Time Handling:** Always use UTC for server, DB, and API time.
- **Golang ORM:** Use `gorm.io/gorm` for writes; prefer SQL for reads to minimize DB load.

## Handover Guidance

- **Claude Messages API:** Fully production-ready, with universal conversion and billing parity. See `docs/arch/api_billing.md` and `docs/arch/api_convert.md` for details.
- **Billing Architecture:** Four-layer pricing (channel overrides > adapter defaults > global > fallback).
- **Adaptor Pattern:** All new API formats should follow the Claude Messages pattern: interface method + universal conversion + context marking.
- **Critical Files:**
  - `relay/controller/claude_messages.go`
  - `relay/adaptor/interface.go`
  - `common/ctxkey/key.go`
  - `docs/arch/api_billing.md`
  - `docs/arch/api_convert.md`

---
**Recent Developments (2025-07):**

- Major refactor to unify and clarify model pricing logic, reduce duplication, and standardize on "per 1M tokens" as the pricing unit. All adaptors now use shared pricing maps and fallback logic. This change is critical for maintainability and billing accuracy.
- When handing over, ensure the new assistant is aware of the pricing unit change, the centralized pricing logic, and the importance of keeping documentation and UI in sync with backend logic.
