import { useState, useCallback } from 'react';
import { ReactFlow, applyNodeChanges, applyEdgeChanges, addEdge, NodeChange, EdgeChange} from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import Sidebar from '../components/Sidebar';
import SidebarTextField from '../components/SidebarTextField';
import SidebarArrayField from '../components/SidebarArrayField';
import SidebarMaxField from '../components/SidebarMaxField';

const SIDEBAR_HEADER = "Nodes"
const initialNodes = [
    {id: 'n1', position: {x: 0, y: 0}, data: {label: 'Node 1'}},
    {id: 'n2', position: {x: 0, y: 100}, data: {label: 'Node 2'}},
];

const initialEdges = [{id: 'n1-n2', source: 'n1', target: 'n2'}];

function StoryPage() {
    const [nodes, setNodes] = useState(initialNodes);
    const [edges, setEdges] = useState(initialEdges);
    const [name, setName] = useState('Opening');
    const [text, setText] = useState('You wake up at the edge of a quiet forest.');
    const [childrenIds, setChildrenIds] = useState(['n2']);
    const [stateChanges, setStateChanges] = useState<Record<string, string>>({awake: '1'});

    const onNodesChange = useCallback((changes: NodeChange<{
        id: string;
        position: { x: number; y: number; };
        data: { label: string; };
    }>[]) => setNodes((nodesSnapshot) =>
        applyNodeChanges(changes, nodesSnapshot)), []);
    const onEdgesChange =
        useCallback((changes: EdgeChange<{ id: string; source: string; target: string; }>[]) =>
            setEdges((edgesSnapshot) =>
                applyEdgeChanges(changes, edgesSnapshot)), []);
    const onConnect = useCallback((params: any) =>
        setEdges((edgesSnapshot) => addEdge(params, edgesSnapshot)), []);

    return (
        <div style={{ display: 'flex', width: '100vw', height: '100vh' }}>
            <Sidebar header={SIDEBAR_HEADER}>
                <SidebarTextField label="Name" value={name} onChange={setName} />
                <SidebarTextField label="Text" value={text} onChange={setText} />
                <SidebarArrayField label="Children" values={childrenIds} onChange={setChildrenIds} />
                <SidebarMaxField label="State changes" values={stateChanges} onChange={setStateChanges} />
            </Sidebar>
            <div style={{ flex: 1 }}>
                <ReactFlow
                    nodes={nodes}
                    edges={edges}
                    onNodesChange={onNodesChange}
                    onEdgesChange={onEdgesChange}
                    onConnect={onConnect}
                    fitView
                />
            </div>
        </div>
    );
}

export default StoryPage