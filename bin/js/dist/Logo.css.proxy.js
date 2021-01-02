// [snowpack] add styles to the page (skip if no document exists)
if (typeof document !== 'undefined') {
  const code = ".background.svelte-1ce9b86{fill:#0e34a0ff}.icon.svelte-1ce9b86{fill:#63a375ff}.text.svelte-1ce9b86{fill:#ffffff}";

  const styleEl = document.createElement("style");
  const codeEl = document.createTextNode(code);
  styleEl.type = 'text/css';
  styleEl.appendChild(codeEl);
  document.head.appendChild(styleEl);
}