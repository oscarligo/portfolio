import { describe, it, expect, vi } from 'vitest';
import { apiFetch } from './apiClient.ts';

describe('apiClient', () => {
	it('should fetch data from the API', async () => {
		const mockData = [{ id: 1, title: 'ClassControl' }];
		
		const fetchMock = vi.fn().mockResolvedValue({
			ok: true,
			status: 200,
			json: async () => mockData
		});
		vi.stubGlobal('fetch', fetchMock);

		const result = await apiFetch('/projects');

		expect(fetchMock).toHaveBeenCalledWith('http://localhost:8080/api/projects', expect.any(Object));
		expect(result).toEqual(mockData);
	});
});