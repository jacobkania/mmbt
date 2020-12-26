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
    port: 3000,
    output: "stream",
    open: "none"
  },
  buildOptions: {
    out: "../bin/js"
  },
  proxy: {
    /* ... */
  },
  alias: {
    /* ... */
  },
};
