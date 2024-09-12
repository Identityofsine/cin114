import { VideoMetadata } from "@/types/video";
import './styles/film.scss';

type FilmTemplateProps = {
	children: React.ReactNode | React.ReactNode[];
	metadata: VideoMetadata;
}

function FilmTemplate() {
	return (
		<div className="film">
		</div>
	)
}

export default FilmTemplate;
