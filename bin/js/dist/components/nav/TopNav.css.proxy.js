// [snowpack] add styles to the page (skip if no document exists)
if (typeof document !== 'undefined') {
  const code = ".topNav.svelte-w13mqq{padding:1rem;margin-bottom:1rem}.link-container.svelte-w13mqq{display:flex;align-items:center;flex-direction:row;justify-content:space-between;width:20rem}.link.svelte-w13mqq{background-color:green;color:white}";

  const styleEl = document.createElement("style");
  const codeEl = document.createTextNode(code);
  styleEl.type = 'text/css';
  styleEl.appendChild(codeEl);
  document.head.appendChild(styleEl);
}