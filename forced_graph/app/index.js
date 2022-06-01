const { app, BrowserWindow } = require('electron');

const createWindow = (options = {}) => {
  const winOptions = {
    width: 800,
    height: 600,
    ...options,
  };
  const win = new BrowserWindow(winOptions);
  win.loadFile('index.html');
};

app.whenReady().then(() => {
  createWindow({
    webPreferences: {
      nodeIntegration: false
    }
  });
});
