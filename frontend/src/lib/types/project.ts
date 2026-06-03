export interface ProjectImage {
	id: number;
	image_url: string;
}

export interface ProjectTechnology {
	id: number;
	name: string;
	icon_slug: string;
}

export interface Project {
	id: number;
	title: string;
	translation_key: string;
	repo_url?: string;
	live_url?: string;
	video_url?: string;
	created_at: string;
	featured: boolean;
	images: ProjectImage[];
	technologies: ProjectTechnology[];
}
