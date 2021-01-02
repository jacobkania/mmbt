// [snowpack] add styles to the page (skip if no document exists)
if (typeof document !== 'undefined') {
  const code = ".topNav.svelte-1b8k0pz{margin-bottom:1rem;background-color:#0e34a0ff}.logo-container.svelte-1b8k0pz{height:80px;width:80px}.link-container.svelte-1b8k0pz{padding:0 1rem;display:flex;align-items:center;flex-direction:row;justify-content:space-between;max-width:40rem}.link.svelte-1b8k0pz{color:#ffffff;text-decoration:none;font-family:\"titillium\";letter-spacing:1px}";

  const styleEl = document.createElement("style");
  const codeEl = document.createTextNode(code);
  styleEl.type = 'text/css';
  styleEl.appendChild(codeEl);
  document.head.appendChild(styleEl);
}