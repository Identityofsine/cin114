'use client';
import { BrandSettings } from "@/brand.settings";
import React, { useEffect, useState } from "react"

export function TextChange() {

	const [text, setText] = useState<string>('Contact Us')
	const ref = React.useRef<HTMLDivElement>(null);
	const oref = React.useRef<IntersectionObserver>();
	const iref = React.useRef<Promise<any>>();

	useEffect(() => {
		if (oref.current) {
			return;
		}
		oref.current = new IntersectionObserver((entries) => {
			entries.forEach(entry => {
				(async () => {
					if (entry.isIntersecting) {
						if (iref.current) {
							return;
						}
						iref.current = (async () => {
							const text = BrandSettings.contact.email;
							let cancel_ = false;
							for (let i = 0; i < text.length; i++) {
								if (cancel_) {
									return;
								}
								await new Promise(r => setTimeout(r, 133));
								setText(text.substring(0, i + 1));
							}
							iref.current = undefined;
							return {
								cancel() {
									cancel_ = true;
									iref.current = undefined;
									setText(text);
								}
							}
						})();
					} else {
						if (iref.current) {
							Promise.resolve(iref.current)
						} else {
							setText('Contact Us');
						}
					}
				})();
			})

		},
			{ root: null, rootMargin: '2px', threshold: 0.0 });
		oref.current.observe(ref.current!);
		() => {
			iref.current?.then(x => x.cancel);
		}
	}, [])

	return (
		<h1 ref={ref}>{text}</h1>
	)
}
