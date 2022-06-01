// TODO: put archive note in one color
// TODO: put reference note in another color
// TODO: put tags in another color


const express = require('express');
const http = require('http');
const fs = require('fs');
const path = require('path');
const EventEmitter = require('events');
const chokidar = require('chokidar');
const socketIO = require('socket.io');
const app = express();
const { initWatcher } = require("./src/watcher");

const server = http.createServer(app);
const io = socketIO(server, {
  cors: {
    origin: "http://localhost:3000",
    methods: ["POST", "GET"]
  }
});


initWatcher();

// TODO: update this to not take into account ctags but actual folder content
const file = fs.readFileSync(path.resolve(__dirname, "zettelkasten", "tags"), "utf8");
const split = file.split('\n').filter(file => !(!file || file.startsWith('!'))).map(file => {
  const s = file.split('\t')
  const regex = /[ar]_[0-9a-z]/;
  const v = {
    noteId: s[0],
    noteName: s[1].split('/')[1],
    refId: s[2]?.match(regex)[0],
    type: s[3],
  }
  // console.log({ v });
  return v;
});
// console.log({ split });

const linksFound = [];

const colors = {
  archive: "#5DADE2",
  references: "#27AE60",
  tags: "#5D6D7E",
};

const formatDataToGraphData = (rawData) => {
  // console.log({ rawData });
  return rawData.map((file, index) => {
    const rfile = fs.readFileSync(path.resolve(__dirname, "zettelkasten", file.type === "r" ? "references" : "slipbox", file.noteName), "utf8");
    const regex = /@[ar]_[0-9a-z]*/g;
    const matched = rfile.match(regex);
     // console.log({ matched, rfile });

    if(!!matched) {
      matched.forEach(m => {
	linksFound.push({
	  source: file.noteId,
	  target: m.replace('@', '')
	});
      });
    }

    return {
      val: 1,
      id: file.noteId,
      name: file.noteName,
      color: file.type === "r" ? colors.references : colors.archive,
      fileContent: rfile
    };
  });
};

const findNoteId = (file) => {
  const id = file.split("\n").find(line => line.match(/ID/)).split(':')[1].trim();
  return id.replace('@', '');
};

const findNoteReferences = (file) => {
    const regex = /@[ar]_[0-9a-z]*/g;
    const matched = file.match(regex);
    return matched || [];
};

const buildGraphLinks = (id, links) => {
  const graphLinks = links.map(link => ({ source: id, target: link.replace('@', '') }));
  return graphLinks;
};



const fileWatcherEvent = new EventEmitter();
const ref_dirWatcher = chokidar.watch('zettelkasten/references', { ignored: /^.*\.swp/ });
const slip_dirWatcher = chokidar.watch('zettelkasten/slipbox', { ignored: /^.*\.swp/ });

const defaultNodes = formatDataToGraphData(split);

const updateGraphNode = (graphNode) => {
  const oldNode = defaultNodes.find(node => node.id === graphNode.id);
  const updatedGraphNode = defaultNodes.filter(node => node.id !== graphNode.id);
  updatedGraphNode.push(graphNode);
  return updatedGraphNode;
};


const update = (path) => {
  const file = fs.readFileSync(path, "utf8");
  const noteId = findNoteId(file);
  const refs = findNoteReferences(file);
  const graphLinks = buildGraphLinks(noteId, refs);
  const graphNode = {
    id: noteId,
    name: path,
    val: 1,
    fileContent: file
  };
  const updatedNodes = updateGraphNode(graphNode);
  const data = { nodes: updatedNodes, links: [...linksFound, ...graphLinks ] }
  io.emit('FILE_CHANGED', data);
};

slip_dirWatcher.on('change', update);
slip_dirWatcher.on('unlink', update);
slip_dirWatcher.on('add', update);

app.get('/', (req, res, next) => {
  fileWatcherEvent.emit('ev');

  const nodes = JSON.stringify(defaultNodes);
  const links = JSON.stringify(linksFound);
  const html = `
    <!DOCTYPE html>
    <html lang="en">
      <head>
	<title>Vincent</title>
	<script src="//unpkg.com/force-graph"></script>
	<script src="/socket.io/socket.io.js"></script>
      </head>
      <body style="background: #212121; position: relative">
	<div id="text-container" style="position: absolute; z-index: 999; top: 16px; left: 16px; background: #F2F3F4; border-radius: 8px"><p id="text-content" style="white-space: pre-line; font-size: 11px; padding: 8px"></p></div>
	<div id="root"></div>
	<script type="module">

	const domElement = document.querySelector("#root");
	let OPENED_NODE = null;

	const data = {
	  nodes: ${nodes},
	  links: ${links}
	};
	
	const myGraph = ForceGraph();

	myGraph(domElement).graphData(data).onNodeClick((node, event) => {
	  // console.log(node, event);
	  const textContent = document.querySelector("#text-content");
	  textContent.textContent = node.fileContent;
	  OPENED_NODE = node.id;
	}).linkColor(() => "#ffffff");

	const socket = io();

	socket.on("FILE_CHANGED", (msg) => {
	  const { data } = msg;
	  const newData = {
	    nodes: data.graphNodes,
	    links: data.graphLinks
	  };
	  console.log({ newData, oldData: myGraph.graphData() });
	  myGraph(domElement).graphData(newData)
	  if(!!OPENED_NODE) {
	    const node = data.graphNodes.find(n => n.id === OPENED_NODE);
	    const textContent = document.querySelector("#text-content");
	    textContent.textContent = node?.fileContent || "";
	  }
	});

	</script>
      </body>
    </html>
  `;
  res.send(html);
});

server.listen(8000, () => {
  console.log("Server started at http://localhost:8000");
});







const initData = () => {
  const zettelsFileName = fs.readdirSync(path.resolve(__dirname, "zettelkasten", "slipbox"), "utf8");
  const referencesFileName = fs.readdirSync(path.resolve(__dirname, "zettelkasten", "references"), "utf8");
  const zettels = [];
  const references = [];
  const links = [];

  zettelsFileName.forEach(fileName => {
    console.log({ fileName });
    const fileContent = fs.readFileSync(path.resolve(__dirname, "zettelkasten", "slipbox", fileName), "utf8");
    const noteId = `a_${fileName.split("_")[0]}`;

    zettels.push({
      id: noteId,
      name: `a.${fileName}`,
      color: colors.archive,
      val: 1,
      fileContent
    });

    const regex = /@[ar]_[0-9a-z]*/g;
    const matched = fileContent.match(regex);

    if(!!matched) {
      matched.forEach(m => {
	links.push({
	  source: noteId,
	  target: m.replace('@', '')
	});
      });
    }
  });

  referencesFileName.forEach(fileName => {
    const fileContent = fs.readFileSync(path.resolve(__dirname, "zettelkasten", "references", fileName), "utf8");
    console.log({ fileName });
    const noteId = `r_${fileName.split("_")[0]}`;

    references.push({
      id: noteId,
      name: `r.${fileName}`,
      color: colors.references,
      val: 1,
      fileContent
    });
  });

  return {
    nodes: [...zettels, ...references],
    links,
  };
};

// console.log(initData());

io.on('connection', (socket) => {
  const initializedData = initData();
  // console.log({ initializedData });
  socket.emit("INITIALIZATION", initializedData);
});
