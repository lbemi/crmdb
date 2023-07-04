import { kubernetesInfo } from '@/stores/kubernetes';
import { Session } from '@/utils/storage';

export function useWebsocketApi() {
	const k8sStore = kubernetesInfo();
	const token = Session.get('token');
	return {
		createWebsocket: (resource: string) => {
			const dns = import.meta.env.VITE_API_WEBSOCKET + 'ws/' + k8sStore.state.activeCluster + '/' + resource;
			const ws = new WebSocket(dns, [token]);
			return ws;
		},
		createLogWebsocket: (namespace: string, podName: string, containerName: string) => {
			const dns =
				import.meta.env.VITE_API_WEBSOCKET + `pods/namespaces/${namespace}/logs/${podName}/${containerName}?cloud=${k8sStore.state.activeCluster}`;
			const ws = new WebSocket(dns, [token]);
			return ws;
		},
		createShellWebsocket: (namespace: string, podName: string, containerName: string) => {
			const dns =
				import.meta.env.VITE_API_WEBSOCKET + `pods/namespaces/${namespace}/exec/${podName}/${containerName}?cloud=${k8sStore.state.activeCluster}`;
			const ws = new WebSocket(dns, [token]);
			return ws;
		},
		createTestShellWebsocket: () => {
			const dns = 'ws://127.0.0.1:8080/api/v1/hosts/1/shell/1';
			const ws = new WebSocket(dns, [token]);
			return ws;
		},
	};
}
