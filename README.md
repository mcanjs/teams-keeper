# Teams Active Keeper 🟢

A simple and lightweight tool that prevents Microsoft Teams, Slack, or similar applications from showing your status as "Away" when you step away from your computer for a short time.

It keeps your computer and status active by moving your mouse a tiny, unnoticeable amount at time intervals you specify.

## 🚀 Installation

To run this project, you need to have **Go** installed on your computer. If you don't have it installed, you can download it from [here](https://go.dev/dl/).

1. Open your Terminal (Mac/Linux) or Command Prompt (Windows) and navigate to the project directory.
2. Run the following command to download the required libraries:
   ```bash
   go mod tidy
   ```
3. To build the application and make it ready to run:
   ```bash
   go build -o teams-keeper main.go
   ```

*(Alternatively, you can run it directly without building by typing `go run main.go`.)*

## 💻 Usage

When you start the application, it moves the mouse **every 4 minutes** by default.

To run it, type this command in your terminal:
```bash
./teams-keeper
```

### Changing the Time Interval (Optional)
If you want the mouse to move more or less frequently, you can use the `-time` parameter to set the time in minutes. For example, to move it **every 10 minutes**:
```bash
./teams-keeper -time 10
```

### How to Stop
When you want to close the application, simply press **`Ctrl + C`** (or **`Control + C`** on Mac) on your keyboard in the terminal window where it's running. Once closed, your Teams status will return to normal.

## 🛠️ Troubleshooting

**1. Error: "Application cannot move the mouse" or permission issues**
The application needs permission from your operating system to move your mouse.
- **Mac Users:** Go to `System Settings > Privacy & Security > Accessibility` and grant permission to your Terminal (or iTerm, VSCode, etc.).

**2. Error: "go command not found"**
Go is not installed on your computer, or it's missing from your PATH settings. Install Go and restart your terminal.

**3. The program is running but I still appear as "Away"**
The time interval you set might be too long. Try closing the application and lowering the time to 2 minutes by running `./teams-keeper -time 2`.

## 🔒 Security
This tool runs completely locally and does not send your data anywhere. It only reads your current mouse position, moves it 1 pixel to the right, and then brings it back.
