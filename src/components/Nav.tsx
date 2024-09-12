'use client';

import React from 'react';
import './styles/navbar.scss';


export function Navbar() {

	const [curPage, setCurPage] = React.useState<string>('');

	React.useEffect(() => {
		setCurPage(window.document.title);
	}, []);

	return (
		<div className="nav flex align-center">
			<img src="/ui/navthing.svg" alt="navbar" draggable={false} />
			<span>{curPage}</span>
		</div>
	)
}
