'use client';

import React from 'react';
import './styles/navbar.scss';


export function Navbar() {

	const [curPage, setCurPage] = React.useState<string>('');
	const [open, setOpen] = React.useState<boolean>(false);

	React.useEffect(() => {
		if (!window.document.title.match('-'))
			setCurPage(window.document.title);
	}, []);

	function close(page: string) {
		const current_page = window.location.pathname;
		const raw = current_page.split('#');
		const current = raw[0];
		if (current === page.split('#')[0]) {
			setOpen(false);
		}
	}

	function Link({ href, children }: { href: string, children: React.ReactNode }) {
		return (
			<a href={href} onClick={() => { close(href) }}>{children}</a>
		)
	}

	return (
		<>
			<div id="navmenu" className={`navmenu ${open && 'open' || ''}`}>
				<div className="navmenu__container">
					<Link href="/">Home</Link>
					<Link href="/#about">About</Link>
					<Link href="/#contact">Contact</Link>
				</div>
			</div>
			<div className={`nav flex align-center ${open && 'open' || ''}`} onClick={() => { setOpen(!open) }}>
				<img src="/ui/navthing.svg" alt="navbar" draggable={false} />
				<span>{curPage}</span>
			</div>
		</>
	)
}
