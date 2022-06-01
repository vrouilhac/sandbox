const chokidar = require('chokidar');
const path = require('path');

class Watcher {
  events = [];
  chokidarWatcher = null;

  constructor(props) {
    this.events = props?.events || [];
    this.chokidarWatcher = props?.chokidarWatcher || null;
  }

  watchDirectory(path, options) {
    this.chokidarWatcher = chokidar.watch(path, options);
    this.chokidarWatcher.on('change', (path) => {
      this.events.forEach(ev => ev(path));
    });

    return new Watcher({
      events: this.events,
      chokidarWatcher: this.chokidarWatcher,
    });
  }

  addEvent(watcherEvent) {
    this.events.push(watcherEvent); 
    return this;
  }
}


const initWatcher = () => {
  const slipboxPath = path.resolve(__dirname, "..", "zettelkasten", "slipbox");
  const options = {
    ignored: /^.*\.swp/
  };

 new Watcher()
   .watchDirectory(slipboxPath, options)
   .addEvent((path) => {
     console.log({ path });
   });
};

module.exports = {
  Watcher,
  initWatcher,
};
