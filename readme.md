# SSH Connection Management Program

## Introduction

The SSH Connection Management Program is a web application that allows users to efficiently manage and interact with SSH connections. Users can register on the website to gain access to the various features of the program.

## Features

### User Registration and Authentication

- Users can register on the website to create an account.
- Passwords are securely stored using encryption techniques.
- Only the user with their password can encrypt and decrypt their stored password.

### SSH Connection Creation

- Once logged in, users can create SSH connections to remote servers.
- Connection details such as username and password are saved in a Postgres database.
- Passwords are securely decrypted when needed for authentication.

### Batch Commands

- Users can execute commands and perform file transfers (copying files to/from remote servers) using a simple batch UI.
- Commands can be executed in parallel across multiple SSH connections.
- Created batches can be saved in the database for future use.

### Batch Output Viewing

- The program provides a convenient way to view the output of executed batches.
- Output is categorized based on the selected tab in the UI.

### SSH Session Management

- Users can easily open an SSH session through the shell tab.
- Clicking on a saved SSH connection allows users to log into the remote server directly.

### Multiple SSH Connection Management

- Users can manage multiple SSH connections.
- New SSH connections can be added by clicking on the "Add" tab.

## Technology Stack

- Backend: Go
- Frontend: React with TypeScript
- Database: Postgres

## Usage

To use the SSH Connection Management Program:
1. Register or log in to the website.
2. Create SSH connections with connection details.
3. Use the batch UI to execute commands and transfer files.
4. View batch outputs based on selected tabs.
5. Open SSH sessions through the shell tab.
6. Manage multiple SSH connections by adding new connections.

## Conclusion

The SSH Connection Management Program simplifies SSH connection management by offering a user-friendly web interface and various features for efficient interaction with remote servers.

