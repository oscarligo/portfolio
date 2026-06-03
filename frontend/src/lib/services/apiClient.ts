const BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080/api';

export async function apiFetch<T>(endpoint: string): Promise<T> {
	const url = `${BASE_URL}${endpoint}`;

	const response = await fetch(url, {
		method: 'GET',
		headers: {
			'Accept': 'application/json'
		}
	});

	if (!response.ok) {
		throw new Error(`[API Error ${response.status}]: ${response.statusText}`);
	}

	return response.json();
}
