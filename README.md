# CamelMailer with Go

This example shows how to use [CamelMailer](https://camelmailer.com) with [Go](https://go.dev) through the [camelmailer-go](https://github.com/camelmailer/camelmailer-go) SDK: send an email, then read the server's message stats.

## Prerequisites

- Go 1.21+
- A CamelMailer server API key (dashboard → your server → **Credentials** → new credential of type **API**)

## Instructions

1. Set your environment:

   ```sh
   export CAMELMAILER_API_KEY="cm_xxxx"
   # Self-hosted instance? Point the SDK at it (defaults to https://app.camelmailer.com):
   export CAMELMAILER_BASE_URL="https://mail.example.com"
   # Sender/recipient used by the example:
   export CAMELMAILER_FROM="you@yourdomain.com"
   export CAMELMAILER_TO="delivered@example.com"
   ```

2. Run it:

   ```sh
   go run .
   ```

## License

MIT License
