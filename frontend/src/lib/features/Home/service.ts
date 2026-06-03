import { apiFetch } from '../../services/apiClient.ts';

export interface Project {
	id: number;
	title: string;
	translation_key: string;
	repo_url?: string;
	live_url?: string;
	video_url?: string;
	created_at: string;
	featured: boolean;
	images: string[];       
	technologies: string[]; 
}

export const projectService = {
	
    // Get all projects from the API
	getAll: (): Promise<Project[]> => apiFetch<Project[]>('/projects'),

	// Get featured projects for the Hero or Main Section
	getFeatured: (): Promise<Project[]> => apiFetch<Project[]>('/projects/featured')
};