<script>
  import { Link } from "svelte-navigator";
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

  .link {
    background-color: green;
    color: white;
  }
</style>

<div class="topNav">
  <div class="link-container">
    <Link to="/"><span class="link">Home</span></Link>
    <Link to="/overview"><span class="link">Overview</span></Link>
  </div>
</div>
