/** @type {import("snowpack").SnowpackUserConfig } */
module.exports = {
  mount: {
	  pub: '/',
	  src: '/dist',
  },
  plugins: [
	  '@snowpack/plugin-svelte',
  ],
  install: [
    /* ... */
  ],
  installOptions: {
    /* ... */
  },
  devOptions: {
    /* ... */
  },
  buildOptions: {
    /* ... */
  },
  proxy: {
    /* ... */
  },
  alias: {
    /* ... */
  },
};
