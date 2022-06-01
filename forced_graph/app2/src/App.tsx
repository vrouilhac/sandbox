import React, { useEffect, useState } from 'react';
import './App.css';
import ForceGraph from 'force-graph';
import socketIOClient from "socket.io-client";

const API_URL = "http://127.0.0.1:8000";

const style = {
  backgroundColor: "#212121"
};

const styles: any = {
  textContent: {
    position: "absolute",
    top: 16,
    left: 16,
    fontSize: 10,
    zIndex: 999,
    backgroundColor: "#FFFFFF",
    borderRadius: 8,
    whiteSpace: "pre-line",
    padding: 8,
    boxShadow: "1px 2px 2px rgba(0, 0, 0, 0.14)"
  }
};

interface TextContentProps {
  text: string;
}

const TextContent: React.FC<TextContentProps> = (props) => {
  const there = (event: any) => {
    event.stopPropagation();
  };
  return (
    <p style={styles.textContent} onClick={there}>{props.text}</p>
  );
};

export const App: React.FC = () => {
  const [nodes, setNodes] = useState([]);
  const [links, setLinks] = useState([]);
  const [textContent, setTextContent] = useState<string | null>(null);

  useEffect(() => {
    const socket = socketIOClient(API_URL);

    socket.on("INITIALIZATION", (data) => {
      setNodes(data.nodes);
      setLinks(data.links);
    });

    socket.on("FILE_CHANGED", (data) => {
      setNodes(data.nodes);
      setLinks(data.links);
    });
  }, []);

  useEffect(() => {
    const elem = document.querySelector("#graph-root");
    if(elem) {
      ForceGraph()(elem as HTMLElement)
	.graphData({ nodes, links })
	.onNodeClick((node, event) => {
	  event.stopPropagation();
	  // @ts-ignore
	  setTextContent(node?.fileContent || null); 
	}).linkColor(() => "#AAB7B8");
    }
  }, [nodes, links]);

  const reset = () => {
    setTextContent(null);
  };
  return (
    <div onClick={reset}>
      {textContent && <TextContent text={textContent} />}
      <div id="graph-root" style={style}></div>
    </div>
  );
};

export default App;
