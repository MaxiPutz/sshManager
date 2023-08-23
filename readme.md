# SSH Connection Management Program

## Introduction

The SSH Connection Management Program is a web application that allows users to efficiently manage and interact with SSH connections. Users can register on the website to gain access to the various features of the program.

## Features
![shell](https://github.com/MaxiPutz/sshManager/assets/48091139/7f82479b-ae54-4959-8298-ca8923b1ed27)![shell_fast](https://github.com/MaxiPutz/sshManager/assets/48091139/401e2257-fcd1-448e-a008-4ca44c67ff46)

![batch](https://github.com/MaxiPutz/sshManager/assets/48091139/77b64d90-aef0-4e89-a9c7-b10dd0c08a30)![batch_fast](https://github.com/MaxiPutz/sshManager/assets/48091139/46c6d3ff-bd11-4787-9816-2faef70d8df3)


### User Registration and Authentication![register_fast](https://github.com/MaxiPutz/sshManager/assets/48091139/150a1432-cce8-4270-8a53-2b16a2631a67)

![register](https://github.com/MaxiPutz/sshManager/assets/48091139/2cec4b91-5c2a-49d0-9190-6f66ae979339)
![login](https://github.com/MaxiPutz/sshManager/assets/48091139/07a61ce7-4fb9-445b-ac8a-fde31603454a)![login_fast](https://github.com/MaxiPutz/sshManager/assets/48091139/0ccfd24b-231e-4fd9-b2c1-3053c9865b7e)


- Users can register on the website to create an account.
- Passwords are securely stored using encryption techniques.
- Only the user with their password can encrypt and decrypt their stored password.

### SSH Connection Creation



- Once logged in, users can create SSH connections to remote servers.
- Connection details such as username and password are saved in a Postgres database.
- Passwords are securely decrypted when needed for authentication.

### Batch Commands
![batch](https://github.com/MaxiPutz/sshManager/assets/48091139/7cd47e79-78c2-4291-86ea-c22790992d22)![login_fast](https://github.com/MaxiPutz/sshManager/assets/48091139/ef5cdb24-c8d8-4b40-b4f5-7d9a5dabde62)


- Users can execute commands and perform file transfers (copying files to/from remote servers) using a simple batch UI.
- Commands can be executed in parallel across multiple SSH connections.
- Created batches can be saved in the database for future use.

### Batch Output Viewing


- The program provides a convenient way to view the output of executed batches.
- Output is categorized based on the selected tab in the UI.

### SSH Session Management
![shell](https://github.com/MaxiPutz/sshManager/assets/48091139/1b22ea43-2af1-4322-bf0e-45dd05649e26)![shell_fast](https://github.com/MaxiPutz/sshManager/assets/48091139/d2938656-3569-4f33-a406-b6c152f27d9c)


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

