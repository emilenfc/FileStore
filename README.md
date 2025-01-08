# FileStore

FileStore is a secure file uploading and management system that allows users to store and access files within user-specific directories. It uses API credentials to authenticate users and ensures their data is safe and organized. This document explains how the system works and how to use its API endpoints.

---

## Features

- **Secure Uploads**: Authenticated users can upload files to their private folders.
- **Folder Management**: Automatically creates and manages folders for users.
- **File Retrieval**: Users can access their files using secure URLs.
- **User  Authentication**: API keys and secrets authenticate and manage users.

---

# FileStore API Endpoints

Below is a comprehensive list of all API endpoints provided by the FileStore application, including request details and expected responses.

---

## Public Endpoints

### 1. **Register a User**
   - **URL**: `/register`
   - **Method**: `POST`
   - **Description**: Register a new user and generate API credentials (`X-API-Key` and `X-API-Secret`).
   - **Request Body**:
     ```json
     {
         "email": "user@example.com",
         "password": "password123"
     }
     ```
   - **Response**:
     ```json
     {
         "message": "User registered successfully",
         "api_key": "USER_API_KEY",
         "api_secret": "USER_API_SECRET"
     }
     ```

---

### 2. **Login**
   - **URL**: `/login`
   - **Method**: `POST`
   - **Description**: Login an existing user and retrieve API credentials.
   - **Request Body**:
     ```json
     {
         "email": "user@example.com",
         "password": "password123"
     }
     ```
   - **Response**:
     ```json
     {
         "message": "Login successful",
         "api_key": "USER_API_KEY",
         "api_secret": "USER_API_SECRET"
     }
     ```

---

## Protected Endpoints (Require Authentication)

  **NOTE**:all this end point require you to login first

### 3. **Get User Info**
   - **URL**: `/api/user`
   - **Method**: `GET`
   - **Description**: Retrieve details about the authenticated user.
   - **Response**:
     ```json
     {
         "id": 1,
         "email": "user@example.com",
         "folders": [...]
     }
     ```

---

### 4. **Regenerate API Secret**
   - **URL**: `/api/regenerate-secret`
   - **Method**: `POST`
   - **Description**: Generate a new API secret for the authenticated user.
   - **Response**:
     ```json
     {
         "message": "API secret regenerated",
         "api_secret": "NEW_API_SECRET"
     }
     ```

---

### 5. **List Folders**
   - **URL**: `/api/folders`
   - **Method**: `GET`
   - **Description**: Retrieve a list of folders belonging to the authenticated user.
   - **Response**:
     ```json
     [
         {
             "id": 1,
             "name": "FOLDER_NAME",
             "created_at": "2025-01-01T12:00:00Z"
         }
     ]
     ```

---

### 6. **Get Folder Contents**
   - **URL**: `/api/folders/:folder`
   - **Method**: `GET`
   - **Description**: Retrieve the contents of a specific folder.
   - **Path Parameter**:
     - `:folder`: The name of the folder to retrieve.
   - **Response**:
     ```json
     [
       {
            "name": "string",
            "url": "string",
            "created_at": "datetime",
            "size": "number"
       }
     ]
     
  ```

---

## File Upload and Retrieval

### 7. **Upload File**
**Note**: Add the following headers to this endpoint requests:
- `X-API-Key`: User's API key
- `X-API-Secret`: User's API secret
---
   - **URL**: `/upload`
   - **Method**: `POST`
   - **Description**: Upload a file to a specific folder.
   - **Headers**:
     - `X-API-Key`: User's API key
     - `X-API-Secret`: User's API secret
   - **Form Data**:
     - `folder`: The name of the folder to upload the file to.
     - `file`: The file to upload.
   - **Response**:
     ```json
     {
         "message": "File uploaded successfully",
         "file_name": "FILE_NAME",
         "dir": "FOLDER_NAME",
         "file_url": "http://localhost:8085/uploads/USER_API_KEY/FOLDER_NAME/FILE_NAME",
         "full_path": "/uploads/USER_API_KEY/FOLDER_NAME/FILE_NAME"
     }
     ```

    - **Example Request**:
    ```bash
    curl -X POST \
    http://localhost:8085/upload \
    -H 'Content-Type: multipart/form-data' \
    -H 'X-API-Key: USER_API_KEY' \
    -H 'X-API-Secret: USER_API_SECRET' \
    -F 'folder=FOLDER_NAME' \
    -F 'file=@/path/to/your/file.txt'
    ```
---

### 8. **Retrieve File**
   - **URL**: `/uploads/*path`
   - **Method**: `GET`
   - **Description**: Retrieve a file by its unique public URL.
   - **Path Parameter**:
     - `*path`: The path to the file, including the API key, folder name, and file name.
   - **Response**: The requested file is served directly for download.

---

## Error Responses

For all endpoints, the following error structures may apply:

- **Unauthorized**:
  ```json
  {
      "error": "Invalid API credentials"
  }
