
// Base url from .env file
const BASE_URL = import.meta.env.VITE_API_BASE_URL;

/**
 * Custom options for apiFetch.
 */
interface FetchOptions extends RequestInit {
    
    json?: Record<string, any>; 
}

/**
 * Helper function to make API requests with error handling and JSON parsing.
 */
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