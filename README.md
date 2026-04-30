# Gototp
Generate time-based one-time passwords

## Install

### Build from source
`go build -o gototp`

## Commands
```bash
gototp register <name> <secret> [--group <group>]
gototp generate <name> [--group <group>]
gototp list
```

## Register
Register a new account
### Usage
```bash
gototp register <name> <secret> [--group <group>]
```

- `<name>`: the account name (e.g. 'github', 'google' etc)
- `<secret>`: the secret token provided by the account provider
- `--group`, `-g`: Optional group name (e.g. 'work', 'personal' etc)

Names bust me unique within a group

## Generate
Generate a one-time password for a registered account
### Usage
```bash
gototp generate <name> [--group <group>]
```

- `<name>`: the account name (e.g. 'github', 'google' etc)
- `--group`, `-g`: Optional group name (e.g. 'work', 'personal' etc)

### Single match
If only one account exists with the given name:

```bash
gototp generate github
123456
```

### Multiple matches (no group specified)

If multiple accounts exist with the same name across different groups, you will be prompted to choose:
```text
Multiple accounts found for 'github':

1) github (work)
2) github (personal)

Select an option:
```
Enter the number corresponding to the account to generate the code.

### Specify group directly

To skip the selection prompt, provide a group:

```bash
gototp generate github --group work
123456
```

## List
List all registered accounts
### Usage
```bash
gototp list
personal:
  github
  google

work:
  google
  github
```


## Storage
NEVER share this file.  
This tool stores secrets in plain text JSON.  
Do not use on shared machines or production environments.

Secret is stored in the current user's home directory.  
On Unix / MacOS - `$HOME/.totp_2fa.json`.  
On Windows - `%USERPROFILE%/.totp_2fa.json`.