import type { Metadata } from "next";
import "./globals.scss";
import { Navbar } from "@/components/Nav";
import Footer from "@/components/Footer";
import { TemplateMetadata } from "./template_metadata";
import { getBuildInfo } from "@/services/build";
import dynamic from "next/dynamic";

// Dynamically import the client component with SSR disabled
const EventCallToActionClient = dynamic(
  () => import("@/components/home/EventCallToActionClient"),
  { ssr: false }
);

export const metadata: Metadata = {
  ...TemplateMetadata
};


export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  const build = getBuildInfo();

  return (
    <html lang="en">
      <head>
        <link href="https://fonts.googleapis.com/css2?family=Rubik:wght@700;900&display=swap" rel="stylesheet" />
        <link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
          integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
          crossOrigin="" />
      </head>
      <body>
        <div id="event-banner-root"></div>
        <EventCallToActionClient />
        <Navbar
          buildInfo={build}
        />
        {children}
        <Footer />
      </body>
    </html>
  );
}
