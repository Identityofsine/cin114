import { NextApiRequest, NextApiResponse } from "next";

const bc = new (class BuildController {

	//nextjs styled endpoint
	public async getBuildInfo(_: NextApiRequest, res: NextApiResponse) {
		const buildId = process.env.BUILD_ID || 'local';
		const buildDate = new Date(process.env.BUILD_DATE || Date.now()).toISOString();
		const branch = process.env.BRANCH_NAME === 'main' ? 'prod' : 'dev';

		res.status(200).json({
			branch,
			buildId,
			buildDate,
		});
	}

})

export default bc;
