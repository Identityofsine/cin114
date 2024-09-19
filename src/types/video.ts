type Credit = {
	name: string | string[];
	role: string;
}

export type VideoMetadata = {
	title: string;
	description?: string;
	useboxartaspreview?: boolean;
	boxart: {
		title: string;
		caption: string;
		img: string;
		video?: string;
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
