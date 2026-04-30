# Gototp
Generate time-based one-time passwords

## Install

### Build from source
`go build -o gototp`

## Usage
Register a secret token  
`gototp register name secret`  
name is a label for the service (e.g github, google)  
secret must be Base32 encoded (as provided by most 2FA providers / QR codes)   

Output a time-based one-time code for 2FA  
`gototp generate name`  

List registered accounts / services  
`gototp list`

## Example
`gototp register github AHUVY3FPAPNK8GYL`  
`gototp generate github`  
output:
`123456`

## Storage
NEVER share this file.  
This tool stores secrets in plain text JSON.  
Do not use on shared machines or production environments.

Secret is stored in the current user's home directory.  
On Unix / MacOS - `$HOME/.totp_2fa.json`.  
On Windows - `%USERPROFILE%/.totp_2fa.json`.