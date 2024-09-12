import { VideoMetadata } from "@/types/video";
import './styles/film.scss';

type FilmTemplateProps = {
	children?: React.ReactNode | React.ReactNode[];
	metadata: VideoMetadata;
}

function FilmTemplate({ children, metadata }: FilmTemplateProps) {
	return (
		<div className="film">
			{children}
			<div className="film__bg" style={{ backgroundImage: `url("${metadata.boxart.img}")` }} />
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
									<p>
										{credit.name}
									</p>
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
