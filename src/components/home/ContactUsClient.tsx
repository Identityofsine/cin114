'use client';
import { BrandSettings } from "@/brand.settings";
import React, { useEffect, useState } from "react"

export function TextChange() {

	const [text, setText] = useState<string>('Contact Us')
	const ref = React.useRef<HTMLDivElement>(null);
	const oref = React.useRef<IntersectionObserver>();
	const iref = React.useRef<{ cancel: () => void }>();

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
						iref.current = (() => {
							const text_ = BrandSettings.contact.email;
							const original = "" + text; //avoid reference
							let cancel_ = false;

							const generic_cancel = () => {
								cancel_ = true;
								iref.current = undefined;
								setText(original);
							}

							let timeout = setTimeout(async () => {
								await new Promise(r => setTimeout(r, 225));
								if (cancel_) {
									generic_cancel();
									return { cancel: () => { generic_cancel() } };
								}
								for (let i = 0; i < original.length; i++) {
									if (cancel_) {
										generic_cancel();
										return { cancel: () => { generic_cancel() } };
									}
									await new Promise(r => setTimeout(r, 225));
									let newText = original.substring(0, original.length - i - 1);
									if (newText === "") newText = "   \n";
									setText(newText);
								}
								for (let i = 0; i < text_.length; i++) {
									if (cancel_) {
										generic_cancel();
										return { cancel: () => { generic_cancel() } };
									}
									await new Promise(r => setTimeout(r, 225));
									setText(text_.substring(0, i + 1));
								}
								iref.current = undefined;
							}, 2150);
							return {
								cancel() {
									setText(original);
									clearTimeout(timeout);
									cancel_ = true;
									iref.current = undefined;
								}
							}
						})();
					} else {
						if (iref.current) {
							iref.current.cancel();
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
			iref.current!.cancel();
		}
	}, [])

	return (
		<h1 ref={ref}>{text}</h1>
	)
}
