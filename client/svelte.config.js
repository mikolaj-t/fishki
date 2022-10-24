import preprocess from 'svelte-preprocess';
import adapter from '@sveltejs/adapter-node';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter()
	},
	env: {
		dir: process.cwd(),
		publicPrefix: 'PUBLIC_' // lo
	},
	preprocess: [
		preprocess({
			postcss: true
		})
	]
};

export default config;
