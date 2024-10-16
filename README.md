# Real-time Messaging App (ChatChat)

This project is a real-time messaging app, similar to LINE, built using Flutter for the frontend, Go for the backend, and SQL Server for the database. The app supports real-time communication through WebSockets (Gorilla WebSocket).

## Technologies Used

- **Frontend**: Flutter
- **Backend**: Go
- **Database**: SQL Server
- **Real-time Communication**: WebSockets (Gorilla WebSocket)

## Key Features

- **Real-time Messaging**: Send and receive messages instantly using WebSocket connections.
- **User Management**: Register and login functionality.
- **Chat History**: Store chat history in SQL Server for future retrieval.
- **Instant Notifications**: Get notified immediately of new messages via WebSocket.

## System Architecture

- **Frontend (Flutter)**: The user interface for sending and receiving messages.
- **Backend (Go)**: Manages WebSocket communication and database operations.
- **Database (SQL Server)**: Stores user data and chat history.

## Real-time Communication

- Utilizes **WebSocket** to maintain real-time communication between the Flutter app and Go server.
- **Gorilla WebSocket** is used on the Go server to manage WebSocket connections.

## Installation and Setup

### 1. Frontend (Flutter)

- Install Flutter SDK.
- Run the Flutter app on an emulator or a real device.

### 2. Backend (Go)

- Install Go.
- Set up the WebSocket server using Gorilla WebSocket.

### 3. Database (SQL Server)

- Set up SQL Server to store user data and chat history.

## How It Works

1. Users send messages through the Flutter app.
2. Messages are transmitted to the Go server via WebSocket.
3. The server processes the message and broadcasts it to the intended recipient in real-time.
4. Chat history is stored in SQL Server for future reference.

## Future Development

- Add image and file-sharing capabilities.
- Integrate push notifications for new messages.

---

