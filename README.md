# Scratch Pad

A note-taking application made in [Go](https://go.dev) and [Pico.css](https://picocss.com).

| Home Page                                                                                                      | Show Page                                                                                                      | Create Scratch Form                                                                                            |
|----------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------|
| ![image](https://user-images.githubusercontent.com/6123841/149582325-4d5c9274-84d7-478b-88b7-ba85abee82d4.png) | ![image](https://user-images.githubusercontent.com/6123841/149582567-21b9d4df-230f-4625-9c40-a4caad9b112e.png) | ![image](https://user-images.githubusercontent.com/6123841/149582414-1fabfd28-c3fc-4aa9-9e06-55af611d161d.png) |

## Setup

Setup localhost TLS certificates by following instruction inside `/tls` directory.

MariaDB is used for database. Either use docker compose container or a local installation.

Before starting application you need to create additional tables and indices. When using docker compose a provisioning
script will be executed setting up the db.

<details>
<summary>SQL Init</summary>

```sql
CREATE DATABASE scratchpad CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE DATABASE test_scratchpad CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE scratchpad;

CREATE TABLE scratches
(
    id      INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title   VARCHAR(100) NOT NULL,
    content TEXT         NOT NULL,
    created DATETIME     NOT NULL,
    expires DATETIME     NOT NULL
);

CREATE INDEX idx_scratches_created ON scratches (created);

INSERT INTO scratches (title, content, created, expires)
VALUES ('One World',
        'The view of the earth from the moon fascinated me - a small disk, 240,000 miles away. It was hard to think that that little thing held so many problems, so many frustrations. Raging nationalistic interests, famines, wars, pestilence don\'t show from that distance. I\'m convinced that some wayward stranger in a space-craft, coming from some other part of the heavens, could look at earth and never know that it was inhabited at all. But the same wayward stranger would certainly know instinctively that if the earth were inhabited, then the destinies of all who lived on it must inevitably be interwoven and joined. We are one hunk of ground, water, air, clouds, floating around in space. From out there it really is \'one world\'.\n\n??? Frank Borman',
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY));

INSERT INTO scratches (title, content, created, expires)
VALUES ('Magnificent Desolation',
        'I believe that space travel will one day become as common as airline travel is today. I???m convinced, however, that the true future of space travel does not lie with government agencies ??? NASA is still obsessed with the idea that the primary purpose of the space program is science ??? but real progress will come from private companies competing to provide the ultimate adventure ride, and NASA will receive the trickle-down benefits.\n\n??? Buzz Aldrin',
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY));

CREATE TABLE users
(
    id              INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name            VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    hashed_password CHAR(60)     NOT NULL,
    created         DATETIME     NOT NULL,
    active          BOOLEAN      NOT NULL DEFAULT TRUE
);

ALTER TABLE users
    ADD CONSTRAINT users_uc_email UNIQUE (email);
```

</details>

## Run

Build and run app on port :4000.

```bash
go build -o ./ ./...
./web
```

## Dev

Install [Air](https://github.com/cosmtrek/air) for easy development and live-reloading.

```bash
air init

# adjust .air.toml config, might need to update:
# bin = "./tmp/web"
# cmd = "go build -o ./tmp/ ./..."

air
```

