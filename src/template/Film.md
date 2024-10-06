# Film.tsx
FilmTemplate is a template component that renders the film's metadata and background, it allows for children to be passed in to render the film's content (usually a logo).
#### FilmTemplate
FilmTemplate is a template component that renders the film's metadata and background, it allows for children to be passed in to render the film's content (usually a logo).
```tsx
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
```
