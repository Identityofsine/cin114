'use client';

import React, { useMemo } from 'react';
import './styles/navbar.scss';
import { BuildInfo } from '@/services/build';
import { ServerHealth } from '@/types/health';
import { getHealth } from '@/api';
import { Tooltip } from 'react-tooltip';

type NavbarProps = {
  buildInfo: BuildInfo;
}

export function Navbar({ buildInfo }: NavbarProps) {

  const [curPage, setCurPage] = React.useState<string>('');
  const [open, setOpen] = React.useState<boolean>(false);
  const [serverHealth, setServerHealth] = React.useState<ServerHealth | undefined>();

  React.useEffect(() => {
    if (!window.document.title.match('-'))
      setCurPage(window.document.title);

    getHealth().then((health) => health && setServerHealth(health))

  }, []);

  const serverDetails = useMemo(() => {
    if (!serverHealth) return ['No server health data available'];
    return `Server: ${serverHealth.server_name || 'unknown'}\n
            Build Date: ${serverHealth.build_date || 'unknown'}\n
            Version: ${serverHealth.version || 'unknown'}\n
            Branch: ${serverHealth.branch || 'unknown'} \n
            Commit: ${serverHealth.commit || 'unknown'}\n
            Environment: ${serverHealth.environment || 'unknown'}`.split('\n');
  }, [serverHealth]);

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
      <div
        id="navmenu"
        className={`navmenu ${open && 'open' || ''}`}
      >
        <div className="navmenu__container">
          <Link href="/">Home</Link>
          <Link href="/#about">About</Link>
          <Link href="/#contact">Contact</Link>
        </div>
      </div>
      <div
        className={`nav flex align-center ${open && 'open' || ''}`}
        onClick={() => { setOpen(!open) }}
      >
        <img
          src="/ui/navthing.svg"
          alt="navbar"
          draggable={false} />
        <span>{curPage}</span>
        {buildInfo.branch !== 'prod' && (
          <span
            data-tooltip-id="server-health"
            className="build"
          >
            {buildInfo.buildDate}
            {buildInfo.branch && ` (${buildInfo.branch})`}
            {buildInfo.buildId && ` - ${buildInfo.buildId}`}
          </span>
        )}
        <Tooltip
          id="server-health"
          place="bottom"
          className="tooltip"
        >
          <div className="flex column">
            {serverDetails.map((line, index) => (
              <span key={index}>{line}</span>
            ))}
          </div>
        </Tooltip>
      </div>
    </>
  )
}
