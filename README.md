# Creative Dream Bot

A Telegram bot for AI image generation with support for HD/4K quality output.

## Features

- 🎨 AI-powered image generation (HD/4K quality)
- 💰 User credit system with coin management
- 💳 Integrated payments support
- 🌐 Multi-language support (English/Russian)
- 🎟️ Promo code system
- ✅ Subscription verification

## Tech Stack

- **Language**: Go 1.23
- **Database**: PostgreSQL
- **Image Generation**: WebSocket API connection

## Configuration

Create a `.env` file with the following variables:

```env
DSN=postgres://user:password@host:port/database
CHANNEL_ID=your_channel_id
TG_API=your_telegram_bot_token
WEBSOCKET_API_KEY=your_websocket_api_key
WEBSOCKET_API_URL=your_websocket_api_url
```

## Getting Started

### Prerequisites

- Docker & Docker Compose
- PostgreSQL database (running separately)

### Running the Bot

1. Clone the repository
2. Configure your `.env` file
3. Start the bot:

```bash
docker compose up -d
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
