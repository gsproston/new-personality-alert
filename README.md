# new-personality-alert
App that notifies people when a new personality is available to download.

## Setup
[Install Go](https://go.dev/doc/install).

Get a [TMDB account](https://www.themoviedb.org/settings/profile).

Set TMDB config:
1. Copy "config.go.example" in the "tmdb" directory.
2. Rename the file to "config.go".
3. Set the `TMDB_TOKEN` value to your TMDB API read access token, found in your [settings](https://www.themoviedb.org/settings/api).

Set email config:
1. Copy "config.go.example" in the "email" directory.
2. Rename the file to "config.go".
3. Set the values as follows:
   - `EMAIL_ADDRESS` - The email address that will send the emails.
   - `EMAIL_APP_PASSWORD` - The [app password](https://support.google.com/accounts/answer/185833?hl=en) for the email address.
   - `EMAIL_RECIPIENTS_CSL` - Comma separated list of the recipients' email addresses.
   - `EMAIL_AUTH_HOST` - The email authorisation host, e.g. "smtp.gmail.com".
   - `EMAIL_SEND_MAIL_ADDRESS` - The address and port that the emails will be sent from, e.g. "smtp.gmail.com:587".

## Run

### Run once
```bash
go run .
```

### Run periodically
Follow [these instructions](https://www.baeldung.com/linux/schedule-script-execution).
