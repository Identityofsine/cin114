type Credit = {
	name: string;
	role: string;
}

export type VideoMetadata = {
	title: string;
	description?: string;
	credits?: Credit[];
	url: string;
	date: string;
	img?: string;
}
