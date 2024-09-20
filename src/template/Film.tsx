import { VideoMetadata } from "@/types/video";
import './styles/film.scss';
import { FilmBackground, FilmCredit } from "./FilmClient";

type FilmTemplateProps = {
	children?: React.ReactNode | React.ReactNode[];
	metadata: VideoMetadata;
}

function FilmTemplate({ children, metadata }: FilmTemplateProps) {
	return (
		<div className="film">
			{children}
			<FilmBackground metadata={metadata} />
			<div className="film__static">
				<div className="film__static__meta">
					<h2 className="film__static__meta__title">
						{metadata.boxart.title}
					</h2>
					<p className="film__static__meta__description">
						{metadata.boxart.caption}
					</p>
				</div>
				<FilmCredit metadata={metadata} />
			</div>
		</div >
	)
}

export default FilmTemplate;
