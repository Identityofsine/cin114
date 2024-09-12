type Credit = {
	name: string;
	role: string;
}

export type VideoMetadata = {
	title: string;
	description?: string;
	boxart: {
		title: string;
		caption: string;
		img: string;
	};
	links?: {
		youtube?: string;
		vimeo?: string;
	}
	credits?: Credit[];
	url: string;
	date: string;
	img?: string;
}
