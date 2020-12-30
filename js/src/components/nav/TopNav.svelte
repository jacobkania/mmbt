<script>
  import { Link } from "svelte-navigator";
  import Button from "@smui/button";
</script>

<style>
  .topNav {
    padding: 1rem;
    margin-bottom: 1rem;
  }

  .link-container {
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: space-between;
    width: 20rem;
  }
</style>

<div class="topNav">
  <div class="link-container">
    <Link to="/">
      <Button>Home</Button>
    </Link>
    <Link to="/overview">
      <Button>Overview</Button>
    </Link>
  </div>
</div>
