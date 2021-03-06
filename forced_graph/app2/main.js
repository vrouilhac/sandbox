const { app, BrowserWindow } = require("electron");
const path = require("path");
const url = require("url");

const URL = "http://localhost:3000";

let mainWindow;

const createWindow = () => {
	const windowOptions = {
		width: 800,
		height: 600,
	};

	mainWindow = new BrowserWindow(windowOptions);
	mainWindow.loadURL(URL);
	mainWindow.webContents.openDevTools();
	mainWindow.on("closed", () => {
		mainWindow = null;
	});
};

app.on("ready", createWindow);

app.on("window-all-closed", function () {
	// On OS X it is common for applications and their menu bar
	// to stay active until the user quits explicitly with Cmd + Q
	if (process.platform !== "darwin") {
		app.quit();
	}
});

app.on("activate", function () {
	// On OS X it's common to re-create a window in the app when the
	// dock icon is clicked and there are no other windows open.
	if (mainWindow === null) {
		createWindow();
	}
});
