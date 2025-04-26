export interface BuildInfo {
	branch: 'prod' | 'dev'; // Branch name (prod or dev)
	buildId: string; // Unique identifier for the build
	buildDate: string;
}

export function getBuildInfo(): BuildInfo {
	//this needs to be set in the build pipeline (e.g. GitHub Actions, etc.)
	const buildId = process.env.BUILD_ID || 'local';
	const buildDate = new Date(process.env.BUILD_DATE || Date.now()).toISOString();
	const branch = process.env.BRANCH_NAME === 'main' ? 'prod' : 'dev';

	return {
		branch,
		buildId,
		buildDate,
	};
}
