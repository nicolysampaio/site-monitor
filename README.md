# Site Monitoring Tool

A simple and lightweight site monitoring tool written in Go. It checks the status of websites periodically and logs the results for easy monitoring.

## 🚀 Features

- Monitor website statuses in real time.
- Log the status of each website with timestamps.
- Display detailed logs.
- Easy configuration and setup.

## 📂 Project Structure

| File        | Description                                     |
| ----------- | ----------------------------------------------- |
| `main.go`   | Main program logic for monitoring and logging.  |
| `sites.txt` | List of websites to monitor (one URL per line). |
| `log.txt`   | Logs the monitoring results with timestamps.    |

## 🛠️ How to Use

1. **Install Go:**

   Make sure you have Go installed. You can download it from [here](https://golang.org/dl/).

2. **Run the Program:**

   Use the following command to start the tool:

   ```bash
   go run main.go
   ```

3. **Menu Options:**
   Once the program is running, you can choose from the following options:

   - **1**: Start monitoring the websites listed in `sites.txt`.
     - The program will periodically check the status of each website and log the results.
   - **2**: Display the logs stored in `log.txt`.
     - This will show all the monitoring logs with timestamps and the status of each website.
   - **0**: Exit the program.
     - Stops the execution and closes the application.

## ⚙️ Configuration

- **Monitoring Frequency**:  
  Customize the number of cycles and the delay between each cycle by modifying the constants in `main.go`.

- **Adding Websites**:  
  Add or update URLs in `sites.txt` to include websites for monitoring. Ensure each URL is written on a new line.

## 📝 Log Format

The monitoring results are logged in `log.txt` in the following format.

```sh
DD/MM/YYYY HH:MM:SS - <site_url> - online: <true/false>
```

Example log entry:

```sh
15/01/2025 12:30:45 - https://example.com - online: true
```

## 🌟 Why Use This Tool?

- Written in Go for efficiency and simplicity.
- Provides a minimalistic approach to site monitoring.
- Easily extensible and configurable.

## 📦 Requirements

- Go programming language (version 1.16 or later).

## 🤝 Contributions

Feel free to fork this repository, create new features, or report issues. Contributions are welcome!

## 📜 License

This project is licensed under the MIT License.
