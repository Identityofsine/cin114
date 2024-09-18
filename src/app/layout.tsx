import type { Metadata } from "next";
import "./globals.scss";
import { Navbar } from "@/components/Nav";
import Footer from "@/components/Footer";
import { TemplateMetadata } from "./template_metadata";

export const metadata: Metadata = {
	...TemplateMetadata
};


export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
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
