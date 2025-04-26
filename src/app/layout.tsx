import type { Metadata } from "next";
import "./globals.scss";
import { Navbar } from "@/components/Nav";
import Footer from "@/components/Footer";
import { TemplateMetadata } from "./template_metadata";
import { getBuildInfo } from "@/services/build";

export const metadata: Metadata = {
	...TemplateMetadata
};


export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {

	const build = getBuildInfo();

	return (
		<html lang="en">
			<body>
				<Navbar />
				{children}
				<Footer />
			</body>
		</html>
	);
}
