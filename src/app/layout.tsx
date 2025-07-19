import type { Metadata } from "next";
import "./globals.scss";
import { Navbar } from "@/components/Nav";
import Footer from "@/components/Footer";
import { TemplateMetadata } from "./template_metadata";
import { getBuildInfo } from "@/services/build";

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
      </head>
      <body>
        <Navbar
          buildInfo={build}
        />
        {children}
        <Footer />
      </body>
    </html>
  );
}
