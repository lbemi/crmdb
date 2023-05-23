import { kubernetesInfo } from '/@/stores/kubernetes';

export function useWebsocketApi() {
	const k8sStore = kubernetesInfo();
	return {
		createWebsocket: (resource: string) => {
			const dns = import.meta.env.VITE_API_WEBSOCKET + k8sStore.state.activeCluster + '/' + resource;
			return new WebSocket(dns);
		},
		createLogWebsocket: (namespace: string, podName: string, containerName: string) => {
			const dns = import.meta.env.VITE_API_WEBSOCKET + `pod/${namespace}/${podName}/${containerName}/log?cloud=${k8sStore.state.activeCluster}`;
			return new WebSocket(dns);
		},
		createShellWebsocket: (namespace: string, podName: string, containerName: string) => {
			const dns = import.meta.env.VITE_API_WEBSOCKET + `pod/${namespace}/${podName}/${containerName}?cloud=${k8sStore.state.activeCluster}`;
			return new WebSocket(dns);
		},
	};
}
