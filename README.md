# Gototp
Module to generate time-based one-time passwords

Experimental for now, only works with a single account secret at a time.

## Usage
NEVER share secrets! This is just experimental.

`-register <secret>` Registers the secret token from your third-party provider  
`-generate` Outputs a time-based one-time code for 2FA

## Storage
NEVER share this file.  
Secret is stored in the current user's home directory.  
On Unix / MacOS - `$HOME/.totp_2fa.json`.  
On Windows - `%USERPROFILE%/.totp_2fa.json`.