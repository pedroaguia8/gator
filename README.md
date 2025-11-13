# gator 🐊

**gator** (as in, *aggre-gator*) is a command-line RSS feed aggregator.

It's a multi-user CLI tool that allows users on the same machine to add, follow, and read posts from their favorite RSS feeds, all from the comfort of the terminal. All aggregated posts are stored in a local PostgreSQL database.

## Features

- Add RSS feeds to be collected.
- Store collected posts in a PostgreSQL database.
- Follow and unfollow RSS feeds that other users have added.
- View summaries of aggregated posts, with a link to the full post.
- Run a long-running service (`gator agg`) to continuously fetch new posts.

## Prerequisites

Before you begin, you will need:

- **Git**: For cloning the repository.
- **Go**: Version 1.21 or later.
- **PostgreSQL**: Version 15 or later.

## Installation

1. Clone the repository:
```bash
git clone https://github.com/pedroaguia8/gator.git
cd gator
```

2. Install the binary:
```bash
go install
```

This will compile the `gator` binary and place it in your `$GOPATH/bin` directory (e.g., `~/go/bin/gator`).

3. **Ensure it's in your PATH**: Make sure your `$GOPATH/bin` directory is in your system's PATH environment variable so you can run `gator` from anywhere.

## Setup

### 1. Database Setup

You must have a PostgreSQL server running.

1. Connect to `psql`:
```bash
# On macOS
psql postgres

# On Linux
sudo -u postgres psql
```

2. Create the database:
```sql
CREATE DATABASE gator;
```

### 2. Configuration File

`gator` uses a JSON config file to store database credentials and the current user.

1. Manually create the file `~/.gatorconfig.json`:
```bash
touch ~/.gatorconfig.json
```

2. Open the file and add your database connection string.

**Connection String Format**: `postgres://USERNAME:PASSWORD@HOST:PORT/DATABASE?sslmode=disable`

**Examples:**

- **macOS** (default, no password): `postgres://your_username:@localhost:5432/gator?sslmode=disable`
- **Linux** (default 'postgres' user/pass): `postgres://postgres:postgres@localhost:5432/gator?sslmode=disable`

Your `~/.gatorconfig.json` should look like this:
```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

*The `current_user_name` field will be added automatically when you register or log in.*

### 3. Database Migrations

Before you can use the app, you need to set up the database schema. This project uses **goose** for migrations.

1. Install Goose:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

2. Run "Up" Migrations: From the root of the gator project directory, run:
```bash
# You'll need to pass goose your connection string
# (Note: goose doesn't need the ?sslmode=disable part)
goose -dir "sql/schema" postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

This will create all the necessary tables (users, feeds, posts, etc.).

## Usage

Now you can run the `gator` CLI from your terminal.

### First-Time Use: Register a User
```bash
# Register your first user
gator register your_username
```

This creates the user in the database and automatically logs you in by updating `~/.gatorconfig.json`.

### Common Commands

- **Register a new user**:
```bash
  gator register <username>
```

- **Log in as an existing user**:
```bash
  gator login <username>
```

- **List all registered users**:
```bash
  gator users
```

- **Add a new feed**: (This also automatically follows the feed for you.)
```bash
  gator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
```

- **Follow a feed someone else added**:
```bash
  gator follow https://techcrunch.com/feed/
```

- **List all feeds you are following**:
```bash
  gator following
```

- **Unfollow a feed**:
```bash
  gator unfollow https://techcrunch.com/feed/
```

- **Browse your posts**: (Shows the latest 2 posts by default.)
```bash
  gator browse
  
  # Show the 10 latest posts
  gator browse 10
```

### Running the Aggregator

This is the most important command. It's a long-running service that continuously fetches feeds. You should run this in its own terminal window (or a tmux session).
```bash
# Fetch all feeds every 1 minute
gator agg 1m

# Fetch all feeds every 10 minutes
gator agg 10m
```

The aggregator will loop forever, finding the least-recently-fetched feed, fetching its posts, saving new ones to the database, and then sleeping until the next cycle.

## Development

- **Run for development**: `go run . <command>`
- **Generate SQLC code**: `sqlc generate` (run from project root)
