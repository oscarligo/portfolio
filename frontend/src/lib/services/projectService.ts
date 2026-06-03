import type { Project, ProjectImage, ProjectTechnology } from '../types/project.js';
import { apiFetch } from './apiClient.ts';

type RawProject = Omit<Project, 'images' | 'technologies'> & {
	repo_url?: string | null;
	live_url?: string | null;
	video_url?: string | null;
	images?: unknown;
	technologies?: unknown;
};

function isRecord(value: unknown): value is Record<string, unknown> {
	return typeof value === 'object' && value !== null;
}

function normalizeImages(value: unknown): ProjectImage[] {
	if (!Array.isArray(value)) {
		return [];
	}

	return value.flatMap((image) => {
		if (!isRecord(image) || typeof image.image_url !== 'string') {
			return [];
		}

		return [
			{
				id: Number(image.id) || 0,
				image_url: image.image_url
			}
		];
	});
}

function normalizeTechnologies(value: unknown): ProjectTechnology[] {
	if (!Array.isArray(value)) {
		return [];
	}

	return value.flatMap((technology) => {
		if (!isRecord(technology) || typeof technology.name !== 'string') {
			return [];
		}

		return [
			{
				id: Number(technology.id) || 0,
				name: technology.name,
				icon_slug: typeof technology.icon_slug === 'string' ? technology.icon_slug : ''
			}
		];
	});
}

function normalizeProject(project: RawProject): Project {
	return {
		id: Number(project.id),
		title: project.title,
		translation_key: project.translation_key,
		repo_url: project.repo_url ?? undefined,
		live_url: project.live_url ?? undefined,
		video_url: project.video_url ?? undefined,
		created_at: project.created_at,
		featured: Boolean(project.featured),
		images: normalizeImages(project.images),
		technologies: normalizeTechnologies(project.technologies)
	};
}

async function fetchProjects(endpoint: string): Promise<Project[]> {
	const response = await apiFetch<RawProject[]>(endpoint);
	return response.map(normalizeProject);
}

export const projectService = {
	getAll: (): Promise<Project[]> => fetchProjects('/projects'),
	getFeatured: (): Promise<Project[]> => fetchProjects('/projects/featured')
};
