<script>
  import Grid from "svelte-grid-responsive";

  export let xs;
  export let sm;
  export let md;
  export let lg;
  export let xl;
</script>

<Grid {xs} {sm} {md} {lg} {xl}>
  <slot />
</Grid>
