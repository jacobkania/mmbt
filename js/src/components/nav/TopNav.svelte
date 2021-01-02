<script>
  import Button from "@smui/button";

  import Logo from "app/Logo.svelte";
</script>

<style type="text/scss">
  @import "theme";

  .topNav {
    margin-bottom: 1rem;
    background-color: $primary;
  }

  .logo-container {
    height: 80px;
    width: 80px;
  }

  .link-container {
    padding: 0 1rem;
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: space-between;
    max-width: 40rem;
  }

  .link {
    color: $white;
    text-decoration: none;
    font-family: "titillium";
    letter-spacing: 1px;
  }
</style>

<div class="topNav">
  <div class="link-container">
    <span class="logo-container">
      <a href="/">
        <Logo />
      </a>
    </span>
    <a class="link" href="/">Home</a>
    <a class="link" href="/overview">Overview</a>
    <a class="link" href="/snapshot/new">New snapshot</a>
  </div>
</div>
