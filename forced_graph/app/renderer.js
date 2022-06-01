const root = document.querySelector("#root");

const data = {
  nodes: [{ id: 1, name: "Vincent", val: 1 }],
  links: []
};

const graph = ForceGraph();
graph(root).graphData(data);
