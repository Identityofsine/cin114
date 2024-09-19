import { VideoMetadata } from "@/types/video";
import './styles/film.scss';
import { FilmBackground } from "./FilmClient";
import { isArray } from "util";

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
				<div className="film__credits">
					<div className="film__credits__container">
						{metadata?.credits && metadata.credits.map((credit, index) => (
							<div className="film__credits__credit" key={credit.role + credit.name + index}>
								<div className="film__credits__credit__title" key={credit.role + index}>
									<p>
										{credit.role}
									</p>
								</div>
								<div className="film__credits__credit__name">
									{!Array.isArray(credit.name) ? <p>{credit.name}</p> :
										credit.name.map((name, index) => (
											<p key={name + index}>{name}</p>
										))
									}
								</div>
							</div>
						))}
					</div>
				</div>
			</div>
		</div >
	)
}

export default FilmTemplate;
