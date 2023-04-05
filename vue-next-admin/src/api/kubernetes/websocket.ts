import { kubernetesInfo } from '/@/stores/kubernetes';

export function useWebsocketApi() {
	const k8sStore = kubernetesInfo();
	return {
		createWebsocket: (resource: string) => {
			const dns = import.meta.env.VITE_API_WEBSOCKET + k8sStore.state.activeCluster + '/' + resource;
			return   new WebSocket(dns);

		},
	};
}
