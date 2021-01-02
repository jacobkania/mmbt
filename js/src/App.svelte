<script>
  import { Router, Link, Route } from "svelte-navigator";
  import Home from "pages/home/Home.svelte";
  import Overview from "pages/overview/Overview.svelte";
  import TopNav from "components/nav/TopNav.svelte";
</script>

<style type="text/scss">
  @import "theme";

  .main {
    height: 100vh;
    margin: 0 auto;
    color: $text;
  }
</style>

<Router>
  <div class="main" id="main">
    <TopNav />
    <div>
      <Route path="/">
        <Home />
      </Route>
      <Route path="overview">
        <Overview />
      </Route>
    </div>
  </div>
</Router>
