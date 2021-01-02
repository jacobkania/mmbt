<script>
  import Grid from "svelte-grid-responsive";

  export let pad = 2;

  const padMultiplier = 8;
</script>

<Grid container gutter={pad * padMultiplier}>
  <slot />
</Grid>
