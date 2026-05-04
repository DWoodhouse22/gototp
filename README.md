# Gototp
Generate time-based one-time passwords

## Install

### Build from source
`go build -o gototp`

## Commands

- `gototp use [group]`
  - [docs](#use)  
- `gototp register <name> <secret> [--group <group>]`
  - [docs](#register)
- `gototp generate <name> [--group <group> --copy]`
  - [docs](#generate)
- `gototp remove <name> [--group <group> --force]`
  - [docs](#remove)
- `gototp list`
  - [docs](#list)

## Use
Switch context to a group for future register and generate commands
### Usage
```
gotop use [group]
```

- `[group]` Optional group name (e.g. 'work', 'personal' etc)

### Group does not exist
```
gototp use work
Current group: personal

Group 'work' does not exist.
Create and switch to it? (y/N): n
```

### Group exists
```
gototp use personal
Current group: work

Switched to group 'personal'
```

### No group specified
```text
gototp use
Current group: personal

Available groups:
1) personal
2) work

Select a group:
```
## Register
Register a new account
### Usage
```
gototp register <name> <secret> [--group <group>]
```

- `<name>`: the account name (e.g. 'github', 'google' etc)
- `<secret>`: the secret token provided by the account provider
- `--group`, `-g`: Optional group name (e.g. 'work', 'personal' etc)

Names must be unique within a group

## Generate
Generate a one-time password for a registered account
### Usage
```
gototp generate <name> [--group <group>]
```

- `<name>`: the account name (e.g. 'github', 'google' etc)
- `--group`, `-g`: Optional group name (e.g. 'work', 'personal' etc)
- `--copy`, `-c`: Optional, copies the generated code to your clipboard

### Single match
If only one account exists with the given name:

```
gototp generate github
123456
```

### Multiple matches (no group specified)

If multiple accounts exist with the same name across different groups, you will be prompted to choose:
```
Multiple accounts found for 'github':

1) github (work)
2) github (personal)

Select an option:
```
Enter the number corresponding to the account to generate the code.

### Specify group directly

To skip the selection prompt, provide a group:

```
gototp generate github --group work
123456
```

## Remove
Remove a registered account
### Usage
```
gototp remove <name> [--group <group> --force]
```
- `<name>`: The account name (e.g. 'github'. 'google' etc)
- `--group`, `-g`: Optional group name (e.g. 'work', 'personal' etc)
- `--force`, `-f`: Optional, force removal and skip confirmation

## List
List all registered accounts
### Usage
```
gototp list
personal:
  github
  google

work:
  github
  google
```

## Storage
NEVER share this file.  
This tool stores secrets in plain text JSON.  
Do not use on shared machines or production environments.

Secret is stored in the current user's home directory.  
On Unix / MacOS - `$HOME/.totp_2fa.json`.  
On Windows - `%USERPROFILE%/.totp_2fa.json`.
