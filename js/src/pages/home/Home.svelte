<script>
  import Card from "@smui/card";
  import Button from "@smui/button";

  let clicked = 0;
</script>

<style>
</style>

<Button on:click={() => clicked++} variant="raised" color="primary">
  Click me
</Button>
<div>Clicks: {clicked}</div>
<Card>Fancy card clicks: {clicked}</Card>
