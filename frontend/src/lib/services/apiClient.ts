
// Base url from .env file
const BASE_URL = import.meta.env.VITE_API_BASE_URL;

/**
 * Custom options for apiFetch.
 */
interface FetchOptions extends RequestInit {
    
    json?: Record<string, any>; 
}

/**
 * apiFetch es un wrapper seguro sobre el fetch nativo del navegador.
 * Se encarga de inyectar cabeceras, manejar tokens de sesión y capturar errores.
 */
export async function apiFetch<T>(endpoint: string, options: FetchOptions = {}): Promise<T> {
    const url = `${BASE_URL}${endpoint}`;
    
    // 1. Inicializar cabeceras heredadas o vacías
    const headers = new Headers(options.headers);

    // 2. Extraer el token JWT de localStorage e inyectarlo si existe
    const token = localStorage.getItem('token');
    if (token) {
        headers.set('Authorization', `Bearer ${token}`);
    }

    // 3. Configurar el cuerpo de la petición (JSON vs Multipart/FormData)
    let body = options.body;

    if (options.json) {
        // Si enviamos JSON plano, aseguramos la cabecera correspondiente
        headers.set('Content-Type', 'application/json');
        body = JSON.stringify(options.json);
    } 
    // NOTA: Si pasas un FormData (para R2), NO debes setear 'Content-Type'.
    // El navegador lo detecta automáticamente e inyecta el "boundary" del multipart.

    // 4. Ejecutar la petición nativa
    const response = await fetch(url, {
        ...options,
        headers,
        body
    });

    // 5. Manejo defensivo de errores globales (4xx y 5xx)
    if (!response.ok) {
        // Si tu backend de Go devuelve un error formateado en texto plano o JSON, lo capturamos
        const errorText = await response.text().catch(() => 'Unknown error');
        
        // Manejo específico si el token expiró o es inválido
        if (response.status === 401) {
            localStorage.removeItem('token'); 
        }

        throw new Error(`[API Error ${response.status}]: ${errorText || response.statusText}`);
    }

	// 6. Manejo correcto para respuestas vacías (como tu DELETE que devuelve 204 No Content)
	if (response.status === 204) {
		return null as T;
	}

	return response.json();
}