<script>
  import Grid from "svelte-grid-responsive";

  export let xs = undefined;
  export let sm = undefined;
  export let md = undefined;
  export let lg = undefined;
  export let xl = undefined;
</script>

<Grid {xs} {sm} {md} {lg} {xl}>
  <slot />
</Grid>
