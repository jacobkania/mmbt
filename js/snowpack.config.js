/** @type {import("snowpack").SnowpackUserConfig } */
module.exports = {
  mount: {
	  pub: '/',
	  src: '/dist',
  },
  plugins: [
    '@snowpack/plugin-svelte',
    [
      '@snowpack/plugin-webpack',
      {
      }
    ]
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
    out: "."
  },
  proxy: {
    /* ... */
  },
  alias: {
    /* ... */
  },
};
