//@FILEDESC FilmTemplate is a template component that renders the film's metadata and background, it allows for children to be passed in to render the film's content (usually a logo).
import { VideoMetadata } from "@/types/video";
import './styles/film.scss';
import { FilmBackground, FilmCredit } from "./FilmClient";

type FilmTemplateProps = {
	children?: React.ReactNode | React.ReactNode[];
	metadata: VideoMetadata;
}

//@BLOCK
//@TITLE FilmTemplate
//@DESC FilmTemplate is a template component that renders the film's metadata and background, it allows for children to be passed in to render the film's content (usually a logo).
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
//@END

export default FilmTemplate;
