import { displayDate } from "@/util/date";

export interface BuildInfo {
	branch: 'prod' | 'dev'; // Branch name (prod or dev)
	buildId: string; // Unique identifier for the build
	buildDate: string;
}

export function getBuildInfo(): BuildInfo {
	//this needs to be set in the build pipeline (e.g. GitHub Actions, etc.)
	const buildId = process.env.BUILD_ID || '';
	const buildDate = displayDate(
		new Date(process.env.BUILD_DATE || Date.now())
	)
	const branch = process.env.NEXT_PUBLIC_BRANCH || 'dev';
	if (branch !== 'prod' && branch !== 'dev') {
		throw new Error(`Invalid branch name: ${branch}`);
	}

	return {
		branch,
		buildId,
		buildDate,
	};
}
